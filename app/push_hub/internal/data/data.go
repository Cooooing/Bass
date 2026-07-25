package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	bizrepo "push_hub/internal/biz/repo"
	"push_hub/internal/config"
	"push_hub/internal/data/repo"

	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(
	ProvideRedis,
	ProvideConsul,
	ProvideNats,
	commonClient.NewObservability,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,
	repo.NewNodeRegistryRepo,
	wire.Bind(new(bizrepo.NodeRegistry), new(*repo.NodeRegistryRepo)),
	wire.Bind(new(commonClient.Publisher), new(*commonClient.NatsClient)),
)

func ProvideRedis(
	c *config.Bootstrap,
) *common.Redis {
	return c.Redis
}

func ProvideConsul(
	c *config.Bootstrap,
) *common.Consul {
	return c.Consul
}

func ProvideNats(
	c *config.Bootstrap,
) *common.Nats {
	return c.Nats
}
