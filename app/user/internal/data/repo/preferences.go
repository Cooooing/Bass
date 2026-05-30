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
		Language:    new(enum.Language(p.Language)),
		Timezone:    new(p.Timezone),
		Theme:       new(p.Theme),
		MobileTheme: new(p.MobileTheme),
	}, nil
}

func (r *PreferencesRepo) UpsertByUserID(ctx context.Context, p *model.Preferences) (*model.Preferences, error) {
	existing, err := r.FindByUserID(ctx, p.UserID)
	if err != nil {
		return nil, err
	}
	if existing == nil {
		tx := r.getClient(ctx)
		create := tx.Preferences.Create().
			SetUserID(p.UserID)
		if p.Language != nil {
			create.SetLanguage(preferences.Language(*p.Language))
		}
		if p.Timezone != nil {
			create.SetTimezone(*p.Timezone)
		}
		if p.Theme != nil {
			create.SetTheme(*p.Theme)
		}
		if p.MobileTheme != nil {
			create.SetMobileTheme(*p.MobileTheme)
		}
		saved, err := create.Save(ctx)
		if err != nil {
			return nil, err
		}
		return &model.Preferences{
			ID:          saved.ID,
			UserID:      saved.UserID,
			Language:    new(enum.Language(saved.Language)),
			Timezone:    new(saved.Timezone),
			Theme:       new(saved.Theme),
			MobileTheme: new(saved.MobileTheme),
		}, nil
	}
	p.ID = existing.ID
	return r.Update(ctx, p)
}

func (r *PreferencesRepo) Update(ctx context.Context, p *model.Preferences) (*model.Preferences, error) {
	tx := r.getClient(ctx)
	if p.Language == nil && p.Timezone == nil && p.Theme == nil && p.MobileTheme == nil {
		saved, err := tx.Preferences.Get(ctx, p.ID)
		if err != nil {
			return nil, err
		}
		return &model.Preferences{
			ID:          saved.ID,
			UserID:      saved.UserID,
			Language:    new(enum.Language(saved.Language)),
			Timezone:    new(saved.Timezone),
			Theme:       new(saved.Theme),
			MobileTheme: new(saved.MobileTheme),
		}, nil
	}
	update := tx.Preferences.UpdateOneID(p.ID)
	if p.Language != nil {
		update.SetLanguage(preferences.Language(*p.Language))
	}
	if p.Timezone != nil {
		update.SetTimezone(*p.Timezone)
	}
	if p.Theme != nil {
		update.SetTheme(*p.Theme)
	}
	if p.MobileTheme != nil {
		update.SetMobileTheme(*p.MobileTheme)
	}
	saved, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.Preferences{
		ID:          saved.ID,
		UserID:      saved.UserID,
		Language:    new(enum.Language(saved.Language)),
		Timezone:    new(saved.Timezone),
		Theme:       new(saved.Theme),
		MobileTheme: new(saved.MobileTheme),
	}, nil
}
