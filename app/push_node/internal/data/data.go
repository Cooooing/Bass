package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	bizrepo "push_node/internal/biz/repo"
	"push_node/internal/config"
	"push_node/internal/data/repo"

	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(
	ProvideConsul,
	ProvideNats,
	commonClient.NewObservability,
	commonClient.NewConsulClient,
	commonClient.NewNatsClient,
	repo.NewConnectionRegistryRepo,
	wire.Bind(new(bizrepo.ConnectionRegistry), new(*repo.ConnectionRegistryRepo)),
	wire.Bind(new(commonClient.Subscriber), new(*commonClient.NatsClient)),
)

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
