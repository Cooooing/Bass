package data

import (
	commonclient "common/pkg/client"
	"common/pkg/client/rpc"
	"common/proto/gen/common"
	"realtime_node/internal/config"

	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(commonclient.NewObservability, commonclient.NewConsulClient, ProvideConsul, rpc.ProvideRealtimeHubClient, NewHubClient)

func ProvideConsul(config *config.Bootstrap) *common.Consul {
	return config.Consul
}
