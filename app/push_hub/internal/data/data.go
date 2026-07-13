package data

import (
	commonClient "common/pkg/client"
	"common/proto/gen/common"
	bizrepo "push_hub/internal/biz/repo"
	"push_hub/internal/conf"
	"push_hub/internal/data/repo"

	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(
	ProvideRedis,
	ProvideConsul,
	ProvideNats,
	commonClient.NewObservability,
	commonClient.NewConsulClient,
	commonClient.NewRedisClient,
	commonClient.NewNatsClient,
	repo.NewNodeRegistryRepo,
	wire.Bind(new(bizrepo.NodeRegistry), new(*repo.NodeRegistryRepo)),
	wire.Bind(new(commonClient.Publisher), new(*commonClient.NatsClient)),
)

func ProvideRedis(c *conf.Bootstrap) *common.Redis   { return c.Redis }
func ProvideConsul(c *conf.Bootstrap) *common.Consul { return c.Consul }
func ProvideNats(c *conf.Bootstrap) *common.Nats     { return c.Nats }
