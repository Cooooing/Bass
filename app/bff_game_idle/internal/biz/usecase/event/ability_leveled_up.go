package event

import (
	commonenum "common/pkg/enum"
	"context"

	"bff_game_idle/internal/biz/usecase"
	"bff_game_idle/internal/enum"
)

type AbilityLeveledUpHandler struct {
}

func NewAbilityLeveledUpHandler() *AbilityLeveledUpHandler {
	return &AbilityLeveledUpHandler{}
}

func (h *AbilityLeveledUpHandler) Type() commonenum.EventType {
	return commonenum.EventTypeGameIdleAbilityLeveledUp
}

func (h *AbilityLeveledUpHandler) Handle(ctx context.Context, req *usecase.WebSocketEventReq) (*usecase.WebSocketEventResult, error) {
	return &usecase.WebSocketEventResult{
		Type:              enum.WebSocketMessageTypeAbilityLeveledUp,
		Payload:           req.Event.AbilityLeveledUp,
		TargetCharacterID: req.Event.AbilityLeveledUp.CharacterID,
	}, nil
}
