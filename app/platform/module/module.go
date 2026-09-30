package module

import (
	"common/pkg/client/rpc"
	commonmodule "common/pkg/module"
	"common/pkg/server"
	"platform/internal/biz"
	"platform/internal/config"
	"platform/internal/data"
	platformserver "platform/internal/server"
	"platform/internal/service"

	"github.com/go-kratos/kratos/v3/transport"
	"github.com/google/wire"
)

// ProviderSet 提供平台模块的依赖项。
var ProviderSet = wire.NewSet(
	provideBootstrap,
	data.ModuleProviderSet,
	commonmodule.InfrastructureProviderSet,
	biz.BizProviderSet,
	service.ServiceProviderSet,
	platformserver.NewAssetEventConsumerServer,
	provideServers,
	newModule,
)

type Config = commonmodule.Config[*config.Bootstrap]

type Module struct {
	Name     string
	Services []server.Service
	Servers  []transport.Server
}

func newModule(config *Config, services []server.Service, servers []transport.Server) *Module {
	return &Module{Name: config.Server().GetName(), Services: services, Servers: servers}
}

func provideBootstrap(c *Config) *config.Bootstrap { return c.Bootstrap() }

func provideServers(assetEventConsumerServer *platformserver.AssetEventConsumerServer) []transport.Server {
	return []transport.Server{assetEventConsumerServer}
}

// Build 构造平台模块并返回单体可收集的模块能力。
func Build(runtime *commonmodule.Runtime, name string) (commonmodule.Mounted, func(), error) {
	values, err := runtime.Values(name)
	if err != nil {
		return commonmodule.Mounted{}, func() {}, err
	}
	moduleConfig, err := commonmodule.NewConfig(runtime.Config, values, name, func() *config.Bootstrap { return &config.Bootstrap{} })
	if err != nil {
		return commonmodule.Mounted{}, func() {}, err
	}
	module, cleanup, err := wireModule(moduleConfig, runtime.Logger, runtime.Infrastructure)
	if err != nil {
		return commonmodule.Mounted{}, cleanup, err
	}
	return commonmodule.Mounted{Module: module, Services: module.Services, Servers: module.Servers}, cleanup, nil
}

func Descriptor() commonmodule.Descriptor {
	return commonmodule.NewDescriptor(
		Build,
		commonmodule.WithLocalClient(rpc.NewPlatformClient),
		commonmodule.WithMount(rpc.MountPlatformServices[server.Service]),
	)
}
