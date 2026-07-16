package repo

import (
	"common/pkg/client"
	"common/pkg/constant"
	"context"
	"errors"
	"strconv"
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

func (c *TotpSecretCache) Save(ctx context.Context, req *repo.TotpSecretCacheSaveReq) (*repo.TotpSecretCacheSaveResponse, error) {
	err := c.redisClient.Client.SetEx(ctx, constant.GetKeyTotpSecret(strconv.FormatInt(req.UserID, 10)), req.Secret, req.TTL).Err()
	if err != nil {
		return nil, err
	}
	return &repo.TotpSecretCacheSaveResponse{}, nil
}

func (c *TotpSecretCache) Get(ctx context.Context, req *repo.TotpSecretCacheGetReq) (*repo.TotpSecretCacheGetResponse, error) {
	secret, err := c.redisClient.Client.Get(ctx, constant.GetKeyTotpSecret(strconv.FormatInt(req.UserID, 10))).Result()
	if errors.Is(err, redis.Nil) {
		return &repo.TotpSecretCacheGetResponse{}, nil
	}
	if err != nil {
		return nil, err
	}
	return &repo.TotpSecretCacheGetResponse{Secret: secret}, nil
}
