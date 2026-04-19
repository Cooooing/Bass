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
	commonClient.NewRabbitMQClient,
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
	conf     *conf.Bootstrap
	log      *log.Helper
	consul   *commonClient.ConsulClient
	redis    *commonClient.RedisClient
	rabbitmq *commonClient.RabbitMQClient
}

func NewBaseData(conf *conf.Bootstrap, log *log.Helper, consul *commonClient.ConsulClient, redis *commonClient.RedisClient, rabbitmq *commonClient.RabbitMQClient) *BaseData {
	return &BaseData{
		conf:     conf,
		log:      log,
		consul:   consul,
		redis:    redis,
		rabbitmq: rabbitmq,
	}
}
