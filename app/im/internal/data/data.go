package data

import (
	"common/api/gen/common"
	commonClient "common/pkg/client"
	"im/internal/conf"
	database "im/internal/data/base"
	"im/internal/data/client"
	"im/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet is data providers.
var DataProviderSet = wire.NewSet(
	database.NewBaseData,

	client.NewDataBaseClient,
	ProvideRedis,
	ProvideConsul,
	ProvideNats,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,

	repo.NewChatGroupRepo,
	repo.NewChatSessionRepo,
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
