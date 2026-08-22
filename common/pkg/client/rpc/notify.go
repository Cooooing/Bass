package rpc

import (
	"common/pkg/client/localrpc"
	notifyv1 "common/proto/gen/notify/v1"

	"google.golang.org/grpc"
)

type NotifyClient struct {
	StationMessage notifyv1.NotifyStationMessageServiceClient
	RateLimit      notifyv1.NotifyRateLimitServiceClient
	Template       notifyv1.NotifyTemplateServiceClient
}

func NewNotifyClient(
	conn grpc.ClientConnInterface,
) *NotifyClient {
	return &NotifyClient{
		StationMessage: notifyv1.NewNotifyStationMessageServiceClient(conn),
		RateLimit:      notifyv1.NewNotifyRateLimitServiceClient(conn),
		Template:       notifyv1.NewNotifyTemplateServiceClient(conn),
	}
}

func NewLocalNotifyClient[T any](services []T) *NotifyClient {
	conn := localrpc.NewConn()
	MountNotifyServices(conn, services)
	return NewNotifyClient(conn)
}

func MountNotifyServices[T any](conn *localrpc.Conn, services []T) {
	for _, service := range services {
		conn.RegisterMatching(&notifyv1.NotifyStationMessageService_ServiceDesc, service)
		conn.RegisterMatching(&notifyv1.NotifyRateLimitService_ServiceDesc, service)
		conn.RegisterMatching(&notifyv1.NotifyTemplateService_ServiceDesc, service)
	}
}
