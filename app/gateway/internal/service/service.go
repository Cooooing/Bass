package service

import (
	"common/pkg/util/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	ProvideServices,
)

func ProvideServices(
	systemService *SystemService,
) []server.HttpService {
	return []server.HttpService{
		systemService,
	}
}
