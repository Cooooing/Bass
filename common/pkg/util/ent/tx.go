package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

type txKey struct{}

// Propagation 事务传播行为
type Propagation int

const (
	// PropagationRequired 默认：如果当前存在事务则加入，否则创建新事务。
	PropagationRequired Propagation = iota
	// PropagationRequiresNew 总是创建新事务，挂起当前事务。
	PropagationRequiresNew
	// PropagationNested 如果当前存在事务则创建 savepoint，否则等同于 Required。
	PropagationNested
	// PropagationNotSupported 以非事务方式运行，挂起当前事务。
	PropagationNotSupported
	// PropagationNever 以非事务方式运行，如果当前存在事务则报错。
	PropagationNever
	// PropagationSupports 如果当前存在事务则加入，否则以非事务方式运行。
	PropagationSupports
)

// TxStarter 创建事务的函数，由各模块提供。
type TxStarter[C any] func(ctx context.Context) (Tx[C], error)

// Tx 事务操作接口，各模块的 *gen.Tx 均通过 wrapper 实现。
type Tx[C any] interface {
	Commit() error
	Rollback() error
	Client() C
}

type suspendedTx[C any] struct {
	tx Tx[C]
}

func (s suspendedTx[C]) Commit() error {
	return nil
}

func (s suspendedTx[C]) Rollback() error {
	return nil
}

func (s suspendedTx[C]) Client() C {
	return s.tx.Client()
}

// WithTx 开启事务，支持事务传播。
func WithTx[C any](ctx context.Context, starter TxStarter[C], fn func(ctx context.Context) error, opts ...TxOption) error {
	cfg := &TxOptionConfig{
		Propagation: PropagationRequired,
	}
	for _, o := range opts {
		o(cfg)
	}
	return doWithTx(ctx, starter, fn, cfg)
}

// ClientFromCtx 从事务上下文中提取 typed client。
func ClientFromCtx[C any](ctx context.Context) (C, bool) {
	tx, ok := ctx.Value(txKey{}).(Tx[C])
	if !ok {
		var zero C
		return zero, false
	}
	return tx.Client(), true
}

func doWithTx[C any](ctx context.Context, starter TxStarter[C], fn func(ctx context.Context) error, cfg *TxOptionConfig) error {
	currentTx, hasTx := ctx.Value(txKey{}).(Tx[C])

	switch cfg.Propagation {
	case PropagationRequired:
		if hasTx {
			return fn(ctx)
		}
		return startTx(ctx, starter, fn)

	case PropagationRequiresNew:
		if hasTx {
			suspended := context.WithValue(ctx, txKey{}, suspendedTx[C]{
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
			clean := context.WithValue(ctx, txKey{}, suspendedTx[C]{
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

func startTx[C any](ctx context.Context, starter TxStarter[C], fn func(ctx context.Context) error) error {
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

func startSavepoint[C any](ctx context.Context, tx Tx[C], fn func(ctx context.Context) error) error {
	type saver interface {
		TxExec(ctx context.Context, sql string, args ...driver.Value) error
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

// TxOption 事务选项。
type TxOption func(*TxOptionConfig)

// TxOptionConfig 事务选项配置。
type TxOptionConfig struct {
	Propagation Propagation
}

// WithPropagation 设置事务传播行为。
func WithPropagation(p Propagation) TxOption {
	return func(c *TxOptionConfig) {
		c.Propagation = p
	}
}
