package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewEconomyService,
)

func ProvideServices(commonSystemService *CommonSystemService, economyService *EconomyService) []server.Service {
	return []server.Service{
		commonSystemService,
		economyService,
	}
}
