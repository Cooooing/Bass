package cache

import (
	"context"
)

type SessionCache interface {
	SetTicket(ctx context.Context, ticket string, userId int64) error
	GetTicket(ctx context.Context, ticket string) (int64, error)

	SetSession(ctx context.Context, sessionId string, userId int64, nodeKey string) error
	SetNodeSessionIds(ctx context.Context, nodeKey string, sessionIds []string) error
	RenewalSessions(ctx context.Context, sessionIds []string) error
	GetNodeKeyByUserIds(ctx context.Context, userIds []int64) ([]string, error)
}
