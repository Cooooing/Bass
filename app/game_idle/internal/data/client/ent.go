package client

import (
	commonClient "common/pkg/client"
	"common/pkg/client/driver"
	"common/proto/gen/common"
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"game_idle/internal/data/gen"
	"game_idle/internal/data/gen/migrate"

	_ "github.com/lib/pq" // register PostgreSQL driver

	_ "game_idle/internal/data/gen/runtime" // register Ent runtime hooks

	entsql "entgo.io/ent/dialect/sql"
)

// EntClient 管理模块的 Ent 客户端。
type EntClient struct {
	Client         *gen.Client
	logger         *slog.Logger
	databaseDriver string
	databaseConfig driver.Config
}

func NewEntClient(logger *slog.Logger, conf *common.Database, sqlClient *commonClient.SQLClient) (*EntClient, error) {
	sqlDriver := entsql.OpenDB(sqlClient.Driver, sqlClient.DB)
	client := gen.NewClient(gen.Driver(driver.Wrap(sqlDriver, logger, conf)))
	if conf.GetMerge() {
		if err := client.Schema.Create(context.Background(), migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
			return nil, fmt.Errorf("create schema resources: %w", err)
		}
	}
	return &EntClient{Client: client, logger: logger, databaseDriver: sqlClient.Driver, databaseConfig: conf}, nil
}

func (c *EntClient) ClientFromTx(tx *sql.Tx) *gen.Client {
	sqlDriver := entsql.NewDriver(c.databaseDriver, entsql.Conn{ExecQuerier: tx})
	return gen.NewClient(gen.Driver(driver.Wrap(sqlDriver, c.logger, c.databaseConfig)))
}
