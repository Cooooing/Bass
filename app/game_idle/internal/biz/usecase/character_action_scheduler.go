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

// CharacterActionSchedulerUsecase 负责行动队列首位任务的调度、完成推进和离线到期处理。
type CharacterActionSchedulerUsecase struct {
	mutex                sync.Mutex
	logger               *slog.Logger
	characterRepo        repo.CharacterRepo
	actionRepo           repo.MetaActionRepo
	characterSessionRepo repo.CharacterSessionRepo
	actionQueueRepo      repo.CharacterActionQueueRepo
	abilityUsecase       *CharacterAbilityUsecase
	eventUsecase         *GameIdleEventUsecase
	stateUsecase         *CharacterStateUsecase
	timeWheel            *timewheel.TimeWheel
	actionTasks          map[enum.ActionKind]CharacterActionTaskBuilder
	// pendingTasks 接收时间轮完成的行动任务，统一回到角色状态通道里结算。
	pendingTasks chan *PendingCharacterActionTask
	// offlineTasks 接收离线收益到期任务，用于停止行动、刷库并释放热状态。
	offlineTasks chan *OfflineCharacterActionTask
	// scheduledTasks 记录当前已放入时间轮的队首任务，避免重复调度同一个队列项。
	scheduledTasks map[int64]string
	stop           context.CancelFunc
	running        bool
}

func NewCharacterActionSchedulerUsecase(
	logger *slog.Logger,
	characterRepo repo.CharacterRepo,
	actionRepo repo.MetaActionRepo,
	characterSessionRepo repo.CharacterSessionRepo,
	actionQueueRepo repo.CharacterActionQueueRepo,
	abilityUsecase *CharacterAbilityUsecase,
	eventUsecase *GameIdleEventUsecase,
	stateUsecase *CharacterStateUsecase,
	timeWheel *timewheel.TimeWheel,
	actionTasks map[enum.ActionKind]CharacterActionTaskBuilder,
) *CharacterActionSchedulerUsecase {
	return &CharacterActionSchedulerUsecase{
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
		pendingTasks:         make(chan *PendingCharacterActionTask, 1024),
		offlineTasks:         make(chan *OfflineCharacterActionTask, 1024),
		scheduledTasks:       make(map[int64]string),
	}
}

func (u *CharacterActionSchedulerUsecase) Start(ctx context.Context) error {
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

	// 调度器只消费通道，具体状态变更仍进入 CharacterStateUsecase 串行执行。
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
		err = u.stateUsecase.Execute(runCtx, characterID, func(jobCtx context.Context) error {
			return u.startCurrent(jobCtx, characterID)
		})
		if err != nil {
			u.logger.ErrorContext(runCtx, "game idle restore action queue failed", constant.LogFieldErr, err, "character_id", characterID)
		}
	}
	return nil
}

func (u *CharacterActionSchedulerUsecase) Stop(ctx context.Context) error {
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
		if err = u.stateUsecase.Execute(ctx, characterID, func(jobCtx context.Context) error {
			u.stopCurrent(jobCtx, characterID)
			u.stateUsecase.flushAndClear(jobCtx, characterID)
			return nil
		}); err != nil {
			u.logger.ErrorContext(ctx, "game idle stop current action failed", constant.LogFieldErr, err, "character_id", characterID)
		}
	}
	return nil
}

func (u *CharacterActionSchedulerUsecase) Resume(ctx context.Context, characterID int64) error {
	return u.stateUsecase.Execute(ctx, characterID, func(jobCtx context.Context) error {
		u.timeWheel.Remove(u.offlineTaskID(characterID))
		return u.startCurrent(jobCtx, characterID)
	})
}

// ScheduleOfflineCheck 安排角色离线收益到期检查。
func (u *CharacterActionSchedulerUsecase) ScheduleOfflineCheck(ctx context.Context, characterID int64) error {
	return u.stateUsecase.Execute(ctx, characterID, func(jobCtx context.Context) error {
		_, err := u.scheduleOfflineTimeout(jobCtx, characterID)
		return err
	})
}

func (u *CharacterActionSchedulerUsecase) consumePendingTasks(ctx context.Context) {
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

func (u *CharacterActionSchedulerUsecase) consumeOfflineTasks(ctx context.Context) {
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
func (u *CharacterActionSchedulerUsecase) handlePendingTask(ctx context.Context, task *PendingCharacterActionTask) error {
	return u.stateUsecase.Execute(ctx, task.CharacterID, func(jobCtx context.Context) error {
		syncQueue := func(changeSet *CharacterStateChangeSet) error {
			if changeSet.QueueChanged != nil {
				u.eventUsecase.PublishActionQueueUpdated(jobCtx, changeSet.QueueChanged.Queue)
				if changeSet.QueueChanged.NewHeadID != "" {
					return u.startCurrent(jobCtx, task.CharacterID)
				}
				return nil
			}
			// 无限行动结算不会改变队列，但仍需要重新调度下一轮。
			return u.startCurrent(jobCtx, task.CharacterID)
		}
		if task.StopReason != enum.ActionStopReasonNone {
			changeSet, err := u.stateUsecase.apply(jobCtx, &CharacterActionQueueCompleteHeadCommand{
				CharacterID: task.CharacterID,
				TaskID:      task.TaskID,
				ActionID:    task.ActionID,
				RemoveHead:  true,
			})
			if err != nil || !changeSet.QueueCommandApplied {
				return err
			}
			return syncQueue(changeSet)
		}
		changeSet, err := u.stateUsecase.apply(jobCtx, &CharacterSettlementCommand{
			CharacterID: task.CharacterID,
			TaskID:      task.TaskID,
			ActionID:    task.ActionID,
			Items:       task.Items,
			AbilityID:   task.AbilityID,
			ExpReward:   task.ExpReward,
		})
		if code, ok := apperror.BusinessCode(err); ok && code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_BACKPACK_INSUFFICIENT {
			task.StopReason = enum.ActionStopReasonInsufficientItems
			changeSet, err = u.stateUsecase.apply(jobCtx, &CharacterActionQueueCompleteHeadCommand{
				CharacterID: task.CharacterID,
				TaskID:      task.TaskID,
				ActionID:    task.ActionID,
				RemoveHead:  true,
			})
			if err != nil || !changeSet.QueueCommandApplied {
				return err
			}
			return syncQueue(changeSet)
		}
		if err != nil || !changeSet.QueueCommandApplied {
			return err
		}
		u.eventUsecase.PublishActionCompleted(jobCtx, &model.ActionCompletedEvent{
			CharacterID:    task.CharacterID,
			ActionID:       task.ActionID,
			TimesFinished:  1,
			TimesRemaining: changeSet.ActionTimesRemaining,
			StartedAt:      task.StartedAt,
			CompletedAt:    task.CompletedAt,
			ItemChanges:    changeSet.ItemChanges,
			AbilityChanges: changeSet.AbilityChanges,
		})
		for _, event := range changeSet.AbilityLeveledUp {
			u.eventUsecase.PublishAbilityLeveledUp(jobCtx, event)
		}
		return syncQueue(changeSet)
	})
}

func (u *CharacterActionSchedulerUsecase) handleOfflineTask(ctx context.Context, task *OfflineCharacterActionTask) error {
	return u.stateUsecase.Execute(ctx, task.CharacterID, func(jobCtx context.Context) error {
		online, err := u.characterSessionRepo.IsOnline(jobCtx, task.CharacterID)
		if err != nil || online {
			return err
		}
		character, err := u.characterRepo.Get(jobCtx, task.CharacterID)
		if err != nil {
			return err
		}
		if character == nil {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_NOT_FOUND)
		}
		if time.Since(task.LastLogoutAt) <= character.MaxOfflineDuration {
			return nil
		}
		u.stopCurrent(jobCtx, task.CharacterID)
		u.stateUsecase.flushAndClear(jobCtx, task.CharacterID)
		return nil
	})
}

// startCurrent 启动当前队首行动；物品不足时移除队首并继续尝试下一项。
func (u *CharacterActionSchedulerUsecase) startCurrent(ctx context.Context, characterID int64) error {
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
		u.mutex.Lock()
		taskScheduled := u.scheduledTasks[characterID] == current.ID
		u.mutex.Unlock()
		if taskScheduled {
			return nil
		}
		actionConfig, err := u.actionRepo.Get(ctx, current.ActionID)
		if err != nil {
			return err
		}
		if actionConfig == nil {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
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
		task, err := actionTask.BuildTask(ctx, &BuildCharacterActionTaskReq{
			CharacterID:  queue.CharacterID,
			QueueItem:    current,
			Action:       actionConfig,
			Now:          time.Now(),
			PendingTasks: u.pendingTasks,
		})
		if code, ok := apperror.BusinessCode(err); ok && code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_BACKPACK_INSUFFICIENT {
			// 当前行动已无法启动时直接移除队首，并尝试启动后续行动。
			changeSet, err := u.stateUsecase.apply(ctx, &CharacterActionQueueCompleteHeadCommand{
				CharacterID: characterID,
				TaskID:      current.ID,
				ActionID:    current.ActionID,
				RemoveHead:  true,
			})
			if err != nil {
				return err
			}
			if changeSet.QueueChanged != nil {
				u.eventUsecase.PublishActionQueueUpdated(ctx, changeSet.QueueChanged.Queue)
			}
			continue
		}
		if err != nil {
			return err
		}
		u.wrapActionTask(task)
		if err = u.timeWheel.Add(task); err != nil {
			return err
		}
		u.mutex.Lock()
		u.scheduledTasks[characterID] = task.ID
		u.mutex.Unlock()
		return nil
	}
}

// wrapActionTask 在结算前确认任务仍是队首，避免玩家调整队列后旧任务继续产出。
func (u *CharacterActionSchedulerUsecase) wrapActionTask(task *timewheel.Task) {
	job := task.Job
	task.Job = func(jobCtx context.Context, item *timewheel.Task) error {
		payload := item.Payload.(*model.CharacterActionTask)
		u.unmarkCurrentTask(payload.CharacterID, payload.TaskID)
		return u.stateUsecase.Execute(jobCtx, payload.CharacterID, func(runCtx context.Context) error {
			queue, err := u.actionQueueRepo.Load(runCtx, payload.CharacterID)
			if err != nil {
				return err
			}
			if len(queue.Items) == 0 || queue.Items[0].ID != payload.TaskID {
				return nil
			}
			return job(runCtx, item)
		})
	}
}

// stopCurrent 停止当前队首行动和离线到期检查，不修改队列内容。
func (u *CharacterActionSchedulerUsecase) stopCurrent(ctx context.Context, characterID int64) {
	queue, err := u.actionQueueRepo.Load(ctx, characterID)
	if err != nil {
		u.logger.ErrorContext(ctx, "game idle load action queue before stop failed", constant.LogFieldErr, err, "character_id", characterID)
		return
	}
	if len(queue.Items) > 0 {
		u.timeWheel.Remove(queue.Items[0].ID)
		u.unmarkCurrentTask(characterID, queue.Items[0].ID)
	}
	u.timeWheel.Remove(u.offlineTaskID(characterID))
}

// scheduleOfflineTimeout 为离线角色设置收益上限检查，到期后刷盘并释放本地热状态。
func (u *CharacterActionSchedulerUsecase) scheduleOfflineTimeout(ctx context.Context, characterID int64) (bool, error) {
	online, err := u.characterSessionRepo.IsOnline(ctx, characterID)
	if err != nil {
		return false, err
	}
	if online {
		u.timeWheel.Remove(u.offlineTaskID(characterID))
		return false, nil
	}
	character, err := u.characterRepo.Get(ctx, characterID)
	if err != nil {
		return false, err
	}
	if character == nil {
		return false, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_CHARACTER_NOT_FOUND)
	}
	if character.LastOfflineAt == nil {
		return false, nil
	}
	dueAt := character.LastOfflineAt.Add(character.MaxOfflineDuration)
	if !time.Now().Before(dueAt) {
		u.timeWheel.Remove(u.offlineTaskID(characterID))
		u.stateUsecase.flushAndClear(ctx, characterID)
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
			case u.offlineTasks <- &OfflineCharacterActionTask{
				CharacterID:  characterID,
				LastLogoutAt: offlineAt,
				Now:          time.Now(),
			}:
				return nil
			}
		},
	})
}

func (u *CharacterActionSchedulerUsecase) offlineTaskID(characterID int64) string {
	return fmt.Sprintf("character:%d:offline-timeout", characterID)
}

func (u *CharacterActionSchedulerUsecase) unmarkCurrentTask(characterID int64, taskID string) {
	u.mutex.Lock()
	if u.scheduledTasks[characterID] == taskID {
		delete(u.scheduledTasks, characterID)
	}
	u.mutex.Unlock()
}
