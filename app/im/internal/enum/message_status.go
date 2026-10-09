package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/im/v1/enum"
)

type MessageStatus string

const (
	MessageStatusUnknown MessageStatus = "unknown" // unknown
	MessageStatusNormal  MessageStatus = "normal"  // normal
	MessageStatusRevoke  MessageStatus = "revoke"  // revoke
)

var MessageStatusMap = enum.NewMapping[MessageStatus, v1.MessageStatus](map[MessageStatus]enum.Entry[MessageStatus, v1.MessageStatus]{
	MessageStatusUnknown: {Proto: v1.MessageStatus_MESSAGE_STATUS_UNKNOWN},
	MessageStatusNormal:  {Proto: v1.MessageStatus_MESSAGE_STATUS_NORMAL},
	MessageStatusRevoke:  {Proto: v1.MessageStatus_MESSAGE_STATUS_REVOKE},
})

func (e MessageStatus) String() string { return string(e) }
