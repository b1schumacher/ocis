package http

import (
	"fmt"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/owncloud/ocis/v2/ocis-pkg/config"
	ocismiddleware "github.com/owncloud/ocis/v2/ocis-pkg/middleware"
	"github.com/owncloud/ocis/v2/ocis-pkg/service/http"
	"github.com/owncloud/ocis/v2/ocis-pkg/version"
	svc "github.com/owncloud/ocis/v2/services/thumbnails/pkg/service/http/v0"
	"github.com/owncloud/ocis/v2/services/thumbnails/pkg/thumbnail/storage"
	"go-micro.dev/v4"
)

// Server initializes the http service and server.
func Server(opts ...Option) (http.Service, error) {
	options := newOptions(opts...)

	tlsConfig, err := config.BuildTLSConfig(
		options.Logger,
		options.Config.HTTP.TLS.Enabled,
		options.Config.HTTP.TLS.Cert,
		options.Config.HTTP.TLS.Key,
		options.Config.HTTP.Addr,
	)
	if err != nil {
		options.Logger.Error().
			Err(err).
			Msg("could not build certificate")
		return http.Service{}, fmt.Errorf("could not build certificate: %w", err)
	}

	service, err := http.NewService(
		http.TLSConfig(tlsConfig),
		http.Logger(options.Logger),
		http.Name(options.Config.Service.Name),
		http.Version(version.GetString()),
		http.Namespace(options.Config.HTTP.Namespace),
		http.Address(options.Config.HTTP.Addr),
		http.Context(options.Context),
	)
	if err != nil {
		options.Logger.Error().
			Err(err).
			Msg("Error initializing http service")
		return http.Service{}, fmt.Errorf("could not initialize http service: %w", err)
	}

	handle := svc.NewService(
		svc.Logger(options.Logger),
		svc.Config(options.Config),
		svc.Middleware(
			middleware.RealIP,
			middleware.RequestID,
			// ocismiddleware.Secure,
			ocismiddleware.Version(
				options.Config.Service.Name,
				version.GetString(),
			),
			ocismiddleware.Logger(options.Logger),
		),
		svc.ThumbnailStorage(
			storage.NewFileSystemStorage(
				options.Config.Thumbnail.FileSystemStorage,
				options.Logger,
			),
		),
	)

	{
		handle = svc.NewInstrument(handle, options.Metrics)
		handle = svc.NewLogging(handle, options.Logger)
		handle = svc.NewTracing(handle)
	}

	if err := micro.RegisterHandler(service.Server(), handle); err != nil {
		return http.Service{}, err
	}

	return service, nil
}
