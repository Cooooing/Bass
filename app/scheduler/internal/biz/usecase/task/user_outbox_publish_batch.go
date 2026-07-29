package task

import (
	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	"context"
	"encoding/json"
	schedulerenum "scheduler/internal/enum"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

type UserOutboxPublishBatch struct {
	userClient     *rpc.UserClient
	name           string
	title          string
	description    string
	defaultLimit   int32
	defaultTimeout time.Duration
	defaultRetry   int32
}

func NewUserOutboxPublishBatch(
	userClient *rpc.UserClient,
) *UserOutboxPublishBatch {
	return &UserOutboxPublishBatch{
		userClient:     userClient,
		name:           "user.outbox_publish_batch",
		title:          "用户 Outbox 批量投递",
		description:    "调用 user.OutboxService.PublishBatch 投递待发送 outbox 事件。",
		defaultLimit:   1000,
		defaultTimeout: 1 * time.Minute,
		defaultRetry:   3,
	}
}

func (t *UserOutboxPublishBatch) Name() string {
	return t.name
}

func (t *UserOutboxPublishBatch) Title() string {
	return t.title
}

func (t *UserOutboxPublishBatch) Description() string {
	return t.description
}

func (t *UserOutboxPublishBatch) DefaultScheduledTasks() []*DefaultScheduledTask {
	return []*DefaultScheduledTask{
		{
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

func (t *UserOutboxPublishBatch) DefaultDelayedTasks() []*DefaultDelayedTask {
	return []*DefaultDelayedTask{
		{
			Title:         "延迟用户 Outbox 批量投递",
			Description:   t.Description(),
			Enabled:       true,
			Timeout:       1 * time.Minute,
			MaxAttempts:   3,
			MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteAll,
		},
	}
}

type userOutboxPublishBatchPayload struct {
	Limit          int32         `json:"limit"`
	PublishTimeout time.Duration `json:"publish_timeout"`
	MaxRetry       int32         `json:"max_retry"`
}

func (t *UserOutboxPublishBatch) Execute(ctx context.Context, payload string) error {
	limit := t.defaultLimit
	publishTimeout := t.defaultTimeout
	maxRetry := t.defaultRetry
	payload = strings.TrimSpace(payload)
	if payload != "" && payload != "{}" {
		data := &userOutboxPublishBatchPayload{}
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
	_, err := t.userClient.Outbox.PublishBatch(ctx, &userv1.PublishOutboxEvents_Req{
		Limit:          limit,
		PublishTimeout: durationpb.New(publishTimeout),
		MaxRetry:       maxRetry,
	})
	return err
}
