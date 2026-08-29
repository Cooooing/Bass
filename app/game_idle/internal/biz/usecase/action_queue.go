package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"fmt"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/enum"
	"time"
)

// ActionQueueUsecase 处理玩家对行动队列的主动编辑。
type ActionQueueUsecase struct {
	characterRepo   repo.CharacterRepo
	actionRepo      repo.ActionRepo
	actionQueueRepo repo.ActionQueueRepo
	abilityUsecase  *CharacterAbilityUsecase
	eventUsecase    *GameIdleEventUsecase
	scheduler       *ActionSchedulerUsecase
	locker          *ActionQueueLocker
	actionTasks     map[enum.ActionKind]ActionTask
}

func NewActionQueueUsecase(
	characterRepo repo.CharacterRepo,
	actionRepo repo.ActionRepo,
	actionQueueRepo repo.ActionQueueRepo,
	abilityUsecase *CharacterAbilityUsecase,
	eventUsecase *GameIdleEventUsecase,
	scheduler *ActionSchedulerUsecase,
	locker *ActionQueueLocker,
	actionTasks map[enum.ActionKind]ActionTask,
) *ActionQueueUsecase {
	return &ActionQueueUsecase{
		characterRepo:   characterRepo,
		actionRepo:      actionRepo,
		actionQueueRepo: actionQueueRepo,
		abilityUsecase:  abilityUsecase,
		eventUsecase:    eventUsecase,
		scheduler:       scheduler,
		locker:          locker,
		actionTasks:     actionTasks,
	}
}

type ListActionQueueResp struct {
	Queue *model.ActionQueue
}

func (u *ActionQueueUsecase) List(ctx context.Context, characterID int64) (*ListActionQueueResp, error) {
	queue, err := u.actionQueueRepo.Load(ctx, characterID)
	if err != nil {
		return nil, err
	}
	return &ListActionQueueResp{Queue: queue}, nil
}

func (u *ActionQueueUsecase) Persist(ctx context.Context, characterID int64) error {
	return u.actionQueueRepo.Persist(ctx, characterID)
}

type AddActionReq struct {
	CharacterID int64
	ActionID    string
	Times       int64 // Times 表示执行次数，-1 表示无限执行，直到条件不满足或玩家主动调整队列。
	Position    *int32
}

func (u *ActionQueueUsecase) Add(ctx context.Context, req *AddActionReq) error {
	character, err := u.characterRepo.Get(ctx, req.CharacterID)
	if err != nil {
		return err
	}
	if character.Status != enum.CharacterStatusActive {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_INVALID)
	}
	actionConfig, err := u.actionRepo.Get(ctx, req.ActionID)
	if err != nil {
		return err
	}
	if !actionConfig.Enabled || req.Times == 0 || req.Times < -1 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
	}
	if u.actionTasks[actionConfig.ActionKind] == nil {
		return fmt.Errorf("game idle action task is required: %s", actionConfig.ActionKind)
	}
	if err = u.abilityUsecase.CheckLevel(
		ctx,
		req.CharacterID,
		enum.Ability(actionConfig.AbilityID),
		actionConfig.RequiredAbilityLevel,
	); err != nil {
		return err
	}

	return u.locker.withCharacterLock(req.CharacterID, func() error {
		return u.editQueue(ctx, req.CharacterID, func(queue *model.ActionQueue) error {
			if len(queue.Items) >= int(character.ActionQueueCapacity) {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_QUEUE_FULL)
			}
			position := len(queue.Items)
			if req.Position != nil {
				position = int(*req.Position)
			}
			if position < 0 || position > len(queue.Items) {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
			}

			now := time.Now()
			queueItem := &model.ActionQueueItem{
				ID:        fmt.Sprintf("character:%d:action:%s:at:%d", req.CharacterID, req.ActionID, now.UnixNano()),
				ActionID:  req.ActionID,
				Times:     req.Times,
				CreatedAt: now,
			}
			queue.Items = append(queue.Items, nil)
			copy(queue.Items[position+1:], queue.Items[position:])
			queue.Items[position] = queueItem
			return nil
		})
	})
}

type MoveActionReq struct {
	CharacterID     int64
	CurrentPosition int32
	TargetPosition  int32
}

func (u *ActionQueueUsecase) Move(ctx context.Context, req *MoveActionReq) error {
	return u.locker.withCharacterLock(req.CharacterID, func() error {
		return u.editQueue(ctx, req.CharacterID, func(queue *model.ActionQueue) error {
			currentPosition := int(req.CurrentPosition)
			targetPosition := int(req.TargetPosition)
			if currentPosition < 0 || targetPosition < 0 || currentPosition >= len(queue.Items) || targetPosition >= len(queue.Items) {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
			}
			queueItem := queue.Items[currentPosition]
			queue.Items = append(queue.Items[:currentPosition], queue.Items[currentPosition+1:]...)
			if targetPosition >= len(queue.Items) {
				queue.Items = append(queue.Items, queueItem)
				return nil
			}
			queue.Items = append(queue.Items, nil)
			copy(queue.Items[targetPosition+1:], queue.Items[targetPosition:])
			queue.Items[targetPosition] = queueItem
			return nil
		})
	})
}

type RemoveActionReq struct {
	CharacterID int64
	Position    int32
}

func (u *ActionQueueUsecase) Remove(ctx context.Context, req *RemoveActionReq) error {
	return u.locker.withCharacterLock(req.CharacterID, func() error {
		return u.editQueue(ctx, req.CharacterID, func(queue *model.ActionQueue) error {
			if len(queue.Items) == 0 || req.Position < 0 || int(req.Position) >= len(queue.Items) {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
			}
			position := int(req.Position)
			queue.Items = append(queue.Items[:position], queue.Items[position+1:]...)
			return nil
		})
	})
}

func (u *ActionQueueUsecase) Clear(ctx context.Context, characterID int64) error {
	return u.locker.withCharacterLock(characterID, func() error {
		return u.editQueue(ctx, characterID, func(queue *model.ActionQueue) error {
			queue.Items = make([]*model.ActionQueueItem, 0)
			return nil
		})
	})
}

func (u *ActionQueueUsecase) editQueue(
	ctx context.Context,
	characterID int64,
	edit func(queue *model.ActionQueue) error,
) error {
	queue, err := u.actionQueueRepo.Load(ctx, characterID)
	if err != nil {
		return err
	}
	oldHeadID := u.headTaskID(queue)
	if err = edit(queue); err != nil {
		return err
	}
	headChanged := oldHeadID != u.headTaskID(queue)
	if headChanged && oldHeadID != "" {
		u.scheduler.stopCurrent(ctx, characterID)
	}
	if err = u.actionQueueRepo.Save(ctx, queue); err != nil {
		if headChanged && oldHeadID != "" {
			_ = u.scheduler.startCurrent(ctx, characterID)
		}
		return err
	}
	u.eventUsecase.PublishActionQueueUpdated(ctx, queue, ActionQueueUpdateReasonManualChanged)
	if headChanged && len(queue.Items) > 0 {
		return u.scheduler.startCurrent(ctx, characterID)
	}
	return nil
}

func (u *ActionQueueUsecase) headTaskID(queue *model.ActionQueue) string {
	if len(queue.Items) == 0 {
		return ""
	}
	return queue.Items[0].ID
}
