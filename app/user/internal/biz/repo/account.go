package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/enum"
)

type AccountRepo interface {
	Create(ctx context.Context, account *model.Account) (*model.Account, error)

	Update(ctx context.Context, account *model.Account) (*model.Account, error)
	UpdateProfile(ctx context.Context, profile *model.AccountProfileUpdate) (*model.Account, error)
	AddStat(ctx context.Context, req *AccountAddStatReq) (*model.Account, error)

	UpdateStatus(ctx context.Context, userID int64, status enum.AccountStatus) (*model.Account, error)

	ExistsByAccount(ctx context.Context, account string) (bool, error)
	Get(ctx context.Context, req *AccountGetReq) (*model.Account, error)
	List(ctx context.Context, req *AccountGetReq) ([]*model.Account, error)
	Map(ctx context.Context, req *AccountGetReq) (map[int64]*model.Account, error)
	Count(ctx context.Context, req *AccountGetReq) (int, error)
	Page(ctx context.Context, req *AccountPageReq) (*AccountPageResp, error)
}

type AccountAddStatReq struct {
	UserID   int64
	StatType enum.AccountStatType
	Num      int32
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
	Account   *string
}

type AccountPageReq struct {
	Page  PageReq
	Query AccountGetReq
}

type AccountPageResp struct {
	Rows []*model.Account
	Page PageResp
}
