package client

import (
	"common/pkg/client/driver"
	"common/pkg/constant"
	"context"
	"fmt"
	"log/slog"
	"platform/internal/config"
	"platform/internal/data/gen"
	"platform/internal/data/gen/migrate"

	_ "platform/internal/data/gen/runtime"

	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
)

func NewDataBaseClient(logger *slog.Logger, conf *config.Bootstrap) (*gen.Client, func(), error) {
	drv, err := sql.Open(conf.Database.Driver, conf.Database.Source)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open db: %w", err)
	}
	observedDrv := driver.Wrap(drv, logger, conf.Database)
	client := gen.NewClient(gen.Driver(observedDrv))
	logger.Info("database client created", constant.LogFieldKind, constant.LogKindDatabase, constant.LogFieldDriver, conf.Database.Driver)
	if conf.Database.Merge {
		if err := client.Schema.Create(context.Background(), migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
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
