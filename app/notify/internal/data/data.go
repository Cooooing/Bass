package data

import (
	"common/api/gen/common"
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"net/http"
	bizchannel "notify/internal/biz/channel"
	"notify/internal/conf"
	datachannel "notify/internal/data/channel"
	"notify/internal/data/client"
	"notify/internal/data/oss"
	"notify/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet 是 data 层依赖集合。
var DataProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	ProvideRedis,
	ProvideConsul,
	ProvideNats,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,
	NewHTTPClient,

	client.ProvideTx,

	repo.NewObjectStorageRepo,
	repo.NewInboxEventRepo,
	repo.NewNotificationRuleRepo,
	repo.NewNotificationStationMessageRepo,
	repo.NewNotificationEmailDeliveryRepo,
	repo.NewNotificationTencentSMSDeliveryRepo,
	repo.NewNotificationLarkWebhookDeliveryRepo,
	repo.NewUserClient,
	repo.NewContentClient,

	datachannel.NewEmailClient,
	wire.Bind(new(bizchannel.EmailClient), new(*datachannel.EmailClient)),
	datachannel.NewTencentSMSClient,
	wire.Bind(new(bizchannel.TencentSMSClient), new(*datachannel.TencentSMSClient)),
	datachannel.NewLarkWebhookClient,
	wire.Bind(new(bizchannel.LarkWebhookClient), new(*datachannel.LarkWebhookClient)),

	oss.ProviderSet,

	rpc.ProvideUserClient,
	rpc.ProvideContentClient,
)

func ProvideRedis(c *conf.Bootstrap) *common.Redis {
	return c.Data.Redis
}

func ProvideConsul(c *conf.Bootstrap) *common.Consul {
	return c.Data.Consul
}

func ProvideNats(c *conf.Bootstrap) *common.Nats {
	return c.Data.Nats
}

func NewHTTPClient() *http.Client {
	return http.DefaultClient
}
