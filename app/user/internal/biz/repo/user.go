package repo

import (
	"common/api/gen/common"
	v1 "common/api/gen/user/v1"
	"context"
	"user/internal/biz/model"
)

type UserRepo interface {
	Save(ctx context.Context, u *model.User) (*model.User, error)

	Update(ctx context.Context, u *model.User) (*model.User, error)
	UpdateStat(ctx context.Context, userId int64, statType v1.UserStatType, num int32) (*model.User, error)

	ConstantAccount(ctx context.Context, account string) (bool, error)
	GetOne(ctx context.Context, req *UserGetReq) (*model.User, error)
	GetByAccount(ctx context.Context, account string) (*model.User, error)
	GetList(ctx context.Context, req *UserGetReq) ([]*model.User, error)
	GetPage(ctx context.Context, page *common.PageRequest, req *UserGetReq) ([]*model.User, *common.PageReply, error)
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
