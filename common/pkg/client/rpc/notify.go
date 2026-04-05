package rpc

import (
	notifyv1 "common/gen/notify/v1"

	"google.golang.org/grpc"
)

type NotifyClient struct {
	Template notifyv1.NotifyNotificationTemplateServiceClient
	Meta     notifyv1.NotifyNotificationMetaServiceClient
	Record   notifyv1.NotifyNotificationRecordServiceClient
}

func NewNotifyClient(conn *grpc.ClientConn) *NotifyClient {
	return &NotifyClient{
		Template: notifyv1.NewNotifyNotificationTemplateServiceClient(conn),
		Meta:     notifyv1.NewNotifyNotificationMetaServiceClient(conn),
		Record:   notifyv1.NewNotifyNotificationRecordServiceClient(conn),
	}
}
