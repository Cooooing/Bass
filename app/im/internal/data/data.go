package data

import (
	commonClient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"

	"im/internal/config"
	"im/internal/data/client"
	"im/internal/data/gen"
	"im/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet 提供微服务模式的数据层依赖。
var DataProviderSet = wire.NewSet(
	ModuleProviderSet,
	commonClient.NewObservability,
	ProvideConsul,
	commonClient.NewConsulClient,
	rpc.ProvidePlatformClient,
)

// ModuleProviderSet 提供不依赖服务发现的模块数据层依赖。
var ModuleProviderSet = wire.NewSet(
	commonClient.NewSQLClient,
	ProvideDatabase,
	client.NewEntClient,
	ProvideGenClient,
	repo.NewChatGroupRepo,
	repo.NewChatGroupMemberRepo,
	repo.NewChatSessionRepo,
	repo.NewChatMessageRepo,
	repo.NewAssetClient,
)

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}

func ProvideDatabase(c *config.Bootstrap) *common.Database {
	return c.Database
}

func ProvideGenClient(c *client.EntClient) *gen.Client {
	return c.Client
}
