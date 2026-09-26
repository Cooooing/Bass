package repo

import (
	"context"
	"time"
)

// WebSocketTicketRepo 管理 WS 建连凭证。
type WebSocketTicketRepo interface {
	Save(ctx context.Context, characterID int64, ticket string, ttl time.Duration) error
	Consume(ctx context.Context, characterID int64, ticket string) (bool, error)
}
