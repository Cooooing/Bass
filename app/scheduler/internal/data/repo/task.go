package repo

import (
	"common/pkg/apperror"
	"common/pkg/server"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
	"scheduler/internal/biz/model"
	bizrepo "scheduler/internal/biz/repo"
	"scheduler/internal/data/gen"
	"scheduler/internal/data/gen/task"

	utilent "common/pkg/util/ent"
	"entgo.io/ent/dialect/sql"
)

var _ bizrepo.TaskRepo = (*TaskRepo)(nil)

type TaskRepo struct {
	db *gen.Client
}

func NewTaskRepo(db *gen.Client) bizrepo.TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *TaskRepo) Get(ctx context.Context, req *bizrepo.TaskGetReq) (*bizrepo.TaskGetResponse, error) {
	query := r.getClient(ctx).Task.Query().Where(task.DeletedAtIsNil())
	query = r.getQuery(query, req)
	row, err := query.Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.TaskGetResponse{Row: r.model(row)}, nil
}

func (r *TaskRepo) List(ctx context.Context, req *bizrepo.TaskGetReq) (*bizrepo.TaskListResponse, error) {
	query := r.getClient(ctx).Task.Query().Where(task.DeletedAtIsNil())
	query = r.getQuery(query, req)
	rows, err := query.Order(task.ByID()).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Task, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return &bizrepo.TaskListResponse{Rows: result}, nil
}

func (r *TaskRepo) Map(ctx context.Context, req *bizrepo.TaskGetReq) (*bizrepo.TaskMapResponse, error) {
	resp, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.Task, len(resp.Rows))
	for _, row := range resp.Rows {
		result[row.ID] = row
	}
	return &bizrepo.TaskMapResponse{Rows: result}, nil
}

func (r *TaskRepo) Count(ctx context.Context, req *bizrepo.TaskGetReq) (*bizrepo.TaskCountResponse, error) {
	query := r.getClient(ctx).Task.Query().Where(task.DeletedAtIsNil())
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.TaskCountResponse{Count: count}, nil
}

func (r *TaskRepo) Page(ctx context.Context, req *bizrepo.TaskPageReq) (*bizrepo.TaskPageResponse, error) {
	page := server.PageValid(req.Page)
	query := r.getClient(ctx).Task.Query().Where(task.DeletedAtIsNil())
	query = r.getQuery(query, &req.TaskGetReq)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := query.Order(task.ByID()).Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Task, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return &bizrepo.TaskPageResponse{Rows: result, Page: &common.PageResponse{Total: uint32(total), Page: page.Page, Size: page.Size}}, nil
}

func (r *TaskRepo) Upsert(ctx context.Context, req *bizrepo.TaskUpsertReq) (*bizrepo.TaskUpsertResponse, error) {
	row := req.Row
	var result *model.Task
	if row.ID > 0 {
		current, err := r.getClient(ctx).Task.Query().Where(task.ID(row.ID), task.DeletedAtIsNil()).Only(ctx)
		if gen.IsNotFound(err) {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		if err != nil {
			return nil, err
		}
		updated, err := r.getClient(ctx).Task.UpdateOneID(row.ID).
			SetName(row.Name).
			SetTitle(row.Title).
			SetDescription(row.Description).
			SetEnabled(row.Enabled).
			SetCronSpec(row.CronSpec).
			SetPayload(row.Payload).
			SetTimeoutSeconds(row.TimeoutSeconds).
			SetAllowOverlap(row.AllowOverlap).
			SetAlertEnabled(row.AlertEnabled).
			SetVersion(current.Version + 1).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		result = r.model(updated)
	} else {
		created, err := r.getClient(ctx).Task.Create().
			SetName(row.Name).
			SetTitle(row.Title).
			SetDescription(row.Description).
			SetEnabled(row.Enabled).
			SetCronSpec(row.CronSpec).
			SetPayload(row.Payload).
			SetTimeoutSeconds(row.TimeoutSeconds).
			SetAllowOverlap(row.AllowOverlap).
			SetAlertEnabled(row.AlertEnabled).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		result = r.model(created)
	}
	return &bizrepo.TaskUpsertResponse{Row: result}, nil
}

func (r *TaskRepo) Lock(ctx context.Context, req *bizrepo.TaskLockReq) (*bizrepo.TaskLockResponse, error) {
	_, err := r.getClient(ctx).Task.Query().
		Where(
			task.ID(req.ID),
			task.DeletedAtIsNil(),
			func(s *sql.Selector) {
				s.ForUpdate()
			},
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.TaskLockResponse{}, nil
}

func (r *TaskRepo) getQuery(query *gen.TaskQuery, req *bizrepo.TaskGetReq) *gen.TaskQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(task.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(task.IDIn(req.IDs...))
	}
	if req.Name != nil {
		query = query.Where(task.NameContains(*req.Name))
	}
	if req.Title != nil {
		query = query.Where(task.TitleContains(*req.Title))
	}
	if req.Enabled != nil {
		query = query.Where(task.Enabled(*req.Enabled))
	}
	return query
}

func (r *TaskRepo) model(row *gen.Task) *model.Task {
	return &model.Task{
		ID:             row.ID,
		Name:           row.Name,
		Title:          row.Title,
		Description:    row.Description,
		Enabled:        row.Enabled,
		CronSpec:       row.CronSpec,
		Payload:        row.Payload,
		TimeoutSeconds: row.TimeoutSeconds,
		AllowOverlap:   row.AllowOverlap,
		AlertEnabled:   row.AlertEnabled,
		Version:        row.Version,
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
	}
}
