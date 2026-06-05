package repo

import (
	"context"

	userv1 "common/api/gen/user/v1"
)

type UserClient interface {
	MapAccounts(ctx context.Context, userIDs []int64) (map[int64]*userv1.AccountBasic, error)
}
