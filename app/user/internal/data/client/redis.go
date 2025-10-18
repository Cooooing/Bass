package client

import (
	"context"
	"fmt"
	"user/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	log    *log.Helper
	Client *redis.Client
}

// NewRedisClient 初始化单机 Redis 客户端
func NewRedisClient(log *log.Helper, conf *conf.Bootstrap) (*RedisClient, func(), error) {
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Data.Redis.Addr,
		Password:     conf.Data.Redis.Password,
		DB:           int(conf.Data.Redis.Db),
		ReadTimeout:  conf.Data.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Data.Redis.WriteTimeout.AsDuration(),
	})
	ctx := context.Background()

	// 测试连接
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	r := &RedisClient{
		log:    log,
		Client: client,
	}
	log.Infof("redis: connected to [%s]", conf.Data.Redis.Addr)

	// 清理函数
	cleanup := func() {
		if err := client.Close(); err != nil {
			log.Errorf("failed to close redis client: %s", err.Error())
		} else {
			log.Infof("redis client closed")
		}
	}

	return r, cleanup, nil
}
