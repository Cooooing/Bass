package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(ProvideServices, NewCommonSystemService, NewRealtimeService)

func ProvideServices(system *CommonSystemService, realtime *RealtimeService) []server.Service {
	return []server.Service{system, realtime}
}
