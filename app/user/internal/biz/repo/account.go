package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/enum"
)

type AccountRepo interface {
	Create(ctx context.Context, req *AccountCreateReq) (*AccountCreateResponse, error)

	Update(ctx context.Context, req *AccountUpdateReq) (*AccountUpdateResponse, error)
	UpdateProfile(ctx context.Context, req *AccountUpdateProfileReq) (*AccountUpdateProfileResponse, error)
	AddStat(ctx context.Context, req *AccountAddStatReq) (*AccountAddStatResponse, error)

	ExistsByAccount(ctx context.Context, req *AccountExistsByAccountReq) (*AccountExistsByAccountResponse, error)
	Get(ctx context.Context, req *AccountGetReq) (*AccountGetResponse, error)
	List(ctx context.Context, req *AccountGetReq) (*AccountListResponse, error)
	Map(ctx context.Context, req *AccountGetReq) (*AccountMapResponse, error)
	Count(ctx context.Context, req *AccountGetReq) (*AccountCountResponse, error)
	Page(ctx context.Context, req *AccountPageReq) (*AccountPageResponse, error)
}

type AccountCreateReq struct {
	Account *model.Account
}

type AccountCreateResponse struct {
	Account *model.Account
}

type AccountUpdateReq struct {
	Account *model.Account
}

type AccountUpdateResponse struct {
	Account *model.Account
}

type AccountUpdateProfileReq struct {
	Profile *model.AccountProfileUpdate
}

type AccountUpdateProfileResponse struct {
	Account *model.Account
}

type AccountAddStatReq struct {
	UserID   int64
	StatType enum.AccountStatType
	Num      int32
}

type AccountAddStatResponse struct {
	Account *model.Account
}

type AccountExistsByAccountReq struct {
	Account string
}

type AccountExistsByAccountResponse struct {
	Exists bool
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

type AccountGetResponse struct {
	Account *model.Account
}

type AccountListResponse struct {
	Rows []*model.Account
}

type AccountMapResponse struct {
	Rows map[int64]*model.Account
}

type AccountCountResponse struct {
	Count int
}

type AccountPageReq struct {
	Page  PageReq
	Query AccountGetReq
}

type AccountPageResponse struct {
	Rows []*model.Account
	Page PageResponse
}
