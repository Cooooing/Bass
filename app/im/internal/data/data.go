package data

import (
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"

	"im/internal/config"
	"im/internal/data/client"
	"im/internal/data/gen"
	"im/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet 提供微服务模式的数据层依赖。
var DataProviderSet = wire.NewSet(
	ModuleProviderSet,
	commonClient.NewObservability,
	commonClient.NewRedisClient,
	ProvideConsul,
	ProvideRedis,
	ProvideNats,
	commonClient.NewConsulClient,
	commonClient.NewNatsClient,
	rpc.ProvideUserClient,
	rpc.ProvidePlatformClient,
	rpc.ProvideEconomyClient,
)

// ModuleProviderSet 提供不依赖服务发现的模块数据层依赖。
var ModuleProviderSet = wire.NewSet(
	commonClient.NewSQLClient,
	ProvideDatabase,
	client.NewEntClient,
	ProvideGenClient,
	ProvideDTMConfig,
	ProvideBusinessConfig,
	commonClient.NewDTMClient,
	commonClient.NewTCCBranchCaller,
	ProvideTCCBranchExecutor,
	client.ProvideTx,
	repo.NewPrivateChatSessionRepo,
	repo.NewPrivateChatMessageRepo,
	repo.NewGlobalChatMessageRepo,
	repo.NewStatisticsRepo,
	repo.NewGlobalChatMuteRepo,
	repo.NewGlobalChatTopicRepo,
	repo.NewGlobalChatTopicRateLimiter,
	repo.NewGlobalChatBarrageRepo,
	repo.NewGlobalChatBarrageRateLimiter,
	repo.NewRelationClient,
	repo.NewRealtimeEventPublisher,
	repo.NewEmojiGroupRepo,
	repo.NewEmojiShareCache,
	repo.NewEmojiAssetClient,
)

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}

func ProvideNats(c *config.Bootstrap) *common.Nats {
	return c.Nats
}

func ProvideDatabase(c *config.Bootstrap) *common.Database {
	return c.Database
}

func ProvideRedis(c *config.Bootstrap) *common.Redis             { return c.Redis }
func ProvideDTMConfig(c *config.Bootstrap) *common.DTM           { return c.Dtm }
func ProvideBusinessConfig(c *config.Bootstrap) *config.Business { return c.Business }

func ProvideGenClient(c *client.EntClient) *gen.Client {
	return c.Client
}

func ProvideTCCBranchExecutor(sqlClient *commonClient.SQLClient, entClient *client.EntClient, dtm *commonClient.DTMClient) (commonClient.TCCBranchExecutor, error) {
	return commonClient.NewTCCBranchExecutor(dtm, sqlClient.DB, sqlClient.Driver, entClient.ClientFromTx)
}
