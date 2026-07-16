package repo

import (
	"context"
	"im/internal/biz/base"
	"im/internal/biz/model"
	"im/internal/enum"
)

type ChatMessageRepo interface {
	Save(ctx context.Context, req *ChatMessageSaveReq) (*ChatMessageSaveResponse, error)
	UpdateStatus(ctx context.Context, req *ChatMessageUpdateStatusReq) (*ChatMessageUpdateStatusResponse, error)

	Get(ctx context.Context, req *ChatMessageGetReq) (*ChatMessageGetResponse, error)
	List(ctx context.Context, req *ChatMessageListReq) (*ChatMessageListResponse, error)
	Map(ctx context.Context, req *ChatMessageMapReq) (*ChatMessageMapResponse, error)
	Count(ctx context.Context, req *ChatMessageCountReq) (*ChatMessageCountResponse, error)
	Page(ctx context.Context, req *ChatMessagePageReq) (*ChatMessagePageResponse, error)
}

type ChatMessageSaveReq struct {
	ChatMessage *model.ChatMessage
}

type ChatMessageSaveResponse struct {
	ChatMessage *model.ChatMessage
}

type ChatMessageUpdateStatusReq struct {
	ChatMessageID int64
	Status        enum.MessageStatus
	UpdatedBy     int64
}

type ChatMessageUpdateStatusResponse struct {
	ChatMessage *model.ChatMessage
}

type ChatMessageQuery struct {
	Page      *base.PageRequest
	IDs       []int64
	SessionID *int64
	SenderID  *int64
}

type ChatMessageGetReq struct {
	ChatMessageQuery
}

type ChatMessageGetResponse struct {
	ChatMessage *model.ChatMessage
}

type ChatMessageListReq struct {
	ChatMessageQuery
}

type ChatMessageListResponse struct {
	Rows []*model.ChatMessage
}

type ChatMessageMapReq struct {
	ChatMessageQuery
}

type ChatMessageMapResponse struct {
	Rows map[int64]*model.ChatMessage
}

type ChatMessageCountReq struct {
	ChatMessageQuery
}

type ChatMessageCountResponse struct {
	Count int
}

type ChatMessagePageReq struct {
	ChatMessageQuery
}

type ChatMessagePageResponse struct {
	Rows []*model.ChatMessage
	Page *base.PageResponse
}
