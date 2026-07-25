package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewAgentConfigService,
	NewPlayerService,
	NewWorldService,
	NewWorldMemberService,
	NewNpcService,
	NewLocationService,
	NewFactionService,
	NewEventService,
)

func ProvideServices(system *CommonSystemService, agentConfig *AgentConfigService, player *PlayerService, world *WorldService, worldMember *WorldMemberService, npc *NpcService, location *LocationService, faction *FactionService, event *EventService) []server.Service {
	return []server.Service{
		system,
		agentConfig,
		player,
		world,
		worldMember,
		npc,
		location,
		faction,
		event,
	}
}
