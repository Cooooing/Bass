package repo

import (
	"context"
	"im/internal/biz/base"
	"im/internal/biz/model"
	"im/internal/enum"
)

type ChatGroupRepo interface {
	Save(ctx context.Context, chatGroup *model.ChatGroup) (*model.ChatGroup, error)

	UpdateAvatar(ctx context.Context, req *ChatGroupUpdateAvatarReq) (*model.ChatGroup, error)
	UpdateOwner(ctx context.Context, req *ChatGroupUpdateOwnerReq) (*model.ChatGroup, error)
	UpdateLastMessage(ctx context.Context, req *ChatGroupUpdateLastMessageReq) (*model.ChatGroup, error)
	UpdateMemberCount(ctx context.Context, req *ChatGroupUpdateMemberCountReq) (*model.ChatGroup, error)
	UpdateStatus(ctx context.Context, req *ChatGroupUpdateStatusReq) (*model.ChatGroup, error)

	Get(ctx context.Context, query *ChatGroupQuery) (*model.ChatGroup, error)
	List(ctx context.Context, query *ChatGroupQuery) ([]*model.ChatGroup, error)
	Map(ctx context.Context, query *ChatGroupQuery) (map[int64]*model.ChatGroup, error)
	Count(ctx context.Context, query *ChatGroupQuery) (int, error)
	Page(ctx context.Context, query *ChatGroupQuery) (*ChatGroupPageResp, error)
}

type ChatGroupUpdateAvatarReq struct {
	ChatGroupID int64
	Avatar      string
	UpdatedBy   int64
}

type ChatGroupUpdateOwnerReq struct {
	ChatGroupID int64
	OwnerID     int64
	UpdatedBy   int64
}

type ChatGroupUpdateLastMessageReq struct {
	Message   *model.ChatMessage
	UpdatedBy int64
}

type ChatGroupUpdateMemberCountReq struct {
	ChatGroupID          int64
	OperationMemberCount int32
	UpdatedBy            int64
}

type ChatGroupUpdateStatusReq struct {
	ChatGroupID int64
	Status      enum.ChatGroupStatus
	UpdatedBy   int64
}

type ChatGroupQuery struct {
	Page   *base.PageRequest
	IDs    []int64
	Status *enum.ChatGroupStatus
}

type ChatGroupPageResp struct {
	Rows []*model.ChatGroup
	Page *base.PageResp
}
