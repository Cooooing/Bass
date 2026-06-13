package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	ProvideHttpServices,
)

// ProvideHttpServices 提供 HTTP 服务列表。
func ProvideHttpServices(
	systemService *SystemService,
) []server.HttpService {
	return []server.HttpService{
		systemService,
	}
}
