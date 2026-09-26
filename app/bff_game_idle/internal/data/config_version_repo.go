package data

import (
	"bff_game_idle/internal/biz/repo"
	commonclient "common/pkg/client"
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var _ repo.ConfigVersionRepo = (*ConfigVersionRepo)(nil)

type ConfigVersionRepo struct {
	redisClient *commonclient.RedisClient
	redisKey    string
	redisTTL    time.Duration
}

func NewConfigVersionRepo(redisClient *commonclient.RedisClient) repo.ConfigVersionRepo {
	return &ConfigVersionRepo{
		redisClient: redisClient,
		redisKey:    "bff_game_idle:config:version",
		redisTTL:    time.Hour,
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
	return r.redisClient.Client.Set(ctx, r.redisKey, version, r.redisTTL).Err()
}
