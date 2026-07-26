package task

import (
	"common/pkg/client/rpc"
	userv1 "common/proto/gen/user/v1"
	"context"
	"encoding/json"
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
		title:          "User outbox publish batch",
		description:    "Call user.OutboxService.PublishBatch to publish pending outbox events.",
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

func (t *UserOutboxPublishBatch) DefaultSchedules() []*DefaultSchedule {
	return []*DefaultSchedule{
		{
			Title:          t.Title(),
			Description:    t.Description(),
			Enabled:        true,
			CronSpec:       "0/10 * * * * ?",
			Payload:        `{"limit":1000,"publish_timeout_seconds":300,"max_retry":10}`,
			TimeoutSeconds: 30,
			AllowOverlap:   false,
			AlertEnabled:   false,
		},
	}
}

type userOutboxPublishBatchPayload struct {
	Limit                 int32 `json:"limit"`
	PublishTimeoutSeconds int64 `json:"publish_timeout_seconds"`
	MaxRetry              int32 `json:"max_retry"`
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
		if data.PublishTimeoutSeconds > 0 {
			publishTimeout = time.Duration(data.PublishTimeoutSeconds) * time.Second
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
