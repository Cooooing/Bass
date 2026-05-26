package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/biz/repo"
	"user/internal/data/gen"
	"user/internal/data/gen/preferences"
	"user/internal/enum"

	utilent "common/pkg/util/ent"
)

var _ repo.PreferencesRepo = (*PreferencesRepo)(nil)

type PreferencesRepo struct {
	db *gen.Client
}

func NewPreferencesRepo(db *gen.Client) repo.PreferencesRepo {
	return &PreferencesRepo{
		db: db,
	}
}

func (r *PreferencesRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
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
	return &model.Preferences{
		ID:          p.ID,
		UserID:      p.UserID,
		Language:    (*enum.Language)(p.Language),
		Timezone:    p.Timezone,
		Theme:       p.Theme,
		MobileTheme: p.MobileTheme,
	}, nil
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
			SetNillableLanguage((*preferences.Language)(p.Language)).
			SetNillableTimezone(p.Timezone).
			SetNillableTheme(p.Theme).
			SetNillableMobileTheme(p.MobileTheme).
			Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.Preferences{
			ID:          saved.ID,
			UserID:      saved.UserID,
			Language:    (*enum.Language)(saved.Language),
			Timezone:    saved.Timezone,
			Theme:       saved.Theme,
			MobileTheme: saved.MobileTheme,
		}, nil
	}
	p.ID = existing.ID
	return r.Update(ctx, p)
}

func (r *PreferencesRepo) Update(ctx context.Context, p *model.Preferences) (*model.Preferences, error) {
	tx := r.getClient(ctx)
	saved, err := tx.Preferences.UpdateOneID(p.ID).
		SetNillableLanguage((*preferences.Language)(p.Language)).
		SetNillableTimezone(p.Timezone).
		SetNillableTheme(p.Theme).
		SetNillableMobileTheme(p.MobileTheme).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Preferences{
		ID:          saved.ID,
		UserID:      saved.UserID,
		Language:    (*enum.Language)(saved.Language),
		Timezone:    saved.Timezone,
		Theme:       saved.Theme,
		MobileTheme: saved.MobileTheme,
	}, nil
}
