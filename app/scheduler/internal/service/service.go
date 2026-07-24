package service

import (
	"common/pkg/server"

	"github.com/google/wire"
)

var ServiceProviderSet = wire.NewSet(
	ProvideServices,
	NewCommonSystemService,
	NewSchedulerTaskService,
	NewSchedulerDelayedTaskService,
)

func ProvideServices(
	commonSystemService *CommonSystemService,
	schedulerTaskService *SchedulerTaskService,
	schedulerDelayedTaskService *SchedulerDelayedTaskService,
) []server.Service {
	return []server.Service{commonSystemService, schedulerTaskService, schedulerDelayedTaskService}
}
