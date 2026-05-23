package repo

import (
	"context"
	"user/internal/biz/model"
)

type PreferencesRepo interface {
	FindByUserID(ctx context.Context, userID int64) (*model.Preferences, error)
	UpsertByUserID(ctx context.Context, p *model.Preferences) (*model.Preferences, error)
	Update(ctx context.Context, p *model.Preferences) (*model.Preferences, error)
}
