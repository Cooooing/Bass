package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewCommonSystemService,
	NewIpResolutionService,
	NewOssService,

	ProvideGrpcServices,
	ProvideHttpServices,
)

func ProvideGrpcServices(
	commonSystemService *CommonSystemService,
	ipResolutionService *IpResolutionService,
	ossService *OssService,
) []server.GrpcService {
	return []server.GrpcService{
		commonSystemService,
		ipResolutionService,
		ossService,
	}
}

func ProvideHttpServices(
	commonSystemService *CommonSystemService,
	ossService *OssService,
) []server.HttpService {
	return []server.HttpService{
		commonSystemService,
		ossService,
	}
}
