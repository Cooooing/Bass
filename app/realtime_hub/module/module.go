package module

import (
	"common/pkg/client/rpc"
	commonmodule "common/pkg/module"
	"common/pkg/server"
	"common/proto/gen/common"
	"realtime_hub/internal/biz"
	"realtime_hub/internal/config"
	"realtime_hub/internal/data"
	realtimeserver "realtime_hub/internal/server"
	"realtime_hub/internal/service"

	"github.com/go-kratos/kratos/v3/transport"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(ProvideBootstrap, ProvideCommonServer, commonmodule.InfrastructureProviderSet, data.ModuleProviderSet, biz.BizProviderSet, service.ServiceProviderSet, realtimeserver.NewNotificationDeliveryConsumer, realtimeserver.NewControlConsumer, provideServers, newModule)

type Config = commonmodule.Config[*config.Bootstrap]

type Module struct {
	Name     string
	Services []server.Service
	Servers  []transport.Server
}

func newModule(config *Config, services []server.Service, servers []transport.Server) *Module {
	return &Module{Name: config.Server().GetName(), Services: services, Servers: servers}
}

func provideServers(notificationConsumer *realtimeserver.NotificationDeliveryConsumer, controlConsumer *realtimeserver.ControlConsumer) []transport.Server {
	return []transport.Server{notificationConsumer, controlConsumer}
}

func ProvideBootstrap(config *Config) *config.Bootstrap {
	return config.Bootstrap()
}

func ProvideCommonServer(config *Config) *common.Server {
	return config.Server()
}

func Build(runtime *commonmodule.Runtime, name string) (commonmodule.Mounted, func(), error) {
	values, err := runtime.Values(name)
	if err != nil {
		return commonmodule.Mounted{}, func() {}, err
	}
	config, err := commonmodule.NewConfig(runtime.Config, values, name, func() *config.Bootstrap { return &config.Bootstrap{} })
	if err != nil {
		return commonmodule.Mounted{}, func() {}, err
	}
	module, err := wireModule(config, runtime.Logger, runtime.Infrastructure)
	if err != nil {
		return commonmodule.Mounted{}, func() {}, err
	}
	return commonmodule.Mounted{Module: module, Services: module.Services, Servers: module.Servers}, func() {}, nil
}

func Descriptor() commonmodule.Descriptor {
	return commonmodule.NewDescriptor(Build, commonmodule.WithLocalClient(rpc.NewRealtimeHubClient), commonmodule.WithMount(rpc.MountRealtimeHubServices[server.Service]))
}
