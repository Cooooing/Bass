package usecase

import (
	"context"
	"encoding/json"
	"game_idle_bff/internal/enum"
)

// WebSocketCommandReq 是客户端上行命令上下文。
type WebSocketCommandReq struct {
	CharacterID int64
	SessionID   string
	Connection  *WebSocketConnection
	Command     *WebSocketCommand
}

// WebSocketCommandHandler 处理一种客户端上行命令。
type WebSocketCommandHandler interface {
	Type() enum.WebSocketMessageType
	Handle(ctx context.Context, req *WebSocketCommandReq) error
}

// WebSocketCommandHandlers 按 WS 消息类型索引命令处理器。
type WebSocketCommandHandlers map[enum.WebSocketMessageType]WebSocketCommandHandler

// WebSocketCommand 是客户端上行消息。
type WebSocketCommand struct {
	Type    enum.WebSocketMessageType `json:"type"`
	Payload json.RawMessage           `json:"payload,omitempty"`
}

// WebSocketCommandError 是客户端命令失败提示。
type WebSocketCommandError struct {
	Message string `json:"message"`
}
