package repo

import (
	"common/proto/gen/common"
	"context"
	"im/internal/biz/model"
	"time"
)

type ChatGroupMemberRepo interface {
	Save(ctx context.Context, chatGroupMember *model.ChatGroupMember) (*model.ChatGroupMember, error)

	UpdateMuteEndAt(ctx context.Context, groupId int64, userId int64, muteEndAt time.Duration, updatedBy int64) (*model.ChatGroupMember, error)

	Get(ctx context.Context, req *ChatGroupMemberGetReq) (*model.ChatGroupMember, error)
	List(ctx context.Context, req *ChatGroupMemberGetReq) ([]*model.ChatGroupMember, error)
	Map(ctx context.Context, req *ChatGroupMemberGetReq) (map[int64]*model.ChatGroupMember, error)
	Count(ctx context.Context, req *ChatGroupMemberGetReq) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *ChatGroupMemberGetReq) ([]*model.ChatGroupMember, *common.PageReply, error)
}

type ChatGroupMemberGetReq struct {
	IDs     []int64
	GroupID *int64
	UserID  *int64
}
