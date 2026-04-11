package model

import (
	v1 "common/api/gen/notify/v1"
	"fmt"
	"notify/internal/data/ent/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type NotificationTemplate struct {
	*gen.NotificationTemplate
}

func GetKeyNotificationTemplate(notificationType *v1.NotificationType, channel *v1.NotificationChannel) string {
	if notificationType == nil || channel == nil {
		return ""
	}
	return fmt.Sprintf("%s_%s", notificationType.String(), channel.String())
}

func (n *NotificationTemplate) GetKey() string {
	return GetKeyNotificationTemplate(new(v1.NotificationType(n.NotificationType)), new(v1.NotificationChannel(n.Channel)))
}

func (n *NotificationTemplate) ConvertToRpc() *v1.NotificationTemplate {
	return &v1.NotificationTemplate{
		CreatedAt:        timestamppb.New(*n.CreatedAt),
		UpdatedAt:        timestamppb.New(*n.UpdatedAt),
		Id:               n.ID,
		NotificationType: v1.NotificationType(n.NotificationType),
		Channel:          v1.NotificationChannel(n.Channel),
		Title:            n.Title,
		Content:          n.Content,
		Processors:       n.Processors,
		Enable:           n.Enable,
	}
}
