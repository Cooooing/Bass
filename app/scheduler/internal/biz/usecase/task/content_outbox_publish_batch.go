package task

import (
	"common/pkg/client/rpc"
	contentv1 "common/proto/gen/content/v1"
	"context"
	"encoding/json"
	schedulerenum "scheduler/internal/enum"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

type ContentOutboxPublishBatch struct {
	contentClient  *rpc.ContentClient
	title          string
	description    string
	defaultLimit   int32
	defaultTimeout time.Duration
	defaultRetry   int32
}

func NewContentOutboxPublishBatch(
	contentClient *rpc.ContentClient,
) *ContentOutboxPublishBatch {
	return &ContentOutboxPublishBatch{
		contentClient:  contentClient,
		title:          "内容 Outbox 批量投递",
		description:    "调用 content.OutboxService.PublishBatch 投递待发送 outbox 事件。",
		defaultLimit:   1000,
		defaultTimeout: 1 * time.Minute,
		defaultRetry:   3,
	}
}

func (t *ContentOutboxPublishBatch) HandlerName() schedulerenum.TaskHandlerName {
	return schedulerenum.TaskHandlerNameContentOutboxPublishBatch
}

func (t *ContentOutboxPublishBatch) Title() string {
	return t.title
}

func (t *ContentOutboxPublishBatch) Description() string {
	return t.description
}

func (t *ContentOutboxPublishBatch) DefaultScheduledTasks() []*DefaultScheduledTask {
	return []*DefaultScheduledTask{
		{
			TaskKey:       schedulerenum.TaskKeyContentOutboxPublishBatchDefault,
			Title:         t.Title(),
			Description:   t.Description(),
			Enabled:       true,
			CronSpec:      "0 0/1 * * * ?",
			Payload:       `{"limit":1000,"max_retry":1}`,
			Timeout:       30 * time.Second,
			MaxAttempts:   1,
			MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteLatest,
			AllowOverlap:  false,
		},
	}
}

func (t *ContentOutboxPublishBatch) DefaultDelayedTasks() []*DefaultDelayedTask {
	return []*DefaultDelayedTask{
		{
			TaskKey:       schedulerenum.TaskKeyContentOutboxPublishBatchDelayedDefault,
			Title:         t.Title(),
			Description:   t.Description(),
			Enabled:       true,
			Timeout:       1 * time.Minute,
			MaxAttempts:   3,
			MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteAll,
		},
	}
}

type contentOutboxPublishBatchPayload struct {
	Limit          int32         `json:"limit"`
	PublishTimeout time.Duration `json:"publish_timeout"`
	MaxRetry       int32         `json:"max_retry"`
}

func (t *ContentOutboxPublishBatch) Execute(ctx context.Context, payload string) error {
	limit := t.defaultLimit
	publishTimeout := t.defaultTimeout
	maxRetry := t.defaultRetry
	payload = strings.TrimSpace(payload)
	if payload != "" && payload != "{}" {
		data := &contentOutboxPublishBatchPayload{}
		if err := json.Unmarshal([]byte(payload), data); err != nil {
			return err
		}
		if data.Limit > 0 {
			limit = data.Limit
		}
		if data.PublishTimeout > 0 {
			publishTimeout = data.PublishTimeout
		}
		if data.MaxRetry > 0 {
			maxRetry = data.MaxRetry
		}
	}
	_, err := t.contentClient.Outbox.PublishBatch(ctx, &contentv1.PublishOutboxEvents_Req{
		Limit:          limit,
		PublishTimeout: durationpb.New(publishTimeout),
		MaxRetry:       maxRetry,
	})
	return err
}
