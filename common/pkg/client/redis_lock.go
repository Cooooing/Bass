package client

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const redisLockReleaseScript = `
if redis.call("GET", KEYS[1]) == ARGV[1] then
	return redis.call("DEL", KEYS[1])
end
return 0
`

// RedisLock 提供基于 Redis 的短期互斥锁。
type RedisLock struct {
	redisClient *RedisClient
}

// RedisLockHandle 表示一次成功获取的 Redis 锁。
type RedisLockHandle struct {
	key         string
	token       string
	redisClient *RedisClient
}

func NewRedisLock(redisClient *RedisClient) *RedisLock {
	return &RedisLock{redisClient: redisClient}
}

func (l *RedisLock) TryAcquire(ctx context.Context, key string, ttl time.Duration) (*RedisLockHandle, bool, error) {
	if l == nil || l.redisClient == nil || l.redisClient.Client == nil || key == "" || ttl <= 0 {
		return nil, false, nil
	}
	token := uuid.NewString()
	ok, err := l.redisClient.Client.SetNX(ctx, key, token, ttl).Result()
	if err != nil || !ok {
		return nil, ok, err
	}
	return &RedisLockHandle{
		key:         key,
		token:       token,
		redisClient: l.redisClient,
	}, true, nil
}

func (h *RedisLockHandle) Release(ctx context.Context) error {
	if h == nil || h.redisClient == nil || h.redisClient.Client == nil || h.key == "" || h.token == "" {
		return nil
	}
	return h.redisClient.Client.Eval(ctx, redisLockReleaseScript, []string{h.key}, h.token).Err()
}
