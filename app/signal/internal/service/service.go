package service

import (
	"common/pkg/util/server"

	"github.com/google/wire"
)

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	ProvideServices,

	NewNodeService,
)

func ProvideServices(
	systemService *SystemService,
	nodeService *NodeService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
		nodeService,
	}
}
