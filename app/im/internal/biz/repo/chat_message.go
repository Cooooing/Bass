package repo

import (
	"common/api/gen/common"
	"context"
	"im/internal/biz/model"
	"im/internal/data/ent/gen"
)

type ChatMessageRepo interface {
	Save(ctx context.Context, tx *gen.Client, chatMessage *model.ChatMessage) (*model.ChatMessage, error)

	// UpdateStatus(ctx context.Context, tx *gen.Client, messageId int64, status model.ChatMessageStatus) (*model.ChatMessage, error)

	GetOne(ctx context.Context, tx *gen.Client, req *ChatMessageGetReq) (*model.ChatMessage, error)
	GetList(ctx context.Context, tx *gen.Client, req *ChatMessageGetReq) ([]*model.ChatMessage, error)
	GetPage(ctx context.Context, tx *gen.Client, page *common.PageRequest, req *ChatMessageGetReq) ([]*model.ChatMessage, *common.PageReply, error)
}

type ChatMessageGetReq struct {
}
