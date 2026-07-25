package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
) []server.Service {
	return []server.Service{commonSystemService}
}
