package repo

import "context"

type RelationClient interface {
	BlockedCounterparts(ctx context.Context, userID int64) ([]int64, error)
}
