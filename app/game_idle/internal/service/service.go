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
	NewRegionService,
	NewActionService,
	NewItemService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	characterService *CharacterService,
	characterAbilityService *CharacterAbilityService,
	backpackService *BackpackService,
	actionQueueService *ActionQueueService,
	chatService *ChatService,
	regionService *RegionService,
	actionService *ActionService,
	itemService *ItemService,
) []server.Service {
	return []server.Service{
		commonSystemService,
		characterService,
		characterAbilityService,
		backpackService,
		actionQueueService,
		chatService,
		regionService,
		actionService,
		itemService,
	}
}
