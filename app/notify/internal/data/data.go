package data

import (
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	"net/http"
	bizchannel "notify/internal/biz/channel"
	"notify/internal/conf"
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

func ProvideRedis(c *conf.Bootstrap) *common.Redis   { return c.Data.Redis }
func ProvideConsul(c *conf.Bootstrap) *common.Consul { return c.Data.Consul }
func ProvideNats(c *conf.Bootstrap) *common.Nats     { return c.Data.Nats }
