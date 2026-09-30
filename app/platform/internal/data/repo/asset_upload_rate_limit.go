package repo

import (
	"common/pkg/client"
	"context"
	"fmt"
	"platform/internal/biz/repo"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var _ repo.AssetUploadRateLimitCache = (*AssetUploadRateLimitCache)(nil)

// AssetUploadRateLimitCache uses one Redis sorted set per uploader. The Lua
// script removes expired entries and consumes a new allowance atomically.
type AssetUploadRateLimitCache struct {
	redisClient *client.RedisClient
	allowScript *redis.Script
}

func NewAssetUploadRateLimitCache(redisClient *client.RedisClient) repo.AssetUploadRateLimitCache {
	return &AssetUploadRateLimitCache{
		redisClient: redisClient,
		allowScript: redis.NewScript(`
local key = KEYS[1]
local now = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
local max_count = tonumber(ARGV[3])
local member = ARGV[4]
redis.call("ZREMRANGEBYSCORE", key, 0, now - window)
if redis.call("ZCARD", key) >= max_count then
	local first = redis.call("ZRANGE", key, 0, 0, "WITHSCORES")
	local retry_after = 0
	if first[2] ~= nil then
		retry_after = tonumber(first[2]) + window - now
		if retry_after < 0 then
			retry_after = 0
		end
	end
	return {0, retry_after}
end
redis.call("ZADD", key, now, member)
redis.call("PEXPIRE", key, window)
return {1, 0}
`),
	}
}

func (c *AssetUploadRateLimitCache) Allow(
	ctx context.Context,
	uploadByID int64,
	window time.Duration,
	maxCount int64,
) (*repo.AssetUploadRateLimitState, error) {
	if c == nil || c.redisClient == nil || c.redisClient.Client == nil {
		return nil, fmt.Errorf("asset upload rate limit Redis client is required")
	}
	if uploadByID <= 0 || window <= 0 || maxCount <= 0 {
		return nil, fmt.Errorf("invalid asset upload rate limit specification")
	}
	result, err := c.allowScript.Run(
		ctx,
		c.redisClient.Client,
		[]string{fmt.Sprintf("platform:asset_upload_rate_limit:{user:%d}", uploadByID)},
		time.Now().UnixMilli(),
		window.Milliseconds(),
		maxCount,
		uuid.NewString(),
	).Int64Slice()
	if err != nil {
		return nil, err
	}
	state := &repo.AssetUploadRateLimitState{}
	if len(result) > 0 {
		state.Allowed = result[0] == 1
	}
	if len(result) > 1 && result[1] > 0 {
		state.RetryAfter = time.Duration(result[1]) * time.Millisecond
	}
	return state, nil
}
