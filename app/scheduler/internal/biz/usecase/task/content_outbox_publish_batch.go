package task

import (
	"common/pkg/client/rpc"
	contentv1 "common/proto/gen/content/v1"
	"context"
	"encoding/json"
	"strings"
	"time"

	schedulerenum "scheduler/internal/enum"

	"google.golang.org/protobuf/types/known/durationpb"
)

type ContentOutboxPublishBatch struct {
	contentClient      *rpc.ContentClient
	title, description string
	defaultLimit       int32
	defaultTimeout     time.Duration
	defaultRetry       int32
}

func NewContentOutboxPublishBatch(contentClient *rpc.ContentClient) *ContentOutboxPublishBatch {
	return &ContentOutboxPublishBatch{contentClient: contentClient, title: "内容 Outbox 批量投递", description: "调用 content.OutboxService.PublishBatch 投递待发送 outbox 事件。", defaultLimit: 1000, defaultTimeout: time.Minute, defaultRetry: 3}
}

func (*ContentOutboxPublishBatch) HandlerName() schedulerenum.TaskHandlerName {
	return schedulerenum.TaskHandlerNameContentOutboxPublishBatch
}
func (t *ContentOutboxPublishBatch) Title() string       { return t.title }
func (t *ContentOutboxPublishBatch) Description() string { return t.description }
func (t *ContentOutboxPublishBatch) DefaultScheduledTasks() []*DefaultScheduledTask {
	return []*DefaultScheduledTask{{TaskKey: schedulerenum.TaskKeyContentOutboxPublishBatchDefault, Title: t.title, Description: t.description, Enabled: true, CronSpec: "0 0/1 * * * ?", Payload: `{"limit":1000,"max_retry":1}`, Timeout: 30 * time.Second, MaxAttempts: 1, MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteLatest}}
}
func (t *ContentOutboxPublishBatch) DefaultDelayedTasks() []*DefaultDelayedTask {
	return []*DefaultDelayedTask{{TaskKey: schedulerenum.TaskKeyContentOutboxPublishBatchDelayedDefault, Title: t.title, Description: t.description, Enabled: true, Timeout: time.Minute, MaxAttempts: 3, MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteAll}}
}

type contentOutboxPublishBatchPayload struct {
	Limit          int32         `json:"limit"`
	PublishTimeout time.Duration `json:"publish_timeout"`
	MaxRetry       int32         `json:"max_retry"`
}

func (t *ContentOutboxPublishBatch) Execute(ctx context.Context, payload string) error {
	data := contentOutboxPublishBatchPayload{Limit: t.defaultLimit, PublishTimeout: t.defaultTimeout, MaxRetry: t.defaultRetry}
	if payload = strings.TrimSpace(payload); payload != "" && payload != "{}" {
		var override contentOutboxPublishBatchPayload
		if err := json.Unmarshal([]byte(payload), &override); err != nil {
			return err
		}
		if override.Limit > 0 {
			data.Limit = override.Limit
		}
		if override.PublishTimeout > 0 {
			data.PublishTimeout = override.PublishTimeout
		}
		if override.MaxRetry > 0 {
			data.MaxRetry = override.MaxRetry
		}
	}
	_, err := t.contentClient.Outbox.PublishBatch(ctx, &contentv1.PublishOutboxEvents_Req{Limit: data.Limit, PublishTimeout: durationpb.New(data.PublishTimeout), MaxRetry: data.MaxRetry})
	return err
}
