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
	recipeRepo    repo.RecipeRepo
	backpackRepo  repo.BackpackRepo
	recipeUsecase *usecase.RecipeUsecase
	stateEngine   *usecase.StateEngine
}

func NewRecipeActionTask(
	recipeRepo repo.RecipeRepo,
	backpackRepo repo.BackpackRepo,
	recipeUsecase *usecase.RecipeUsecase,
	stateEngine *usecase.StateEngine,
) *RecipeActionTask {
	return &RecipeActionTask{
		recipeRepo:    recipeRepo,
		backpackRepo:  backpackRepo,
		recipeUsecase: recipeUsecase,
		stateEngine:   stateEngine,
	}
}

func (t *RecipeActionTask) BuildTask(ctx context.Context, req *usecase.BuildActionTaskReq) (*timewheel.Task, error) {
	inputQuantities := make(map[string]int64, len(req.Action.Recipes))
	for _, relation := range req.Action.Recipes {
		recipe, err := t.recipeRepo.Get(ctx, relation.RecipeID)
		if err != nil {
			return nil, err
		}
		for _, input := range recipe.Inputs {
			inputQuantities[input.ItemID] += input.Quantity
		}
	}
	if len(inputQuantities) > 0 {
		// 行动每一轮进入时间轮前都校验消耗条件。
		if err := t.backpackRepo.CheckItems(ctx, &repo.BackpackCheckReq{
			CharacterID: req.CharacterID,
			Items:       inputQuantities,
		}); err != nil {
			return nil, err
		}
	}

	task := &model.ActionTask{
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
				// 结算前再次校验，避免等待期间背包被其他链路消耗。
				if err := t.backpackRepo.CheckItems(jobCtx, &repo.BackpackCheckReq{
					CharacterID: req.CharacterID,
					Items:       inputQuantities,
				}); err != nil {
					if code, ok := apperror.BusinessCode(err); !ok || code != cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_BACKPACK_INSUFFICIENT {
						return err
					}
					select {
					case <-jobCtx.Done():
						return jobCtx.Err()
					case req.PendingTasks <- &usecase.PendingActionTask{
						CharacterID: req.CharacterID,
						TaskID:      task.TaskID,
						ActionID:    task.ActionID,
						StopReason:  enum.ActionStopReasonInsufficientItems,
					}:
						return nil
					}
				}
			}

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

			items := make([]*model.BackpackItemChange, 0, len(inputQuantities)+len(outputQuantities))
			for itemID, quantity := range inputQuantities {
				items = append(items, &model.BackpackItemChange{
					ItemID:   itemID,
					Quantity: -quantity,
				})
			}
			for itemID, quantity := range outputQuantities {
				items = append(items, &model.BackpackItemChange{
					ItemID:   itemID,
					Quantity: quantity,
				})
			}
			// 状态机同步完成扣物品、加产物、加经验等核心状态变化。
			// TODO 后续在命令里接入钓鱼速度、产量、稀有率等 Buff 对结算的影响。
			changeSet, err := t.stateEngine.Apply(jobCtx, &usecase.ActionSettlementCommand{
				CharacterID: req.CharacterID,
				ActionID:    task.ActionID,
				Items:       items,
				AbilityID:   enum.Ability(req.Action.AbilityID),
				ExpReward:   req.Action.ExpReward,
			})
			if err != nil {
				if code, ok := apperror.BusinessCode(err); ok && code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_BACKPACK_INSUFFICIENT {
					stopReason = enum.ActionStopReasonInsufficientItems
					changeSet = &usecase.StateChangeSet{}
				} else {
					return err
				}
			}
			select {
			case <-jobCtx.Done():
				return jobCtx.Err()
			case req.PendingTasks <- &usecase.PendingActionTask{
				CharacterID:  req.CharacterID,
				TaskID:       task.TaskID,
				ActionID:     task.ActionID,
				StopReason:   stopReason,
				StartedAt:    req.Now,
				CompletedAt:  time.Now(),
				StateChanges: changeSet,
			}:
				return nil
			}
		},
	}, nil
}
