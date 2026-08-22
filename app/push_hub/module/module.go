package module

import (
	"common/pkg/client/rpc"
	commonmodule "common/pkg/module"
	"common/pkg/server"
	"common/proto/gen/common"
	"push_hub/internal/biz"
	"push_hub/internal/config"
	"push_hub/internal/data"
	"push_hub/internal/service"

	"github.com/google/wire"
)

// ProviderSet 提供推送中心模块的依赖项。
var ProviderSet = wire.NewSet(
	provideBootstrap,
	provideCommonServer,
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
func provideCommonServer(c *Config) *common.Server { return c.Server() }

// Build 构造推送中心模块并返回单体可收集的模块能力。
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
	return commonmodule.Mounted{Module: module, Services: module.Services}, cleanup, nil
}

func Descriptor() commonmodule.Descriptor {
	return commonmodule.NewDescriptor(
		Build,
		commonmodule.WithLocalClient(rpc.NewPushHubClient),
		commonmodule.WithMount(rpc.MountPushHubServices[server.Service]),
	)
}
