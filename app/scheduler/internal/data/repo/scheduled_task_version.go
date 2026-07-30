package repo

import (
	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"context"
	"scheduler/internal/biz/model"
	bizrepo "scheduler/internal/biz/repo"
	"scheduler/internal/data/gen"
	"scheduler/internal/data/gen/scheduledtaskversion"
	schedulerenum "scheduler/internal/enum"
	"time"

	entsql "entgo.io/ent/dialect/sql"
)

type ScheduledTaskVersionRepo struct {
	db *gen.Client
}

func NewScheduledTaskVersionRepo(
	db *gen.Client,
) bizrepo.ScheduledTaskVersionRepo {
	return &ScheduledTaskVersionRepo{db: db}
}

func (r *ScheduledTaskVersionRepo) Get(ctx context.Context, req *bizrepo.ScheduledTaskVersionGetReq) (*model.ScheduledTaskVersion, error) {
	row, err := r.getQuery(r.getClient(ctx).ScheduledTaskVersion.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err != nil {
		return nil, err
	}
	return r.model(row), nil
}

func (r *ScheduledTaskVersionRepo) List(ctx context.Context, req *bizrepo.ScheduledTaskVersionGetReq) ([]*model.ScheduledTaskVersion, error) {
	rows, err := r.getQuery(r.getClient(ctx).ScheduledTaskVersion.Query(), req).
		Order(
			scheduledtaskversion.ByScheduledTaskID(),
			scheduledtaskversion.ByVersion(entsql.OrderDesc()),
		).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.ScheduledTaskVersion, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return result, nil
}

func (r *ScheduledTaskVersionRepo) Create(ctx context.Context, task *model.ScheduledTask) (*model.ScheduledTaskVersion, error) {
	var staleAfterSeconds *int32
	if task.StaleAfter != nil {
		staleAfterSeconds = new(int32(*task.StaleAfter / time.Second))
	}
	created, err := r.getClient(ctx).ScheduledTaskVersion.Create().
		SetScheduledTaskID(task.ID).
		SetVersion(task.Version).
		SetTaskKey(task.TaskKey).
		SetHandlerName(task.HandlerName.String()).
		SetTitle(task.Title).
		SetDescription(task.Description).
		SetEnabled(task.Enabled).
		SetCronSpec(task.CronSpec).
		SetPayload(task.Payload).
		SetTimeoutSeconds(int32(task.Timeout / time.Second)).
		SetNillableStaleAfterSeconds(staleAfterSeconds).
		SetMaxAttempts(task.MaxAttempts).
		SetMisfirePolicy(scheduledtaskversion.MisfirePolicy(task.MisfirePolicy)).
		SetAllowOverlap(task.AllowOverlap).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.model(created), nil
}

func (r *ScheduledTaskVersionRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *ScheduledTaskVersionRepo) getQuery(query *gen.ScheduledTaskVersionQuery, req *bizrepo.ScheduledTaskVersionGetReq) *gen.ScheduledTaskVersionQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(scheduledtaskversion.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(scheduledtaskversion.IDIn(req.IDs...))
	}
	if req.ScheduledTaskID != nil {
		query = query.Where(scheduledtaskversion.ScheduledTaskIDEQ(*req.ScheduledTaskID))
	}
	if len(req.ScheduledTaskIDs) > 0 {
		query = query.Where(scheduledtaskversion.ScheduledTaskIDIn(req.ScheduledTaskIDs...))
	}
	if req.Version != nil {
		query = query.Where(scheduledtaskversion.VersionEQ(*req.Version))
	}
	if len(req.Versions) > 0 {
		query = query.Where(scheduledtaskversion.VersionIn(req.Versions...))
	}
	return query
}

func (r *ScheduledTaskVersionRepo) model(row *gen.ScheduledTaskVersion) *model.ScheduledTaskVersion {
	var staleAfter *time.Duration
	if row.StaleAfterSeconds != nil {
		staleAfter = new(time.Duration(*row.StaleAfterSeconds) * time.Second)
	}
	return &model.ScheduledTaskVersion{
		ID:              row.ID,
		ScheduledTaskID: row.ScheduledTaskID,
		Version:         row.Version,
		TaskKey:         row.TaskKey,
		HandlerName:     schedulerenum.TaskHandlerName(row.HandlerName),
		Title:           row.Title,
		Description:     row.Description,
		Enabled:         row.Enabled,
		CronSpec:        row.CronSpec,
		Payload:         row.Payload,
		Timeout:         time.Duration(row.TimeoutSeconds) * time.Second,
		StaleAfter:      staleAfter,
		MaxAttempts:     row.MaxAttempts,
		MisfirePolicy:   schedulerenum.TaskMisfirePolicy(row.MisfirePolicy),
		AllowOverlap:    row.AllowOverlap,
		CreatedAt:       row.CreatedAt,
		UpdatedAt:       row.UpdatedAt,
	}
}
