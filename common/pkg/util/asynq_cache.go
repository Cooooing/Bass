package util

import (
	"common/pkg/client"
	"common/pkg/constant"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type AsynqCache struct {
	log   *log.Helper
	redis *client.RedisClient
}

func NewAsynqCache(log *log.Helper, redis *client.RedisClient) *AsynqCache {
	return &AsynqCache{
		log:   log,
		redis: redis,
	}
}

func (c *AsynqCache) SetAsynqTaskVersion(ctx context.Context, taskName string, version int64, expire time.Duration) error {
	_, err := c.redis.Client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		err := pipe.HSet(ctx, constant.AsynqTaskVersion, taskName, version).Err()
		if err != nil {
			return err
		}
		err = pipe.Expire(ctx, constant.AsynqTaskVersion, expire).Err()
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (c *AsynqCache) GetAsynqTaskVersion(ctx context.Context, taskName string) (int64, error) {
	return c.redis.Client.HGet(ctx, constant.AsynqTaskVersion, taskName).Int64()
}
