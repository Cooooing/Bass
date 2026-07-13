package base

import (
	"context"

	utilent "common/pkg/util/ent"
)

type Tx func(ctx context.Context, fn func(ctx context.Context) error, opts ...utilent.TxOption) error
