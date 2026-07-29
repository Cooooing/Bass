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
	"github.com/samber/lo"
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
	if row == nil || strings.TrimSpace(row.Name) == "" || strings.TrimSpace(row.Title) == "" || strings.TrimSpace(row.CronSpec) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
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
	var previous *model.ScheduledTask
	if row.ID > 0 {
		previous, _ = u.scheduledTaskRepo.Get(ctx, &repo.ScheduledTaskGetReq{ID: &row.ID})
	}
	var saved *model.ScheduledTask
	err := u.tx(ctx, func(txCtx context.Context) error {
		var err error
		saved, err = u.scheduledTaskRepo.Upsert(txCtx, row)
		if err != nil {
			return err
		}
		if _, err = u.scheduledTaskVersionRepo.Create(txCtx, saved); err != nil {
			return err
		}
		subject := u.scheduledTaskScheduleSubject(saved.ID)
		if saved.Enabled {
			return u.scheduledTaskScheduleRepo.Schedule(txCtx, &repo.ScheduledTaskScheduleReq{ScheduledTask: saved, Subject: subject, Target: u.scheduledTaskExecuteSubject()})
		}
		return u.scheduledTaskScheduleRepo.Cancel(txCtx, subject)
	})
	if err != nil {
		if previous != nil {
			if previous.Enabled {
				_ = u.scheduledTaskScheduleRepo.Schedule(context.WithoutCancel(ctx), &repo.ScheduledTaskScheduleReq{
					ScheduledTask: previous,
					Subject:       u.scheduledTaskScheduleSubject(previous.ID),
					Target:        u.scheduledTaskExecuteSubject(),
				})
			} else {
				_ = u.scheduledTaskScheduleRepo.Cancel(context.WithoutCancel(ctx), u.scheduledTaskScheduleSubject(previous.ID))
			}
		} else if saved != nil {
			_ = u.scheduledTaskScheduleRepo.Cancel(context.WithoutCancel(ctx), u.scheduledTaskScheduleSubject(saved.ID))
		}
		return nil, err
	}
	if previous != nil && previous.Title != saved.Title {
		_ = u.scheduledTaskCacheRepo.DeleteScheduledTask(ctx, previous.Title)
	}
	_ = u.scheduledTaskCacheRepo.DeleteScheduledTask(ctx, saved.Title)
	return saved, nil
}

func (u *ScheduledTaskUsecase) SeedDefaultTasks(ctx context.Context) error {
	names := make([]string, 0, len(u.tasks))
	for name := range u.tasks {
		names = append(names, name)
	}
	sort.Strings(names)
	defaultTasks := make([]*model.ScheduledTask, 0)
	for _, name := range names {
		item := u.tasks[name]
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
				Name:          item.Name(),
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
	titles := lo.Map(defaultTasks, func(row *model.ScheduledTask, _ int) string { return row.Title })
	if duplicatedTitles := lo.FindDuplicates(titles); len(duplicatedTitles) > 0 {
		return fmt.Errorf("scheduler default task title duplicated: %s", strings.Join(duplicatedTitles, ","))
	}
	existingTasks, err := u.scheduledTaskRepo.MapByTitle(ctx, titles)
	if err != nil {
		return err
	}
	for _, row := range defaultTasks {
		if existingTasks[row.Title] != nil {
			current := existingTasks[row.Title]
			if current.Enabled {
				if err = u.scheduledTaskScheduleRepo.Schedule(ctx, &repo.ScheduledTaskScheduleReq{
					ScheduledTask: current,
					Subject:       u.scheduledTaskScheduleSubject(current.ID),
					Target:        u.scheduledTaskExecuteSubject(),
				}); err != nil {
					return err
				}
			}
			continue
		}
		if _, err = u.Upsert(ctx, row); err != nil {
			return err
		}
	}
	return nil
}

type ScheduledTaskGetReq struct {
	ID    int64
	Title string
}

func (u *ScheduledTaskUsecase) Get(ctx context.Context, req *ScheduledTaskGetReq) (*model.ScheduledTask, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	title := strings.TrimSpace(req.Title)
	if title != "" {
		row, err := u.scheduledTaskCacheRepo.GetScheduledTask(ctx, title)
		if err != nil || row != nil {
			return row, err
		}
		row, err = u.scheduledTaskRepo.Get(ctx, &repo.ScheduledTaskGetReq{Title: &title})
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
	Page    *common.PageReq
	IDs     []int64
	Name    *string
	Title   *string
	Enabled *bool
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
			IDs:     req.IDs,
			Name:    req.Name,
			Title:   req.Title,
			Enabled: req.Enabled,
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
		row := &model.AvailableTask{Name: item.Name(), Title: item.Title(), Description: item.Description()}
		if keyword == "" ||
			strings.Contains(strings.ToLower(row.Name), keyword) ||
			strings.Contains(strings.ToLower(row.Title), keyword) ||
			strings.Contains(strings.ToLower(row.Description), keyword) {
			rows = append(rows, row)
		}
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].Name < rows[j].Name })
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
	ID      int64
	Payload string
}

func (u *ScheduledTaskUsecase) Trigger(ctx context.Context, req *TaskTriggerReq) (*model.ScheduledTaskExecutionRecord, error) {
	if req == nil || req.ID == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	task, err := u.Get(ctx, &ScheduledTaskGetReq{ID: req.ID})
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
			resp, err := u.scheduledTaskExecutionRecordRepo.Create(ctx, &repo.ScheduledTaskExecutionRecordCreateReq{
				Status: schedulerenum.TaskExecutionStatusSkipped,
				Record: &model.ScheduledTaskExecutionRecord{
					ScheduledTaskID:      task.ID,
					ScheduledTaskVersion: task.Version,
					TriggerType:          triggerType,
					ScheduleKey:          scheduleKey,
					ScheduledAt:          scheduledAt,
					FinishedAt:           new(time.Now()),
					Duration:             new(time.Duration(0)),
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
	resp, err := u.scheduledTaskExecutionRecordRepo.Create(ctx, &repo.ScheduledTaskExecutionRecordCreateReq{
		Status: schedulerenum.TaskExecutionStatusRunning,
		Record: &model.ScheduledTaskExecutionRecord{
			ScheduledTaskID:      task.ID,
			ScheduledTaskVersion: task.Version,
			TriggerType:          triggerType,
			ScheduleKey:          scheduleKey,
			ScheduledAt:          scheduledAt,
			StartedAt:            new(time.Now()),
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
	currentTask, ok := u.tasks[task.Name]
	var err error
	if ok {
		err = currentTask.Execute(execCtx, record.Payload)
	} else {
		err = fmt.Errorf("unknown scheduler task: %s", task.Name)
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
	task, err := u.Get(ctx, &ScheduledTaskGetReq{ID: message.ScheduledTaskID, Title: message.ScheduledTaskTitle})
	if err != nil {
		return nil, err
	}
	if task == nil || !task.Enabled || task.Version != message.ScheduledTaskVersion {
		return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionComplete}, nil
	}
	// 过期消息按任务策略处理：跳过、只补最新一次或逐条补齐。
	if task.StaleAfter != nil && time.Since(message.ScheduledAt) > *task.StaleAfter+u.clockSkewGrace {
		if task.MisfirePolicy == schedulerenum.TaskMisfirePolicySkip || task.MisfirePolicy == schedulerenum.TaskMisfirePolicyExecuteLatest && !message.LatestForSubject {
			_, _ = u.scheduledTaskExecutionRecordRepo.Create(ctx, &repo.ScheduledTaskExecutionRecordCreateReq{
				Status: schedulerenum.TaskExecutionStatusSkipped,
				Record: &model.ScheduledTaskExecutionRecord{
					ScheduledTaskID:      task.ID,
					ScheduledTaskVersion: task.Version,
					TriggerType:          schedulerenum.TaskTriggerTypeSchedule,
					ScheduleKey:          message.ScheduleKey,
					ScheduledAt:          message.ScheduledAt,
					FinishedAt:           new(time.Now()),
					Duration:             new(time.Duration(0)),
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
		return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionRetry, RetryAfter: u.retryBackoff(updated.Attempt)}, nil
	}
	return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionComplete}, nil
}

func (u *ScheduledTaskUsecase) EnsureSchedule(ctx context.Context) error {
	// 启动时重建调度源，避免 NATS 残留旧配置继续投递。
	if err := u.scheduledTaskScheduleRepo.Ensure(ctx); err != nil {
		return err
	}
	prefix := "scheduler.schedule.scheduled_task"
	if u.conf.GetScheduler() != nil && u.conf.GetScheduler().GetScheduledTaskScheduleSubjectPrefix() != "" {
		prefix = u.conf.GetScheduler().GetScheduledTaskScheduleSubjectPrefix()
	}
	if err := u.scheduledTaskScheduleRepo.Cancel(ctx, prefix+".>"); err != nil {
		return err
	}
	rows, err := u.scheduledTaskRepo.List(ctx, &repo.ScheduledTaskGetReq{Enabled: new(true)})
	if err != nil {
		return err
	}
	for _, row := range rows {
		if row == nil {
			continue
		}
		if err = u.scheduledTaskScheduleRepo.Schedule(ctx, &repo.ScheduledTaskScheduleReq{
			ScheduledTask: row,
			Subject:       u.scheduledTaskScheduleSubject(row.ID),
			Target:        u.scheduledTaskExecuteSubject(),
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
	stopRunningCtx, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()
	return u.stopRunning(stopRunningCtx)
}

func (u *ScheduledTaskUsecase) scheduledTaskScheduleSubject(taskID int64) string {
	prefix := "scheduler.schedule.scheduled_task"
	if u.conf.GetScheduler() != nil && u.conf.GetScheduler().GetScheduledTaskScheduleSubjectPrefix() != "" {
		prefix = u.conf.GetScheduler().GetScheduledTaskScheduleSubjectPrefix()
	}
	return fmt.Sprintf("%s.%d", prefix, taskID)
}

func (u *ScheduledTaskUsecase) scheduledTaskExecuteSubject() string {
	if u.conf.GetScheduler() != nil && u.conf.GetScheduler().GetScheduledTaskExecuteSubject() != "" {
		return u.conf.GetScheduler().GetScheduledTaskExecuteSubject()
	}
	return "scheduler.execute.scheduled_task"
}

func (u *ScheduledTaskUsecase) retryBackoff(attempt int32) time.Duration {
	base := 5 * time.Second
	capValue := 5 * time.Minute
	factor := 1 << max(0, int(attempt)-1)
	delay := time.Duration(factor) * base
	if delay > capValue {
		delay = capValue
	}
	return delay + time.Duration(rand.Int64N(int64(delay/2)+1))
}

func (u *ScheduledTaskUsecase) stopRunning(ctx context.Context) error {
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
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
