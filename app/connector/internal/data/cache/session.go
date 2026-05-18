package cache

import (
	commonClient "common/pkg/client"
	"common/pkg/constant"
	"connector/internal/biz/cache"
	"connector/internal/conf"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type SessionCache struct {
	conf  *conf.Bootstrap
	log   *log.Helper
	redis *commonClient.RedisClient
}

func NewSessionCache(
	conf *conf.Bootstrap,
	logger log.Logger,
	redis *commonClient.RedisClient,
) cache.SessionCache {
	return &SessionCache{
		conf:  conf,
		log:   log.NewHelper(logger),
		redis: redis,
	}
}

func (c *SessionCache) SetSessions(ctx context.Context, sessionIds []string) error {
	members := make([]interface{}, len(sessionIds))
	for i, sessionId := range sessionIds {
		members[i] = sessionId
	}
	return c.redis.Client.SAdd(ctx, constant.ConnectorSession, members...).Err()
}

func (c *SessionCache) RemoveSessions(ctx context.Context, sessionIds []string) error {
	members := make([]interface{}, len(sessionIds))
	for i, sessionId := range sessionIds {
		members[i] = sessionId
	}
	return c.redis.Client.SRem(ctx, constant.ConnectorSession, members...).Err()
}
