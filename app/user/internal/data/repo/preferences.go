package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/conf"
	"user/internal/data/gen"
	"user/internal/data/gen/preferences"

	commonClient "common/pkg/client"
	utilent "common/pkg/util/ent"

	"github.com/go-kratos/kratos/v2/log"
)

var _ repo.PreferencesRepo = (*PreferencesRepo)(nil)

type PreferencesRepo struct {
	conf   *conf.Bootstrap
	log    *log.Helper
	db     *gen.Client
	consul *commonClient.ConsulClient
	redis  *commonClient.RedisClient
	nats   *commonClient.NatsClient
}

func NewPreferencesRepo(
	conf *conf.Bootstrap,
	logger log.Logger,
	db *gen.Client,
	consul *commonClient.ConsulClient,
	redis *commonClient.RedisClient,
	nats *commonClient.NatsClient,
) repo.PreferencesRepo {
	return &PreferencesRepo{
		conf:   conf,
		log:    log.NewHelper(logger),
		db:     db,
		consul: consul,
		redis:  redis,
		nats:   nats,
	}
}

func (r *PreferencesRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func preferencesToDomain(p *gen.Preferences) *model.Preferences {
	return &model.Preferences{
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

func (r *PreferencesRepo) FindByUserID(ctx context.Context, userID int64) (*model.Preferences, error) {
	tx := r.getClient(ctx)
	p, err := tx.Preferences.Query().Where(preferences.UserID(userID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return preferencesToDomain(p), nil
}

func (r *PreferencesRepo) UpsertByUserID(ctx context.Context, p *model.Preferences) (*model.Preferences, error) {
	existing, err := r.FindByUserID(ctx, p.UserID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		tx := r.getClient(ctx)
		saved, err := tx.Preferences.Create().
			SetUserID(p.UserID).
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
	p.ID = existing.ID
	return r.Update(ctx, p)
}

func (r *PreferencesRepo) Update(ctx context.Context, p *model.Preferences) (*model.Preferences, error) {
	tx := r.getClient(ctx)
	saved, err := tx.Preferences.UpdateOneID(p.ID).
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
