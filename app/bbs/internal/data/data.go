package data

import (
	"bbs/internal/conf"
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	userv1 "common/proto/gen/user/v1"

	"github.com/google/wire"
)

// DataProviderSet 是 data 层依赖集合。
var DataProviderSet = wire.NewSet(
	ProvideConsul,
	ProvideRedis,
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
	return c.Data.Redis
}

func ProvideConsul(c *conf.Bootstrap) *common.Consul {
	return c.Data.Consul
}

func ProvideUserAuthClient(userClient *rpc.UserClient) userv1.AuthServiceClient {
	return userClient.Auth
}
