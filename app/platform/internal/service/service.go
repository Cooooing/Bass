package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewIpResolutionService,
	NewAssetService,
)

func ProvideServices(commonSystemService *CommonSystemService, ipResolutionService *IpResolutionService, assetService *AssetService) []server.Service {
	return []server.Service{
		commonSystemService,
		ipResolutionService,
		assetService,
	}
}
