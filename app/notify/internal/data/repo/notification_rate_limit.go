package repo

import (
	"common/pkg/client"
	"context"
	"fmt"
	bizrepo "notify/internal/biz/repo"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var _ bizrepo.NotificationRateLimitCache = (*NotificationRateLimitCache)(nil)

const notificationRateLimitKeyPrefix = "notify:rate_limit"

var notificationRateLimitAllowScript = redis.NewScript(`
local key = KEYS[1]
local now = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
local max_count = tonumber(ARGV[3])
local member = ARGV[4]
redis.call("ZREMRANGEBYSCORE", key, 0, now - window)
local count = redis.call("ZCARD", key)
if count >= max_count then
	redis.call("PEXPIRE", key, window)
	return 0
end
redis.call("ZADD", key, now, member)
redis.call("PEXPIRE", key, window)
return 1
`)

var notificationRateLimitCheckScript = redis.NewScript(`
local key = KEYS[1]
local now = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
local max_count = tonumber(ARGV[3])
local min_score = "(" .. tostring(now - window)
local count = redis.call("ZCOUNT", key, min_score, "+inf")
if count <= 0 then
	return {0, 0, max_count}
end
local remaining = max_count - count
if remaining < 0 then
	remaining = 0
end
if count < max_count then
	return {0, 0, remaining}
end
local first = redis.call("ZRANGEBYSCORE", key, min_score, "+inf", "WITHSCORES", "LIMIT", 0, 1)
local retry_after = 0
if first[2] ~= nil then
	retry_after = tonumber(first[2]) + window - now
	if retry_after < 0 then
		retry_after = 0
	end
end
return {1, retry_after, 0}
`)

type NotificationRateLimitCache struct {
	redisClient *client.RedisClient
}

func NewNotificationRateLimitCache(redisClient *client.RedisClient) bizrepo.NotificationRateLimitCache {
	return &NotificationRateLimitCache{redisClient: redisClient}
}

func (c *NotificationRateLimitCache) Allow(ctx context.Context, spec *bizrepo.NotificationRateLimitSpec) (bool, error) {

	key, err := notificationRateLimitKey(spec)
	if err != nil {
		return false, err
	}
	allowed, err := notificationRateLimitAllowScript.Run(
		ctx,
		c.redisClient.Client,
		[]string{key},
		time.Now().UnixMilli(), spec.
			Window.Milliseconds(), spec.
			MaxCount, uuid.NewString(),
	).Int()
	if err != nil {
		return false, err
	}
	return allowed == 1, nil
}

func (c *NotificationRateLimitCache) Check(ctx context.Context, spec *bizrepo.NotificationRateLimitSpec) (*bizrepo.NotificationRateLimitState, error) {

	key, err := notificationRateLimitKey(spec)
	if err != nil {
		return nil, err
	}
	values, err := notificationRateLimitCheckScript.Run(
		ctx,
		c.redisClient.Client,
		[]string{key},
		time.Now().UnixMilli(), spec.
			Window.Milliseconds(), spec.
			MaxCount,
	).Int64Slice()
	if err != nil {
		return nil, err
	}
	state := &bizrepo.NotificationRateLimitState{}
	if len(values) > 0 {
		state.Limited = values[0] == 1
	}
	if len(values) > 1 && values[1] > 0 {
		state.RetryAfter = time.Duration(values[1]) * time.Millisecond
	}
	if len(values) > 2 && values[2] > 0 {
		state.RemainingCount = values[2]
	}
	return state, nil
}

func notificationRateLimitKey(spec *bizrepo.NotificationRateLimitSpec) (string, error) {
	if spec == nil {
		return "", fmt.Errorf("notification rate limit spec is nil")
	}
	if spec.Channel == "" || spec.Window <= 0 || spec.MaxCount <= 0 {
		return "", fmt.Errorf("notification rate limit spec is invalid")
	}
	recipient := strings.TrimSpace(strings.ToLower(spec.Recipient))
	if recipient == "" {
		return "", fmt.Errorf("notification rate limit recipient is empty")
	}
	return fmt.Sprintf("%s:%s:%s", notificationRateLimitKeyPrefix, spec.Channel, recipient), nil
}
