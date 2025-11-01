package repo

import (
	cv1 "common/api/common/v1"
	"context"
	"user/internal/biz/model"
	"user/internal/data/ent/gen"
)

type UserRepo interface {
	Save(ctx context.Context, client *gen.Client, u *model.User) (*model.User, error)

	GetUserById(ctx context.Context, client *gen.Client, id int64) (*model.User, error)
	GetUserByAccount(ctx context.Context, client *gen.Client, account string) (*model.User, error)
	GetUserList(ctx context.Context, client *gen.Client, page *cv1.PageRequest, ids []int64) ([]*model.User, error)
	ConstantAccount(ctx context.Context, client *gen.Client, account string) (bool, error)
}
