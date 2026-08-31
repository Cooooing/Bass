package biz

import (
	"game_idle/internal/biz/usecase"
	"game_idle/internal/biz/usecase/task"
	"game_idle/internal/enum"

	"github.com/google/wire"
)

// BizProviderSet 提供业务层依赖集合。
var BizProviderSet = wire.NewSet(
	task.NewRecipeActionTask,
	ProvideStateCommandHandlers,
	ProvideActionTasks,
	usecase.NewCharacterUsecase,
	usecase.NewCharacterBackpackUsecase,
	usecase.NewCharacterAbilityUsecase,
	usecase.NewCharacterStateUsecase,
	usecase.NewGameIdleEventUsecase,
	usecase.NewCharacterSettlementHandler,
	usecase.NewCharacterActionQueueCommandHandler,
	usecase.NewMetaRecipeUsecase,
	usecase.NewCharacterActionSchedulerUsecase,
	usecase.NewCharacterActionQueueUsecase,
	usecase.NewChatUsecase,
	usecase.NewMetaRegionUsecase,
	usecase.NewMetaActionUsecase,
	usecase.NewMetaItemUsecase,
)

func ProvideStateCommandHandlers(
	actionSettlementHandler *usecase.CharacterSettlementHandler,
	actionQueueCommandHandler *usecase.CharacterActionQueueCommandHandler,
) map[enum.StateCommandType]usecase.CharacterStateCommandHandler {
	return map[enum.StateCommandType]usecase.CharacterStateCommandHandler{
		enum.StateCommandTypeActionSettlement:        actionSettlementHandler,
		enum.StateCommandTypeActionQueueAdd:          actionQueueCommandHandler,
		enum.StateCommandTypeActionQueueMove:         actionQueueCommandHandler,
		enum.StateCommandTypeActionQueueRemove:       actionQueueCommandHandler,
		enum.StateCommandTypeActionQueueClear:        actionQueueCommandHandler,
		enum.StateCommandTypeActionQueueCompleteHead: actionQueueCommandHandler,
	}
}

func ProvideActionTasks(
	recipeActionTask *task.RecipeActionTask,
) map[enum.ActionKind]usecase.CharacterActionTaskBuilder {
	return map[enum.ActionKind]usecase.CharacterActionTaskBuilder{
		enum.ActionKindWoodcutting: recipeActionTask,
		enum.ActionKindForaging:    recipeActionTask,
		enum.ActionKindMining:      recipeActionTask,
		enum.ActionKindFishing:     recipeActionTask,
		enum.ActionKindCrafting:    recipeActionTask,
		enum.ActionKindSewing:      recipeActionTask,
		enum.ActionKindSmithing:    recipeActionTask,
		enum.ActionKindCooking:     recipeActionTask,
		enum.ActionKindEnhancing:   recipeActionTask,
		enum.ActionKindAlchemy:     recipeActionTask,
	}
}
