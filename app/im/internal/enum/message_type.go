package enum

import (
	commonenum "common/pkg/enum"
	commonenums "common/proto/gen/common/enums"
)

// MessageType 表示 IM 消息的渲染语义。
type MessageType string

const (
	MessageTypeText      MessageType = "text"       // 普通 Markdown 文本消息。
	MessageTypeRedPacket MessageType = "red_packet" // 红包消息。
)

var MessageTypeMap = commonenum.NewMapping[MessageType, commonenums.IMMessageType](map[MessageType]commonenum.Entry[MessageType, commonenums.IMMessageType]{
	MessageTypeText:      {Proto: commonenums.IMMessageType_IM_MESSAGE_TYPE_TEXT},
	MessageTypeRedPacket: {Proto: commonenums.IMMessageType_IM_MESSAGE_TYPE_RED_PACKET},
})
