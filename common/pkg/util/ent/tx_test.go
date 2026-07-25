package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"strings"
	"sync"
	"testing"
)

type mockTx struct {
	mu         sync.Mutex
	committed  bool
	rolledBack bool
	client     string
}

func newMockTx(client string) *mockTx {
	return &mockTx{
		client: client,
	}
}

func (m *mockTx) Commit() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.committed = true
	return nil
}

func (m *mockTx) Rollback() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.rolledBack = true
	return nil
}

func (m *mockTx) Client() string {
	return m.client
}

func (m *mockTx) isCommitted() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.committed
}

func (m *mockTx) isRolledBack() bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.rolledBack
}

type failingTx struct {
	*mockTx
	failCommit   bool
	failRollback bool
}

func (f *failingTx) Commit() error {
	if f.failCommit {
		return errors.New("commit failed")
	}
	return f.mockTx.Commit()
}

func (f *failingTx) Rollback() error {
	if f.failRollback {
		return errors.New("rollback failed")
	}
	return f.mockTx.Rollback()
}

type savableTx struct {
	*mockTx
	mu   sync.Mutex
	sqls []string
}

func newSavableTx(client string) *savableTx {
	return &savableTx{
		mockTx: newMockTx(client),
	}
}

func (s *savableTx) TxExec(_ context.Context, sql string, _ ...driver.Value) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sqls = append(s.sqls, sql)
	return nil
}

func (s *savableTx) getSQLs() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	cp := make([]string, len(s.sqls))
	copy(cp, s.sqls)
	return cp
}

type failingSavableTx struct {
	*savableTx
	failOnPrefix string
}

func (f *failingSavableTx) TxExec(_ context.Context, sql string, _ ...driver.Value) error {
	if f.failOnPrefix != "" && strings.HasPrefix(sql, f.failOnPrefix) {
		return errors.New("savepoint exec failed")
	}
	f.mu.Lock()
	defer f.mu.Unlock()
	f.sqls = append(f.sqls, sql)
	return nil
}

func newStarter(tx Tx[string]) TxStarter[string] {
	return func(_ context.Context) (Tx[string], error) {
		return tx, nil
	}
}

func newFailingStarter(err error) TxStarter[string] {
	return func(_ context.Context) (Tx[string], error) {
		return nil, err
	}
}

func withTx(ctx context.Context, tx Tx[string]) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func TestWithTx_CommitOnSuccess(t *testing.T) {
	tx := newMockTx("db")
	err := WithTx(context.Background(), newStarter(tx), func(ctx context.Context) error {
		c, ok := ClientFromCtx[string](ctx)
		if !ok || c != "db" {
			t.Fatal("expected client in context")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !tx.isCommitted() {
		t.Fatal("tx should be committed")
	}
	if tx.isRolledBack() {
		t.Fatal("tx should not be rolled back")
	}
}

func TestWithTx_RollbackOnError(t *testing.T) {
	tx := newMockTx("db")
	bizErr := errors.New("business error")
	err := WithTx(context.Background(), newStarter(tx), func(ctx context.Context) error {
		return bizErr
	})
	if !errors.Is(err, bizErr) {
		t.Fatalf("expected business error, got: %v", err)
	}
	if tx.isCommitted() {
		t.Fatal("tx should not be committed")
	}
	if !tx.isRolledBack() {
		t.Fatal("tx should be rolled back")
	}
}

func TestWithTx_RollbackOnPanic(t *testing.T) {
	tx := newMockTx("db")
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("expected panic to propagate")
			}
		}()
		_ = WithTx(context.Background(), newStarter(tx), func(ctx context.Context) error {
			panic("boom")
		})
	}()
	if !tx.isRolledBack() {
		t.Fatal("tx should be rolled back after panic")
	}
}

func TestWithTx_RollbackFails(t *testing.T) {
	tx := &failingTx{
		mockTx:       newMockTx("db"),
		failRollback: true,
	}
	err := WithTx(context.Background(), newStarter(tx), func(ctx context.Context) error {
		return errors.New("biz err")
	})
	if err == nil {
		t.Fatal("expected error when rollback fails")
	}
}

func TestWithTx_CommitFails(t *testing.T) {
	tx := &failingTx{
		mockTx:     newMockTx("db"),
		failCommit: true,
	}
	err := WithTx(context.Background(), newStarter(tx), func(ctx context.Context) error {
		return nil
	})
	if err == nil {
		t.Fatal("expected error when commit fails")
	}
}

func TestWithTx_StarterFails(t *testing.T) {
	starterErr := errors.New("connection refused")
	err := WithTx(context.Background(), newFailingStarter(starterErr), func(ctx context.Context) error {
		t.Fatal("fn should not be called")
		return nil
	})
	if !errors.Is(err, starterErr) {
		t.Fatalf("expected starter error, got: %v", err)
	}
}

func TestWithTx_DefaultPropagationIsRequired(t *testing.T) {
	tx := newMockTx("db")
	err := WithTx(context.Background(), newStarter(tx), func(ctx context.Context) error {
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !tx.isCommitted() {
		t.Fatal("default should create and commit tx")
	}
}

func TestRequired_NoExistingTx(t *testing.T) {
	tx := newMockTx("db")
	err := WithTx(context.Background(), newStarter(tx), func(ctx context.Context) error {
		if _, ok := ClientFromCtx[string](ctx); !ok {
			t.Fatal("expected tx in context")
		}
		return nil
	}, WithPropagation(PropagationRequired))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !tx.isCommitted() {
		t.Fatal("tx should be committed")
	}
}

func TestRequired_WithExistingTx_JoinsOuter(t *testing.T) {
	outer := newMockTx("outer")
	innerCalled := false
	starter := func(_ context.Context) (Tx[string], error) {
		innerCalled = true
		return newMockTx("inner"), nil
	}

	ctx := withTx(context.Background(), outer)
	err := WithTx(ctx, starter, func(ctx context.Context) error {
		c, ok := ClientFromCtx[string](ctx)
		if !ok || c != "outer" {
			t.Fatal("expected outer tx client")
		}
		return nil
	}, WithPropagation(PropagationRequired))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if innerCalled {
		t.Fatal("starter should not be called")
	}
}

func TestRequiresNew_WithExistingTx_CreatesNew(t *testing.T) {
	outer := newMockTx("outer")
	inner := newMockTx("inner")

	ctx := withTx(context.Background(), outer)
	err := WithTx(ctx, newStarter(inner), func(ctx context.Context) error {
		c, ok := ClientFromCtx[string](ctx)
		if !ok || c != "inner" {
			t.Fatal("expected inner tx")
		}
		return nil
	}, WithPropagation(PropagationRequiresNew))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !inner.isCommitted() {
		t.Fatal("inner should be committed")
	}
	if outer.isCommitted() || outer.isRolledBack() {
		t.Fatal("outer should not be touched")
	}
}

func TestRequiresNew_InnerFail_OuterUnaffected(t *testing.T) {
	outer := newMockTx("outer")
	inner := newMockTx("inner")

	ctx := withTx(context.Background(), outer)
	err := WithTx(ctx, newStarter(inner), func(ctx context.Context) error {
		return errors.New("inner error")
	}, WithPropagation(PropagationRequiresNew))
	if err == nil {
		t.Fatal("expected error")
	}
	if !inner.isRolledBack() {
		t.Fatal("inner should be rolled back")
	}
	if outer.isRolledBack() {
		t.Fatal("outer should not be rolled back")
	}
}

func TestRequiresNew_NoExistingTx(t *testing.T) {
	tx := newMockTx("db")
	err := WithTx(context.Background(), newStarter(tx), func(ctx context.Context) error {
		return nil
	}, WithPropagation(PropagationRequiresNew))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !tx.isCommitted() {
		t.Fatal("tx should be committed")
	}
}

func TestNested_WithExistingTx_Savepoint(t *testing.T) {
	outer := newSavableTx("outer")

	ctx := withTx(context.Background(), outer)
	err := WithTx(ctx, newStarter(newMockTx("inner")), func(ctx context.Context) error {
		return nil
	}, WithPropagation(PropagationNested))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	sqls := outer.getSQLs()
	if len(sqls) != 2 {
		t.Fatalf("expected 2 SQLs, got %d: %v", len(sqls), sqls)
	}
	if !strings.HasPrefix(sqls[0], "SAVEPOINT sp_") {
		t.Fatalf("expected SAVEPOINT, got: %s", sqls[0])
	}
	if !strings.HasPrefix(sqls[1], "RELEASE SAVEPOINT sp_") {
		t.Fatalf("expected RELEASE, got: %s", sqls[1])
	}
}

func TestNested_InnerFail_RollbackToSavepoint(t *testing.T) {
	outer := newSavableTx("outer")

	ctx := withTx(context.Background(), outer)
	err := WithTx(ctx, newStarter(newMockTx("inner")), func(ctx context.Context) error {
		return errors.New("inner error")
	}, WithPropagation(PropagationNested))
	if err == nil {
		t.Fatal("expected error")
	}
	sqls := outer.getSQLs()
	if len(sqls) != 3 {
		t.Fatalf("expected 3 SQLs, got %d: %v", len(sqls), sqls)
	}
	if !strings.HasPrefix(sqls[1], "ROLLBACK TO SAVEPOINT sp_") {
		t.Fatalf("expected ROLLBACK TO, got: %s", sqls[1])
	}
	if !strings.HasPrefix(sqls[2], "RELEASE SAVEPOINT sp_") {
		t.Fatalf("expected RELEASE, got: %s", sqls[2])
	}
	if outer.isCommitted() || outer.isRolledBack() {
		t.Fatal("outer tx should not be committed or rolled back")
	}
}

func TestNested_NoExistingTx_CreatesTx(t *testing.T) {
	tx := newMockTx("db")
	err := WithTx(context.Background(), newStarter(tx), func(ctx context.Context) error {
		return nil
	}, WithPropagation(PropagationNested))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !tx.isCommitted() {
		t.Fatal("tx should be committed")
	}
}

func TestNested_NonSaver_Fallback(t *testing.T) {
	outer := newMockTx("outer")
	ctx := withTx(context.Background(), outer)

	err := WithTx(ctx, newStarter(newMockTx("inner")), func(ctx context.Context) error {
		return nil
	}, WithPropagation(PropagationNested))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestNested_SavepointCreateFails(t *testing.T) {
	outer := &failingSavableTx{
		savableTx:    newSavableTx("outer"),
		failOnPrefix: "SAVEPOINT",
	}

	ctx := withTx(context.Background(), outer)
	err := WithTx(ctx, newStarter(newMockTx("inner")), func(ctx context.Context) error {
		return nil
	}, WithPropagation(PropagationNested))
	if err == nil {
		t.Fatal("expected error when savepoint creation fails")
	}
	if !strings.Contains(err.Error(), "savepoint") {
		t.Fatalf("expected savepoint error, got: %v", err)
	}
}

func TestNotSupported_WithExistingTx_Suspends(t *testing.T) {
	outer := newMockTx("outer")
	fnCalled := false

	ctx := withTx(context.Background(), outer)
	err := WithTx(ctx, newStarter(newMockTx("inner")), func(ctx context.Context) error {
		fnCalled = true
		return nil
	}, WithPropagation(PropagationNotSupported))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !fnCalled {
		t.Fatal("fn should be called")
	}
	if outer.isCommitted() || outer.isRolledBack() {
		t.Fatal("outer should not be touched")
	}
}

func TestNotSupported_InnerError_Propagates(t *testing.T) {
	outer := newMockTx("outer")
	ctx := withTx(context.Background(), outer)

	err := WithTx(ctx, newStarter(newMockTx("inner")), func(ctx context.Context) error {
		return errors.New("error")
	}, WithPropagation(PropagationNotSupported))
	if err == nil {
		t.Fatal("expected error")
	}
	if outer.isRolledBack() {
		t.Fatal("outer should not be rolled back")
	}
}

func TestNotSupported_NoExistingTx(t *testing.T) {
	fnCalled := false
	err := WithTx(context.Background(), newStarter(newMockTx("db")), func(ctx context.Context) error {
		fnCalled = true
		if _, ok := ClientFromCtx[string](ctx); ok {
			t.Fatal("should not have tx in context")
		}
		return nil
	}, WithPropagation(PropagationNotSupported))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !fnCalled {
		t.Fatal("fn should be called")
	}
}

func TestNever_WithExistingTx_Error(t *testing.T) {
	outer := newMockTx("outer")
	ctx := withTx(context.Background(), outer)

	err := WithTx(ctx, newStarter(newMockTx("inner")), func(ctx context.Context) error {
		t.Fatal("fn should not be called")
		return nil
	}, WithPropagation(PropagationNever))
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "never") {
		t.Fatalf("expected 'never' in error, got: %v", err)
	}
}

func TestNever_NoExistingTx_Runs(t *testing.T) {
	fnCalled := false
	err := WithTx(context.Background(), newStarter(newMockTx("db")), func(ctx context.Context) error {
		fnCalled = true
		return nil
	}, WithPropagation(PropagationNever))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !fnCalled {
		t.Fatal("fn should be called")
	}
}

func TestSupports_WithExistingTx_Joins(t *testing.T) {
	outer := newMockTx("outer")
	starterCalled := false

	ctx := withTx(context.Background(), outer)
	err := WithTx(ctx, func(_ context.Context) (Tx[string], error) {
		starterCalled = true
		return newMockTx("inner"), nil
	}, func(ctx context.Context) error {
		c, ok := ClientFromCtx[string](ctx)
		if !ok || c != "outer" {
			t.Fatal("expected outer tx")
		}
		return nil
	}, WithPropagation(PropagationSupports))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if starterCalled {
		t.Fatal("starter should not be called")
	}
}

func TestSupports_NoExistingTx_RunsWithoutTx(t *testing.T) {
	tx := newMockTx("db")
	err := WithTx(context.Background(), newStarter(tx), func(ctx context.Context) error {
		if _, ok := ClientFromCtx[string](ctx); ok {
			t.Fatal("should not have tx in context")
		}
		return nil
	}, WithPropagation(PropagationSupports))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tx.isCommitted() {
		t.Fatal("tx should not be committed")
	}
}

func TestNestedScenario_RequiredFlatten(t *testing.T) {
	outer := newMockTx("outer")
	err := WithTx(context.Background(), newStarter(outer), func(ctx context.Context) error {
		return WithTx(ctx, newStarter(newMockTx("inner")), func(ctx context.Context) error {
			c, _ := ClientFromCtx[string](ctx)
			if c != "outer" {
				t.Fatal("nested Required should use outer tx")
			}
			return nil
		}, WithPropagation(PropagationRequired))
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !outer.isCommitted() {
		t.Fatal("outer should be committed")
	}
}

func TestNestedScenario_RequiresNewIsolation(t *testing.T) {
	outer := newMockTx("outer")
	inner := newMockTx("inner")

	err := WithTx(context.Background(), newStarter(outer), func(ctx context.Context) error {
		return WithTx(ctx, newStarter(inner), func(ctx context.Context) error {
			c, _ := ClientFromCtx[string](ctx)
			if c != "inner" {
				t.Fatal("RequiresNew should use inner tx")
			}
			return nil
		}, WithPropagation(PropagationRequiresNew))
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !outer.isCommitted() || !inner.isCommitted() {
		t.Fatal("both should be committed")
	}
}

func TestNestedScenario_RequiresNew_InnerFailOuterSucceeds(t *testing.T) {
	outer := newMockTx("outer")
	inner := newMockTx("inner")

	err := WithTx(context.Background(), newStarter(outer), func(ctx context.Context) error {
		_ = WithTx(ctx, newStarter(inner), func(ctx context.Context) error {
			return errors.New("inner fail")
		}, WithPropagation(PropagationRequiresNew))
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !outer.isCommitted() {
		t.Fatal("outer should be committed")
	}
	if !inner.isRolledBack() {
		t.Fatal("inner should be rolled back")
	}
}

func TestNestedScenario_NestedSavepoint_PartialRollback(t *testing.T) {
	outer := newSavableTx("outer")

	err := WithTx(context.Background(), newStarter(outer), func(ctx context.Context) error {
		_ = WithTx(ctx, newStarter(newMockTx("inner")), func(ctx context.Context) error {
			return errors.New("inner fail")
		}, WithPropagation(PropagationNested))
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !outer.isCommitted() {
		t.Fatal("outer should be committed")
	}
	sqls := outer.getSQLs()
	hasRollbackTo := false
	for _, sql := range sqls {
		if strings.HasPrefix(sql, "ROLLBACK TO SAVEPOINT") {
			hasRollbackTo = true
		}
	}
	if !hasRollbackTo {
		t.Fatalf("expected ROLLBACK TO SAVEPOINT: %v", sqls)
	}
}

func TestNestedScenario_ThreeLevelsDeep(t *testing.T) {
	l1 := newSavableTx("l1")
	l2 := newSavableTx("l2")

	err := WithTx(context.Background(), newStarter(l1), func(ctx context.Context) error {
		return WithTx(ctx, newStarter(l2), func(ctx context.Context) error {
			return WithTx(ctx, newStarter(newMockTx("unused")), func(ctx context.Context) error {
				c, _ := ClientFromCtx[string](ctx)
				if c != "l2" {
					t.Fatal("expected l2 client from ctx")
				}
				return nil
			}, WithPropagation(PropagationNested))
		}, WithPropagation(PropagationRequiresNew))
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !l1.isCommitted() {
		t.Fatal("l1 should be committed")
	}
	if !l2.isCommitted() {
		t.Fatal("l2 should be committed")
	}
	sqls := l2.getSQLs()
	if len(sqls) < 2 {
		t.Fatalf("l2 should have savepoint SQLs, got: %v", sqls)
	}
}

func TestClientFromCtx_NoTx(t *testing.T) {
	c, ok := ClientFromCtx[string](context.Background())
	if ok {
		t.Fatal("should not find client")
	}
	if c != "" {
		t.Fatal("zero value expected")
	}
}

type intClient struct{ v int }

func (i intClient) Commit() error {
	return nil
}

func (i intClient) Rollback() error {
	return nil
}

func (i intClient) Client() int {
	return i.v
}

func TestClientFromCtx_TypeMismatch(t *testing.T) {
	tx := intClient{
		v: 42,
	}
	ctx := context.WithValue(context.Background(), txKey{}, tx)
	c, ok := ClientFromCtx[string](ctx)
	if ok {
		t.Fatal("type mismatch should return false")
	}
	if c != "" {
		t.Fatal("zero value expected")
	}
}

func TestClientFromCtx_WithSuspendedTx(t *testing.T) {
	suspended := suspendedTx[string]{
		tx: newMockTx("suspended"),
	}
	ctx := context.WithValue(context.Background(), txKey{}, suspended)
	c, ok := ClientFromCtx[string](ctx)
	if !ok || c != "suspended" {
		t.Fatal("should extract client from suspended tx")
	}
}

func TestWithPropagation_SetsOption(t *testing.T) {
	tx := newMockTx("db")
	fnCalled := false

	err := WithTx(context.Background(), newStarter(tx), func(ctx context.Context) error {
		fnCalled = true
		return nil
	}, WithPropagation(PropagationNever))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !fnCalled {
		t.Fatal("fn should be called")
	}
}
