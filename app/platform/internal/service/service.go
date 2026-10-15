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
	NewHolidayService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	ipResolutionService *IPResolutionService,
	assetService *AssetService,
	holidayService *HolidayService,
) []server.Service {
	return []server.Service{
		commonSystemService,
		ipResolutionService,
		assetService,
		holidayService,
	}
}
