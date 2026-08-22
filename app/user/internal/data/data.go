package data

import (
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	"user/internal/config"
	"user/internal/data/client"
	"user/internal/data/repo"

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
	rpc.ProvideNotifyClient,
	rpc.ProvidePlatformClient,
	rpc.ProvideSchedulerClient,
)

// ModuleProviderSet 提供不依赖服务发现的模块数据层依赖。
var ModuleProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	ProvideRedis,
	ProvideNats,
	client.ProvideTx,
	repo.NewAccountRepo,
	repo.NewRelationRepo,
	repo.NewPreferencesRepo,
	repo.NewPrivacySettingRepo,
	repo.NewLocationRepo,
	repo.NewTotpRepo,
	repo.NewCheckinRecordRepo,
	repo.NewCheckinStatRepo,
	repo.NewLoginLogRepo,
	repo.NewBanRecordRepo,
	repo.NewRbacRepo,
	repo.NewOutboxEventRepo,
	repo.NewTotpSecretCache,
	repo.NewAuthCacheRepo,
	repo.NewNotificationRateLimitClient,
	repo.NewDelayedTaskClient,
	repo.NewIPClient,
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
