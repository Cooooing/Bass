package client

import (
	"common/pkg/client/db/driver"
	"common/pkg/util"
	"context"
	"fmt"
	"im/internal/conf"
	"im/internal/data/ent/gen"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/lib/pq"
)

func NewDataBaseClient(logger log.Logger, conf *conf.Bootstrap) (*gen.Client, func(), error) {
	l := log.NewHelper(logger)
	drv, err := sql.Open(conf.Data.Database.Driver, conf.Data.Database.Source)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open db: %w", err)
	}
	debugDrv := driver.NewDriver(logger, conf.Server.Mode, drv, conf.Data.Database)
	client := gen.NewClient(gen.Driver(debugDrv))
	l.Infof("database: ent created database client [%s]", conf.Data.Database.Driver)
	// 可选：自动迁移
	if conf.Data.Database.Merge {
		ctx := context.Background()
		if err := client.Schema.Create(ctx); err != nil {
			return nil, nil, fmt.Errorf("failed creating schema resources: %w", err)
		}
	}

	// 注册审计 Hook
	client.Use(util.AuditHook())

	cleanup := func() {
		if err := client.Close(); err != nil {
			l.Errorf("failed to close ent client: %v", err)
		}
		if err := debugDrv.Close(); err != nil {
			l.Errorf("failed to close db driver: %v", err)
		}
		l.Infof("database client closed")
	}
	return client, cleanup, nil
}
