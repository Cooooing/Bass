package client

import (
	bizbase "content/internal/biz/base"
	"content/internal/data/gen"
	"context"

	utilent "common/pkg/util/ent"
)

// genTxWrapper 将 gen.Tx 适配为 utilent.Tx 接口
type genTxWrapper struct{ tx *gen.Tx }

func (w *genTxWrapper) Commit() error       { return w.tx.Commit() }
func (w *genTxWrapper) Rollback() error     { return w.tx.Rollback() }
func (w *genTxWrapper) Client() interface{} { return w.tx.Client() }

// ProvideTx 提供事务执行器，供 usecase 层使用
func ProvideTx(db *gen.Client) bizbase.Tx {
	starter := func(ctx context.Context) (utilent.Tx, error) {
		tx, err := db.Tx(ctx)
		if err != nil {
			return nil, err
		}
		return &genTxWrapper{tx: tx}, nil
	}
	return func(ctx context.Context, fn func(ctx context.Context) error, opts ...utilent.TxOption) error {
		return utilent.WithTx(ctx, starter, fn, opts...)
	}
}
