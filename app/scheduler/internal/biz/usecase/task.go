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
	"scheduler/internal/conf"
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
	conf                    *conf.Bootstrap
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

func NewTaskUsecase(logger *slog.Logger, conf *conf.Bootstrap, tx base.Tx, taskRepo repo.TaskRepo, taskVersionRepo repo.TaskVersionRepo, executionRepo repo.TaskExecutionRecordRepo, taskLockRepo repo.TaskLockRepo, tasks map[string]taskimpl.Task, taskEventBus repo.TaskEventBus, alert repo.TaskAlert) *TaskUsecase {
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

func (u *TaskUsecase) Upsert(ctx context.Context, row *model.Task) (*model.Task, error) {
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
	_ = u.taskEventBus.PublishTaskChanged(ctx, &repo.TaskChangedMessage{TaskID: saved.ID, Version: saved.Version})
	return saved, nil
}

func (u *TaskUsecase) Get(ctx context.Context, id int64) (*model.Task, error) {
	return u.taskRepo.Get(ctx, &repo.TaskGetReq{ID: &id})
}

func (u *TaskUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.TaskGetReq) ([]*model.Task, *common.PageReply, error) {
	return u.taskRepo.Page(ctx, page, req)
}

func (u *TaskUsecase) PageExecutionRecords(ctx context.Context, page *common.PageRequest, req *repo.TaskExecutionRecordGetReq) ([]*model.TaskExecutionRecord, *common.PageReply, error) {
	return u.taskExecutionRecordRepo.Page(ctx, page, req)
}

func (u *TaskUsecase) ListAvailableTasks(keyword string) []*model.AvailableTask {
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
	return rows
}

func (u *TaskUsecase) Trigger(ctx context.Context, id int64, payload string) (*model.TaskExecutionRecord, error) {
	if id == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	task, err := u.taskRepo.Get(ctx, &repo.TaskGetReq{ID: &id})
	if err != nil {
		return nil, err
	}
	if payload != "" {
		if !json.Valid([]byte(payload)) {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		task.Payload = payload
	}
	record, err := u.ScheduleExecution(ctx, task, time.Now(), schedulerenum.TaskTriggerTypeManual)
	if err == nil && record == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
	}
	return record, err
}

func (u *TaskUsecase) CancelExecution(ctx context.Context, id int64) (*model.TaskExecutionRecord, error) {
	if id == 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	record, err := u.taskExecutionRecordRepo.Get(ctx, &repo.TaskExecutionRecordGetReq{ID: &id})
	if err != nil {
		return nil, err
	}
	if record.Status != schedulerenum.TaskExecutionStatusRunning {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := u.taskEventBus.PublishExecutionCanceled(ctx, &repo.TaskExecutionCanceledMessage{ExecutionRecordID: id}); err != nil {
		return nil, err
	}
	return record, nil
}

func (u *TaskUsecase) CancelExecutionLocally(id int64) {
	if cancelValue, ok := u.runningCancels.Load(id); ok {
		cancelValue.(context.CancelFunc)()
	}
}

func (u *TaskUsecase) ScheduleExecution(ctx context.Context, task *model.Task, scheduledAt time.Time, triggerType schedulerenum.TaskTriggerType) (record *model.TaskExecutionRecord, err error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
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
			result, created, conflict, err := u.StartExecution(ctx, task, scheduledAt, triggerType, "")
			if err != nil {
				return result, err
			}
			if !created || conflict {
				if conflict {
					skipReason = "db_unique_conflict"
				} else {
					skipReason = "db_create_skipped"
				}
				return nil, nil
			}
			return result, nil
		}
		var result *model.TaskExecutionRecord
		var created bool
		var conflict bool
		err = u.tx(ctx, func(txCtx context.Context) error {
			if err := u.taskRepo.Lock(txCtx, task.ID); err != nil {
				return err
			}
			exists, err := u.taskExecutionRecordRepo.ExistsPeriod(txCtx, task.ID, scheduledAt)
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
			running, err := u.taskExecutionRecordRepo.HasUnexpiredRunning(txCtx, task.ID, time.Now().Add(-(timeout + period)))
			if err != nil {
				return err
			}
			if running {
				traceID := ""
				if oteltrace.SpanContextFromContext(txCtx).IsValid() {
					traceID = oteltrace.SpanContextFromContext(txCtx).TraceID().String()
				}
				record, recordCreated, recordConflict, err := u.taskExecutionRecordRepo.Create(txCtx, &model.TaskExecutionRecord{
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
				}, schedulerenum.TaskExecutionStatusOverlapSkipped)
				result = record
				created = recordCreated
				conflict = recordConflict
				if recordConflict {
					skipReason = "db_unique_conflict"
				}
				return err
			}
			traceID := ""
			if oteltrace.SpanContextFromContext(txCtx).IsValid() {
				traceID = oteltrace.SpanContextFromContext(txCtx).TraceID().String()
			}
			record, recordCreated, recordConflict, err := u.taskExecutionRecordRepo.Create(txCtx, &model.TaskExecutionRecord{
				TaskID:      task.ID,
				ScheduledAt: scheduledAt,
				StartedAt:   new(time.Now()),
				Status:      schedulerenum.TaskExecutionStatusRunning,
				TriggerType: triggerType,
				TaskVersion: task.Version,
				Payload:     task.Payload,
				TraceID:     traceID,
			}, schedulerenum.TaskExecutionStatusRunning)
			result = record
			created = recordCreated
			conflict = recordConflict
			if recordConflict {
				skipReason = "db_unique_conflict"
			}
			return err
		})
		if err != nil {
			return result, err
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
				_ = u.alert.Alert(ctx, task, result, "overlap")
			}
			return result, nil
		}
		result, _, _, err = u.startCreatedExecution(ctx, task, result, scheduledAt, triggerType, "", 0)
		return result, err
	}
	if acquired == nil {
		return nil, fmt.Errorf("scheduler task lock acquire result is nil")
	}
	switch acquired.Decision {
	case schedulerenum.TaskScheduleDecisionSkip:
		skipReason = "redis_schedule_skip"
		return nil, nil
	case schedulerenum.TaskScheduleDecisionRun:
		result, created, conflict, err := u.StartExecution(ctx, task, scheduledAt, triggerType, acquired.RunningToken)
		if err != nil {
			return result, err
		}
		if !created || conflict {
			if conflict {
				skipReason = "db_unique_conflict"
			} else {
				skipReason = "db_create_skipped"
			}
			return nil, nil
		}
		return result, nil
	case schedulerenum.TaskScheduleDecisionOverlap:
		result, created, conflict, err := u.CreateOverlapSkipped(ctx, task, scheduledAt, triggerType)
		if err != nil {
			return result, err
		}
		if !created || conflict {
			if conflict {
				skipReason = "db_unique_conflict"
			} else {
				skipReason = "db_create_skipped"
			}
			return nil, nil
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unknown scheduler schedule decision: %s", acquired.Decision)
	}
}

func (u *TaskUsecase) CreateOverlapSkipped(ctx context.Context, task *model.Task, scheduledAt time.Time, triggerType schedulerenum.TaskTriggerType) (*model.TaskExecutionRecord, bool, bool, error) {
	traceID := ""
	if oteltrace.SpanContextFromContext(ctx).IsValid() {
		traceID = oteltrace.SpanContextFromContext(ctx).TraceID().String()
	}
	record, created, conflict, err := u.taskExecutionRecordRepo.Create(ctx, &model.TaskExecutionRecord{
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
	}, schedulerenum.TaskExecutionStatusOverlapSkipped)
	if conflict {
		u.logger.WarnContext(ctx, "scheduler execution record already exists", constant.LogFieldTaskID, task.ID, constant.LogFieldTaskName, task.Name)
	}
	if err == nil && created && task.AlertEnabled {
		_ = u.alert.Alert(ctx, task, record, "overlap")
	}
	return record, created, conflict, err
}

func (u *TaskUsecase) StartExecution(ctx context.Context, task *model.Task, scheduledAt time.Time, triggerType schedulerenum.TaskTriggerType, runningToken string) (*model.TaskExecutionRecord, bool, bool, error) {
	traceID := ""
	if oteltrace.SpanContextFromContext(ctx).IsValid() {
		traceID = oteltrace.SpanContextFromContext(ctx).TraceID().String()
	}
	record, created, conflict, err := u.taskExecutionRecordRepo.Create(ctx, &model.TaskExecutionRecord{
		TaskID:      task.ID,
		ScheduledAt: scheduledAt,
		StartedAt:   new(time.Now()),
		Status:      schedulerenum.TaskExecutionStatusRunning,
		TriggerType: triggerType,
		TaskVersion: task.Version,
		Payload:     task.Payload,
		TraceID:     traceID,
	}, schedulerenum.TaskExecutionStatusRunning)
	if conflict {
		u.logger.WarnContext(ctx, "scheduler execution record already exists", constant.LogFieldTaskID, task.ID, constant.LogFieldTaskName, task.Name)
	}
	if err != nil || !created {
		if releaseErr := u.taskLockRepo.ReleaseRunning(context.WithoutCancel(ctx), task.ID, 0, runningToken, !task.AllowOverlap && runningToken != ""); releaseErr != nil {
			u.logger.ErrorContext(context.WithoutCancel(ctx), "release scheduler running lock failed", constant.LogFieldTaskID, task.ID, constant.LogFieldErr, releaseErr)
		}
		return record, created, conflict, err
	}
	runningTTL := time.Duration(0)
	if runningToken != "" {
		schedule, parseErr := u.cronParser.Parse(task.CronSpec)
		if parseErr == nil {
			next := schedule.Next(scheduledAt)
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
			failed, _, markErr := u.taskExecutionRecordRepo.MarkFinished(context.WithoutCancel(ctx), record.ID, schedulerenum.TaskExecutionStatusFailed, finishedAt, durationMS, parseErr.Error())
			if releaseErr := u.taskLockRepo.ReleaseRunning(context.WithoutCancel(ctx), task.ID, record.ID, runningToken, !task.AllowOverlap); releaseErr != nil {
				u.logger.ErrorContext(context.WithoutCancel(ctx), "release scheduler running lock failed", constant.LogFieldTaskID, task.ID, constant.LogFieldErr, releaseErr)
			}
			if markErr != nil {
				return failed, true, false, markErr
			}
			return failed, true, false, parseErr
		}
		ok, registerErr := u.taskLockRepo.RegisterRunning(context.WithoutCancel(ctx), task.ID, record.ID, runningToken, !task.AllowOverlap, runningTTL)
		if registerErr != nil || !ok {
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
			failed, _, markErr := u.taskExecutionRecordRepo.MarkFinished(context.WithoutCancel(ctx), record.ID, schedulerenum.TaskExecutionStatusFailed, finishedAt, durationMS, lastError)
			if releaseErr := u.taskLockRepo.ReleaseRunning(context.WithoutCancel(ctx), task.ID, record.ID, runningToken, !task.AllowOverlap); releaseErr != nil {
				u.logger.ErrorContext(context.WithoutCancel(ctx), "release scheduler running lock failed", constant.LogFieldTaskID, task.ID, constant.LogFieldErr, releaseErr)
			}
			if markErr != nil {
				return failed, true, false, markErr
			}
			if registerErr != nil {
				return failed, true, false, registerErr
			}
			return failed, true, false, fmt.Errorf("scheduler redis running registration failed")
		}
	}
	return u.startCreatedExecution(ctx, task, record, scheduledAt, triggerType, runningToken, runningTTL)
}

func (u *TaskUsecase) startCreatedExecution(ctx context.Context, task *model.Task, record *model.TaskExecutionRecord, scheduledAt time.Time, triggerType schedulerenum.TaskTriggerType, runningToken string, runningTTL time.Duration) (*model.TaskExecutionRecord, bool, bool, error) {
	runCtx, span := otel.Tracer("scheduler.task").Start(
		ctx,
		"scheduler.task "+task.Name,
		oteltrace.WithSpanKind(oteltrace.SpanKindInternal),
		oteltrace.WithAttributes(
			attribute.Int64("scheduler.task.id", task.ID),
			attribute.String("scheduler.task.name", task.Name),
			attribute.String("scheduler.task.title", task.Title),
			attribute.String("scheduler.trigger_type", string(triggerType)),
			attribute.String("scheduler.scheduled_at", scheduledAt.Format(time.RFC3339Nano)),
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
			return nil, false, false, fmt.Errorf("scheduler task_timeout is invalid")
		}
		timeout = u.conf.GetScheduler().GetTaskTimeout().AsDuration()
	}
	span.SetAttributes(attribute.Int64("scheduler.execution.id", record.ID))
	callCtx, cancel := context.WithTimeout(context.WithoutCancel(runCtx), timeout)
	u.runningCancels.Store(record.ID, cancel)
	heartbeatInterval := time.Duration(0)
	if runningToken != "" {
		var configErr error
		if runningTTL <= 0 {
			configErr = fmt.Errorf("scheduler cron period is invalid")
		} else {
			heartbeatInterval = runningTTL / 3
			if heartbeatInterval < 100*time.Millisecond {
				heartbeatInterval = 100 * time.Millisecond
			}
			if heartbeatInterval >= runningTTL {
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
			failed, _, markErr := u.taskExecutionRecordRepo.MarkFinished(context.WithoutCancel(runCtx), record.ID, schedulerenum.TaskExecutionStatusFailed, finishedAt, durationMS, configErr.Error())
			span.End()
			cancel()
			u.runningCancels.Delete(record.ID)
			if releaseErr := u.taskLockRepo.ReleaseRunning(context.WithoutCancel(runCtx), task.ID, record.ID, runningToken, !task.AllowOverlap); releaseErr != nil {
				u.logger.ErrorContext(context.WithoutCancel(runCtx), "release scheduler running lock failed", constant.LogFieldTaskID, task.ID, constant.LogFieldErr, releaseErr)
			}
			if markErr != nil {
				return failed, true, false, markErr
			}
			return failed, true, false, configErr
		}
	}
	u.runningWG.Add(1)
	go u.executeRecord(callCtx, cancel, span, task, record, scheduledAt, runningToken, heartbeatInterval, runningTTL)
	return record, true, false, nil
}

func (u *TaskUsecase) executeRecord(ctx context.Context, cancel context.CancelFunc, span oteltrace.Span, task *model.Task, record *model.TaskExecutionRecord, scheduledAt time.Time, runningToken string, heartbeatInterval time.Duration, runningTTL time.Duration) {
	defer span.End()
	defer cancel()
	defer u.runningCancels.Delete(record.ID)
	defer u.runningWG.Done()
	if runningToken != "" {
		go u.heartbeatRunningLock(ctx, cancel, task.ID, record.ID, runningToken, !task.AllowOverlap, heartbeatInterval, runningTTL)
		defer func() {
			if err := u.taskLockRepo.ReleaseRunning(context.WithoutCancel(ctx), task.ID, record.ID, runningToken, !task.AllowOverlap); err != nil {
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
	finished, updated, markErr := u.taskExecutionRecordRepo.MarkFinished(context.WithoutCancel(ctx), record.ID, statusValue, finishedAt, durationMS, lastError)
	if markErr != nil || !updated {
		return
	}
	if task.AlertEnabled {
		if statusValue != schedulerenum.TaskExecutionStatusSuccess {
			_ = u.alert.Alert(ctx, task, finished, statusValue.String())
		}
		if statusValue == schedulerenum.TaskExecutionStatusSuccess {
			schedule, parseErr := u.cronParser.Parse(task.CronSpec)
			if parseErr == nil {
				next := schedule.Next(scheduledAt)
				if record.StartedAt != nil && next.After(scheduledAt) && finishedAt.Sub(*record.StartedAt) > next.Sub(scheduledAt) {
					_ = u.alert.Alert(ctx, task, finished, "duration_exceeded")
				}
			}
		}
	}
}

func (u *TaskUsecase) heartbeatRunningLock(ctx context.Context, cancel context.CancelFunc, taskID int64, executionRecordID int64, runningToken string, exclusive bool, interval time.Duration, ttl time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			ok, err := u.taskLockRepo.RefreshRunning(context.WithoutCancel(ctx), taskID, executionRecordID, runningToken, exclusive, ttl)
			if err == nil && ok {
				continue
			}
			if err != nil {
				u.logger.ErrorContext(context.WithoutCancel(ctx), "refresh scheduler running lock failed", constant.LogFieldTaskID, taskID, constant.LogFieldErr, err)
			}
			// 续期失败表示运行互斥所有权已丢失或 Redis 状态不可确认，本地执行必须尽快停止。
			cancel()
			return
		}
	}
}

func (u *TaskUsecase) CheckExecutionRuntimes(ctx context.Context, ids []int64) ([]*model.TaskExecutionRuntime, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	records, err := u.taskExecutionRecordRepo.List(ctx, &repo.TaskExecutionRecordGetReq{IDs: ids})
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
		activeRows, mapErr := u.taskLockRepo.MapRunning(ctx, taskID, executionIDs)
		if mapErr != nil {
			redisFailedTask[taskID] = true
			u.logger.ErrorContext(ctx, "map scheduler running lock failed", constant.LogFieldTaskID, taskID, constant.LogFieldErr, mapErr)
			continue
		}
		for id, active := range activeRows {
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
		versionRows, err := u.taskVersionRepo.List(ctx, &repo.TaskVersionGetReq{TaskIDs: taskIDs, Versions: versions})
		if err != nil {
			return nil, err
		}
		for _, row := range versionRows {
			key := taskVersionKey{taskID: row.TaskID, version: row.Version}
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
	return result, nil
}

func (u *TaskUsecase) MarkExecutionsUnknown(ctx context.Context, ids []int64) ([]*model.TaskExecutionRecord, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	records, err := u.taskExecutionRecordRepo.List(ctx, &repo.TaskExecutionRecordGetReq{IDs: ids})
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
		activeRows, err := u.taskLockRepo.MapRunning(ctx, taskID, executionIDs)
		if err != nil {
			return nil, err
		}
		for id, active := range activeRows {
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
		versionRows, err := u.taskVersionRepo.List(ctx, &repo.TaskVersionGetReq{TaskIDs: taskIDs, Versions: versions})
		if err != nil {
			return nil, err
		}
		for _, row := range versionRows {
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
	return u.taskExecutionRecordRepo.MarkUnknown(ctx, unknownIDs, now, "scheduler execution lost redis running heartbeat")
}

func (u *TaskUsecase) StopRunning(ctx context.Context) error {
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
