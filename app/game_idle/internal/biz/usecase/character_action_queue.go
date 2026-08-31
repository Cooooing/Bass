package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/enum"
	"time"
)

// CharacterActionQueueUsecase 处理玩家对行动队列的主动编辑。
type CharacterActionQueueUsecase struct {
	characterRepo   repo.CharacterRepo
	actionRepo      repo.MetaActionRepo
	actionQueueRepo repo.CharacterActionQueueRepo
	abilityUsecase  *CharacterAbilityUsecase
	eventUsecase    *GameIdleEventUsecase
	scheduler       *CharacterActionSchedulerUsecase
	stateUsecase    *CharacterStateUsecase
}

func NewCharacterActionQueueUsecase(
	characterRepo repo.CharacterRepo,
	actionRepo repo.MetaActionRepo,
	actionQueueRepo repo.CharacterActionQueueRepo,
	abilityUsecase *CharacterAbilityUsecase,
	eventUsecase *GameIdleEventUsecase,
	scheduler *CharacterActionSchedulerUsecase,
	stateUsecase *CharacterStateUsecase,
) *CharacterActionQueueUsecase {
	return &CharacterActionQueueUsecase{
		characterRepo:   characterRepo,
		actionRepo:      actionRepo,
		actionQueueRepo: actionQueueRepo,
		abilityUsecase:  abilityUsecase,
		eventUsecase:    eventUsecase,
		scheduler:       scheduler,
		stateUsecase:    stateUsecase,
	}
}

type ListActionQueueResp struct {
	Queue *model.CharacterActionQueue
}

func (u *CharacterActionQueueUsecase) List(ctx context.Context, characterID int64) (*ListActionQueueResp, error) {
	var queue *model.CharacterActionQueue
	err := u.stateUsecase.Execute(ctx, characterID, func(runCtx context.Context) error {
		var err error
		queue, err = u.actionQueueRepo.Load(runCtx, characterID)
		return err
	})
	return &ListActionQueueResp{Queue: queue}, err
}

func (u *CharacterActionQueueUsecase) Persist(ctx context.Context, characterID int64) error {
	return u.stateUsecase.Persist(ctx, characterID)
}

type AddActionReq struct {
	CharacterID int64
	ActionID    string
	Times       int64 // Times 表示执行次数，-1 表示无限执行，直到条件不满足或玩家主动调整队列。
	Position    *int32
}

func (u *CharacterActionQueueUsecase) Add(ctx context.Context, req *AddActionReq) error {
	character, err := u.characterRepo.Get(ctx, req.CharacterID)
	if err != nil {
		return err
	}
	if character == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_NOT_FOUND)
	}
	if character.Status != enum.CharacterStatusActive {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_INVALID)
	}
	actionConfig, err := u.actionRepo.Get(ctx, req.ActionID)
	if err != nil {
		return err
	}
	if actionConfig == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
	}
	if !actionConfig.Enabled || req.Times == 0 || req.Times < -1 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
	}
	command := &CharacterActionQueueAddCommand{
		CharacterID: req.CharacterID,
		ActionID:    req.ActionID,
		Times:       req.Times,
		Capacity:    character.ActionQueueCapacity,
		Position:    req.Position,
		Now:         time.Now(),
	}
	return u.stateUsecase.Execute(ctx, req.CharacterID, func(runCtx context.Context) error {
		if err := u.abilityUsecase.CheckLevel(
			runCtx,
			req.CharacterID,
			enum.Ability(actionConfig.AbilityID),
			actionConfig.RequiredAbilityLevel,
		); err != nil {
			return err
		}
		changeSet, err := u.stateUsecase.apply(runCtx, command)
		if err != nil {
			return err
		}
		return u.syncQueueChanged(runCtx, changeSet.QueueChanged)
	})
}

type MoveActionReq struct {
	CharacterID     int64
	CurrentPosition int32
	TargetPosition  int32
}

func (u *CharacterActionQueueUsecase) Move(ctx context.Context, req *MoveActionReq) error {
	command := &CharacterActionQueueMoveCommand{
		CharacterID:     req.CharacterID,
		CurrentPosition: req.CurrentPosition,
		TargetPosition:  req.TargetPosition,
	}
	_, err := u.stateUsecase.ExecuteCommand(ctx, command, func(runCtx context.Context, changeSet *CharacterStateChangeSet) error {
		return u.syncQueueChanged(runCtx, changeSet.QueueChanged)
	})
	return err
}

type RemoveActionReq struct {
	CharacterID int64
	Position    int32
}

func (u *CharacterActionQueueUsecase) Remove(ctx context.Context, req *RemoveActionReq) error {
	command := &CharacterActionQueueRemoveCommand{
		CharacterID: req.CharacterID,
		Position:    req.Position,
	}
	_, err := u.stateUsecase.ExecuteCommand(ctx, command, func(runCtx context.Context, changeSet *CharacterStateChangeSet) error {
		return u.syncQueueChanged(runCtx, changeSet.QueueChanged)
	})
	return err
}

func (u *CharacterActionQueueUsecase) Clear(ctx context.Context, characterID int64) error {
	command := &CharacterActionQueueClearCommand{CharacterID: characterID}
	_, err := u.stateUsecase.ExecuteCommand(ctx, command, func(runCtx context.Context, changeSet *CharacterStateChangeSet) error {
		return u.syncQueueChanged(runCtx, changeSet.QueueChanged)
	})
	return err
}

func (u *CharacterActionQueueUsecase) syncQueueChanged(ctx context.Context, event *CharacterActionQueueChangedStateEvent) error {
	if event == nil {
		return nil
	}
	if event.HeadChanged() && event.OldHeadID != "" {
		u.scheduler.timeWheel.Remove(event.OldHeadID)
		u.scheduler.unmarkCurrentTask(event.CharacterID, event.OldHeadID)
	}
	u.eventUsecase.PublishActionQueueUpdated(ctx, event.Queue)
	if event.HeadChanged() && event.NewHeadID != "" {
		return u.scheduler.startCurrent(ctx, event.CharacterID)
	}
	return nil
}
