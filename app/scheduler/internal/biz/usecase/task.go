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

	"github.com/robfig/cron/v3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	otelcodes "go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskUsecase struct {
	logger                  *slog.Logger
	conf                    *config.Bootstrap
	cronParser              cron.Parser
	tx                      base.Tx
	taskRepo                repo.TaskRepo
	taskVersionRepo         repo.TaskVersionRepo
	taskExecutionRecordRepo repo.TaskExecutionRecordRepo
	taskLockRepo            repo.TaskLockRepo
	tasks                   map[string]taskimpl.Task
	taskEventBus            repo.TaskEventBus
	alert                   repo.TaskAlert
	runningCancels          sync.Map
	runningWG               sync.WaitGroup
}

func NewTaskUsecase(logger *slog.Logger, conf *config.Bootstrap, tx base.Tx, taskRepo repo.TaskRepo, taskVersionRepo repo.TaskVersionRepo, executionRepo repo.TaskExecutionRecordRepo, taskLockRepo repo.TaskLockRepo, tasks map[string]taskimpl.Task, taskEventBus repo.TaskEventBus, alert repo.TaskAlert) *TaskUsecase {
	return &TaskUsecase{
		logger:                  logger,
		conf:                    conf,
		cronParser:              cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow),
		tx:                      tx,
		taskRepo:                taskRepo,
		taskVersionRepo:         taskVersionRepo,
		taskExecutionRecordRepo: executionRepo,
		taskLockRepo:            taskLockRepo,
		tasks:                   tasks,
		taskEventBus:            taskEventBus,
		alert:                   alert,
	}
}

type TaskUpsertReq struct {
	Row *model.Task
}

type TaskUpsertResponse struct {
	Row *model.Task
}

func (u *TaskUsecase) Upsert(ctx context.Context, req *TaskUpsertReq) (*TaskUpsertResponse, error) {
	if req == nil || req.Row == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := u.upsert(ctx, &taskUpsertReq{Row: req.Row})
	if err != nil {
		return nil, err
	}
	return &TaskUpsertResponse{Row: resp.Row}, nil
}

type taskUpsertReq struct {
	Row *model.Task
}

type taskUpsertResponse struct {
	Row *model.Task
}

func (u *TaskUsecase) upsert(ctx context.Context, req *taskUpsertReq) (*taskUpsertResponse, error) {
	row := req.Row
	if strings.TrimSpace(row.Name) == "" || strings.TrimSpace(row.Title) == "" || strings.TrimSpace(row.CronSpec) == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if row.TimeoutSeconds <= 0 {
		if u.conf.GetScheduler() == nil || u.conf.GetScheduler().GetTaskTimeout() == nil || u.conf.GetScheduler().GetTaskTimeout().AsDuration() <= 0 {
			return nil, fmt.Errorf("scheduler task_timeout is invalid")
		}
		seconds := u.conf.GetScheduler().GetTaskTimeout().AsDuration().Seconds()
		if seconds <= 0 || seconds > float64(1<<31-1) {
			return nil, fmt.Errorf("scheduler task_timeout is invalid")
		}
		row.TimeoutSeconds = int32(seconds)
	}
	if row.Payload == "" {
		row.Payload = "{}"
	}
	if !json.Valid([]byte(row.Payload)) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if _, err := u.cronParser.Parse(row.CronSpec); err != nil {
		return nil, err
	}
	var saved *model.Task
	err := u.tx(ctx, func(txCtx context.Context) error {
		var err error
		savedResp, err := u.taskRepo.Upsert(txCtx, &repo.TaskUpsertReq{Row: row})
		if err != nil {
			return err
		}
		saved = savedResp.Row
		_, err = u.taskVersionRepo.Create(txCtx, &repo.TaskVersionCreateReq{Task: saved})
		return err
	})
	if err != nil {
		return nil, err
	}
	_, _ = u.taskEventBus.PublishTaskChanged(ctx, &repo.PublishTaskChangedReq{Message: &repo.TaskChangedMessage{TaskID: saved.ID, Version: saved.Version}})
	return &taskUpsertResponse{Row: saved}, nil
}

type TaskGetReq struct {
	ID int64
}

type TaskGetResponse struct {
	Row *model.Task
}

func (u *TaskUsecase) Get(ctx context.Context, req *TaskGetReq) (*TaskGetResponse, error) {
	if req == nil {
		req = &TaskGetReq{}
	}
	taskResp, err := u.taskRepo.Get(ctx, &repo.TaskGetReq{ID: &req.ID})
	if err != nil {
		return nil, err
	}
	return &TaskGetResponse{Row: taskResp.Row}, nil
}

type TaskPageReq struct {
	Page    *common.PageRequest
	IDs     []int64
	Name    *string
	Title   *string
	Enabled *bool
}

type TaskPageResponse struct {
	Rows []*model.Task
	Page *common.PageResponse
}

func (u *TaskUsecase) Page(ctx context.Context, req *TaskPageReq) (*TaskPageResponse, error) {
	if req == nil {
		req = &TaskPageReq{}
	}
	pageResp, err := u.taskRepo.Page(ctx, &repo.TaskPageReq{
		Page: req.Page,
		TaskGetReq: repo.TaskGetReq{
			IDs:     req.IDs,
			Name:    req.Name,
			Title:   req.Title,
			Enabled: req.Enabled,
		},
	})
	if err != nil {
		return nil, err
	}
	return &TaskPageResponse{Rows: pageResp.Rows, Page: pageResp.Page}, nil
}

type TaskExecutionRecordPageReq struct {
	Page        *common.PageRequest
	IDs         []int64
	TaskID      *int64
	Status      *schedulerenum.TaskExecutionStatus
	TriggerType *schedulerenum.TaskTriggerType
}

type TaskExecutionRecordPageResponse struct {
	Rows []*model.TaskExecutionRecord
	Page *common.PageResponse
}

func (u *TaskUsecase) PageExecutionRecords(ctx context.Context, req *TaskExecutionRecordPageReq) (*TaskExecutionRecordPageResponse, error) {
	if req == nil {
		req = &TaskExecutionRecordPageReq{}
	}
	pageResp, err := u.taskExecutionRecordRepo.Page(ctx, &repo.TaskExecutionRecordPageReq{
		Page: req.Page,
		TaskExecutionRecordGetReq: repo.TaskExecutionRecordGetReq{
			IDs:         req.IDs,
			TaskID:      req.TaskID,
			Status:      req.Status,
			TriggerType: req.TriggerType,
		},
	})
	if err != nil {
		return nil, err
	}
	return &TaskExecutionRecordPageResponse{Rows: pageResp.Rows, Page: pageResp.Page}, nil
}

type TaskListAvailableTasksReq struct {
	Keyword string
}

type TaskListAvailableTasksResponse struct {
	Rows []*model.AvailableTask
}

func (u *TaskUsecase) ListAvailableTasks(ctx context.Context, req *TaskListAvailableTasksReq) (*TaskListAvailableTasksResponse, error) {
	if req == nil {
		req = &TaskListAvailableTasksReq{}
	}
	resp, err := u.listAvailableTasks(ctx, &taskListAvailableTasksReq{Keyword: req.Keyword})
	if err != nil {
		return nil, err
	}
	return &TaskListAvailableTasksResponse{Rows: resp.Rows}, nil
}

type taskListAvailableTasksReq struct {
	Keyword string
}

type taskListAvailableTasksResponse struct {
	Rows []*model.AvailableTask
}

func (u *TaskUsecase) listAvailableTasks(ctx context.Context, req *taskListAvailableTasksReq) (*taskListAvailableTasksResponse, error) {
	_ = ctx
	keyword := strings.ToLower(strings.TrimSpace(req.Keyword))
	rows := make([]*model.AvailableTask, 0, len(u.tasks))
	for _, item := range u.tasks {
		row := &model.AvailableTask{
			Name:        item.Name(),
			Title:       item.Title(),
			Description: item.Description(),
		}
		if keyword == "" ||
			strings.Contains(strings.ToLower(row.Name), keyword) ||
			strings.Contains(strings.ToLower(row.Title), keyword) ||
			strings.Contains(strings.ToLower(row.Description), keyword) {
			rows = append(rows, row)
		}
	}
	sort.Slice(rows, func(i, j int) bool {
		return rows[i].Name < rows[j].Name
	})
	return &taskListAvailableTasksResponse{Rows: rows}, nil
}

type TaskTriggerReq struct {
	ID      int64
	Payload string
}

type TaskTriggerResponse struct {
	Row *model.TaskExecutionRecord
}

func (u *TaskUsecase) Trigger(ctx context.Context, req *TaskTriggerReq) (*TaskTriggerResponse, error) {
	if req == nil {
		req = &TaskTriggerReq{}
	}
	resp, err := u.trigger(ctx, &taskTriggerReq{ID: req.ID, Payload: req.Payload})
	if err != nil {
		return nil, err
	}
	return &TaskTriggerResponse{Row: resp.Row}, nil
}

type taskTriggerReq struct {
	ID      int64
	Payload string
}

type taskTriggerResponse struct {
	Row *model.TaskExecutionRecord
}

func (u *TaskUsecase) trigger(ctx context.Context, req *taskTriggerReq) (*taskTriggerResponse, error) {
	id := req.ID
	if id == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	taskResp, err := u.taskRepo.Get(ctx, &repo.TaskGetReq{ID: &id})
	if err != nil {
		return nil, err
	}
	task := taskResp.Row
	if req.Payload != "" {
		if !json.Valid([]byte(req.Payload)) {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		task.Payload = req.Payload
	}
	scheduleResp, err := u.scheduleExecution(ctx, &taskScheduleExecutionReq{Task: task, ScheduledAt: time.Now(), TriggerType: schedulerenum.TaskTriggerTypeManual})
	if err == nil && (scheduleResp == nil || scheduleResp.Record == nil) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
	}
	if err != nil {
		return nil, err
	}
	return &taskTriggerResponse{Row: scheduleResp.Record}, nil
}

type TaskCancelExecutionReq struct {
	ID int64
}

type TaskCancelExecutionResponse struct {
	Row *model.TaskExecutionRecord
}

func (u *TaskUsecase) CancelExecution(ctx context.Context, req *TaskCancelExecutionReq) (*TaskCancelExecutionResponse, error) {
	if req == nil {
		req = &TaskCancelExecutionReq{}
	}
	resp, err := u.cancelExecution(ctx, &taskCancelExecutionReq{ID: req.ID})
	if err != nil {
		return nil, err
	}
	return &TaskCancelExecutionResponse{Row: resp.Row}, nil
}

type taskCancelExecutionReq struct {
	ID int64
}

type taskCancelExecutionResponse struct {
	Row *model.TaskExecutionRecord
}

func (u *TaskUsecase) cancelExecution(ctx context.Context, req *taskCancelExecutionReq) (*taskCancelExecutionResponse, error) {
	id := req.ID
	if id == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	recordResp, err := u.taskExecutionRecordRepo.Get(ctx, &repo.TaskExecutionRecordGetReq{ID: &id})
	if err != nil {
		return nil, err
	}
	record := recordResp.Row
	if record.Status != schedulerenum.TaskExecutionStatusRunning {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if _, err := u.taskEventBus.PublishExecutionCanceled(ctx, &repo.PublishExecutionCanceledReq{Message: &repo.TaskExecutionCanceledMessage{ExecutionRecordID: id}}); err != nil {
		return nil, err
	}
	return &taskCancelExecutionResponse{Row: record}, nil
}

type TaskCancelExecutionLocallyReq struct {
	ID int64
}

func (u *TaskUsecase) CancelExecutionLocally(ctx context.Context, req *TaskCancelExecutionLocallyReq) error {
	if req == nil {
		return nil
	}
	return u.cancelExecutionLocally(ctx, &taskCancelExecutionLocallyReq{ID: req.ID})
}

type taskCancelExecutionLocallyReq struct {
	ID int64
}

func (u *TaskUsecase) cancelExecutionLocally(ctx context.Context, req *taskCancelExecutionLocallyReq) error {
	_ = ctx
	if cancelValue, ok := u.runningCancels.Load(req.ID); ok {
		cancelValue.(context.CancelFunc)()
	}
	return nil
}

type TaskScheduleExecutionReq struct {
	Task        *model.Task
	ScheduledAt time.Time
	TriggerType schedulerenum.TaskTriggerType
}

type TaskScheduleExecutionResponse struct {
	Record *model.TaskExecutionRecord
}

func (u *TaskUsecase) ScheduleExecution(ctx context.Context, req *TaskScheduleExecutionReq) (*TaskScheduleExecutionResponse, error) {
	if req == nil || req.Task == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := u.scheduleExecution(ctx, &taskScheduleExecutionReq{Task: req.Task, ScheduledAt: req.ScheduledAt, TriggerType: req.TriggerType})
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return &TaskScheduleExecutionResponse{}, nil
	}
	return &TaskScheduleExecutionResponse{Record: resp.Record}, nil
}

type taskScheduleExecutionReq struct {
	Task        *model.Task
	ScheduledAt time.Time
	TriggerType schedulerenum.TaskTriggerType
}

type taskScheduleExecutionResponse struct {
	Record *model.TaskExecutionRecord
}

func (u *TaskUsecase) scheduleExecution(ctx context.Context, req *taskScheduleExecutionReq) (resp *taskScheduleExecutionResponse, err error) {
	task := req.Task
	scheduledAt := req.ScheduledAt
	triggerType := req.TriggerType
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	var record *model.TaskExecutionRecord
	finish := func(row *model.TaskExecutionRecord, returnErr error) (*taskScheduleExecutionResponse, error) {
		record = row
		if returnErr != nil {
			return nil, returnErr
		}
		if row == nil {
			return nil, nil
		}
		return &taskScheduleExecutionResponse{Record: row}, nil
	}
	scheduleCtx, span := otel.Tracer("scheduler.task").Start(
		ctx,
		"scheduler.schedule "+task.Name,
		oteltrace.WithSpanKind(oteltrace.SpanKindInternal),
		oteltrace.WithAttributes(
			attribute.Int64("scheduler.task.id", task.ID),
			attribute.String("scheduler.task.name", task.Name),
			attribute.String("scheduler.task.title", task.Title),
			attribute.String("scheduler.trigger_type", string(triggerType)),
			attribute.String("scheduler.scheduled_at", scheduledAt.Format(time.RFC3339Nano)),
		),
	)
	defer span.End()
	ctx = scheduleCtx
	skipReason := "unknown"
	defer func() {
		if err != nil {
			span.SetStatus(otelcodes.Error, err.Error())
			u.logger.ErrorContext(ctx, "scheduler task scheduling failed", constant.LogFieldTaskID, task.ID, constant.LogFieldTaskName, task.Name, constant.LogFieldErr, err, "scheduled_at", scheduledAt.Format(time.RFC3339Nano), "trigger_type", triggerType)
			return
		}
		if record == nil {
			u.logger.DebugContext(ctx, "scheduler task skipped", constant.LogFieldTaskID, task.ID, constant.LogFieldTaskName, task.Name, "scheduled_at", scheduledAt.Format(time.RFC3339Nano), "trigger_type", triggerType, "skip_reason", skipReason)
			return
		}
		u.logger.DebugContext(ctx, "scheduler execution accepted", constant.LogFieldTaskID, task.ID, constant.LogFieldTaskName, task.Name, constant.LogFieldExecutionID, record.ID, constant.LogFieldStatus, record.Status, "scheduled_at", scheduledAt.Format(time.RFC3339Nano), "trigger_type", triggerType)
	}()
	u.logger.DebugContext(ctx, "scheduler task triggered", constant.LogFieldTaskID, task.ID, constant.LogFieldTaskName, task.Name, "scheduled_at", scheduledAt.Format(time.RFC3339Nano), "trigger_type", triggerType)
	schedule, parseErr := u.cronParser.Parse(task.CronSpec)
	if parseErr != nil {
		return nil, parseErr
	}
	next := schedule.Next(scheduledAt)
	nextAfter := schedule.Next(next)
	if next.IsZero() || nextAfter.IsZero() || !nextAfter.After(next) {
		return nil, fmt.Errorf("scheduler cron period is invalid")
	}
	period := nextAfter.Sub(next)
	acquired, err := u.taskLockRepo.TryAcquireSchedule(ctx, &repo.TaskScheduleAcquireReq{
		TaskID:            task.ID,
		ScheduledAt:       scheduledAt,
		AllowOverlap:      task.AllowOverlap,
		SchedulePeriodTTL: period,
		RunningLockTTL:    period,
	})
	if err != nil {
		u.logger.ErrorContext(ctx, "acquire scheduler redis task lock failed", constant.LogFieldTaskID, task.ID, constant.LogFieldTaskName, task.Name, constant.LogFieldErr, err)
		if task.AllowOverlap {
			startResp, err := u.startExecution(ctx, &taskStartExecutionReq{Task: task, ScheduledAt: scheduledAt, TriggerType: triggerType})
			if err != nil {
				if startResp != nil {
					return finish(startResp.Record, err)
				}
				return nil, err
			}
			if startResp == nil || !startResp.Created || startResp.Conflict {
				if startResp != nil && startResp.Conflict {
					skipReason = "db_unique_conflict"
				} else {
					skipReason = "db_create_skipped"
				}
				return nil, nil
			}
			return finish(startResp.Record, nil)
		}
		var result *model.TaskExecutionRecord
		var created bool
		var conflict bool
		err = u.tx(ctx, func(txCtx context.Context) error {
			if _, err := u.taskRepo.Lock(txCtx, &repo.TaskLockReq{ID: task.ID}); err != nil {
				return err
			}
			existsResp, err := u.taskExecutionRecordRepo.ExistsPeriod(txCtx, &repo.TaskExecutionRecordExistsPeriodReq{TaskID: task.ID, ScheduledAt: scheduledAt})
			if err != nil {
				return err
			}
			if existsResp.Exists {
				conflict = true
				skipReason = "db_period_exists"
				return nil
			}
			timeout := time.Duration(task.TimeoutSeconds) * time.Second
			if timeout <= 0 {
				if u.conf.GetScheduler() == nil || u.conf.GetScheduler().GetTaskTimeout() == nil || u.conf.GetScheduler().GetTaskTimeout().AsDuration() <= 0 {
					return fmt.Errorf("scheduler task_timeout is invalid")
				}
				timeout = u.conf.GetScheduler().GetTaskTimeout().AsDuration()
			}
			runningResp, err := u.taskExecutionRecordRepo.HasUnexpiredRunning(txCtx, &repo.TaskExecutionRecordHasUnexpiredRunningReq{TaskID: task.ID, StartedAfter: time.Now().Add(-(timeout + period))})
			if err != nil {
				return err
			}
			if runningResp.Exists {
				traceID := ""
				if oteltrace.SpanContextFromContext(txCtx).IsValid() {
					traceID = oteltrace.SpanContextFromContext(txCtx).TraceID().String()
				}
				createResp, err := u.createExecutionRecord(txCtx, &taskCreateExecutionRecordReq{
					Record: &model.TaskExecutionRecord{
						TaskID:      task.ID,
						ScheduledAt: scheduledAt,
						FinishedAt:  new(time.Now()),
						DurationMS:  new(int64(0)),
						Status:      schedulerenum.TaskExecutionStatusOverlapSkipped,
						TriggerType: triggerType,
						TaskVersion: task.Version,
						Payload:     task.Payload,
						TraceID:     traceID,
						LastError:   "scheduler task has another running execution",
					},
					Status: schedulerenum.TaskExecutionStatusOverlapSkipped,
				})
				if createResp != nil {
					result = createResp.Record
					created = createResp.Created
					conflict = createResp.Conflict
					if createResp.Conflict {
						skipReason = "db_unique_conflict"
					}
				}
				return err
			}
			traceID := ""
			if oteltrace.SpanContextFromContext(txCtx).IsValid() {
				traceID = oteltrace.SpanContextFromContext(txCtx).TraceID().String()
			}
			createResp, err := u.createExecutionRecord(txCtx, &taskCreateExecutionRecordReq{
				Record: &model.TaskExecutionRecord{
					TaskID:      task.ID,
					ScheduledAt: scheduledAt,
					StartedAt:   new(time.Now()),
					Status:      schedulerenum.TaskExecutionStatusRunning,
					TriggerType: triggerType,
					TaskVersion: task.Version,
					Payload:     task.Payload,
					TraceID:     traceID,
				},
				Status: schedulerenum.TaskExecutionStatusRunning,
			})
			if createResp != nil {
				result = createResp.Record
				created = createResp.Created
				conflict = createResp.Conflict
				if createResp.Conflict {
					skipReason = "db_unique_conflict"
				}
			}
			return err
		})
		if err != nil {
			return finish(result, err)
		}
		if conflict {
			u.logger.WarnContext(ctx, "scheduler execution record already exists", constant.LogFieldTaskID, task.ID, constant.LogFieldTaskName, task.Name)
			if skipReason == "unknown" {
				skipReason = "db_unique_conflict"
			}
			return nil, nil
		}
		if result == nil || !created {
			skipReason = "db_create_skipped"
			return nil, nil
		}
		if result.Status == schedulerenum.TaskExecutionStatusOverlapSkipped {
			if task.AlertEnabled {
				_, _ = u.alert.Alert(ctx, &repo.TaskAlertReq{Task: task, Record: result, Reason: "overlap"})
			}
			return finish(result, nil)
		}
		startCreatedResp, err := u.startCreatedExecution(ctx, &taskStartCreatedExecutionReq{Task: task, Record: result, ScheduledAt: scheduledAt, TriggerType: triggerType})
		if startCreatedResp != nil {
			result = startCreatedResp.Record
		}
		return finish(result, err)
	}
	if acquired == nil {
		return nil, fmt.Errorf("scheduler task lock acquire result is nil")
	}
	switch acquired.Decision {
	case schedulerenum.TaskScheduleDecisionSkip:
		skipReason = "redis_schedule_skip"
		return nil, nil
	case schedulerenum.TaskScheduleDecisionRun:
		startResp, err := u.startExecution(ctx, &taskStartExecutionReq{Task: task, ScheduledAt: scheduledAt, TriggerType: triggerType, RunningToken: acquired.RunningToken})
		if err != nil {
			if startResp != nil {
				return finish(startResp.Record, err)
			}
			return nil, err
		}
		if startResp == nil || !startResp.Created || startResp.Conflict {
			if startResp != nil && startResp.Conflict {
				skipReason = "db_unique_conflict"
			} else {
				skipReason = "db_create_skipped"
			}
			return nil, nil
		}
		return finish(startResp.Record, nil)
	case schedulerenum.TaskScheduleDecisionOverlap:
		overlapResp, err := u.createOverlapSkipped(ctx, &taskCreateOverlapSkippedReq{Task: task, ScheduledAt: scheduledAt, TriggerType: triggerType})
		if err != nil {
			if overlapResp != nil {
				return finish(overlapResp.Record, err)
			}
			return nil, err
		}
		if overlapResp == nil || !overlapResp.Created || overlapResp.Conflict {
			if overlapResp != nil && overlapResp.Conflict {
				skipReason = "db_unique_conflict"
			} else {
				skipReason = "db_create_skipped"
			}
			return nil, nil
		}
		return finish(overlapResp.Record, nil)
	default:
		return nil, fmt.Errorf("unknown scheduler schedule decision: %s", acquired.Decision)
	}
}

type TaskCreateOverlapSkippedReq struct {
	Task        *model.Task
	ScheduledAt time.Time
	TriggerType schedulerenum.TaskTriggerType
}

type TaskCreateOverlapSkippedResponse struct {
	Record   *model.TaskExecutionRecord
	Created  bool
	Conflict bool
}

func (u *TaskUsecase) CreateOverlapSkipped(ctx context.Context, req *TaskCreateOverlapSkippedReq) (*TaskCreateOverlapSkippedResponse, error) {
	if req == nil || req.Task == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := u.createOverlapSkipped(ctx, &taskCreateOverlapSkippedReq{Task: req.Task, ScheduledAt: req.ScheduledAt, TriggerType: req.TriggerType})
	if err != nil {
		return nil, err
	}
	return &TaskCreateOverlapSkippedResponse{Record: resp.Record, Created: resp.Created, Conflict: resp.Conflict}, nil
}

type taskCreateOverlapSkippedReq struct {
	Task        *model.Task
	ScheduledAt time.Time
	TriggerType schedulerenum.TaskTriggerType
}

type taskCreateOverlapSkippedResponse struct {
	Record   *model.TaskExecutionRecord
	Created  bool
	Conflict bool
}

func (u *TaskUsecase) createOverlapSkipped(ctx context.Context, req *taskCreateOverlapSkippedReq) (*taskCreateOverlapSkippedResponse, error) {
	task := req.Task
	traceID := ""
	if oteltrace.SpanContextFromContext(ctx).IsValid() {
		traceID = oteltrace.SpanContextFromContext(ctx).TraceID().String()
	}
	createResp, err := u.createExecutionRecord(ctx, &taskCreateExecutionRecordReq{
		Record: &model.TaskExecutionRecord{
			TaskID:      task.ID,
			ScheduledAt: req.ScheduledAt,
			FinishedAt:  new(time.Now()),
			DurationMS:  new(int64(0)),
			Status:      schedulerenum.TaskExecutionStatusOverlapSkipped,
			TriggerType: req.TriggerType,
			TaskVersion: task.Version,
			Payload:     task.Payload,
			TraceID:     traceID,
			LastError:   "scheduler task has another running execution",
		},
		Status: schedulerenum.TaskExecutionStatusOverlapSkipped,
	})
	if createResp != nil && createResp.Conflict {
		u.logger.WarnContext(ctx, "scheduler execution record already exists", constant.LogFieldTaskID, task.ID, constant.LogFieldTaskName, task.Name)
	}
	if err == nil && createResp != nil && createResp.Created && task.AlertEnabled {
		_, _ = u.alert.Alert(ctx, &repo.TaskAlertReq{Task: task, Record: createResp.Record, Reason: "overlap"})
	}
	if createResp == nil {
		return nil, err
	}
	return &taskCreateOverlapSkippedResponse{Record: createResp.Record, Created: createResp.Created, Conflict: createResp.Conflict}, err
}

type TaskStartExecutionReq struct {
	Task         *model.Task
	ScheduledAt  time.Time
	TriggerType  schedulerenum.TaskTriggerType
	RunningToken string
}

type TaskStartExecutionResponse struct {
	Record   *model.TaskExecutionRecord
	Created  bool
	Conflict bool
}

func (u *TaskUsecase) StartExecution(ctx context.Context, req *TaskStartExecutionReq) (*TaskStartExecutionResponse, error) {
	if req == nil || req.Task == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := u.startExecution(ctx, &taskStartExecutionReq{Task: req.Task, ScheduledAt: req.ScheduledAt, TriggerType: req.TriggerType, RunningToken: req.RunningToken})
	if err != nil {
		return nil, err
	}
	return &TaskStartExecutionResponse{Record: resp.Record, Created: resp.Created, Conflict: resp.Conflict}, nil
}

type taskStartExecutionReq struct {
	Task         *model.Task
	ScheduledAt  time.Time
	TriggerType  schedulerenum.TaskTriggerType
	RunningToken string
}

type taskStartExecutionResponse struct {
	Record   *model.TaskExecutionRecord
	Created  bool
	Conflict bool
}

func (u *TaskUsecase) startExecution(ctx context.Context, req *taskStartExecutionReq) (*taskStartExecutionResponse, error) {
	task := req.Task
	traceID := ""
	if oteltrace.SpanContextFromContext(ctx).IsValid() {
		traceID = oteltrace.SpanContextFromContext(ctx).TraceID().String()
	}
	createResp, err := u.createExecutionRecord(ctx, &taskCreateExecutionRecordReq{
		Record: &model.TaskExecutionRecord{
			TaskID:      task.ID,
			ScheduledAt: req.ScheduledAt,
			StartedAt:   new(time.Now()),
			Status:      schedulerenum.TaskExecutionStatusRunning,
			TriggerType: req.TriggerType,
			TaskVersion: task.Version,
			Payload:     task.Payload,
			TraceID:     traceID,
		},
		Status: schedulerenum.TaskExecutionStatusRunning,
	})
	if createResp != nil && createResp.Conflict {
		u.logger.WarnContext(ctx, "scheduler execution record already exists", constant.LogFieldTaskID, task.ID, constant.LogFieldTaskName, task.Name)
	}
	if err != nil || createResp == nil || !createResp.Created {
		if releaseErr := u.releaseRunning(context.WithoutCancel(ctx), &taskReleaseRunningReq{TaskID: task.ID, RunningToken: req.RunningToken, Exclusive: !task.AllowOverlap && req.RunningToken != ""}); releaseErr != nil {
			u.logger.ErrorContext(context.WithoutCancel(ctx), "release scheduler running lock failed", constant.LogFieldTaskID, task.ID, constant.LogFieldErr, releaseErr)
		}
		if createResp == nil {
			return nil, err
		}
		return &taskStartExecutionResponse{Record: createResp.Record, Created: createResp.Created, Conflict: createResp.Conflict}, err
	}
	record := createResp.Record
	runningTTL := time.Duration(0)
	if req.RunningToken != "" {
		schedule, parseErr := u.cronParser.Parse(task.CronSpec)
		if parseErr == nil {
			next := schedule.Next(req.ScheduledAt)
			nextAfter := schedule.Next(next)
			if next.IsZero() || nextAfter.IsZero() || !nextAfter.After(next) {
				parseErr = fmt.Errorf("scheduler cron period is invalid")
			} else {
				runningTTL = nextAfter.Sub(next)
			}
		}
		if parseErr != nil {
			finishedAt := time.Now()
			durationMS := int64(0)
			if record.StartedAt != nil {
				durationMS = finishedAt.Sub(*record.StartedAt).Milliseconds()
				if durationMS < 0 {
					durationMS = 0
				}
			}
			failedResp, markErr := u.markExecutionFinished(context.WithoutCancel(ctx), &taskMarkExecutionFinishedReq{ID: record.ID, Status: schedulerenum.TaskExecutionStatusFailed, FinishedAt: finishedAt, DurationMS: durationMS, LastError: parseErr.Error()})
			if releaseErr := u.releaseRunning(context.WithoutCancel(ctx), &taskReleaseRunningReq{TaskID: task.ID, ExecutionRecordID: record.ID, RunningToken: req.RunningToken, Exclusive: !task.AllowOverlap}); releaseErr != nil {
				u.logger.ErrorContext(context.WithoutCancel(ctx), "release scheduler running lock failed", constant.LogFieldTaskID, task.ID, constant.LogFieldErr, releaseErr)
			}
			if markErr != nil {
				return &taskStartExecutionResponse{Record: failedResp.Record, Created: true}, markErr
			}
			return &taskStartExecutionResponse{Record: failedResp.Record, Created: true}, parseErr
		}
		registerResp, registerErr := u.registerRunning(context.WithoutCancel(ctx), &taskRegisterRunningReq{TaskID: task.ID, ExecutionRecordID: record.ID, RunningToken: req.RunningToken, Exclusive: !task.AllowOverlap, TTL: runningTTL})
		if registerErr != nil || registerResp == nil || !registerResp.OK {
			finishedAt := time.Now()
			durationMS := int64(0)
			if record.StartedAt != nil {
				durationMS = finishedAt.Sub(*record.StartedAt).Milliseconds()
				if durationMS < 0 {
					durationMS = 0
				}
			}
			lastError := "scheduler redis running registration failed"
			if registerErr != nil {
				lastError = registerErr.Error()
			}
			failedResp, markErr := u.markExecutionFinished(context.WithoutCancel(ctx), &taskMarkExecutionFinishedReq{ID: record.ID, Status: schedulerenum.TaskExecutionStatusFailed, FinishedAt: finishedAt, DurationMS: durationMS, LastError: lastError})
			if releaseErr := u.releaseRunning(context.WithoutCancel(ctx), &taskReleaseRunningReq{TaskID: task.ID, ExecutionRecordID: record.ID, RunningToken: req.RunningToken, Exclusive: !task.AllowOverlap}); releaseErr != nil {
				u.logger.ErrorContext(context.WithoutCancel(ctx), "release scheduler running lock failed", constant.LogFieldTaskID, task.ID, constant.LogFieldErr, releaseErr)
			}
			if markErr != nil {
				return &taskStartExecutionResponse{Record: failedResp.Record, Created: true}, markErr
			}
			if registerErr != nil {
				return &taskStartExecutionResponse{Record: failedResp.Record, Created: true}, registerErr
			}
			return &taskStartExecutionResponse{Record: failedResp.Record, Created: true}, fmt.Errorf("scheduler redis running registration failed")
		}
	}
	startCreatedResp, err := u.startCreatedExecution(ctx, &taskStartCreatedExecutionReq{Task: task, Record: record, ScheduledAt: req.ScheduledAt, TriggerType: req.TriggerType, RunningToken: req.RunningToken, RunningTTL: runningTTL})
	if startCreatedResp == nil {
		return nil, err
	}
	return &taskStartExecutionResponse{Record: startCreatedResp.Record, Created: startCreatedResp.Created, Conflict: startCreatedResp.Conflict}, err
}

type taskStartCreatedExecutionReq struct {
	Task         *model.Task
	Record       *model.TaskExecutionRecord
	ScheduledAt  time.Time
	TriggerType  schedulerenum.TaskTriggerType
	RunningToken string
	RunningTTL   time.Duration
}

type taskStartCreatedExecutionResponse struct {
	Record   *model.TaskExecutionRecord
	Created  bool
	Conflict bool
}

func (u *TaskUsecase) startCreatedExecution(ctx context.Context, req *taskStartCreatedExecutionReq) (*taskStartCreatedExecutionResponse, error) {
	task := req.Task
	record := req.Record
	runCtx, span := otel.Tracer("scheduler.task").Start(
		ctx,
		"scheduler.task "+task.Name,
		oteltrace.WithSpanKind(oteltrace.SpanKindInternal),
		oteltrace.WithAttributes(
			attribute.Int64("scheduler.task.id", task.ID),
			attribute.String("scheduler.task.name", task.Name),
			attribute.String("scheduler.task.title", task.Title),
			attribute.String("scheduler.trigger_type", string(req.TriggerType)),
			attribute.String("scheduler.scheduled_at", req.ScheduledAt.Format(time.RFC3339Nano)),
		),
	)
	traceID := ""
	if span.SpanContext().IsValid() {
		traceID = span.SpanContext().TraceID().String()
	}
	if traceID != "" && record.TraceID == "" {
		record.TraceID = traceID
	}
	timeout := time.Duration(task.TimeoutSeconds) * time.Second
	if timeout <= 0 {
		if u.conf.GetScheduler() == nil || u.conf.GetScheduler().GetTaskTimeout() == nil || u.conf.GetScheduler().GetTaskTimeout().AsDuration() <= 0 {
			span.End()
			return nil, fmt.Errorf("scheduler task_timeout is invalid")
		}
		timeout = u.conf.GetScheduler().GetTaskTimeout().AsDuration()
	}
	span.SetAttributes(attribute.Int64("scheduler.execution.id", record.ID))
	callCtx, cancel := context.WithTimeout(context.WithoutCancel(runCtx), timeout)
	u.runningCancels.Store(record.ID, cancel)
	heartbeatInterval := time.Duration(0)
	if req.RunningToken != "" {
		var configErr error
		if req.RunningTTL <= 0 {
			configErr = fmt.Errorf("scheduler cron period is invalid")
		} else {
			heartbeatInterval = req.RunningTTL / 3
			if heartbeatInterval < 100*time.Millisecond {
				heartbeatInterval = 100 * time.Millisecond
			}
			if heartbeatInterval >= req.RunningTTL {
				configErr = fmt.Errorf("scheduler running heartbeat interval must be less than cron period")
			}
		}
		if configErr != nil {
			finishedAt := time.Now()
			durationMS := int64(0)
			if record.StartedAt != nil {
				durationMS = finishedAt.Sub(*record.StartedAt).Milliseconds()
				if durationMS < 0 {
					durationMS = 0
				}
			}
			failedResp, markErr := u.markExecutionFinished(context.WithoutCancel(runCtx), &taskMarkExecutionFinishedReq{ID: record.ID, Status: schedulerenum.TaskExecutionStatusFailed, FinishedAt: finishedAt, DurationMS: durationMS, LastError: configErr.Error()})
			span.End()
			cancel()
			u.runningCancels.Delete(record.ID)
			if releaseErr := u.releaseRunning(context.WithoutCancel(runCtx), &taskReleaseRunningReq{TaskID: task.ID, ExecutionRecordID: record.ID, RunningToken: req.RunningToken, Exclusive: !task.AllowOverlap}); releaseErr != nil {
				u.logger.ErrorContext(context.WithoutCancel(runCtx), "release scheduler running lock failed", constant.LogFieldTaskID, task.ID, constant.LogFieldErr, releaseErr)
			}
			failedRecord := record
			if failedResp != nil && failedResp.Record != nil {
				failedRecord = failedResp.Record
			}
			if markErr != nil {
				return &taskStartCreatedExecutionResponse{Record: failedRecord, Created: true}, markErr
			}
			return &taskStartCreatedExecutionResponse{Record: failedRecord, Created: true}, configErr
		}
	}
	u.runningWG.Add(1)
	go u.executeRecord(callCtx, &taskExecuteRecordReq{Cancel: cancel, Span: span, Task: task, Record: record, ScheduledAt: req.ScheduledAt, RunningToken: req.RunningToken, HeartbeatInterval: heartbeatInterval, RunningTTL: req.RunningTTL})
	return &taskStartCreatedExecutionResponse{Record: record, Created: true}, nil
}

type taskExecuteRecordReq struct {
	Cancel            context.CancelFunc
	Span              oteltrace.Span
	Task              *model.Task
	Record            *model.TaskExecutionRecord
	ScheduledAt       time.Time
	RunningToken      string
	HeartbeatInterval time.Duration
	RunningTTL        time.Duration
}

func (u *TaskUsecase) executeRecord(ctx context.Context, req *taskExecuteRecordReq) {
	task := req.Task
	record := req.Record
	span := req.Span
	defer span.End()
	defer req.Cancel()
	defer u.runningCancels.Delete(record.ID)
	defer u.runningWG.Done()
	if req.RunningToken != "" {
		go u.heartbeatRunningLock(ctx, &taskHeartbeatRunningLockReq{Cancel: req.Cancel, TaskID: task.ID, ExecutionRecordID: record.ID, RunningToken: req.RunningToken, Exclusive: !task.AllowOverlap, Interval: req.HeartbeatInterval, TTL: req.RunningTTL})
		defer func() {
			if err := u.releaseRunning(context.WithoutCancel(ctx), &taskReleaseRunningReq{TaskID: task.ID, ExecutionRecordID: record.ID, RunningToken: req.RunningToken, Exclusive: !task.AllowOverlap}); err != nil {
				u.logger.ErrorContext(context.WithoutCancel(ctx), "release scheduler running lock failed", constant.LogFieldTaskID, task.ID, constant.LogFieldErr, err)
			}
		}()
	}
	currentTask, ok := u.tasks[task.Name]
	var err error
	if ok {
		err = currentTask.Execute(ctx, task.Payload)
	} else {
		err = fmt.Errorf("unknown scheduler task: %s", task.Name)
	}
	finishedAt := time.Now()
	durationMS := int64(0)
	if record.StartedAt != nil {
		durationMS = finishedAt.Sub(*record.StartedAt).Milliseconds()
	}
	statusValue := schedulerenum.TaskExecutionStatusSuccess
	lastError := ""
	if errors.Is(ctx.Err(), context.Canceled) {
		lastError = ctx.Err().Error()
		span.SetStatus(otelcodes.Error, lastError)
		statusValue = schedulerenum.TaskExecutionStatusCanceled
	} else if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		lastError = ctx.Err().Error()
		span.SetStatus(otelcodes.Error, lastError)
		statusValue = schedulerenum.TaskExecutionStatusTimeout
	} else if err != nil {
		lastError = err.Error()
		span.SetStatus(otelcodes.Error, lastError)
		statusValue = schedulerenum.TaskExecutionStatusFailed
		if status.Code(err) == codes.DeadlineExceeded {
			statusValue = schedulerenum.TaskExecutionStatusTimeout
		}
	}
	markResp, markErr := u.markExecutionFinished(context.WithoutCancel(ctx), &taskMarkExecutionFinishedReq{ID: record.ID, Status: statusValue, FinishedAt: finishedAt, DurationMS: durationMS, LastError: lastError})
	if markErr != nil || markResp == nil || !markResp.Updated {
		return
	}
	finished := markResp.Record
	if task.AlertEnabled {
		if statusValue != schedulerenum.TaskExecutionStatusSuccess {
			_, _ = u.alert.Alert(ctx, &repo.TaskAlertReq{Task: task, Record: finished, Reason: statusValue.String()})
		}
		if statusValue == schedulerenum.TaskExecutionStatusSuccess {
			schedule, parseErr := u.cronParser.Parse(task.CronSpec)
			if parseErr == nil {
				next := schedule.Next(req.ScheduledAt)
				if record.StartedAt != nil && next.After(req.ScheduledAt) && finishedAt.Sub(*record.StartedAt) > next.Sub(req.ScheduledAt) {
					_, _ = u.alert.Alert(ctx, &repo.TaskAlertReq{Task: task, Record: finished, Reason: "duration_exceeded"})
				}
			}
		}
	}
}

type taskHeartbeatRunningLockReq struct {
	Cancel            context.CancelFunc
	TaskID            int64
	ExecutionRecordID int64
	RunningToken      string
	Exclusive         bool
	Interval          time.Duration
	TTL               time.Duration
}

func (u *TaskUsecase) heartbeatRunningLock(ctx context.Context, req *taskHeartbeatRunningLockReq) {
	ticker := time.NewTicker(req.Interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			refreshResp, err := u.refreshRunning(context.WithoutCancel(ctx), &taskRefreshRunningReq{TaskID: req.TaskID, ExecutionRecordID: req.ExecutionRecordID, RunningToken: req.RunningToken, Exclusive: req.Exclusive, TTL: req.TTL})
			if err == nil && refreshResp != nil && refreshResp.OK {
				continue
			}
			if err != nil {
				u.logger.ErrorContext(context.WithoutCancel(ctx), "refresh scheduler running lock failed", constant.LogFieldTaskID, req.TaskID, constant.LogFieldErr, err)
			}
			// Redis 心跳失败时主动取消执行，避免本地任务继续占用资源。
			req.Cancel()
			return
		}
	}
}

type TaskCheckExecutionRuntimesReq struct {
	IDs []int64
}

type TaskCheckExecutionRuntimesResponse struct {
	Rows []*model.TaskExecutionRuntime
}

func (u *TaskUsecase) CheckExecutionRuntimes(ctx context.Context, req *TaskCheckExecutionRuntimesReq) (*TaskCheckExecutionRuntimesResponse, error) {
	if req == nil {
		req = &TaskCheckExecutionRuntimesReq{}
	}
	resp, err := u.checkExecutionRuntimes(ctx, &taskCheckExecutionRuntimesReq{IDs: req.IDs})
	if err != nil {
		return nil, err
	}
	return &TaskCheckExecutionRuntimesResponse{Rows: resp.Rows}, nil
}

type taskCheckExecutionRuntimesReq struct {
	IDs []int64
}

type taskCheckExecutionRuntimesResponse struct {
	Rows []*model.TaskExecutionRuntime
}

func (u *TaskUsecase) checkExecutionRuntimes(ctx context.Context, req *taskCheckExecutionRuntimesReq) (*taskCheckExecutionRuntimesResponse, error) {
	if len(req.IDs) == 0 {
		return &taskCheckExecutionRuntimesResponse{}, nil
	}
	recordsResp, err := u.listExecutionRecords(ctx, &taskListExecutionRecordsReq{Query: &repo.TaskExecutionRecordGetReq{IDs: req.IDs}})
	if err != nil {
		return nil, err
	}
	records := recordsResp.Rows
	recordByID := make(map[int64]*model.TaskExecutionRecord, len(records))
	runningByTask := make(map[int64][]int64)
	for _, row := range records {
		recordByID[row.ID] = row
		if row.Status == schedulerenum.TaskExecutionStatusRunning {
			runningByTask[row.TaskID] = append(runningByTask[row.TaskID], row.ID)
		}
	}
	activeByID := make(map[int64]bool)
	redisFailedTask := make(map[int64]bool)
	for taskID, executionIDs := range runningByTask {
		mapResp, mapErr := u.mapRunning(ctx, &taskMapRunningReq{TaskID: taskID, ExecutionRecordIDs: executionIDs})
		if mapErr != nil {
			redisFailedTask[taskID] = true
			u.logger.ErrorContext(ctx, "map scheduler running lock failed", constant.LogFieldTaskID, taskID, constant.LogFieldErr, mapErr)
			continue
		}
		for id, active := range mapResp.Rows {
			activeByID[id] = active
		}
	}
	taskIDs := make([]int64, 0)
	versions := make([]int64, 0)
	type taskVersionKey struct {
		taskID  int64
		version int64
	}
	needVersion := make(map[taskVersionKey]struct{})
	for _, row := range records {
		if row.Status != schedulerenum.TaskExecutionStatusRunning || activeByID[row.ID] || redisFailedTask[row.TaskID] {
			continue
		}
		key := taskVersionKey{taskID: row.TaskID, version: row.TaskVersion}
		if _, ok := needVersion[key]; ok {
			continue
		}
		needVersion[key] = struct{}{}
		taskIDs = append(taskIDs, row.TaskID)
		versions = append(versions, row.TaskVersion)
	}
	versionByKey := make(map[taskVersionKey]*model.TaskVersion, len(needVersion))
	if len(needVersion) > 0 {
		versionResp, err := u.listTaskVersions(ctx, &taskListTaskVersionsReq{Query: &repo.TaskVersionGetReq{TaskIDs: taskIDs, Versions: versions}})
		if err != nil {
			return nil, err
		}
		for _, row := range versionResp.Rows {
			key := taskVersionKey{taskID: row.TaskID, version: row.Version}
			if _, ok := needVersion[key]; ok {
				versionByKey[key] = row
			}
		}
	}
	now := time.Now()
	result := make([]*model.TaskExecutionRuntime, 0, len(req.IDs))
	for _, id := range req.IDs {
		row := recordByID[id]
		if row == nil {
			continue
		}
		state := schedulerenum.TaskExecutionRuntimeStateTerminal
		if row.Status == schedulerenum.TaskExecutionStatusRunning {
			state = schedulerenum.TaskExecutionRuntimeStateUnknown
			if activeByID[row.ID] {
				state = schedulerenum.TaskExecutionRuntimeStateActive
			} else if !redisFailedTask[row.TaskID] && row.StartedAt != nil {
				version := versionByKey[taskVersionKey{taskID: row.TaskID, version: row.TaskVersion}]
				if version != nil {
					timeout := time.Duration(version.TimeoutSeconds) * time.Second
					if timeout <= 0 && u.conf.GetScheduler() != nil && u.conf.GetScheduler().GetTaskTimeout() != nil && u.conf.GetScheduler().GetTaskTimeout().AsDuration() > 0 {
						timeout = u.conf.GetScheduler().GetTaskTimeout().AsDuration()
					}
					schedule, parseErr := u.cronParser.Parse(version.CronSpec)
					if timeout > 0 && parseErr == nil {
						next := schedule.Next(row.ScheduledAt)
						nextAfter := schedule.Next(next)
						if !next.IsZero() && !nextAfter.IsZero() && nextAfter.After(next) && now.After(row.StartedAt.Add(timeout+nextAfter.Sub(next))) {
							state = schedulerenum.TaskExecutionRuntimeStateStale
						}
					}
				}
			}
		}
		result = append(result, &model.TaskExecutionRuntime{
			ExecutionRecordID: row.ID,
			TaskID:            row.TaskID,
			State:             state,
		})
	}
	return &taskCheckExecutionRuntimesResponse{Rows: result}, nil
}

type TaskMarkExecutionsUnknownReq struct {
	IDs []int64
}

type TaskMarkExecutionsUnknownResponse struct {
	Rows []*model.TaskExecutionRecord
}

func (u *TaskUsecase) MarkExecutionsUnknown(ctx context.Context, req *TaskMarkExecutionsUnknownReq) (*TaskMarkExecutionsUnknownResponse, error) {
	if req == nil {
		req = &TaskMarkExecutionsUnknownReq{}
	}
	resp, err := u.markExecutionsUnknown(ctx, &taskMarkExecutionsUnknownReq{IDs: req.IDs})
	if err != nil {
		return nil, err
	}
	return &TaskMarkExecutionsUnknownResponse{Rows: resp.Rows}, nil
}

type taskMarkExecutionsUnknownReq struct {
	IDs []int64
}

type taskMarkExecutionsUnknownResponse struct {
	Rows []*model.TaskExecutionRecord
}

func (u *TaskUsecase) markExecutionsUnknown(ctx context.Context, req *taskMarkExecutionsUnknownReq) (*taskMarkExecutionsUnknownResponse, error) {
	if len(req.IDs) == 0 {
		return &taskMarkExecutionsUnknownResponse{}, nil
	}
	recordsResp, err := u.listExecutionRecords(ctx, &taskListExecutionRecordsReq{Query: &repo.TaskExecutionRecordGetReq{IDs: req.IDs}})
	if err != nil {
		return nil, err
	}
	records := recordsResp.Rows
	runningByTask := make(map[int64][]int64)
	for _, row := range records {
		if row.Status == schedulerenum.TaskExecutionStatusRunning {
			runningByTask[row.TaskID] = append(runningByTask[row.TaskID], row.ID)
		}
	}
	activeByID := make(map[int64]bool)
	for taskID, executionIDs := range runningByTask {
		mapResp, err := u.mapRunning(ctx, &taskMapRunningReq{TaskID: taskID, ExecutionRecordIDs: executionIDs})
		if err != nil {
			return nil, err
		}
		for id, active := range mapResp.Rows {
			activeByID[id] = active
		}
	}
	type taskVersionKey struct {
		taskID  int64
		version int64
	}
	taskIDs := make([]int64, 0)
	versions := make([]int64, 0)
	needVersion := make(map[taskVersionKey]struct{})
	for _, row := range records {
		if row.Status != schedulerenum.TaskExecutionStatusRunning || activeByID[row.ID] {
			continue
		}
		key := taskVersionKey{taskID: row.TaskID, version: row.TaskVersion}
		if _, ok := needVersion[key]; ok {
			continue
		}
		needVersion[key] = struct{}{}
		taskIDs = append(taskIDs, row.TaskID)
		versions = append(versions, row.TaskVersion)
	}
	versionByKey := make(map[taskVersionKey]*model.TaskVersion, len(needVersion))
	if len(needVersion) > 0 {
		versionResp, err := u.listTaskVersions(ctx, &taskListTaskVersionsReq{Query: &repo.TaskVersionGetReq{TaskIDs: taskIDs, Versions: versions}})
		if err != nil {
			return nil, err
		}
		for _, row := range versionResp.Rows {
			key := taskVersionKey{taskID: row.TaskID, version: row.Version}
			if _, ok := needVersion[key]; ok {
				versionByKey[key] = row
			}
		}
	}
	now := time.Now()
	unknownIDs := make([]int64, 0)
	for _, row := range records {
		if row.Status != schedulerenum.TaskExecutionStatusRunning || activeByID[row.ID] || row.StartedAt == nil {
			continue
		}
		version := versionByKey[taskVersionKey{taskID: row.TaskID, version: row.TaskVersion}]
		if version != nil {
			timeout := time.Duration(version.TimeoutSeconds) * time.Second
			if timeout <= 0 && u.conf.GetScheduler() != nil && u.conf.GetScheduler().GetTaskTimeout() != nil && u.conf.GetScheduler().GetTaskTimeout().AsDuration() > 0 {
				timeout = u.conf.GetScheduler().GetTaskTimeout().AsDuration()
			}
			schedule, parseErr := u.cronParser.Parse(version.CronSpec)
			if timeout > 0 && parseErr == nil {
				next := schedule.Next(row.ScheduledAt)
				nextAfter := schedule.Next(next)
				if !next.IsZero() && !nextAfter.IsZero() && nextAfter.After(next) && now.After(row.StartedAt.Add(timeout+nextAfter.Sub(next))) {
					unknownIDs = append(unknownIDs, row.ID)
				}
			}
		}
	}
	unknownResp, err := u.markUnknownExecutionRecords(ctx, &taskMarkUnknownExecutionRecordsReq{IDs: unknownIDs, FinishedAt: now, LastError: "scheduler execution lost redis running heartbeat"})
	if err != nil {
		return nil, err
	}
	return &taskMarkExecutionsUnknownResponse{Rows: unknownResp.Rows}, nil
}

type taskCreateExecutionRecordReq struct {
	Record *model.TaskExecutionRecord
	Status schedulerenum.TaskExecutionStatus
}

type taskCreateExecutionRecordResponse struct {
	Record   *model.TaskExecutionRecord
	Created  bool
	Conflict bool
}

func (u *TaskUsecase) createExecutionRecord(ctx context.Context, req *taskCreateExecutionRecordReq) (*taskCreateExecutionRecordResponse, error) {
	resp, err := u.taskExecutionRecordRepo.Create(ctx, &repo.TaskExecutionRecordCreateReq{Record: req.Record, Status: req.Status})
	if err != nil {
		return nil, err
	}
	return &taskCreateExecutionRecordResponse{Record: resp.Row, Created: resp.Created, Conflict: resp.Conflict}, nil
}

type taskMarkExecutionFinishedReq struct {
	ID         int64
	Status     schedulerenum.TaskExecutionStatus
	FinishedAt time.Time
	DurationMS int64
	LastError  string
}

type taskMarkExecutionFinishedResponse struct {
	Record  *model.TaskExecutionRecord
	Updated bool
}

func (u *TaskUsecase) markExecutionFinished(ctx context.Context, req *taskMarkExecutionFinishedReq) (*taskMarkExecutionFinishedResponse, error) {
	resp, err := u.taskExecutionRecordRepo.MarkFinished(ctx, &repo.TaskExecutionRecordMarkFinishedReq{ID: req.ID, Status: req.Status, FinishedAt: req.FinishedAt, DurationMS: req.DurationMS, LastError: req.LastError})
	if err != nil {
		return nil, err
	}
	return &taskMarkExecutionFinishedResponse{Record: resp.Row, Updated: resp.Updated}, nil
}

type taskReleaseRunningReq struct {
	TaskID            int64
	ExecutionRecordID int64
	RunningToken      string
	Exclusive         bool
}

func (u *TaskUsecase) releaseRunning(ctx context.Context, req *taskReleaseRunningReq) error {
	_, err := u.taskLockRepo.ReleaseRunning(ctx, &repo.TaskRunningLockReq{TaskID: req.TaskID, ExecutionRecordID: req.ExecutionRecordID, RunningToken: req.RunningToken, Exclusive: req.Exclusive})
	return err
}

type taskRegisterRunningReq struct {
	TaskID            int64
	ExecutionRecordID int64
	RunningToken      string
	Exclusive         bool
	TTL               time.Duration
}

type taskRegisterRunningResponse struct {
	OK bool
}

func (u *TaskUsecase) registerRunning(ctx context.Context, req *taskRegisterRunningReq) (*taskRegisterRunningResponse, error) {
	resp, err := u.taskLockRepo.RegisterRunning(ctx, &repo.TaskRunningLockReq{TaskID: req.TaskID, ExecutionRecordID: req.ExecutionRecordID, RunningToken: req.RunningToken, Exclusive: req.Exclusive, TTL: req.TTL})
	if err != nil {
		return nil, err
	}
	return &taskRegisterRunningResponse{OK: resp.OK}, nil
}

type taskRefreshRunningReq struct {
	TaskID            int64
	ExecutionRecordID int64
	RunningToken      string
	Exclusive         bool
	TTL               time.Duration
}

type taskRefreshRunningResponse struct {
	OK bool
}

func (u *TaskUsecase) refreshRunning(ctx context.Context, req *taskRefreshRunningReq) (*taskRefreshRunningResponse, error) {
	resp, err := u.taskLockRepo.RefreshRunning(ctx, &repo.TaskRunningLockReq{TaskID: req.TaskID, ExecutionRecordID: req.ExecutionRecordID, RunningToken: req.RunningToken, Exclusive: req.Exclusive, TTL: req.TTL})
	if err != nil {
		return nil, err
	}
	return &taskRefreshRunningResponse{OK: resp.OK}, nil
}

type taskMapRunningReq struct {
	TaskID             int64
	ExecutionRecordIDs []int64
}

type taskMapRunningResponse struct {
	Rows map[int64]bool
}

func (u *TaskUsecase) mapRunning(ctx context.Context, req *taskMapRunningReq) (*taskMapRunningResponse, error) {
	resp, err := u.taskLockRepo.MapRunning(ctx, &repo.TaskRunningMapReq{TaskID: req.TaskID, ExecutionRecordIDs: req.ExecutionRecordIDs})
	if err != nil {
		return nil, err
	}
	return &taskMapRunningResponse{Rows: resp.Rows}, nil
}

type taskListExecutionRecordsReq struct {
	Query *repo.TaskExecutionRecordGetReq
}

type taskListExecutionRecordsResponse struct {
	Rows []*model.TaskExecutionRecord
}

func (u *TaskUsecase) listExecutionRecords(ctx context.Context, req *taskListExecutionRecordsReq) (*taskListExecutionRecordsResponse, error) {
	resp, err := u.taskExecutionRecordRepo.List(ctx, req.Query)
	if err != nil {
		return nil, err
	}
	return &taskListExecutionRecordsResponse{Rows: resp.Rows}, nil
}

type taskListTaskVersionsReq struct {
	Query *repo.TaskVersionGetReq
}

type taskListTaskVersionsResponse struct {
	Rows []*model.TaskVersion
}

func (u *TaskUsecase) listTaskVersions(ctx context.Context, req *taskListTaskVersionsReq) (*taskListTaskVersionsResponse, error) {
	resp, err := u.taskVersionRepo.List(ctx, req.Query)
	if err != nil {
		return nil, err
	}
	return &taskListTaskVersionsResponse{Rows: resp.Rows}, nil
}

type taskMarkUnknownExecutionRecordsReq struct {
	IDs        []int64
	FinishedAt time.Time
	LastError  string
}

type taskMarkUnknownExecutionRecordsResponse struct {
	Rows []*model.TaskExecutionRecord
}

func (u *TaskUsecase) markUnknownExecutionRecords(ctx context.Context, req *taskMarkUnknownExecutionRecordsReq) (*taskMarkUnknownExecutionRecordsResponse, error) {
	resp, err := u.taskExecutionRecordRepo.MarkUnknown(ctx, &repo.TaskExecutionRecordMarkUnknownReq{IDs: req.IDs, FinishedAt: req.FinishedAt, LastError: req.LastError})
	if err != nil {
		return nil, err
	}
	return &taskMarkUnknownExecutionRecordsResponse{Rows: resp.Rows}, nil
}

type TaskStopRunningReq struct{}

func (u *TaskUsecase) StopRunning(ctx context.Context, req *TaskStopRunningReq) error {
	if req == nil {
		req = &TaskStopRunningReq{}
	}
	return u.stopRunning(ctx, &taskStopRunningReq{})
}

type taskStopRunningReq struct{}

func (u *TaskUsecase) stopRunning(ctx context.Context, req *taskStopRunningReq) error {
	_ = req
	u.runningCancels.Range(func(_, value any) bool {
		value.(context.CancelFunc)()
		return true
	})
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
