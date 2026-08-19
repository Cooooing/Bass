package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	"economy/internal/config"
	"economy/internal/data/client"
	"economy/internal/data/repo"

	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	commonClient.NewObservability,
	commonClient.NewConsulClient,
	client.ProvideTx,
	client.NewTransactionNoGenerator,
	repo.NewAccountRepo,
	repo.NewRecordRepo,
	ProvideConsul,
)

func ProvideConsul(c *config.Bootstrap) *common.Consul {
	return c.Consul
}
