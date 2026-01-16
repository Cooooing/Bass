package model

import (
	v1 "common/api/im/v1"
	userv1 "common/api/user/v1"
	"im/internal/data/ent/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ChatGroupMember struct {
	*gen.ChatGroupMember

	Member *userv1.User
}

// ConvertToRpc 转换为RPC返回格式
func (a *ChatGroupMember) ConvertToRpc() *v1.ChatGroupMember {
	groupMember := &v1.ChatGroupMember{
		CreatedAt: timestamppb.New(*a.CreatedAt),
		UpdatedAt: timestamppb.New(*a.UpdatedAt),
		Id:        a.ID,
		GroupId:   a.GroupID,
		Member:    a.Member,
		Nickname:  a.Nickname,
		Role:      a.Role,
	}
	if a.MuteEndAt != nil {
		groupMember.MuteEndAt = timestamppb.New(*a.MuteEndAt)
	}
	return groupMember
}
