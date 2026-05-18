package data

import (
	"common/api/gen/common"
	commonClient "common/pkg/client"
	"signal/internal/conf"
	"signal/internal/data/gen/cache"
	"signal/internal/data/gen/client"
	"signal/internal/data/gen/repo"

	"github.com/google/wire"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	ProvideRedis,
	ProvideConsul,
	ProvideNats,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,
	commonClient.NewHttpClient(),

	repo.NewNodeRepo,
	cache.NewNodeCache,
	cache.NewSessionCache,
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
