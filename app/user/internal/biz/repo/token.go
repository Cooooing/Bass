package repo

import (
	"context"
	"user/internal/biz/model"
)

type TokenRepo interface {
	SaveEmailToken(ctx context.Context, token string, code string, user *model.User) error
	GetEmailToken(ctx context.Context, token string) (string, *model.User, error)
	DelEmailToken(ctx context.Context, token string) error

	SaveToken(ctx context.Context, token string, user *model.User) error
	GetToken(ctx context.Context, token string) (*model.User, error)
	DelToken(ctx context.Context, token string) error
}
