package module

import (
	"common/pkg/client/rpc"
	commonmodule "common/pkg/module"
	"common/pkg/server"
	pushhubv1 "common/proto/gen/push_hub/v1"
	"push_node/internal/biz"
	"push_node/internal/config"
	"push_node/internal/data"
	"push_node/internal/service"

	"github.com/google/wire"
)

// ProviderSet 提供推送节点模块的依赖项。
var ProviderSet = wire.NewSet(
	provideBootstrap,
	rpc.LocalProviderSet,
	commonmodule.InfrastructureProviderSet,
	providePushHubNodeClient,
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

func providePushHubNodeClient(client *rpc.PushHubClient) pushhubv1.PushHubNodeServiceClient {
	return client.Node
}

// Build 构造推送节点模块并返回单体可收集的模块能力。
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
	)
}
