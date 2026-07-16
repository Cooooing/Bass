package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewPlayerService,
	NewSessionService,
	NewWorldService,
	NewAgentConfigService,
	NewCommandService,
	NewNpcService,
	NewMemoryService,
	NewEventService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	playerService *PlayerService,
	sessionService *SessionService,
	worldService *WorldService,
	agentConfigService *AgentConfigService,
	commandService *CommandService,
	npcService *NpcService,
	memoryService *MemoryService,
	eventService *EventService,
) []server.Service {
	return []server.Service{
		commonSystemService,
		playerService,
		sessionService,
		worldService,
		agentConfigService,
		commandService,
		npcService,
		memoryService,
		eventService,
	}
}
