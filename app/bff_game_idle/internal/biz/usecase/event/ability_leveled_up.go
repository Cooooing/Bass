package event

import (
	event "common/proto/gen/common/enums/event/v1"
	"context"

	"bff_game_idle/internal/biz/usecase"
	"bff_game_idle/internal/enum"
)

type AbilityLeveledUpHandler struct {
}

func NewAbilityLeveledUpHandler() *AbilityLeveledUpHandler {
	return &AbilityLeveledUpHandler{}
}

func (h *AbilityLeveledUpHandler) Type() event.EventSubject {
	return event.EventSubject_EVENT_SUBJECT_GAME_IDLE_ABILITY_LEVELED_UP
}

func (h *AbilityLeveledUpHandler) Handle(ctx context.Context, req *usecase.WebSocketEventReq) (*usecase.WebSocketEventResult, error) {
	return &usecase.WebSocketEventResult{
		Type:              enum.WebSocketMessageTypeAbilityLeveledUp,
		Payload:           req.Event.AbilityLeveledUp,
		TargetCharacterID: req.Event.AbilityLeveledUp.CharacterID,
	}, nil
}
