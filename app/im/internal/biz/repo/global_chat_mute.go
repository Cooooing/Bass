package repo

import (
	"context"

	"im/internal/biz/model"
)

type GlobalChatMuteRepo interface {
	Get(context.Context, int64) (*model.GlobalChatMute, error)
	Save(context.Context, *model.GlobalChatMute) error
	Delete(context.Context, int64) error
}
