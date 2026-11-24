package repo

import (
	"context"

	"im/internal/biz/model"
)

type GlobalChatMuteRepo interface {
	Get(ctx context.Context, userID int64) (*model.GlobalChatMute, error)
	Save(ctx context.Context, mute *model.GlobalChatMute) error
	Delete(ctx context.Context, userID int64) error
}
