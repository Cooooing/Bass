package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	bizrepo "push_hub/internal/biz/repo"
	"push_hub/internal/config"
	"push_hub/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet 提供微服务模式的数据层依赖。
var DataProviderSet = wire.NewSet(
	ModuleProviderSet,
	commonClient.NewObservability,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,
	ProvideConsul,
	commonClient.NewConsulClient,
)

// ModuleProviderSet 提供不依赖服务发现的模块数据层依赖。
var ModuleProviderSet = wire.NewSet(
	ProvideRedis,
	ProvideNats,
	repo.NewNodeRegistryRepo,
	wire.Bind(new(bizrepo.NodeRegistry), new(*repo.NodeRegistryRepo)),
	ProvidePublisher,
)

func ProvideRedis(c *config.Bootstrap) *common.Redis {
	return c.Redis
}

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}

func ProvideNats(c *config.Bootstrap) *common.Nats {
	return c.Nats
}

func ProvidePublisher(natsClient *commonClient.NatsClient) commonClient.Publisher {
	return natsClient
}
