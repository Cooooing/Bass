package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewPointsService,
	NewPointsTccService,
	NewPointsTransferTccService,
)

func ProvideServices(commonSystemService *CommonSystemService, pointsService *PointsService, pointsTccService *PointsTccService, pointsTransferTccService *PointsTransferTccService) []server.Service {
	return []server.Service{
		commonSystemService,
		pointsService,
		pointsTccService,
		pointsTransferTccService,
	}
}
