package model

import (
	"common/api/gen/common/enums"
	v1 "common/api/gen/notify/v1"
	"fmt"
	"notify/internal/data/ent/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type NotificationTemplate struct {
	*gen.NotificationTemplate
}

func (n *NotificationTemplate) GetKey() string {
	return fmt.Sprintf("%s_%s", n.EventType, n.Channel)
}

func (n *NotificationTemplate) ConvertToRpc() *v1.NotificationTemplate {
	return &v1.NotificationTemplate{
		CreatedAt: timestamppb.New(*n.CreatedAt),
		UpdatedAt: timestamppb.New(*n.UpdatedAt),
		Id:        n.ID,
		EventType: enums.EventType(enums.EventType_value[string(n.EventType)]),
		Channel:   v1.NotificationChannel(v1.NotificationChannel_value[string(n.Channel)]),
		Title:     n.Title,
		Content:   n.Content,
		Enable:    n.Enable,
	}
}
