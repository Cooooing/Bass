package event

import (
	commonenum "common/pkg/enum"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
)

type CloseSessionHandler struct {
}

func NewCloseSessionHandler() *CloseSessionHandler {
	return &CloseSessionHandler{}
}

func (h *CloseSessionHandler) Type() commonenum.EventType {
	return commonenum.EventTypeGameIdleCloseSession
}

func (h *CloseSessionHandler) Handle(ctx context.Context, req *usecase.WebSocketEventReq) (*usecase.WebSocketEventResult, error) {
	silentClose := req.Event.CloseSession.Reason == "timeout"
	payload := req.Event.CloseSession
	if silentClose {
		payload = nil
	}
	return &usecase.WebSocketEventResult{
		Type:            enum.WebSocketMessageTypeSessionClose,
		Payload:         payload,
		TargetSessionID: req.Event.CloseSession.SessionID,
		Close:           true,
		SilentClose:     silentClose,
	}, nil
}
