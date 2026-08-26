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
	ProvideActionTasks,
	usecase.NewCharacterUsecase,
	usecase.NewBackpackUsecase,
	usecase.NewRecipeUsecase,
	usecase.NewActionQueueUsecase,
)

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
