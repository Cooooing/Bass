package repo

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"common/pkg/client"

	"im/internal/biz/model"
	bizrepo "im/internal/biz/repo"

	"github.com/redis/go-redis/v9"
)

// EmojiShareCache 将分享快照放入 Redis；表情分组事实不依赖该缓存。
type EmojiShareCache struct {
	redis     *client.RedisClient
	keyFormat string
}

func NewEmojiShareCache(redisClient *client.RedisClient) bizrepo.EmojiShareCache {
	return &EmojiShareCache{redis: redisClient, keyFormat: "im:emoji_share:{token_sha256:%s}"}
}

func (c *EmojiShareCache) Save(ctx context.Context, code string, snapshot *model.EmojiShare, ttl time.Duration) error {
	if snapshot == nil || ttl <= 0 {
		return fmt.Errorf("emoji share snapshot and TTL are required")
	}
	payload, err := json.Marshal(snapshot)
	if err != nil {
		return fmt.Errorf("marshal emoji share snapshot: %w", err)
	}
	return c.redis.Client.Set(ctx, c.key(code), payload, ttl).Err()
}

func (c *EmojiShareCache) Get(ctx context.Context, code string) (*model.EmojiShare, error) {
	payload, err := c.redis.Client.Get(ctx, c.key(code)).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var snapshot model.EmojiShare
	if err = json.Unmarshal(payload, &snapshot); err != nil {
		return nil, fmt.Errorf("unmarshal emoji share snapshot: %w", err)
	}
	return &snapshot, nil
}

func (c *EmojiShareCache) key(code string) string {
	digest := sha256.Sum256([]byte(code))
	return fmt.Sprintf(c.keyFormat, hex.EncodeToString(digest[:]))
}
