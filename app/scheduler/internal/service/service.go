package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewSchedulerTaskService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	schedulerTaskService *SchedulerTaskService,
) []server.Service {
	return []server.Service{
		commonSystemService,
		schedulerTaskService,
	}
}
