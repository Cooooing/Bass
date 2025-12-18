package repo

import (
	cv1 "common/api/common/v1"
	v1 "common/api/user/v1"
	"context"
	"user/internal/biz/model"
	"user/internal/data/ent/gen"
)

type UserRepo interface {
	Save(ctx context.Context, tx *gen.Client, u *model.User) (*model.User, error)

	Update(ctx context.Context, tx *gen.Client, u *model.User) (*model.User, error)
	UpdateStat(ctx context.Context, tx *gen.Client, userId int64, statType v1.UserStatType, num int32) (*model.User, error)

	EnableTwoFactorAuthentication(ctx context.Context, tx *gen.Client, name string, secret string) (int, error)
	DisableTwoFactorAuthentication(ctx context.Context, tx *gen.Client, name string) (int, error)

	ConstantAccount(ctx context.Context, tx *gen.Client, account string) (bool, error)
	GetOne(ctx context.Context, tx *gen.Client, req *UserGetReq) (*model.User, error)
	GetByAccount(ctx context.Context, tx *gen.Client, account string) (*model.User, error)
	GetList(ctx context.Context, tx *gen.Client, req *UserGetReq) ([]*model.User, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *UserGetReq) ([]*model.User, *cv1.PageReply, error)
}

type UserGetReq struct {
	UserId    *int64
	UserIds   []int64
	Name      *string
	Names     []string
	Nickname  *string
	Nicknames []string
	Email     *string
	Emails    []string
	Phone     *string
	Phones    []string
}
