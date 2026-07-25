package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type txKey struct{}

// Propagation 事务传播行为
type Propagation int

const (
	// PropagationRequired（默认）如果当前存在事务则加入，否则创建新事务
	PropagationRequired Propagation = iota
	// PropagationRequiresNew 总是创建新事务，挂起当前事务
	PropagationRequiresNew
	// PropagationNested 如果当前存在事务则创建 savepoint（支持部分回滚），否则等同于 PropagationRequired
	PropagationNested
	// PropagationNotSupported 以非事务方式运行，挂起当前事务
	PropagationNotSupported
	// PropagationNever 以非事务方式运行，如果当前存在事务则抛出异常
	PropagationNever
	// PropagationSupports 如果当前存在事务则加入，否则以非事务方式运行
	PropagationSupports
)

// TxStarter 创建事务的函数，由各模块提供
type TxStarter func(ctx context.Context) (Tx, error)

// Tx 事务操作接口，各模块的 *gen.Tx 均满足此接口
type Tx interface {
	Commit() error
	Rollback() error
	Client() interface{}
}

// suspendedTx 挂起的事务上下文
type suspendedTx struct {
	tx Tx
}

func (s suspendedTx) Commit() error {
	return nil
}

func (s suspendedTx) Rollback() error {
	return nil
}

func (s suspendedTx) Client() interface{} {
	return s.tx.Client()
}

// WithTx 开启事务，支持事务传播
func WithTx(
	ctx context.Context,
	starter TxStarter,
	fn func(ctx context.Context) error,
	opts ...TxOption,
) error {
	cfg := &TxOptionConfig{
		Propagation: PropagationRequired,
	}
	for _, o := range opts {
		o(cfg)
	}
	return doWithTx(ctx, starter, fn, cfg)
}

// ClientFromCtx 从事务上下文中提取 typed client
func ClientFromCtx[C any](
	ctx context.Context,
) (C, bool) {
	tx, ok := ctx.Value(txKey{}).(Tx)
	if !ok {
		var zero C
		return zero, false
	}
	c, ok := tx.Client().(C)
	return c, ok
}

func doWithTx(
	ctx context.Context,
	starter TxStarter,
	fn func(ctx context.Context) error,
	cfg *TxOptionConfig,
) error {
	currentTx, hasTx := ctx.Value(txKey{}).(Tx)

	switch cfg.Propagation {
	case PropagationRequired:
		if hasTx {
			return fn(ctx)
		}
		return startTx(ctx, starter, fn)

	case PropagationRequiresNew:
		if hasTx {
			suspended := context.WithValue(ctx, txKey{}, suspendedTx{
				tx: currentTx,
			})
			return startTx(suspended, starter, fn)
		}
		return startTx(ctx, starter, fn)

	case PropagationNested:
		if hasTx {
			return startSavepoint(ctx, currentTx, fn)
		}
		return startTx(ctx, starter, fn)

	case PropagationNotSupported:
		if hasTx {
			clean := context.WithValue(ctx, txKey{}, suspendedTx{
				tx: currentTx,
			})
			err := fn(clean)
			_ = context.WithValue(ctx, txKey{}, currentTx)
			return err
		}
		return fn(ctx)

	case PropagationNever:
		if hasTx {
			return fmt.Errorf("tx propagation never: existing transaction found")
		}
		return fn(ctx)

	case PropagationSupports:
		return fn(ctx)

	default:
		return fn(ctx)
	}
}

func startTx(
	ctx context.Context,
	starter TxStarter,
	fn func(ctx context.Context) error,
) error {
	tx, err := starter(ctx)
	if err != nil {
		return errors.Join(err, fmt.Errorf("create tx failed"))
	}

	ctx = context.WithValue(ctx, txKey{}, tx)

	defer func() {
		if p := recover(); p != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				panic(errors.Join(rbErr, fmt.Errorf("tx panic rollback failed: %v", p)))
			}
			panic(p)
		}
	}()

	if err := fn(ctx); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return errors.Join(rbErr, fmt.Errorf("tx rollback failed: %v", err))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return errors.Join(err, fmt.Errorf("tx commit failed"))
	}

	return nil
}

// startSavepoint 在已有事务内创建 savepoint，支持部分回滚
func startSavepoint(
	ctx context.Context,
	tx Tx,
	fn func(ctx context.Context) error,
) error {
	type saver interface {
		TxExec(ctx context.Context, sql string, args ...any) error
	}
	if s, ok := tx.(saver); ok {
		sp := fmt.Sprintf("sp_%d", time.Now().UnixNano())
		if err := s.TxExec(ctx, "SAVEPOINT "+sp); err != nil {
			return fmt.Errorf("create savepoint failed: %w", err)
		}

		err := fn(ctx)
		if err != nil {
			_ = s.TxExec(ctx, "ROLLBACK TO SAVEPOINT "+sp)
			_ = s.TxExec(ctx, "RELEASE SAVEPOINT "+sp)
			return err
		}

		if rbErr := s.TxExec(ctx, "RELEASE SAVEPOINT "+sp); rbErr != nil {
			return fmt.Errorf("release savepoint failed: %w", rbErr)
		}
		return nil
	}

	return fn(ctx)
}

// TxOption 事务选项
type TxOption func(*TxOptionConfig)

// TxOptionConfig 事务选项配置
type TxOptionConfig struct {
	Propagation Propagation
}

// WithPropagation 设置事务传播行为
func WithPropagation(
	p Propagation,
) TxOption {
	return func(c *TxOptionConfig) {
		c.Propagation = p
	}
}
