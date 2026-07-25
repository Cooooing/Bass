package repo

import (
	"context"
	"fmt"

	"common/pkg/apperror"
	utilent "common/pkg/util/ent"
	cerrors "common/proto/gen/common/errors"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/location"
	"game_town/internal/enum"

	"github.com/samber/lo"
)

var _ bizrepo.LocationRepo = (*LocationRepo)(nil)

type LocationRepo struct {
	db *gen.Client
}

func NewLocationRepo(
	db *gen.Client,
) bizrepo.LocationRepo {
	return &LocationRepo{
		db: db,
	}
}

func (r *LocationRepo) getClient(
	ctx context.Context,
) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *LocationRepo) Save(
	ctx context.Context,
	row *model.Location,
) (*model.Location, error) {
	saved, err := r.getClient(ctx).Location.Create().
		SetWorldID(row.WorldID).
		SetCode(row.Code).
		SetName(row.Name).
		SetDescription(row.Description).
		SetStatus(location.Status(row.Status)).
		SetNillableControllingFactionID(row.ControllingFactionID).
		SetEnvironmentTags(row.EnvironmentTags).
		SetAttributes(row.Attributes).
		SetAccessible(row.Accessible).
		SetVersion(row.Version).
		SetSort(row.Sort).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Location{
		ID:                   saved.ID,
		WorldID:              saved.WorldID,
		Code:                 saved.Code,
		Name:                 saved.Name,
		Description:          saved.Description,
		Status:               enum.LocationStatus(saved.Status),
		ControllingFactionID: saved.ControllingFactionID,
		EnvironmentTags:      saved.EnvironmentTags,
		Attributes:           saved.Attributes,
		Accessible:           saved.Accessible,
		Version:              saved.Version,
		Sort:                 saved.Sort,
		CreatedAt:            saved.CreatedAt,
		UpdatedAt:            saved.UpdatedAt,
	}, nil
}

func locationQuery(
	q *gen.LocationQuery,
	req *bizrepo.LocationQuery,
) *gen.LocationQuery {
	q = q.Where(location.DeletedAtIsNil())
	if req == nil {
		return q
	}
	if req.ID != nil {
		q = q.Where(location.ID(*req.ID))
	}
	if len(req.IDs) > 0 {
		q = q.Where(location.IDIn(req.IDs...))
	}
	if req.WorldID != nil {
		q = q.Where(location.WorldID(*req.WorldID))
	}
	if req.Code != nil {
		q = q.Where(location.Code(*req.Code))
	}
	return q
}

func (r *LocationRepo) Get(
	ctx context.Context,
	req *bizrepo.LocationQuery,
) (*model.Location, error) {
	row, err := locationQuery(r.getClient(ctx).Location.Query(), req).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_LOCATION_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &model.Location{
		ID:                   row.ID,
		WorldID:              row.WorldID,
		Code:                 row.Code,
		Name:                 row.Name,
		Description:          row.Description,
		Sort:                 row.Sort,
		Status:               enum.LocationStatus(row.Status),
		ControllingFactionID: row.ControllingFactionID,
		EnvironmentTags:      row.EnvironmentTags,
		Attributes:           row.Attributes,
		Accessible:           row.Accessible,
		Version:              row.Version,
		CreatedAt:            row.CreatedAt,
		UpdatedAt:            row.UpdatedAt,
	}, nil
}

func (r *LocationRepo) List(
	ctx context.Context,
	req *bizrepo.LocationQuery,
) ([]*model.Location, error) {
	rows, err := locationQuery(r.getClient(ctx).Location.Query(), req).Order(location.BySort(), location.ByID()).All(ctx)
	if err != nil {
		return nil, err
	}
	out := lo.Map(rows, func(row *gen.Location, _ int) *model.Location {
		return &model.Location{
			ID:                   row.ID,
			WorldID:              row.WorldID,
			Code:                 row.Code,
			Name:                 row.Name,
			Description:          row.Description,
			Sort:                 row.Sort,
			Status:               enum.LocationStatus(row.Status),
			ControllingFactionID: row.ControllingFactionID,
			EnvironmentTags:      row.EnvironmentTags,
			Attributes:           row.Attributes,
			Accessible:           row.Accessible,
			Version:              row.Version,
			CreatedAt:            row.CreatedAt,
			UpdatedAt:            row.UpdatedAt,
		}
	})
	return out, nil
}

func (r *LocationRepo) Map(
	ctx context.Context,
	req *bizrepo.LocationQuery,
) (map[int64]*model.Location, error) {
	rows, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make(map[int64]*model.Location, len(rows))
	for _, row := range rows {
		out[row.ID] = row
	}
	return out, nil
}

func (r *LocationRepo) Count(
	ctx context.Context,
	req *bizrepo.LocationQuery,
) (int, error) {
	return locationQuery(r.getClient(ctx).Location.Query(), req).Count(ctx)
}

func (r *LocationRepo) Page(
	ctx context.Context,
	req *bizrepo.LocationPageReq,
) (*bizrepo.LocationPageResp, error) {
	p := page(req.Page)
	q := locationQuery(r.getClient(ctx).Location.Query(), &req.Query)
	total, err := q.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := q.Order(location.BySort(), location.ByID()).Offset(pageOffset(p)).Limit(pageLimit(p)).All(ctx)
	if err != nil {
		return nil, err
	}
	out := lo.Map(rows, func(row *gen.Location, _ int) *model.Location {
		return &model.Location{
			ID:                   row.ID,
			WorldID:              row.WorldID,
			Code:                 row.Code,
			Name:                 row.Name,
			Description:          row.Description,
			Sort:                 row.Sort,
			Status:               enum.LocationStatus(row.Status),
			ControllingFactionID: row.ControllingFactionID,
			EnvironmentTags:      row.EnvironmentTags,
			Attributes:           row.Attributes,
			Accessible:           row.Accessible,
			Version:              row.Version,
			CreatedAt:            row.CreatedAt,
			UpdatedAt:            row.UpdatedAt,
		}
	})
	return &bizrepo.LocationPageResp{
		Rows: out,
		Page: basePage(total, p),
	}, nil
}

func (r *LocationRepo) UpdateState(
	ctx context.Context,
	req *bizrepo.LocationStateUpdateReq,
) (*model.Location, error) {
	update := r.getClient(ctx).Location.Update().Where(location.ID(req.LocationID), location.Version(req.Version), location.DeletedAtIsNil())
	if req.Status != nil {
		update.SetStatus(location.Status(*req.Status))
	}
	if req.Description != "" {
		update.SetDescription(req.Description)
	}
	if req.Accessible != nil {
		update.SetAccessible(*req.Accessible)
	}
	if req.ControllingFactionID != nil {
		update.SetControllingFactionID(*req.ControllingFactionID)
	}
	if req.EnvironmentTags != nil {
		update.SetEnvironmentTags(req.EnvironmentTags)
	}
	if req.Attributes != nil {
		update.SetAttributes(req.Attributes)
	}
	update.AddVersion(1)
	count, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("location version conflict")
	}
	return r.Get(ctx, &bizrepo.LocationQuery{
		ID: new(req.LocationID),
	})
}
