package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/im/v1/enum"
)

type ChatGroupStatus string

const (
	ChatGroupStatusNormal   ChatGroupStatus = "normal"
	ChatGroupStatusDissolve ChatGroupStatus = "dissolve"
)

var ChatGroupStatusMap = enum.NewMapping[ChatGroupStatus, v1.ChatGroupStatus](map[ChatGroupStatus]enum.Entry[ChatGroupStatus, v1.ChatGroupStatus]{
	ChatGroupStatusNormal:   {Proto: v1.ChatGroupStatus_CHAT_GROUP_STATUS_NORMAL},
	ChatGroupStatusDissolve: {Proto: v1.ChatGroupStatus_CHAT_GROUP_STATUS_DISSOLVE},
})

type ChatGroupMemberRole string

const (
	ChatGroupMemberRoleMember ChatGroupMemberRole = "member"
	ChatGroupMemberRoleAdmin  ChatGroupMemberRole = "admin"
	ChatGroupMemberRoleOwner  ChatGroupMemberRole = "owner"
)

var ChatGroupMemberRoleMap = enum.NewMapping[ChatGroupMemberRole, v1.ChatGroupMemberRole](map[ChatGroupMemberRole]enum.Entry[ChatGroupMemberRole, v1.ChatGroupMemberRole]{
	ChatGroupMemberRoleMember: {Proto: v1.ChatGroupMemberRole_CHAT_GROUP_MEMBER_ROLE_MEMBER},
	ChatGroupMemberRoleAdmin:  {Proto: v1.ChatGroupMemberRole_CHAT_GROUP_MEMBER_ROLE_ADMIN},
	ChatGroupMemberRoleOwner:  {Proto: v1.ChatGroupMemberRole_CHAT_GROUP_MEMBER_ROLE_OWNER},
})

type MessageType string

const (
	MessageTypeNormal MessageType = "normal"
)

var MessageTypeMap = enum.NewMapping[MessageType, v1.MessageType](map[MessageType]enum.Entry[MessageType, v1.MessageType]{
	MessageTypeNormal: {Proto: v1.MessageType_MESSAGE_TYPE_NORMAL},
})

type MessageStatus string

const (
	MessageStatusNormal MessageStatus = "normal"
	MessageStatusRevoke MessageStatus = "revoke"
)

var MessageStatusMap = enum.NewMapping[MessageStatus, v1.MessageStatus](map[MessageStatus]enum.Entry[MessageStatus, v1.MessageStatus]{
	MessageStatusNormal: {Proto: v1.MessageStatus_MESSAGE_STATUS_NORMAL},
	MessageStatusRevoke: {Proto: v1.MessageStatus_MESSAGE_STATUS_REVOKE},
})

type ReceiverType string

const (
	ReceiverTypeUser  ReceiverType = "user"
	ReceiverTypeGroup ReceiverType = "group"
)

var ReceiverTypeMap = enum.NewMapping[ReceiverType, v1.ReceiverType](map[ReceiverType]enum.Entry[ReceiverType, v1.ReceiverType]{
	ReceiverTypeUser:  {Proto: v1.ReceiverType_RECEIVER_TYPE_USER},
	ReceiverTypeGroup: {Proto: v1.ReceiverType_RECEIVER_TYPE_GROUP},
})

func (e ChatGroupStatus) String() string {
	return string(e)
}

func (e ChatGroupMemberRole) String() string {
	return string(e)
}

func (e MessageType) String() string {
	return string(e)
}

func (e MessageStatus) String() string {
	return string(e)
}

func (e ReceiverType) String() string {
	return string(e)
}
