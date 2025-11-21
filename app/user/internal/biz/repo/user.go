package repo

import (
	cv1 "common/api/common/v1"
	"context"
	"user/internal/biz/model"
	"user/internal/data/ent/gen"
)

type UserRepo interface {
	Save(ctx context.Context, client *gen.Client, u *model.User) (*model.User, error)

	Update(ctx context.Context, tx *gen.Client, u *model.User) (*model.User, error)

	ConstantAccount(ctx context.Context, client *gen.Client, account string) (bool, error)
	GetById(ctx context.Context, client *gen.Client, id int64) (*model.User, error)
	GetByAccount(ctx context.Context, client *gen.Client, account string) (*model.User, error)
	GetList(ctx context.Context, tx *gen.Client, req *UserGetReq) ([]*model.User, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *UserGetReq) ([]*model.User, *cv1.PageReply, error)
}

type UserGetReq struct {
	ArticleIds []int64
}
