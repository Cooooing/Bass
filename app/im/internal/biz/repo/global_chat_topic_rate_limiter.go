package repo

import (
	"context"
	"time"
)

// GlobalChatTopicRateLimiter 限制聊天室话题设置频率。
type GlobalChatTopicRateLimiter interface {
	Limited(ctx context.Context, userID int64) (bool, error)
	Mark(ctx context.Context, userID int64, ttl, cooldown time.Duration) error
}
