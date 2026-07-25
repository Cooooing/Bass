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
	runningCancelMu         sync.Mutex
	runningCancels          map[int64]context.CancelFunc
	runningWG               sync.WaitGroup
}

func NewTaskUsecase(
	logger *slog.Logger,
	conf *config.Bootstrap,
	tx base.Tx,
	taskRepo repo.TaskRepo,
	taskVersionRepo repo.TaskVersionRepo,
	executionRepo repo.TaskExecutionRecordRepo,
	taskLockRepo repo.TaskLockRepo,
	tasks map[string]taskimpl.Task,
	taskEventBus repo.TaskEventBus,
	alert repo.TaskAlert,
) *TaskUsecase {
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
		runningCancels:          make(map[int64]context.CancelFunc),
	}
}

func (u *TaskUsecase) Upsert(ctx context.Context, row *model.Task) (*model.Task, error) {
	if row == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.upsert(ctx, row)
}

func (u *TaskUsecase) upsert(ctx context.Context, row *model.Task) (*model.Task, error) {
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
		saved, err = u.taskRepo.Upsert(txCtx, row)
		if err != nil {
			return err
		}
		_, err = u.taskVersionRepo.Create(txCtx, saved)
		return err
	})
	if err != nil {
		return nil, err
	}
	_ = u.taskEventBus.PublishTaskChanged(ctx, &repo.TaskChangedMessage{
		TaskID:  saved.ID,
		Version: saved.Version,
	})
	return saved, nil
}

func (u *TaskUsecase) Get(ctx context.Context, id int64) (*model.Task, error) {
	return u.taskRepo.Get(ctx, &repo.TaskGetReq{
		ID: new(id),
	})
}

type TaskPageReq struct {
	Page    *common.PageReq
	IDs     []int64
	Name    *string
	Title   *string
	Enabled *bool
}

type TaskPageResp struct {
	Rows []*model.Task
	Page *common.PageResp
}

func (u *TaskUsecase) Page(ctx context.Context, req *TaskPageReq) (*TaskPageResp, error) {
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
	return &TaskPageResp{
		Rows: pageResp.Rows,
		Page: pageResp.Page,
	}, nil
}

type TaskExecutionRecordPageReq struct {
	Page        *common.PageReq
	IDs         []int64
	TaskID      *int64
	Status      *schedulerenum.TaskExecutionStatus
	TriggerType *schedulerenum.TaskTriggerType
}

type TaskExecutionRecordPageResp struct {
	Rows []*model.TaskExecutionRecord
	Page *common.PageResp
}

func (u *TaskUsecase) PageExecutionRecords(ctx context.Context, req *TaskExecutionRecordPageReq) (*TaskExecutionRecordPageResp, error) {
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
	return &TaskExecutionRecordPageResp{
		Rows: pageResp.Rows,
		Page: pageResp.Page,
	}, nil
}

func (u *TaskUsecase) ListAvailableTasks(ctx context.Context, keyword string) ([]*model.AvailableTask, error) {
	return u.listAvailableTasks(ctx, keyword)
}

func (u *TaskUsecase) listAvailableTasks(ctx context.Context, keyword string) ([]*model.AvailableTask, error) {
	_ = ctx
	keyword = strings.ToLower(strings.TrimSpace(keyword))
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
	return rows, nil
}

type TaskTriggerReq struct {
	ID      int64
	Payload string
}

func (u *TaskUsecase) Trigger(ctx context.Context, req *TaskTriggerReq) (*model.TaskExecutionRecord, error) {
	if req == nil {
		req = &TaskTriggerReq{}
	}
	resp, err := u.trigger(ctx, &taskTriggerReq{
		ID:      req.ID,
		Payload: req.Payload,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type taskTriggerReq struct {
	ID      int64
	Payload string
}

func (u *TaskUsecase) trigger(ctx context.Context, req *taskTriggerReq) (*model.TaskExecutionRecord, error) {
	id := req.ID
	if id == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	task, err := u.taskRepo.Get(ctx, &repo.TaskGetReq{
		ID: new(id),
	})
	if err != nil {
		return nil, err
	}
	if req.Payload != "" {
		if !json.Valid([]byte(req.Payload)) {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		task.Payload = req.Payload
	}
	scheduleResp, err := u.scheduleExecution(ctx, &taskScheduleExecutionReq{
		Task:        task,
		ScheduledAt: time.Now(),
		TriggerType: schedulerenum.TaskTriggerTypeManual,
	})
	if err == nil && (scheduleResp == nil) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
	}
	if err != nil {
		return nil, err
	}
	return scheduleResp, nil
}

func (u *TaskUsecase) CancelExecution(ctx context.Context, id int64) (*model.TaskExecutionRecord, error) {
	return u.cancelExecution(ctx, id)
}

func (u *TaskUsecase) cancelExecution(ctx context.Context, id int64) (*model.TaskExecutionRecord, error) {
	if id == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	record, err := u.taskExecutionRecordRepo.Get(ctx, &repo.TaskExecutionRecordGetReq{
		ID: new(id),
	})
	if err != nil {
		return nil, err
	}
	if record.Status != schedulerenum.TaskExecutionStatusRunning {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := u.taskEventBus.PublishExecutionCanceled(ctx, &repo.TaskExecutionCanceledMessage{
		ExecutionRecordID: id,
	}); err != nil {
		return nil, err
	}
	return record, nil
}

func (u *TaskUsecase) CancelExecutionLocally(ctx context.Context, id int64) error {
	return u.cancelExecutionLocally(ctx, id)
}

func (u *TaskUsecase) cancelExecutionLocally(ctx context.Context, id int64) error {
	_ = ctx
	u.runningCancelMu.Lock()
	cancel := u.runningCancels[id]
	u.runningCancelMu.Unlock()
	if cancel != nil {
		cancel()
	}
	return nil
}

type TaskScheduleExecutionReq struct {
	Task        *model.Task
	ScheduledAt time.Time
	TriggerType schedulerenum.TaskTriggerType
}

func (u *TaskUsecase) ScheduleExecution(ctx context.Context, req *TaskScheduleExecutionReq) (*model.TaskExecutionRecord, error) {
	if req == nil || req.Task == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := u.scheduleExecution(ctx, &taskScheduleExecutionReq{
		Task:        req.Task,
		ScheduledAt: req.ScheduledAt,
		TriggerType: req.TriggerType,
	})
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, nil
	}
	return resp, nil
}

type taskScheduleExecutionReq struct {
	Task        *model.Task
	ScheduledAt time.Time
	TriggerType schedulerenum.TaskTriggerType
}

func (u *TaskUsecase) scheduleExecution(ctx context.Context, req *taskScheduleExecutionReq) (resp *model.TaskExecutionRecord, err error) {
	task := req.Task
	scheduledAt := req.ScheduledAt
	triggerType := req.TriggerType
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}
	var record *model.TaskExecutionRecord
	finish := func(row *model.TaskExecutionRecord, returnErr error) (*model.TaskExecutionRecord, error) {
		record = row
		if returnErr != nil {
			return nil, returnErr
		}
		if row == nil {
			return nil, nil
		}
		return row, nil
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
			startResp, err := u.startExecution(ctx, &taskStartExecutionReq{
				Task:        task,
				ScheduledAt: scheduledAt,
				TriggerType: triggerType,
			})
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
			if err := u.taskRepo.Lock(txCtx, task.ID); err != nil {
				return err
			}
			exists, err := u.taskExecutionRecordRepo.ExistsPeriod(txCtx, &repo.TaskExecutionRecordExistsPeriodReq{
				TaskID:      task.ID,
				ScheduledAt: scheduledAt,
			})
			if err != nil {
				return err
			}
			if exists {
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
			running, err := u.taskExecutionRecordRepo.HasUnexpiredRunning(txCtx, &repo.TaskExecutionRecordHasUnexpiredRunningReq{
				TaskID:       task.ID,
				StartedAfter: time.Now().Add(-(timeout + period)),
			})
			if err != nil {
				return err
			}
			if running {
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
				_ = u.alert.Alert(ctx, &repo.TaskAlertReq{
					Task:   task,
					Record: result,
					Reason: "overlap",
				})
			}
			return finish(result, nil)
		}
		startCreatedResp, err := u.startCreatedExecution(ctx, &taskStartCreatedExecutionReq{
			Task:        task,
			Record:      result,
			ScheduledAt: scheduledAt,
			TriggerType: triggerType,
		})
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
		startResp, err := u.startExecution(ctx, &taskStartExecutionReq{
			Task:         task,
			ScheduledAt:  scheduledAt,
			TriggerType:  triggerType,
			RunningToken: acquired.RunningToken,
		})
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
		overlapResp, err := u.createOverlapSkipped(ctx, &taskCreateOverlapSkippedReq{
			Task:        task,
			ScheduledAt: scheduledAt,
			TriggerType: triggerType,
		})
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

type TaskCreateOverlapSkippedResp struct {
	Record   *model.TaskExecutionRecord
	Created  bool
	Conflict bool
}

func (u *TaskUsecase) CreateOverlapSkipped(ctx context.Context, req *TaskCreateOverlapSkippedReq) (*TaskCreateOverlapSkippedResp, error) {
	if req == nil || req.Task == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := u.createOverlapSkipped(ctx, &taskCreateOverlapSkippedReq{
		Task:        req.Task,
		ScheduledAt: req.ScheduledAt,
		TriggerType: req.TriggerType,
	})
	if err != nil {
		return nil, err
	}
	return &TaskCreateOverlapSkippedResp{
		Record:   resp.Record,
		Created:  resp.Created,
		Conflict: resp.Conflict,
	}, nil
}

type taskCreateOverlapSkippedReq struct {
	Task        *model.Task
	ScheduledAt time.Time
	TriggerType schedulerenum.TaskTriggerType
}

type taskCreateOverlapSkippedResp struct {
	Record   *model.TaskExecutionRecord
	Created  bool
	Conflict bool
}

func (u *TaskUsecase) createOverlapSkipped(ctx context.Context, req *taskCreateOverlapSkippedReq) (*taskCreateOverlapSkippedResp, error) {
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
		_ = u.alert.Alert(ctx, &repo.TaskAlertReq{
			Task:   task,
			Record: createResp.Record,
			Reason: "overlap",
		})
	}
	if createResp == nil {
		return nil, err
	}
	return &taskCreateOverlapSkippedResp{
		Record:   createResp.Record,
		Created:  createResp.Created,
		Conflict: createResp.Conflict,
	}, err
}

type TaskStartExecutionReq struct {
	Task         *model.Task
	ScheduledAt  time.Time
	TriggerType  schedulerenum.TaskTriggerType
	RunningToken string
}

type TaskStartExecutionResp struct {
	Record   *model.TaskExecutionRecord
	Created  bool
	Conflict bool
}

func (u *TaskUsecase) StartExecution(ctx context.Context, req *TaskStartExecutionReq) (*TaskStartExecutionResp, error) {
	if req == nil || req.Task == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := u.startExecution(ctx, &taskStartExecutionReq{
		Task:         req.Task,
		ScheduledAt:  req.ScheduledAt,
		TriggerType:  req.TriggerType,
		RunningToken: req.RunningToken,
	})
	if err != nil {
		return nil, err
	}
	return &TaskStartExecutionResp{
		Record:   resp.Record,
		Created:  resp.Created,
		Conflict: resp.Conflict,
	}, nil
}

type taskStartExecutionReq struct {
	Task         *model.Task
	ScheduledAt  time.Time
	TriggerType  schedulerenum.TaskTriggerType
	RunningToken string
}

type taskStartExecutionResp struct {
	Record   *model.TaskExecutionRecord
	Created  bool
	Conflict bool
}

func (u *TaskUsecase) startExecution(ctx context.Context, req *taskStartExecutionReq) (*taskStartExecutionResp, error) {
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
		if releaseErr := u.releaseRunning(context.WithoutCancel(ctx), &taskReleaseRunningReq{
			TaskID:       task.ID,
			RunningToken: req.RunningToken,
			Exclusive:    !task.AllowOverlap && req.RunningToken != "",
		}); releaseErr != nil {
			u.logger.ErrorContext(context.WithoutCancel(ctx), "release scheduler running lock failed", constant.LogFieldTaskID, task.ID, constant.LogFieldErr, releaseErr)
		}
		if createResp == nil {
			return nil, err
		}
		return &taskStartExecutionResp{
			Record:   createResp.Record,
			Created:  createResp.Created,
			Conflict: createResp.Conflict,
		}, err
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
			failedResp, markErr := u.markExecutionFinished(context.WithoutCancel(ctx), &taskMarkExecutionFinishedReq{
				ID:         record.ID,
				Status:     schedulerenum.TaskExecutionStatusFailed,
				FinishedAt: finishedAt,
				DurationMS: durationMS,
				LastError:  parseErr.Error(),
			})
			if releaseErr := u.releaseRunning(context.WithoutCancel(ctx), &taskReleaseRunningReq{
				TaskID:            task.ID,
				ExecutionRecordID: record.ID,
				RunningToken:      req.RunningToken,
				Exclusive:         !task.AllowOverlap,
			}); releaseErr != nil {
				u.logger.ErrorContext(context.WithoutCancel(ctx), "release scheduler running lock failed", constant.LogFieldTaskID, task.ID, constant.LogFieldErr, releaseErr)
			}
			if markErr != nil {
				return &taskStartExecutionResp{
					Record:  failedResp.Record,
					Created: true,
				}, markErr
			}
			return &taskStartExecutionResp{
				Record:  failedResp.Record,
				Created: true,
			}, parseErr
		}
		registered, registerErr := u.registerRunning(context.WithoutCancel(ctx), &taskRegisterRunningReq{
			TaskID:            task.ID,
			ExecutionRecordID: record.ID,
			RunningToken:      req.RunningToken,
			Exclusive:         !task.AllowOverlap,
			TTL:               runningTTL,
		})
		if registerErr != nil || !registered {
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
			failedResp, markErr := u.markExecutionFinished(context.WithoutCancel(ctx), &taskMarkExecutionFinishedReq{
				ID:         record.ID,
				Status:     schedulerenum.TaskExecutionStatusFailed,
				FinishedAt: finishedAt,
				DurationMS: durationMS,
				LastError:  lastError,
			})
			if releaseErr := u.releaseRunning(context.WithoutCancel(ctx), &taskReleaseRunningReq{
				TaskID:            task.ID,
				ExecutionRecordID: record.ID,
				RunningToken:      req.RunningToken,
				Exclusive:         !task.AllowOverlap,
			}); releaseErr != nil {
				u.logger.ErrorContext(context.WithoutCancel(ctx), "release scheduler running lock failed", constant.LogFieldTaskID, task.ID, constant.LogFieldErr, releaseErr)
			}
			if markErr != nil {
				return &taskStartExecutionResp{
					Record:  failedResp.Record,
					Created: true,
				}, markErr
			}
			if registerErr != nil {
				return &taskStartExecutionResp{
					Record:  failedResp.Record,
					Created: true,
				}, registerErr
			}
			return &taskStartExecutionResp{
				Record:  failedResp.Record,
				Created: true,
			}, fmt.Errorf("scheduler redis running registration failed")
		}
	}
	startCreatedResp, err := u.startCreatedExecution(ctx, &taskStartCreatedExecutionReq{
		Task:         task,
		Record:       record,
		ScheduledAt:  req.ScheduledAt,
		TriggerType:  req.TriggerType,
		RunningToken: req.RunningToken,
		RunningTTL:   runningTTL,
	})
	if startCreatedResp == nil {
		return nil, err
	}
	return &taskStartExecutionResp{
		Record:   startCreatedResp.Record,
		Created:  startCreatedResp.Created,
		Conflict: startCreatedResp.Conflict,
	}, err
}

type taskStartCreatedExecutionReq struct {
	Task         *model.Task
	Record       *model.TaskExecutionRecord
	ScheduledAt  time.Time
	TriggerType  schedulerenum.TaskTriggerType
	RunningToken string
	RunningTTL   time.Duration
}

type taskStartCreatedExecutionResp struct {
	Record   *model.TaskExecutionRecord
	Created  bool
	Conflict bool
}

func (u *TaskUsecase) startCreatedExecution(ctx context.Context, req *taskStartCreatedExecutionReq) (*taskStartCreatedExecutionResp, error) {
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
	u.runningCancelMu.Lock()
	u.runningCancels[record.ID] = cancel
	u.runningCancelMu.Unlock()
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
			failedResp, markErr := u.markExecutionFinished(context.WithoutCancel(runCtx), &taskMarkExecutionFinishedReq{
				ID:         record.ID,
				Status:     schedulerenum.TaskExecutionStatusFailed,
				FinishedAt: finishedAt,
				DurationMS: durationMS,
				LastError:  configErr.Error(),
			})
			span.End()
			cancel()
			u.runningCancelMu.Lock()
			delete(u.runningCancels, record.ID)
			u.runningCancelMu.Unlock()
			if releaseErr := u.releaseRunning(context.WithoutCancel(runCtx), &taskReleaseRunningReq{
				TaskID:            task.ID,
				ExecutionRecordID: record.ID,
				RunningToken:      req.RunningToken,
				Exclusive:         !task.AllowOverlap,
			}); releaseErr != nil {
				u.logger.ErrorContext(context.WithoutCancel(runCtx), "release scheduler running lock failed", constant.LogFieldTaskID, task.ID, constant.LogFieldErr, releaseErr)
			}
			failedRecord := record
			if failedResp != nil && failedResp.Record != nil {
				failedRecord = failedResp.Record
			}
			if markErr != nil {
				return &taskStartCreatedExecutionResp{
					Record:  failedRecord,
					Created: true,
				}, markErr
			}
			return &taskStartCreatedExecutionResp{
				Record:  failedRecord,
				Created: true,
			}, configErr
		}
	}
	u.runningWG.Add(1)
	go u.executeRecord(callCtx, &taskExecuteRecordReq{
		Cancel:            cancel,
		Span:              span,
		Task:              task,
		Record:            record,
		ScheduledAt:       req.ScheduledAt,
		RunningToken:      req.RunningToken,
		HeartbeatInterval: heartbeatInterval,
		RunningTTL:        req.RunningTTL,
	})
	return &taskStartCreatedExecutionResp{
		Record:  record,
		Created: true,
	}, nil
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
	defer func() {
		u.runningCancelMu.Lock()
		delete(u.runningCancels, record.ID)
		u.runningCancelMu.Unlock()
	}()
	defer u.runningWG.Done()
	if req.RunningToken != "" {
		go u.heartbeatRunningLock(ctx, &taskHeartbeatRunningLockReq{
			Cancel:            req.Cancel,
			TaskID:            task.ID,
			ExecutionRecordID: record.ID,
			RunningToken:      req.RunningToken,
			Exclusive:         !task.AllowOverlap,
			Interval:          req.HeartbeatInterval,
			TTL:               req.RunningTTL,
		})
		defer func() {
			if err := u.releaseRunning(context.WithoutCancel(ctx), &taskReleaseRunningReq{
				TaskID:            task.ID,
				ExecutionRecordID: record.ID,
				RunningToken:      req.RunningToken,
				Exclusive:         !task.AllowOverlap,
			}); err != nil {
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
	markResp, markErr := u.markExecutionFinished(context.WithoutCancel(ctx), &taskMarkExecutionFinishedReq{
		ID:         record.ID,
		Status:     statusValue,
		FinishedAt: finishedAt,
		DurationMS: durationMS,
		LastError:  lastError,
	})
	if markErr != nil || markResp == nil || !markResp.Updated {
		return
	}
	finished := markResp.Record
	if task.AlertEnabled {
		if statusValue != schedulerenum.TaskExecutionStatusSuccess {
			_ = u.alert.Alert(ctx, &repo.TaskAlertReq{
				Task:   task,
				Record: finished,
				Reason: statusValue.String(),
			})
		}
		if statusValue == schedulerenum.TaskExecutionStatusSuccess {
			schedule, parseErr := u.cronParser.Parse(task.CronSpec)
			if parseErr == nil {
				next := schedule.Next(req.ScheduledAt)
				if record.StartedAt != nil && next.After(req.ScheduledAt) && finishedAt.Sub(*record.StartedAt) > next.Sub(req.ScheduledAt) {
					_ = u.alert.Alert(ctx, &repo.TaskAlertReq{
						Task:   task,
						Record: finished,
						Reason: "duration_exceeded",
					})
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
			refreshed, err := u.refreshRunning(context.WithoutCancel(ctx), &taskRefreshRunningReq{
				TaskID:            req.TaskID,
				ExecutionRecordID: req.ExecutionRecordID,
				RunningToken:      req.RunningToken,
				Exclusive:         req.Exclusive,
				TTL:               req.TTL,
			})
			if err == nil && refreshed {
				continue
			}
			if err != nil {
				u.logger.ErrorContext(context.WithoutCancel(ctx), "refresh scheduler running lock failed", constant.LogFieldTaskID, req.TaskID, constant.LogFieldErr, err)
			}
			// Redis 蹇冭烦澶辫触鏃朵富鍔ㄥ彇娑堟墽琛岋紝閬垮厤鏈湴浠诲姟缁х画鍗犵敤璧勬簮銆?			req.Cancel()
			return
		}
	}
}

func (u *TaskUsecase) CheckExecutionRuntimes(ctx context.Context, ids []int64) ([]*model.TaskExecutionRuntime, error) {
	return u.checkExecutionRuntimes(ctx, ids)
}

func (u *TaskUsecase) checkExecutionRuntimes(ctx context.Context, ids []int64) ([]*model.TaskExecutionRuntime, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	records, err := u.listExecutionRecords(ctx, &repo.TaskExecutionRecordGetReq{
		IDs: ids,
	})
	if err != nil {
		return nil, err
	}
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
		runningMap, mapErr := u.mapRunning(ctx, &taskMapRunningReq{
			TaskID:             taskID,
			ExecutionRecordIDs: executionIDs,
		})
		if mapErr != nil {
			redisFailedTask[taskID] = true
			u.logger.ErrorContext(ctx, "map scheduler running lock failed", constant.LogFieldTaskID, taskID, constant.LogFieldErr, mapErr)
			continue
		}
		for id, active := range runningMap {
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
		key := taskVersionKey{
			taskID:  row.TaskID,
			version: row.TaskVersion,
		}
		if _, ok := needVersion[key]; ok {
			continue
		}
		needVersion[key] = struct{}{}
		taskIDs = append(taskIDs, row.TaskID)
		versions = append(versions, row.TaskVersion)
	}
	versionByKey := make(map[taskVersionKey]*model.TaskVersion, len(needVersion))
	if len(needVersion) > 0 {
		versionsRows, err := u.listTaskVersions(ctx, &repo.TaskVersionGetReq{
			TaskIDs:  taskIDs,
			Versions: versions,
		})
		if err != nil {
			return nil, err
		}
		for _, row := range versionsRows {
			key := taskVersionKey{
				taskID:  row.TaskID,
				version: row.Version,
			}
			if _, ok := needVersion[key]; ok {
				versionByKey[key] = row
			}
		}
	}
	now := time.Now()
	result := make([]*model.TaskExecutionRuntime, 0, len(ids))
	for _, id := range ids {
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
				version := versionByKey[taskVersionKey{
					taskID:  row.TaskID,
					version: row.TaskVersion,
				}]
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
	return result, nil
}

func (u *TaskUsecase) MarkExecutionsUnknown(ctx context.Context, ids []int64) ([]*model.TaskExecutionRecord, error) {
	return u.markExecutionsUnknown(ctx, ids)
}

func (u *TaskUsecase) markExecutionsUnknown(ctx context.Context, ids []int64) ([]*model.TaskExecutionRecord, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	records, err := u.listExecutionRecords(ctx, &repo.TaskExecutionRecordGetReq{
		IDs: ids,
	})
	if err != nil {
		return nil, err
	}
	runningByTask := make(map[int64][]int64)
	for _, row := range records {
		if row.Status == schedulerenum.TaskExecutionStatusRunning {
			runningByTask[row.TaskID] = append(runningByTask[row.TaskID], row.ID)
		}
	}
	activeByID := make(map[int64]bool)
	for taskID, executionIDs := range runningByTask {
		runningMap, err := u.mapRunning(ctx, &taskMapRunningReq{
			TaskID:             taskID,
			ExecutionRecordIDs: executionIDs,
		})
		if err != nil {
			return nil, err
		}
		for id, active := range runningMap {
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
		key := taskVersionKey{
			taskID:  row.TaskID,
			version: row.TaskVersion,
		}
		if _, ok := needVersion[key]; ok {
			continue
		}
		needVersion[key] = struct{}{}
		taskIDs = append(taskIDs, row.TaskID)
		versions = append(versions, row.TaskVersion)
	}
	versionByKey := make(map[taskVersionKey]*model.TaskVersion, len(needVersion))
	if len(needVersion) > 0 {
		versionsRows, err := u.listTaskVersions(ctx, &repo.TaskVersionGetReq{
			TaskIDs:  taskIDs,
			Versions: versions,
		})
		if err != nil {
			return nil, err
		}
		for _, row := range versionsRows {
			key := taskVersionKey{
				taskID:  row.TaskID,
				version: row.Version,
			}
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
		version := versionByKey[taskVersionKey{
			taskID:  row.TaskID,
			version: row.TaskVersion,
		}]
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
	unknownResp, err := u.markUnknownExecutionRecords(ctx, &taskMarkUnknownExecutionRecordsReq{
		IDs:        unknownIDs,
		FinishedAt: now,
		LastError:  "scheduler execution lost redis running heartbeat",
	})
	if err != nil {
		return nil, err
	}
	return unknownResp, nil
}

type taskCreateExecutionRecordReq struct {
	Record *model.TaskExecutionRecord
	Status schedulerenum.TaskExecutionStatus
}

type taskCreateExecutionRecordResp struct {
	Record   *model.TaskExecutionRecord
	Created  bool
	Conflict bool
}

func (u *TaskUsecase) createExecutionRecord(ctx context.Context, req *taskCreateExecutionRecordReq) (*taskCreateExecutionRecordResp, error) {
	resp, err := u.taskExecutionRecordRepo.Create(ctx, &repo.TaskExecutionRecordCreateReq{
		Record: req.Record,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &taskCreateExecutionRecordResp{
		Record:   resp.Row,
		Created:  resp.Created,
		Conflict: resp.Conflict,
	}, nil
}

type taskMarkExecutionFinishedReq struct {
	ID         int64
	Status     schedulerenum.TaskExecutionStatus
	FinishedAt time.Time
	DurationMS int64
	LastError  string
}

type taskMarkExecutionFinishedResp struct {
	Record  *model.TaskExecutionRecord
	Updated bool
}

func (u *TaskUsecase) markExecutionFinished(ctx context.Context, req *taskMarkExecutionFinishedReq) (*taskMarkExecutionFinishedResp, error) {
	resp, err := u.taskExecutionRecordRepo.MarkFinished(ctx, &repo.TaskExecutionRecordMarkFinishedReq{
		ID:         req.ID,
		Status:     req.Status,
		FinishedAt: req.FinishedAt,
		DurationMS: req.DurationMS,
		LastError:  req.LastError,
	})
	if err != nil {
		return nil, err
	}
	return &taskMarkExecutionFinishedResp{
		Record:  resp.Row,
		Updated: resp.Updated,
	}, nil
}

type taskReleaseRunningReq struct {
	TaskID            int64
	ExecutionRecordID int64
	RunningToken      string
	Exclusive         bool
}

func (u *TaskUsecase) releaseRunning(ctx context.Context, req *taskReleaseRunningReq) error {
	err := u.taskLockRepo.ReleaseRunning(ctx, &repo.TaskRunningLockReq{
		TaskID:            req.TaskID,
		ExecutionRecordID: req.ExecutionRecordID,
		RunningToken:      req.RunningToken,
		Exclusive:         req.Exclusive,
	})
	return err
}

type taskRegisterRunningReq struct {
	TaskID            int64
	ExecutionRecordID int64
	RunningToken      string
	Exclusive         bool
	TTL               time.Duration
}

func (u *TaskUsecase) registerRunning(ctx context.Context, req *taskRegisterRunningReq) (bool, error) {
	ok, err := u.taskLockRepo.RegisterRunning(ctx, &repo.TaskRunningLockReq{
		TaskID:            req.TaskID,
		ExecutionRecordID: req.ExecutionRecordID,
		RunningToken:      req.RunningToken,
		Exclusive:         req.Exclusive,
		TTL:               req.TTL,
	})
	if err != nil {
		return false, err
	}
	return ok, nil
}

type taskRefreshRunningReq struct {
	TaskID            int64
	ExecutionRecordID int64
	RunningToken      string
	Exclusive         bool
	TTL               time.Duration
}

func (u *TaskUsecase) refreshRunning(ctx context.Context, req *taskRefreshRunningReq) (bool, error) {
	ok, err := u.taskLockRepo.RefreshRunning(ctx, &repo.TaskRunningLockReq{
		TaskID:            req.TaskID,
		ExecutionRecordID: req.ExecutionRecordID,
		RunningToken:      req.RunningToken,
		Exclusive:         req.Exclusive,
		TTL:               req.TTL,
	})
	if err != nil {
		return false, err
	}
	return ok, nil
}

type taskMapRunningReq struct {
	TaskID             int64
	ExecutionRecordIDs []int64
}

func (u *TaskUsecase) mapRunning(ctx context.Context, req *taskMapRunningReq) (map[int64]bool, error) {
	rows, err := u.taskLockRepo.MapRunning(ctx, &repo.TaskRunningMapReq{
		TaskID:             req.TaskID,
		ExecutionRecordIDs: req.ExecutionRecordIDs,
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (u *TaskUsecase) listExecutionRecords(ctx context.Context, query *repo.TaskExecutionRecordGetReq) ([]*model.TaskExecutionRecord, error) {
	rows, err := u.taskExecutionRecordRepo.List(ctx, query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (u *TaskUsecase) listTaskVersions(ctx context.Context, query *repo.TaskVersionGetReq) ([]*model.TaskVersion, error) {
	rows, err := u.taskVersionRepo.List(ctx, query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

type taskMarkUnknownExecutionRecordsReq struct {
	IDs        []int64
	FinishedAt time.Time
	LastError  string
}

func (u *TaskUsecase) markUnknownExecutionRecords(ctx context.Context, req *taskMarkUnknownExecutionRecordsReq) ([]*model.TaskExecutionRecord, error) {
	rows, err := u.taskExecutionRecordRepo.MarkUnknown(ctx, &repo.TaskExecutionRecordMarkUnknownReq{
		IDs:        req.IDs,
		FinishedAt: req.FinishedAt,
		LastError:  req.LastError,
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (u *TaskUsecase) StopRunning(ctx context.Context) error {
	return u.stopRunning(ctx)
}

func (u *TaskUsecase) stopRunning(ctx context.Context) error {
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
