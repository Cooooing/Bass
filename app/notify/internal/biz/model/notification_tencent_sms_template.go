package model

import (
	"bytes"
	"context"
	"text/template"
	"time"
)

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

type NotificationTencentSMSRendered struct {
	SMSSDKAppID        string
	SignName           string
	ProviderTemplateID string
	TemplateParams     []string
}

func (t *NotificationTencentSMSTemplate) Render(ctx context.Context, data any) (*NotificationTencentSMSRendered, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	params := make([]string, 0, len(t.ParamTemplates))
	for _, paramTemplateText := range t.ParamTemplates {
		paramTemplate, err := template.New("").Option("missingkey=error").Parse(paramTemplateText)
		if err != nil {
			return nil, err
		}
		var param bytes.Buffer
		if err := paramTemplate.Execute(&param, data); err != nil {
			return nil, err
		}
		params = append(params, param.String())
	}
	return &NotificationTencentSMSRendered{
		SMSSDKAppID:        t.SMSSDKAppID,
		SignName:           t.SignName,
		ProviderTemplateID: t.ProviderTemplateID,
		TemplateParams:     params,
	}, nil
}
