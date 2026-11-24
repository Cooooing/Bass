package repo

import (
	"context"

	"im/internal/biz/model"
)

type PrivateChatSessionRepo interface {
	Get(ctx context.Context, userID, counterpartID int64) (*model.PrivateChatSession, error)
	List(ctx context.Context, userID int64, limit, offset int) ([]*model.PrivateChatSession, int64, error)
	CountUnread(ctx context.Context, userID int64) (uint32, error)
	Save(ctx context.Context, session *model.PrivateChatSession) (*model.PrivateChatSession, error)
	UpdateAfterSend(ctx context.Context, userID, counterpartID, messageID int64, incrementUnread bool) error
	MarkRead(ctx context.Context, userID, counterpartID, lastReadMessageID int64) error
	SetMuted(ctx context.Context, userID, counterpartID int64, muted bool) error
	SetPinned(ctx context.Context, userID, counterpartID int64, pinned bool) error
}
