package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/base"
	"user/internal/data/ent/gen"
	"user/internal/data/ent/gen/userpreferences"

	utilent "common/pkg/util/ent"
)

type UserPreferencesRepo struct {
	*base.BaseData
}

func NewUserPreferencesRepo(repo *base.BaseData) repo.UserPreferencesRepo {
	return &UserPreferencesRepo{
		BaseData: repo,
	}
}

func (r *UserPreferencesRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.Db
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
