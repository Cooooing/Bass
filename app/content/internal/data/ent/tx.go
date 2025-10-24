package ent

import (
	"content/internal/data/ent/gen"
	"context"
)

func WithTx(ctx context.Context, client *gen.Client, fn func(tx *gen.Client) error) (err error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	txClient := tx.Client()

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback() // 尝试回滚
			panic(p)          // 继续抛出 panic
		} else if err != nil {
			_ = tx.Rollback() // 出错时回滚
		} else {
			err = tx.Commit() // 成功时提交
		}
	}()

	err = fn(txClient)
	return err
}
