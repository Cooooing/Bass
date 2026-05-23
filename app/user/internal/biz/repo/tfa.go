package repo

import (
	"context"
	"user/internal/biz/model"
)

type TfaRepo interface {
	FindByUserID(ctx context.Context, userID int64) (*model.TFA, error)
	UpsertEnabledByUserID(ctx context.Context, userID int64, secret string) (*model.TFA, error)
	DisableByUserID(ctx context.Context, userID int64) (*model.TFA, error)
}
