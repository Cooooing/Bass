package task

import (
	"common/pkg/apperror"
	"common/pkg/client/timewheel"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/biz/usecase"
	"game_idle/internal/enum"
	"time"
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
	// 构建任务时先汇总所有配方输入，避免每个配方单独查询背包造成重复读。
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
	if len(inputQuantities) > 0 {
		// 行动每一轮进入时间轮前只校验消耗条件，不冻结也不扣减。
		sufficient, err := t.backpackRepo.CheckItems(ctx, &repo.BackpackCheckReq{
			CharacterID: req.CharacterID,
			Items:       inputQuantities,
		})
		if err != nil {
			return nil, err
		}
		if !sufficient {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_BACKPACK_INSUFFICIENT)
		}
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
			stopReason := enum.ActionStopReasonNone
			if len(inputQuantities) > 0 {
				// 结算前再次校验，真正扣减会在状态结算事务里和产出一起原子提交。
				sufficient, err := t.backpackRepo.CheckItems(jobCtx, &repo.BackpackCheckReq{
					CharacterID: req.CharacterID,
					Items:       inputQuantities,
				})
				if err != nil {
					return err
				}
				if !sufficient {
					select {
					case <-jobCtx.Done():
						return jobCtx.Err()
					case req.PendingTasks <- &usecase.PendingCharacterActionTask{
						CharacterID: req.CharacterID,
						TaskID:      task.TaskID,
						ActionID:    task.ActionID,
						StopReason:  enum.ActionStopReasonInsufficientItems,
					}:
						return nil
					}
				}
			}

			// 产出先按配方聚合，最终由结算命令一次性应用到角色状态。
			outputQuantities := make(map[string]int64, len(req.Action.Recipes))
			for _, relation := range req.Action.Recipes {
				itemQuantities, err := t.recipeUsecase.RollNormal(jobCtx, &usecase.RollRecipeReq{
					RecipeID: relation.RecipeID,
				})
				if err != nil {
					return err
				}
				for itemID, quantity := range itemQuantities {
					outputQuantities[itemID] += quantity
				}
			}

			items := make([]*model.CharacterBackpackItemChange, 0, len(inputQuantities)+len(outputQuantities))
			for itemID, quantity := range inputQuantities {
				items = append(items, &model.CharacterBackpackItemChange{
					ItemID:   itemID,
					Quantity: -quantity,
				})
			}
			for itemID, quantity := range outputQuantities {
				items = append(items, &model.CharacterBackpackItemChange{
					ItemID:   itemID,
					Quantity: quantity,
				})
			}
			select {
			case <-jobCtx.Done():
				return jobCtx.Err()
			case req.PendingTasks <- &usecase.PendingCharacterActionTask{
				CharacterID: req.CharacterID,
				TaskID:      task.TaskID,
				ActionID:    task.ActionID,
				StopReason:  stopReason,
				Items:       items,
				AbilityID:   enum.Ability(req.Action.AbilityID),
				ExpReward:   req.Action.ExpReward,
				StartedAt:   req.Now,
				CompletedAt: time.Now(),
			}:
				return nil
			}
		},
	}, nil
}
