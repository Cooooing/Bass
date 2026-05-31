package repo

import (
	"context"
	"time"
)

type NotificationRateLimitState struct {
	Limited        bool
	RetryAfter     time.Duration
	RemainingCount int64
}

type NotificationRateLimitClient interface {
	CheckEmail(ctx context.Context, email string) (*NotificationRateLimitState, error)
	CheckPhone(ctx context.Context, phone string) (*NotificationRateLimitState, error)
}
