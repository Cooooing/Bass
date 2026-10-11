package task

import (
	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	"context"
	"encoding/json"
	"strings"
	"time"

	schedulerenum "scheduler/internal/enum"

	"google.golang.org/protobuf/types/known/durationpb"
)

type UserOutboxPublishBatch struct {
	userClient         *rpc.UserClient
	title, description string
	defaultLimit       int32
	defaultTimeout     time.Duration
	defaultRetry       int32
}

func NewUserOutboxPublishBatch(userClient *rpc.UserClient) *UserOutboxPublishBatch {
	return &UserOutboxPublishBatch{userClient: userClient, title: "用户 Outbox 批量投递", description: "调用 user.OutboxService.PublishBatch 投递待发送 outbox 事件。", defaultLimit: 1000, defaultTimeout: time.Minute, defaultRetry: 3}
}

func (*UserOutboxPublishBatch) HandlerName() schedulerenum.TaskHandlerName {
	return schedulerenum.TaskHandlerNameUserOutboxPublishBatch
}
func (t *UserOutboxPublishBatch) Title() string       { return t.title }
func (t *UserOutboxPublishBatch) Description() string { return t.description }
func (t *UserOutboxPublishBatch) DefaultScheduledTasks() []*DefaultScheduledTask {
	return []*DefaultScheduledTask{{TaskKey: schedulerenum.TaskKeyUserOutboxPublishBatchDefault, Title: t.title, Description: t.description, Enabled: true, CronSpec: "0 0/1 * * * ?", Payload: `{"limit":1000,"max_retry":1}`, Timeout: 30 * time.Second, MaxAttempts: 1, MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteLatest}}
}
func (t *UserOutboxPublishBatch) DefaultDelayedTasks() []*DefaultDelayedTask {
	return []*DefaultDelayedTask{{TaskKey: schedulerenum.TaskKeyUserOutboxPublishBatchDelayedDefault, Title: t.title, Description: t.description, Enabled: true, Timeout: time.Minute, MaxAttempts: 3, MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteAll}}
}

type userOutboxPublishBatchPayload struct {
	Limit          int32         `json:"limit"`
	PublishTimeout time.Duration `json:"publish_timeout"`
	MaxRetry       int32         `json:"max_retry"`
}

func (t *UserOutboxPublishBatch) Execute(ctx context.Context, payload string) error {
	data := userOutboxPublishBatchPayload{Limit: t.defaultLimit, PublishTimeout: t.defaultTimeout, MaxRetry: t.defaultRetry}
	if payload = strings.TrimSpace(payload); payload != "" && payload != "{}" {
		var override userOutboxPublishBatchPayload
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
	_, err := t.userClient.Outbox.PublishBatch(ctx, &userv1.PublishOutboxEvents_Req{Limit: data.Limit, PublishTimeout: durationpb.New(data.PublishTimeout), MaxRetry: data.MaxRetry})
	return err
}
