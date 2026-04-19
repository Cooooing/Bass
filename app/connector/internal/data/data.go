package data

import (
	commonClient "common/pkg/client"
	commonModel "common/pkg/model"
	"connector/internal/conf"
	database "connector/internal/data/base"
	"connector/internal/data/cache"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	NewConsulClient,
	database.NewBaseData,
	NewRedisClient,

	cache.NewSessionCache,
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

func NewConsulClient(log *log.Helper, conf *conf.Bootstrap) (*commonClient.ConsulClient, func(), error) {
	return nil, func() {

	}, nil
}

func NewRedisClient(log *log.Helper, conf *conf.Bootstrap) (*commonClient.RedisClient, func(), error) {
	c := &commonModel.RedisConf{}
	err := copier.Copy(c, conf.Data.Redis)
	if err != nil {
		return nil, nil, err
	}
	return commonClient.NewRedisClient(log, c)
}
