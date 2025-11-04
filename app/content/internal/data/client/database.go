package client

import (
	"common/pkg"
	"content/internal/conf"
	"content/internal/data/ent/gen"
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/lib/pq"
)

func NewDataBaseClient(log *log.Helper, conf *conf.Bootstrap) (*gen.Client, func(), error) {
	drv, err := sql.Open(conf.Data.Database.Driver, conf.Data.Database.Source)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open db: %w", err)
	}
	debugDrv := dialect.DebugWithContext(drv, func(ctx context.Context, args ...any) {
		text := fmt.Sprint(args...)
		log.WithContext(ctx).Debugf("%s", text)
	})
	client := gen.NewClient(gen.Driver(debugDrv))
	log.Infof("database: ent created database client [%s]", conf.Data.Database.Driver)
	// 可选：自动迁移
	if conf.Data.Database.Merge {
		ctx := context.Background()
		if err := client.Schema.Create(ctx); err != nil {
			return nil, nil, fmt.Errorf("failed creating schema resources: %w", err)
		}
	}

	// 注册审计 Hook
	client.Use(pkg.AuditHook())

	cleanup := func() {
		err := client.Close()
		if err != nil {
			log.Errorf("failed to close %s db: %s", conf.Data.Database.Driver, err.Error())
		}
		log.Infof("database client closed")
	}
	return client, cleanup, nil
}
