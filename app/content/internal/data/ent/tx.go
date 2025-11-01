package ent

import (
	"content/internal/data/ent/gen"
	"context"
	"errors"
	"fmt"
)

func WithTx(ctx context.Context, client *gen.Client, fn func(tx *gen.Client) error) error {
	tx, err := client.Tx(ctx)
	// 如果已经存在事务，则直接使用传入的client执行
	if errors.Is(err, gen.ErrTxStarted) {
		return fn(client)
	}
	if err != nil {
		return errors.Join(err, fmt.Errorf("create tx failed"))
	}

	defer func() {
		if p := recover(); p != nil {
			// 如果发生 panic，尝试回滚并继续抛出
			if rbErr := tx.Rollback(); rbErr != nil {
				panic(errors.Join(rbErr, fmt.Errorf("tx panic: %v", p)))
			}
			panic(p)
		}
	}()

	// 执行业务逻辑
	if err := fn(tx.Client()); err != nil {
		// 出错时回滚
		if rbErr := tx.Rollback(); rbErr != nil {
			return errors.Join(rbErr, fmt.Errorf("tx commit failed: %v", err))
		}
		return err
	}

	// 正常提交
	if err := tx.Commit(); err != nil {
		return errors.Join(err, fmt.Errorf("tx commit failed"))
	}

	return nil
}
