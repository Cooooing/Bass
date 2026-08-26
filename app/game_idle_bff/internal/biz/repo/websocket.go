package repo

import (
	"context"
	"game_idle_bff/internal/biz/model"
)

type WebSocketEventHandler func(ctx context.Context, event *model.WebSocketEvent) error

type WebSocketEventSubscription interface {
	Unsubscribe() error
}

type WebSocketEventRepo interface {
	Subscribe(ctx context.Context, handler WebSocketEventHandler) (WebSocketEventSubscription, error)
}
