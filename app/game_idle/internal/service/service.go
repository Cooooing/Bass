package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewCharacterService,
	NewCharacterAbilityService,
	NewBackpackService,
	NewActionQueueService,
	NewChatService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	characterService *CharacterService,
	characterAbilityService *CharacterAbilityService,
	backpackService *BackpackService,
	actionQueueService *ActionQueueService,
	chatService *ChatService,
) []server.Service {
	return []server.Service{
		commonSystemService,
		characterService,
		characterAbilityService,
		backpackService,
		actionQueueService,
		chatService,
	}
}
