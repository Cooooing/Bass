package data

import (
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	"scheduler/internal/config"
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
	rpc.ProvideUserClient,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,
	repo.NewScheduledTaskCacheRepo,
	repo.NewDelayedTaskCacheRepo,
	repo.NewScheduledTaskRepo,
	repo.NewScheduledTaskVersionRepo,
	repo.NewScheduledTaskExecutionRecordRepo,
	repo.NewScheduledTaskScheduleNatsRepo,
	repo.NewDelayedTaskScheduleNatsRepo,
	repo.NewDelayedTaskRepo,
	repo.NewDelayedTaskVersionRepo,
	repo.NewDelayedTaskExecutionRecordRepo,
)

func ProvideRedis(
	c *config.Bootstrap,
) *common.Redis {
	return c.Redis
}

func ProvideNats(
	c *config.Bootstrap,
) *common.Nats {
	return c.Nats
}

func ProvideConsul(
	c *config.Bootstrap,
) *common.Consul {
	return c.Consul
}
