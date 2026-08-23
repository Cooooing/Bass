package module

import (
	"common/pkg/client/rpc"
	commonmodule "common/pkg/module"
	commonserver "common/pkg/server"
	"common/proto/gen/common"
	"scheduler/internal/biz"
	"scheduler/internal/config"
	"scheduler/internal/data"
	schedulerserver "scheduler/internal/server"
	"scheduler/internal/service"

	"github.com/go-kratos/kratos/v3/transport"
	"github.com/google/wire"
)

// ProviderSet 提供调度模块的依赖项。
var ProviderSet = wire.NewSet(
	provideBootstrap,
	provideCommonServer,
	rpc.LocalProviderSet,
	commonmodule.InfrastructureProviderSet,
	data.ModuleProviderSet,
	biz.BizProviderSet,
	service.ServiceProviderSet,
	schedulerserver.NewSchedulerBootstrapServer,
	schedulerserver.NewScheduledTaskConsumerServer,
	schedulerserver.NewDelayedTaskConsumerServer,
	provideServers,
	newModule,
)

type Config = commonmodule.Config[*config.Bootstrap]

type Module struct {
	Name     string
	Services []commonserver.Service
	Servers  []transport.Server
}

func newModule(config *Config, services []commonserver.Service, servers []transport.Server) *Module {
	return &Module{Name: config.Server().GetName(), Services: services, Servers: servers}
}

func provideBootstrap(c *Config) *config.Bootstrap { return c.Bootstrap() }
func provideCommonServer(c *Config) *common.Server { return c.Server() }

func provideServers(
	bootstrapServer *schedulerserver.SchedulerBootstrapServer,
	scheduledTaskConsumerServer *schedulerserver.ScheduledTaskConsumerServer,
	delayedTaskConsumerServer *schedulerserver.DelayedTaskConsumerServer,
) []transport.Server {
	return []transport.Server{
		bootstrapServer,
		scheduledTaskConsumerServer,
		delayedTaskConsumerServer,
	}
}

// Build 构造调度模块并返回单体可收集的模块能力。
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
	return commonmodule.Mounted{Module: module, Services: module.Services, Servers: module.Servers}, cleanup, nil
}

func Descriptor() commonmodule.Descriptor {
	return commonmodule.NewDescriptor(
		Build,
		commonmodule.WithLocalClient(rpc.NewSchedulerClient),
		commonmodule.WithMount(rpc.MountSchedulerServices[commonserver.Service]),
	)
}
