package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/conf"
	"user/internal/data/gen"
	"user/internal/data/gen/userprivacy"

	commonClient "common/pkg/client"
	utilent "common/pkg/util/ent"

	"github.com/go-kratos/kratos/v2/log"
)

type UserPrivacyRepo struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	db     *gen.Client
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
	nats   *commonClient.NatsClient
}

func NewUserPrivacyRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
	nats *commonClient.NatsClient,
) repo.UserPrivacyRepo {
	return &UserPrivacyRepo{
		conf:   conf,
		log:    log.NewHelper(logger),
		db:     db,
		consul: consul,
		redis:  redis,
		nats:   nats,
	}
}

func (r *UserPrivacyRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func privacyToDomain(p *gen.UserPrivacy) *model.UserPrivacy {
	return &model.UserPrivacy{
		ID:                 p.ID,
		UserID:             p.UserID,
		PublicPoints:       p.PublicPoints,
		PublicFollowers:    p.PublicFollowers,
		PublicArticles:     p.PublicArticles,
		PublicComments:     p.PublicComments,
		PublicOnlineStatus: p.PublicOnlineStatus,
		PublicLocation:     p.PublicLocation,
	}
}

func (r *UserPrivacyRepo) GetByUserID(ctx context.Context, userID int64) (*model.UserPrivacy, error) {
	tx := r.getClient(ctx)
	p, err := tx.UserPrivacy.Query().Where(userprivacy.UserID(userID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return privacyToDomain(p), nil
}

func (r *UserPrivacyRepo) Update(ctx context.Context, p *model.UserPrivacy) (*model.UserPrivacy, error) {
	tx := r.getClient(ctx)
	saved, err := tx.UserPrivacy.UpdateOneID(p.ID).
		SetNillablePublicPoints(p.PublicPoints).
		SetNillablePublicFollowers(p.PublicFollowers).
		SetNillablePublicArticles(p.PublicArticles).
		SetNillablePublicComments(p.PublicComments).
		SetNillablePublicOnlineStatus(p.PublicOnlineStatus).
		SetNillablePublicLocation(p.PublicLocation).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return privacyToDomain(saved), nil
}
