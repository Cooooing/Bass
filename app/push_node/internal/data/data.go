package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	bizrepo "push_node/internal/biz/repo"
	"push_node/internal/config"
	"push_node/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet 提供微服务模式的数据层依赖。
var DataProviderSet = wire.NewSet(
	ModuleProviderSet,
	commonClient.NewObservability,
	commonClient.NewNatsClient,
	ProvideConsul,
	commonClient.NewConsulClient,
)

// ModuleProviderSet 提供不依赖服务发现的模块数据层依赖。
var ModuleProviderSet = wire.NewSet(
	ProvideNats,
	repo.NewConnectionRegistryRepo,
	wire.Bind(new(bizrepo.ConnectionRegistry), new(*repo.ConnectionRegistryRepo)),
	ProvideSubscriber,
)

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}

func ProvideNats(c *config.Bootstrap) *common.Nats {
	return c.Nats
}

func ProvideSubscriber(natsClient *commonClient.NatsClient) commonClient.Subscriber {
	return natsClient
}
