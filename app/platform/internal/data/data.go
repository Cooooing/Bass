package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	"platform/internal/conf"
	"platform/internal/data/client"
	"platform/internal/data/oss"
	"platform/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet 是 data 层依赖集合。
var DataProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	ProvideRedis,
	ProvideConsul,
	ProvideNats,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,

	client.ProvideTx,

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
