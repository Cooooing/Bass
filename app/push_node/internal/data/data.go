package data

import (
	"common/proto/gen/common"
	commonClient "common/pkg/client"
	"push_node/internal/biz/repo"
	"push_node/internal/conf"
	datarepo "push_node/internal/data/repo"

	"github.com/google/wire"
)

// DataProviderSet 是 data 层依赖集合。
var DataProviderSet = wire.NewSet(
	ProvideConsul,
	ProvideNats,
	commonClient.NewConsulClient,
	commonClient.NewNatsClient,

	datarepo.NewConnectionRegistryRepo,
	wire.Bind(new(repo.ConnectionRegistry), new(*datarepo.ConnectionRegistryRepo)),
	wire.Bind(new(commonClient.Subscriber), new(*commonClient.NatsClient)),
)

// ProvideConsul 从配置中提取 Consul 配置。
func ProvideConsul(c *conf.Bootstrap) *common.Consul {
	return c.Data.Consul
}

// ProvideNats 从配置中提取 NATS 配置。
func ProvideNats(c *conf.Bootstrap) *common.Nats {
	return c.Data.Nats
}
