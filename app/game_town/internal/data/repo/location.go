package repo

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/location"
)

type LocationRepo struct{ *baseRepo }

func NewLocationRepo(db *gen.Client) bizrepo.LocationRepo {
	return &LocationRepo{baseRepo: &baseRepo{db: db}}
}

func (r *LocationRepo) GetLocationByCode(ctx context.Context, req *bizrepo.GetLocationByCodeReq) (*model.Location, error) {
	row, err := r.db.Location.Query().Where(location.WorldID(req.WorldID), location.Code(req.Code), location.DeletedAtIsNil()).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_LOCATION_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.location(row), nil
}

func (r *LocationRepo) GetLocation(ctx context.Context, id int64) (*model.Location, error) {
	row, err := r.db.Location.Query().Where(location.ID(id), location.DeletedAtIsNil()).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_LOCATION_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.location(row), nil
}
