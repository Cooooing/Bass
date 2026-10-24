package repo

import (
	"context"
	"time"
)

type GlobalChatBarrageRateLimiter interface {
	Limited(context.Context, int64) (bool, error)
	Mark(context.Context, int64, time.Duration) error
}
