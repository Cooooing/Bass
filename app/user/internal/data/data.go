package data

import (
	"common/api/gen/common"
	commonClient "common/pkg/client"
	"user/internal/conf"
	"user/internal/data/client"
	"user/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	ProvideRedis,
	ProvideConsul,
	ProvideNats,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,

	client.ProvideTx,

	repo.NewUserRepo,
	repo.NewUserRelationRepo,
	repo.NewUserPreferencesRepo,
	repo.NewUserPrivacyRepo,
	repo.NewUserLocationRepo,
	repo.NewUserTfaRepo,
	repo.NewUserCheckinRepo,
)

func ProvideRedis(c *conf.Bootstrap) *common.Redis {
	return c.Data.Redis
}

func ProvideConsul(c *conf.Bootstrap) *common.Consul {
	return c.Data.Consul
}

func ProvideNats(c *conf.Bootstrap) *common.Nats {
	return c.Data.Nats
}
