package client

import (
	"common/pkg"
	"content/internal/conf"
	"content/internal/data/ent/gen"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/lib/pq"
)

func NewDataBaseClient(log *log.Helper, conf *conf.Bootstrap) (*gen.Client, func(), error) {
	logFunc := func(args ...any) {
		text := fmt.Sprint(args[0])
		log.Debugf("%s", text)
	}

	client, err := gen.Open(conf.Data.Database.Driver, conf.Data.Database.Source, gen.Log(logFunc))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open db: %w", err)
	}
	log.Infof("database: ent created database client [%s]", conf.Data.Database.Driver)
	client = client.Debug()
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
