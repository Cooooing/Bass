package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/im/v1/enum"
)

type MessageType string

const (
	MessageTypeNormal MessageType = "normal" // normal
)

var MessageTypeMap = enum.NewMapping[MessageType, v1.MessageType](map[MessageType]enum.Entry[MessageType, v1.MessageType]{
	MessageTypeNormal: {Proto: v1.MessageType_MESSAGE_TYPE_NORMAL},
})

func (e MessageType) String() string { return string(e) }
