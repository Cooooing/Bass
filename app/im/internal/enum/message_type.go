package enum

import (
	commonenum "common/pkg/enum"
	imv1enum "common/proto/gen/im/v1/enum"
)

// MessageType 表示 IM 消息的渲染语义。
type MessageType string

const (
	MessageTypeText      MessageType = "text"       // 普通 Markdown 文本消息。
	MessageTypeRedPacket MessageType = "red_packet" // 红包消息。
)

var MessageTypeMap = commonenum.NewMapping[MessageType, imv1enum.IMMessageType](map[MessageType]commonenum.Entry[MessageType, imv1enum.IMMessageType]{
	MessageTypeText:      {Proto: imv1enum.IMMessageType_IM_MESSAGE_TYPE_TEXT},
	MessageTypeRedPacket: {Proto: imv1enum.IMMessageType_IM_MESSAGE_TYPE_RED_PACKET},
})

func (e MessageType) String() string { return string(e) }
