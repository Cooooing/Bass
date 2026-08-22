package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	"game_town/internal/config"
	"game_town/internal/data/client"
	"game_town/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet 提供微服务模式的数据层依赖。
var DataProviderSet = wire.NewSet(
	ModuleProviderSet,
	ProvideConsul,
	commonClient.NewConsulClient,
)

// ModuleProviderSet 提供不依赖服务发现的模块数据层依赖。
var ModuleProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	client.ProvideTx,
	client.NewAgentClient,
	client.NewEmbeddingClient,
	repo.NewEventNotifier,
	repo.NewAgentConfigRepo,
	repo.NewPlayerRepo,
	repo.NewWorldRepo,
	repo.NewWorldMemberRepo,
	repo.NewLocationRepo,
	repo.NewNpcRepo,
	repo.NewWorldStateRepo,
	repo.NewWorldRuleRepo,
	repo.NewEventRepo,
	repo.NewObservationRepo,
	repo.NewClaimRepo,
	repo.NewNpcBeliefRepo,
	repo.NewRelationshipRepo,
	repo.NewFactionRepo,
	repo.NewFactionMembershipRepo,
	repo.NewNpcMemoryRepo,
	repo.NewAgentJobRepo,
)

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}
