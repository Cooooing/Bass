package client

import (
	"common/pkg/model"
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	log    *log.Helper
	Client *redis.Client
}

// NewRedisClient 初始化单机 Redis 客户端
func NewRedisClient(log *log.Helper, conf *model.RedisConf) (*RedisClient, func(), error) {
	client := redis.NewClient(&redis.Options{
		Addr:         conf.Addr,
		Password:     conf.Password,
		DB:           int(conf.Db),
		ReadTimeout:  conf.ReadTimeout.AsDuration(),
		WriteTimeout: conf.WriteTimeout.AsDuration(),
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
	log.Infof("redis: connected to [%s]", conf.Addr)

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
