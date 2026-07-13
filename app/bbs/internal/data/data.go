package data

import (
	"bbs/internal/conf"
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
	rpc.ProvideNotifyClient,
	NewAuthClient,
	NewAccountClient,
	NewPreferencesClient,
	NewPrivacySettingClient,
	NewLocationClient,
	NewRelationClient,
	NewTotpClient,
	NewContentArticleClient,
	NewContentPostscriptClient,
	NewContentCommentClient,
	NewContentDomainClient,
	NewContentTagClient,
	NewNotificationClient,
)

func ProvideRedis(c *conf.Bootstrap) *common.Redis {
	return c.Redis
}

func ProvideConsul(c *conf.Bootstrap) *common.Consul {
	return c.Consul
}

func ProvideUserAuthClient(userClient *rpc.UserClient) userv1.AuthServiceClient {
	return userClient.Auth
}
