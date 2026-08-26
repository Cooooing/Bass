package repo

import (
	"context"
	"game_idle/internal/biz/model"
)

// GameIdleEventRepo 发布挂机游戏内部事件。
type GameIdleEventRepo interface {
	Publish(ctx context.Context, event *model.GameIdleEvent) error
}
