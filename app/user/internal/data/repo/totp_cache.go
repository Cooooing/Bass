package repo

import (
	"common/pkg/client"
	"context"
	"errors"
	"fmt"
	"user/internal/biz/repo"

	"github.com/redis/go-redis/v9"
)

var _ repo.TotpSecretCache = (*TotpSecretCache)(nil)

type TotpSecretCache struct {
	redisClient       *client.RedisClient
	authTotpSecretKey string
}

func NewTotpSecretCache(
	redisClient *client.RedisClient,
) repo.TotpSecretCache {
	return &TotpSecretCache{
		redisClient:       redisClient,
		authTotpSecretKey: "Auth:TotpSecret:{%d}",
	}
}

func (c *TotpSecretCache) authTotpSecretRedisKey(userID int64) string {
	return fmt.Sprintf(c.authTotpSecretKey, userID)
}

func (c *TotpSecretCache) Save(ctx context.Context, req *repo.TotpSecretCacheSaveReq) error {
	err := c.redisClient.Client.SetEx(ctx, c.authTotpSecretRedisKey(req.UserID), req.Secret, req.TTL).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *TotpSecretCache) Get(ctx context.Context, userID int64) (string, error) {
	secret, err := c.redisClient.Client.Get(ctx, c.authTotpSecretRedisKey(userID)).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return secret, nil
}
