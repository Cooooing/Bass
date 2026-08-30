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
	ProvideStateEventHandlers,
	ProvideActionTasks,
	usecase.NewCharacterUsecase,
	usecase.NewBackpackUsecase,
	usecase.NewCharacterAbilityUsecase,
	usecase.NewCharacterStateUsecase,
	usecase.NewGameIdleEventUsecase,
	usecase.NewActionSettlementHandler,
	usecase.NewActionQueueCommandHandler,
	usecase.NewStateEngine,
	usecase.NewActionQueueLocker,
	usecase.NewRecipeUsecase,
	usecase.NewMetadataCacheUsecase,
	usecase.NewActionSchedulerUsecase,
	usecase.NewActionQueueUsecase,
	usecase.NewChatUsecase,
	usecase.NewRegionUsecase,
	usecase.NewActionUsecase,
	usecase.NewItemUsecase,
)

func ProvideStateCommandHandlers(
	actionSettlementHandler *usecase.ActionSettlementHandler,
	actionQueueCommandHandler *usecase.ActionQueueCommandHandler,
) map[usecase.StateCommandType]usecase.StateCommandHandler {
	return map[usecase.StateCommandType]usecase.StateCommandHandler{
		usecase.StateCommandTypeActionSettlement:        actionSettlementHandler,
		usecase.StateCommandTypeActionQueueAdd:          actionQueueCommandHandler,
		usecase.StateCommandTypeActionQueueMove:         actionQueueCommandHandler,
		usecase.StateCommandTypeActionQueueRemove:       actionQueueCommandHandler,
		usecase.StateCommandTypeActionQueueClear:        actionQueueCommandHandler,
		usecase.StateCommandTypeActionQueueCompleteHead: actionQueueCommandHandler,
	}
}

func ProvideStateEventHandlers() map[usecase.StateEventType][]usecase.StateEventHandler {
	return map[usecase.StateEventType][]usecase.StateEventHandler{}
}

func ProvideActionTasks(
	recipeActionTask *task.RecipeActionTask,
) map[enum.ActionKind]usecase.ActionTask {
	return map[enum.ActionKind]usecase.ActionTask{
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
