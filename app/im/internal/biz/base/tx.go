package base

import (
	utilent "common/pkg/util/ent"
	"context"
)

type Tx func(context.Context, func(context.Context) error, ...utilent.TxOption) error
