package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	NewIpResolutionService,
	NewOssService,

	ProvideGrpcServices,
	ProvideHttpServices,
)

func ProvideGrpcServices(
	systemService *SystemService,
	ipResolutionService *IpResolutionService,
	ossService *OssService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
		ipResolutionService,
		ossService,
	}
}

func ProvideHttpServices(
	ossService *OssService,
) []server.HttpService {
	return []server.HttpService{
		ossService,
	}
}
