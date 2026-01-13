package model

import (
	"im/internal/data/ent/gen"
)

type ChatMessage struct {
	*gen.ChatMessage
}

// ConvertToRpc 转换为RPC返回格式
func (a *ChatMessage) ConvertToRpc() {
}
