package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	"user/internal/config"
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
	repo.NewBanRecordRepo,
	repo.NewRbacRepo,
	repo.NewOutboxEventRepo,
	repo.NewTotpSecretCache,
	repo.NewAuthCacheRepo,
	repo.NewDelayedTaskClient,
	repo.NewIPResolver,
	repo.NewNatsEventClient,
)

func ProvideRedis(c *config.Bootstrap) *common.Redis   { return c.Redis }
func ProvideConsul(c *config.Bootstrap) *common.Consul { return c.Consul }
func ProvideNats(c *config.Bootstrap) *common.Nats     { return c.Nats }
