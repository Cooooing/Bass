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
	stateEngine     *StateEngine
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
	stateEngine *StateEngine,
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
		stateEngine:     stateEngine,
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

	return u.applyQueueCommand(ctx, req.CharacterID, &ActionQueueAddCommand{
		CharacterID: req.CharacterID,
		ActionID:    req.ActionID,
		Times:       req.Times,
		Capacity:    character.ActionQueueCapacity,
		Position:    req.Position,
		Now:         time.Now(),
	})
}

type MoveActionReq struct {
	CharacterID     int64
	CurrentPosition int32
	TargetPosition  int32
}

func (u *ActionQueueUsecase) Move(ctx context.Context, req *MoveActionReq) error {
	return u.applyQueueCommand(ctx, req.CharacterID, &ActionQueueMoveCommand{
		CharacterID:     req.CharacterID,
		CurrentPosition: req.CurrentPosition,
		TargetPosition:  req.TargetPosition,
	})
}

type RemoveActionReq struct {
	CharacterID int64
	Position    int32
}

func (u *ActionQueueUsecase) Remove(ctx context.Context, req *RemoveActionReq) error {
	return u.applyQueueCommand(ctx, req.CharacterID, &ActionQueueRemoveCommand{
		CharacterID: req.CharacterID,
		Position:    req.Position,
	})
}

func (u *ActionQueueUsecase) Clear(ctx context.Context, characterID int64) error {
	return u.applyQueueCommand(ctx, characterID, &ActionQueueClearCommand{CharacterID: characterID})
}

func (u *ActionQueueUsecase) applyQueueCommand(ctx context.Context, characterID int64, command StateCommand) error {
	return u.locker.withCharacterLock(characterID, func() error {
		changeSet, err := u.stateEngine.Apply(ctx, command)
		if err != nil {
			return err
		}
		return u.syncQueueChanged(ctx, changeSet.QueueChanged)
	})
}

func (u *ActionQueueUsecase) syncQueueChanged(ctx context.Context, event *ActionQueueChangedStateEvent) error {
	if event == nil {
		return nil
	}
	if event.HeadChanged() && event.OldHeadID != "" {
		u.scheduler.stopTask(event.OldHeadID)
	}
	u.eventUsecase.PublishActionQueueUpdated(ctx, event.Queue)
	if event.HeadChanged() && event.NewHeadID != "" {
		return u.scheduler.startCurrent(ctx, event.CharacterID)
	}
	return nil
}
