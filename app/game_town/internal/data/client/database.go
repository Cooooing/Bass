package client

import (
	"context"
	dbsql "database/sql"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"common/pkg/client/driver"
	"common/pkg/constant"
	"game_town/internal/config"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/migrate"
	_ "game_town/internal/data/gen/runtime"

	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

func checkPgvector(
	ctx context.Context,
	drv *sql.Driver,
) (string, error) {
	var version string
	err := drv.DB().QueryRowContext(ctx, "SELECT extversion FROM pg_extension WHERE extname = 'vector'").Scan(&version)
	if err == nil {
		return version, nil
	}
	return "", pgvectorCheckError(err)
}

func pgvectorCheckError(
	err error,
) error {
	if errors.Is(err, dbsql.ErrNoRows) {
		return fmt.Errorf("pgvector extension is not initialized; initialize it once as a database administrator with CREATE EXTENSION IF NOT EXISTS vector")
	}
	return fmt.Errorf("failed checking pgvector extension: %w", err)
}

func NewDataBaseClient(
	logger *slog.Logger,
	conf *config.Bootstrap,
) (*gen.Client, func(), error) {
	drv, err := sql.Open(conf.Database.Driver, conf.Database.Source)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open db: %w", err)
	}
	drv.DB().SetMaxOpenConns(20)
	drv.DB().SetMaxIdleConns(10)
	drv.DB().SetConnMaxLifetime(30 * time.Minute)
	drv.DB().SetConnMaxIdleTime(5 * time.Minute)

	ctx := context.Background()
	vectorVersion, err := checkPgvector(ctx, drv)
	if err != nil {
		return nil, nil, err
	}
	observedDrv := driver.Wrap(drv, logger, conf.Database)
	client := gen.NewClient(gen.Driver(observedDrv))
	logger.Info("database client created", constant.LogFieldKind, constant.LogKindDatabase, constant.LogFieldDriver, conf.Database.Driver, "pgvector_version", vectorVersion)
	if conf.Database.Merge {
		if err = client.Schema.Create(ctx, migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
			return nil, nil, fmt.Errorf("failed creating schema resources: %w", err)
		}
		if _, err = drv.DB().ExecContext(ctx, `CREATE INDEX IF NOT EXISTS game_town_npc_memories_embedding_hnsw_idx
			ON game_town_npc_memories USING hnsw (embedding vector_cosine_ops)
			WHERE embedding IS NOT NULL`); err != nil {
			return nil, nil, fmt.Errorf("failed creating npc memory vector index: %w", err)
		}
	}
	cleanup := func() {
		if err := client.Close(); err != nil {
			logger.Error("close ent client failed", constant.LogFieldKind, constant.LogKindDatabase, constant.LogFieldErr, err)
		}
		if err := observedDrv.Close(); err != nil {
			logger.Error("close db driver failed", constant.LogFieldKind, constant.LogKindDatabase, constant.LogFieldErr, err)
		}
		logger.Info("database client closed", constant.LogFieldKind, constant.LogKindDatabase)
	}
	return client, cleanup, nil
}
