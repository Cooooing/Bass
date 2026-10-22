package repo

import (
	"context"

	"im/internal/biz/model"
)

type PrivateChatMessageRepo interface {
	Get(context.Context, int64) (*model.PrivateChatMessage, error)
	Map(context.Context, []int64) (map[int64]*model.PrivateChatMessage, error)
	GetByClientMessageID(context.Context, int64, string) (*model.PrivateChatMessage, error)
	Save(context.Context, *model.PrivateChatMessage) (*model.PrivateChatMessage, error)
	List(context.Context, int64, int64, int64, int, int) ([]*model.PrivateChatMessage, int64, error)
	Revoke(context.Context, int64, int64) (*model.PrivateChatMessage, error)
}
