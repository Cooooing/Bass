package repo

import (
	"context"
	"user/internal/biz/model"
)

type UserRepo interface {
	Save(ctx context.Context, u *model.User) (*model.User, error)

	GetUserById(ctx context.Context, id int) (*model.User, error)
	GetUserByAccount(ctx context.Context, account string) (*model.User, error)
	ConstantAccount(ctx context.Context, account string) (bool, error)
}
