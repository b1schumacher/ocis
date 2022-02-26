// Package store implements the go-micro store interface
package store

import (
	"context"
	"log"
	"sync"

	"github.com/cs3org/reva/pkg/storage/utils/metadata"
	olog "github.com/owncloud/ocis/ocis-pkg/log"
	"github.com/owncloud/ocis/settings/pkg/config"
	"github.com/owncloud/ocis/settings/pkg/settings"
)

var (
	// Name is the default name for the settings store
	Name                   = "ocis-settings"
	managerName            = "metadata"
	settingsSpaceID        = "f1bdd61a-da7c-49fc-8203-0558109d1b4f" // uuid.Must(uuid.NewV4()).String()
	rootFolderLocation     = "settings"
	bundleFolderLocation   = "settings/bundles"
	accountsFolderLocation = "settings/accounts"
	valuesFolderLocation   = "settings/values"
)

// MetadataClient is the interface to talk to metadata service
type MetadataClient interface {
	SimpleDownload(ctx context.Context, id string) ([]byte, error)
	SimpleUpload(ctx context.Context, id string, content []byte) error
	Delete(ctx context.Context, id string) error
	ReadDir(ctx context.Context, id string) ([]string, error)
	MakeDirIfNotExist(ctx context.Context, id string) error
	Init(ctx context.Context, id string) error
}

// Store interacts with the filesystem to manage settings information
type Store struct {
	Logger olog.Logger

	mdc       MetadataClient
	cfg       *config.Config
	initStore func(settings.Manager)

	init *sync.Once
	l    *sync.Mutex
}

// Init initialize the store once, later calls are noops
func (s Store) Init() {
	if s.mdc != nil {
		return
	}

	s.l.Lock()
	defer s.l.Unlock()
	var err error
	s.init.Do(func() {
		//b := backoff.NewExponentialBackOff()
		//b.MaxElapsedTime = 4 * time.Second
		//backoff.Retry(func() error {
		err = s.initMetadataClient()
		//return err

		//}, b)

	})
	if err != nil {
		log.Fatal(err)
	}
}

// New creates a new store
func New(cfg *config.Config, initstore func(settings.Manager)) settings.Manager {
	s := Store{
		Logger: olog.NewLogger(
			olog.Color(cfg.Log.Color),
			olog.Pretty(cfg.Log.Pretty),
			olog.Level(cfg.Log.Level),
			olog.File(cfg.Log.File),
		),
		initStore: initstore,
		l:         &sync.Mutex{},
		init:      &sync.Once{},
	}

	return &s
}

// NewMetadataClient returns the MetadataClient
func NewMetadataClient(cfg *config.Config) MetadataClient {
	mdc, err := metadata.NewCS3Storage("127.0.0.1:9142", "127.0.0.1:9215", "058bff95-6708-4fe5-91e4-9ea3d377588b", "change-me-please")
	if err != nil {
		log.Fatal("error connecting to mdc:", err)
	}
	return mdc

}

// we need to lazy initialize the MetadataClient because metadata service might not be ready
func (s Store) initMetadataClient() error {
	s.mdc = NewMetadataClient(s.cfg)

	// TODO: this fails because of authentication issues
	err := s.mdc.Init(nil, settingsSpaceID)
	if err != nil {
		return err
	}

	for _, p := range []string{
		rootFolderLocation,
		accountsFolderLocation,
		bundleFolderLocation,
		valuesFolderLocation,
	} {
		err = s.mdc.MakeDirIfNotExist(nil, p)
		if err != nil {
			return err
		}
	}

	s.initStore(s)
	return nil
}

func init() {
	settings.Registry[managerName] = New
}
