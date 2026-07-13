package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	"user/internal/conf"
	"user/internal/data/client"
	"user/internal/data/repo"

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
	commonClient.NewLarkWebhookClient,
	commonClient.NewDeadLetterAlertClient,

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
	repo.NewOutboxEventRepo,
	repo.NewTotpSecretCache,
	repo.NewNatsEventClient,
)

func ProvideRedis(c *conf.Bootstrap) *common.Redis {
	return c.Redis
}

func ProvideConsul(c *conf.Bootstrap) *common.Consul {
	return c.Consul
}

func ProvideNats(c *conf.Bootstrap) *common.Nats {
	return c.Nats
}
