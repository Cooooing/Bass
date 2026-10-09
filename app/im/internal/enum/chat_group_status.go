package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/im/v1/enum"
)

type ChatGroupStatus string

const (
	ChatGroupStatusNormal   ChatGroupStatus = "normal"   // normal
	ChatGroupStatusDissolve ChatGroupStatus = "dissolve" // dissolve
)

var ChatGroupStatusMap = enum.NewMapping[ChatGroupStatus, v1.ChatGroupStatus](map[ChatGroupStatus]enum.Entry[ChatGroupStatus, v1.ChatGroupStatus]{
	ChatGroupStatusNormal:   {Proto: v1.ChatGroupStatus_CHAT_GROUP_STATUS_NORMAL},
	ChatGroupStatusDissolve: {Proto: v1.ChatGroupStatus_CHAT_GROUP_STATUS_DISSOLVE},
})

func (e ChatGroupStatus) String() string { return string(e) }
