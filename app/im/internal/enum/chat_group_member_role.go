package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/im/v1/enum"
)

type ChatGroupMemberRole string

const (
	ChatGroupMemberRoleMember ChatGroupMemberRole = "member" // member
	ChatGroupMemberRoleAdmin  ChatGroupMemberRole = "admin"  // admin
	ChatGroupMemberRoleOwner  ChatGroupMemberRole = "owner"  // owner
)

var ChatGroupMemberRoleMap = enum.NewMapping[ChatGroupMemberRole, v1.ChatGroupMemberRole](map[ChatGroupMemberRole]enum.Entry[ChatGroupMemberRole, v1.ChatGroupMemberRole]{
	ChatGroupMemberRoleMember: {Proto: v1.ChatGroupMemberRole_CHAT_GROUP_MEMBER_ROLE_MEMBER},
	ChatGroupMemberRoleAdmin:  {Proto: v1.ChatGroupMemberRole_CHAT_GROUP_MEMBER_ROLE_ADMIN},
	ChatGroupMemberRoleOwner:  {Proto: v1.ChatGroupMemberRole_CHAT_GROUP_MEMBER_ROLE_OWNER},
})

func (e ChatGroupMemberRole) String() string { return string(e) }
