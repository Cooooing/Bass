package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/location"

	utilent "common/pkg/util/ent"
)

var _ repo.LocationRepo = (*LocationRepo)(nil)

type LocationRepo struct {
	db *gen.Client
}

func NewLocationRepo(db *gen.Client) repo.LocationRepo {
	return &LocationRepo{
		db: db,
	}
}

func (r *LocationRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *LocationRepo) FindByUserID(ctx context.Context, userID int64) (*model.Location, error) {
	tx := r.getClient(ctx)
	l, err := tx.Location.Query().Where(location.UserID(userID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &model.Location{
		ID:       l.ID,
		UserID:   l.UserID,
		Country:  l.Country,
		Province: l.Province,
		City:     l.City,
	}, nil
}

func (r *LocationRepo) UpsertByUserID(ctx context.Context, l *model.Location) (*model.Location, error) {
	existing, err := r.FindByUserID(ctx, l.UserID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		tx := r.getClient(ctx)
		saved, err := tx.Location.Create().
			SetUserID(l.UserID).
			SetNillableCountry(l.Country).
			SetNillableProvince(l.Province).
			SetNillableCity(l.City).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.Location{
			ID:       saved.ID,
			UserID:   saved.UserID,
			Country:  saved.Country,
			Province: saved.Province,
			City:     saved.City,
		}, nil
	}
	l.ID = existing.ID
	return r.Update(ctx, l)
}

func (r *LocationRepo) Update(ctx context.Context, l *model.Location) (*model.Location, error) {
	tx := r.getClient(ctx)
	saved, err := tx.Location.UpdateOneID(l.ID).
		SetNillableCountry(l.Country).
		SetNillableProvince(l.Province).
		SetNillableCity(l.City).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Location{
		ID:       saved.ID,
		UserID:   saved.UserID,
		Country:  saved.Country,
		Province: saved.Province,
		City:     saved.City,
	}, nil
}
