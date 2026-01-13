package model

import (
	"im/internal/data/ent/gen"
)

type ChatGroupMember struct {
	*gen.ChatGroupMember
}

// ConvertToRpc 转换为RPC返回格式
func (a *ChatGroupMember) ConvertToRpc() {
}
