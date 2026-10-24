package repo

import (
	"context"
	"fmt"
	"time"

	"common/pkg/client"

	bizrepo "im/internal/biz/repo"
)

var _ bizrepo.GlobalChatTopicRateLimiter = (*GlobalChatTopicRateLimiter)(nil)

// GlobalChatTopicRateLimiter 通过 Redis 短期键限制话题设置频率。
type GlobalChatTopicRateLimiter struct {
	redis *client.RedisClient
}

func NewGlobalChatTopicRateLimiter(redisClient *client.RedisClient) bizrepo.GlobalChatTopicRateLimiter {
	return &GlobalChatTopicRateLimiter{redis: redisClient}
}

func (r *GlobalChatTopicRateLimiter) Limited(ctx context.Context, userID int64) (bool, error) {
	if r == nil || r.redis == nil || r.redis.Client == nil || userID <= 0 {
		return false, fmt.Errorf("global chat topic rate limit request is invalid")
	}
	return r.redis.Client.Exists(ctx, r.userKey(userID), r.globalKey()).Val() > 0, nil
}

func (r *GlobalChatTopicRateLimiter) Mark(ctx context.Context, userID int64, userInterval, globalInterval time.Duration) error {
	if r == nil || r.redis == nil || r.redis.Client == nil || userID <= 0 || userInterval <= 0 || globalInterval <= 0 {
		return fmt.Errorf("global chat topic rate limit mark request is invalid")
	}
	return r.redis.Client.Eval(ctx, `
redis.call("SET", KEYS[1], "1", "EX", ARGV[1])
redis.call("SET", KEYS[2], "1", "EX", ARGV[2])
return 1
`, []string{r.userKey(userID), r.globalKey()}, int64(userInterval.Seconds()), int64(globalInterval.Seconds())).Err()
}

func (*GlobalChatTopicRateLimiter) userKey(userID int64) string {
	return fmt.Sprintf("im:global_chat_topic:user:{user_id:%d}", userID)
}

func (*GlobalChatTopicRateLimiter) globalKey() string {
	return "im:global_chat_topic:global"
}
