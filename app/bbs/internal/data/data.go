package data

import (
	"bbs/internal/config"
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	userv1 "common/proto/gen/user/v1"

	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(
	ProvideConsul,
	ProvideRedis,
	commonClient.NewObservability,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	rpc.ProvideUserClient,
	ProvideUserAuthClient,
	rpc.ProvideContentClient,
	rpc.ProvideEconomyClient,
	rpc.ProvideNotifyClient,
	rpc.ProvidePlatformClient,
	NewAssetClient,
	NewAuthClient,
	NewAccountClient,
	NewEconomyClient,
	NewPreferencesClient,
	NewPrivacySettingClient,
	NewLocationClient,
	NewRelationClient,
	NewOtpClient,
	NewCheckinClient,
	NewContentArticleClient,
	NewContentPostscriptClient,
	NewContentCommentClient,
	NewContentDomainClient,
	NewContentTagClient,
	NewNotificationClient,
)

func ProvideRedis(c *config.Bootstrap) *common.Redis {
	return c.Redis
}

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}

func ProvideUserAuthClient(userClient *rpc.UserClient) userv1.AuthServiceClient {
	return userClient.Auth
}
