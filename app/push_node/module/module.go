package module

import (
	"common/pkg/server"
	"common/proto/gen/common"
	"fmt"
	"push_node/internal/config"
	"push_node/internal/service"

	"github.com/google/wire"
)

// ProviderSet 提供推送节点模块的依赖项。
var ProviderSet = wire.NewSet(
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

type Clients struct{}

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

func newClients() *Clients { return &Clients{} }
