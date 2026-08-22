package data

import (
	"bbs/internal/config"
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	userv1 "common/proto/gen/user/v1"

	"github.com/google/wire"
)

// DataProviderSet 提供微服务模式的数据层依赖。
var DataProviderSet = wire.NewSet(
	ModuleProviderSet,
	ProvideConsul,
	commonClient.NewConsulClient,
	rpc.ProvideUserClient,
	rpc.ProvideContentClient,
	rpc.ProvideEconomyClient,
	rpc.ProvideNotifyClient,
	rpc.ProvidePlatformClient,
)

// ModuleProviderSet 提供不依赖服务发现的模块数据层依赖。
var ModuleProviderSet = wire.NewSet(
	ProvideUserAuthClient,
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

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}

func ProvideUserAuthClient(userClient *rpc.UserClient) userv1.AuthServiceClient {
	return userClient.Auth
}
