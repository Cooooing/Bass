package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_idle/v1/enum"
)

// ChatChannelType 表示挂机游戏聊天频道类型。
type ChatChannelType string

const (
	ChatChannelTypeWorld ChatChannelType = "world" // 世界频道
)

var ChatChannelTypeMap = commonenum.NewMapping[ChatChannelType, v1.ChatChannelType](
	map[ChatChannelType]commonenum.Entry[ChatChannelType, v1.ChatChannelType]{
		ChatChannelTypeWorld: {
			Proto: v1.ChatChannelType_CHAT_CHANNEL_TYPE_WORLD,
		},
	},
)

func (e ChatChannelType) String() string {
	return string(e)
}

func ChatChannelTypeValues() []string {
	return ChatChannelTypeMap.EnumValues()
}
