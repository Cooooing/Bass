package model

import (
	"bytes"
	"context"
	"text/template"
	"time"
)

type NotificationLarkWebhookTemplate struct {
	ID              int64
	RuleID          int64
	WebhookID       string
	Token           string
	Secret          string
	MsgType         string
	ContentTemplate string
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}

type NotificationLarkWebhookRendered struct {
	WebhookID string
	Token     string
	Secret    string
	MsgType   string
	Content   string
}

func (t *NotificationLarkWebhookTemplate) Render(ctx context.Context, data any) (*NotificationLarkWebhookRendered, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	contentTemplate, err := template.New("").Option("missingkey=error").Parse(t.ContentTemplate)
	if err != nil {
		return nil, err
	}
	var content bytes.Buffer
	if err := contentTemplate.Execute(&content, data); err != nil {
		return nil, err
	}
	return &NotificationLarkWebhookRendered{
		WebhookID: t.WebhookID,
		Token:     t.Token,
		Secret:    t.Secret,
		MsgType:   t.MsgType,
		Content:   content.String(),
	}, nil
}
