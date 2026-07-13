package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	"scheduler/internal/conf"
	"scheduler/internal/data/client"
	"scheduler/internal/data/repo"

	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	client.ProvideTx,
	ProvideRedis,
	ProvideNats,
	ProvideConsul,
	commonClient.NewObservability,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,
	commonClient.NewLarkWebhookClient,
	repo.NewTaskRepo,
	repo.NewTaskVersionRepo,
	repo.NewTaskExecutionRecordRepo,
	repo.NewTaskLockRepo,
	repo.NewTaskEventBus,
	repo.NewTaskAlert,
)

func ProvideRedis(c *conf.Bootstrap) *common.Redis {
	return c.Redis
}

func ProvideNats(c *conf.Bootstrap) *common.Nats {
	return c.Nats
}

func ProvideConsul(c *conf.Bootstrap) *common.Consul {
	return c.Consul
}
