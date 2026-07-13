package repo

import (
	"common/pkg/apperror"
	"common/pkg/server"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
	"os"
	"scheduler/internal/biz/model"
	bizrepo "scheduler/internal/biz/repo"
	"scheduler/internal/data/gen"
	"scheduler/internal/data/gen/taskexecutionrecord"
	schedulerenum "scheduler/internal/enum"
	"strings"
	"time"

	utilent "common/pkg/util/ent"
	entsql "entgo.io/ent/dialect/sql"
)

var _ bizrepo.TaskExecutionRecordRepo = (*TaskExecutionRecordRepo)(nil)

type TaskExecutionRecordRepo struct {
	db       *gen.Client
	workerID string
}

func NewTaskExecutionRecordRepo(db *gen.Client) bizrepo.TaskExecutionRecordRepo {
	hostname, _ := os.Hostname()
	return &TaskExecutionRecordRepo{db: db, workerID: hostname}
}

func (r *TaskExecutionRecordRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *TaskExecutionRecordRepo) Get(ctx context.Context, req *bizrepo.TaskExecutionRecordGetReq) (*model.TaskExecutionRecord, error) {
	query := r.getClient(ctx).TaskExecutionRecord.Query()
	query = r.getQuery(query, req)
	row, err := query.Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *TaskExecutionRecordRepo) List(ctx context.Context, req *bizrepo.TaskExecutionRecordGetReq) ([]*model.TaskExecutionRecord, error) {
	query := r.getClient(ctx).TaskExecutionRecord.Query()
	query = r.getQuery(query, req)
	rows, err := query.Order(taskexecutionrecord.ByScheduledAt(entsql.OrderDesc())).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.TaskExecutionRecord, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return result, nil
}

func (r *TaskExecutionRecordRepo) Map(ctx context.Context, req *bizrepo.TaskExecutionRecordGetReq) (map[int64]*model.TaskExecutionRecord, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.TaskExecutionRecord, len(rows))
	for _, row := range rows {
		result[row.ID] = row
	}
	return result, nil
}

func (r *TaskExecutionRecordRepo) Count(ctx context.Context, req *bizrepo.TaskExecutionRecordGetReq) (int, error) {
	query := r.getClient(ctx).TaskExecutionRecord.Query()
	query = r.getQuery(query, req)
	return query.Count(ctx)
}

func (r *TaskExecutionRecordRepo) Page(ctx context.Context, page *common.PageRequest, req *bizrepo.TaskExecutionRecordGetReq) ([]*model.TaskExecutionRecord, *common.PageReply, error) {
	page = server.PageValid(page)
	query := r.getClient(ctx).TaskExecutionRecord.Query()
	query = r.getQuery(query, req)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, nil, err
	}
	rows, err := query.Order(taskexecutionrecord.ByScheduledAt(entsql.OrderDesc())).Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, nil, err
	}
	result := make([]*model.TaskExecutionRecord, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return result, &common.PageReply{Total: uint32(total), Page: page.Page, Size: page.Size}, nil
}

func (r *TaskExecutionRecordRepo) ExistsPeriod(ctx context.Context, taskID int64, scheduledAt time.Time) (bool, error) {
	return r.getClient(ctx).TaskExecutionRecord.Query().
		Where(
			taskexecutionrecord.TaskIDEQ(taskID),
			taskexecutionrecord.ScheduledAtEQ(scheduledAt),
		).
		Exist(ctx)
}

func (r *TaskExecutionRecordRepo) Create(ctx context.Context, record *model.TaskExecutionRecord, statusValue schedulerenum.TaskExecutionStatus) (*model.TaskExecutionRecord, bool, bool, error) {
	created, err := r.getClient(ctx).TaskExecutionRecord.Create().
		SetTaskID(record.TaskID).
		SetScheduledAt(record.ScheduledAt).
		SetNillableStartedAt(record.StartedAt).
		SetNillableFinishedAt(record.FinishedAt).
		SetNillableDurationMs(record.DurationMS).
		SetStatus(taskexecutionrecord.Status(statusValue)).
		SetTriggerType(taskexecutionrecord.TriggerType(record.TriggerType)).
		SetTaskVersion(record.TaskVersion).
		SetWorkerID(r.workerID).
		SetPayload(record.Payload).
		SetLastError(record.LastError).
		SetTraceID(record.TraceID).
		Save(ctx)
	if gen.IsConstraintError(err) {
		if strings.Contains(err.Error(), "scheduler_task_execution_records_task_period_unique") {
			return nil, false, true, nil
		}
		return nil, false, false, err
	}
	if err != nil {
		return nil, false, false, err
	}
	return r.model(created), true, false, nil
}

func (r *TaskExecutionRecordRepo) HasUnexpiredRunning(ctx context.Context, taskID int64, startedAfter time.Time) (bool, error) {
	return r.getClient(ctx).TaskExecutionRecord.Query().
		Where(
			taskexecutionrecord.TaskIDEQ(taskID),
			taskexecutionrecord.StatusEQ(taskexecutionrecord.StatusRunning),
			taskexecutionrecord.StartedAtGTE(startedAfter),
		).
		Exist(ctx)
}

func (r *TaskExecutionRecordRepo) MarkUnknown(ctx context.Context, ids []int64, finishedAt time.Time, lastError string) ([]*model.TaskExecutionRecord, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	rows, err := r.getClient(ctx).TaskExecutionRecord.Query().
		Where(taskexecutionrecord.IDIn(ids...), taskexecutionrecord.StatusEQ(taskexecutionrecord.StatusRunning)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.TaskExecutionRecord, 0, len(rows))
	for _, row := range rows {
		duration := int64(0)
		if row.StartedAt != nil {
			duration = finishedAt.Sub(*row.StartedAt).Milliseconds()
			if duration < 0 {
				duration = 0
			}
		}
		updated, err := r.getClient(ctx).TaskExecutionRecord.UpdateOneID(row.ID).
			Where(taskexecutionrecord.StatusEQ(taskexecutionrecord.StatusRunning)).
			SetStatus(taskexecutionrecord.StatusUnknown).
			SetFinishedAt(finishedAt).
			SetDurationMs(duration).
			SetLastError(lastError).
			Save(ctx)
		if gen.IsNotFound(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		result = append(result, r.model(updated))
	}
	return result, nil
}

func (r *TaskExecutionRecordRepo) MarkFinished(ctx context.Context, id int64, statusValue schedulerenum.TaskExecutionStatus, finishedAt time.Time, durationMS int64, lastError string) (*model.TaskExecutionRecord, bool, error) {
	row, err := r.getClient(ctx).TaskExecutionRecord.UpdateOneID(id).
		Where(taskexecutionrecord.StatusEQ(taskexecutionrecord.StatusRunning)).
		SetStatus(taskexecutionrecord.Status(statusValue)).
		SetFinishedAt(finishedAt).
		SetDurationMs(durationMS).
		SetLastError(lastError).
		Save(ctx)
	if gen.IsNotFound(err) {
		current, getErr := r.Get(ctx, &bizrepo.TaskExecutionRecordGetReq{ID: &id})
		return current, false, getErr
	}
	if err != nil {
		return nil, false, err
	}
	return r.model(row), true, nil
}

func (r *TaskExecutionRecordRepo) getQuery(query *gen.TaskExecutionRecordQuery, req *bizrepo.TaskExecutionRecordGetReq) *gen.TaskExecutionRecordQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(taskexecutionrecord.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(taskexecutionrecord.IDIn(req.IDs...))
	}
	if req.TaskID != nil {
		query = query.Where(taskexecutionrecord.TaskIDEQ(*req.TaskID))
	}
	if req.ScheduledAt != nil {
		query = query.Where(taskexecutionrecord.ScheduledAtEQ(*req.ScheduledAt))
	}
	if req.Status != nil {
		query = query.Where(taskexecutionrecord.StatusEQ(taskexecutionrecord.Status(*req.Status)))
	}
	if req.TriggerType != nil {
		query = query.Where(taskexecutionrecord.TriggerTypeEQ(taskexecutionrecord.TriggerType(*req.TriggerType)))
	}
	return query
}

func (r *TaskExecutionRecordRepo) model(row *gen.TaskExecutionRecord) *model.TaskExecutionRecord {
	return &model.TaskExecutionRecord{
		ID:          row.ID,
		TaskID:      row.TaskID,
		ScheduledAt: row.ScheduledAt,
		StartedAt:   row.StartedAt,
		FinishedAt:  row.FinishedAt,
		DurationMS:  row.DurationMs,
		Status:      schedulerenum.TaskExecutionStatus(row.Status),
		TriggerType: schedulerenum.TaskTriggerType(row.TriggerType),
		TaskVersion: row.TaskVersion,
		WorkerID:    row.WorkerID,
		Payload:     row.Payload,
		LastError:   row.LastError,
		TraceID:     row.TraceID,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}
