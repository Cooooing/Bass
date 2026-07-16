package repo

import (
	"context"

	"content/internal/biz/model"
)

type UserClient interface {
	MapAccounts(ctx context.Context, req *MapAccountsReq) (*MapAccountsResponse, error)
}

type MapAccountsReq struct {
	UserIDs []int64
}

type MapAccountsResponse struct {
	Rows map[int64]*model.UserAccountBasic
}
