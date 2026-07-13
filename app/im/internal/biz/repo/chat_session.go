package repo

import (
	"common/proto/gen/common"
	"context"
	"im/internal/biz/model"
)

type ChatSessionRepo interface {
	Save(ctx context.Context, chatSession *model.ChatSession) (*model.ChatSession, error)

	UpdateLastReadMessage(ctx context.Context, chatSessionId int64, messageId int64, operationReadCount int32, updatedBy int64) (*model.ChatSession, error)
	UpdateMuted(ctx context.Context, chatSessionId int64, muted bool, updatedBy int64) (*model.ChatSession, error)
	UpdatePinned(ctx context.Context, chatSessionId int64, pinned bool, updatedBy int64) (*model.ChatSession, error)

	Get(ctx context.Context, req *ChatSessionGetReq) (*model.ChatSession, error)
	List(ctx context.Context, req *ChatSessionGetReq) ([]*model.ChatSession, error)
	Map(ctx context.Context, req *ChatSessionGetReq) (map[int64]*model.ChatSession, error)
	Count(ctx context.Context, req *ChatSessionGetReq) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *ChatSessionGetReq) ([]*model.ChatSession, *common.PageReply, error)
}

type ChatSessionGetReq struct {
	IDs        []int64
	CreatedBy  *int64
	GroupID    *int64
	ReceiverID *int64
}
