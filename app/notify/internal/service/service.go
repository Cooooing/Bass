package service

import (
	"common/pkg/util/server"

	"github.com/google/wire"
)

// ServiceProviderSet 是 service 层依赖集合。
var ServiceProviderSet = wire.NewSet(
	NewSystemService,
	NewStationMessageService,
	NewOssService,
	ProvideServices,
)

func ProvideServices(
	systemService *SystemService,
	stationMessageService *StationMessageService,
	ossService *OssService,
) []server.GrpcService {
	return []server.GrpcService{
		systemService,
		stationMessageService,
		ossService,
	}
}
