package client

import (
	"content/internal/data/gen"
	"context"

	utilent "common/pkg/util/ent"
)

type genTxWrapper struct{ tx *gen.Tx }

func (w *genTxWrapper) Commit() error       { return w.tx.Commit() }
func (w *genTxWrapper) Rollback() error     { return w.tx.Rollback() }
func (w *genTxWrapper) Client() interface{} { return w.tx.Client() }

func WithTx(ctx context.Context, client *gen.Client, fn func(tx *gen.Client) error, opts ...utilent.TxOption) error {
	starter := func(ctx context.Context) (utilent.Tx, error) {
		tx, err := client.Tx(ctx)
		if err != nil {
			return nil, err
		}
		return &genTxWrapper{tx: tx}, nil
	}

	return utilent.WithTx(ctx, starter, func(ctx context.Context) error {
		c, ok := utilent.ClientFromCtx[*gen.Client](ctx)
		if !ok {
			return fn(client)
		}
		return fn(c)
	}, opts...)
}
