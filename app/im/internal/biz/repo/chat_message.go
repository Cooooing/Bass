package repo

import (
	"context"
	"im/internal/biz/base"
	"im/internal/biz/model"
	"im/internal/enum"
)

type ChatMessageRepo interface {
	Save(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error)
	UpdateStatus(ctx context.Context, req *ChatMessageUpdateStatusReq) (*model.ChatMessage, error)

	Get(ctx context.Context, query *ChatMessageQuery) (*model.ChatMessage, error)
	List(ctx context.Context, query *ChatMessageQuery) ([]*model.ChatMessage, error)
	Map(ctx context.Context, query *ChatMessageQuery) (map[int64]*model.ChatMessage, error)
	Count(ctx context.Context, query *ChatMessageQuery) (int, error)
	Page(ctx context.Context, query *ChatMessageQuery) (*ChatMessagePageResp, error)
}

type ChatMessageUpdateStatusReq struct {
	ChatMessageID int64
	Status        enum.MessageStatus
	UpdatedBy     int64
}

type ChatMessageQuery struct {
	Page      *base.PageRequest
	IDs       []int64
	SessionID *int64
	SenderID  *int64
}

type ChatMessagePageResp struct {
	Rows []*model.ChatMessage
	Page *base.PageResp
}
