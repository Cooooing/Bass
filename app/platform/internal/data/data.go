package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	"platform/internal/config"
	"platform/internal/data/client"
	"platform/internal/data/oss"
	"platform/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet 提供微服务模式的数据层依赖。
var DataProviderSet = wire.NewSet(
	ModuleProviderSet,
	commonClient.NewObservability,
	ProvideRedis,
	commonClient.NewRedisClient,
	ProvideNats,
	commonClient.NewNatsClient,
	ProvideConsul,
	commonClient.NewConsulClient,
)

// ModuleProviderSet 提供不依赖服务发现的模块数据层依赖。
var ModuleProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	commonClient.NewHttpClient,
	client.ProvideTx,
	repo.NewAssetRepo,
	repo.NewAssetUploadRateLimitCache,
	repo.NewAssetEventNatsRepo,
	oss.ProviderSet,
)

func ProvideNats(c *config.Bootstrap) *common.Nats {
	return c.Nats
}

func ProvideRedis(c *config.Bootstrap) *common.Redis {
	return c.Redis
}

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}
