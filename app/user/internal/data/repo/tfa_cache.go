package repo

import (
	"common/pkg/client"
	"common/pkg/constant"
	"context"
	"strconv"
	"time"
	"user/internal/biz/repo"
)

var _ repo.TfaSecretCache = (*TfaSecretCache)(nil)

type TfaSecretCache struct {
	redisClient *client.RedisClient
}

func NewTfaSecretCache(redisClient *client.RedisClient) repo.TfaSecretCache {
	return &TfaSecretCache{redisClient: redisClient}
}

func (c *TfaSecretCache) Save(ctx context.Context, userID int64, secret string, ttl time.Duration) error {
	return c.redisClient.Client.SetEx(ctx, constant.GetKeyTwoFactorAuth(strconv.FormatInt(userID, 10)), secret, ttl).Err()
}

func (c *TfaSecretCache) Get(ctx context.Context, userID int64) (string, error) {
	return c.redisClient.Client.Get(ctx, constant.GetKeyTwoFactorAuth(strconv.FormatInt(userID, 10))).Result()
}
