package repo

import (
	"context"
	"fmt"
	"time"

	"common/pkg/client"

	bizrepo "im/internal/biz/repo"
)

var _ bizrepo.GlobalChatBarrageRateLimiter = (*GlobalChatBarrageRateLimiter)(nil)

// GlobalChatBarrageRateLimiter 使用 Redis 短期键限制同一账号的成功弹幕频率。
type GlobalChatBarrageRateLimiter struct {
	redis *client.RedisClient
}

func NewGlobalChatBarrageRateLimiter(redisClient *client.RedisClient) bizrepo.GlobalChatBarrageRateLimiter {
	return &GlobalChatBarrageRateLimiter{redis: redisClient}
}

func (r *GlobalChatBarrageRateLimiter) Limited(ctx context.Context, userID int64) (bool, error) {
	if r == nil || r.redis == nil || r.redis.Client == nil || userID <= 0 {
		return false, fmt.Errorf("global chat barrage rate limit request is invalid")
	}
	return r.redis.Client.Exists(ctx, r.userKey(userID)).Val() > 0, nil
}

func (r *GlobalChatBarrageRateLimiter) Mark(ctx context.Context, userID int64, interval time.Duration) error {
	if r == nil || r.redis == nil || r.redis.Client == nil || userID <= 0 || interval <= 0 {
		return fmt.Errorf("global chat barrage rate limit mark request is invalid")
	}
	return r.redis.Client.Set(ctx, r.userKey(userID), "1", interval).Err()
}

func (*GlobalChatBarrageRateLimiter) userKey(userID int64) string {
	return fmt.Sprintf("im:global_chat_barrage:user:{user_id:%d}", userID)
}
