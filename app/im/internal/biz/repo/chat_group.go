package repo

import (
	"common/proto/gen/common"
	"context"
	"im/internal/biz/model"
	"im/internal/enum"
)

type ChatGroupRepo interface {
	Save(ctx context.Context, chatGroup *model.ChatGroup) (*model.ChatGroup, error)

	UpdateAvatar(ctx context.Context, chatGroupId int64, avatar string, updatedBy int64) (*model.ChatGroup, error)
	UpdateOwner(ctx context.Context, chatGroupId int64, ownerId int64, updatedBy int64) (*model.ChatGroup, error)
	UpdateLastMessage(ctx context.Context, message *model.ChatMessage, updatedBy int64) (*model.ChatGroup, error)
	UpdateMemberCount(ctx context.Context, chatGroupId int64, operationMemberCount int32, updatedBy int64) (*model.ChatGroup, error)
	UpdateStatus(ctx context.Context, chatGroupId int64, status enum.ChatGroupStatus, updatedBy int64) (*model.ChatGroup, error)

	Get(ctx context.Context, req *ChatGroupGetReq) (*model.ChatGroup, error)
	List(ctx context.Context, req *ChatGroupGetReq) ([]*model.ChatGroup, error)
	Map(ctx context.Context, req *ChatGroupGetReq) (map[int64]*model.ChatGroup, error)
	Count(ctx context.Context, req *ChatGroupGetReq) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *ChatGroupGetReq) ([]*model.ChatGroup, *common.PageReply, error)
}

type ChatGroupGetReq struct {
	IDs    []int64
	Status *enum.ChatGroupStatus
}
