package data

import (
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	userv1 "common/proto/gen/user/v1"

	"bff_bbs/internal/config"

	"github.com/google/wire"
)

// DataProviderSet 提供微服务模式的数据层依赖。
var DataProviderSet = wire.NewSet(
	ModuleProviderSet,
	commonClient.NewObservability,
	commonClient.NewRedisClient,
	ProvideConsul,
	ProvideRedis,
	commonClient.NewConsulClient,
	rpc.ProvideUserClient,
	rpc.ProvideContentClient,
	rpc.ProvideEconomyClient,
	rpc.ProvideNotifyClient,
	rpc.ProvidePlatformClient,
	rpc.ProvideRealtimeHubClient,
	rpc.ProvideIMClient,
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
	NewIPResolutionClient,
	NewLocationClient,
	NewRelationClient,
	NewOtpClient,
	NewLivenessClient,
	NewContentArticleClient,
	NewArticleSearchRateLimitCache,
	NewContentPostscriptClient,
	NewContentCommentClient,
	NewContentDomainClient,
	NewContentTagClient,
	NewContentMoonbreezeClient,
	NewNotificationClient,
	NewHolidayRepo,
)

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}

func ProvideRedis(c *config.Bootstrap) *common.Redis {
	return c.Redis
}

func ProvideUserAuthClient(userClient *rpc.UserClient) userv1.AuthServiceClient {
	return userClient.Auth
}
