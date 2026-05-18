package data

import (
	"common/api/gen/common"
	commonClient "common/pkg/client"
	"connector/internal/conf"
	"connector/internal/data/gen/cache"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	NewConsulClient,

	ProvideRedis,
	commonClient.NewRedisClient,

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

func NewConsulClient(logger log.Logger, conf *conf.Bootstrap) (*commonClient.ConsulClient, func(), error) {
	return nil, func() {

	}, nil
}
