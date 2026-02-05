package cache

import (
	"common/pkg/constant"
	"connector/internal/biz/cache"
	database "connector/internal/data/base"
	"context"
)

type SessionCache struct {
	*database.BaseData
}

func NewSessionCache(BaseData *database.BaseData) cache.SessionCache {
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
