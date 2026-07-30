package usecase

import (
	"common/pkg/apperror"
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
)

type DelayedTaskUsecase struct {
	logger               *slog.Logger
	conf                 *config.Bootstrap
	tx                   base.Tx
	repo                 repo.DelayedTaskRepo
	delayedTaskCacheRepo repo.DelayedTaskCacheRepo
	versionRepo          repo.DelayedTaskVersionRepo
	executionRepo        repo.DelayedTaskExecutionRecordRepo
	scheduleRepo         repo.DelayedTaskScheduleRepo
	tasks                map[string]taskimpl.Task
	workerID             string
	runningMu            sync.Mutex
	runningCancel        map[int64]context.CancelFunc
	clockSkewGrace       time.Duration
}

func NewDelayedTaskUsecase(
	logger *slog.Logger,
	conf *config.Bootstrap,
	tx base.Tx,
	delayedTaskRepo repo.DelayedTaskRepo,
	delayedTaskCacheRepo repo.DelayedTaskCacheRepo,
	versionRepo repo.DelayedTaskVersionRepo,
	executionRepo repo.DelayedTaskExecutionRecordRepo,
	scheduleRepo repo.DelayedTaskScheduleRepo,
	tasks map[string]taskimpl.Task,
) *DelayedTaskUsecase {
	workerID := "scheduler"
	if host, err := os.Hostname(); err == nil && host != "" {
		workerID = host
	}
	return &DelayedTaskUsecase{
		logger:               logger,
		conf:                 conf,
		tx:                   tx,
		repo:                 delayedTaskRepo,
		delayedTaskCacheRepo: delayedTaskCacheRepo,
		versionRepo:          versionRepo,
		executionRepo:        executionRepo,
		scheduleRepo:         scheduleRepo,
		tasks:                tasks,
		workerID:             workerID,
		runningCancel:        make(map[int64]context.CancelFunc),
		clockSkewGrace:       10 * time.Second,
	}
}

func (u *DelayedTaskUsecase) Upsert(ctx context.Context, row *model.DelayedTask) (*model.DelayedTask, error) {
	if row == nil ||
		row.HandlerName == "" ||
		strings.TrimSpace(row.Title) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row.TaskKey = strings.TrimSpace(row.TaskKey)
	if row.TaskKey == "" {
		row.TaskKey = uuid.NewString()
	}
	if row.Timeout <= 0 {
		row.Timeout = 30 * time.Second
	}
	if row.MaxAttempts <= 0 {
		row.MaxAttempts = 3
	}
	if row.MisfirePolicy == "" {
		row.MisfirePolicy = schedulerenum.TaskMisfirePolicyExecuteAll
	}
	var saved *model.DelayedTask
	err := u.tx(ctx, func(txCtx context.Context) error {
		var err error
		saved, err = u.repo.Upsert(txCtx, row)
		if err != nil {
			return err
		}
		_, err = u.versionRepo.Create(txCtx, saved)
		return err
	})
	if err != nil {
		return nil, err
	}
	_ = u.delayedTaskCacheRepo.DeleteDelayedTask(ctx, saved.TaskKey)
	return saved, nil
}

func (u *DelayedTaskUsecase) SeedDefaultTasks(ctx context.Context) error {
	handlerNames := make([]string, 0, len(u.tasks))
	for handlerName := range u.tasks {
		handlerNames = append(handlerNames, handlerName)
	}
	sort.Strings(handlerNames)
	defaultTasks := make([]*model.DelayedTask, 0)
	for _, handlerName := range handlerNames {
		item := u.tasks[handlerName]
		for _, task := range item.DefaultDelayedTasks() {
			if task == nil {
				continue
			}
			title := strings.TrimSpace(task.Title)
			if title == "" {
				title = item.Title()
			}
			description := strings.TrimSpace(task.Description)
			if description == "" {
				description = item.Description()
			}
			defaultTasks = append(defaultTasks, &model.DelayedTask{
				TaskKey:       task.TaskKey.String(),
				HandlerName:   item.HandlerName(),
				Title:         title,
				Description:   description,
				Enabled:       task.Enabled,
				Timeout:       task.Timeout,
				StaleAfter:    task.StaleAfter,
				MaxAttempts:   task.MaxAttempts,
				MisfirePolicy: task.MisfirePolicy,
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
	existingTasks, err := u.repo.MapByTaskKey(ctx, taskKeys)
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

type DelayedTaskGetReq struct {
	ID      int64
	TaskKey string
}

func (u *DelayedTaskUsecase) Get(ctx context.Context, req *DelayedTaskGetReq) (*model.DelayedTask, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	taskKey := strings.TrimSpace(req.TaskKey)
	if taskKey != "" {
		row, err := u.delayedTaskCacheRepo.GetDelayedTask(ctx, taskKey)
		if err != nil || row != nil {
			return row, err
		}
		row, err = u.repo.Get(ctx, &repo.DelayedTaskGetReq{TaskKey: &taskKey})
		if err == nil && row != nil {
			_ = u.delayedTaskCacheRepo.SetDelayedTask(ctx, row)
		}
		return row, err
	}
	if req.ID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row, err := u.repo.Get(ctx, &repo.DelayedTaskGetReq{ID: &req.ID})
	if err == nil && row != nil {
		_ = u.delayedTaskCacheRepo.SetDelayedTask(ctx, row)
	}
	return row, err
}

type DelayedTaskPageReq struct {
	Page        *common.PageReq
	IDs         []int64
	TaskKey     *string
	HandlerName *schedulerenum.TaskHandlerName
	Title       *string
	Enabled     *bool
}

type DelayedTaskPageResp struct {
	Rows []*model.DelayedTask
	Page *common.PageResp
}

func (u *DelayedTaskUsecase) Page(ctx context.Context, req *DelayedTaskPageReq) (*DelayedTaskPageResp, error) {
	if req == nil {
		req = &DelayedTaskPageReq{}
	}
	resp, err := u.repo.Page(ctx, &repo.DelayedTaskPageReq{
		Page: req.Page,
		DelayedTaskGetReq: repo.DelayedTaskGetReq{
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
	return &DelayedTaskPageResp{Rows: resp.Rows, Page: resp.Page}, nil
}

func (u *DelayedTaskUsecase) ListAvailableTasks(ctx context.Context, keyword string) ([]*model.AvailableTask, error) {
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

type DelayedTaskScheduleReq struct {
	TaskKey     string
	Payload     string
	ScheduledAt time.Time
	TriggerType schedulerenum.TaskTriggerType
}

func (u *DelayedTaskUsecase) Schedule(ctx context.Context, req *DelayedTaskScheduleReq) (*model.DelayedTaskExecutionRecord, error) {
	if req == nil || req.ScheduledAt.IsZero() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if strings.TrimSpace(req.TaskKey) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	payload := strings.TrimSpace(req.Payload)
	if payload == "" {
		payload = "{}"
	}
	if !json.Valid([]byte(payload)) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	task, err := u.Get(ctx, &DelayedTaskGetReq{TaskKey: req.TaskKey})
	if err != nil {
		return nil, err
	}
	if task == nil || !task.Enabled {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	triggerType := req.TriggerType
	if triggerType == "" {
		triggerType = schedulerenum.TaskTriggerTypeSchedule
	}
	// 延迟任务配置不保存单次执行参数，执行实例在记录表中保存完整快照。
	record := &model.DelayedTaskExecutionRecord{
		DelayedTaskID:      task.ID,
		DelayedTaskVersion: task.Version,
		IdempotencyKey:     uuid.NewString(),
		TriggerType:        triggerType,
		ScheduleKey:        uuid.NewString(),
		ScheduledAt:        req.ScheduledAt.Truncate(time.Second),
		Status:             schedulerenum.TaskExecutionStatusPending,
		Attempt:            0,
		MaxAttempts:        task.MaxAttempts,
		Timeout:            task.Timeout,
		StaleAfter:         task.StaleAfter,
		MisfirePolicy:      task.MisfirePolicy,
		Payload:            payload,
	}
	var created *repo.DelayedTaskExecutionRecordCreateResp
	err = u.tx(ctx, func(txCtx context.Context) error {
		var err error
		created, err = u.executionRepo.CreatePending(txCtx, record)
		if err != nil || created == nil || created.Row == nil || !created.Created {
			return err
		}
		return u.scheduleRepo.Schedule(txCtx, &repo.DelayedTaskScheduleReq{
			DelayedTask: task,
			Record:      created.Row,
			Subject:     u.delayedTaskScheduleSubject(created.Row.ID),
			Target:      u.delayedTaskExecuteSubject(),
		})
	})
	if err != nil {
		if created != nil && created.Created && created.Row != nil {
			_ = u.scheduleRepo.Cancel(context.WithoutCancel(ctx), u.delayedTaskScheduleSubject(created.Row.ID))
		}
		return nil, err
	}
	return created.Row, nil
}

func (u *DelayedTaskUsecase) Trigger(ctx context.Context, taskKey string, payload string) (*model.DelayedTaskExecutionRecord, error) {
	return u.Schedule(ctx, &DelayedTaskScheduleReq{
		TaskKey:     taskKey,
		Payload:     payload,
		ScheduledAt: time.Now().Truncate(time.Second),
		TriggerType: schedulerenum.TaskTriggerTypeManual,
	})
}

func (u *DelayedTaskUsecase) CancelExecution(ctx context.Context, id int64, idempotencyKey string) (*model.DelayedTaskExecutionRecord, error) {
	var query repo.DelayedTaskExecutionRecordGetReq
	if id > 0 {
		query.ID = &id
	} else {
		query.IdempotencyKey = new(strings.TrimSpace(idempotencyKey))
	}
	row, err := u.executionRepo.MarkCanceled(ctx, &query, time.Now())
	if err != nil || row == nil {
		return row, err
	}
	if row.Status == schedulerenum.TaskExecutionStatusCanceled {
		_ = u.scheduleRepo.Cancel(ctx, u.delayedTaskScheduleSubject(row.ID))
	}
	u.runningMu.Lock()
	cancel := u.runningCancel[row.ID]
	u.runningMu.Unlock()
	if cancel != nil {
		cancel()
	}
	return row, nil
}

func (u *DelayedTaskUsecase) HandleDelayedTaskMessage(ctx context.Context, message *repo.DelayedTaskScheduleMessage) (*repo.MessageHandleResult, error) {
	if message == nil {
		return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionComplete}, nil
	}
	record, err := u.executionRepo.Get(ctx, &repo.DelayedTaskExecutionRecordGetReq{ID: &message.ExecutionRecordID})
	if err != nil {
		return nil, err
	}
	if record == nil {
		return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionComplete}, nil
	}
	if record.Status == schedulerenum.TaskExecutionStatusSuccess ||
		record.Status == schedulerenum.TaskExecutionStatusFailed ||
		record.Status == schedulerenum.TaskExecutionStatusTimeout ||
		record.Status == schedulerenum.TaskExecutionStatusCanceled ||
		record.Status == schedulerenum.TaskExecutionStatusSkipped {
		return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionComplete}, nil
	}
	// 延迟任务每条执行记录都是独立业务意图，只有跳过策略会因过期直接跳过。
	if record.StaleAfter != nil && record.MisfirePolicy == schedulerenum.TaskMisfirePolicySkip && time.Since(record.ScheduledAt) > *record.StaleAfter+u.clockSkewGrace {
		finishedAt := time.Now()
		_, err = u.executionRepo.MarkFinished(ctx, &repo.DelayedTaskExecutionRecordMarkFinishedReq{
			ID:         record.ID,
			WorkerID:   record.WorkerID,
			Attempt:    record.Attempt,
			Status:     schedulerenum.TaskExecutionStatusSkipped,
			FinishedAt: finishedAt,
			Duration:   0,
			LastError:  "scheduler delayed task message is stale",
		})
		if err != nil {
			return nil, err
		}
		return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionComplete}, nil
	}
	// 运行中记录未超时时延迟重投，超时后由新消费者接管或终结为超时。
	if record.Status == schedulerenum.TaskExecutionStatusRunning {
		if record.StartedAt != nil && record.StartedAt.Add(record.Timeout).After(time.Now()) {
			return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionRetry, RetryAfter: time.Until(record.StartedAt.Add(record.Timeout))}, nil
		}
		if record.Attempt >= record.MaxAttempts {
			finishedAt := time.Now()
			duration := time.Duration(0)
			if record.StartedAt != nil {
				duration = finishedAt.Sub(*record.StartedAt)
			}
			_, err = u.executionRepo.MarkFinished(ctx, &repo.DelayedTaskExecutionRecordMarkFinishedReq{
				ID:         record.ID,
				WorkerID:   record.WorkerID,
				Attempt:    record.Attempt,
				Status:     schedulerenum.TaskExecutionStatusTimeout,
				FinishedAt: finishedAt,
				Duration:   duration,
				LastError:  "scheduler delayed task running lease expired",
			})
			if err != nil {
				return nil, err
			}
			return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionDiscard}, nil
		}
	}
	claim, err := u.executionRepo.Claim(ctx, &repo.DelayedTaskExecutionRecordClaimReq{ID: record.ID, WorkerID: u.workerID, StartedAt: time.Now()})
	if err != nil {
		return nil, err
	}
	if claim == nil || !claim.Claimed || claim.Row == nil {
		return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionRetry, RetryAfter: 5 * time.Second}, nil
	}
	updated, err := u.executeDelayedRecord(ctx, claim.Row, message.DelayedTaskKey)
	if err != nil {
		return nil, err
	}
	if updated == nil {
		return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionComplete}, nil
	}
	// 执行失败但未达到最大次数时，返回重试结果交给 NATS 延迟重投。
	if updated.Status == schedulerenum.TaskExecutionStatusRetryPending {
		return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionRetry, RetryAfter: u.retryBackoff(updated.Attempt)}, nil
	}
	return &repo.MessageHandleResult{Action: schedulerenum.MessageHandleActionComplete}, nil
}

func (u *DelayedTaskUsecase) executeDelayedRecord(ctx context.Context, record *model.DelayedTaskExecutionRecord, taskKey string) (*model.DelayedTaskExecutionRecord, error) {
	var taskConfig *model.DelayedTask
	var err error
	if taskKey != "" {
		taskConfig, err = u.Get(ctx, &DelayedTaskGetReq{TaskKey: taskKey})
	} else {
		taskConfig, err = u.Get(ctx, &DelayedTaskGetReq{ID: record.DelayedTaskID})
	}
	if err != nil || taskConfig == nil || !taskConfig.Enabled || taskConfig.Version != record.DelayedTaskVersion {
		return u.executionRepo.MarkFinished(context.WithoutCancel(ctx), &repo.DelayedTaskExecutionRecordMarkFinishedReq{
			ID:         record.ID,
			WorkerID:   record.WorkerID,
			Attempt:    record.Attempt,
			Status:     schedulerenum.TaskExecutionStatusFailed,
			FinishedAt: time.Now(),
			Duration:   0,
			LastError:  "scheduler delayed task config is unavailable",
		})
	}
	task, ok := u.tasks[taskConfig.HandlerName.String()]
	if !ok {
		return u.executionRepo.MarkFinished(context.WithoutCancel(ctx), &repo.DelayedTaskExecutionRecordMarkFinishedReq{
			ID:         record.ID,
			WorkerID:   record.WorkerID,
			Attempt:    record.Attempt,
			Status:     schedulerenum.TaskExecutionStatusFailed,
			FinishedAt: time.Now(),
			Duration:   0,
			LastError:  fmt.Sprintf("unknown delayed task handler: %s", taskConfig.HandlerName),
		})
	}
	timeout := record.Timeout
	if timeout <= 0 {
		timeout = 30 * time.Second
	}
	execCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	u.runningMu.Lock()
	u.runningCancel[record.ID] = cancel
	u.runningMu.Unlock()
	defer func() {
		u.runningMu.Lock()
		delete(u.runningCancel, record.ID)
		u.runningMu.Unlock()
	}()
	startedAt := time.Now()
	err = task.Execute(execCtx, record.Payload)
	duration := time.Since(startedAt)
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
	return u.executionRepo.MarkFinished(context.WithoutCancel(ctx), &repo.DelayedTaskExecutionRecordMarkFinishedReq{
		ID:         record.ID,
		WorkerID:   record.WorkerID,
		Attempt:    record.Attempt,
		Status:     statusValue,
		FinishedAt: time.Now(),
		Duration:   duration,
		LastError:  lastError,
	})
}

type DelayedTaskExecutionRecordPageReq struct {
	Page           *common.PageReq
	DelayedTaskID  *int64
	IdempotencyKey *string
	Status         *schedulerenum.TaskExecutionStatus
	TriggerType    *schedulerenum.TaskTriggerType
}

type DelayedTaskExecutionRecordPageResp struct {
	Rows []*model.DelayedTaskExecutionRecord
	Page *common.PageResp
}

func (u *DelayedTaskUsecase) PageExecutionRecords(ctx context.Context, req *DelayedTaskExecutionRecordPageReq) (*DelayedTaskExecutionRecordPageResp, error) {
	if req == nil {
		req = &DelayedTaskExecutionRecordPageReq{}
	}
	resp, err := u.executionRepo.Page(ctx, &repo.DelayedTaskExecutionRecordPageReq{
		Page: req.Page,
		DelayedTaskExecutionRecordGetReq: repo.DelayedTaskExecutionRecordGetReq{
			DelayedTaskID:  req.DelayedTaskID,
			IdempotencyKey: req.IdempotencyKey,
			Status:         req.Status,
			TriggerType:    req.TriggerType,
		},
	})
	if err != nil {
		return nil, err
	}
	return &DelayedTaskExecutionRecordPageResp{Rows: resp.Rows, Page: resp.Page}, nil
}

func (u *DelayedTaskUsecase) EnsureSchedule(ctx context.Context) error {
	return u.scheduleRepo.Ensure(ctx)
}

func (u *DelayedTaskUsecase) StartConsuming(ctx context.Context) error {
	return u.scheduleRepo.Consume(ctx, u.HandleDelayedTaskMessage)
}

func (u *DelayedTaskUsecase) StopConsuming(ctx context.Context) error {
	return u.scheduleRepo.Stop(ctx)
}

func (u *DelayedTaskUsecase) delayedTaskScheduleSubject(executionRecordID int64) string {
	prefix := "scheduler.schedule.delayed_task_execution"
	if u.conf.GetScheduler() != nil && u.conf.GetScheduler().GetDelayedTaskScheduleSubjectPrefix() != "" {
		prefix = u.conf.GetScheduler().GetDelayedTaskScheduleSubjectPrefix()
	}
	return fmt.Sprintf("%s.%d", prefix, executionRecordID)
}

func (u *DelayedTaskUsecase) delayedTaskExecuteSubject() string {
	if u.conf.GetScheduler() != nil && u.conf.GetScheduler().GetDelayedTaskExecuteSubject() != "" {
		return u.conf.GetScheduler().GetDelayedTaskExecuteSubject()
	}
	return "scheduler.execute.delayed_task"
}

func (u *DelayedTaskUsecase) retryBackoff(attempt int32) time.Duration {
	base := 5 * time.Second
	capValue := 5 * time.Minute
	factor := 1 << max(0, int(attempt)-1)
	delay := time.Duration(factor) * base
	if delay > capValue {
		delay = capValue
	}
	return delay + time.Duration(rand.Int64N(int64(delay/2)+1))
}
