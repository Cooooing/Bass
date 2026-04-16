package repo

import (
	"common/api/gen/common"
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
	GetPage(ctx context.Context, tx *gen.Client, page *common.PageRequest, req *ChatSessionGetReq) ([]*model.ChatSession, *common.PageReply, error)
}

type ChatSessionGetReq struct {
}
