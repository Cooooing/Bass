package data

import (
	commonClient "common/pkg/client"
	"common/pkg/client/timewheel"
	"common/proto/gen/common"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/config"
	"game_idle/internal/data/client"
	"game_idle/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet 提供微服务模式的数据层依赖。
var DataProviderSet = wire.NewSet(
	ModuleProviderSet,
	commonClient.NewNatsClient,
	ProvideConsul,
	commonClient.NewConsulClient,
)

// ModuleProviderSet 提供不依赖服务发现的模块数据层依赖。
var ModuleProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	timewheel.NewTimeWheel,
	ProvideNats,
	ProvideTimeWheel,
	repo.NewCharacterStateCache,
	ProvideCharacterStateRepo,
	client.ProvideTx,
	repo.NewCharacterRepo,
	repo.NewCharacterSessionRepo,
	repo.NewGameIdleEventRepo,
	repo.NewCharacterBackpackRepo,
	repo.NewCharacterAbilityRepo,
	repo.NewMetaItemRepo,
	repo.NewMetaRecipeRepo,
	repo.NewMetaActionRepo,
	repo.NewMetaRegionRepo,
	repo.NewCharacterActionQueueRepo,
	repo.NewChatMessageRepo,
)

func ProvideTimeWheel(c *config.Bootstrap) *common.TimeWheel {
	return c.GetTimewheel()
}

func ProvideCharacterStateRepo(cache *repo.CharacterStateCache) bizrepo.CharacterStateRepo {
	return cache
}

func ProvideNats(c *config.Bootstrap) *common.Nats {
	return c.Nats
}

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}
