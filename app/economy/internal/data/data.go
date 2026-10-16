package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"

	"economy/internal/config"
	"economy/internal/data/client"
	"economy/internal/data/gen"
	"economy/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet 提供微服务模式的数据层依赖。
var DataProviderSet = wire.NewSet(
	ModuleProviderSet,
	commonClient.NewObservability,
	ProvideConsul,
	commonClient.NewConsulClient,
)

// ModuleProviderSet 提供不依赖服务发现的模块数据层依赖。
var ModuleProviderSet = wire.NewSet(
	commonClient.NewSQLClient,
	ProvideDatabase,
	client.NewEntClient,
	ProvideGenClient,
	ProvideTCCBranchExecutor,
	client.ProvideTx,
	client.NewTransactionNoGenerator,
	ProvideDTMConfig,
	commonClient.NewDTMClient,
	repo.NewAccountRepo,
	repo.NewRecordRepo,
)

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}

// ProvideDTMConfig 从模块启动配置中提取 DTM 配置。
func ProvideDTMConfig(c *config.Bootstrap) *common.DTM {
	return c.Dtm
}

func ProvideDatabase(c *config.Bootstrap) *common.Database {
	return c.Database
}

func ProvideGenClient(c *client.EntClient) *gen.Client {
	return c.Client
}

func ProvideTCCBranchExecutor(sqlClient *commonClient.SQLClient, entClient *client.EntClient, dtm *commonClient.DTMClient) (commonClient.TCCBranchExecutor, error) {
	return commonClient.NewTCCBranchExecutor(dtm, sqlClient.DB, sqlClient.Driver, entClient.ClientFromTx)
}
