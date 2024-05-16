package command

import (
	"context"
	"crypto/tls"
	"fmt"
	"os/signal"

	"github.com/owncloud/ocis/v2/ocis-pkg/config/configlog"
	pkgcrypto "github.com/owncloud/ocis/v2/ocis-pkg/crypto"
	"github.com/owncloud/ocis/v2/ocis-pkg/handlers"
	"github.com/owncloud/ocis/v2/ocis-pkg/runner"
	"github.com/owncloud/ocis/v2/ocis-pkg/service/debug"
	"github.com/owncloud/ocis/v2/ocis-pkg/version"
	"github.com/owncloud/ocis/v2/services/nats/pkg/config"
	"github.com/owncloud/ocis/v2/services/nats/pkg/config/parser"
	"github.com/owncloud/ocis/v2/services/nats/pkg/logging"
	"github.com/owncloud/ocis/v2/services/nats/pkg/server/nats"
	"github.com/urfave/cli/v2"
)

// Server is the entrypoint for the server command.
func Server(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:     "server",
		Usage:    fmt.Sprintf("start the %s service without runtime (unsupervised mode)", cfg.Service.Name),
		Category: "server",
		Before: func(c *cli.Context) error {
			return configlog.ReturnFatal(parser.ParseConfig(cfg))
		},
		Action: func(c *cli.Context) error {
			logger := logging.Configure(cfg.Service.Name, cfg.Log)

			var cancel context.CancelFunc
			ctx := cfg.Context
			if ctx == nil {
				ctx, cancel = signal.NotifyContext(context.Background(), runner.StopSignals...)
				defer cancel()
			}

			gr := runner.NewGroup()
			{
				server := debug.NewService(
					debug.Logger(logger),
					debug.Name(cfg.Service.Name),
					debug.Version(version.GetString()),
					debug.Address(cfg.Debug.Addr),
					debug.Token(cfg.Debug.Token),
					debug.Pprof(cfg.Debug.Pprof),
					debug.Zpages(cfg.Debug.Zpages),
					debug.Health(handlers.Health),
					debug.Ready(handlers.Ready),
				)

				gr.Add(runner.NewGolangHttpServerRunner("nats_debug", server))
			}

			var tlsConf *tls.Config
			if cfg.Nats.EnableTLS {
				// Generate a self-signing cert if no certificate is present
				if err := pkgcrypto.GenCert(cfg.Nats.TLSCert, cfg.Nats.TLSKey, logger); err != nil {
					logger.Fatal().Err(err).Msgf("Could not generate test-certificate")
				}

				crt, err := tls.LoadX509KeyPair(cfg.Nats.TLSCert, cfg.Nats.TLSKey)
				if err != nil {
					return err
				}

				clientAuth := tls.RequireAndVerifyClientCert
				if cfg.Nats.TLSSkipVerifyClientCert {
					clientAuth = tls.NoClientCert
				}

				tlsConf = &tls.Config{
					MinVersion:   tls.VersionTLS12,
					ClientAuth:   clientAuth,
					Certificates: []tls.Certificate{crt},
				}
			}
			natsServer, err := nats.NewNATSServer(
				logging.NewLogWrapper(logger),
				nats.Host(cfg.Nats.Host),
				nats.Port(cfg.Nats.Port),
				nats.ClusterID(cfg.Nats.ClusterID),
				nats.StoreDir(cfg.Nats.StoreDir),
				nats.TLSConfig(tlsConf),
				nats.AllowNonTLS(!cfg.Nats.EnableTLS),
			)
			if err != nil {
				return err
			}

			gr.Add(runner.New("nats_svc", func() error {
				return natsServer.ListenAndServe()
			}, func() {
				natsServer.Shutdown()
			}))

			grResults := gr.Run(ctx)

			// return the first non-nil error found in the results
			for _, grResult := range grResults {
				if grResult.RunnerError != nil {
					return grResult.RunnerError
				}
			}
			return nil
		},
	}
}
