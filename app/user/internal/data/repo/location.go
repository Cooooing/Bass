package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/conf"
	"user/internal/data/gen"
	"user/internal/data/gen/location"

	commonClient "common/pkg/client"
	utilent "common/pkg/util/ent"

	"github.com/go-kratos/kratos/v2/log"
)

var _ repo.LocationRepo = (*LocationRepo)(nil)

type LocationRepo struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	db     *gen.Client
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
	nats   *commonClient.NatsClient
}

func NewLocationRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
	nats *commonClient.NatsClient,
) repo.LocationRepo {
	return &LocationRepo{
		conf:   conf,
		log:    log.NewHelper(logger),
		db:     db,
		consul: consul,
		redis:  redis,
		nats:   nats,
	}
}

func (r *LocationRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func locationToDomain(l *gen.Location) *model.Location {
	return &model.Location{
		ID:       l.ID,
		UserID:   l.UserID,
		Country:  l.Country,
		Province: l.Province,
		City:     l.City,
	}
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
	return locationToDomain(l), nil
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
		return locationToDomain(saved), nil
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
	return locationToDomain(saved), nil
}
