package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewCommonSystemService,
	NewPlayerService,
	NewSessionService,
	NewWorldService,
	NewAgentConfigService,
	NewCommandService,
	NewNpcService,
	NewMemoryService,
	NewEventService,
	ProvideGrpcServices,
	ProvideHttpServices,
)

func ProvideGrpcServices(commonSystemService *CommonSystemService, playerService *PlayerService, sessionService *SessionService, worldService *WorldService, agentConfigService *AgentConfigService, commandService *CommandService, npcService *NpcService, memoryService *MemoryService, eventService *EventService) []server.GrpcService {
	return []server.GrpcService{commonSystemService, playerService, sessionService, worldService, agentConfigService, commandService, npcService, memoryService, eventService}
}

func ProvideHttpServices(commonSystemService *CommonSystemService) []server.HttpService {
	return []server.HttpService{commonSystemService}
}
