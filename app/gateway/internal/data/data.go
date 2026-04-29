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
	ProvideRabbitMQ,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
)

func ProvideRedis(c *conf.Bootstrap) *common.Redis {
	return c.Data.Redis
}

func ProvideConsul(c *conf.Bootstrap) *common.Consul {
	return c.Data.Consul
}

func ProvideRabbitMQ(c *conf.Bootstrap) *common.RabbitMQ {
	return c.Data.Rabbitmq
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
