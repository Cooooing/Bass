package repo

import (
	"context"

	"im/internal/biz/model"
)

type GlobalChatMessageRepo interface {
	Get(ctx context.Context, messageID int64) (*model.GlobalChatMessage, error)
	GetByClientMessageID(ctx context.Context, senderID int64, clientMessageID string) (*model.GlobalChatMessage, error)
	Save(ctx context.Context, message *model.GlobalChatMessage) (*model.GlobalChatMessage, error)
	List(ctx context.Context, afterID int64, limit, offset int) ([]*model.GlobalChatMessage, int64, error)
	Revoke(ctx context.Context, messageID, userID int64) (*model.GlobalChatMessage, error)
}
