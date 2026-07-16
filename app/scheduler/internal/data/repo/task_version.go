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
	"scheduler/internal/data/gen/taskversion"

	utilent "common/pkg/util/ent"
	entsql "entgo.io/ent/dialect/sql"
)

var _ bizrepo.TaskVersionRepo = (*TaskVersionRepo)(nil)

type TaskVersionRepo struct {
	db *gen.Client
}

func NewTaskVersionRepo(db *gen.Client) bizrepo.TaskVersionRepo {
	return &TaskVersionRepo{db: db}
}

func (r *TaskVersionRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *TaskVersionRepo) Get(ctx context.Context, req *bizrepo.TaskVersionGetReq) (*bizrepo.TaskVersionGetResponse, error) {
	query := r.getClient(ctx).TaskVersion.Query()
	query = r.getQuery(query, req)
	row, err := query.Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.TaskVersionGetResponse{Row: r.model(row)}, nil
}

func (r *TaskVersionRepo) List(ctx context.Context, req *bizrepo.TaskVersionGetReq) (*bizrepo.TaskVersionListResponse, error) {
	query := r.getClient(ctx).TaskVersion.Query()
	query = r.getQuery(query, req)
	rows, err := query.Order(taskversion.ByTaskID(), taskversion.ByVersion(entsql.OrderDesc())).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.TaskVersion, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return &bizrepo.TaskVersionListResponse{Rows: result}, nil
}

func (r *TaskVersionRepo) Map(ctx context.Context, req *bizrepo.TaskVersionGetReq) (*bizrepo.TaskVersionMapResponse, error) {
	resp, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.TaskVersion, len(resp.Rows))
	for _, row := range resp.Rows {
		result[row.ID] = row
	}
	return &bizrepo.TaskVersionMapResponse{Rows: result}, nil
}

func (r *TaskVersionRepo) Count(ctx context.Context, req *bizrepo.TaskVersionGetReq) (*bizrepo.TaskVersionCountResponse, error) {
	query := r.getClient(ctx).TaskVersion.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.TaskVersionCountResponse{Count: count}, nil
}

func (r *TaskVersionRepo) Page(ctx context.Context, req *bizrepo.TaskVersionPageReq) (*bizrepo.TaskVersionPageResponse, error) {
	page := server.PageValid(req.Page)
	query := r.getClient(ctx).TaskVersion.Query()
	query = r.getQuery(query, &req.TaskVersionGetReq)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := query.Order(taskversion.ByTaskID(), taskversion.ByVersion(entsql.OrderDesc())).Limit(int(page.Size)).Offset(int((page.Page - 1) * page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.TaskVersion, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.model(row))
	}
	return &bizrepo.TaskVersionPageResponse{Rows: result, Page: &common.PageResponse{Total: uint32(total), Page: page.Page, Size: page.Size}}, nil
}

func (r *TaskVersionRepo) Create(ctx context.Context, req *bizrepo.TaskVersionCreateReq) (*bizrepo.TaskVersionCreateResponse, error) {
	row := req.Task
	created, err := r.getClient(ctx).TaskVersion.Create().
		SetTaskID(row.ID).
		SetVersion(row.Version).
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
	return &bizrepo.TaskVersionCreateResponse{Row: r.model(created)}, nil
}

func (r *TaskVersionRepo) getQuery(query *gen.TaskVersionQuery, req *bizrepo.TaskVersionGetReq) *gen.TaskVersionQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(taskversion.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(taskversion.IDIn(req.IDs...))
	}
	if req.TaskID != nil {
		query = query.Where(taskversion.TaskIDEQ(*req.TaskID))
	}
	if len(req.TaskIDs) > 0 {
		query = query.Where(taskversion.TaskIDIn(req.TaskIDs...))
	}
	if req.Version != nil {
		query = query.Where(taskversion.VersionEQ(*req.Version))
	}
	if len(req.Versions) > 0 {
		query = query.Where(taskversion.VersionIn(req.Versions...))
	}
	return query
}

func (r *TaskVersionRepo) model(row *gen.TaskVersion) *model.TaskVersion {
	return &model.TaskVersion{
		ID:             row.ID,
		TaskID:         row.TaskID,
		Version:        row.Version,
		Name:           row.Name,
		Title:          row.Title,
		Description:    row.Description,
		Enabled:        row.Enabled,
		CronSpec:       row.CronSpec,
		Payload:        row.Payload,
		TimeoutSeconds: row.TimeoutSeconds,
		AllowOverlap:   row.AllowOverlap,
		AlertEnabled:   row.AlertEnabled,
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
	}
}
