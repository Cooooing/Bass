package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/conf"
	"user/internal/data/gen"
	"user/internal/data/gen/userlocation"

	commonClient "common/pkg/client"
	utilent "common/pkg/util/ent"

	"github.com/go-kratos/kratos/v2/log"
)

var _ repo.UserLocationRepo = (*UserLocationRepo)(nil)

type UserLocationRepo struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	db     *gen.Client
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
	nats   *commonClient.NatsClient
}

func NewUserLocationRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
	nats *commonClient.NatsClient,
) repo.UserLocationRepo {
	return &UserLocationRepo{
		conf:   conf,
		log:    log.NewHelper(logger),
		db:     db,
		consul: consul,
		redis:  redis,
		nats:   nats,
	}
}

func (r *UserLocationRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
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
