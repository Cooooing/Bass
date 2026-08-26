package data

import (
	commonClient "common/pkg/client"
	"common/pkg/client/timewheel"
	"common/proto/gen/common"
	"game_idle/internal/config"
	"game_idle/internal/data/client"
	"game_idle/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet 提供微服务模式的数据层依赖。
var DataProviderSet = wire.NewSet(
	ModuleProviderSet,
	commonClient.NewRedisClient,
	ProvideConsul,
	commonClient.NewConsulClient,
)

// ModuleProviderSet 提供不依赖服务发现的模块数据层依赖。
var ModuleProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	timewheel.NewTimeWheel,
	ProvideRedis,
	ProvideTimeWheel,
	client.ProvideTx,
	repo.NewCharacterRepo,
	repo.NewBackpackRepo,
	repo.NewItemRepo,
	repo.NewRecipeRepo,
	repo.NewActionRepo,
	repo.NewActionQueueRepo,
)

func ProvideTimeWheel(c *config.Bootstrap) *common.TimeWheel {
	return c.GetTimewheel()
}

func ProvideRedis(c *config.Bootstrap) *common.Redis {
	return c.Redis
}

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}
