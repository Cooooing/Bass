package repo

import (
	cv1 "common/api/common/v1"
	"context"
	"im/internal/biz/model"
	"im/internal/data/ent/gen"
	"time"
)

type ChatGroupMemberRepo interface {
	Save(ctx context.Context, tx *gen.Client, chatGroupMember *model.ChatGroupMember) (*model.ChatGroupMember, error)

	//UpdateRole(ctx context.Context, tx *gen.Client, chatGroupMemberId int64, role v1.ChatGroupMemberRole) (*model.ChatGroupMember, error)
	UpdateMuteEndAt(ctx context.Context, tx *gen.Client, groupId int64, UserId int64, muteEndAt time.Duration) (*model.ChatGroupMember, error)

	GetOne(ctx context.Context, tx *gen.Client, req *ChatGroupMemberGetReq) (*model.ChatGroupMember, error)
	GetList(ctx context.Context, tx *gen.Client, req *ChatGroupMemberGetReq) ([]*model.ChatGroupMember, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *ChatGroupMemberGetReq) ([]*model.ChatGroupMember, *cv1.PageReply, error)
}

type ChatGroupMemberGetReq struct {
}
