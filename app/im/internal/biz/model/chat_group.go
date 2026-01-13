package model

import (
	"im/internal/data/ent/gen"
)

type ChatGroup struct {
	*gen.ChatGroup
}

// ConvertToRpc 转换为RPC返回格式
func (a *ChatGroup) ConvertToRpc() {
}
