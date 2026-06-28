package client

import (
	"common/pkg/util"
	"common/proto/gen/common"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/types/known/durationpb"
	"log/slog"
)

// RedisClient 封装 Redis 客户端
type RedisClient struct {
	log    *util.LogHelper
	Client *redis.Client
}

// NewRedisClient 初始化 Redis 客户端
func NewRedisClient(logger *slog.Logger, conf *common.Redis) (*RedisClient, func(), error) {
	helper := util.NewLogHelper(logger)

	// 默认值
	if conf.DialTimeout == nil {
		conf.DialTimeout = durationpb.New(5 * time.Second)
	}
	if conf.ReadTimeout == nil {
		conf.ReadTimeout = durationpb.New(3 * time.Second)
	}
	if conf.WriteTimeout == nil {
		conf.WriteTimeout = durationpb.New(3 * time.Second)
	}
	if conf.PoolSize == 0 {
		conf.PoolSize = 10
	}
	if conf.PoolTimeout == nil {
		conf.PoolTimeout = durationpb.New(5 * time.Second)
	}
	if conf.ConnMaxIdleTime == nil {
		conf.ConnMaxIdleTime = durationpb.New(10 * time.Minute)
	}
	if conf.ConnMaxLifeTime == nil {
		conf.ConnMaxLifeTime = durationpb.New(30 * time.Minute)
	}

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
