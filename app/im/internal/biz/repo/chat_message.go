package repo

import (
	"common/api/gen/common"
	"context"
	"im/internal/biz/model"
)

type ChatMessageRepo interface {
	Save(ctx context.Context, chatMessage *model.ChatMessage) (*model.ChatMessage, error)

	Get(ctx context.Context, req *ChatMessageGetReq) (*model.ChatMessage, error)
	GetList(ctx context.Context, req *ChatMessageGetReq) ([]*model.ChatMessage, error)
	Page(ctx context.Context, page *common.PageRequest, req *ChatMessageGetReq) ([]*model.ChatMessage, *common.PageReply, error)
}

type ChatMessageGetReq struct {
}
