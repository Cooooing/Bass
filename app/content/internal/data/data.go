package data

import (
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"

	"content/internal/config"
	"content/internal/data/client"
	"content/internal/data/gen"
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
	commonClient.NewSQLClient,
	ProvideDatabase,
	client.NewEntClient,
	ProvideGenClient,
	ProvideRedis,
	ProvideNats,
	client.ProvideTx,
	repo.NewArticleRepo,
	repo.NewStatisticsRepo,
	repo.NewArticleQARepo,
	repo.NewArticlePollRepo,
	repo.NewArticleLotteryRepo,
	repo.NewPostgresArticleSearchRepo,
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
	repo.NewMoonbreezeRepo,
	repo.NewMoonbreezeRateLimitCache,
	ProvideDTMConfig,
	ProvideBusinessConfig,
	commonClient.NewDTMClient,
	commonClient.NewTCCBranchCaller,
	ProvideTCCBranchExecutor,
)

func ProvideRedis(c *config.Bootstrap) *common.Redis {
	return c.Redis
}

func ProvideDatabase(c *config.Bootstrap) *common.Database {
	return c.Database
}

func ProvideGenClient(c *client.EntClient) *gen.Client {
	return c.Client
}

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}

func ProvideNats(c *config.Bootstrap) *common.Nats {
	return c.Nats
}

func ProvideDTMConfig(c *config.Bootstrap) *common.DTM           { return c.Dtm }
func ProvideBusinessConfig(c *config.Bootstrap) *config.Business { return c.Business }

func ProvideTCCBranchExecutor(sqlClient *commonClient.SQLClient, entClient *client.EntClient, dtm *commonClient.DTMClient) (commonClient.TCCBranchExecutor, error) {
	return commonClient.NewTCCBranchExecutor(dtm, sqlClient.DB, sqlClient.Driver, entClient.ClientFromTx)
}
