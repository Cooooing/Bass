package repo

import (
	"context"

	"notify/internal/biz/model"
)

type UserAccountRepo interface {
	MapAccounts(ctx context.Context, userIDs []int64) (map[int64]*model.UserAccount, error)
	MapAccountsByNames(ctx context.Context, names []string) (map[int64]*model.UserAccount, error)
	ListFollowerIDs(ctx context.Context, userID int64) ([]int64, error)
	MapRelationStatuses(ctx context.Context, userID int64, targetIDs []int64) (map[int64]*model.RelationStatus, error)
}
