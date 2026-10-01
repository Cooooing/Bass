package base

import (
	utilent "common/pkg/util/ent"
	"context"
)

type Tx func(ctx context.Context, fn func(ctx context.Context) error, opts ...utilent.TxOption) error
