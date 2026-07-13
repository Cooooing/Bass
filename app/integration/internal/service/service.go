package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewCommonSystemService,

	ProvideGrpcServices,
	ProvideHttpServices,
)

func ProvideGrpcServices(
	commonSystemService *CommonSystemService,
) []server.GrpcService {
	return []server.GrpcService{
		commonSystemService,
	}
}

func ProvideHttpServices(
	commonSystemService *CommonSystemService,
) []server.HttpService {
	return []server.HttpService{
		commonSystemService,
	}
}
