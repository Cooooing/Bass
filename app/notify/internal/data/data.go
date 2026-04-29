package data

import (
	"common/api/gen/common"
	commonClient "common/pkg/client"
	"notify/internal/conf"
	database "notify/internal/data/base"
	"notify/internal/data/client"
	"notify/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	database.NewBaseData,

	client.NewDataBaseClient,
	ProvideRedis,
	ProvideConsul,
	ProvideRabbitMQ,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,

	repo.NewNotificationMetaRepo,
	repo.NewNotificationRecordRepo,
	repo.NewNotificationTemplateRepo,
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
