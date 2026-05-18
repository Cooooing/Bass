package data

import (
	"common/api/gen/common"
	commonClient "common/pkg/client"
	"notify/internal/conf"
	"notify/internal/data/client"
	"notify/internal/data/oss"
	"notify/internal/data/repo"

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

	repo.NewNotificationMetaRepo,
	repo.NewNotificationRecordRepo,
	repo.NewNotificationTemplateRepo,
	repo.NewNotificationSettingRepo,
	repo.NewObjectStorageRepo,
	oss.ProviderSet,
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
