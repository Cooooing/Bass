package repo

import (
	"common/api/gen/common"
	"context"
	"im/internal/biz/model"
)

type ChatSessionRepo interface {
	Save(ctx context.Context, chatSession *model.ChatSession) (*model.ChatSession, error)

	UpdateLastReadMessage(ctx context.Context, chatSessionId int64, messageId int64, operationReadCount int32) (*model.ChatSession, error)
	UpdateMuted(ctx context.Context, chatSessionId int64, muted bool) (*model.ChatSession, error)
	UpdatePinned(ctx context.Context, chatSessionId int64, pinned bool) (*model.ChatSession, error)

	Get(ctx context.Context, req *ChatSessionGetReq) (*model.ChatSession, error)
	GetList(ctx context.Context, req *ChatSessionGetReq) ([]*model.ChatSession, error)
	GetPage(ctx context.Context, page *common.PageRequest, req *ChatSessionGetReq) ([]*model.ChatSession, *common.PageReply, error)
}

type ChatSessionGetReq struct {
}
