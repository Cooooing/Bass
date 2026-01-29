package cache

import "context"

type SessionCache interface {
	SetSessions(ctx context.Context, sessionIds []string) error
	RemoveSessions(ctx context.Context, sessionIds []string) error
}
