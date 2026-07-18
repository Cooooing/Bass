package repo

import (
	"context"
	"im/internal/biz/base"
	"im/internal/biz/model"
)

type ChatSessionRepo interface {
	Save(ctx context.Context, chatSession *model.ChatSession) (*model.ChatSession, error)

	UpdateLastReadMessage(ctx context.Context, req *ChatSessionUpdateLastReadMessageReq) (*model.ChatSession, error)
	UpdateMuted(ctx context.Context, req *ChatSessionUpdateMutedReq) (*model.ChatSession, error)
	UpdatePinned(ctx context.Context, req *ChatSessionUpdatePinnedReq) (*model.ChatSession, error)

	Get(ctx context.Context, query *ChatSessionQuery) (*model.ChatSession, error)
	List(ctx context.Context, query *ChatSessionQuery) ([]*model.ChatSession, error)
	Map(ctx context.Context, query *ChatSessionQuery) (map[int64]*model.ChatSession, error)
	Count(ctx context.Context, query *ChatSessionQuery) (int, error)
	Page(ctx context.Context, query *ChatSessionQuery) (*ChatSessionPageResp, error)
}

type ChatSessionUpdateLastReadMessageReq struct {
	ChatSessionID      int64
	MessageID          int64
	OperationReadCount int32
	UpdatedBy          int64
}

type ChatSessionUpdateMutedReq struct {
	ChatSessionID int64
	Muted         bool
	UpdatedBy     int64
}

type ChatSessionUpdatePinnedReq struct {
	ChatSessionID int64
	Pinned        bool
	UpdatedBy     int64
}

type ChatSessionQuery struct {
	Page       *base.PageRequest
	IDs        []int64
	CreatedBy  *int64
	GroupID    *int64
	ReceiverID *int64
}

type ChatSessionPageResp struct {
	Rows []*model.ChatSession
	Page *base.PageResp
}
