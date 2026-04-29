package client

import (
	"common/api/gen/common"
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

// RedisClient 封装 Redis 客户端
type RedisClient struct {
	log    *log.Helper
	Client *redis.Client
}

// NewRedisClient 初始化 Redis 客户端
func NewRedisClient(logger log.Logger, conf *common.Redis) (*RedisClient, func(), error) {
	helper := log.NewHelper(logger)
	client := redis.NewClient(&redis.Options{
		Addr:            conf.Addr,
		Password:        conf.Password,
		DB:              int(conf.Db),
		DialTimeout:     conf.DialTimeout.AsDuration(),
		ReadTimeout:     conf.ReadTimeout.AsDuration(),
		WriteTimeout:    conf.WriteTimeout.AsDuration(),
		PoolSize:        int(conf.PoolSize),
		MinIdleConns:    int(conf.MinIdleConns),
		PoolTimeout:     conf.PoolTimeout.AsDuration(),
		ConnMaxIdleTime: conf.ConnMaxIdleTime.AsDuration(),
		ConnMaxLifetime: conf.ConnMaxLifeTime.AsDuration(),
		MaxRetries:      3,                      // 网络闪断自动重试
		MinRetryBackoff: 100 * time.Millisecond, // 重试退避下限
		MaxRetryBackoff: 500 * time.Millisecond, // 重试退避上限
		OnConnect: func(ctx context.Context, conn *redis.Conn) error {
			helper.Infof("redis conn created: %s", conn.String())
			return nil
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		_ = client.Close()
		return nil, nil, fmt.Errorf("redis ping [%s]: %w", conf.Addr, err)
	}

	helper.Infof("redis connected: %s (db=%d)", conf.Addr, conf.Db)

	r := &RedisClient{
		log:    helper,
		Client: client,
	}

	return r, func() {
		if err := client.Close(); err != nil {
			helper.Errorf("redis close: %s", err)
		} else {
			helper.Infof("redis closed: %s", conf.Addr)
		}
	}, nil
}
