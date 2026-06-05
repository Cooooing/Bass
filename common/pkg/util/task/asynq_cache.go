package task

import (
	"common/pkg/client"
	"common/pkg/constant"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type AsynqCache struct {
	log         *log.Helper
	redisClient *client.RedisClient
}

func NewAsynqCache(logger log.Logger, redisClient *client.RedisClient) *AsynqCache {
	return &AsynqCache{
		log:         log.NewHelper(logger),
		redisClient: redisClient,
	}
}

func (c *AsynqCache) SetAsynqTaskVersion(ctx context.Context, taskName string, version int64, expire time.Duration) error {
	_, err := c.redisClient.Client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		err := pipe.HSet(ctx, constant.AsynqTaskVersion, taskName, version).Err()
		if err != nil {
			return err
		}
		err = pipe.HExpire(ctx, constant.AsynqTaskVersion, expire, taskName).Err()
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (c *AsynqCache) GetAsynqTaskVersion(ctx context.Context, taskName string) (int64, error) {
	return c.redisClient.Client.HGet(ctx, constant.AsynqTaskVersion, taskName).Int64()
}

func (c *AsynqCache) SetAsynqTaskExpire(ctx context.Context, taskName string, expire time.Duration) error {
	return c.redisClient.Client.HExpire(ctx, constant.AsynqTaskVersion, expire, taskName).Err()
}
