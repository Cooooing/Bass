package repo

import (
	"common/pkg/client"
	"context"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
)

var _ bizrepo.GameIdleEventRepo = (*GameIdleEventRepo)(nil)

// GameIdleEventRepo 发布已经编码完成的挂机游戏事件消息。
type GameIdleEventRepo struct {
	natsClient *client.NatsClient
}

func NewGameIdleEventRepo(natsClient *client.NatsClient) bizrepo.GameIdleEventRepo {
	return &GameIdleEventRepo{
		natsClient: natsClient,
	}
}

func (r *GameIdleEventRepo) Publish(ctx context.Context, message *model.GameIdleEventMessage) error {
	return r.natsClient.Publish(ctx, message.Subject, &client.Message{
		Subject: message.Subject,
		Data:    message.Data,
		Header:  message.Header,
	})
}
