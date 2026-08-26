package module

import (
	"common/pkg/client/rpc"
	commonmodule "common/pkg/module"
	commonserver "common/pkg/server"
	"game_idle/internal/biz"
	"game_idle/internal/config"
	"game_idle/internal/data"
	gameidleserver "game_idle/internal/server"
	"game_idle/internal/service"

	"github.com/go-kratos/kratos/v3/transport"
	"github.com/google/wire"
)

// ProviderSet 提供挂机游戏模块的依赖项。
var ProviderSet = wire.NewSet(
	provideBootstrap,
	commonmodule.InfrastructureProviderSet,
	data.ModuleProviderSet,
	biz.BizProviderSet,
	service.ServiceProviderSet,
	gameidleserver.NewTimeWheelServer,
	gameidleserver.NewActionQueueServer,
	ProvideModuleServers,
	newModule,
)

type Config = commonmodule.Config[*config.Bootstrap]

type Module struct {
	Name     string
	Services []commonserver.Service
	Servers  []transport.Server
}

func newModule(
	config *Config,
	services []commonserver.Service,
	servers []transport.Server,
) *Module {
	return &Module{
		Name:     config.Server().GetName(),
		Services: services,
		Servers:  servers,
	}
}

func provideBootstrap(c *Config) *config.Bootstrap { return c.Bootstrap() }

// ProvideModuleServers 提供模块内部随生命周期启动的后台服务。
func ProvideModuleServers(
	timeWheelServer *gameidleserver.TimeWheelServer,
	actionQueueServer *gameidleserver.ActionQueueServer,
	webSocketService *service.WebSocketService,
) []transport.Server {
	return []transport.Server{
		timeWheelServer,
		actionQueueServer,
		webSocketService,
	}
}

// Build 构造挂机游戏模块并返回单体可收集的模块能力。
func Build(runtime *commonmodule.Runtime, name string) (commonmodule.Mounted, func(), error) {
	values, err := runtime.Values(name)
	if err != nil {
		return commonmodule.Mounted{}, func() {}, err
	}
	moduleConfig, err := commonmodule.NewConfig(
		runtime.Config,
		values,
		name,
		func() *config.Bootstrap { return &config.Bootstrap{} },
	)
	if err != nil {
		return commonmodule.Mounted{}, func() {}, err
	}
	module, cleanup, err := wireModule(moduleConfig, runtime.Logger, runtime.Infrastructure)
	if err != nil {
		return commonmodule.Mounted{}, cleanup, err
	}
	return commonmodule.Mounted{
		Module:   module,
		Services: module.Services,
		Servers:  module.Servers,
	}, cleanup, nil
}

func Descriptor() commonmodule.Descriptor {
	return commonmodule.NewDescriptor(
		Build,
		commonmodule.WithMount(rpc.MountGameIdleServices[commonserver.Service]),
	)
}
