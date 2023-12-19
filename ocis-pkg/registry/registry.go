package registry

import (
	"os"
	"strings"
	"sync"
	"time"

	rRegistry "github.com/cs3org/reva/v2/pkg/registry"
	consulr "github.com/go-micro/plugins/v4/registry/consul"
	etcdr "github.com/go-micro/plugins/v4/registry/etcd"
	kubernetesr "github.com/go-micro/plugins/v4/registry/kubernetes"
	mdnsr "github.com/go-micro/plugins/v4/registry/mdns"
	memr "github.com/go-micro/plugins/v4/registry/memory"
	natsr "github.com/go-micro/plugins/v4/registry/nats"
	"github.com/owncloud/ocis/v2/ocis-pkg/natsjsregistry"
	mRegistry "go-micro.dev/v4/registry"
	"go-micro.dev/v4/registry/cache"
)

const (
	registryEnv         = "MICRO_REGISTRY"
	registryAddressEnv  = "MICRO_REGISTRY_ADDRESS"
	regisryUsernameEnv  = "MICRO_REGISTRY_AUTH_USERNAME"
	registryPasswordEnv = "MICRO_REGISTRY_AUTH_PASSWORD"
)

var (
	once      sync.Once
	regPlugin string
	reg       mRegistry.Registry
)

// Configure configures the registry
func Configure(plugin string) {
	if reg == nil {
		regPlugin = plugin
	}
}

// GetRegistry returns a configured micro registry based on Micro env vars.
// It defaults to mDNS, so mind that systems with mDNS disabled by default (i.e SUSE) will have a hard time
// and it needs to explicitly use etcd. Os awareness for providing a working registry out of the box should be done.
func GetRegistry() mRegistry.Registry {
	once.Do(func() {
		addresses := strings.Split(os.Getenv(registryAddressEnv), ",")
		// prefer env of setting from Configure()
		plugin := os.Getenv(registryEnv)
		if plugin == "" {
			plugin = regPlugin
		}

		switch plugin {
		case "nats":
			reg = natsr.NewRegistry(
				mRegistry.Addrs(addresses...),
				natsr.RegisterAction("put"),
			)
		case "kubernetes":
			reg = kubernetesr.NewRegistry(
				mRegistry.Addrs(addresses...),
			)
		case "etcd":
			reg = etcdr.NewRegistry(
				mRegistry.Addrs(addresses...),
			)
		case "consul":
			reg = consulr.NewRegistry(
				mRegistry.Addrs(addresses...),
			)
		case "memory":
			reg = memr.NewRegistry()
		case "natsjs", "nats-js", "nats-js-kv": // for backwards compatibility - we will stick with one of those
			reg = natsjsregistry.NewRegistry(
				mRegistry.Addrs(addresses...),
				natsjsregistry.Authenticate(os.Getenv(regisryUsernameEnv), os.Getenv(registryPasswordEnv)),
			)
		default:
			reg = mdnsr.NewRegistry()
		}

		// No cache needed for in-memory registry
		if plugin != "memory" {
			reg = cache.New(reg, cache.WithTTL(30*time.Second))
		}

		// fixme: lazy initialization of reva registry, needs refactor to a explicit call per service
		_ = rRegistry.Init(reg)
	})
	// always use cached registry to prevent registry
	// lookup for every request
	return reg
}
