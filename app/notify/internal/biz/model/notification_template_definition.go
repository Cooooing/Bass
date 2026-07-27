package model

import (
	"bytes"
	commonenum "common/pkg/enum"
	"context"
	notifyenum "notify/internal/enum"
	"text/template"
)

type NotificationTemplateDefinition struct {
	EventType commonenum.EventType
	Channel   notifyenum.NotificationChannel
	Language  notifyenum.Language
	Enabled   bool

	StationTemplate     *NotificationStationTemplateDefinition
	EmailTemplate       *NotificationEmailTemplateDefinition
	TencentSMSTemplate  *NotificationTencentSMSTemplateDefinition
	LarkWebhookTemplate *NotificationLarkWebhookTemplateDefinition
}

type NotificationStationTemplateDefinition struct {
	TitleTemplate   string
	ContentTemplate string
}

type NotificationEmailTemplateDefinition struct {
	SubjectTemplate string
	BodyTemplate    string
	ContentType     string
}

type NotificationTencentSMSTemplateDefinition struct {
	SMSSDKAppID        string
	SignName           string
	ProviderTemplateID string
	ParamTemplates     []string
}

type NotificationLarkWebhookTemplateDefinition struct {
	WebhookID       string
	Token           string
	Secret          string
	MsgType         string
	ContentTemplate string
}

type NotificationTemplatePreview struct {
	Template string
}

func (t *NotificationTemplatePreview) Render(ctx context.Context, data any) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", err
	}
	previewTemplate, err := template.New("").Option("missingkey=error").Parse(t.Template)
	if err != nil {
		return "", err
	}
	var rendered bytes.Buffer
	if err := previewTemplate.Execute(&rendered, data); err != nil {
		return "", err
	}
	return rendered.String(), nil
}
