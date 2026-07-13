package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewCommonSystemService,
	NewPushEventService,
	NewNodeService,

	ProvideGrpcServices,
	ProvideHttpServices,
)

func ProvideGrpcServices(
	commonSystemService *CommonSystemService,
	pushEventService *PushEventService,
	nodeService *NodeService,
) []server.GrpcService {
	return []server.GrpcService{
		commonSystemService,
		pushEventService,
		nodeService,
	}
}

func ProvideHttpServices(
	commonSystemService *CommonSystemService,
	pushEventService *PushEventService,
	nodeService *NodeService,
) []server.HttpService {
	return []server.HttpService{
		commonSystemService,
		pushEventService,
		nodeService,
	}
}
