package client

import (
	"context"
	"database/sql"
	"fmt"

	utilent "common/pkg/util/ent"

	"github.com/dtm-labs/client/dtmcli"
)

// BarrierStore 提供屏障事务所需的数据库连接与方言。
type BarrierStore struct {
	DB     *sql.DB
	DBType string // dtmcli.DBTypeMysql / dtmcli.DBTypePostgres
}

// TCCBranchExecutor 在屏障事务中执行 TCC 分支业务。
type TCCBranchExecutor interface {
	Execute(ctx context.Context, fn func(context.Context) error) error
}

type barrierExecutor[C any] struct {
	dtm          *DTMClient
	store        *BarrierStore
	clientFromTx func(*sql.Tx) C
}

// NewTCCBranchExecutor 创建 TCC 分支执行器。
func NewTCCBranchExecutor[C any](dtm *DTMClient, db *sql.DB, dbType string, clientFromTx func(*sql.Tx) C) (TCCBranchExecutor, error) {
	if db == nil {
		return nil, fmt.Errorf("barrier database is required")
	}
	if dbType == dtmcli.DBTypePostgres {
		if _, err := db.ExecContext(context.Background(), fmt.Sprintf(
			`DO $$
	DECLARE
		constraint_name text;
	BEGIN
		SELECT c.conname INTO constraint_name
		FROM pg_constraint AS c
		WHERE c.conrelid = '%[1]s'::regclass
			AND c.contype = 'u'
			AND c.conname = '%[2]s';

		IF constraint_name IS NULL THEN
			SELECT c.conname INTO constraint_name
			FROM pg_constraint AS c
			JOIN pg_class AS i ON i.oid = c.conindid
			WHERE c.conrelid = '%[1]s'::regclass
				AND c.contype = 'u'
				AND i.relname = '%[2]s';
		END IF;

		IF constraint_name IS NULL THEN
			ALTER TABLE %[1]s ADD CONSTRAINT %[2]s UNIQUE USING INDEX %[2]s;
		ELSIF constraint_name <> '%[2]s' THEN
			EXECUTE format('ALTER TABLE %%I RENAME CONSTRAINT %%I TO %%I', '%[1]s', constraint_name, '%[2]s');
		END IF;
	EXCEPTION WHEN duplicate_object THEN NULL;
	END $$`,
			utilent.DTMBarrierTableName,
			utilent.DTMBarrierConstraintName,
		)); err != nil {
			return nil, fmt.Errorf("ensure barrier constraint: %w", err)
		}
	}
	return &barrierExecutor[C]{
		dtm:          dtm,
		store:        &BarrierStore{DB: db, DBType: dbType},
		clientFromTx: clientFromTx,
	}, nil
}

func (e *barrierExecutor[C]) Execute(ctx context.Context, fn func(context.Context) error) error {
	if e == nil || e.dtm == nil || e.clientFromTx == nil || fn == nil {
		return fmt.Errorf("tcc branch executor invalid")
	}
	return e.dtm.BarrierCall(ctx, e.store, func(tx *sql.Tx) error {
		return fn(utilent.WithClient(ctx, e.clientFromTx(tx)))
	})
}
