package repo

import (
	"common/pkg/client"
	"content/internal/biz/repo"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var _ repo.MoonbreezeRateLimitCache = (*MoonbreezeRateLimitCache)(nil)

// MoonbreezeRateLimitCache atomically trims and consumes each author's window.
type MoonbreezeRateLimitCache struct {
	redisClient *client.RedisClient
	allowScript *redis.Script
}

func NewMoonbreezeRateLimitCache(redisClient *client.RedisClient) repo.MoonbreezeRateLimitCache {
	return &MoonbreezeRateLimitCache{
		redisClient: redisClient,
		allowScript: redis.NewScript(`
local key = KEYS[1]
local now = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
local maximum = tonumber(ARGV[3])
local member = ARGV[4]
redis.call("ZREMRANGEBYSCORE", key, 0, now - window)
if redis.call("ZCARD", key) >= maximum then
  local first = redis.call("ZRANGE", key, 0, 0, "WITHSCORES")
  return {0, math.max(0, tonumber(first[2]) + window - now)}
end
redis.call("ZADD", key, now, member)
redis.call("PEXPIRE", key, window)
return {1, 0}
	`),
	}
}

func (c *MoonbreezeRateLimitCache) Allow(ctx context.Context, authorID int64, window time.Duration, maxCount int64) (*repo.MoonbreezeRateLimitState, error) {
	if c == nil || c.redisClient == nil || c.redisClient.Client == nil || authorID <= 0 || window <= 0 || maxCount <= 0 {
		return nil, fmt.Errorf("invalid moonbreeze rate limit specification")
	}
	result, err := c.allowScript.Run(
		ctx,
		c.redisClient.Client,
		[]string{c.publishRateLimitKey(authorID)},
		time.Now().UnixMilli(),
		window.Milliseconds(),
		maxCount,
		uuid.NewString(),
	).Int64Slice()
	if err != nil {
		return nil, err
	}
	state := &repo.MoonbreezeRateLimitState{Allowed: len(result) > 0 && result[0] == 1}
	if len(result) > 1 && result[1] > 0 {
		state.RetryAfter = time.Duration(result[1]) * time.Millisecond
	}
	return state, nil
}

func (*MoonbreezeRateLimitCache) publishRateLimitKey(authorID int64) string {
	return fmt.Sprintf("content:moonbreeze:publish_rate_limit:{author_id:%d}", authorID)
}
