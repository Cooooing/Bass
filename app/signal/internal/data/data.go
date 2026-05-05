package data

import (
	"common/api/gen/common"
	commonClient "common/pkg/client"
	"common/pkg/util/jwt"
	"signal/internal/conf"
	"signal/internal/data/base"
	"signal/internal/data/cache"
	"signal/internal/data/client"
	"signal/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	base.NewBaseData,

	client.NewDataBaseClient,
	ProvideRedis,
	ProvideConsul,
	ProvideNats,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,
	commonClient.NewHTTPClient(),

	jwt.NewTokenCache,

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
