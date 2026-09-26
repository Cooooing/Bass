package service

import (
	commonserver "common/pkg/server"

	"github.com/google/wire"
)

// ServiceProviderSet 提供管理端当前可暴露的基础服务。
var ServiceProviderSet = wire.NewSet(ProvideServices, NewCommonSystemService)

func ProvideServices(commonSystemService *CommonSystemService) []commonserver.Service {
	return []commonserver.Service{commonSystemService}
}
