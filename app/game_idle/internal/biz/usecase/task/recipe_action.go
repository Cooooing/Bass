package task

import (
	"common/pkg/apperror"
	"common/pkg/client/timewheel"
	"common/pkg/constant"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/biz/usecase"
	"game_idle/internal/enum"
	"log/slog"
	"time"
)

// RecipeActionTask 构建基于配方结算的行动任务。
type RecipeActionTask struct {
	logger          *slog.Logger
	recipeRepo      repo.RecipeRepo
	backpackRepo    repo.BackpackRepo
	settlementRepo  repo.ActionSettlementRepo
	actionQueueRepo repo.ActionQueueRepo
	recipeUsecase   *usecase.RecipeUsecase
}

func NewRecipeActionTask(
	logger *slog.Logger,
	recipeRepo repo.RecipeRepo,
	backpackRepo repo.BackpackRepo,
	settlementRepo repo.ActionSettlementRepo,
	actionQueueRepo repo.ActionQueueRepo,
	recipeUsecase *usecase.RecipeUsecase,
) *RecipeActionTask {
	return &RecipeActionTask{
		logger:          logger,
		recipeRepo:      recipeRepo,
		backpackRepo:    backpackRepo,
		settlementRepo:  settlementRepo,
		actionQueueRepo: actionQueueRepo,
		recipeUsecase:   recipeUsecase,
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
			queue, err := t.actionQueueRepo.Load(jobCtx, req.CharacterID)
			if err != nil {
				return err
			}
			if len(queue.Items) == 0 || queue.Items[0].ID != task.TaskID {
				return nil
			}

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
			// 原子变更是最终防线，任意物品变更后为负数则整次结算失败。
			// TODO 后续在这里接入钓鱼速度、产量、稀有率等 Buff 对结算的影响。
			settlement, err := t.settlementRepo.Apply(jobCtx, &repo.ActionSettlementReq{
				CharacterID: req.CharacterID,
				Items:       items,
				AbilityID:   enum.Ability(req.Action.AbilityID),
				ExpReward:   req.Action.ExpReward,
			})
			if err != nil {
				if code, ok := apperror.BusinessCode(err); ok && code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_BACKPACK_INSUFFICIENT {
					stopReason = enum.ActionStopReasonInsufficientItems
					settlement = &model.ActionSettlement{}
				} else if settlement == nil {
					return err
				} else {
					t.logger.ErrorContext(jobCtx, "game idle action settlement persist failed", constant.LogFieldErr, err, "character_id", req.CharacterID)
				}
			}
			select {
			case <-jobCtx.Done():
				return jobCtx.Err()
			case req.PendingTasks <- &usecase.PendingActionTask{
				CharacterID:      req.CharacterID,
				TaskID:           task.TaskID,
				ActionID:         task.ActionID,
				StopReason:       stopReason,
				StartedAt:        req.Now,
				CompletedAt:      time.Now(),
				ItemChanges:      settlement.ItemChanges,
				AbilityChanges:   settlement.AbilityChanges,
				AbilityLeveledUp: settlement.AbilityLeveledUp,
			}:
				return nil
			}
		},
	}, nil
}
