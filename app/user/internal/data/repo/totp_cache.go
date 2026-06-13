package repo

import (
	"common/pkg/client"
	"common/pkg/constant"
	"context"
	"errors"
	"strconv"
	"time"
	"user/internal/biz/repo"

	"github.com/redis/go-redis/v9"
)

var _ repo.TotpSecretCache = (*TotpSecretCache)(nil)

type TotpSecretCache struct {
	redisClient *client.RedisClient
}

func NewTotpSecretCache(redisClient *client.RedisClient) repo.TotpSecretCache {
	return &TotpSecretCache{redisClient: redisClient}
}

func (c *TotpSecretCache) Save(ctx context.Context, userID int64, secret string, ttl time.Duration) error {
	return c.redisClient.Client.SetEx(ctx, constant.GetKeyTotpSecret(strconv.FormatInt(userID, 10)), secret, ttl).Err()
}

func (c *TotpSecretCache) Get(ctx context.Context, userID int64) (string, error) {
	secret, err := c.redisClient.Client.Get(ctx, constant.GetKeyTotpSecret(strconv.FormatInt(userID, 10))).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	return secret, err
}
