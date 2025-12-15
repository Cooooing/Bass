package model

import (
	v1 "common/api/notify/v1"
	"common/pkg/constant"
	"common/pkg/cutil/base"
	"notify/internal/data/ent/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type NotificationTemplate struct {
	*gen.NotificationTemplate
}

func (n *NotificationTemplate) GetKey() string {
	return constant.GetKeyNotificationTemplate(base.Ptr(v1.NotificationType(n.NotificationType)), base.Ptr(v1.NotificationChannel(n.Channel)))
}

func (n *NotificationTemplate) ConvertToRpc() *v1.NotificationTemplate {
	return &v1.NotificationTemplate{
		CreatedAt:        timestamppb.New(*n.CreatedAt),
		UpdatedAt:        timestamppb.New(*n.UpdatedAt),
		Id:               n.ID,
		NotificationType: n.NotificationType,
		Channel:          n.Channel,
		Content:          n.Content,
		Processors:       n.Processors,
		Enable:           n.Enable,
	}
}
