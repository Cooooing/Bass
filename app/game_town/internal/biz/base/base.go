package base

import (
	"context"
)

type Tx func(ctx context.Context, fn func(ctx context.Context) error) error

type PageRequest struct {
	Page int64
	Size int64
}

type PageResp struct {
	Page  int64
	Size  int64
	Total int64
}
