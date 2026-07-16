package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewPushEventService,
	NewNodeService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	pushEventService *PushEventService,
	nodeService *NodeService,
) []server.Service {
	return []server.Service{
		commonSystemService,
		pushEventService,
		nodeService,
	}
}
