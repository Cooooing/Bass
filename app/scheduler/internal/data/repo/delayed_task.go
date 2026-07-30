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
	"scheduler/internal/data/gen/delayedtask"
	schedulerenum "scheduler/internal/enum"
	"time"

	"entgo.io/ent/dialect/sql"
)

type DelayedTaskRepo struct {
	db *gen.Client
}

func NewDelayedTaskRepo(
	db *gen.Client,
) bizrepo.DelayedTaskRepo {
	return &DelayedTaskRepo{db: db}
}

func (r *DelayedTaskRepo) Get(ctx context.Context, req *bizrepo.DelayedTaskGetReq) (*model.DelayedTask, error) {
	query := r.getClient(ctx).DelayedTask.Query().Where(delayedtask.DeletedAtIsNil())
	row, err := r.getQuery(query, req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *DelayedTaskRepo) List(ctx context.Context, req *bizrepo.DelayedTaskGetReq) ([]*model.DelayedTask, error) {
	query := r.getClient(ctx).DelayedTask.Query().Where(delayedtask.DeletedAtIsNil())
	rows, err := r.getQuery(query, req).Order(delayedtask.ByID()).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.DelayedTask, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return result, nil
}

func (r *DelayedTaskRepo) Page(ctx context.Context, req *bizrepo.DelayedTaskPageReq) (*bizrepo.DelayedTaskPageResp, error) {
	page := server.PageValid(req.Page)
	query := r.getClient(ctx).DelayedTask.Query().Where(delayedtask.DeletedAtIsNil())
	query = r.getQuery(query, &req.DelayedTaskGetReq)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := query.Order(delayedtask.ByID()).Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.DelayedTask, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return &bizrepo.DelayedTaskPageResp{Rows: result, Page: &common.PageResp{Total: uint32(total), Page: page.Page, Size: page.Size}}, nil
}

func (r *DelayedTaskRepo) MapByTaskKey(ctx context.Context, taskKeys []string) (map[string]*model.DelayedTask, error) {
	if len(taskKeys) == 0 {
		return map[string]*model.DelayedTask{}, nil
	}
	rows, err := r.List(ctx, &bizrepo.DelayedTaskGetReq{TaskKeys: taskKeys})
	if err != nil {
		return nil, err
	}
	result := make(map[string]*model.DelayedTask, len(rows))
	for _, row := range rows {
		result[row.TaskKey] = row
	}
	return result, nil
}

func (r *DelayedTaskRepo) Upsert(ctx context.Context, row *model.DelayedTask) (*model.DelayedTask, error) {
	var staleAfterSeconds *int32
	if row.StaleAfter != nil {
		staleAfterSeconds = new(int32(*row.StaleAfter / time.Second))
	}
	if row.ID > 0 {
		current, err := r.getClient(ctx).DelayedTask.Query().Where(delayedtask.ID(row.ID), delayedtask.DeletedAtIsNil()).Only(ctx)
		if gen.IsNotFound(err) {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		if err != nil {
			return nil, err
		}
		updated, err := r.getClient(ctx).DelayedTask.UpdateOneID(row.ID).
			SetTaskKey(row.TaskKey).
			SetHandlerName(row.HandlerName.String()).
			SetTitle(row.Title).
			SetDescription(row.Description).
			SetEnabled(row.Enabled).
			SetTimeoutSeconds(int32(row.Timeout / time.Second)).
			SetNillableStaleAfterSeconds(staleAfterSeconds).
			SetMaxAttempts(row.MaxAttempts).
			SetMisfirePolicy(delayedtask.MisfirePolicy(row.MisfirePolicy)).
			SetVersion(current.Version + 1).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return r.model(updated), nil
	}
	created, err := r.getClient(ctx).DelayedTask.Create().
		SetTaskKey(row.TaskKey).
		SetHandlerName(row.HandlerName.String()).
		SetTitle(row.Title).
		SetDescription(row.Description).
		SetEnabled(row.Enabled).
		SetTimeoutSeconds(int32(row.Timeout / time.Second)).
		SetNillableStaleAfterSeconds(staleAfterSeconds).
		SetMaxAttempts(row.MaxAttempts).
		SetMisfirePolicy(delayedtask.MisfirePolicy(row.MisfirePolicy)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.model(created), nil
}

func (r *DelayedTaskRepo) Lock(ctx context.Context, id int64) error {
	_, err := r.getClient(ctx).DelayedTask.Query().Where(delayedtask.ID(id), delayedtask.DeletedAtIsNil(), func(s *sql.Selector) { s.ForUpdate() }).Only(ctx)
	return err
}

func (r *DelayedTaskRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *DelayedTaskRepo) getQuery(query *gen.DelayedTaskQuery, req *bizrepo.DelayedTaskGetReq) *gen.DelayedTaskQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(delayedtask.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(delayedtask.IDIn(req.IDs...))
	}
	if req.TaskKey != nil {
		query = query.Where(delayedtask.TaskKeyEQ(*req.TaskKey))
	}
	if len(req.TaskKeys) > 0 {
		query = query.Where(delayedtask.TaskKeyIn(req.TaskKeys...))
	}
	if req.HandlerName != nil {
		query = query.Where(delayedtask.HandlerNameEQ(req.HandlerName.String()))
	}
	if req.Title != nil {
		query = query.Where(delayedtask.TitleContains(*req.Title))
	}
	if req.Enabled != nil {
		query = query.Where(delayedtask.Enabled(*req.Enabled))
	}
	return query
}

func (r *DelayedTaskRepo) model(row *gen.DelayedTask) *model.DelayedTask {
	var staleAfter *time.Duration
	if row.StaleAfterSeconds != nil {
		staleAfter = new(time.Duration(*row.StaleAfterSeconds) * time.Second)
	}
	return &model.DelayedTask{
		ID:            row.ID,
		TaskKey:       row.TaskKey,
		HandlerName:   schedulerenum.TaskHandlerName(row.HandlerName),
		Title:         row.Title,
		Description:   row.Description,
		Enabled:       row.Enabled,
		Timeout:       time.Duration(row.TimeoutSeconds) * time.Second,
		StaleAfter:    staleAfter,
		MaxAttempts:   row.MaxAttempts,
		MisfirePolicy: schedulerenum.TaskMisfirePolicy(row.MisfirePolicy),
		Version:       row.Version,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}
}
