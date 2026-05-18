package repo

import (
	"common/api/gen/common"
	"context"
	"im/internal/biz/model"
	"im/internal/data/gen"
	"time"
)

type ChatGroupMemberRepo interface {
	Save(ctx context.Context, tx *gen.Client, chatGroupMember *model.ChatGroupMember) (*model.ChatGroupMember, error)

	// UpdateRole(ctx context.Context, tx *gen.Client, chatGroupMemberId int64, role v1.ChatGroupMemberRole) (*model.ChatGroupMember, error)
	UpdateMuteEndAt(ctx context.Context, tx *gen.Client, groupId int64, UserId int64, muteEndAt time.Duration) (*model.ChatGroupMember, error)

	GetOne(ctx context.Context, tx *gen.Client, req *ChatGroupMemberGetReq) (*model.ChatGroupMember, error)
	GetList(ctx context.Context, tx *gen.Client, req *ChatGroupMemberGetReq) ([]*model.ChatGroupMember, error)
	GetPage(ctx context.Context, tx *gen.Client, page *common.PageRequest, req *ChatGroupMemberGetReq) ([]*model.ChatGroupMember, *common.PageReply, error)
}

type ChatGroupMemberGetReq struct {
}
