package data

import (
	"common/proto/gen/common"
	commonClient "common/pkg/client"
	"push_hub/internal/biz/repo"
	"push_hub/internal/conf"
	datarepo "push_hub/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet 是 data 层依赖集合。
var DataProviderSet = wire.NewSet(
	ProvideRedis,
	ProvideConsul,
	ProvideNats,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,

	datarepo.NewNodeRegistryRepo,
	wire.Bind(new(repo.NodeRegistry), new(*datarepo.NodeRegistryRepo)),
	wire.Bind(new(commonClient.Publisher), new(*commonClient.NatsClient)),
)

func ProvideRedis(c *conf.Bootstrap) *common.Redis {
	return c.Data.Redis
}

func ProvideConsul(c *conf.Bootstrap) *common.Consul {
	return c.Data.Consul
}

func ProvideNats(c *conf.Bootstrap) *common.Nats {
	return c.Data.Nats
}
