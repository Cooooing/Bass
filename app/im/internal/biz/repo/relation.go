package repo

import "context"

type RelationClient interface {
	BlockedCounterparts(context.Context, int64) ([]int64, error)
}
