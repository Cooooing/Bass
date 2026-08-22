package data

import (
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	"content/internal/config"
	"content/internal/data/client"
	"content/internal/data/repo"

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
	rpc.ProvideUserClient,
	rpc.ProvideSchedulerClient,
)

// ModuleProviderSet 提供不依赖服务发现的模块数据层依赖。
var ModuleProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	ProvideRedis,
	ProvideNats,
	client.ProvideTx,
	repo.NewArticleRepo,
	repo.NewCommentRepo,
	repo.NewCommentActionRecordRepo,
	repo.NewPostscriptRepo,
	repo.NewArticleActionRecordRepo,
	repo.NewArticleViewCacheRepo,
	repo.NewArticleViewRecordRepo,
	repo.NewOutboxEventRepo,
	repo.NewContentModerationRecordRepo,
	repo.NewDomainRepo,
	repo.NewTagRepo,
	repo.NewUserClient,
	repo.NewDelayedTaskClient,
	repo.NewNatsEventClient,
)

func ProvideRedis(c *config.Bootstrap) *common.Redis {
	return c.Redis
}

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}

func ProvideNats(c *config.Bootstrap) *common.Nats {
	return c.Nats
}
