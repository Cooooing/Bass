package repo

import (
	"common/api/gen/common"
	"context"
	"im/internal/biz/model"
	"time"
)

type ChatGroupMemberRepo interface {
	Save(ctx context.Context, chatGroupMember *model.ChatGroupMember) (*model.ChatGroupMember, error)

	UpdateMuteEndAt(ctx context.Context, groupId int64, UserId int64, muteEndAt time.Duration) (*model.ChatGroupMember, error)

	Get(ctx context.Context, req *ChatGroupMemberGetReq) (*model.ChatGroupMember, error)
	GetList(ctx context.Context, req *ChatGroupMemberGetReq) ([]*model.ChatGroupMember, error)
	Page(ctx context.Context, page *common.PageRequest, req *ChatGroupMemberGetReq) ([]*model.ChatGroupMember, *common.PageReply, error)
}

type ChatGroupMemberGetReq struct {
}
