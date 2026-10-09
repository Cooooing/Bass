package repo

import (
	"context"

	"bff_game_idle/internal/biz/model"
)

type WebSocketEventHandler func(ctx context.Context, event *model.WebSocketEvent) error

type WebSocketEventSubscription interface {
	Unsubscribe() error
}

type WebSocketEventRepo interface {
	Subscribe(ctx context.Context, handler WebSocketEventHandler) (WebSocketEventSubscription, error)
}
