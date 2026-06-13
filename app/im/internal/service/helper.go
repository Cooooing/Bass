package service

import (
	v1 "common/proto/gen/im/v1"
	"im/internal/biz/model"
	"im/internal/enum"
)

// toProtoChatGroup 将 biz 模型转换为 proto 模型。
func toProtoChatGroup(biz *model.ChatGroup) *v1.ChatGroup {
	if biz == nil {
		return nil
	}
	status, _ := enum.ChatGroupStatusMap.ToProto(biz.Status)
	return &v1.ChatGroup{
		Id:           biz.ID,
		Name:         biz.Name,
		Avatar:       biz.Avatar,
		Introduction: biz.Introduction,
		Status:       status,
		MemberCount:  biz.MemberCount,
	}
}

// toProtoChatMessage 将 biz 模型转换为 proto 模型。
func toProtoChatMessage(biz *model.ChatMessage) *v1.ChatMessage {
	if biz == nil {
		return nil
	}
	msgType, _ := enum.MessageTypeMap.ToProto(biz.Type)
	status, _ := enum.MessageStatusMap.ToProto(biz.Status)
	return &v1.ChatMessage{
		Id:              biz.ID,
		ReceiverUserId:  derefInt64(biz.ReceiverID),
		ReceiverGroupId: derefInt64(biz.GroupID),
		Type:            msgType,
		Content:         biz.Content,
		Status:          status,
	}
}

// toProtoChatSession 将 biz 模型转换为 proto 模型。
func toProtoChatSession(biz *model.ChatSession) *v1.ChatSession {
	if biz == nil {
		return nil
	}
	p := &v1.ChatSession{
		Id:          biz.ID,
		IsMuted:     biz.IsMuted,
		IsPinned:    biz.IsPinned,
		UnreadCount: biz.UnreadCount,
	}
	if biz.ReceiverID != nil {
		p.RelationId = *biz.ReceiverID
	}
	if biz.GroupID != nil {
		p.GroupId = *biz.GroupID
	}
	if biz.LastReadMessageID != nil {
		p.LastReadMessageId = *biz.LastReadMessageID
	}
	return p
}

// toProtoChatGroupMember 将 biz 模型转换为 proto 模型。
func toProtoChatGroupMember(biz *model.ChatGroupMember) *v1.ChatGroupMember {
	if biz == nil {
		return nil
	}
	role, _ := enum.ChatGroupMemberRoleMap.ToProto(biz.Role)
	return &v1.ChatGroupMember{
		Id:       biz.ID,
		GroupId:  biz.GroupID,
		Nickname: biz.Nickname,
		Role:     role,
	}
}

// derefInt64 安全解引用 int64 指针。
func derefInt64(p *int64) int64 {
	if p != nil {
		return *p
	}
	return 0
}
