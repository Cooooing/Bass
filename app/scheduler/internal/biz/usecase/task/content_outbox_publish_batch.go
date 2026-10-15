package task

import (
	"common/pkg/client/rpc"
	commonscheduler "common/pkg/scheduler"
	contentv1 "common/proto/gen/content/v1"
	schedulerv1 "common/proto/gen/scheduler/v1"
	"context"
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
	return []*DefaultScheduledTask{{TaskKey: schedulerenum.TaskKeyContentOutboxPublishBatchDefault, Title: t.title, Description: t.description, Enabled: true, CronSpec: "0 0/1 * * * ?", Payload: commonscheduler.MustMarshalPayload(&schedulerv1.OutboxPublishBatchTaskPayload{Limit: t.defaultLimit, MaxRetry: 1}), Timeout: 30 * time.Second, MaxAttempts: 1, MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteLatest}}
}
func (t *ContentOutboxPublishBatch) DefaultDelayedTasks() []*DefaultDelayedTask {
	return []*DefaultDelayedTask{{TaskKey: schedulerenum.TaskKeyContentOutboxPublishBatchDelayedDefault, Title: t.title, Description: t.description, Enabled: true, Timeout: time.Minute, MaxAttempts: 3, MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteAll}}
}

func (t *ContentOutboxPublishBatch) Execute(ctx context.Context, payload string) error {
	data := &schedulerv1.OutboxPublishBatchTaskPayload{}
	if err := commonscheduler.UnmarshalPayload(payload, data); err != nil {
		return err
	}
	limit, timeout, maxRetry := t.defaultLimit, t.defaultTimeout, t.defaultRetry
	if data.GetLimit() > 0 {
		limit = data.GetLimit()
	}
	if data.GetPublishTimeout() > 0 {
		timeout = time.Duration(data.GetPublishTimeout())
	}
	if data.GetMaxRetry() > 0 {
		maxRetry = data.GetMaxRetry()
	}
	_, err := t.contentClient.Outbox.PublishBatch(ctx, &contentv1.PublishOutboxEvents_Req{Limit: limit, PublishTimeout: durationpb.New(timeout), MaxRetry: maxRetry})
	return err
}
