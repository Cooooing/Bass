package repo

import (
	"context"
	"im/internal/biz/base"
	"im/internal/biz/model"
	"time"
)

type ChatGroupMemberRepo interface {
	Save(ctx context.Context, chatGroupMember *model.ChatGroupMember) (*model.ChatGroupMember, error)

	UpdateMuteEndAt(ctx context.Context, req *ChatGroupMemberUpdateMuteEndAtReq) (*model.ChatGroupMember, error)

	Get(ctx context.Context, query *ChatGroupMemberQuery) (*model.ChatGroupMember, error)
	List(ctx context.Context, query *ChatGroupMemberQuery) ([]*model.ChatGroupMember, error)
	Map(ctx context.Context, query *ChatGroupMemberQuery) (map[int64]*model.ChatGroupMember, error)
	Count(ctx context.Context, query *ChatGroupMemberQuery) (int, error)
	Page(ctx context.Context, query *ChatGroupMemberQuery) (*ChatGroupMemberPageResp, error)
}

type ChatGroupMemberUpdateMuteEndAtReq struct {
	GroupID   int64
	UserID    int64
	MuteEndAt time.Duration
	UpdatedBy int64
}

type ChatGroupMemberQuery struct {
	Page    *base.PageRequest
	IDs     []int64
	GroupID *int64
	UserID  *int64
}

type ChatGroupMemberPageResp struct {
	Rows []*model.ChatGroupMember
	Page *base.PageResp
}
