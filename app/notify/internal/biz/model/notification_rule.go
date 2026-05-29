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

type NotificationStationTemplate struct {
	ID              int64
	RuleID          int64
	TitleTemplate   string
	ContentTemplate string
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}

type NotificationEmailTemplate struct {
	ID              int64
	RuleID          int64
	SubjectTemplate string
	BodyTemplate    string
	ContentType     string
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}

type NotificationTencentSMSTemplate struct {
	ID                 int64
	RuleID             int64
	SMSSDKAppID        string
	SignName           string
	ProviderTemplateID string
	ParamTemplates     []string
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}

type NotificationLarkWebhookTemplate struct {
	ID           int64
	RuleID       int64
	WebhookID    string
	Token        string
	BodyTemplate string
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
