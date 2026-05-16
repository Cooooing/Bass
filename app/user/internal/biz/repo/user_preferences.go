package repo

import (
	"context"
	"user/internal/biz/model"
)

type UserPreferencesRepo interface {
	GetByUserID(ctx context.Context, userID int64) (*model.UserPreferences, error)
	Update(ctx context.Context, p *model.UserPreferences) (*model.UserPreferences, error)
}
