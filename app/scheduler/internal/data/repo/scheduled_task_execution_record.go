package repo

import (
	"common/pkg/apperror"
	"common/pkg/server"
	utilent "common/pkg/util/ent"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
	"scheduler/internal/biz/model"
	bizrepo "scheduler/internal/biz/repo"
	"scheduler/internal/data/gen"
	"scheduler/internal/data/gen/scheduledtaskexecutionrecord"
	schedulerenum "scheduler/internal/enum"
	"strings"
	"time"

	entsql "entgo.io/ent/dialect/sql"
)

type ScheduledTaskExecutionRecordRepo struct {
	db *gen.Client
}

func NewScheduledTaskExecutionRecordRepo(
	db *gen.Client,
) bizrepo.ScheduledTaskExecutionRecordRepo {
	return &ScheduledTaskExecutionRecordRepo{db: db}
}

func (r *ScheduledTaskExecutionRecordRepo) Get(ctx context.Context, req *bizrepo.ScheduledTaskExecutionRecordGetReq) (*model.ScheduledTaskExecutionRecord, error) {
	row, err := r.getQuery(r.getClient(ctx).ScheduledTaskExecutionRecord.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *ScheduledTaskExecutionRecordRepo) List(ctx context.Context, req *bizrepo.ScheduledTaskExecutionRecordGetReq) ([]*model.ScheduledTaskExecutionRecord, error) {
	rows, err := r.getQuery(r.getClient(ctx).ScheduledTaskExecutionRecord.Query(), req).Order(scheduledtaskexecutionrecord.ByScheduledAt(entsql.OrderDesc())).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.ScheduledTaskExecutionRecord, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return result, nil
}

func (r *ScheduledTaskExecutionRecordRepo) Page(ctx context.Context, req *bizrepo.ScheduledTaskExecutionRecordPageReq) (*bizrepo.ScheduledTaskExecutionRecordPageResp, error) {
	page := server.PageValid(req.Page)
	query := r.getQuery(r.getClient(ctx).ScheduledTaskExecutionRecord.Query(), &req.ScheduledTaskExecutionRecordGetReq)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := query.Order(scheduledtaskexecutionrecord.ByScheduledAt(entsql.OrderDesc())).Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.ScheduledTaskExecutionRecord, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return &bizrepo.ScheduledTaskExecutionRecordPageResp{Rows: result, Page: &common.PageResp{Total: uint32(total), Page: page.Page, Size: page.Size}}, nil
}

func (r *ScheduledTaskExecutionRecordRepo) HasRunning(ctx context.Context, req *bizrepo.ScheduledTaskExecutionRecordHasRunningReq) (bool, error) {
	rows, err := r.getClient(ctx).ScheduledTaskExecutionRecord.Query().
		Where(
			scheduledtaskexecutionrecord.ScheduledTaskIDEQ(req.ScheduledTaskID),
			scheduledtaskexecutionrecord.StatusEQ(scheduledtaskexecutionrecord.Status(schedulerenum.TaskExecutionStatusRunning)),
		).
		All(ctx)
	if err != nil {
		return false, err
	}
	now := time.Now()
	for _, row := range rows {
		if row.StartedAt == nil {
			return true, nil
		}
		timeout := time.Duration(row.TimeoutSeconds) * time.Second
		if timeout <= 0 {
			timeout = 300 * time.Second
		}
		if row.StartedAt.Add(timeout).After(now) {
			return true, nil
		}
	}
	return false, nil
}

func (r *ScheduledTaskExecutionRecordRepo) Create(
	ctx context.Context,
	req *bizrepo.ScheduledTaskExecutionRecordCreateReq,
) (*bizrepo.ScheduledTaskExecutionRecordCreateResp, error) {
	record := req.Record
	var durationMs *int64
	if record.Duration != nil {
		durationMs = new(record.Duration.Milliseconds())
	}
	var staleAfterSeconds *int32
	if record.StaleAfter != nil {
		staleAfterSeconds = new(int32(*record.StaleAfter / time.Second))
	}
	created, err := r.getClient(ctx).ScheduledTaskExecutionRecord.Create().
		SetScheduledTaskID(record.ScheduledTaskID).
		SetScheduledTaskVersion(record.ScheduledTaskVersion).
		SetTriggerType(scheduledtaskexecutionrecord.TriggerType(record.TriggerType)).
		SetScheduleKey(record.ScheduleKey).
		SetScheduledAt(record.ScheduledAt).
		SetNillableStartedAt(record.StartedAt).
		SetNillableFinishedAt(record.FinishedAt).
		SetNillableDurationMs(durationMs).
		SetStatus(scheduledtaskexecutionrecord.Status(req.Status)).
		SetAttempt(record.Attempt).
		SetMaxAttempts(record.MaxAttempts).
		SetTimeoutSeconds(int32(record.Timeout / time.Second)).
		SetNillableStaleAfterSeconds(staleAfterSeconds).
		SetMisfirePolicy(scheduledtaskexecutionrecord.MisfirePolicy(record.MisfirePolicy)).
		SetWorkerID(record.WorkerID).
		SetPayload(record.Payload).
		SetLastError(record.LastError).
		SetTraceID(record.TraceID).
		Save(ctx)
	if gen.IsConstraintError(err) && strings.Contains(err.Error(), "scheduler_scheduled_task_execution_records_task_schedule_key_unique") {
		return &bizrepo.ScheduledTaskExecutionRecordCreateResp{Conflict: true}, nil
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.ScheduledTaskExecutionRecordCreateResp{Row: r.model(created), Created: true}, nil
}

func (r *ScheduledTaskExecutionRecordRepo) Claim(ctx context.Context, req *bizrepo.ScheduledTaskExecutionRecordClaimReq) (*bizrepo.ScheduledTaskExecutionRecordClaimResp, error) {
	affected, err := r.getClient(ctx).ScheduledTaskExecutionRecord.Update().
		Where(
			scheduledtaskexecutionrecord.ID(req.ID),
			scheduledtaskexecutionrecord.StatusIn(
				scheduledtaskexecutionrecord.Status(schedulerenum.TaskExecutionStatusPending),
				scheduledtaskexecutionrecord.Status(schedulerenum.TaskExecutionStatusRetryPending),
				scheduledtaskexecutionrecord.Status(schedulerenum.TaskExecutionStatusRunning),
			),
		).
		SetStatus(scheduledtaskexecutionrecord.Status(schedulerenum.TaskExecutionStatusRunning)).
		SetStartedAt(req.StartedAt).
		ClearFinishedAt().
		ClearDurationMs().
		SetWorkerID(req.WorkerID).
		SetLastError("").
		AddAttempt(1).
		Save(ctx)
	if err != nil || affected == 0 {
		return &bizrepo.ScheduledTaskExecutionRecordClaimResp{}, err
	}
	row, err := r.Get(ctx, &bizrepo.ScheduledTaskExecutionRecordGetReq{ID: &req.ID})
	if err != nil {
		return nil, err
	}
	return &bizrepo.ScheduledTaskExecutionRecordClaimResp{Row: row, Claimed: true}, nil
}

func (r *ScheduledTaskExecutionRecordRepo) MarkFinished(
	ctx context.Context,
	req *bizrepo.ScheduledTaskExecutionRecordMarkFinishedReq,
) (*bizrepo.ScheduledTaskExecutionRecordMarkFinishedResp, error) {
	row, err := r.getClient(ctx).ScheduledTaskExecutionRecord.UpdateOneID(req.ID).
		Where(
			scheduledtaskexecutionrecord.StatusEQ(scheduledtaskexecutionrecord.Status(schedulerenum.TaskExecutionStatusRunning)),
			scheduledtaskexecutionrecord.WorkerIDEQ(req.WorkerID),
			scheduledtaskexecutionrecord.AttemptEQ(req.Attempt),
		).
		SetStatus(scheduledtaskexecutionrecord.Status(req.Status)).
		SetFinishedAt(req.FinishedAt).
		SetDurationMs(req.Duration.Milliseconds()).
		SetLastError(req.LastError).
		Save(ctx)
	if gen.IsNotFound(err) {
		current, getErr := r.Get(ctx, &bizrepo.ScheduledTaskExecutionRecordGetReq{ID: &req.ID})
		if getErr != nil {
			return nil, getErr
		}
		return &bizrepo.ScheduledTaskExecutionRecordMarkFinishedResp{Row: current}, nil
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.ScheduledTaskExecutionRecordMarkFinishedResp{Row: r.model(row), Updated: true}, nil
}

func (r *ScheduledTaskExecutionRecordRepo) MarkCanceled(ctx context.Context, id int64, finishedAt time.Time) (*model.ScheduledTaskExecutionRecord, error) {
	row, err := r.getClient(ctx).ScheduledTaskExecutionRecord.UpdateOneID(id).
		Where(scheduledtaskexecutionrecord.StatusEQ(scheduledtaskexecutionrecord.StatusRunning)).
		SetStatus(scheduledtaskexecutionrecord.Status(schedulerenum.TaskExecutionStatusCanceled)).
		SetFinishedAt(finishedAt).
		SetLastError("scheduler execution canceled").
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *ScheduledTaskExecutionRecordRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *ScheduledTaskExecutionRecordRepo) getQuery(
	query *gen.ScheduledTaskExecutionRecordQuery,
	req *bizrepo.ScheduledTaskExecutionRecordGetReq,
) *gen.ScheduledTaskExecutionRecordQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(scheduledtaskexecutionrecord.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(scheduledtaskexecutionrecord.IDIn(req.IDs...))
	}
	if req.ScheduledTaskID != nil {
		query = query.Where(scheduledtaskexecutionrecord.ScheduledTaskIDEQ(*req.ScheduledTaskID))
	}
	if req.ScheduledAt != nil {
		query = query.Where(scheduledtaskexecutionrecord.ScheduledAtEQ(*req.ScheduledAt))
	}
	if req.ScheduleKey != nil {
		query = query.Where(scheduledtaskexecutionrecord.ScheduleKey(*req.ScheduleKey))
	}
	if req.Status != nil {
		query = query.Where(scheduledtaskexecutionrecord.StatusEQ(scheduledtaskexecutionrecord.Status(*req.Status)))
	}
	if req.TriggerType != nil {
		query = query.Where(scheduledtaskexecutionrecord.TriggerTypeEQ(scheduledtaskexecutionrecord.TriggerType(*req.TriggerType)))
	}
	return query
}

func (r *ScheduledTaskExecutionRecordRepo) model(row *gen.ScheduledTaskExecutionRecord) *model.ScheduledTaskExecutionRecord {
	var duration *time.Duration
	if row.DurationMs != nil {
		duration = new(time.Duration(*row.DurationMs) * time.Millisecond)
	}
	var staleAfter *time.Duration
	if row.StaleAfterSeconds != nil {
		staleAfter = new(time.Duration(*row.StaleAfterSeconds) * time.Second)
	}
	return &model.ScheduledTaskExecutionRecord{
		ID:                   row.ID,
		ScheduledTaskID:      row.ScheduledTaskID,
		ScheduledTaskVersion: row.ScheduledTaskVersion,
		TriggerType:          schedulerenum.TaskTriggerType(row.TriggerType),
		ScheduleKey:          row.ScheduleKey,
		ScheduledAt:          row.ScheduledAt,
		StartedAt:            row.StartedAt,
		FinishedAt:           row.FinishedAt,
		Duration:             duration,
		Status:               schedulerenum.TaskExecutionStatus(row.Status),
		Attempt:              row.Attempt,
		MaxAttempts:          row.MaxAttempts,
		Timeout:              time.Duration(row.TimeoutSeconds) * time.Second,
		StaleAfter:           staleAfter,
		MisfirePolicy:        schedulerenum.TaskMisfirePolicy(row.MisfirePolicy),
		WorkerID:             row.WorkerID,
		Payload:              row.Payload,
		LastError:            row.LastError,
		TraceID:              row.TraceID,
		CreatedAt:            row.CreatedAt,
		UpdatedAt:            row.UpdatedAt,
	}
}
