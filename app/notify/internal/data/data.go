package data

import (
	"common/api/gen/common"
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"fmt"
	"notify/internal/conf"
	datachannel "notify/internal/data/channel"
	"notify/internal/data/client"
	"notify/internal/data/oss"
	"notify/internal/data/repo"
	"notify/internal/enum"

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

	client.ProvideTx,

	repo.NewNotificationMetaRepo,
	repo.NewNotificationRecordRepo,
	repo.NewNotificationTemplateRepo,
	repo.NewNotificationSettingRepo,
	repo.NewObjectStorageRepo,
	repo.NewInboxEventRepo,
	repo.NewNotificationDeliveryRepo,
	repo.NewUserClient,

	datachannel.NewClient,
	datachannel.NewEmailChannel,
	NewSMSClient,
	datachannel.NewSMSChannel,
	datachannel.NewWebhookChannel,

	datachannel.NewTencentSMSClient,

	oss.ProviderSet,

	rpc.ProvideUserClient,
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

func NewSMSClient(
	conf *conf.Bootstrap,
	tencentSMS *datachannel.TencentSMSClient,
) (datachannel.SMSClient, error) {
	if conf == nil || conf.Server == nil || conf.Server.Sms == nil {
		return nil, fmt.Errorf("sms config is required")
	}
	switch enum.SMSType(conf.Server.Sms.Provider) {
	case enum.SMSTypeTencent:
		if tencentSMS == nil {
			return nil, fmt.Errorf("tencent sms client is required")
		}
		return tencentSMS, nil
	default:
		return nil, fmt.Errorf("unsupported sms provider: %s", conf.Server.Sms.Provider)
	}
}
