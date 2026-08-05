package repo

import (
	"context"
	"economy/internal/biz/base"
	"economy/internal/biz/model"
)

type AccountRepo interface {
	Save(ctx context.Context, account *model.Account) (*model.Account, error)
	UpdateBalance(ctx context.Context, req *AccountUpdateBalanceReq) (*model.Account, error)
	Get(ctx context.Context, req *AccountGetReq) (*model.Account, error)
	List(ctx context.Context, req *AccountGetReq) ([]*model.Account, error)
	Map(ctx context.Context, req *AccountGetReq) (map[int64]*model.Account, error)
	Count(ctx context.Context, req *AccountGetReq) (int, error)
	Page(ctx context.Context, req *AccountGetReq) (*AccountPageResp, error)
}

type AccountGetReq struct {
	Page    *base.PageRequest
	ID      *int64
	IDs     []int64
	UserID  *int64
	UserIDs []int64
}

type AccountUpdateBalanceReq struct {
	UserID       int64
	BalanceDelta int64
	IncomeDelta  int64
	ExpenseDelta int64
}

type AccountPageResp struct {
	Rows []*model.Account
	Page *base.PageResp
}
