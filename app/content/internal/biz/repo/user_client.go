package repo

import (
	"context"

	"content/internal/biz/model"
)

type UserClient interface {
	MapAccounts(ctx context.Context, userIDs []int64) (map[int64]*model.UserAccountBasic, error)
}
