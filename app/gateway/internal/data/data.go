package data

import (
	"common/api/gen/common"
	commonClient "common/pkg/client"
	"gateway/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	NewBaseData,

	ProvideRedis,
	ProvideConsul,
	ProvideNats,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,
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

type BaseData struct {
	Conf   *conf.Bootstrap
	Log    *log.Helper
	Consul *commonClient.ConsulClient
	Redis  *commonClient.RedisClient
}

func NewBaseData(conf *conf.Bootstrap, logger log.Logger, consul *commonClient.ConsulClient, redis *commonClient.RedisClient) *BaseData {
	return &BaseData{
		Conf:   conf,
		Log:    log.NewHelper(logger),
		Consul: consul,
		Redis:  redis,
	}
}
