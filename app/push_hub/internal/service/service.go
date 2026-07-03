package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	NewPushEventService,
	NewNodeService,

	ProvideGrpcServices,
	ProvideHttpServices,
)

func ProvideGrpcServices(
	systemService *SystemService,
	pushEventService *PushEventService,
	nodeService *NodeService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
		pushEventService,
		nodeService,
	}
}

func ProvideHttpServices(
	systemService *SystemService,
	pushEventService *PushEventService,
	nodeService *NodeService,
) []server.HttpService {
	return []server.HttpService{
		systemService,
		pushEventService,
		nodeService,
	}
}
