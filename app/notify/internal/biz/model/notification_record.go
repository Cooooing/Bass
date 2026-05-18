package model

import (
	v1 "common/api/gen/notify/v1"
	"notify/internal/data/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type NotificationRecord struct {
	*gen.NotificationRecord

	WithMeta bool
}

func (n *NotificationRecord) ConvertToRpc() *v1.NotificationRecord {
	notificationRecord := &v1.NotificationRecord{
		CreatedAt:      timestamppb.New(*n.CreatedAt),
		UpdatedAt:      timestamppb.New(*n.UpdatedAt),
		Id:             n.ID,
		NotificationId: n.NotificationID,
		ReceiverId:     n.ReceiverID,
	}
	if n.ReadTime != nil {
		notificationRecord.ReadTime = timestamppb.New(*n.ReadTime)
	}
	if n.WithMeta {
		notificationRecord.NotificationMeta = (&NotificationMeta{NotificationMeta: n.Edges.NotificationMeta}).ConvertToRpc()
	}
	return notificationRecord
}
