package data

import (
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	"net/http"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/config"
	"notify/internal/data/channel"
	"notify/internal/data/client"
	"notify/internal/data/repo"

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
	wire.Value(http.DefaultClient),
	rpc.ProvideUserClient,
	rpc.ProvideContentClient,
	client.ProvideTx,
	repo.NewInboxEventRepo,
	repo.NewNotificationRuleRepo,
	repo.NewNotificationStationTemplateRepo,
	repo.NewNotificationEmailTemplateRepo,
	repo.NewNotificationTencentSMSTemplateRepo,
	repo.NewNotificationLarkWebhookTemplateRepo,
	repo.NewNotificationStationMessageRepo,
	repo.NewNotificationEmailDeliveryRepo,
	repo.NewNotificationTencentSMSDeliveryRepo,
	repo.NewNotificationLarkWebhookDeliveryRepo,
	repo.NewNotificationRateLimitCache,
	repo.NewUserAccountRepo,
	channel.NewEmailClient,
	channel.NewTencentSMSClient,
	channel.NewLarkWebhookClient,
	repo.NewContentClient,
	wire.Bind(new(bizrepo.EmailClient), new(*channel.EmailClient)),
	wire.Bind(new(bizrepo.TencentSMSClient), new(*channel.TencentSMSClient)),
	wire.Bind(new(bizrepo.LarkWebhookClient), new(*channel.LarkWebhookClient)),
)

func ProvideRedis(
	c *config.Bootstrap,
) *common.Redis {
	return c.Redis
}

func ProvideConsul(
	c *config.Bootstrap,
) *common.Consul {
	return c.Consul
}

func ProvideNats(
	c *config.Bootstrap,
) *common.Nats {
	return c.Nats
}
