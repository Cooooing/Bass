package data

import (
	commonclient "common/pkg/client"
	"common/proto/gen/common"
	"realtime_hub/internal/config"
	datarepo "realtime_hub/internal/data/repo"

	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(ModuleProviderSet, commonclient.NewObservability, commonclient.NewRedisClient, commonclient.NewNatsClient, commonclient.NewConsulClient, ProvideNats, ProvideConsul)

var ModuleProviderSet = wire.NewSet(ProvideRedis, ProvideSubscriber, ProvidePublisher, datarepo.NewRegistry, datarepo.NewControlPublisher)

func ProvideRedis(config *config.Bootstrap) *common.Redis {
	return config.Redis
}

func ProvideNats(config *config.Bootstrap) *common.Nats {
	return config.Nats
}

func ProvideConsul(config *config.Bootstrap) *common.Consul {
	return config.Consul
}

func ProvideSubscriber(client *commonclient.NatsClient) commonclient.Subscriber {
	return client
}

func ProvidePublisher(client *commonclient.NatsClient) commonclient.Publisher {
	return client
}
