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

// ActionSchedulerUsecase 负责行动队列首位任务的调度、完成推进和离线到期处理。
type ActionSchedulerUsecase struct {
	mutex                sync.Mutex
	logger               *slog.Logger
	characterRepo        repo.CharacterRepo
	actionRepo           repo.ActionRepo
	characterSessionRepo repo.CharacterSessionRepo
	actionQueueRepo      repo.ActionQueueRepo
	abilityUsecase       *CharacterAbilityUsecase
	eventUsecase         *GameIdleEventUsecase
	stateUsecase         *CharacterStateUsecase
	timeWheel            *timewheel.TimeWheel
	actionTasks          map[enum.ActionKind]ActionTask
	locker               *ActionQueueLocker
	pendingTasks         chan *PendingActionTask
	offlineTasks         chan *OfflineActionTask
	stop                 context.CancelFunc
	running              bool
}

func NewActionSchedulerUsecase(
	logger *slog.Logger,
	characterRepo repo.CharacterRepo,
	actionRepo repo.ActionRepo,
	characterSessionRepo repo.CharacterSessionRepo,
	actionQueueRepo repo.ActionQueueRepo,
	abilityUsecase *CharacterAbilityUsecase,
	eventUsecase *GameIdleEventUsecase,
	stateUsecase *CharacterStateUsecase,
	timeWheel *timewheel.TimeWheel,
	actionTasks map[enum.ActionKind]ActionTask,
	locker *ActionQueueLocker,
) *ActionSchedulerUsecase {
	return &ActionSchedulerUsecase{
		logger:               logger,
		characterRepo:        characterRepo,
		actionRepo:           actionRepo,
		characterSessionRepo: characterSessionRepo,
		actionQueueRepo:      actionQueueRepo,
		abilityUsecase:       abilityUsecase,
		eventUsecase:         eventUsecase,
		stateUsecase:         stateUsecase,
		timeWheel:            timeWheel,
		actionTasks:          actionTasks,
		locker:               locker,
		pendingTasks:         make(chan *PendingActionTask, 1024),
		offlineTasks:         make(chan *OfflineActionTask, 1024),
	}
}

func (u *ActionSchedulerUsecase) Start(ctx context.Context) error {
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

	go u.consumePendingTasks(runCtx)
	go u.consumeOfflineTasks(runCtx)

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
		if err = u.locker.withCharacterLock(characterID, func() error {
			return u.startCurrent(runCtx, characterID)
		}); err != nil {
			u.logger.ErrorContext(runCtx, "game idle restore action queue failed", constant.LogFieldErr, err, "character_id", characterID)
		}
	}
	return nil
}

func (u *ActionSchedulerUsecase) Stop(ctx context.Context) error {
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
		if err = u.locker.withCharacterLock(characterID, func() error {
			u.stopCurrent(ctx, characterID)
			u.stateUsecase.FlushAndClear(ctx, characterID)
			return nil
		}); err != nil {
			u.logger.ErrorContext(ctx, "game idle stop current action failed", constant.LogFieldErr, err, "character_id", characterID)
		}
	}
	return nil
}

func (u *ActionSchedulerUsecase) Resume(ctx context.Context, characterID int64) error {
	return u.locker.withCharacterLock(characterID, func() error {
		u.stopOfflineCheck(characterID)
		return u.startCurrent(ctx, characterID)
	})
}

// ScheduleOfflineCheck 安排角色离线收益到期检查。
func (u *ActionSchedulerUsecase) ScheduleOfflineCheck(ctx context.Context, characterID int64) error {
	return u.locker.withCharacterLock(characterID, func() error {
		_, err := u.scheduleOfflineTimeout(ctx, characterID)
		return err
	})
}

func (u *ActionSchedulerUsecase) consumePendingTasks(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case task := <-u.pendingTasks:
			if err := u.handlePendingTask(ctx, task); err != nil && !errors.Is(err, context.Canceled) {
				u.logger.ErrorContext(ctx, "game idle process pending action failed", constant.LogFieldErr, err, "character_id", task.CharacterID)
			}
		}
	}
}

func (u *ActionSchedulerUsecase) consumeOfflineTasks(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case task := <-u.offlineTasks:
			if err := u.handleOfflineTask(ctx, task); err != nil && !errors.Is(err, context.Canceled) {
				u.logger.ErrorContext(ctx, "game idle process offline action failed", constant.LogFieldErr, err, "character_id", task.CharacterID)
			}
		}
	}
}

// handlePendingTask 在单角色锁内推进队首，保证扣次数、移除和重新调度是一组连续操作。
func (u *ActionSchedulerUsecase) handlePendingTask(ctx context.Context, task *PendingActionTask) error {
	return u.locker.withCharacterLock(task.CharacterID, func() error {
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
			u.eventUsecase.PublishActionQueueUpdated(ctx, queue, u.queueUpdateReason(task))
		}
		if len(queue.Items) > 0 {
			return u.startCurrent(ctx, task.CharacterID)
		}
		return nil
	})
}

func (u *ActionSchedulerUsecase) applyPendingTask(queue *model.ActionQueue, task *PendingActionTask) (int64, bool) {
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

func (u *ActionSchedulerUsecase) publishActionCompleted(ctx context.Context, task *PendingActionTask, timesRemaining int64) {
	u.eventUsecase.PublishActionCompleted(ctx, &model.ActionCompletedEvent{
		CharacterID:    task.CharacterID,
		ActionID:       task.ActionID,
		TimesFinished:  1,
		TimesRemaining: timesRemaining,
		StartedAt:      task.StartedAt,
		CompletedAt:    task.CompletedAt,
		ItemChanges:    task.ItemChanges,
		AbilityChanges: task.AbilityChanges,
	})
	u.eventUsecase.PublishAbilityLeveledUp(ctx, task.AbilityLeveledUp)
}

func (u *ActionSchedulerUsecase) queueUpdateReason(task *PendingActionTask) string {
	if task.StopReason == enum.ActionStopReasonInsufficientItems {
		return ActionQueueUpdateReasonInsufficientItems
	}
	return ActionQueueUpdateReasonActionCompleted
}

func (u *ActionSchedulerUsecase) handleOfflineTask(ctx context.Context, task *OfflineActionTask) error {
	return u.locker.withCharacterLock(task.CharacterID, func() error {
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
		u.stopCurrent(ctx, task.CharacterID)
		u.stateUsecase.FlushAndClear(ctx, task.CharacterID)
		return nil
	})
}

// startCurrent 启动当前队首行动；物品不足时移除队首并继续尝试下一项。
func (u *ActionSchedulerUsecase) startCurrent(ctx context.Context, characterID int64) error {
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
		})
		if code, ok := apperror.BusinessCode(err); ok && code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_BACKPACK_INSUFFICIENT {
			queue.Items = queue.Items[1:]
			if err = u.actionQueueRepo.Save(ctx, queue); err != nil {
				return err
			}
			u.eventUsecase.PublishActionQueueUpdated(ctx, queue, ActionQueueUpdateReasonInsufficientItems)
			continue
		}
		if err != nil {
			return err
		}
		u.wrapActionTask(task)
		return u.timeWheel.Add(task)
	}
}

// wrapActionTask 在结算前确认任务仍是队首，避免玩家调整队列后旧任务继续产出。
func (u *ActionSchedulerUsecase) wrapActionTask(task *timewheel.Task) {
	job := task.Job
	task.Job = func(jobCtx context.Context, item *timewheel.Task) error {
		payload := item.Payload.(*model.ActionTask)
		return u.locker.withCharacterLock(payload.CharacterID, func() error {
			queue, err := u.actionQueueRepo.Load(jobCtx, payload.CharacterID)
			if err != nil {
				return err
			}
			if len(queue.Items) == 0 || queue.Items[0].ID != payload.TaskID {
				return nil
			}
			return job(jobCtx, item)
		})
	}
}

// stopCurrent 停止当前队首行动和离线到期检查，不修改队列内容。
func (u *ActionSchedulerUsecase) stopCurrent(ctx context.Context, characterID int64) {
	queue, err := u.actionQueueRepo.Load(ctx, characterID)
	if err != nil {
		u.logger.ErrorContext(ctx, "game idle load action queue before stop failed", constant.LogFieldErr, err, "character_id", characterID)
		return
	}
	if len(queue.Items) > 0 {
		u.timeWheel.Remove(queue.Items[0].ID)
	}
	u.stopOfflineCheck(characterID)
}

// scheduleOfflineTimeout 为离线角色设置收益上限检查，到期后刷盘并释放 Redis 热状态。
func (u *ActionSchedulerUsecase) scheduleOfflineTimeout(ctx context.Context, characterID int64) (bool, error) {
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
		u.stateUsecase.FlushAndClear(ctx, characterID)
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

func (u *ActionSchedulerUsecase) stopOfflineCheck(characterID int64) {
	u.timeWheel.Remove(u.offlineTaskID(characterID))
}

func (u *ActionSchedulerUsecase) offlineTaskID(characterID int64) string {
	return fmt.Sprintf("character:%d:offline-timeout", characterID)
}
