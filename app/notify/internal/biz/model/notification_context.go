package model

import (
	commonenum "common/pkg/enum"
	templatedata "notify/internal/biz/model/template_data"
	notifyenum "notify/internal/enum"
)

type NotificationContext struct {
	EventID      string
	EventType    commonenum.EventType
	Language     notifyenum.Language
	TemplateData templatedata.NotificationTemplateData
	Recipients   []*NotificationRecipient
}
