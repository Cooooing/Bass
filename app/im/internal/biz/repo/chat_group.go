package repo

import (
	cv1 "common/api/gen/common/v1"
	v1 "common/api/gen/im/v1"
	"context"
	"im/internal/biz/model"
	"im/internal/data/ent/gen"
)

type ChatGroupRepo interface {
	Save(ctx context.Context, tx *gen.Client, chatGroup *model.ChatGroup) (*model.ChatGroup, error)

	UpdateAvatar(ctx context.Context, tx *gen.Client, chatGroupId int64, avatar string) (*model.ChatGroup, error)
	UpdateOwner(ctx context.Context, tx *gen.Client, chatGroupId int64, ownerId int64) (*model.ChatGroup, error)
	UpdateLastMessage(ctx context.Context, tx *gen.Client, message *model.ChatMessage) (*model.ChatGroup, error)
	UpdateMemberCount(ctx context.Context, tx *gen.Client, chatGroupId int64, operationMemberCount int32) (*model.ChatGroup, error)
	UpdateStatus(ctx context.Context, tx *gen.Client, chatGroupId int64, status v1.ChatGroupStatus) (*model.ChatGroup, error)

	GetOne(ctx context.Context, tx *gen.Client, req *ChatGroupGetReq) (*model.ChatGroup, error)
	GetList(ctx context.Context, tx *gen.Client, req *ChatGroupGetReq) ([]*model.ChatGroup, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *ChatGroupGetReq) ([]*model.ChatGroup, *cv1.PageReply, error)
}

type ChatGroupGetReq struct {
}
