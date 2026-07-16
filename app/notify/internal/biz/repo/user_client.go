package repo

import (
	"context"
	"notify/internal/biz/model"
)

type UserClient interface {
	MapAccounts(ctx context.Context, req *UserMapAccountsReq) (*UserMapAccountsResponse, error)
	ListFollowerIDs(ctx context.Context, req *UserListFollowerIDsReq) (*UserListFollowerIDsResponse, error)
}

type UserMapAccountsReq struct {
	UserIDs []int64
}

type UserMapAccountsResponse struct {
	Rows map[int64]*model.UserAccount
}

type UserListFollowerIDsReq struct {
	UserID int64
}

type UserListFollowerIDsResponse struct {
	UserIDs []int64
}
