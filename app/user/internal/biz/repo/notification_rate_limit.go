package repo

import "context"

type NotificationRateLimitState struct {
	Limited           bool
	RetryAfterSeconds int64
}

type NotificationRateLimitClient interface {
	CheckEmail(ctx context.Context, email string) (*NotificationRateLimitState, error)
	CheckPhone(ctx context.Context, phone string) (*NotificationRateLimitState, error)
}
