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
	return fmt.Sprintf("%s_%s", n.NotificationType, n.Channel)
}

func (n *NotificationTemplate) ConvertToRpc() *v1.NotificationTemplate {
	return &v1.NotificationTemplate{
		CreatedAt:        timestamppb.New(*n.CreatedAt),
		UpdatedAt:        timestamppb.New(*n.UpdatedAt),
		Id:               n.ID,
		NotificationType: v1.NotificationType(v1.NotificationType_value[n.NotificationType]),
		Channel:          v1.NotificationChannel(v1.NotificationChannel_value[n.Channel]),
		Title:            n.Title,
		Content:          n.Content,
		Enable:           n.Enable,
	}
}
