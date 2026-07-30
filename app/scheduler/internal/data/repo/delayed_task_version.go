package repo

import (
	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"context"
	"scheduler/internal/biz/model"
	bizrepo "scheduler/internal/biz/repo"
	"scheduler/internal/data/gen"
	"scheduler/internal/data/gen/delayedtaskversion"
	schedulerenum "scheduler/internal/enum"
	"time"

	entsql "entgo.io/ent/dialect/sql"
)

type DelayedTaskVersionRepo struct {
	db *gen.Client
}

func NewDelayedTaskVersionRepo(
	db *gen.Client,
) bizrepo.DelayedTaskVersionRepo {
	return &DelayedTaskVersionRepo{db: db}
}

func (r *DelayedTaskVersionRepo) Get(ctx context.Context, req *bizrepo.DelayedTaskVersionGetReq) (*model.DelayedTaskVersion, error) {
	row, err := r.getQuery(r.getClient(ctx).DelayedTaskVersion.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *DelayedTaskVersionRepo) List(ctx context.Context, req *bizrepo.DelayedTaskVersionGetReq) ([]*model.DelayedTaskVersion, error) {
	rows, err := r.getQuery(r.getClient(ctx).DelayedTaskVersion.Query(), req).Order(delayedtaskversion.ByDelayedTaskID(), delayedtaskversion.ByVersion(entsql.OrderDesc())).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.DelayedTaskVersion, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return result, nil
}

func (r *DelayedTaskVersionRepo) Create(ctx context.Context, task *model.DelayedTask) (*model.DelayedTaskVersion, error) {
	var staleAfterSeconds *int32
	if task.StaleAfter != nil {
		staleAfterSeconds = new(int32(*task.StaleAfter / time.Second))
	}
	created, err := r.getClient(ctx).DelayedTaskVersion.Create().
		SetDelayedTaskID(task.ID).
		SetVersion(task.Version).
		SetTaskKey(task.TaskKey).
		SetHandlerName(task.HandlerName.String()).
		SetTitle(task.Title).
		SetDescription(task.Description).
		SetEnabled(task.Enabled).
		SetTimeoutSeconds(int32(task.Timeout / time.Second)).
		SetNillableStaleAfterSeconds(staleAfterSeconds).
		SetMaxAttempts(task.MaxAttempts).
		SetMisfirePolicy(delayedtaskversion.MisfirePolicy(task.MisfirePolicy)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.model(created), nil
}

func (r *DelayedTaskVersionRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *DelayedTaskVersionRepo) getQuery(query *gen.DelayedTaskVersionQuery, req *bizrepo.DelayedTaskVersionGetReq) *gen.DelayedTaskVersionQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(delayedtaskversion.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(delayedtaskversion.IDIn(req.IDs...))
	}
	if req.DelayedTaskID != nil {
		query = query.Where(delayedtaskversion.DelayedTaskIDEQ(*req.DelayedTaskID))
	}
	if len(req.DelayedTaskIDs) > 0 {
		query = query.Where(delayedtaskversion.DelayedTaskIDIn(req.DelayedTaskIDs...))
	}
	if req.Version != nil {
		query = query.Where(delayedtaskversion.VersionEQ(*req.Version))
	}
	if len(req.Versions) > 0 {
		query = query.Where(delayedtaskversion.VersionIn(req.Versions...))
	}
	return query
}

func (r *DelayedTaskVersionRepo) model(row *gen.DelayedTaskVersion) *model.DelayedTaskVersion {
	var staleAfter *time.Duration
	if row.StaleAfterSeconds != nil {
		staleAfter = new(time.Duration(*row.StaleAfterSeconds) * time.Second)
	}
	return &model.DelayedTaskVersion{
		ID:            row.ID,
		DelayedTaskID: row.DelayedTaskID,
		Version:       row.Version,
		TaskKey:       row.TaskKey,
		HandlerName:   schedulerenum.TaskHandlerName(row.HandlerName),
		Title:         row.Title,
		Description:   row.Description,
		Enabled:       row.Enabled,
		Timeout:       time.Duration(row.TimeoutSeconds) * time.Second,
		StaleAfter:    staleAfter,
		MaxAttempts:   row.MaxAttempts,
		MisfirePolicy: schedulerenum.TaskMisfirePolicy(row.MisfirePolicy),
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}
}
