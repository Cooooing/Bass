package client

import (
	utilent "common/pkg/util/ent"
	"context"
	"platform/internal/biz/base"
	"platform/internal/data/gen"
)

type genTxWrapper struct{ tx *gen.Tx }

func (w *genTxWrapper) Commit() error {
	return w.tx.Commit()
}

func (w *genTxWrapper) Rollback() error {
	return w.tx.Rollback()
}

func (w *genTxWrapper) Client() interface{} {
	return w.tx.Client()
}

func ProvideTx(
	db *gen.Client,
) base.Tx {
	starter := func(ctx context.Context) (utilent.Tx, error) {
		tx, err := db.Tx(ctx)
		if err != nil {
			return nil, err
		}
		return &genTxWrapper{
			tx: tx,
		}, nil
	}
	return func(ctx context.Context, fn func(ctx context.Context) error, opts ...utilent.TxOption) error {
		return utilent.WithTx(ctx, starter, fn, opts...)
	}
}
