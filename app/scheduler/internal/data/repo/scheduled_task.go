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
	"scheduler/internal/data/gen/scheduledtask"
	schedulerenum "scheduler/internal/enum"
	"time"

	"entgo.io/ent/dialect/sql"
)

type ScheduledTaskRepo struct {
	db *gen.Client
}

func NewScheduledTaskRepo(
	db *gen.Client,
) bizrepo.ScheduledTaskRepo {
	return &ScheduledTaskRepo{db: db}
}

func (r *ScheduledTaskRepo) Get(ctx context.Context, req *bizrepo.ScheduledTaskGetReq) (*model.ScheduledTask, error) {
	query := r.getClient(ctx).ScheduledTask.Query().Where(scheduledtask.DeletedAtIsNil())
	row, err := r.getQuery(query, req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *ScheduledTaskRepo) List(ctx context.Context, req *bizrepo.ScheduledTaskGetReq) ([]*model.ScheduledTask, error) {
	query := r.getClient(ctx).ScheduledTask.Query().Where(scheduledtask.DeletedAtIsNil())
	rows, err := r.getQuery(query, req).Order(scheduledtask.ByID()).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.ScheduledTask, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return result, nil
}

func (r *ScheduledTaskRepo) Page(ctx context.Context, req *bizrepo.ScheduledTaskPageReq) (*bizrepo.ScheduledTaskPageResp, error) {
	page := server.PageValid(req.Page)
	query := r.getClient(ctx).ScheduledTask.Query().Where(scheduledtask.DeletedAtIsNil())
	query = r.getQuery(query, &req.ScheduledTaskGetReq)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := query.Order(scheduledtask.ByID()).Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.ScheduledTask, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return &bizrepo.ScheduledTaskPageResp{Rows: result, Page: &common.PageResp{Total: uint32(total), Page: page.Page, Size: page.Size}}, nil
}

func (r *ScheduledTaskRepo) MapByTitle(ctx context.Context, titles []string) (map[string]*model.ScheduledTask, error) {
	if len(titles) == 0 {
		return map[string]*model.ScheduledTask{}, nil
	}
	rows, err := r.List(ctx, &bizrepo.ScheduledTaskGetReq{Titles: titles})
	if err != nil {
		return nil, err
	}
	result := make(map[string]*model.ScheduledTask, len(rows))
	for _, row := range rows {
		result[row.Title] = row
	}
	return result, nil
}

func (r *ScheduledTaskRepo) Upsert(ctx context.Context, row *model.ScheduledTask) (*model.ScheduledTask, error) {
	var staleAfterSeconds *int32
	if row.StaleAfter != nil {
		staleAfterSeconds = new(int32(*row.StaleAfter / time.Second))
	}
	if row.ID > 0 {
		current, err := r.getClient(ctx).ScheduledTask.Query().Where(scheduledtask.ID(row.ID), scheduledtask.DeletedAtIsNil()).Only(ctx)
		if gen.IsNotFound(err) {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		if err != nil {
			return nil, err
		}
		updated, err := r.getClient(ctx).ScheduledTask.UpdateOneID(row.ID).
			SetName(row.Name).
			SetTitle(row.Title).
			SetDescription(row.Description).
			SetEnabled(row.Enabled).
			SetCronSpec(row.CronSpec).
			SetPayload(row.Payload).
			SetTimeoutSeconds(int32(row.Timeout / time.Second)).
			SetNillableStaleAfterSeconds(staleAfterSeconds).
			SetMaxAttempts(row.MaxAttempts).
			SetMisfirePolicy(scheduledtask.MisfirePolicy(row.MisfirePolicy)).
			SetAllowOverlap(row.AllowOverlap).
			SetVersion(current.Version + 1).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return r.model(updated), nil
	}
	created, err := r.getClient(ctx).ScheduledTask.Create().
		SetName(row.Name).
		SetTitle(row.Title).
		SetDescription(row.Description).
		SetEnabled(row.Enabled).
		SetCronSpec(row.CronSpec).
		SetPayload(row.Payload).
		SetTimeoutSeconds(int32(row.Timeout / time.Second)).
		SetNillableStaleAfterSeconds(staleAfterSeconds).
		SetMaxAttempts(row.MaxAttempts).
		SetMisfirePolicy(scheduledtask.MisfirePolicy(row.MisfirePolicy)).
		SetAllowOverlap(row.AllowOverlap).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.model(created), nil
}

func (r *ScheduledTaskRepo) Lock(ctx context.Context, id int64) error {
	_, err := r.getClient(ctx).ScheduledTask.Query().Where(scheduledtask.ID(id), scheduledtask.DeletedAtIsNil(), func(s *sql.Selector) { s.ForUpdate() }).Only(ctx)
	return err
}

func (r *ScheduledTaskRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *ScheduledTaskRepo) getQuery(query *gen.ScheduledTaskQuery, req *bizrepo.ScheduledTaskGetReq) *gen.ScheduledTaskQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(scheduledtask.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(scheduledtask.IDIn(req.IDs...))
	}
	if req.Name != nil {
		query = query.Where(scheduledtask.NameContains(*req.Name))
	}
	if req.Title != nil {
		query = query.Where(scheduledtask.TitleContains(*req.Title))
	}
	if len(req.Titles) > 0 {
		query = query.Where(scheduledtask.TitleIn(req.Titles...))
	}
	if req.Enabled != nil {
		query = query.Where(scheduledtask.Enabled(*req.Enabled))
	}
	return query
}

func (r *ScheduledTaskRepo) model(row *gen.ScheduledTask) *model.ScheduledTask {
	var staleAfter *time.Duration
	if row.StaleAfterSeconds != nil {
		staleAfter = new(time.Duration(*row.StaleAfterSeconds) * time.Second)
	}
	return &model.ScheduledTask{
		ID:            row.ID,
		Name:          row.Name,
		Title:         row.Title,
		Description:   row.Description,
		Enabled:       row.Enabled,
		CronSpec:      row.CronSpec,
		Payload:       row.Payload,
		Timeout:       time.Duration(row.TimeoutSeconds) * time.Second,
		StaleAfter:    staleAfter,
		MaxAttempts:   row.MaxAttempts,
		MisfirePolicy: schedulerenum.TaskMisfirePolicy(row.MisfirePolicy),
		AllowOverlap:  row.AllowOverlap,
		Version:       row.Version,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}
}
