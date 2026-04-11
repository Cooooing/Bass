package model

import (
	v1 "common/api/gen/notify/v1"
	"encoding/json"
	"notify/internal/data/ent/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type NotificationMeta struct {
	*gen.NotificationMeta
}

func (n *NotificationMeta) ConvertToRpc() *v1.NotificationMeta {
	m := &v1.NotificationMeta{
		CreatedAt:        timestamppb.New(*n.CreatedAt),
		UpdatedAt:        timestamppb.New(*n.UpdatedAt),
		Id:               n.ID,
		Uuid:             n.UUID,
		NotificationType: v1.NotificationType(n.NotificationType),
		SenderId:         n.SenderID,
		Status:           v1.NotificationStatus(n.Status),
		ContentRender:    n.Content,
	}
	s, _ := json.Marshal(n.Meta)
	m.Meta = string(s)
	return m
}
