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

// DataProviderSet 提供微服务模式的数据层依赖。
var DataProviderSet = wire.NewSet(
	ModuleProviderSet,
	commonClient.NewObservability,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,
	ProvideConsul,
	commonClient.NewConsulClient,
	rpc.ProvideContentClient,
	rpc.ProvideUserClient,
)

// ModuleProviderSet 提供不依赖服务发现的模块数据层依赖。
var ModuleProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	client.ProvideTx,
	ProvideRedis,
	ProvideNats,
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
