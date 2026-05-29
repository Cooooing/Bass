package rpc

import (
	notifyv1 "common/api/gen/notify/v1"

	"google.golang.org/grpc"
)

type NotifyClient struct {
	StationMessage notifyv1.NotifyStationMessageServiceClient
}

func NewNotifyClient(conn *grpc.ClientConn) *NotifyClient {
	return &NotifyClient{
		StationMessage: notifyv1.NewNotifyStationMessageServiceClient(conn),
	}
}
