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
	NewTemplateService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	stationMessageService *StationMessageService,
	rateLimitService *RateLimitService,
	templateService *TemplateService,
) []server.Service {
	return []server.Service{
		commonSystemService,
		stationMessageService,
		rateLimitService,
		templateService,
	}
}
