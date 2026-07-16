package repo

import (
	"context"
	"im/internal/biz/base"
	"im/internal/biz/model"
	"im/internal/enum"
)

type ChatGroupRepo interface {
	Save(ctx context.Context, req *ChatGroupSaveReq) (*ChatGroupSaveResponse, error)

	UpdateAvatar(ctx context.Context, req *ChatGroupUpdateAvatarReq) (*ChatGroupUpdateAvatarResponse, error)
	UpdateOwner(ctx context.Context, req *ChatGroupUpdateOwnerReq) (*ChatGroupUpdateOwnerResponse, error)
	UpdateLastMessage(ctx context.Context, req *ChatGroupUpdateLastMessageReq) (*ChatGroupUpdateLastMessageResponse, error)
	UpdateMemberCount(ctx context.Context, req *ChatGroupUpdateMemberCountReq) (*ChatGroupUpdateMemberCountResponse, error)
	UpdateStatus(ctx context.Context, req *ChatGroupUpdateStatusReq) (*ChatGroupUpdateStatusResponse, error)

	Get(ctx context.Context, req *ChatGroupGetReq) (*ChatGroupGetResponse, error)
	List(ctx context.Context, req *ChatGroupListReq) (*ChatGroupListResponse, error)
	Map(ctx context.Context, req *ChatGroupMapReq) (*ChatGroupMapResponse, error)
	Count(ctx context.Context, req *ChatGroupCountReq) (*ChatGroupCountResponse, error)
	Page(ctx context.Context, req *ChatGroupPageReq) (*ChatGroupPageResponse, error)
}

type ChatGroupSaveReq struct {
	ChatGroup *model.ChatGroup
}

type ChatGroupSaveResponse struct {
	ChatGroup *model.ChatGroup
}

type ChatGroupUpdateAvatarReq struct {
	ChatGroupID int64
	Avatar      string
	UpdatedBy   int64
}

type ChatGroupUpdateAvatarResponse struct {
	ChatGroup *model.ChatGroup
}

type ChatGroupUpdateOwnerReq struct {
	ChatGroupID int64
	OwnerID     int64
	UpdatedBy   int64
}

type ChatGroupUpdateOwnerResponse struct {
	ChatGroup *model.ChatGroup
}

type ChatGroupUpdateLastMessageReq struct {
	Message   *model.ChatMessage
	UpdatedBy int64
}

type ChatGroupUpdateLastMessageResponse struct {
	ChatGroup *model.ChatGroup
}

type ChatGroupUpdateMemberCountReq struct {
	ChatGroupID          int64
	OperationMemberCount int32
	UpdatedBy            int64
}

type ChatGroupUpdateMemberCountResponse struct {
	ChatGroup *model.ChatGroup
}

type ChatGroupUpdateStatusReq struct {
	ChatGroupID int64
	Status      enum.ChatGroupStatus
	UpdatedBy   int64
}

type ChatGroupUpdateStatusResponse struct {
	ChatGroup *model.ChatGroup
}

type ChatGroupQuery struct {
	Page   *base.PageRequest
	IDs    []int64
	Status *enum.ChatGroupStatus
}

type ChatGroupGetReq struct {
	ChatGroupQuery
}

type ChatGroupGetResponse struct {
	ChatGroup *model.ChatGroup
}

type ChatGroupListReq struct {
	ChatGroupQuery
}

type ChatGroupListResponse struct {
	Rows []*model.ChatGroup
}

type ChatGroupMapReq struct {
	ChatGroupQuery
}

type ChatGroupMapResponse struct {
	Rows map[int64]*model.ChatGroup
}

type ChatGroupCountReq struct {
	ChatGroupQuery
}

type ChatGroupCountResponse struct {
	Count int
}

type ChatGroupPageReq struct {
	ChatGroupQuery
}

type ChatGroupPageResponse struct {
	Rows []*model.ChatGroup
	Page *base.PageResponse
}
