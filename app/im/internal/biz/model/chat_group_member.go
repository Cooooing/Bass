package model

import (
	v1 "common/api/gen/im/v1"
	userv1 "common/api/gen/user/v1"
	"im/internal/data/gen"

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
		Role:      v1.ChatGroupMemberRole(a.Role),
	}
	if a.MuteEndAt != nil {
		groupMember.MuteEndAt = timestamppb.New(*a.MuteEndAt)
	}
	return groupMember
}
