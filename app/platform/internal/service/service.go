package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewIPResolutionService,
	NewAssetService,
)

func ProvideServices(commonSystemService *CommonSystemService, ipResolutionService *IPResolutionService, assetService *AssetService) []server.Service {
	return []server.Service{
		commonSystemService,
		ipResolutionService,
		assetService,
	}
}
