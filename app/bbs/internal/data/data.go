package data

import (
	"bbs/internal/conf"
	"common/api/gen/common"
	userv1 "common/api/gen/user/v1"
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"

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
	NewAuthRepo,
	NewAccountRepo,
	NewPreferencesRepo,
	NewPrivacySettingRepo,
	NewLocationRepo,
	NewRelationRepo,
	NewTotpRepo,
	NewContentArticleRepo,
	NewContentPostscriptRepo,
	NewContentCommentRepo,
	NewContentDomainRepo,
	NewContentTagRepo,
	NewNotificationRepo,
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
