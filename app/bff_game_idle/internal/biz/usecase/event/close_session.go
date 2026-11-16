package event

import (
	event "common/proto/gen/common/enums/event"
	"context"

	"bff_game_idle/internal/biz/usecase"
	"bff_game_idle/internal/enum"
)

type CloseSessionHandler struct {
}

func NewCloseSessionHandler() *CloseSessionHandler {
	return &CloseSessionHandler{}
}

func (h *CloseSessionHandler) Type() event.EventSubject {
	return event.EventSubject_EVENT_SUBJECT_GAME_IDLE_CLOSE_SESSION
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
