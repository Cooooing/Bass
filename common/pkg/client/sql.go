package client

import (
	"common/pkg/constant"
	"database/sql"
	"fmt"
	"log/slog"

	"common/proto/gen/common"
)

// SQLClient 管理数据库连接池。
type SQLClient struct {
	DB     *sql.DB
	Driver string
}

func NewSQLClient(logger *slog.Logger, conf *common.Database) (*SQLClient, func(), error) {
	db, err := sql.Open(conf.GetDriver(), conf.GetSource())
	if err != nil {
		return nil, nil, fmt.Errorf("open database: %w", err)
	}
	logger.Info("database client created", constant.LogFieldKind, constant.LogKindDatabase, constant.LogFieldDriver, conf.GetDriver())
	cleanup := func() {
		if err := db.Close(); err != nil {
			logger.Error("close database client failed", constant.LogFieldKind, constant.LogKindDatabase, constant.LogFieldErr, err)
		}
		logger.Info("database client closed", constant.LogFieldKind, constant.LogKindDatabase)
	}
	return &SQLClient{DB: db, Driver: conf.GetDriver()}, cleanup, nil
}
