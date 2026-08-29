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
	mutex                sync.Mutex
	logger               *slog.Logger
	characterRepo        repo.CharacterRepo
	actionRepo           repo.ActionRepo
	characterSessionRepo repo.CharacterSessionRepo
	actionQueueRepo      repo.ActionQueueRepo
	backpackRepo         repo.BackpackRepo
	abilityRepo          repo.CharacterAbilityRepo
	gameIdleEventRepo    repo.GameIdleEventRepo
	abilityUsecase       *CharacterAbilityUsecase
	timeWheel            *timewheel.TimeWheel
	actionTasks          map[enum.ActionKind]ActionTask
	pendingTasks         chan *PendingActionTask
	offlineTasks         chan *OfflineActionTask
	characterLocks       sync.Map
	stop                 context.CancelFunc
	running              bool
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
	characterSessionRepo repo.CharacterSessionRepo,
	actionQueueRepo repo.ActionQueueRepo,
	backpackRepo repo.BackpackRepo,
	abilityRepo repo.CharacterAbilityRepo,
	gameIdleEventRepo repo.GameIdleEventRepo,
	abilityUsecase *CharacterAbilityUsecase,
	timeWheel *timewheel.TimeWheel,
	actionTasks map[enum.ActionKind]ActionTask,
) *ActionQueueUsecase {
	return &ActionQueueUsecase{
		logger:               logger,
		characterRepo:        characterRepo,
		actionRepo:           actionRepo,
		characterSessionRepo: characterSessionRepo,
		actionQueueRepo:      actionQueueRepo,
		backpackRepo:         backpackRepo,
		abilityRepo:          abilityRepo,
		gameIdleEventRepo:    gameIdleEventRepo,
		abilityUsecase:       abilityUsecase,
		timeWheel:            timeWheel,
		actionTasks:          actionTasks,
		pendingTasks:         make(chan *PendingActionTask, 1024),
		offlineTasks:         make(chan *OfflineActionTask, 1024),
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
		if err = u.withCharacterLock(characterID, func() error {
			return u.startHead(runCtx, characterID)
		}); err != nil {
			u.logger.ErrorContext(runCtx, "game idle restore action queue failed", constant.LogFieldErr, err, "character_id", characterID)
		}
	}

	go func() {
		for {
			select {
			case <-runCtx.Done():
				return
			case task := <-u.pendingTasks:
				if err := u.handlePendingTask(runCtx, task); err != nil && !errors.Is(err, context.Canceled) {
					u.logger.ErrorContext(runCtx, "game idle process pending action failed", constant.LogFieldErr, err, "character_id", task.CharacterID)
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
				if err := u.handleOfflineTask(runCtx, task); err != nil && !errors.Is(err, context.Canceled) {
					u.logger.ErrorContext(runCtx, "game idle process offline action failed", constant.LogFieldErr, err, "character_id", task.CharacterID)
				}
			}
		}
	}()

	return nil
}

func (u *ActionQueueUsecase) Stop(ctx context.Context) error {
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
	characterIDs, err := u.actionQueueRepo.ListCharacterIDs(ctx)
	if err != nil {
		u.logger.ErrorContext(ctx, "game idle list action queue characters failed", constant.LogFieldErr, err)
		return nil
	}
	for _, characterID := range characterIDs {
		if err = u.withCharacterLock(characterID, func() error {
			if err = u.stopHead(ctx, characterID); err != nil {
				return err
			}
			u.flushAndClearState(ctx, characterID)
			return nil
		}); err != nil {
			u.logger.ErrorContext(ctx, "game idle stop current action failed", constant.LogFieldErr, err, "character_id", characterID)
			continue
		}
	}
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

func (u *ActionQueueUsecase) Resume(ctx context.Context, characterID int64) error {
	return u.withCharacterLock(characterID, func() error {
		u.stopOfflineCheck(characterID)
		return u.startHead(ctx, characterID)
	})
}

func (u *ActionQueueUsecase) Persist(ctx context.Context, characterID int64) error {
	return u.actionQueueRepo.Persist(ctx, characterID)
}

// ScheduleOfflineCheck 安排角色离线收益到期检查。
func (u *ActionQueueUsecase) ScheduleOfflineCheck(ctx context.Context, characterID int64) error {
	return u.withCharacterLock(characterID, func() error {
		_, err := u.scheduleOfflineTimeout(ctx, characterID)
		return err
	})
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

	return u.withCharacterLock(req.CharacterID, func() error {
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
	return u.withCharacterLock(req.CharacterID, func() error {
		return u.editQueue(ctx, req.CharacterID, func(queue *model.ActionQueue) error {
			currentPosition := int(req.CurrentPosition)
			targetPosition := int(req.TargetPosition)
			if currentPosition < 0 ||
				targetPosition < 0 ||
				currentPosition >= len(queue.Items) ||
				targetPosition >= len(queue.Items) {
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
	return u.withCharacterLock(req.CharacterID, func() error {
		return u.editQueue(ctx, req.CharacterID, func(queue *model.ActionQueue) error {
			if len(queue.Items) == 0 {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
			}
			if req.Position < 0 || int(req.Position) >= len(queue.Items) {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
			}
			position := int(req.Position)
			queue.Items = append(queue.Items[:position], queue.Items[position+1:]...)
			return nil
		})
	})
}

func (u *ActionQueueUsecase) Clear(ctx context.Context, characterID int64) error {
	return u.withCharacterLock(characterID, func() error {
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
		if err = u.stopHead(ctx, characterID); err != nil {
			return err
		}
	}
	if err = u.actionQueueRepo.Save(ctx, queue); err != nil {
		if headChanged && oldHeadID != "" {
			_ = u.startHead(ctx, characterID)
		}
		return err
	}
	u.publishActionQueueUpdated(ctx, queue, actionQueueUpdateReasonManualChanged)
	if headChanged && len(queue.Items) > 0 {
		if err = u.startHead(ctx, characterID); err != nil {
			return err
		}
	}
	return nil
}

func (u *ActionQueueUsecase) headTaskID(queue *model.ActionQueue) string {
	if len(queue.Items) == 0 {
		return ""
	}
	return queue.Items[0].ID
}

func (u *ActionQueueUsecase) handlePendingTask(ctx context.Context, task *PendingActionTask) error {
	return u.withCharacterLock(task.CharacterID, func() error {
		queue, err := u.actionQueueRepo.Load(ctx, task.CharacterID)
		if err != nil {
			return err
		}
		if len(queue.Items) == 0 || queue.Items[0].ID != task.TaskID || queue.Items[0].ActionID != task.ActionID {
			return nil
		}

		timesRemaining, queueChanged := u.applyPendingTask(queue, task)
		if err = u.actionQueueRepo.Save(ctx, queue); err != nil {
			return err
		}
		if task.StopReason == enum.ActionStopReasonNone {
			u.publishActionCompleted(ctx, task, timesRemaining)
		}
		if queueChanged {
			u.publishActionQueueUpdated(ctx, queue, u.queueUpdateReason(task))
		}
		if len(queue.Items) > 0 {
			return u.startHead(ctx, task.CharacterID)
		}
		return nil
	})
}

func (u *ActionQueueUsecase) applyPendingTask(queue *model.ActionQueue, task *PendingActionTask) (int64, bool) {
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
	return timesRemaining, queueChanged
}

func (u *ActionQueueUsecase) publishActionCompleted(ctx context.Context, task *PendingActionTask, timesRemaining int64) {
	err := u.gameIdleEventRepo.Publish(ctx, &model.GameIdleEvent{
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
		u.logger.ErrorContext(ctx, "game idle action completed event publish failed", constant.LogFieldErr, err, "character_id", task.CharacterID)
	}
	if task.AbilityLeveledUp != nil {
		err = u.gameIdleEventRepo.Publish(ctx, &model.GameIdleEvent{
			AbilityLeveledUp: task.AbilityLeveledUp,
		})
		if err != nil {
			u.logger.ErrorContext(ctx, "game idle ability leveled up event publish failed", constant.LogFieldErr, err, "character_id", task.CharacterID)
		}
	}
}

func (u *ActionQueueUsecase) queueUpdateReason(task *PendingActionTask) string {
	if task.StopReason == enum.ActionStopReasonInsufficientItems {
		return actionQueueUpdateReasonInsufficientItems
	}
	return actionQueueUpdateReasonActionCompleted
}

func (u *ActionQueueUsecase) handleOfflineTask(ctx context.Context, task *OfflineActionTask) error {
	return u.withCharacterLock(task.CharacterID, func() error {
		online, err := u.characterSessionRepo.IsOnline(ctx, task.CharacterID)
		if err != nil || online {
			return err
		}
		character, err := u.characterRepo.Get(ctx, task.CharacterID)
		if err != nil {
			return err
		}
		if time.Since(task.LastLogoutAt) <= character.MaxOfflineDuration {
			return nil
		}
		if err = u.stopHead(ctx, task.CharacterID); err != nil {
			return err
		}
		u.flushAndClearState(ctx, task.CharacterID)
		return nil
	})
}

// startHead 启动当前队首行动。
func (u *ActionQueueUsecase) startHead(ctx context.Context, characterID int64) error {
	for {
		expired, err := u.scheduleOfflineTimeout(ctx, characterID)
		if err != nil || expired {
			return err
		}
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

// stopHead 停止当前队首行动。
func (u *ActionQueueUsecase) stopHead(ctx context.Context, characterID int64) error {
	queue, err := u.actionQueueRepo.Load(ctx, characterID)
	if err != nil {
		return err
	}
	if len(queue.Items) > 0 {
		u.timeWheel.Remove(queue.Items[0].ID)
	}
	u.stopOfflineCheck(characterID)
	return nil
}

func (u *ActionQueueUsecase) scheduleOfflineTimeout(ctx context.Context, characterID int64) (bool, error) {
	online, err := u.characterSessionRepo.IsOnline(ctx, characterID)
	if err != nil {
		return false, err
	}
	if online {
		u.stopOfflineCheck(characterID)
		return false, nil
	}
	character, err := u.characterRepo.Get(ctx, characterID)
	if err != nil {
		return false, err
	}
	if character.LastOfflineAt == nil {
		return false, nil
	}
	dueAt := character.LastOfflineAt.Add(character.MaxOfflineDuration)
	if !time.Now().Before(dueAt) {
		u.stopOfflineCheck(characterID)
		u.flushAndClearState(ctx, characterID)
		return true, nil
	}
	offlineAt := *character.LastOfflineAt
	return false, u.timeWheel.Add(&timewheel.Task{
		ID:    u.offlineTaskID(characterID),
		DueAt: dueAt,
		Job: func(jobCtx context.Context, task *timewheel.Task) error {
			select {
			case <-jobCtx.Done():
				return jobCtx.Err()
			case u.offlineTasks <- &OfflineActionTask{
				CharacterID:  characterID,
				LastLogoutAt: offlineAt,
				Now:          time.Now(),
			}:
				return nil
			}
		},
	})
}

func (u *ActionQueueUsecase) stopOfflineCheck(characterID int64) {
	u.timeWheel.Remove(u.offlineTaskID(characterID))
}

func (u *ActionQueueUsecase) offlineTaskID(characterID int64) string {
	return fmt.Sprintf("character:%d:offline-timeout", characterID)
}

func (u *ActionQueueUsecase) characterLock(characterID int64) *sync.Mutex {
	lock, _ := u.characterLocks.LoadOrStore(characterID, &sync.Mutex{})
	return lock.(*sync.Mutex)
}

func (u *ActionQueueUsecase) withCharacterLock(characterID int64, fn func() error) error {
	lock := u.characterLock(characterID)
	lock.Lock()
	defer lock.Unlock()
	return fn()
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

func (u *ActionQueueUsecase) flushAndClearState(ctx context.Context, characterID int64) {
	if err := u.actionQueueRepo.Persist(ctx, characterID); err != nil {
		u.logger.ErrorContext(ctx, "game idle persist action queue failed", constant.LogFieldErr, err, "character_id", characterID)
		return
	}
	if err := u.backpackRepo.PersistItems(ctx, characterID); err != nil {
		u.logger.ErrorContext(ctx, "game idle persist backpack failed", constant.LogFieldErr, err, "character_id", characterID)
		return
	}
	if err := u.abilityRepo.Persist(ctx, characterID); err != nil {
		u.logger.ErrorContext(ctx, "game idle persist character ability failed", constant.LogFieldErr, err, "character_id", characterID)
		return
	}
	if err := u.actionQueueRepo.Clear(ctx, characterID); err != nil {
		u.logger.ErrorContext(ctx, "game idle clear action queue cache failed", constant.LogFieldErr, err, "character_id", characterID)
		return
	}
	if err := u.backpackRepo.Clear(ctx, characterID); err != nil {
		u.logger.ErrorContext(ctx, "game idle clear backpack cache failed", constant.LogFieldErr, err, "character_id", characterID)
		return
	}
	if err := u.abilityRepo.Clear(ctx, characterID); err != nil {
		u.logger.ErrorContext(ctx, "game idle clear character ability cache failed", constant.LogFieldErr, err, "character_id", characterID)
	}
}
