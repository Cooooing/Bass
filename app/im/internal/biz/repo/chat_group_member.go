package repo

import (
	"context"
	"im/internal/biz/base"
	"im/internal/biz/model"
	"time"
)

type ChatGroupMemberRepo interface {
	Save(ctx context.Context, req *ChatGroupMemberSaveReq) (*ChatGroupMemberSaveResponse, error)

	UpdateMuteEndAt(ctx context.Context, req *ChatGroupMemberUpdateMuteEndAtReq) (*ChatGroupMemberUpdateMuteEndAtResponse, error)

	Get(ctx context.Context, req *ChatGroupMemberGetReq) (*ChatGroupMemberGetResponse, error)
	List(ctx context.Context, req *ChatGroupMemberListReq) (*ChatGroupMemberListResponse, error)
	Map(ctx context.Context, req *ChatGroupMemberMapReq) (*ChatGroupMemberMapResponse, error)
	Count(ctx context.Context, req *ChatGroupMemberCountReq) (*ChatGroupMemberCountResponse, error)
	Page(ctx context.Context, req *ChatGroupMemberPageReq) (*ChatGroupMemberPageResponse, error)
}

type ChatGroupMemberSaveReq struct {
	ChatGroupMember *model.ChatGroupMember
}

type ChatGroupMemberSaveResponse struct {
	ChatGroupMember *model.ChatGroupMember
}

type ChatGroupMemberUpdateMuteEndAtReq struct {
	GroupID   int64
	UserID    int64
	MuteEndAt time.Duration
	UpdatedBy int64
}

type ChatGroupMemberUpdateMuteEndAtResponse struct {
	ChatGroupMember *model.ChatGroupMember
}

type ChatGroupMemberQuery struct {
	Page    *base.PageRequest
	IDs     []int64
	GroupID *int64
	UserID  *int64
}

type ChatGroupMemberGetReq struct {
	ChatGroupMemberQuery
}

type ChatGroupMemberGetResponse struct {
	ChatGroupMember *model.ChatGroupMember
}

type ChatGroupMemberListReq struct {
	ChatGroupMemberQuery
}

type ChatGroupMemberListResponse struct {
	Rows []*model.ChatGroupMember
}

type ChatGroupMemberMapReq struct {
	ChatGroupMemberQuery
}

type ChatGroupMemberMapResponse struct {
	Rows map[int64]*model.ChatGroupMember
}

type ChatGroupMemberCountReq struct {
	ChatGroupMemberQuery
}

type ChatGroupMemberCountResponse struct {
	Count int
}

type ChatGroupMemberPageReq struct {
	ChatGroupMemberQuery
}

type ChatGroupMemberPageResponse struct {
	Rows []*model.ChatGroupMember
	Page *base.PageResponse
}
