package data

import (
	commonclient "common/pkg/client"
	nats "common/pkg/client/nats"
	"common/proto/gen/common/v1"

	"realtime_hub/internal/config"
	datarepo "realtime_hub/internal/data/repo"

	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(ModuleProviderSet, commonclient.NewObservability, commonclient.NewRedisClient, nats.NewClient, commonclient.NewConsulClient, ProvideNats, ProvideConsul)

var ModuleProviderSet = wire.NewSet(ProvideRedis, ProvideBusiness, datarepo.NewRegistry, datarepo.NewControlPublisher)

func ProvideRedis(config *config.Bootstrap) *common.Redis {
	return config.GetRedis()
}

func ProvideBusiness(config *config.Bootstrap) *config.Business {
	return config.GetBusiness()
}

func ProvideNats(config *config.Bootstrap) *common.Nats {
	return config.GetNats()
}

func ProvideConsul(config *config.Bootstrap) *common.Consul {
	return config.GetConsul()
}
