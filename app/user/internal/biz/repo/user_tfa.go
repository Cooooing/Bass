package repo

import (
	"context"
	"user/internal/biz/model"
)

type UserTfaRepo interface {
	GetByUserID(ctx context.Context, userID int64) (*model.UserTFA, error)
	Enable(ctx context.Context, userID int64, secret string) (*model.UserTFA, error)
	Disable(ctx context.Context, userID int64) (*model.UserTFA, error)
}
