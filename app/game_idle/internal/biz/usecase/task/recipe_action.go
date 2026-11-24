package task

import (
	"common/pkg/apperror"
	"common/pkg/client/timewheel"
	cerrors "common/proto/gen/common/errors/v1"
	"context"
	"time"

	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/biz/usecase"
	"game_idle/internal/enum"
)

// RecipeActionTask 构建基于配方结算的行动任务。
type RecipeActionTask struct {
	recipeRepo    repo.MetaRecipeRepo
	backpackRepo  repo.CharacterBackpackRepo
	recipeUsecase *usecase.MetaRecipeUsecase
}

func NewRecipeActionTask(
	recipeRepo repo.MetaRecipeRepo,
	backpackRepo repo.CharacterBackpackRepo,
	recipeUsecase *usecase.MetaRecipeUsecase,
) *RecipeActionTask {
	return &RecipeActionTask{
		recipeRepo:    recipeRepo,
		backpackRepo:  backpackRepo,
		recipeUsecase: recipeUsecase,
	}
}

func (t *RecipeActionTask) BuildTask(ctx context.Context, req *usecase.BuildCharacterActionTaskReq) (*timewheel.Task, error) {
	inputQuantities, err := t.collectInputQuantities(ctx, req)
	if err != nil {
		return nil, err
	}
	if err = t.checkInputQuantities(ctx, req.CharacterID, inputQuantities); err != nil {
		return nil, err
	}

	task := &model.CharacterActionTask{
		TaskID:      req.QueueItem.ID,
		CharacterID: req.CharacterID,
		ActionID:    req.QueueItem.ActionID,
		DueAt:       req.Now.Add(req.Action.Duration),
	}
	return &timewheel.Task{
		ID:      task.TaskID,
		DueAt:   task.DueAt,
		Payload: task,
		Job: func(jobCtx context.Context, item *timewheel.Task) error {
			return t.executeTask(jobCtx, req, task, inputQuantities)
		},
	}, nil
}

// collectInputQuantities 汇总行动内全部配方的消耗物品。
func (t *RecipeActionTask) collectInputQuantities(ctx context.Context, req *usecase.BuildCharacterActionTaskReq) (map[string]int64, error) {
	inputQuantities := make(map[string]int64, len(req.Action.Recipes))
	for _, relation := range req.Action.Recipes {
		recipe, err := t.recipeRepo.Get(ctx, relation.RecipeID)
		if err != nil {
			return nil, err
		}
		if recipe == nil {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_RECIPE_INVALID)
		}
		for _, input := range recipe.Inputs {
			inputQuantities[input.ItemID] += input.Quantity
		}
	}
	return inputQuantities, nil
}

// checkInputQuantities 只校验行动消耗条件，不冻结也不扣减物品。
func (t *RecipeActionTask) checkInputQuantities(ctx context.Context, characterID int64, inputQuantities map[string]int64) error {
	if len(inputQuantities) == 0 {
		return nil
	}
	sufficient, err := t.backpackRepo.CheckItems(ctx, &repo.BackpackCheckReq{
		CharacterID: characterID,
		Items:       inputQuantities,
	})
	if err != nil {
		return err
	}
	if !sufficient {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_BACKPACK_INSUFFICIENT)
	}
	return nil
}

// executeTask 在时间轮触发后复核消耗条件、汇总产出并提交结算请求。
func (t *RecipeActionTask) executeTask(
	ctx context.Context,
	req *usecase.BuildCharacterActionTaskReq,
	task *model.CharacterActionTask,
	inputQuantities map[string]int64,
) error {
	if err := t.checkInputQuantities(ctx, req.CharacterID, inputQuantities); err != nil {
		if code, ok := apperror.BusinessCode(err); ok && code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_BACKPACK_INSUFFICIENT {
			return t.sendInsufficientItemsTask(ctx, req, task)
		}
		return err
	}
	outputQuantities, err := t.collectOutputQuantities(ctx, req)
	if err != nil {
		return err
	}
	return t.sendSettlementTask(ctx, req, task, inputQuantities, outputQuantities)
}

// collectOutputQuantities 按配方聚合本次行动的产出物品。
func (t *RecipeActionTask) collectOutputQuantities(ctx context.Context, req *usecase.BuildCharacterActionTaskReq) (map[string]int64, error) {
	outputQuantities := make(map[string]int64, len(req.Action.Recipes))
	for _, relation := range req.Action.Recipes {
		itemQuantities, err := t.recipeUsecase.RollNormal(ctx, &usecase.RollRecipeReq{RecipeID: relation.RecipeID})
		if err != nil {
			return nil, err
		}
		for itemID, quantity := range itemQuantities {
			outputQuantities[itemID] += quantity
		}
	}
	return outputQuantities, nil
}

// sendInsufficientItemsTask 通知调度器当前队首因消耗不足而停止。
func (t *RecipeActionTask) sendInsufficientItemsTask(ctx context.Context, req *usecase.BuildCharacterActionTaskReq, task *model.CharacterActionTask) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case req.PendingTasks <- &usecase.PendingCharacterActionTask{
		CharacterID: req.CharacterID,
		TaskID:      task.TaskID,
		ActionID:    task.ActionID,
		StopReason:  enum.ActionStopReasonInsufficientItems,
	}:
		return nil
	}
}

// sendSettlementTask 通知调度器应用本次行动的物品和能力结算。
func (t *RecipeActionTask) sendSettlementTask(
	ctx context.Context,
	req *usecase.BuildCharacterActionTaskReq,
	task *model.CharacterActionTask,
	inputQuantities map[string]int64,
	outputQuantities map[string]int64,
) error {
	items := make([]*model.CharacterBackpackItemChange, 0, len(inputQuantities)+len(outputQuantities))
	for itemID, quantity := range inputQuantities {
		items = append(items, &model.CharacterBackpackItemChange{ItemID: itemID, Quantity: -quantity})
	}
	for itemID, quantity := range outputQuantities {
		items = append(items, &model.CharacterBackpackItemChange{ItemID: itemID, Quantity: quantity})
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	case req.PendingTasks <- &usecase.PendingCharacterActionTask{
		CharacterID: req.CharacterID,
		TaskID:      task.TaskID,
		ActionID:    task.ActionID,
		StopReason:  enum.ActionStopReasonNone,
		Items:       items,
		AbilityID:   enum.Ability(req.Action.AbilityID),
		ExpReward:   req.Action.ExpReward,
		StartedAt:   req.Now,
		CompletedAt: time.Now(),
	}:
		return nil
	}
}
