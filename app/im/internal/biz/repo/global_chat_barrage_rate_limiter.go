package repo

import (
	"context"
	"time"
)

type GlobalChatBarrageRateLimiter interface {
	Limited(ctx context.Context, userID int64) (bool, error)
	Mark(ctx context.Context, userID int64, ttl time.Duration) error
}
