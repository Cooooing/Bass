package rpc

import (
	schedulerv1 "common/proto/gen/scheduler/v1"

	"google.golang.org/grpc"
)

type SchedulerClient struct {
	Task        schedulerv1.SchedulerTaskServiceClient
	DelayedTask schedulerv1.SchedulerDelayedTaskServiceClient
}

func NewSchedulerClient(
	conn *grpc.ClientConn,
) *SchedulerClient {
	return &SchedulerClient{
		Task:        schedulerv1.NewSchedulerTaskServiceClient(conn),
		DelayedTask: schedulerv1.NewSchedulerDelayedTaskServiceClient(conn),
	}
}
