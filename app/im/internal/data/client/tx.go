package client

import (
	utilent "common/pkg/util/ent"
	"context"

	"im/internal/biz/base"
	"im/internal/data/gen"
)

type txWrapper struct{ tx *gen.Tx }

func (w *txWrapper) Commit() error       { return w.tx.Commit() }
func (w *txWrapper) Rollback() error     { return w.tx.Rollback() }
func (w *txWrapper) Client() *gen.Client { return w.tx.Client() }

func ProvideTx(db *gen.Client) base.Tx {
	starter := func(ctx context.Context) (utilent.Tx[*gen.Client], error) {
		tx, err := db.Tx(ctx)
		if err != nil {
			return nil, err
		}
		return &txWrapper{tx: tx}, nil
	}
	return func(ctx context.Context, fn func(context.Context) error, opts ...utilent.TxOption) error {
		return utilent.WithTx(ctx, starter, fn, opts...)
	}
}
