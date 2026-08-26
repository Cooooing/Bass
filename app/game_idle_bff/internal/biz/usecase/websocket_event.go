package usecase

import (
	commonenum "common/pkg/enum"
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/enum"
)

// WebSocketEventReq 是服务事件推送处理上下文。
type WebSocketEventReq struct {
	CharacterID int64
	Event       *model.WebSocketEvent
}

// WebSocketEventResult 是服务事件转换后的 WS 消息。
type WebSocketEventResult struct {
	Type              enum.WebSocketMessageType
	Payload           any
	TargetCharacterID int64
	TargetSessionID   string
	Broadcast         bool
	Close             bool
	SilentClose       bool
}

// WebSocketEventHandler 处理一种 NATS 事件。
type WebSocketEventHandler interface {
	Type() commonenum.EventType
	Handle(ctx context.Context, req *WebSocketEventReq) (*WebSocketEventResult, error)
}

// WebSocketEventHandlers 按 NATS 事件类型索引 WS 事件处理器。
type WebSocketEventHandlers map[commonenum.EventType]WebSocketEventHandler
