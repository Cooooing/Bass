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
	ProvideConsul,
	commonClient.NewObservability,
	commonClient.NewConsulClient,
	repo.NewPlayerRepo,
	repo.NewAgentConfigRepo,
	repo.NewSessionRepo,
	repo.NewWorldRepo,
	repo.NewWorldMemberRepo,
	repo.NewLocationRepo,
	repo.NewNpcRepo,
	repo.NewCommandRepo,
	repo.NewEventRepo,
	repo.NewWorldStateSnapshotRepo,
	repo.NewWorldMetricDefinitionRepo,
	repo.NewMemoryRepo,
	repo.NewRelationshipRepo,
	repo.NewAgentRunRepo,
)

func ProvideConsul(c *config.Bootstrap) *common.Consul { return c.Consul }
