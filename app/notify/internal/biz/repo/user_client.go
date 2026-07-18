package repo

import (
	"context"
	"notify/internal/biz/model"
)

type UserClient interface {
	MapAccounts(ctx context.Context, userIDs []int64) (map[int64]*model.UserAccount, error)
	ListFollowerIDs(ctx context.Context, userID int64) ([]int64, error)
}
