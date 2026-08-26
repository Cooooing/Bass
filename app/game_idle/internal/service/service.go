package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewCharacterService,
	NewBackpackService,
	NewActionQueueService,
	NewWebSocketService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	characterService *CharacterService,
	backpackService *BackpackService,
	actionQueueService *ActionQueueService,
	webSocketService *WebSocketService,
) []server.Service {
	return []server.Service{
		commonSystemService,
		characterService,
		backpackService,
		actionQueueService,
		webSocketService,
	}
}
