package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/base"
	"user/internal/data/ent/gen"
	"user/internal/data/ent/gen/userlocation"

	utilent "common/pkg/util/ent"
)

type UserLocationRepo struct {
	*base.BaseData
}

func NewUserLocationRepo(repo *base.BaseData) repo.UserLocationRepo {
	return &UserLocationRepo{
		BaseData: repo,
	}
}

func (r *UserLocationRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.Db
}

func locationToDomain(l *gen.UserLocation) *model.UserLocation {
	return &model.UserLocation{
		ID:       l.ID,
		UserID:   l.UserID,
		Country:  l.Country,
		Province: l.Province,
		City:     l.City,
	}
}

func (r *UserLocationRepo) GetByUserID(ctx context.Context, userID int64) (*model.UserLocation, error) {
	tx := r.getClient(ctx)
	l, err := tx.UserLocation.Query().Where(userlocation.UserID(userID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return locationToDomain(l), nil
}

func (r *UserLocationRepo) Update(ctx context.Context, l *model.UserLocation) (*model.UserLocation, error) {
	tx := r.getClient(ctx)
	saved, err := tx.UserLocation.UpdateOneID(l.ID).
		SetNillableCountry(l.Country).
		SetNillableProvince(l.Province).
		SetNillableCity(l.City).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return locationToDomain(saved), nil
}
