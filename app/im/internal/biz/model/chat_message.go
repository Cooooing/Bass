package model

import (
	v1 "common/api/gen/im/v1"
	userv1 "common/api/gen/user/v1"
	"im/internal/data/ent/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ChatMessage struct {
	*gen.ChatMessage

	SenderUser    *userv1.User
	ReceiverUser  *userv1.User
	ReceiverGroup *v1.ChatGroup
}

// ConvertToRpc 转换为RPC返回格式
func (a *ChatMessage) ConvertToRpc() *v1.ChatMessage {
	message := &v1.ChatMessage{
		CreatedAt:     timestamppb.New(*a.CreatedAt),
		UpdatedAt:     timestamppb.New(*a.UpdatedAt),
		CreatedBy:     a.CreatedBy,
		UpdatedBy:     a.UpdatedBy,
		CreatedByName: a.CreatedByName,
		UpdatedByName: a.UpdatedByName,
		Id:            a.ID,
		Sender:        a.SenderUser,
		Type:          v1.MessageType(a.Type),
		Content:       a.Content,
		Status:        v1.MessageStatus(a.Status),
	}
	if a.ReceiverID != nil {
		message.ReceiverUserId = *a.ReceiverID
		message.ReceiverUser = a.ReceiverUser
	}
	if a.GroupID != nil {
		message.ReceiverGroupId = *a.GroupID
		message.ReceiverGroup = (&ChatGroup{ChatGroup: a.Edges.Group}).ConvertToRpc()
	}
	return message
}
