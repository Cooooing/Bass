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

// DataProviderSet 提供微服务模式的数据层依赖。
var DataProviderSet = wire.NewSet(
	ModuleProviderSet,
	commonClient.NewObservability,
	commonClient.NewRedisClient,
	commonClient.NewRedisLock,
	commonClient.NewNatsClient,
	ProvideConsul,
	commonClient.NewConsulClient,
	rpc.ProvideUserClient,
	rpc.ProvideContentClient,
)

// ModuleProviderSet 提供不依赖服务发现的模块数据层依赖。
var ModuleProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	ProvideRedis,
	ProvideNats,
	commonClient.NewLarkWebhookClient,
	wire.Value(http.DefaultClient),
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
