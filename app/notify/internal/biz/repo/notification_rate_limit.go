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
	Allow(ctx context.Context, spec *NotificationRateLimitSpec) (bool, error)
	Check(ctx context.Context, spec *NotificationRateLimitSpec) (*NotificationRateLimitState, error)
}
