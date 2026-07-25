package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewStationMessageService,
	NewRateLimitService,
)

func ProvideServices(commonSystemService *CommonSystemService, stationMessageService *StationMessageService, rateLimitService *RateLimitService) []server.Service {
	return []server.Service{
		commonSystemService,
		stationMessageService,
		rateLimitService,
	}
}
