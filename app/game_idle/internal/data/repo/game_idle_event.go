package repo

import (
	nats "common/pkg/client/nats"
	"context"

	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
)

var _ bizrepo.GameIdleEventRepo = (*GameIdleEventRepo)(nil)

// GameIdleEventRepo 发布已经编码完成的挂机游戏事件消息。
type GameIdleEventRepo struct {
	natsClient *nats.Client
}

func NewGameIdleEventRepo(natsClient *nats.Client) bizrepo.GameIdleEventRepo {
	return &GameIdleEventRepo{
		natsClient: natsClient,
	}
}

func (r *GameIdleEventRepo) Publish(ctx context.Context, message *model.GameIdleEventMessage) error {
	return r.natsClient.Publish(ctx, message.Subject, &nats.Message{
		Subject: message.Subject,
		Data:    message.Data,
		Header:  message.Header,
	})
}
