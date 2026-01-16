package model

import (
	"im/internal/data/ent/gen"
)

type ChatSession struct {
	*gen.ChatSession

	Group *ChatGroup

	UnreadCount uint32
}

// ConvertToRpc 转换为RPC返回格式
func (a *ChatSession) ConvertToRpc() {
}
