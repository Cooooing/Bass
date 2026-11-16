package model

import (
	"time"

	notifyenum "notify/internal/enum"
)

type NotificationRule struct {
	ID               int64
	NotificationType notifyenum.NotificationType
	Channel          notifyenum.NotificationChannel
	Language         notifyenum.Language
	Enabled          bool
	CreatedAt        *time.Time
	UpdatedAt        *time.Time

	StationTemplate     *NotificationStationTemplate
	EmailTemplate       *NotificationEmailTemplate
	TencentSMSTemplate  *NotificationTencentSMSTemplate
	LarkWebhookTemplate *NotificationLarkWebhookTemplate
}
