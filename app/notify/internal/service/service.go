package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	NewStationMessageService,
	NewRateLimitService,
	ProvideServices,
)

func ProvideServices(
	systemService *SystemService,
	stationMessageService *StationMessageService,
	rateLimitService *RateLimitService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
		stationMessageService,
		rateLimitService,
	}
}
