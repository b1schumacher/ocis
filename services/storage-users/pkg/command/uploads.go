package command

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	tw "github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"

	userpb "github.com/cs3org/go-cs3apis/cs3/identity/user/v1beta1"
	"github.com/cs3org/reva/v2/pkg/events"
	"github.com/cs3org/reva/v2/pkg/storage"
	"github.com/cs3org/reva/v2/pkg/storage/fs/registry"
	"github.com/cs3org/reva/v2/pkg/utils"
	"github.com/owncloud/ocis/v2/ocis-pkg/config/configlog"
	"github.com/owncloud/ocis/v2/services/storage-users/pkg/config"
	"github.com/owncloud/ocis/v2/services/storage-users/pkg/config/parser"
	"github.com/owncloud/ocis/v2/services/storage-users/pkg/event"
	"github.com/owncloud/ocis/v2/services/storage-users/pkg/revaconfig"
)

// Session contains the information of an upload session
type Session struct {
	ID         string         `json:"id"`
	Space      string         `json:"space"`
	Filename   string         `json:"filename"`
	Offset     int64          `json:"offset"`
	Size       int64          `json:"size"`
	Executant  userpb.UserId  `json:"executant"`
	SpaceOwner *userpb.UserId `json:"spaceowner,omitempty"`
	Expires    time.Time      `json:"expires"`
	Processing bool           `json:"processing"`
	ScanDate   time.Time      `json:"virus_scan_date"`
	ScanResult string         `json:"virus_scan_result"`
}

// Uploads is the entry point for the uploads command
func Uploads(cfg *config.Config) *cli.Command {
	return &cli.Command{

		Name:  "uploads",
		Usage: "manage unfinished uploads",
		Subcommands: []*cli.Command{
			ListUploads(cfg),
			ListUploadSessions(cfg),
			PurgeExpiredUploads(cfg),
		},
	}
}

// ListUploads prints a list of all incomplete uploads
func ListUploads(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "Print a list of all incomplete uploads (deprecated, use sessions)",
		Before: func(c *cli.Context) error {
			return configlog.ReturnFatal(parser.ParseConfig(cfg))
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Warning: This command is deprecated, use 'uploads sessions' instead")

			f, ok := registry.NewFuncs[cfg.Driver]
			if !ok {
				fmt.Fprintf(os.Stderr, "Unknown filesystem driver '%s'\n", cfg.Driver)
				os.Exit(1)
			}
			drivers := revaconfig.StorageProviderDrivers(cfg)
			fs, err := f(drivers[cfg.Driver].(map[string]interface{}), nil)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to initialize filesystem driver '%s'\n", cfg.Driver)
				return err
			}

			managingFS, ok := fs.(storage.UploadSessionLister)
			if !ok {
				fmt.Fprintf(os.Stderr, "'%s' storage does not support listing upload sessions\n", cfg.Driver)
				os.Exit(1)
			}
			expired := false
			uploads, err := managingFS.ListUploadSessions(c.Context, storage.UploadSessionFilter{Expired: &expired})
			if err != nil {
				return err
			}

			fmt.Println("Incomplete uploads:")
			for _, u := range uploads {
				ref := u.Reference()
				fmt.Printf(" - %s (Space: %s, Name: %s, Size: %d/%d, Expires: %s, Processing: %t)\n", u.ID(), ref.GetResourceId().GetSpaceId(), u.Filename(), u.Offset(), u.Size(), u.Expires(), u.IsProcessing())
			}
			return nil
		},
	}
}

// ListUploadSessions prints a list of upload sessiens
func ListUploadSessions(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:  "sessions",
		Usage: "Print a list of upload sessions",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "id",
				DefaultText: "unset",
				Usage:       "filter sessions by upload session id",
			},
			&cli.BoolFlag{
				Name:        "processing",
				DefaultText: "unset",
				Usage:       "filter sessions by processing status",
			},
			&cli.BoolFlag{
				Name:        "expired",
				DefaultText: "unset",
				Usage:       "filter sessions by expired status",
			},
			&cli.BoolFlag{
				Name:  "json",
				Usage: "output as json",
			},
			&cli.BoolFlag{
				Name:  "restart",
				Usage: "send restart event for all listed sessions",
			},
			&cli.BoolFlag{
				Name:  "clean",
				Usage: "remove uploads",
			},
		},
		Before: func(c *cli.Context) error {
			return configlog.ReturnFatal(parser.ParseConfig(cfg))
		},
		Action: func(c *cli.Context) error {
			f, ok := registry.NewFuncs[cfg.Driver]
			if !ok {
				fmt.Fprintf(os.Stderr, "Unknown filesystem driver '%s'\n", cfg.Driver)
				os.Exit(1)
			}
			drivers := revaconfig.StorageProviderDrivers(cfg)
			fs, err := f(drivers[cfg.Driver].(map[string]interface{}), nil)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to initialize filesystem driver '%s'\n", cfg.Driver)
				return err
			}

			managingFS, ok := fs.(storage.UploadSessionLister)
			if !ok {
				fmt.Fprintf(os.Stderr, "'%s' storage does not support listing upload sessions\n", cfg.Driver)
				os.Exit(1)
			}

			var stream events.Stream
			if c.Bool("restart") {
				stream, err = event.NewStream(cfg)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to create event stream: %v\n", err)
					os.Exit(1)
				}
			}

			filter := buildFilter(c)
			uploads, err := managingFS.ListUploadSessions(c.Context, filter)
			if err != nil {
				return err
			}

			var (
				table *tw.Table
				raw   []Session
			)

			if !c.Bool("json") {
				fmt.Println(buildInfo(filter))

				table = tw.NewWriter(os.Stdout)
				table.SetHeader([]string{"Space", "Upload Id", "Name", "Offset", "Size", "Executant", "Owner", "Expires", "Processing", "Scan Date", "Scan Result"})
				table.SetAutoFormatHeaders(false)
			}

			for _, u := range uploads {
				ref := u.Reference()
				sr, sd := u.ScanData()

				session := Session{
					Space:      ref.GetResourceId().GetSpaceId(),
					ID:         u.ID(),
					Filename:   u.Filename(),
					Offset:     u.Offset(),
					Size:       u.Size(),
					Executant:  u.Executant(),
					SpaceOwner: u.SpaceOwner(),
					Expires:    u.Expires(),
					Processing: u.IsProcessing(),
					ScanDate:   sd,
					ScanResult: sr,
				}

				if c.Bool("json") {
					raw = append(raw, session)
				} else {
					table.Append([]string{
						session.Space,
						session.ID,
						session.Filename,
						strconv.FormatInt(session.Offset, 10),
						strconv.FormatInt(session.Size, 10),
						session.Executant.OpaqueId,
						session.SpaceOwner.GetOpaqueId(),
						session.Expires.Format(time.RFC3339),
						strconv.FormatBool(session.Processing),
						session.ScanDate.Format(time.RFC3339),
						session.ScanResult,
					})
				}

				if c.Bool("restart") {
					if err := events.Publish(context.Background(), stream, events.ResumePostprocessing{
						UploadID:  u.ID(),
						Timestamp: utils.TSNow(),
					}); err != nil {
						fmt.Fprintf(os.Stderr, "Failed to send restart event for upload session '%s'\n", u.ID())
						// if publishing fails there is no need to try publishing other events - they will fail too.
						os.Exit(1)
					}
				}

				if c.Bool("clean") {
					if err := u.Purge(c.Context); err != nil {
						fmt.Fprintf(os.Stderr, "Failed to clean upload session '%s'\n", u.ID())
					}
				}

			}

			if !c.Bool("json") {
				table.Render()
				return nil
			}

			j, err := json.Marshal(raw)
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Println(string(j))
			return nil
		},
	}
}

// PurgeExpiredUploads is the entry point for the clean command
func PurgeExpiredUploads(cfg *config.Config) *cli.Command {
	return &cli.Command{
		Name:  "clean",
		Usage: "Clean up leftovers from expired uploads",
		Before: func(c *cli.Context) error {
			return configlog.ReturnFatal(parser.ParseConfig(cfg))
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Warning: This command is deprecated, use 'uploads sessions --clean' instead")
			f, ok := registry.NewFuncs[cfg.Driver]
			if !ok {
				fmt.Fprintf(os.Stderr, "Unknown filesystem driver '%s'\n", cfg.Driver)
				os.Exit(1)
			}
			drivers := revaconfig.StorageProviderDrivers(cfg)
			fs, err := f(drivers[cfg.Driver].(map[string]interface{}), nil)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to initialize filesystem driver '%s'\n", cfg.Driver)
				return err
			}

			managingFS, ok := fs.(storage.UploadSessionLister)
			if !ok {
				fmt.Fprintf(os.Stderr, "'%s' storage does not support clean expired uploads\n", cfg.Driver)
				os.Exit(1)
			}

			wg := sync.WaitGroup{}
			wg.Add(1)
			processing := false
			expired := true
			uploads, err := managingFS.ListUploadSessions(c.Context, storage.UploadSessionFilter{Expired: &expired, Processing: &processing})
			if err != nil {
				return err
			}

			fmt.Println("purging uploads:")
			go func() {
				for _, u := range uploads {
					ref := u.Reference()
					fmt.Printf(" - %s (Space: %s, Name: %s, Size: %d/%d, Expires: %s, Processing: %t)\n", ref.GetResourceId().GetSpaceId(), u.ID(), u.Filename(), u.Offset(), u.Size(), u.Expires(), u.IsProcessing())
					u.Purge(c.Context)
				}
				wg.Done()
			}()

			wg.Wait()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to clean expired uploads '%s'\n", err)
				return err
			}
			return nil
		},
	}
}

func buildFilter(c *cli.Context) storage.UploadSessionFilter {
	filter := storage.UploadSessionFilter{}
	if c.IsSet("processing") {
		processingValue := c.Bool("processing")
		filter.Processing = &processingValue
	}
	if c.IsSet("expired") {
		expiredValue := c.Bool("expired")
		filter.Expired = &expiredValue
	}
	if c.IsSet("id") {
		idValue := c.String("id")
		filter.ID = &idValue
	}
	return filter
}

func buildInfo(filter storage.UploadSessionFilter) string {
	var b strings.Builder
	if filter.Processing != nil {
		if !*filter.Processing {
			b.WriteString("Not ")
		}
		if b.Len() == 0 {
			b.WriteString("Processing")
		} else {
			b.WriteString("processing")
		}
	}

	if filter.Expired != nil {
		if b.Len() != 0 {
			b.WriteString(", ")
		}
		if !*filter.Expired {
			if b.Len() == 0 {
				b.WriteString("Not ")
			} else {
				b.WriteString("not ")
			}
		}
		if b.Len() == 0 {
			b.WriteString("Expired")
		} else {
			b.WriteString("expired")
		}
	}

	if b.Len() == 0 {
		b.WriteString("Session")
	} else {
		b.WriteString(" session")
	}

	if filter.ID != nil {
		b.WriteString(" with id '" + *filter.ID + "'")
	} else {
		// to make `session` plural
		b.WriteString("s")
	}

	b.WriteString(":")
	return b.String()
}
