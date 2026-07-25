package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewIpResolutionService,
	NewOssService,
)

func ProvideServices(commonSystemService *CommonSystemService, ipResolutionService *IpResolutionService, ossService *OssService) []server.Service {
	return []server.Service{
		commonSystemService,
		ipResolutionService,
		ossService,
	}
}
