package model

import (
	v1 "common/gen/im/v1"
	userv1 "common/gen/user/v1"
	"im/internal/data/ent/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ChatGroup struct {
	*gen.ChatGroup

	WithOwner bool
	Owner     *userv1.User
}

// ConvertToRpc 转换为RPC返回格式
func (a *ChatGroup) ConvertToRpc() *v1.ChatGroup {
	group := &v1.ChatGroup{
		CreatedAt:    timestamppb.New(*a.CreatedAt),
		UpdatedAt:    timestamppb.New(*a.UpdatedAt),
		Id:           a.ID,
		Name:         a.Name,
		Avatar:       a.Avatar,
		Introduction: a.Introduction,
		Status:       v1.ChatGroupStatus(a.Status),
		Owner:        a.Owner,
		MemberCount:  a.MemberCount,
	}
	if a.Edges.LastMessageOfGroup != nil {
		group.LastMessage = (&ChatMessage{ChatMessage: a.Edges.LastMessageOfGroup}).ConvertToRpc()
	}
	return group
}
