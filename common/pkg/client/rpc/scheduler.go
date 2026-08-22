package rpc

import (
	"common/pkg/client/localrpc"
	schedulerv1 "common/proto/gen/scheduler/v1"

	"google.golang.org/grpc"
)

type SchedulerClient struct {
	ScheduledTask schedulerv1.SchedulerScheduledTaskServiceClient
	DelayedTask   schedulerv1.SchedulerDelayedTaskServiceClient
}

func NewSchedulerClient(
	conn grpc.ClientConnInterface,
) *SchedulerClient {
	return &SchedulerClient{
		ScheduledTask: schedulerv1.NewSchedulerScheduledTaskServiceClient(conn),
		DelayedTask:   schedulerv1.NewSchedulerDelayedTaskServiceClient(conn),
	}
}

func MountSchedulerServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&schedulerv1.SchedulerScheduledTaskService_ServiceDesc, service)
		conn.RegisterMatching(&schedulerv1.SchedulerDelayedTaskService_ServiceDesc, service)
	}
}
