package event

import (
	event "common/proto/gen/common/enums/event/v1"
	"context"

	"bff_game_idle/internal/biz/usecase"
	"bff_game_idle/internal/enum"
)

type ActionQueueUpdatedHandler struct {
}

func NewActionQueueUpdatedHandler() *ActionQueueUpdatedHandler {
	return &ActionQueueUpdatedHandler{}
}

func (h *ActionQueueUpdatedHandler) Type() event.EventSubject {
	return event.EventSubject_EVENT_SUBJECT_GAME_IDLE_ACTION_QUEUE_UPDATED
}

func (h *ActionQueueUpdatedHandler) Handle(ctx context.Context, req *usecase.WebSocketEventReq) (*usecase.WebSocketEventResult, error) {
	return &usecase.WebSocketEventResult{
		Type:              enum.WebSocketMessageTypeActionQueueUpdated,
		Payload:           req.Event.ActionQueueUpdated,
		TargetCharacterID: req.Event.ActionQueueUpdated.CharacterID,
	}, nil
}
