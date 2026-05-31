package repo

import (
	"common/api/gen/common"
	"context"
	"user/internal/biz/model"
	"user/internal/enum"
)

type AccountRepo interface {
	Create(ctx context.Context, u *model.Account) (*model.Account, error)

	Update(ctx context.Context, u *model.Account) (*model.Account, error)
	UpdateProfile(ctx context.Context, req *model.AccountProfileUpdate) (*model.Account, error)
	AddStat(ctx context.Context, userId int64, statType enum.AccountStatType, num int32) (*model.Account, error)

	ExistsByAccount(ctx context.Context, account string) (bool, error)
	Get(ctx context.Context, req *AccountGetReq) (*model.Account, error)
	GetByAccount(ctx context.Context, account string) (*model.Account, error)
	List(ctx context.Context, req *AccountGetReq) ([]*model.Account, error)
	Map(ctx context.Context, req *AccountGetReq) (map[int64]*model.Account, error)
	Page(ctx context.Context, page *common.PageRequest, req *AccountGetReq) ([]*model.Account, *common.PageReply, error)
}

type AccountGetReq struct {
	UserID    *int64
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
