package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/conf"
	"user/internal/data/gen"
	"user/internal/data/gen/userpreferences"

	commonClient "common/pkg/client"
	utilent "common/pkg/util/ent"

	"github.com/go-kratos/kratos/v2/log"
)

var _ repo.UserPreferencesRepo = (*UserPreferencesRepo)(nil)

type UserPreferencesRepo struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	db     *gen.Client
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
	nats   *commonClient.NatsClient
}

func NewUserPreferencesRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
	nats *commonClient.NatsClient,
) repo.UserPreferencesRepo {
	return &UserPreferencesRepo{
		conf:   conf,
		log:    log.NewHelper(logger),
		db:     db,
		consul: consul,
		redis:  redis,
		nats:   nats,
	}
}

func (r *UserPreferencesRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func preferencesToDomain(p *gen.UserPreferences) *model.UserPreferences {
	return &model.UserPreferences{
		ID:                   p.ID,
		UserID:               p.UserID,
		Language:             p.Language,
		Timezone:             p.Timezone,
		Theme:                p.Theme,
		MobileTheme:          p.MobileTheme,
		EnableWebNotify:      p.EnableWebNotify,
		EnableEmailSubscribe: p.EnableEmailSubscribe,
	}
}

func (r *UserPreferencesRepo) GetByUserID(ctx context.Context, userID int64) (*model.UserPreferences, error) {
	tx := r.getClient(ctx)
	p, err := tx.UserPreferences.Query().Where(userpreferences.UserID(userID)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return preferencesToDomain(p), nil
}

func (r *UserPreferencesRepo) Update(ctx context.Context, p *model.UserPreferences) (*model.UserPreferences, error) {
	tx := r.getClient(ctx)
	saved, err := tx.UserPreferences.UpdateOneID(p.ID).
		SetNillableLanguage(p.Language).
		SetNillableTimezone(p.Timezone).
		SetNillableTheme(p.Theme).
		SetNillableMobileTheme(p.MobileTheme).
		SetNillableEnableWebNotify(p.EnableWebNotify).
		SetNillableEnableEmailSubscribe(p.EnableEmailSubscribe).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return preferencesToDomain(saved), nil
}
