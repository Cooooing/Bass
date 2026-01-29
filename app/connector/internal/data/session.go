package data

import (
	"common/pkg/constant"
	"connector/internal/biz/cache"
	"context"
)

type SessionCache struct {
	*BaseData
}

func NewSessionCache(BaseData *BaseData) cache.SessionCache {
	return &SessionCache{
		BaseData: BaseData,
	}
}

func (c *SessionCache) SetSessions(ctx context.Context, sessionIds []string) error {
	members := make([]interface{}, len(sessionIds))
	for i, sessionId := range sessionIds {
		members[i] = sessionId
	}
	return c.Redis.Client.SAdd(ctx, constant.ConnectorSession, members...).Err()
}

func (c *SessionCache) RemoveSessions(ctx context.Context, sessionIds []string) error {
	members := make([]interface{}, len(sessionIds))
	for i, sessionId := range sessionIds {
		members[i] = sessionId
	}
	return c.Redis.Client.SRem(ctx, constant.ConnectorSession, members...).Err()
}
