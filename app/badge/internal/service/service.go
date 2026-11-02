package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(ProvideServices, NewBadgeService)

func ProvideServices(badgeService *BadgeService) []server.Service {
	return []server.Service{badgeService}
}
