package model

import (
	"bytes"
	"context"
	"text/template"
	"time"
)

type NotificationEmailTemplate struct {
	ID              int64
	RuleID          int64
	SubjectTemplate string
	BodyTemplate    string
	ContentType     string
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}

type NotificationEmailRendered struct {
	Subject     string
	Body        string
	ContentType string
}

func (t *NotificationEmailTemplate) Render(ctx context.Context, data any) (*NotificationEmailRendered, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	subjectTemplate, err := template.New("").Option("missingkey=error").Parse(t.SubjectTemplate)
	if err != nil {
		return nil, err
	}
	var subject bytes.Buffer
	if err := subjectTemplate.Execute(&subject, data); err != nil {
		return nil, err
	}
	bodyTemplate, err := template.New("").Option("missingkey=error").Parse(t.BodyTemplate)
	if err != nil {
		return nil, err
	}
	var body bytes.Buffer
	if err := bodyTemplate.Execute(&body, data); err != nil {
		return nil, err
	}
	return &NotificationEmailRendered{
		Subject:     subject.String(),
		Body:        body.String(),
		ContentType: t.ContentType,
	}, nil
}
