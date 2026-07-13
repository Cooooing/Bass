package repo

import (
	"common/proto/gen/common"
	"context"
	"im/internal/biz/model"
	"im/internal/enum"
)

type ChatMessageRepo interface {
	Save(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error)
	UpdateStatus(ctx context.Context, chatMessageId int64, status enum.MessageStatus, updatedBy int64) (*model.ChatMessage, error)

	Get(ctx context.Context, req *ChatMessageGetReq) (*model.ChatMessage, error)
	List(ctx context.Context, req *ChatMessageGetReq) ([]*model.ChatMessage, error)
	Map(ctx context.Context, req *ChatMessageGetReq) (map[int64]*model.ChatMessage, error)
	Count(ctx context.Context, req *ChatMessageGetReq) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *ChatMessageGetReq) ([]*model.ChatMessage, *common.PageReply, error)
}

type ChatMessageGetReq struct {
	IDs       []int64
	SessionID *int64
	SenderID  *int64
}
