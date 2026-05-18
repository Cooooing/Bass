package service

import (
	"common/pkg/util/server"

	"github.com/google/wire"
)

// ServiceProviderSet is service providers.
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	ProvideServices,
)

func ProvideServices(
	systemService *SystemService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
	}
}
