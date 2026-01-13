package repo

import (
	cv1 "common/api/common/v1"
	"context"
	"im/internal/biz/model"
	"im/internal/data/ent/gen"
)

type ChatGroupRepo interface {
	Save(ctx context.Context, tx *gen.Client, chatGroup *model.ChatGroup) (*model.ChatGroup, error)

	UpdateAvatar(ctx context.Context, tx *gen.Client, chatGroupId int64, avatar string) error
	UpdateOwner(ctx context.Context, tx *gen.Client, chatGroupId int64, ownerId int64) error
	UpdateLastMessage(ctx context.Context, tx *gen.Client, message *model.ChatMessage) error
	UpdateMemberCount(ctx context.Context, tx *gen.Client, chatGroupId int64, memberCount int32) error
	//UpdateStatus(ctx context.Context, tx *gen.Client, chatGroupId int64, status v1.ChatGroupStatus) error

	GetOne(ctx context.Context, tx *gen.Client, req *ChatGroupGetReq) (*model.ChatGroup, error)
	GetList(ctx context.Context, tx *gen.Client, req *ChatGroupGetReq) ([]*model.ChatGroup, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *ChatGroupGetReq) ([]*model.ChatGroup, *cv1.PageReply, error)
}

type ChatGroupGetReq struct {
}
