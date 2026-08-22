package module

import (
	"common/pkg/client/rpc"
	commonmodule "common/pkg/module"
	"common/pkg/server"
	"notify/internal/biz"
	"notify/internal/config"
	"notify/internal/data"
	"notify/internal/service"

	"github.com/google/wire"
)

// ProviderSet 提供通知模块的依赖项。
var ProviderSet = wire.NewSet(
	provideBootstrap,
	rpc.LocalProviderSet,
	commonmodule.InfrastructureProviderSet,
	data.ModuleProviderSet,
	biz.BizProviderSet,
	service.ServiceProviderSet,
	newModule,
)

type Config = commonmodule.Config[*config.Bootstrap]

type Module struct {
	Name     string
	Services []server.Service
}

func newModule(config *Config, services []server.Service) *Module {
	return &Module{Name: config.Server().GetName(), Services: services}
}

func provideBootstrap(c *Config) *config.Bootstrap { return c.Bootstrap() }

// Build 构造通知模块并返回单体可收集的模块能力。
func Build(runtime *commonmodule.Runtime, name string) (commonmodule.Mounted, func(), error) {
	values, err := runtime.Values(name)
	if err != nil {
		return commonmodule.Mounted{}, func() {}, err
	}
	moduleConfig, err := commonmodule.NewConfig(runtime.Config, values, name, func() *config.Bootstrap { return &config.Bootstrap{} })
	if err != nil {
		return commonmodule.Mounted{}, func() {}, err
	}
	module, cleanup, err := wireModule(moduleConfig, runtime.Logger, runtime.Clients, runtime.Infrastructure)
	if err != nil {
		return commonmodule.Mounted{}, cleanup, err
	}
	return commonmodule.Mounted{Module: module, Services: module.Services}, cleanup, nil
}

func Descriptor() commonmodule.Descriptor {
	return commonmodule.NewDescriptor(
		Build,
		commonmodule.WithLocalClient(rpc.NewNotifyClient),
		commonmodule.WithMount(rpc.MountNotifyServices[server.Service]),
	)
}
