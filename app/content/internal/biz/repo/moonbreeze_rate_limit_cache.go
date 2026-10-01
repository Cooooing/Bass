package repo

import (
	"context"
	"time"
)

// MoonbreezeRateLimitCache keeps the transient publish window in Redis. The
// window argument is also the key TTL; no application invalidation is needed.
// A cache failure is returned so publishing fails closed instead of bypassing
// the per-account rate limit.
type MoonbreezeRateLimitCache interface {
	Allow(ctx context.Context, authorID int64, window time.Duration, maxCount int64) (*MoonbreezeRateLimitState, error)
}

type MoonbreezeRateLimitState struct {
	Allowed    bool
	RetryAfter time.Duration
}
