package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	"game_town/internal/config"
	"game_town/internal/data/client"
	"game_town/internal/data/repo"

	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	client.ProvideTx,
	client.NewAgentClient,
	client.NewEmbeddingClient,
	ProvideConsul,
	commonClient.NewObservability,
	commonClient.NewConsulClient,
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

func ProvideConsul(
	c *config.Bootstrap,
) *common.Consul {
	return c.Consul
}
