package repo

import (
	cv1 "common/api/common/v1"
	"context"
	"im/internal/biz/model"
	"im/internal/data/ent/gen"
)

type ChatSessionRepo interface {
	Save(ctx context.Context, tx *gen.Client, chatSession *model.ChatSession) (*model.ChatSession, error)

	UpdateLastReadMessage(ctx context.Context, tx *gen.Client, chatSessionId int64, messageId int64, operationReadCount int32) (*model.ChatSession, error)
	UpdateMuted(ctx context.Context, tx *gen.Client, chatSessionId int64, muted bool) (*model.ChatSession, error)
	UpdatePinned(ctx context.Context, tx *gen.Client, chatSessionId int64, pinned bool) (*model.ChatSession, error)

	GetOne(ctx context.Context, tx *gen.Client, req *ChatSessionGetReq) (*model.ChatSession, error)
	GetList(ctx context.Context, tx *gen.Client, req *ChatSessionGetReq) ([]*model.ChatSession, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *ChatSessionGetReq) ([]*model.ChatSession, *cv1.PageReply, error)
}

type ChatSessionGetReq struct {
}
