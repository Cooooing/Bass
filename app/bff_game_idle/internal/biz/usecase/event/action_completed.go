package event

import (
	event "common/proto/gen/common/enums/event"
	"context"

	"bff_game_idle/internal/biz/usecase"
	"bff_game_idle/internal/enum"
)

type ActionCompletedHandler struct {
}

func NewActionCompletedHandler() *ActionCompletedHandler {
	return &ActionCompletedHandler{}
}

func (h *ActionCompletedHandler) Type() event.EventSubject {
	return event.EventSubject_EVENT_SUBJECT_GAME_IDLE_ACTION_COMPLETED
}

func (h *ActionCompletedHandler) Handle(ctx context.Context, req *usecase.WebSocketEventReq) (*usecase.WebSocketEventResult, error) {
	return &usecase.WebSocketEventResult{
		Type:              enum.WebSocketMessageTypeActionCompleted,
		Payload:           req.Event.ActionCompleted,
		TargetCharacterID: req.Event.ActionCompleted.CharacterID,
	}, nil
}
