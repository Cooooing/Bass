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
	NewCharacterBackpackService,
	NewCharacterActionQueueService,
	NewChatService,
	NewMetaRegionService,
	NewMetaActionService,
	NewMetaItemService,
	NewMetaRecipeService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	characterService *CharacterService,
	characterAbilityService *CharacterAbilityService,
	backpackService *CharacterBackpackService,
	actionQueueService *CharacterActionQueueService,
	chatService *ChatService,
	regionService *MetaRegionService,
	actionService *MetaActionService,
	itemService *MetaItemService,
	recipeService *MetaRecipeService,
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
		recipeService,
	}
}
