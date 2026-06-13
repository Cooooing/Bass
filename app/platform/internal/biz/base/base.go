package base

import (
	"context"

	utilent "common/pkg/util/ent"
)

// Tx 事务执行器，支持事务传播行为
type Tx func(ctx context.Context, fn func(ctx context.Context) error, opts ...utilent.TxOption) error
