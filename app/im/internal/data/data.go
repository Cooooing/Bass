package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	"im/internal/config"
	"im/internal/data/client"
	"im/internal/data/repo"

	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(
	client.NewDataBaseClient,
	ProvideRedis,
	ProvideConsul,
	ProvideNats,
	commonClient.NewObservability,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewRedisLock,
	commonClient.NewNatsClient,
	repo.NewChatGroupRepo,
	repo.NewChatGroupMemberRepo,
	repo.NewChatSessionRepo,
	repo.NewChatMessageRepo,
)

func ProvideRedis(
	c *config.Bootstrap,
) *common.Redis {
	return c.Redis
}

func ProvideConsul(
	c *config.Bootstrap,
) *common.Consul {
	return c.Consul
}

func ProvideNats(
	c *config.Bootstrap,
) *common.Nats {
	return c.Nats
}
