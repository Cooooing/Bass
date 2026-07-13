package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewCommonSystemService,
	NewStationMessageService,
	NewRateLimitService,
	ProvideServices,
	ProvideHttpServices,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	stationMessageService *StationMessageService,
	rateLimitService *RateLimitService,
) []server.GrpcService {
	return []server.GrpcService{
		commonSystemService,
		stationMessageService,
		rateLimitService,
	}
}

func ProvideHttpServices(
	commonSystemService *CommonSystemService,
	stationMessageService *StationMessageService,
	rateLimitService *RateLimitService,
) []server.HttpService {
	return []server.HttpService{
		commonSystemService,
		stationMessageService,
		rateLimitService,
	}
}
