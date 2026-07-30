package repo

import (
	"common/pkg/server"
	utilent "common/pkg/util/ent"
	"common/proto/gen/common"
	"context"
	"scheduler/internal/biz/model"
	bizrepo "scheduler/internal/biz/repo"
	"scheduler/internal/data/gen"
	"scheduler/internal/data/gen/delayedtaskexecutionrecord"
	schedulerenum "scheduler/internal/enum"
	"strings"
	"time"
)

type DelayedTaskExecutionRecordRepo struct {
	db *gen.Client
}

func NewDelayedTaskExecutionRecordRepo(
	db *gen.Client,
) bizrepo.DelayedTaskExecutionRecordRepo {
	return &DelayedTaskExecutionRecordRepo{db: db}
}

func (r *DelayedTaskExecutionRecordRepo) Get(ctx context.Context, req *bizrepo.DelayedTaskExecutionRecordGetReq) (*model.DelayedTaskExecutionRecord, error) {
	row, err := r.getQuery(r.getClient(ctx).DelayedTaskExecutionRecord.Query(), req).First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *DelayedTaskExecutionRecordRepo) List(ctx context.Context, req *bizrepo.DelayedTaskExecutionRecordGetReq) ([]*model.DelayedTaskExecutionRecord, error) {
	rows, err := r.getQuery(r.getClient(ctx).DelayedTaskExecutionRecord.Query(), req).Order(delayedtaskexecutionrecord.ByScheduledAt()).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.DelayedTaskExecutionRecord, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return result, nil
}

func (r *DelayedTaskExecutionRecordRepo) Page(ctx context.Context, req *bizrepo.DelayedTaskExecutionRecordPageReq) (*bizrepo.DelayedTaskExecutionRecordPageResp, error) {
	page := server.PageValid(req.Page)
	query := r.getQuery(r.getClient(ctx).DelayedTaskExecutionRecord.Query(), &req.DelayedTaskExecutionRecordGetReq)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := query.Order(gen.Desc(delayedtaskexecutionrecord.FieldCreatedAt)).Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.DelayedTaskExecutionRecord, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return &bizrepo.DelayedTaskExecutionRecordPageResp{Rows: result, Page: &common.PageResp{Total: uint32(total), Page: page.Page, Size: page.Size}}, nil
}

func (r *DelayedTaskExecutionRecordRepo) CreatePending(ctx context.Context, record *model.DelayedTaskExecutionRecord) (*bizrepo.DelayedTaskExecutionRecordCreateResp, error) {
	var staleAfterSeconds *int32
	if record.StaleAfter != nil {
		staleAfterSeconds = new(int32(*record.StaleAfter / time.Second))
	}
	created, err := r.getClient(ctx).DelayedTaskExecutionRecord.Create().
		SetDelayedTaskID(record.DelayedTaskID).
		SetDelayedTaskVersion(record.DelayedTaskVersion).
		SetIdempotencyKey(record.IdempotencyKey).
		SetTriggerType(delayedtaskexecutionrecord.TriggerType(record.TriggerType)).
		SetScheduleKey(record.ScheduleKey).
		SetScheduledAt(record.ScheduledAt).
		SetStatus(delayedtaskexecutionrecord.Status(record.Status)).
		SetAttempt(record.Attempt).
		SetMaxAttempts(record.MaxAttempts).
		SetTimeoutSeconds(int32(record.Timeout / time.Second)).
		SetNillableStaleAfterSeconds(staleAfterSeconds).
		SetMisfirePolicy(delayedtaskexecutionrecord.MisfirePolicy(record.MisfirePolicy)).
		SetWorkerID(record.WorkerID).
		SetPayload(record.Payload).
		SetLastError(record.LastError).
		SetTraceID(record.TraceID).
		Save(ctx)
	if gen.IsConstraintError(err) && strings.Contains(err.Error(), "scheduler_delayed_task_execution_records_idempotency_unique") {
		existing, getErr := r.Get(ctx, &bizrepo.DelayedTaskExecutionRecordGetReq{IdempotencyKey: &record.IdempotencyKey})
		return &bizrepo.DelayedTaskExecutionRecordCreateResp{Row: existing, Conflict: true}, getErr
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.DelayedTaskExecutionRecordCreateResp{Row: r.model(created), Created: true}, nil
}

func (r *DelayedTaskExecutionRecordRepo) Claim(ctx context.Context, req *bizrepo.DelayedTaskExecutionRecordClaimReq) (*bizrepo.DelayedTaskExecutionRecordClaimResp, error) {
	affected, err := r.getClient(ctx).DelayedTaskExecutionRecord.Update().
		Where(
			delayedtaskexecutionrecord.ID(req.ID),
			delayedtaskexecutionrecord.StatusIn(
				delayedtaskexecutionrecord.Status(schedulerenum.TaskExecutionStatusPending),
				delayedtaskexecutionrecord.Status(schedulerenum.TaskExecutionStatusRetryPending),
				delayedtaskexecutionrecord.Status(schedulerenum.TaskExecutionStatusRunning),
			),
		).
		SetStatus(delayedtaskexecutionrecord.Status(schedulerenum.TaskExecutionStatusRunning)).
		SetStartedAt(req.StartedAt).
		ClearFinishedAt().
		ClearDurationMs().
		SetWorkerID(req.WorkerID).
		SetLastError("").
		AddAttempt(1).
		Save(ctx)
	if err != nil || affected == 0 {
		return &bizrepo.DelayedTaskExecutionRecordClaimResp{}, err
	}
	row, err := r.Get(ctx, &bizrepo.DelayedTaskExecutionRecordGetReq{ID: &req.ID})
	if err != nil {
		return nil, err
	}
	return &bizrepo.DelayedTaskExecutionRecordClaimResp{Row: row, Claimed: true}, nil
}

func (r *DelayedTaskExecutionRecordRepo) MarkFinished(ctx context.Context, req *bizrepo.DelayedTaskExecutionRecordMarkFinishedReq) (*model.DelayedTaskExecutionRecord, error) {
	row, err := r.getClient(ctx).DelayedTaskExecutionRecord.UpdateOneID(req.ID).
		Where(
			delayedtaskexecutionrecord.StatusEQ(delayedtaskexecutionrecord.Status(schedulerenum.TaskExecutionStatusRunning)),
			delayedtaskexecutionrecord.WorkerIDEQ(req.WorkerID),
			delayedtaskexecutionrecord.AttemptEQ(req.Attempt),
		).
		SetStatus(delayedtaskexecutionrecord.Status(req.Status)).
		SetFinishedAt(req.FinishedAt).
		SetDurationMs(req.Duration.Milliseconds()).
		SetLastError(req.LastError).
		Save(ctx)
	if gen.IsNotFound(err) {
		return r.Get(ctx, &bizrepo.DelayedTaskExecutionRecordGetReq{ID: &req.ID})
	}
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *DelayedTaskExecutionRecordRepo) MarkCanceled(
	ctx context.Context,
	req *bizrepo.DelayedTaskExecutionRecordGetReq,
	finishedAt time.Time,
) (*model.DelayedTaskExecutionRecord, error) {
	row, err := r.Get(ctx, req)
	if err != nil || row == nil {
		return row, err
	}
	updated, err := r.getClient(ctx).DelayedTaskExecutionRecord.UpdateOneID(row.ID).
		Where(delayedtaskexecutionrecord.StatusIn(
			delayedtaskexecutionrecord.Status(schedulerenum.TaskExecutionStatusPending),
			delayedtaskexecutionrecord.Status(schedulerenum.TaskExecutionStatusRunning),
			delayedtaskexecutionrecord.Status(schedulerenum.TaskExecutionStatusRetryPending),
		)).
		SetStatus(delayedtaskexecutionrecord.Status(schedulerenum.TaskExecutionStatusCanceled)).
		SetFinishedAt(finishedAt).
		SetLastError("scheduler delayed execution canceled").
		Save(ctx)
	if gen.IsNotFound(err) {
		return row, nil
	}
	if err != nil {
		return nil, err
	}
	return r.model(updated), nil
}

func (r *DelayedTaskExecutionRecordRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *DelayedTaskExecutionRecordRepo) getQuery(query *gen.DelayedTaskExecutionRecordQuery, req *bizrepo.DelayedTaskExecutionRecordGetReq) *gen.DelayedTaskExecutionRecordQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(delayedtaskexecutionrecord.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(delayedtaskexecutionrecord.IDIn(req.IDs...))
	}
	if req.DelayedTaskID != nil {
		query = query.Where(delayedtaskexecutionrecord.DelayedTaskIDEQ(*req.DelayedTaskID))
	}
	if req.IdempotencyKey != nil {
		query = query.Where(delayedtaskexecutionrecord.IdempotencyKey(*req.IdempotencyKey))
	}
	if req.Status != nil {
		query = query.Where(delayedtaskexecutionrecord.StatusEQ(delayedtaskexecutionrecord.Status(*req.Status)))
	}
	if req.TriggerType != nil {
		query = query.Where(delayedtaskexecutionrecord.TriggerTypeEQ(delayedtaskexecutionrecord.TriggerType(*req.TriggerType)))
	}
	return query
}

func (r *DelayedTaskExecutionRecordRepo) model(row *gen.DelayedTaskExecutionRecord) *model.DelayedTaskExecutionRecord {
	var duration *time.Duration
	if row.DurationMs != nil {
		duration = new(time.Duration(*row.DurationMs) * time.Millisecond)
	}
	var staleAfter *time.Duration
	if row.StaleAfterSeconds != nil {
		staleAfter = new(time.Duration(*row.StaleAfterSeconds) * time.Second)
	}
	return &model.DelayedTaskExecutionRecord{
		ID:                 row.ID,
		DelayedTaskID:      row.DelayedTaskID,
		DelayedTaskVersion: row.DelayedTaskVersion,
		IdempotencyKey:     row.IdempotencyKey,
		TriggerType:        schedulerenum.TaskTriggerType(row.TriggerType),
		ScheduleKey:        row.ScheduleKey,
		ScheduledAt:        row.ScheduledAt,
		StartedAt:          row.StartedAt,
		FinishedAt:         row.FinishedAt,
		Duration:           duration,
		Status:             schedulerenum.TaskExecutionStatus(row.Status),
		Attempt:            row.Attempt,
		MaxAttempts:        row.MaxAttempts,
		Timeout:            time.Duration(row.TimeoutSeconds) * time.Second,
		StaleAfter:         staleAfter,
		MisfirePolicy:      schedulerenum.TaskMisfirePolicy(row.MisfirePolicy),
		WorkerID:           row.WorkerID,
		Payload:            row.Payload,
		LastError:          row.LastError,
		TraceID:            row.TraceID,
		CreatedAt:          row.CreatedAt,
		UpdatedAt:          row.UpdatedAt,
	}
}
