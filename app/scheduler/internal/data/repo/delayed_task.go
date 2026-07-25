package repo

import (
	"common/proto/gen/common"
	"context"
	"scheduler/internal/biz/model"
	bizrepo "scheduler/internal/biz/repo"
	"scheduler/internal/data/gen"
	"scheduler/internal/data/gen/delayedtask"
	schedulerenum "scheduler/internal/enum"
	"time"

	"common/pkg/server"
	utilent "common/pkg/util/ent"
)

var _ bizrepo.DelayedTaskRepo = (*DelayedTaskRepo)(nil)

type DelayedTaskRepo struct{ db *gen.Client }

func NewDelayedTaskRepo(
	db *gen.Client,
) bizrepo.DelayedTaskRepo {
	return &DelayedTaskRepo{
		db: db,
	}
}

func (r *DelayedTaskRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *DelayedTaskRepo) Register(ctx context.Context, row *model.DelayedTask) (*model.DelayedTask, error) {
	if existing, err := r.Get(ctx, &bizrepo.DelayedTaskGetReq{
		IdempotencyKey: &row.IdempotencyKey,
	}); err != nil || existing != nil {
		return existing, err
	}
	created, err := r.getClient(ctx).DelayedTask.Create().
		SetIdempotencyKey(row.IdempotencyKey).
		SetTaskName(row.TaskName).
		SetPayload(row.Payload).
		SetExecuteAt(row.ExecuteAt).
		SetNextRunAt(row.NextRunAt).
		SetStatus(delayedtask.Status(row.Status)).
		SetAttempt(row.Attempt).
		SetMaxAttempts(row.MaxAttempts).
		SetTimeoutSeconds(row.TimeoutSeconds).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return delayedTaskToModel(created), nil
}

func (r *DelayedTaskRepo) Get(ctx context.Context, req *bizrepo.DelayedTaskGetReq) (*model.DelayedTask, error) {
	query := r.getClient(ctx).DelayedTask.Query()
	query = delayedTaskQuery(query, req)
	row, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return delayedTaskToModel(row), nil
}

func (r *DelayedTaskRepo) Page(ctx context.Context, req *bizrepo.DelayedTaskPageReq) (*bizrepo.DelayedTaskPageResp, error) {
	page := server.PageValid(req.Page)
	query := r.getClient(ctx).DelayedTask.Query()
	query = delayedTaskQuery(query, &req.DelayedTaskGetReq)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := query.Order(gen.Desc(delayedtask.FieldCreatedAt)).Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.DelayedTask, 0, len(rows))
	for _, row := range rows {
		result = append(result, delayedTaskToModel(row))
	}
	return &bizrepo.DelayedTaskPageResp{
		Rows: result,
		Page: &common.PageResp{
			Total: uint32(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *DelayedTaskRepo) Cancel(ctx context.Context, req *bizrepo.DelayedTaskGetReq) (bool, error) {
	row, err := r.Get(ctx, req)
	if err != nil || row == nil {
		return false, err
	}
	if row.Status != schedulerenum.DelayedTaskStatusPending {
		return false, nil
	}
	affected, err := r.getClient(ctx).DelayedTask.Update().Where(delayedtask.ID(row.ID), delayedtask.StatusEQ(delayedtask.Status(schedulerenum.DelayedTaskStatusPending))).SetStatus(delayedtask.Status(schedulerenum.DelayedTaskStatusCancelled)).SetFinishedAt(time.Now()).Save(ctx)
	return affected > 0, err
}

func (r *DelayedTaskRepo) ListDue(ctx context.Context, now time.Time, limit int) ([]*model.DelayedTask, error) {
	if limit <= 0 {
		limit = 100
	}
	rows, err := r.getClient(ctx).DelayedTask.Query().Where(delayedtask.Or(
		delayedtask.And(
			delayedtask.StatusEQ(delayedtask.Status(schedulerenum.DelayedTaskStatusPending)),
			delayedtask.NextRunAtLTE(now),
		),
		delayedtask.And(
			delayedtask.StatusEQ(delayedtask.Status(schedulerenum.DelayedTaskStatusRunning)),
			delayedtask.LockExpiresAtNotNil(),
			delayedtask.LockExpiresAtLTE(now),
		),
	)).Order(gen.Asc(delayedtask.FieldNextRunAt)).Limit(limit).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.DelayedTask, 0, len(rows))
	for _, row := range rows {
		result = append(result, delayedTaskToModel(row))
	}
	return result, nil
}

func (r *DelayedTaskRepo) MarkRunning(ctx context.Context, id int64, workerID string, lockExpiresAt time.Time) (bool, *model.DelayedTask, error) {
	now := time.Now()
	affected, err := r.getClient(ctx).DelayedTask.Update().Where(
		delayedtask.ID(id),
		delayedtask.Or(
			delayedtask.StatusEQ(delayedtask.Status(schedulerenum.DelayedTaskStatusPending)),
			delayedtask.And(
				delayedtask.StatusEQ(delayedtask.Status(schedulerenum.DelayedTaskStatusRunning)),
				delayedtask.LockExpiresAtNotNil(),
				delayedtask.LockExpiresAtLTE(now),
			),
		),
	).SetStatus(delayedtask.Status(schedulerenum.DelayedTaskStatusRunning)).SetLockedBy(workerID).SetLockExpiresAt(lockExpiresAt).SetStartedAt(now).Save(ctx)
	if err != nil || affected == 0 {
		return false, nil, err
	}
	row, err := r.Get(ctx, &bizrepo.DelayedTaskGetReq{
		ID: &id,
	})
	return true, row, err
}

func (r *DelayedTaskRepo) MarkSuccess(ctx context.Context, id int64, finishedAt time.Time) (*model.DelayedTask, error) {
	row, err := r.getClient(ctx).DelayedTask.UpdateOneID(id).SetStatus(delayedtask.Status(schedulerenum.DelayedTaskStatusSuccess)).SetFinishedAt(finishedAt).ClearLockExpiresAt().SetLockedBy("").SetLastError("").Save(ctx)
	if err != nil {
		return nil, err
	}
	return delayedTaskToModel(row), nil
}

func (r *DelayedTaskRepo) MarkFailed(ctx context.Context, id int64, attempt int32, final bool, nextRunAt time.Time, lastError string) (*model.DelayedTask, error) {
	update := r.getClient(ctx).DelayedTask.UpdateOneID(id).SetAttempt(attempt).SetLastError(lastError).ClearLockExpiresAt().SetLockedBy("")
	if final {
		update.SetStatus(delayedtask.Status(schedulerenum.DelayedTaskStatusFailed)).SetFinishedAt(time.Now())
	} else {
		update.SetStatus(delayedtask.Status(schedulerenum.DelayedTaskStatusPending)).SetNextRunAt(nextRunAt)
	}
	row, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return delayedTaskToModel(row), nil
}

func delayedTaskQuery(query *gen.DelayedTaskQuery, req *bizrepo.DelayedTaskGetReq) *gen.DelayedTaskQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(delayedtask.ID(*req.ID))
	}
	if req.IdempotencyKey != nil {
		query = query.Where(delayedtask.IdempotencyKey(*req.IdempotencyKey))
	}
	if req.TaskName != nil {
		query = query.Where(delayedtask.TaskName(*req.TaskName))
	}
	if req.Status != nil {
		query = query.Where(delayedtask.StatusEQ(delayedtask.Status(*req.Status)))
	}
	return query
}

func delayedTaskToModel(row *gen.DelayedTask) *model.DelayedTask {
	if row == nil {
		return nil
	}
	return &model.DelayedTask{
		ID:             row.ID,
		IdempotencyKey: row.IdempotencyKey,
		TaskName:       row.TaskName,
		Payload:        row.Payload,
		ExecuteAt:      row.ExecuteAt,
		NextRunAt:      row.NextRunAt,
		Status:         schedulerenum.DelayedTaskStatus(row.Status),
		Attempt:        row.Attempt,
		MaxAttempts:    row.MaxAttempts,
		TimeoutSeconds: row.TimeoutSeconds,
		LockedBy:       row.LockedBy,
		LockExpiresAt:  row.LockExpiresAt,
		StartedAt:      row.StartedAt,
		FinishedAt:     row.FinishedAt,
		LastError:      row.LastError,
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
	}
}
