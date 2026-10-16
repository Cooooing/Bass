package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewPointsService,
	NewTCCPointsService,
	NewTCCPointsTransferService,
)

func ProvideServices(commonSystemService *CommonSystemService, pointsService *PointsService, tccPointsService *TCCPointsService, tccPointsTransferService *TCCPointsTransferService) []server.Service {
	return []server.Service{
		commonSystemService,
		pointsService,
		tccPointsService,
		tccPointsTransferService,
	}
}
