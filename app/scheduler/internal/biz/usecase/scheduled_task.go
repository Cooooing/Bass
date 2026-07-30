package usecase

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"math/rand/v2"
	"os"
	"scheduler/internal/biz/base"
	"scheduler/internal/biz/model"
	"scheduler/internal/biz/repo"
	taskimpl "scheduler/internal/biz/usecase/task"
	"scheduler/internal/config"
	schedulerenum "scheduler/internal/enum"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
)

type ScheduledTaskUsecase struct {
	logger                           *slog.Logger
	conf                             *config.Bootstrap
	cronParser                       cron.Parser
	tx                               base.Tx
	scheduledTaskRepo                repo.ScheduledTaskRepo
	scheduledTaskCacheRepo           repo.ScheduledTaskCacheRepo
	scheduledTaskScheduleRepo        repo.ScheduledTaskScheduleRepo
	scheduledTaskVersionRepo         repo.ScheduledTaskVersionRepo
	scheduledTaskExecutionRecordRepo repo.ScheduledTaskExecutionRecordRepo
	tasks                            map[string]taskimpl.Task
	runningCancelMu                  sync.Mutex
	runningCancels                   map[int64]context.CancelFunc
	runningWG                        sync.WaitGroup
	workerID                         string
	clockSkewGrace                   time.Duration
}

func NewScheduledTaskUsecase(
	logger *slog.Logger,
	conf *config.Bootstrap,
	tx base.Tx,
	scheduledTaskRepo repo.ScheduledTaskRepo,
	scheduledTaskCacheRepo repo.ScheduledTaskCacheRepo,
	scheduledTaskScheduleRepo repo.ScheduledTaskScheduleRepo,
	scheduledTaskVersionRepo repo.ScheduledTaskVersionRepo,
	executionRepo repo.ScheduledTaskExecutionRecordRepo,
	tasks map[string]taskimpl.Task,
) *ScheduledTaskUsecase {
	workerID := "scheduler"
	if host, err := os.Hostname(); err == nil && host != "" {
		workerID = host
	}
	return &ScheduledTaskUsecase{
		logger:                           logger,
		conf:                             conf,
		cronParser:                       cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow),
		tx:                               tx,
		scheduledTaskRepo:                scheduledTaskRepo,
		scheduledTaskCacheRepo:           scheduledTaskCacheRepo,
		scheduledTaskScheduleRepo:        scheduledTaskScheduleRepo,
		scheduledTaskVersionRepo:         scheduledTaskVersionRepo,
		scheduledTaskExecutionRecordRepo: executionRepo,
		tasks:                            tasks,
		runningCancels:                   make(map[int64]context.CancelFunc),
		workerID:                         workerID,
		clockSkewGrace:                   10 * time.Second,
	}
}

func (u *ScheduledTaskUsecase) Upsert(ctx context.Context, row *model.ScheduledTask) (*model.ScheduledTask, error) {
	if row == nil ||
		row.HandlerName == "" ||
		strings.TrimSpace(row.Title) == "" ||
		strings.TrimSpace(row.CronSpec) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row.TaskKey = strings.TrimSpace(row.TaskKey)
	if row.TaskKey == "" {
		row.TaskKey = uuid.NewString()
	}
	if _, err := u.cronParser.Parse(row.CronSpec); err != nil {
		return nil, err
	}
	if strings.TrimSpace(row.Payload) == "" {
		row.Payload = "{}"
	}
	if !json.Valid([]byte(row.Payload)) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if row.Timeout <= 0 {
		row.Timeout = 300 * time.Second
		if u.conf.GetScheduler() != nil && u.conf.GetScheduler().GetDefaultExecutionTimeout() != nil && u.conf.GetScheduler().GetDefaultExecutionTimeout().AsDuration() > 0 {
			row.Timeout = u.conf.GetScheduler().GetDefaultExecutionTimeout().AsDuration()
		}
	}
	if row.MaxAttempts <= 0 {
		row.MaxAttempts = 1
	}
	if row.MisfirePolicy == "" {
		row.MisfirePolicy = schedulerenum.TaskMisfirePolicyExecuteLatest
	}
	// 定时任务默认只补最新一次，过期窗口用于避免服务恢复后回放大量失效消息。
	if row.StaleAfter == nil {
		row.StaleAfter = new(time.Duration)
		*row.StaleAfter = 5 * time.Minute
	}
	var saved *model.ScheduledTask
	err := u.tx(ctx, func(txCtx context.Context) error {
		var err error
		saved, err = u.scheduledTaskRepo.Upsert(txCtx, row)
		if err != nil {
			return err
		}
		_, err = u.scheduledTaskVersionRepo.Create(txCtx, saved)
		return err
	})
	if err != nil {
		return nil, err
	}
	_ = u.scheduledTaskCacheRepo.DeleteScheduledTask(ctx, saved.TaskKey)
	schedulePrefix := "scheduler.schedule.scheduled_task"
	if u.conf.GetScheduler() != nil && u.conf.GetScheduler().GetScheduledTaskScheduleSubjectPrefix() != "" {
		schedulePrefix = u.conf.GetScheduler().GetScheduledTaskScheduleSubjectPrefix()
	}
	executeSubject := "scheduler.execute.scheduled_task"
	if u.conf.GetScheduler() != nil && u.conf.GetScheduler().GetScheduledTaskExecuteSubject() != "" {
		executeSubject = u.conf.GetScheduler().GetScheduledTaskExecuteSubject()
	}
	subject := fmt.Sprintf("%s.%d", schedulePrefix, saved.ID)
	if saved.Enabled {
		err = u.scheduledTaskScheduleRepo.Schedule(
			context.WithoutCancel(ctx),
			&repo.ScheduledTaskScheduleReq{
				ScheduledTask: saved,
				Subject:       subject,
				Target:        executeSubject,
			},
		)
	} else {
		err = u.scheduledTaskScheduleRepo.Cancel(context.WithoutCancel(ctx), subject)
	}
	if err != nil {
		u.logger.WarnContext(
			ctx,
			"同步定时任务 NATS 调度失败，后续由启动补偿修复",
			"error",
			err,
			"task_id",
			saved.ID,
		)
	}
	return saved, nil
}
func (u *ScheduledTaskUsecase) SeedDefaultTasks(ctx context.Context) error {
	handlerNames := make([]string, 0, len(u.tasks))
	for handlerName := range u.tasks {
		handlerNames = append(handlerNames, handlerName)
	}
	sort.Strings(handlerNames)
	defaultTasks := make([]*model.ScheduledTask, 0)
	for _, handlerName := range handlerNames {
		item := u.tasks[handlerName]
		for _, schedule := range item.DefaultScheduledTasks() {
			if schedule == nil {
				continue
			}
			title := strings.TrimSpace(schedule.Title)
			if title == "" {
				title = item.Title()
			}
			description := strings.TrimSpace(schedule.Description)
			if description == "" {
				description = item.Description()
			}
			defaultTasks = append(defaultTasks, &model.ScheduledTask{
				TaskKey:       schedule.TaskKey.String(),
				HandlerName:   item.HandlerName(),
				Title:         title,
				Description:   description,
				Enabled:       schedule.Enabled,
				CronSpec:      schedule.CronSpec,
				Payload:       schedule.Payload,
				Timeout:       schedule.Timeout,
				StaleAfter:    schedule.StaleAfter,
				MaxAttempts:   schedule.MaxAttempts,
				MisfirePolicy: schedule.MisfirePolicy,
				AllowOverlap:  schedule.AllowOverlap,
			})
		}
	}
	if len(defaultTasks) == 0 {
		return nil
	}
	taskKeys := make([]string, 0, len(defaultTasks))
	for _, row := range defaultTasks {
		taskKeys = append(taskKeys, row.TaskKey)
	}
	existingTasks, err := u.scheduledTaskRepo.MapByTaskKey(ctx, taskKeys)
	if err != nil {
		return err
	}
	for _, row := range defaultTasks {
		if existingTasks[row.TaskKey] != nil {
			continue
		}
		if _, err = u.Upsert(ctx, row); err != nil {
			return err
		}
	}
	return nil
}

type ScheduledTaskGetReq struct {
	ID      int64
	TaskKey string
}

func (u *ScheduledTaskUsecase) Get(ctx context.Context, req *ScheduledTaskGetReq) (*model.ScheduledTask, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	taskKey := strings.TrimSpace(req.TaskKey)
	if taskKey != "" {
		row, err := u.scheduledTaskCacheRepo.GetScheduledTask(ctx, taskKey)
		if err != nil || row != nil {
			return row, err
		}
		row, err = u.scheduledTaskRepo.Get(ctx, &repo.ScheduledTaskGetReq{TaskKey: &taskKey})
		if err == nil && row != nil {
			_ = u.scheduledTaskCacheRepo.SetScheduledTask(ctx, row)
		}
		return row, err
	}
	if req.ID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row, err := u.scheduledTaskRepo.Get(ctx, &repo.ScheduledTaskGetReq{ID: &req.ID})
	if err == nil && row != nil {
		_ = u.scheduledTaskCacheRepo.SetScheduledTask(ctx, row)
	}
	return row, err
}

type ScheduledTaskPageReq struct {
	Page        *common.PageReq
	IDs         []int64
	TaskKey     *string
	HandlerName *schedulerenum.TaskHandlerName
	Title       *string
	Enabled     *bool
}

type ScheduledTaskPageResp struct {
	Rows []*model.ScheduledTask
	Page *common.PageResp
}

func (u *ScheduledTaskUsecase) Page(ctx context.Context, req *ScheduledTaskPageReq) (*ScheduledTaskPageResp, error) {
	if req == nil {
		req = &ScheduledTaskPageReq{}
	}
	pageResp, err := u.scheduledTaskRepo.Page(ctx, &repo.ScheduledTaskPageReq{
		Page: req.Page,
		ScheduledTaskGetReq: repo.ScheduledTaskGetReq{
			IDs:         req.IDs,
			TaskKey:     req.TaskKey,
			HandlerName: req.HandlerName,
			Title:       req.Title,
			Enabled:     req.Enabled,
		},
	})
	if err != nil {
		return nil, err
	}
	return &ScheduledTaskPageResp{Rows: pageResp.Rows, Page: pageResp.Page}, nil
}

func (u *ScheduledTaskUsecase) ListAvailableTasks(ctx context.Context, keyword string) ([]*model.AvailableTask, error) {
	_ = ctx
	keyword = strings.ToLower(strings.TrimSpace(keyword))
	rows := make([]*model.AvailableTask, 0, len(u.tasks))
	for _, item := range u.tasks {
		row := &model.AvailableTask{
			HandlerName: item.HandlerName(),
			Title:       item.Title(),
			Description: item.Description(),
		}
		if keyword == "" ||
			strings.Contains(strings.ToLower(row.HandlerName.String()), keyword) ||
			strings.Contains(strings.ToLower(row.Title), keyword) ||
			strings.Contains(strings.ToLower(row.Description), keyword) {
			rows = append(rows, row)
		}
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].HandlerName < rows[j].HandlerName })
	return rows, nil
}

type ScheduledTaskExecutionRecordPageReq struct {
	Page            *common.PageReq
	IDs             []int64
	ScheduledTaskID *int64
	Status          *schedulerenum.TaskExecutionStatus
	TriggerType     *schedulerenum.TaskTriggerType
}

type ScheduledTaskExecutionRecordPageResp struct {
	Rows []*model.ScheduledTaskExecutionRecord
	Page *common.PageResp
}

func (u *ScheduledTaskUsecase) PageExecutionRecords(ctx context.Context, req *ScheduledTaskExecutionRecordPageReq) (*ScheduledTaskExecutionRecordPageResp, error) {
	if req == nil {
		req = &ScheduledTaskExecutionRecordPageReq{}
	}
	pageResp, err := u.scheduledTaskExecutionRecordRepo.Page(ctx, &repo.ScheduledTaskExecutionRecordPageReq{
		Page: req.Page,
		ScheduledTaskExecutionRecordGetReq: repo.ScheduledTaskExecutionRecordGetReq{
			IDs:             req.IDs,
			ScheduledTaskID: req.ScheduledTaskID,
			Status:          req.Status,
			TriggerType:     req.TriggerType,
		},
	})
	if err != nil {
		return nil, err
	}
	return &ScheduledTaskExecutionRecordPageResp{Rows: pageResp.Rows, Page: pageResp.Page}, nil
}

type TaskTriggerReq struct {
	TaskKey string
	Payload string
}

func (u *ScheduledTaskUsecase) Trigger(ctx context.Context, req *TaskTriggerReq) (*model.ScheduledTaskExecutionRecord, error) {
	if req == nil || strings.TrimSpace(req.TaskKey) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	task, err := u.Get(ctx, &ScheduledTaskGetReq{TaskKey: req.TaskKey})
	if err != nil {
		return nil, err
	}
	if req.Payload != "" {
		if !json.Valid([]byte(req.Payload)) {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		task.Payload = req.Payload
	}
	record, err := u.createScheduledExecution(
		ctx,
		task,
		time.Now().Truncate(time.Second),
		schedulerenum.TaskTriggerTypeManual,
		uuid.NewString(),
	)
	if err != nil || record == nil {
		return record, err
	}
	u.runningWG.Add(1)
	go func() {
		_, _ = u.executeScheduledRecord(context.WithoutCancel(ctx), task, record)
	}()
	return record, nil
}

func (u *ScheduledTaskUsecase) CancelExecution(ctx context.Context, id int64) (*model.ScheduledTaskExecutionRecord, error) {
	record, err := u.scheduledTaskExecutionRecordRepo.Get(ctx, &repo.ScheduledTaskExecutionRecordGetReq{ID: &id})
	if err != nil {
		return nil, err
	}
	if record.Status != schedulerenum.TaskExecutionStatusRunning {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	u.runningCancelMu.Lock()
	cancel := u.runningCancels[id]
	u.runningCancelMu.Unlock()
	if cancel != nil {
		cancel()
	}
	if _, err = u.scheduledTaskExecutionRecordRepo.MarkCanceled(context.WithoutCancel(ctx), id, time.Now()); err != nil {
		return nil, err
	}
	return record, nil
}

func (u *ScheduledTaskUsecase) createScheduledExecution(
	ctx context.Context,
	task *model.ScheduledTask,
	scheduledAt time.Time,
	triggerType schedulerenum.TaskTriggerType,
	scheduleKey string,
) (*model.ScheduledTaskExecutionRecord, error) {
	scheduledAt = scheduledAt.Truncate(time.Second)
	if task == nil || !task.Enabled && triggerType == schedulerenum.TaskTriggerTypeSchedule {
		return nil, nil
	}
	if !task.AllowOverlap {
		running, err := u.scheduledTaskExecutionRecordRepo.HasRunning(ctx, &repo.ScheduledTaskExecutionRecordHasRunningReq{ScheduledTaskID: task.ID})
		if err != nil {
			return nil, err
		}
		if running {
			finishedAt := time.Now()
			duration := time.Duration(0)
			resp, err := u.scheduledTaskExecutionRecordRepo.Create(ctx, &repo.ScheduledTaskExecutionRecordCreateReq{
				Status: schedulerenum.TaskExecutionStatusSkipped,
				Record: &model.ScheduledTaskExecutionRecord{
					ScheduledTaskID:      task.ID,
					ScheduledTaskVersion: task.Version,
					TriggerType:          triggerType,
					ScheduleKey:          scheduleKey,
					ScheduledAt:          scheduledAt,
					FinishedAt:           &finishedAt,
					Duration:             &duration,
					Status:               schedulerenum.TaskExecutionStatusSkipped,
					Attempt:              0,
					MaxAttempts:          task.MaxAttempts,
					Timeout:              task.Timeout,
					StaleAfter:           task.StaleAfter,
					MisfirePolicy:        task.MisfirePolicy,
					WorkerID:             u.workerID,
					Payload:              task.Payload,
					LastError:            "scheduler task has another running execution",
				},
			})
			if err != nil || resp == nil || resp.Conflict {
				return nil, err
			}
			return resp.Row, nil
		}
	}
	startedAt := time.Now()
	resp, err := u.scheduledTaskExecutionRecordRepo.Create(ctx, &repo.ScheduledTaskExecutionRecordCreateReq{
		Status: schedulerenum.TaskExecutionStatusRunning,
		Record: &model.ScheduledTaskExecutionRecord{
			ScheduledTaskID:      task.ID,
			ScheduledTaskVersion: task.Version,
			TriggerType:          triggerType,
			ScheduleKey:          scheduleKey,
			ScheduledAt:          scheduledAt,
			StartedAt:            &startedAt,
			Status:               schedulerenum.TaskExecutionStatusRunning,
			Attempt:              1,
			MaxAttempts:          task.MaxAttempts,
			Timeout:              task.Timeout,
			StaleAfter:           task.StaleAfter,
			MisfirePolicy:        task.MisfirePolicy,
			WorkerID:             u.workerID,
			Payload:              task.Payload,
		},
	})
	if err != nil || resp == nil {
		return nil, err
	}
	if resp.Conflict {
		return u.scheduledTaskExecutionRecordRepo.Get(ctx, &repo.ScheduledTaskExecutionRecordGetReq{ScheduledTaskID: &task.ID, ScheduleKey: &scheduleKey})
	}
	return resp.Row, nil
}

func (u *ScheduledTaskUsecase) executeScheduledRecord(
	ctx context.Context,
	task *model.ScheduledTask,
	record *model.ScheduledTaskExecutionRecord,
) (*model.ScheduledTaskExecutionRecord, error) {
	defer u.runningWG.Done()
	timeout := record.Timeout
	if timeout <= 0 {
		timeout = 300 * time.Second
	}
	execCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	u.runningCancelMu.Lock()
	u.runningCancels[record.ID] = cancel
	u.runningCancelMu.Unlock()
	defer func() {
		u.runningCancelMu.Lock()
		delete(u.runningCancels, record.ID)
		u.runningCancelMu.Unlock()
	}()
	currentTask, ok := u.tasks[task.HandlerName.String()]
	var err error
	if ok {
		err = currentTask.Execute(execCtx, record.Payload)
	} else {
		err = fmt.Errorf("unknown scheduler task handler: %s", task.HandlerName)
	}
	finishedAt := time.Now()
	duration := time.Duration(0)
	if record.StartedAt != nil {
		duration = finishedAt.Sub(*record.StartedAt)
	}
	statusValue := schedulerenum.TaskExecutionStatusSuccess
	lastError := ""
	if errors.Is(execCtx.Err(), context.Canceled) {
		statusValue = schedulerenum.TaskExecutionStatusCanceled
		lastError = execCtx.Err().Error()
	} else if errors.Is(execCtx.Err(), context.DeadlineExceeded) {
		statusValue = schedulerenum.TaskExecutionStatusTimeout
		lastError = execCtx.Err().Error()
		if record.Attempt < record.MaxAttempts {
			statusValue = schedulerenum.TaskExecutionStatusRetryPending
		}
	} else if err != nil {
		statusValue = schedulerenum.TaskExecutionStatusFailed
		lastError = err.Error()
		if record.Attempt < record.MaxAttempts {
			statusValue = schedulerenum.TaskExecutionStatusRetryPending
		}
	}
	if len(lastError) > 2048 {
		lastError = lastError[:2048]
	}
	resp, markErr := u.scheduledTaskExecutionRecordRepo.MarkFinished(
		context.WithoutCancel(ctx),
		&repo.ScheduledTaskExecutionRecordMarkFinishedReq{
			ID:         record.ID,
			WorkerID:   record.WorkerID,
			Attempt:    record.Attempt,
			Status:     statusValue,
			FinishedAt: finishedAt,
			Duration:   duration,
			LastError:  lastError,
		},
	)
	if markErr != nil {
		u.logger.ErrorContext(context.WithoutCancel(ctx), "mark scheduled task execution finished failed", constant.LogFieldErr, markErr, constant.LogFieldExecutionID, record.ID)
		return nil, markErr
	}
	if resp == nil {
		return nil, nil
	}
	return resp.Row, nil
}

func (u *ScheduledTaskUsecase) HandleScheduledTaskMessage(ctx context.Context, message *repo.ScheduledTaskScheduleMessage) (*repo.MessageHandleResult, error) {
	if message == nil {
		return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionComplete}, nil
	}
	task, err := u.Get(ctx, &ScheduledTaskGetReq{
		ID:      message.ScheduledTaskID,
		TaskKey: message.ScheduledTaskKey,
	})
	if err != nil {
		return nil, err
	}
	if task == nil || !task.Enabled || task.Version != message.ScheduledTaskVersion {
		return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionComplete}, nil
	}
	// 过期消息按任务策略处理：跳过、只补最新一次或逐条补齐。
	if task.StaleAfter != nil && time.Since(message.ScheduledAt) > *task.StaleAfter+u.clockSkewGrace {
		if task.MisfirePolicy == schedulerenum.TaskMisfirePolicySkip || task.MisfirePolicy == schedulerenum.TaskMisfirePolicyExecuteLatest && !message.LatestForSubject {
			finishedAt := time.Now()
			duration := time.Duration(0)
			_, _ = u.scheduledTaskExecutionRecordRepo.Create(ctx, &repo.ScheduledTaskExecutionRecordCreateReq{
				Status: schedulerenum.TaskExecutionStatusSkipped,
				Record: &model.ScheduledTaskExecutionRecord{
					ScheduledTaskID:      task.ID,
					ScheduledTaskVersion: task.Version,
					TriggerType:          schedulerenum.TaskTriggerTypeSchedule,
					ScheduleKey:          message.ScheduleKey,
					ScheduledAt:          message.ScheduledAt,
					FinishedAt:           &finishedAt,
					Duration:             &duration,
					Status:               schedulerenum.TaskExecutionStatusSkipped,
					Attempt:              0,
					MaxAttempts:          task.MaxAttempts,
					Timeout:              task.Timeout,
					StaleAfter:           task.StaleAfter,
					MisfirePolicy:        task.MisfirePolicy,
					WorkerID:             u.workerID,
					Payload:              task.Payload,
					LastError:            "scheduler task message is stale",
				},
			})
			return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionComplete}, nil
		}
	}
	record, err := u.createScheduledExecution(ctx, task, message.ScheduledAt, schedulerenum.TaskTriggerTypeSchedule, message.ScheduleKey)
	if err != nil {
		return nil, err
	}
	if record == nil || record.Status == schedulerenum.TaskExecutionStatusSkipped {
		return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionComplete}, nil
	}
	if record.Status == schedulerenum.TaskExecutionStatusSuccess ||
		record.Status == schedulerenum.TaskExecutionStatusFailed ||
		record.Status == schedulerenum.TaskExecutionStatusTimeout ||
		record.Status == schedulerenum.TaskExecutionStatusCanceled {
		return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionComplete}, nil
	}
	// 运行中状态表示执行权租约未结束，其他节点不能直接确认消息。
	if record.Status == schedulerenum.TaskExecutionStatusRunning && record.WorkerID != u.workerID {
		if record.StartedAt != nil && record.StartedAt.Add(record.Timeout).After(time.Now()) {
			return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionRetry, RetryAfter: time.Until(record.StartedAt.Add(record.Timeout))}, nil
		}
		if record.Attempt >= record.MaxAttempts {
			finishedAt := time.Now()
			duration := time.Duration(0)
			if record.StartedAt != nil {
				duration = finishedAt.Sub(*record.StartedAt)
			}
			_, err = u.scheduledTaskExecutionRecordRepo.MarkFinished(ctx, &repo.ScheduledTaskExecutionRecordMarkFinishedReq{
				ID:         record.ID,
				WorkerID:   record.WorkerID,
				Attempt:    record.Attempt,
				Status:     schedulerenum.TaskExecutionStatusTimeout,
				FinishedAt: finishedAt,
				Duration:   duration,
				LastError:  "scheduler task running lease expired",
			})
			if err != nil {
				return nil, err
			}
			return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionComplete}, nil
		}
		claim, err := u.scheduledTaskExecutionRecordRepo.Claim(ctx, &repo.ScheduledTaskExecutionRecordClaimReq{ID: record.ID, WorkerID: u.workerID, StartedAt: time.Now()})
		if err != nil || claim == nil || !claim.Claimed || claim.Row == nil {
			return nil, err
		}
		record = claim.Row
	}
	if record.Status == schedulerenum.TaskExecutionStatusRetryPending || record.Status == schedulerenum.TaskExecutionStatusPending {
		claim, err := u.scheduledTaskExecutionRecordRepo.Claim(ctx, &repo.ScheduledTaskExecutionRecordClaimReq{ID: record.ID, WorkerID: u.workerID, StartedAt: time.Now()})
		if err != nil || claim == nil || !claim.Claimed || claim.Row == nil {
			return nil, err
		}
		record = claim.Row
	}
	u.runningWG.Add(1)
	updated, err := u.executeScheduledRecord(ctx, task, record)
	if err != nil {
		return nil, err
	}
	if updated == nil {
		return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionComplete}, nil
	}
	// 业务重试由执行记录状态决定，NATS 只负责按退避时间重新投递。
	if updated.Status == schedulerenum.TaskExecutionStatusRetryPending {
		base := 5 * time.Second
		capValue := 5 * time.Minute
		factor := 1 << max(0, int(updated.Attempt)-1)
		delay := time.Duration(factor) * base
		if delay > capValue {
			delay = capValue
		}
		return &repo.MessageHandleResult{
			Action:     schedulerenum.MessageHandleActionRetry,
			RetryAfter: delay + time.Duration(rand.Int64N(int64(delay/2)+1)),
		}, nil
	}
	return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionComplete}, nil
}

func (u *ScheduledTaskUsecase) EnsureSchedule(ctx context.Context) error {
	// 启动时只重建调度源，不清理已投递的 target 消息和 durable cursor。
	if err := u.scheduledTaskScheduleRepo.Ensure(ctx); err != nil {
		return err
	}
	schedulePrefix := "scheduler.schedule.scheduled_task"
	if u.conf.GetScheduler() != nil && u.conf.GetScheduler().GetScheduledTaskScheduleSubjectPrefix() != "" {
		schedulePrefix = u.conf.GetScheduler().GetScheduledTaskScheduleSubjectPrefix()
	}
	if err := u.scheduledTaskScheduleRepo.Cancel(ctx, schedulePrefix+".>"); err != nil {
		return err
	}
	executeSubject := "scheduler.execute.scheduled_task"
	if u.conf.GetScheduler() != nil && u.conf.GetScheduler().GetScheduledTaskExecuteSubject() != "" {
		executeSubject = u.conf.GetScheduler().GetScheduledTaskExecuteSubject()
	}
	enabled := true
	rows, err := u.scheduledTaskRepo.List(ctx, &repo.ScheduledTaskGetReq{Enabled: &enabled})
	if err != nil {
		return err
	}
	for _, row := range rows {
		if row == nil || !row.Enabled {
			continue
		}
		if err = u.scheduledTaskScheduleRepo.Schedule(ctx, &repo.ScheduledTaskScheduleReq{
			ScheduledTask: row,
			Subject:       fmt.Sprintf("%s.%d", schedulePrefix, row.ID),
			Target:        executeSubject,
		}); err != nil {
			return err
		}
	}
	return nil
}
func (u *ScheduledTaskUsecase) StartConsuming(ctx context.Context) error {
	return u.scheduledTaskScheduleRepo.Consume(ctx, u.HandleScheduledTaskMessage)
}

func (u *ScheduledTaskUsecase) StopConsuming(ctx context.Context) error {
	if err := u.scheduledTaskScheduleRepo.Stop(ctx); err != nil {
		return err
	}
	u.runningCancelMu.Lock()
	cancels := make([]context.CancelFunc, 0, len(u.runningCancels))
	for _, cancel := range u.runningCancels {
		cancels = append(cancels, cancel)
	}
	u.runningCancelMu.Unlock()
	for _, cancel := range cancels {
		cancel()
	}
	done := make(chan struct{})
	go func() {
		u.runningWG.Wait()
		close(done)
	}()
	stopRunningCtx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()
	select {
	case <-done:
		return nil
	case <-stopRunningCtx.Done():
		return stopRunningCtx.Err()
	}
}
