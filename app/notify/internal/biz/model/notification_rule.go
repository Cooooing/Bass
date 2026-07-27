package model

import (
	commonenum "common/pkg/enum"
	notifyenum "notify/internal/enum"
	"time"
)

type NotificationRule struct {
	ID        int64
	EventType commonenum.EventType
	Channel   notifyenum.NotificationChannel
	Language  notifyenum.Language
	Enabled   bool
	CreatedAt *time.Time
	UpdatedAt *time.Time

	StationTemplate     *NotificationStationTemplate
	EmailTemplate       *NotificationEmailTemplate
	TencentSMSTemplate  *NotificationTencentSMSTemplate
	LarkWebhookTemplate *NotificationLarkWebhookTemplate
}
