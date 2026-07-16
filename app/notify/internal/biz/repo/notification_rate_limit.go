package repo

import (
	"context"
	"time"

	notifyenum "notify/internal/enum"
)

type NotificationRateLimitSpec struct {
	Channel   notifyenum.NotificationChannel
	Recipient string
	Window    time.Duration
	MaxCount  int64
}

type NotificationRateLimitState struct {
	Limited        bool
	RetryAfter     time.Duration
	RemainingCount int64
}

type NotificationRateLimitCache interface {
	Allow(ctx context.Context, req *NotificationRateLimitAllowReq) (*NotificationRateLimitAllowResponse, error)
	Check(ctx context.Context, req *NotificationRateLimitCheckReq) (*NotificationRateLimitCheckResponse, error)
}

type NotificationRateLimitAllowReq struct {
	Spec *NotificationRateLimitSpec
}

type NotificationRateLimitAllowResponse struct {
	Allowed bool
}

type NotificationRateLimitCheckReq struct {
	Spec *NotificationRateLimitSpec
}

type NotificationRateLimitCheckResponse struct {
	State *NotificationRateLimitState
}
