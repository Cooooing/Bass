package repo

import (
	"context"
	"im/internal/biz/base"
	"im/internal/biz/model"
)

type ChatSessionRepo interface {
	Save(ctx context.Context, req *ChatSessionSaveReq) (*ChatSessionSaveResponse, error)

	UpdateLastReadMessage(ctx context.Context, req *ChatSessionUpdateLastReadMessageReq) (*ChatSessionUpdateLastReadMessageResponse, error)
	UpdateMuted(ctx context.Context, req *ChatSessionUpdateMutedReq) (*ChatSessionUpdateMutedResponse, error)
	UpdatePinned(ctx context.Context, req *ChatSessionUpdatePinnedReq) (*ChatSessionUpdatePinnedResponse, error)

	Get(ctx context.Context, req *ChatSessionGetReq) (*ChatSessionGetResponse, error)
	List(ctx context.Context, req *ChatSessionListReq) (*ChatSessionListResponse, error)
	Map(ctx context.Context, req *ChatSessionMapReq) (*ChatSessionMapResponse, error)
	Count(ctx context.Context, req *ChatSessionCountReq) (*ChatSessionCountResponse, error)
	Page(ctx context.Context, req *ChatSessionPageReq) (*ChatSessionPageResponse, error)
}

type ChatSessionSaveReq struct {
	ChatSession *model.ChatSession
}

type ChatSessionSaveResponse struct {
	ChatSession *model.ChatSession
}

type ChatSessionUpdateLastReadMessageReq struct {
	ChatSessionID      int64
	MessageID          int64
	OperationReadCount int32
	UpdatedBy          int64
}

type ChatSessionUpdateLastReadMessageResponse struct {
	ChatSession *model.ChatSession
}

type ChatSessionUpdateMutedReq struct {
	ChatSessionID int64
	Muted         bool
	UpdatedBy     int64
}

type ChatSessionUpdateMutedResponse struct {
	ChatSession *model.ChatSession
}

type ChatSessionUpdatePinnedReq struct {
	ChatSessionID int64
	Pinned        bool
	UpdatedBy     int64
}

type ChatSessionUpdatePinnedResponse struct {
	ChatSession *model.ChatSession
}

type ChatSessionQuery struct {
	Page       *base.PageRequest
	IDs        []int64
	CreatedBy  *int64
	GroupID    *int64
	ReceiverID *int64
}

type ChatSessionGetReq struct {
	ChatSessionQuery
}

type ChatSessionGetResponse struct {
	ChatSession *model.ChatSession
}

type ChatSessionListReq struct {
	ChatSessionQuery
}

type ChatSessionListResponse struct {
	Rows []*model.ChatSession
}

type ChatSessionMapReq struct {
	ChatSessionQuery
}

type ChatSessionMapResponse struct {
	Rows map[int64]*model.ChatSession
}

type ChatSessionCountReq struct {
	ChatSessionQuery
}

type ChatSessionCountResponse struct {
	Count int
}

type ChatSessionPageReq struct {
	ChatSessionQuery
}

type ChatSessionPageResponse struct {
	Rows []*model.ChatSession
	Page *base.PageResponse
}
