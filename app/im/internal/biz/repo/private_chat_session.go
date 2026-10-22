package repo

import (
	"context"

	"im/internal/biz/model"
)

type PrivateChatSessionRepo interface {
	Get(context.Context, int64, int64) (*model.PrivateChatSession, error)
	List(context.Context, int64, int, int) ([]*model.PrivateChatSession, int64, error)
	CountUnread(context.Context, int64) (uint32, error)
	Save(context.Context, *model.PrivateChatSession) (*model.PrivateChatSession, error)
	UpdateAfterSend(context.Context, int64, int64, int64, bool) error
	MarkRead(context.Context, int64, int64, int64) error
	SetMuted(context.Context, int64, int64, bool) error
	SetPinned(context.Context, int64, int64, bool) error
}
