package client

import (
	"common/pkg/constant"
	"common/proto/gen/common"
	"context"
	"fmt"
	"time"

	"log/slog"

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/types/known/durationpb"
)

type RedisClient struct {
	logger *slog.Logger
	Client *redis.Client
}

func NewRedisClient(
	logger *slog.Logger,
	conf *common.Redis,
) (*RedisClient, func(), error) {
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
		MaxRetries:      3,
		MinRetryBackoff: 100 * time.Millisecond,
		MaxRetryBackoff: 500 * time.Millisecond,
		OnConnect: func(ctx context.Context, conn *redis.Conn) error {
			logger.DebugContext(ctx, "redis connection created", constant.LogFieldKind, constant.LogKindRedis, constant.LogFieldAddress, conn.String())
			return nil
		},
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		_ = client.Close()
		return nil, nil, fmt.Errorf("redis ping [%s]: %w", conf.Addr, err)
	}

	logger.Info("redis connected", constant.LogFieldKind, constant.LogKindRedis, constant.LogFieldAddress, conf.Addr, constant.LogFieldDB, conf.Db)

	r := &RedisClient{
		logger: logger,
		Client: client,
	}

	return r, func() {
		if err := client.Close(); err != nil {
			logger.Error("redis close failed", constant.LogFieldKind, constant.LogKindRedis, constant.LogFieldErr, err)
		} else {
			logger.Info("redis closed", constant.LogFieldKind, constant.LogKindRedis, constant.LogFieldAddress, conf.Addr)
		}
	}, nil
}
