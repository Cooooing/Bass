package repo

import (
	"context"
	"time"
)

// GlobalChatTopicRateLimiter 限制聊天室话题设置频率。
type GlobalChatTopicRateLimiter interface {
	Limited(context.Context, int64) (bool, error)
	Mark(context.Context, int64, time.Duration, time.Duration) error
}
