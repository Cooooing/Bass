package repo

import (
	"common/api/gen/common"
	v1 "common/api/gen/im/v1"
	"context"
	"im/internal/biz/model"
)

type ChatGroupRepo interface {
	Save(ctx context.Context, chatGroup *model.ChatGroup) (*model.ChatGroup, error)

	UpdateAvatar(ctx context.Context, chatGroupId int64, avatar string) (*model.ChatGroup, error)
	UpdateOwner(ctx context.Context, chatGroupId int64, ownerId int64) (*model.ChatGroup, error)
	UpdateLastMessage(ctx context.Context, message *model.ChatMessage) (*model.ChatGroup, error)
	UpdateMemberCount(ctx context.Context, chatGroupId int64, operationMemberCount int32) (*model.ChatGroup, error)
	UpdateStatus(ctx context.Context, chatGroupId int64, status v1.ChatGroupStatus) (*model.ChatGroup, error)

	Get(ctx context.Context, req *ChatGroupGetReq) (*model.ChatGroup, error)
	GetList(ctx context.Context, req *ChatGroupGetReq) ([]*model.ChatGroup, error)
	GetPage(ctx context.Context, page *common.PageRequest, req *ChatGroupGetReq) ([]*model.ChatGroup, *common.PageReply, error)
}

type ChatGroupGetReq struct {
}
