package client

import (
	"common/pkg/client/driver"
	"common/pkg/constant"
	"context"
	"fmt"
	"notify/internal/conf"
	"notify/internal/data/gen"
	"notify/internal/data/gen/migrate"

	_ "notify/internal/data/gen/runtime"

	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"log/slog"
)

func NewDataBaseClient(logger *slog.Logger, conf *conf.Bootstrap) (*gen.Client, func(), error) {
	drv, err := sql.Open(conf.Data.Database.Driver, conf.Data.Database.Source)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open db: %w", err)
	}
	observedDrv := driver.Wrap(drv, logger, conf.Data.Database)
	client := gen.NewClient(gen.Driver(observedDrv))
	logger.Info("database client created", constant.LogFieldKind, constant.LogKindDatabase, constant.LogFieldDriver, conf.Data.Database.Driver)
	// 可选：自动迁移
	if conf.Data.Database.Merge {
		ctx := context.Background()
		if err := client.Schema.Create(ctx, migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
			return nil, nil, fmt.Errorf("failed creating schema resources: %w", err)
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
