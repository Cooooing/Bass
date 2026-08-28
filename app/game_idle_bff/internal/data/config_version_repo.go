package data

import (
	commonclient "common/pkg/client"
	"context"
	"game_idle_bff/internal/biz/repo"
	"time"

	"github.com/redis/go-redis/v9"
)

var _ repo.ConfigVersionRepo = (*ConfigVersionRepo)(nil)

const configVersionRedisTTL = time.Hour

type ConfigVersionRepo struct {
	redisClient *commonclient.RedisClient
	redisKey    string
}

func NewConfigVersionRepo(redisClient *commonclient.RedisClient) repo.ConfigVersionRepo {
	return &ConfigVersionRepo{
		redisClient: redisClient,
		redisKey:    "game_idle_bff:config:version",
	}
}

func (r *ConfigVersionRepo) Get(ctx context.Context) (string, error) {
	value, err := r.redisClient.Client.Get(ctx, r.redisKey).Result()
	if err == redis.Nil {
		return "", nil
	}
	return value, err
}

func (r *ConfigVersionRepo) Save(ctx context.Context, version string) error {
	return r.redisClient.Client.Set(ctx, r.redisKey, version, configVersionRedisTTL).Err()
}
