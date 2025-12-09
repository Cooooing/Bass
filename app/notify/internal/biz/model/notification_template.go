package model

import (
	v1 "common/api/notify/v1"
	"common/pkg/constant"
	"common/pkg/cutil/base"
	"notify/internal/data/ent/gen"
)

type NotificationTemplate struct {
	*gen.NotificationTemplate
}

func (n *NotificationTemplate) GetKey() string {
	return constant.GetKeyNotificationTemplate(base.Ptr(v1.NotificationType(n.NotificationType)), base.Ptr(v1.NotificationChannel(n.Channel)))
}
