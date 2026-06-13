package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewSystemService,

	ProvideGrpcServices,
)

func ProvideGrpcServices(
	systemService *SystemService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
	}
}
