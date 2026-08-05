package base

import (
	utilent "common/pkg/util/ent"
	"context"
)

type Tx func(ctx context.Context, fn func(ctx context.Context) error, opts ...utilent.TxOption) error

type PageRequest struct {
	Page int64
	Size int64
}

type PageResp struct {
	Total int64
	Page  int64
	Size  int64
}
