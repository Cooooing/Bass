package model

import (
	templatedata "notify/internal/biz/model/template_data"
	notifyenum "notify/internal/enum"
)

type NotificationContext struct {
	EventID          string
	NotificationType notifyenum.NotificationType
	ActorUserID      *int64
	BusinessKey      *string
	Language         notifyenum.Language
	TemplateData     templatedata.NotificationTemplateData
	Recipients       []*NotificationRecipient
}
