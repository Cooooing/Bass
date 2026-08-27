package usecase

import (
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"game_idle_bff/internal/enum"
)

// WebSocketCommandReq 是客户端上行命令上下文。
type WebSocketCommandReq struct {
	CharacterID int64
	SessionID   string
	Connection  *WebSocketConnection
	Command     *v1.WebSocketCommand
}

// WebSocketCommandHandler 处理一种客户端上行命令。
type WebSocketCommandHandler interface {
	Type() enum.WebSocketMessageType
	Validate(command *v1.WebSocketCommand) bool
	Handle(ctx context.Context, req *WebSocketCommandReq) error
}

// WebSocketCommandHandlers 按 WS 消息类型索引命令处理器。
type WebSocketCommandHandlers map[enum.WebSocketMessageType]WebSocketCommandHandler

// WebSocketCommandError 是客户端命令失败提示。
type WebSocketCommandError struct {
	Message string `json:"message"`
}
