package model

import (
	"bytes"
	"context"
	"text/template"
	"time"
)

type NotificationStationTemplate struct {
	ID              int64
	RuleID          int64
	TitleTemplate   string
	ContentTemplate string
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}

type NotificationStationRendered struct {
	Title   string
	Content string
}

func (t *NotificationStationTemplate) Render(ctx context.Context, data any) (*NotificationStationRendered, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	titleTemplate, err := template.New("").Option("missingkey=error").Parse(t.TitleTemplate)
	if err != nil {
		return nil, err
	}
	var title bytes.Buffer
	if err := titleTemplate.Execute(&title, data); err != nil {
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
	return &NotificationStationRendered{
		Title:   title.String(),
		Content: content.String(),
	}, nil
}
