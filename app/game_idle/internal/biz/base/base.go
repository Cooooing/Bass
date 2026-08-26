package base

import (
	"context"

	utilent "common/pkg/util/ent"
)

// Tx 事务执行器，行动结算需要在同一个本地事务里完成条件扣减、奖励入账和队列推进。
type Tx func(ctx context.Context, fn func(ctx context.Context) error, opts ...utilent.TxOption) error
