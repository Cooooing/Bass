package repo

import "context"

// CharacterSessionRepo 管理角色在线会话缓存。
type CharacterSessionRepo interface {
	Online(ctx context.Context, characterID int64, sessionID string, ttlSeconds int64) (string, error)
	Ping(ctx context.Context, characterID int64, sessionID string, ttlSeconds int64) (bool, error)
	Offline(ctx context.Context, characterID int64, sessionID string) (bool, error)
}
