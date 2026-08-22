package module

import (
	commonclient "common/pkg/client"
	"fmt"
	"log/slog"

	"github.com/google/wire"
)

// Infrastructure 保存单体内可复用的基础设施客户端。
type Infrastructure struct {
	observer *commonclient.Observer
	redis    *commonclient.RedisClient
	nats     *commonclient.NatsClient
}

// InfrastructureProviderSet 提供模块化单体模式下复用的基础设施客户端。
var InfrastructureProviderSet = wire.NewSet(
	ProvideObserver,
	ProvideRedisClient,
	ProvideNatsClient,
	commonclient.NewRedisLock,
)

// NewInfrastructure 按需创建单体共享基础设施客户端。
func NewInfrastructure(logger *slog.Logger, config *RuntimeConfig) (*Infrastructure, func(), error) {
	if config == nil {
		config = &RuntimeConfig{}
	}
	observer := commonclient.NewObservability(logger, config.Server)
	var cleanups []func()
	cleanup := func() {
		for i := len(cleanups) - 1; i >= 0; i-- {
			cleanups[i]()
		}
	}

	infra := &Infrastructure{observer: observer}
	if config.Redis != nil {
		redisClient, redisCleanup, err := commonclient.NewRedisClient(logger, config.Redis)
		if err != nil {
			cleanup()
			return nil, func() {}, err
		}
		infra.redis = redisClient
		cleanups = append(cleanups, redisCleanup)
	}
	if config.Nats != nil {
		natsClient, natsCleanup, err := commonclient.NewNatsClient(logger, config.Nats, observer)
		if err != nil {
			cleanup()
			return nil, func() {}, err
		}
		infra.nats = natsClient
		cleanups = append(cleanups, natsCleanup)
	}
	return infra, cleanup, nil
}

// ProvideObserver 返回单体共享观测能力。
func ProvideObserver(infra *Infrastructure) *commonclient.Observer {
	if infra == nil {
		return nil
	}
	return infra.observer
}

// ProvideRedisClient 返回单体共享 Redis 客户端。
func ProvideRedisClient(infra *Infrastructure) (*commonclient.RedisClient, error) {
	if infra == nil || infra.redis == nil {
		return nil, fmt.Errorf("monolith redis client is required")
	}
	return infra.redis, nil
}

// ProvideNatsClient 返回单体共享 NATS 客户端。
func ProvideNatsClient(infra *Infrastructure) (*commonclient.NatsClient, error) {
	if infra == nil || infra.nats == nil {
		return nil, fmt.Errorf("monolith nats client is required")
	}
	return infra.nats, nil
}
