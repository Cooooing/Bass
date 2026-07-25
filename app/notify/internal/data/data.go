package data

import (
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	"net/http"
	bizchannel "notify/internal/biz/channel"
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
	commonClient.NewDeadLetterAlertClient,
	wire.Value(http.DefaultClient),
	rpc.ProvideUserClient,
	rpc.ProvideContentClient,
	client.ProvideTx,
	repo.NewInboxEventRepo,
	repo.NewNotificationRuleRepo,
	repo.NewNotificationStationMessageRepo,
	repo.NewNotificationEmailDeliveryRepo,
	repo.NewNotificationTencentSMSDeliveryRepo,
	repo.NewNotificationLarkWebhookDeliveryRepo,
	repo.NewNotificationRateLimitCache,
	repo.NewUserClient,
	channel.NewEmailClient,
	channel.NewTencentSMSClient,
	channel.NewLarkWebhookClient,
	repo.NewContentClient,
	wire.Bind(new(bizchannel.EmailClient), new(*channel.EmailClient)),
	wire.Bind(new(bizchannel.TencentSMSClient), new(*channel.TencentSMSClient)),
	wire.Bind(new(bizchannel.LarkWebhookClient), new(*channel.LarkWebhookClient)),
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
