package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_idle/v1/enum"
)

// ChatMessageStatus 表示挂机游戏聊天消息状态。
type ChatMessageStatus string

const (
	ChatMessageStatusNormal  ChatMessageStatus = "normal"  // 正常
	ChatMessageStatusDeleted ChatMessageStatus = "deleted" // 已删除
)

var ChatMessageStatusMap = commonenum.NewMapping[ChatMessageStatus, v1.GameIdleChatMessageStatus](
	map[ChatMessageStatus]commonenum.Entry[ChatMessageStatus, v1.GameIdleChatMessageStatus]{
		ChatMessageStatusNormal: {
			Proto: v1.GameIdleChatMessageStatus_GAME_IDLE_CHAT_MESSAGE_STATUS_NORMAL,
		},
		ChatMessageStatusDeleted: {
			Proto: v1.GameIdleChatMessageStatus_GAME_IDLE_CHAT_MESSAGE_STATUS_DELETED,
		},
	},
)

func (e ChatMessageStatus) String() string {
	return string(e)
}

func ChatMessageStatusValues() []string {
	return ChatMessageStatusMap.EnumValues()
}
