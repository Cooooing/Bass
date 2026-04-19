package data

import (
	"common/api/gen/common"
	commonClient "common/pkg/client"
	"common/pkg/util/jwt"
	"infra/internal/conf"
	"infra/internal/data/base"
	"infra/internal/data/client"
	"infra/internal/data/oss"
	"infra/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	base.NewBaseData,
	oss.ProviderSet,
	client.NewDataBaseClient,
	ProvideRedis,
	ProvideConsul,
	ProvideRabbitMQ,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewRabbitMQClient,
	repo.NewObjectStorageRepo,

	jwt.NewTokenCache,
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
