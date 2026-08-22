package module

import (
	commonclient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/pkg/server"
	"common/proto/gen/common"
	"fmt"
	"push_hub/internal/biz"
	bizrepo "push_hub/internal/biz/repo"
	"push_hub/internal/config"
	"push_hub/internal/data"
	"push_hub/internal/data/repo"
	"push_hub/internal/service"

	"github.com/google/wire"
)

// ProviderSet 提供推送中心模块的依赖项。
var ProviderSet = wire.NewSet(
	provideBootstrap,
	provideCommonServer,
	commonclient.NewObservability,
	data.ProvideRedis,
	data.ProvideNats,
	commonclient.NewRedisClient,
	commonclient.NewNatsClient,
	repo.NewNodeRegistryRepo,
	wire.Bind(new(bizrepo.NodeRegistry), new(*repo.NodeRegistryRepo)),
	wire.Bind(new(commonclient.Publisher), new(*commonclient.NatsClient)),
	biz.BizProviderSet,
	service.ServiceProviderSet,
	newClients,
	newModule,
)

type Config struct {
	bootstrap *config.Bootstrap
}

type Module struct {
	Name     string
	Services []server.Service
	Clients  *Clients
}

type Clients struct {
	PushHub *rpc.PushHubClient
}

func newModule(config *Config, services []server.Service, clients *Clients) *Module {
	return &Module{Name: config.Server().GetName(), Services: services, Clients: clients}
}

func LoadConfig(bootstrapPath string, path string) (*Config, *common.Bootstrap, func(), error) {
	c, bc, cleanup, err := config.LoadConfig(bootstrapPath, path)
	if err != nil {
		return nil, nil, cleanup, err
	}
	return &Config{bootstrap: c}, bc, cleanup, nil
}

func LoadRequiredConfig(name string, bootstrapPath string, path string) (*Config, func(), error) {
	if bootstrapPath == "" || path == "" {
		return nil, func() {}, fmt.Errorf("%s module config path is required", name)
	}
	c, _, cleanup, err := LoadConfig(bootstrapPath, path)
	return c, cleanup, err
}

func (c *Config) Server() *common.Server { return c.bootstrap.GetServer() }
func (c *Config) Trace() *common.Trace   { return c.bootstrap.GetTrace() }

func provideBootstrap(c *Config) *config.Bootstrap { return c.bootstrap }
func provideCommonServer(c *Config) *common.Server { return c.Server() }

func newClients(services []server.Service) *Clients {
	return &Clients{PushHub: rpc.NewLocalPushHubClient(services)}
}
