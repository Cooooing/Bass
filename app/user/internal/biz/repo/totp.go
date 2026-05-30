package repo

import (
	"context"
	"user/internal/biz/model"
)

type TotpRepo interface {
	FindByUserID(ctx context.Context, userID int64) (*model.Totp, error)
	UpsertEnabledByUserID(ctx context.Context, userID int64, secret string) (*model.Totp, error)
	DisableByUserID(ctx context.Context, userID int64) (*model.Totp, error)
}
