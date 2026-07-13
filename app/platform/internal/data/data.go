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

var DataProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	ProvideRedis,
	ProvideConsul,
	ProvideNats,
	commonClient.NewObservability,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewRedisLock,
	commonClient.NewNatsClient,
	client.ProvideTx,
	repo.NewObjectStorageRepo,
	oss.ProviderSet,
)

func ProvideRedis(c *conf.Bootstrap) *common.Redis   { return c.Redis }
func ProvideConsul(c *conf.Bootstrap) *common.Consul { return c.Consul }
func ProvideNats(c *conf.Bootstrap) *common.Nats     { return c.Nats }
