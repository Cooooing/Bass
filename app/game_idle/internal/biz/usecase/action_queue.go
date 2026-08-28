package usecase

import (
	"common/pkg/apperror"
	"common/pkg/client/timewheel"
	"common/pkg/constant"
	cerrors "common/proto/gen/common/errors"
	"context"
	"errors"
	"fmt"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/enum"
	"log/slog"
	"sync"
	"time"
)

type ActionQueueUsecase struct {
	mutex             sync.Mutex
	logger            *slog.Logger
	characterRepo     repo.CharacterRepo
	actionRepo        repo.ActionRepo
	actionQueueRepo   repo.ActionQueueRepo
	gameIdleEventRepo repo.GameIdleEventRepo
	abilityUsecase    *CharacterAbilityUsecase
	timeWheel         *timewheel.TimeWheel
	actionTasks       map[enum.ActionKind]ActionTask
	pendingTasks      chan *PendingActionTask
	offlineTasks      chan *OfflineActionTask
	stop              context.CancelFunc
	running           bool
}

const (
	actionQueueUpdateReasonManualChanged     = "manual_changed"
	actionQueueUpdateReasonActionCompleted   = "action_completed"
	actionQueueUpdateReasonInsufficientItems = "insufficient_items"
)

func NewActionQueueUsecase(
	logger *slog.Logger,
	characterRepo repo.CharacterRepo,
	actionRepo repo.ActionRepo,
	actionQueueRepo repo.ActionQueueRepo,
	gameIdleEventRepo repo.GameIdleEventRepo,
	abilityUsecase *CharacterAbilityUsecase,
	timeWheel *timewheel.TimeWheel,
	actionTasks map[enum.ActionKind]ActionTask,
) *ActionQueueUsecase {
	return &ActionQueueUsecase{
		logger:            logger,
		characterRepo:     characterRepo,
		actionRepo:        actionRepo,
		actionQueueRepo:   actionQueueRepo,
		gameIdleEventRepo: gameIdleEventRepo,
		abilityUsecase:    abilityUsecase,
		timeWheel:         timeWheel,
		actionTasks:       actionTasks,
		pendingTasks:      make(chan *PendingActionTask, 1024),
		offlineTasks:      make(chan *OfflineActionTask, 1024),
	}
}

func (u *ActionQueueUsecase) Start(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	u.mutex.Lock()
	if u.running {
		u.mutex.Unlock()
		return nil
	}
	runCtx, cancel := context.WithCancel(ctx)
	u.stop = cancel
	u.running = true
	u.mutex.Unlock()

	characterIDs, err := u.actionQueueRepo.ListCharacterIDs(runCtx)
	if err != nil {
		cancel()
		u.mutex.Lock()
		u.stop = nil
		u.running = false
		u.mutex.Unlock()
		return err
	}
	for _, characterID := range characterIDs {
		if err = u.startCurrent(runCtx, characterID); err != nil {
			u.logger.ErrorContext(runCtx, "game idle restore action queue failed", constant.LogFieldErr, err, "character_id", characterID)
		}
	}

	go func() {
		for {
			select {
			case <-runCtx.Done():
				return
			case task := <-u.pendingTasks:
				queue, err := u.actionQueueRepo.Load(runCtx, task.CharacterID)
				if err != nil {
					u.logger.ErrorContext(runCtx, "game idle load action queue failed", constant.LogFieldErr, err, "character_id", task.CharacterID)
					continue
				}
				if len(queue.Items) == 0 || queue.Items[0].ID != task.TaskID || queue.Items[0].ActionID != task.ActionID {
					continue
				}

				current := queue.Items[0]
				timesRemaining := current.Times
				finishCurrent := task.StopReason != enum.ActionStopReasonNone
				queueChanged := finishCurrent || current.Times != -1
				if !finishCurrent && current.Times != -1 {
					current.Times--
					timesRemaining = current.Times
					finishCurrent = current.Times <= 0
				}
				if finishCurrent {
					queue.Items = queue.Items[1:]
					timesRemaining = 0
				}
				if err = u.actionQueueRepo.Save(runCtx, queue); err != nil {
					u.logger.ErrorContext(runCtx, "game idle save action queue failed", constant.LogFieldErr, err, "character_id", task.CharacterID)
					continue
				}
				if task.StopReason == enum.ActionStopReasonNone {
					err = u.gameIdleEventRepo.Publish(runCtx, &model.GameIdleEvent{
						ActionCompleted: &model.ActionCompletedEvent{
							CharacterID:    task.CharacterID,
							ActionID:       task.ActionID,
							TimesFinished:  1,
							TimesRemaining: timesRemaining,
							StartedAt:      task.StartedAt,
							CompletedAt:    task.CompletedAt,
							ItemChanges:    task.ItemChanges,
							AbilityChanges: task.AbilityChanges,
						},
					})
					if err != nil {
						u.logger.ErrorContext(runCtx, "game idle action completed event publish failed", constant.LogFieldErr, err, "character_id", task.CharacterID)
					}
					if task.AbilityLeveledUp != nil {
						err = u.gameIdleEventRepo.Publish(runCtx, &model.GameIdleEvent{
							AbilityLeveledUp: task.AbilityLeveledUp,
						})
						if err != nil {
							u.logger.ErrorContext(runCtx, "game idle ability leveled up event publish failed", constant.LogFieldErr, err, "character_id", task.CharacterID)
						}
					}
				}
				if queueChanged {
					reason := actionQueueUpdateReasonActionCompleted
					if task.StopReason == enum.ActionStopReasonInsufficientItems {
						reason = actionQueueUpdateReasonInsufficientItems
					}
					u.publishActionQueueUpdated(runCtx, queue, reason)
				}
				if len(queue.Items) > 0 {
					if err = u.startCurrent(runCtx, task.CharacterID); err != nil && !errors.Is(err, context.Canceled) {
						u.logger.ErrorContext(runCtx, "game idle start current action failed", constant.LogFieldErr, err, "character_id", task.CharacterID)
					}
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-runCtx.Done():
				return
			case task := <-u.offlineTasks:
				character, err := u.characterRepo.Get(runCtx, task.CharacterID)
				if err != nil {
					u.logger.ErrorContext(runCtx, "game idle load character failed", constant.LogFieldErr, err, "character_id", task.CharacterID)
					continue
				}
				if time.Since(task.LastLogoutAt) <= character.MaxOfflineDuration {
					continue
				}
				if err = u.stopCurrent(runCtx, task.CharacterID); err != nil {
					u.logger.ErrorContext(runCtx, "game idle stop current action failed", constant.LogFieldErr, err, "character_id", task.CharacterID)
				}
			}
		}
	}()

	return nil
}

func (u *ActionQueueUsecase) Stop(ctx context.Context) error {
	_ = ctx
	u.mutex.Lock()
	if !u.running {
		u.mutex.Unlock()
		return nil
	}
	stop := u.stop
	u.stop = nil
	u.running = false
	u.mutex.Unlock()
	stop()
	return nil
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

	queue, err := u.actionQueueRepo.Load(ctx, req.CharacterID)
	if err != nil {
		return err
	}
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
	oldTaskID := ""
	if len(queue.Items) > 0 {
		oldTaskID = queue.Items[0].ID
	}
	queueItem := &model.ActionQueueItem{
		ID: fmt.Sprintf(
			"character:%d:action:%s:at:%d",
			req.CharacterID,
			req.ActionID,
			now.UnixNano(),
		),
		ActionID:  req.ActionID,
		Times:     req.Times,
		CreatedAt: now,
	}
	queue.Items = append(queue.Items, nil)
	copy(queue.Items[position+1:], queue.Items[position:])
	queue.Items[position] = queueItem
	headChanged := queue.Items[0].ID != oldTaskID
	if headChanged && oldTaskID != "" {
		if err = u.stopCurrent(ctx, req.CharacterID); err != nil {
			return err
		}
	}
	if err = u.actionQueueRepo.Save(ctx, queue); err != nil {
		if headChanged && oldTaskID != "" {
			_ = u.startCurrent(ctx, req.CharacterID)
		}
		return err
	}
	u.publishActionQueueUpdated(ctx, queue, actionQueueUpdateReasonManualChanged)

	if headChanged {
		if err = u.startCurrent(ctx, req.CharacterID); err != nil {
			return err
		}
	}
	return nil
}

type MoveActionReq struct {
	CharacterID     int64
	CurrentPosition int32
	TargetPosition  int32
}

func (u *ActionQueueUsecase) Move(ctx context.Context, req *MoveActionReq) error {
	queue, err := u.actionQueueRepo.Load(ctx, req.CharacterID)
	if err != nil {
		return err
	}
	currentPosition := int(req.CurrentPosition)
	targetPosition := int(req.TargetPosition)
	if currentPosition < 0 ||
		targetPosition < 0 ||
		currentPosition >= len(queue.Items) ||
		targetPosition >= len(queue.Items) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
	}

	oldTaskID := queue.Items[0].ID
	queueItem := queue.Items[currentPosition]
	queue.Items = append(queue.Items[:currentPosition], queue.Items[currentPosition+1:]...)
	if targetPosition >= len(queue.Items) {
		queue.Items = append(queue.Items, queueItem)
	} else {
		queue.Items = append(queue.Items, nil)
		copy(queue.Items[targetPosition+1:], queue.Items[targetPosition:])
		queue.Items[targetPosition] = queueItem
	}
	headChanged := queue.Items[0].ID != oldTaskID
	if headChanged {
		if err = u.stopCurrent(ctx, req.CharacterID); err != nil {
			return err
		}
	}
	if err = u.actionQueueRepo.Save(ctx, queue); err != nil {
		if headChanged {
			_ = u.startCurrent(ctx, req.CharacterID)
		}
		return err
	}
	u.publishActionQueueUpdated(ctx, queue, actionQueueUpdateReasonManualChanged)

	if headChanged {
		if err = u.startCurrent(ctx, req.CharacterID); err != nil {
			return err
		}
	}
	return nil
}

type RemoveActionReq struct {
	CharacterID int64
	Position    int32
}

func (u *ActionQueueUsecase) Remove(ctx context.Context, req *RemoveActionReq) error {
	queue, err := u.actionQueueRepo.Load(ctx, req.CharacterID)
	if err != nil {
		return err
	}
	if len(queue.Items) == 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
	}

	oldTaskID := queue.Items[0].ID
	if req.Position < 0 || int(req.Position) >= len(queue.Items) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
	}

	position := int(req.Position)
	queue.Items = append(queue.Items[:position], queue.Items[position+1:]...)
	headChanged := len(queue.Items) == 0 || queue.Items[0].ID != oldTaskID
	if headChanged {
		if err = u.stopCurrent(ctx, req.CharacterID); err != nil {
			return err
		}
	}
	if err = u.actionQueueRepo.Save(ctx, queue); err != nil {
		if headChanged {
			_ = u.startCurrent(ctx, req.CharacterID)
		}
		return err
	}
	u.publishActionQueueUpdated(ctx, queue, actionQueueUpdateReasonManualChanged)

	if headChanged && len(queue.Items) > 0 {
		if err = u.startCurrent(ctx, req.CharacterID); err != nil {
			return err
		}
	}
	return nil
}

func (u *ActionQueueUsecase) Clear(ctx context.Context, characterID int64) error {
	queue, err := u.actionQueueRepo.Load(ctx, characterID)
	if err != nil {
		return err
	}
	if len(queue.Items) > 0 {
		if err = u.stopCurrent(ctx, characterID); err != nil {
			return err
		}
	}
	queue.Items = make([]*model.ActionQueueItem, 0)
	if err = u.actionQueueRepo.Save(ctx, queue); err != nil {
		_ = u.startCurrent(ctx, characterID)
		return err
	}
	u.publishActionQueueUpdated(ctx, queue, actionQueueUpdateReasonManualChanged)
	return nil
}

// startCurrent 启动当前队首行动。
func (u *ActionQueueUsecase) startCurrent(ctx context.Context, characterID int64) error {
	for {
		queue, err := u.actionQueueRepo.Load(ctx, characterID)
		if err != nil {
			return err
		}
		if len(queue.Items) == 0 {
			return nil
		}
		current := queue.Items[0]
		actionConfig, err := u.actionRepo.Get(ctx, current.ActionID)
		if err != nil {
			return err
		}
		if !actionConfig.Enabled {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
		}
		if err = u.abilityUsecase.CheckLevel(
			ctx,
			characterID,
			enum.Ability(actionConfig.AbilityID),
			actionConfig.RequiredAbilityLevel,
		); err != nil {
			return err
		}
		actionTask := u.actionTasks[actionConfig.ActionKind]
		if actionTask == nil {
			return fmt.Errorf("game idle action task is required: %s", actionConfig.ActionKind)
		}
		task, err := actionTask.BuildTask(ctx, &BuildActionTaskReq{
			CharacterID:  queue.CharacterID,
			QueueItem:    current,
			Action:       actionConfig,
			Now:          time.Now(),
			PendingTasks: u.pendingTasks,
			OfflineTasks: u.offlineTasks,
		})
		if code, ok := apperror.BusinessCode(err); ok && code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_BACKPACK_INSUFFICIENT {
			queue.Items = queue.Items[1:]
			if err = u.actionQueueRepo.Save(ctx, queue); err != nil {
				return err
			}
			u.publishActionQueueUpdated(ctx, queue, actionQueueUpdateReasonInsufficientItems)
			continue
		}
		if err != nil {
			return err
		}
		return u.timeWheel.Add(task)
	}
}

// stopCurrent 停止当前队首行动。
func (u *ActionQueueUsecase) stopCurrent(ctx context.Context, characterID int64) error {
	queue, err := u.actionQueueRepo.Load(ctx, characterID)
	if err != nil {
		return err
	}
	if len(queue.Items) > 0 {
		u.timeWheel.Remove(queue.Items[0].ID)
	}
	return nil
}

func (u *ActionQueueUsecase) publishActionQueueUpdated(ctx context.Context, queue *model.ActionQueue, reason string) {
	err := u.gameIdleEventRepo.Publish(ctx, &model.GameIdleEvent{
		ActionQueueUpdated: &model.ActionQueueUpdatedEvent{
			CharacterID: queue.CharacterID,
			Items:       queue.Items,
			Reason:      reason,
			UpdatedAt:   time.Now(),
		},
	})
	if err != nil {
		u.logger.ErrorContext(ctx, "game idle action queue updated event publish failed", constant.LogFieldErr, err, "character_id", queue.CharacterID)
	}
}
