package client

import (
	commonclient "common/pkg/client"
	"common/pkg/client/driver"
	"common/proto/gen/common"
	"context"
	"fmt"
	"log/slog"

	"badge/internal/data/gen"
	"badge/internal/data/gen/migrate"

	// Register Ent runtime hooks generated for the badge schema.
	_ "badge/internal/data/gen/runtime"

	entsql "entgo.io/ent/dialect/sql"
	// Register the PostgreSQL driver selected by the standard database config.
	_ "github.com/lib/pq"
)

type EntClient struct{ Client *gen.Client }

func NewEntClient(logger *slog.Logger, conf *common.Database, sqlClient *commonclient.SQLClient) (*EntClient, error) {
	client := gen.NewClient(gen.Driver(driver.Wrap(entsql.OpenDB(sqlClient.Driver, sqlClient.DB), logger, conf)))
	if conf.GetMerge() {
		if err := client.Schema.Create(context.Background(), migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
			return nil, fmt.Errorf("create badge schema: %w", err)
		}
	}
	return &EntClient{Client: client}, nil
}
