package repo

import (
	"context"

	"im/internal/biz/model"
)

type PrivateChatMessageRepo interface {
	Get(ctx context.Context, messageID int64) (*model.PrivateChatMessage, error)
	Map(ctx context.Context, messageIDs []int64) (map[int64]*model.PrivateChatMessage, error)
	GetByClientMessageID(ctx context.Context, senderID int64, clientMessageID string) (*model.PrivateChatMessage, error)
	Save(ctx context.Context, message *model.PrivateChatMessage) (*model.PrivateChatMessage, error)
	List(ctx context.Context, userID, counterpartID, afterID int64, limit, offset int) ([]*model.PrivateChatMessage, int64, error)
	Revoke(ctx context.Context, messageID, userID int64) (*model.PrivateChatMessage, error)
}
