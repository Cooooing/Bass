package usecase

import (
	"context"
	"game_idle_bff/internal/enum"

	"google.golang.org/protobuf/proto"
)

// WebSocketCommandReq 是客户端上行命令上下文。
type WebSocketCommandReq struct {
	CharacterID int64
	SessionID   string
	Connection  *WebSocketConnection
	Payload     proto.Message
}

// WebSocketCommandHandler 处理一种客户端上行命令。
type WebSocketCommandHandler interface {
	Type() enum.WebSocketMessageType
	Payload() proto.Message
	Handle(ctx context.Context, req *WebSocketCommandReq) error
}

// WebSocketCommandHandlers 按 WS 消息类型索引命令处理器。
type WebSocketCommandHandlers map[enum.WebSocketMessageType]WebSocketCommandHandler

// WebSocketCommandError 是客户端命令失败提示。
type WebSocketCommandError struct {
	Message string `json:"message"`
}
