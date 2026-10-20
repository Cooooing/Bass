package service

import (
	commonserver "common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(ProvideServices, NewRealtimeService)

func ProvideServices(realtime *RealtimeService) []commonserver.Service {
	return []commonserver.Service{realtime}
}
