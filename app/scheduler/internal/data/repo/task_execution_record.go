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

func (r *TaskExecutionRecordRepo) Get(ctx context.Context, req *bizrepo.TaskExecutionRecordGetReq) (*bizrepo.TaskExecutionRecordGetResponse, error) {
	query := r.getClient(ctx).TaskExecutionRecord.Query()
	query = r.getQuery(query, req)
	row, err := query.Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.TaskExecutionRecordGetResponse{Row: r.model(row)}, nil
}

func (r *TaskExecutionRecordRepo) List(ctx context.Context, req *bizrepo.TaskExecutionRecordGetReq) (*bizrepo.TaskExecutionRecordListResponse, error) {
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
	return &bizrepo.TaskExecutionRecordListResponse{Rows: result}, nil
}

func (r *TaskExecutionRecordRepo) Map(ctx context.Context, req *bizrepo.TaskExecutionRecordGetReq) (*bizrepo.TaskExecutionRecordMapResponse, error) {
	resp, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.TaskExecutionRecord, len(resp.Rows))
	for _, row := range resp.Rows {
		result[row.ID] = row
	}
	return &bizrepo.TaskExecutionRecordMapResponse{Rows: result}, nil
}

func (r *TaskExecutionRecordRepo) Count(ctx context.Context, req *bizrepo.TaskExecutionRecordGetReq) (*bizrepo.TaskExecutionRecordCountResponse, error) {
	query := r.getClient(ctx).TaskExecutionRecord.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.TaskExecutionRecordCountResponse{Count: count}, nil
}

func (r *TaskExecutionRecordRepo) Page(ctx context.Context, req *bizrepo.TaskExecutionRecordPageReq) (*bizrepo.TaskExecutionRecordPageResponse, error) {
	page := server.PageValid(req.Page)
	query := r.getClient(ctx).TaskExecutionRecord.Query()
	query = r.getQuery(query, &req.TaskExecutionRecordGetReq)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := query.Order(taskexecutionrecord.ByScheduledAt(entsql.OrderDesc())).Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.TaskExecutionRecord, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return &bizrepo.TaskExecutionRecordPageResponse{Rows: result, Page: &common.PageResponse{Total: uint32(total), Page: page.Page, Size: page.Size}}, nil
}

func (r *TaskExecutionRecordRepo) ExistsPeriod(ctx context.Context, req *bizrepo.TaskExecutionRecordExistsPeriodReq) (*bizrepo.TaskExecutionRecordExistsPeriodResponse, error) {
	exists, err := r.getClient(ctx).TaskExecutionRecord.Query().
		Where(
			taskexecutionrecord.TaskIDEQ(req.TaskID),
			taskexecutionrecord.ScheduledAtEQ(req.ScheduledAt),
		).
		Exist(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.TaskExecutionRecordExistsPeriodResponse{Exists: exists}, nil
}

func (r *TaskExecutionRecordRepo) Create(ctx context.Context, req *bizrepo.TaskExecutionRecordCreateReq) (*bizrepo.TaskExecutionRecordCreateResponse, error) {
	record := req.Record
	created, err := r.getClient(ctx).TaskExecutionRecord.Create().
		SetTaskID(record.TaskID).
		SetScheduledAt(record.ScheduledAt).
		SetNillableStartedAt(record.StartedAt).
		SetNillableFinishedAt(record.FinishedAt).
		SetNillableDurationMs(record.DurationMS).
		SetStatus(taskexecutionrecord.Status(req.Status)).
		SetTriggerType(taskexecutionrecord.TriggerType(record.TriggerType)).
		SetTaskVersion(record.TaskVersion).
		SetWorkerID(r.workerID).
		SetPayload(record.Payload).
		SetLastError(record.LastError).
		SetTraceID(record.TraceID).
		Save(ctx)
	if gen.IsConstraintError(err) {
		if strings.Contains(err.Error(), "scheduler_task_execution_records_task_period_unique") {
			return &bizrepo.TaskExecutionRecordCreateResponse{Conflict: true}, nil
		}
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.TaskExecutionRecordCreateResponse{Row: r.model(created), Created: true}, nil
}

func (r *TaskExecutionRecordRepo) HasUnexpiredRunning(ctx context.Context, req *bizrepo.TaskExecutionRecordHasUnexpiredRunningReq) (*bizrepo.TaskExecutionRecordHasUnexpiredRunningResponse, error) {
	exists, err := r.getClient(ctx).TaskExecutionRecord.Query().
		Where(
			taskexecutionrecord.TaskIDEQ(req.TaskID),
			taskexecutionrecord.StatusEQ(taskexecutionrecord.StatusRunning),
			taskexecutionrecord.StartedAtGTE(req.StartedAfter),
		).
		Exist(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.TaskExecutionRecordHasUnexpiredRunningResponse{Exists: exists}, nil
}

func (r *TaskExecutionRecordRepo) MarkUnknown(ctx context.Context, req *bizrepo.TaskExecutionRecordMarkUnknownReq) (*bizrepo.TaskExecutionRecordMarkUnknownResponse, error) {
	if len(req.IDs) == 0 {
		return &bizrepo.TaskExecutionRecordMarkUnknownResponse{}, nil
	}
	rows, err := r.getClient(ctx).TaskExecutionRecord.Query().
		Where(taskexecutionrecord.IDIn(req.IDs...), taskexecutionrecord.StatusEQ(taskexecutionrecord.StatusRunning)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.TaskExecutionRecord, 0, len(rows))
	for _, row := range rows {
		duration := int64(0)
		if row.StartedAt != nil {
			duration = req.FinishedAt.Sub(*row.StartedAt).Milliseconds()
			if duration < 0 {
				duration = 0
			}
		}
		updated, err := r.getClient(ctx).TaskExecutionRecord.UpdateOneID(row.ID).
			Where(taskexecutionrecord.StatusEQ(taskexecutionrecord.StatusRunning)).
			SetStatus(taskexecutionrecord.StatusUnknown).
			SetFinishedAt(req.FinishedAt).
			SetDurationMs(duration).
			SetLastError(req.LastError).
			Save(ctx)
		if gen.IsNotFound(err) {
			continue
		}
		if err != nil {
			return nil, err
		}
		result = append(result, r.model(updated))
	}
	return &bizrepo.TaskExecutionRecordMarkUnknownResponse{Rows: result}, nil
}

func (r *TaskExecutionRecordRepo) MarkFinished(ctx context.Context, req *bizrepo.TaskExecutionRecordMarkFinishedReq) (*bizrepo.TaskExecutionRecordMarkFinishedResponse, error) {
	row, err := r.getClient(ctx).TaskExecutionRecord.UpdateOneID(req.ID).
		Where(taskexecutionrecord.StatusEQ(taskexecutionrecord.StatusRunning)).
		SetStatus(taskexecutionrecord.Status(req.Status)).
		SetFinishedAt(req.FinishedAt).
		SetDurationMs(req.DurationMS).
		SetLastError(req.LastError).
		Save(ctx)
	if gen.IsNotFound(err) {
		current, getErr := r.Get(ctx, &bizrepo.TaskExecutionRecordGetReq{ID: &req.ID})
		if getErr != nil {
			return nil, getErr
		}
		return &bizrepo.TaskExecutionRecordMarkFinishedResponse{Row: current.Row}, nil
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.TaskExecutionRecordMarkFinishedResponse{Row: r.model(row), Updated: true}, nil
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
