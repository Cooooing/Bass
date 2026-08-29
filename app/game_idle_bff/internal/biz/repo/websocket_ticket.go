package repo

import (
	"context"
	"time"
)

// WebSocketTicketRepo 管理 WS 建连凭证。
type WebSocketTicketRepo interface {
	Save(ctx context.Context, characterID int64, ticket string, ttl time.Duration) error
	Get(ctx context.Context, characterID int64) (string, time.Duration, error)
}
