package repo

import (
	"context"

	userv1 "common/api/gen/user/v1"
)

type UserClient interface {
	BatchGetBasicAccounts(ctx context.Context, userIDs []int64) (map[int64]*userv1.AccountBasic, error)
}
