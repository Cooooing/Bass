package command

import (
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
)

type ActionClearHandler struct {
	actionQueueUsecase *usecase.ActionQueueUsecase
}

func NewActionClearHandler(actionQueueUsecase *usecase.ActionQueueUsecase) *ActionClearHandler {
	return &ActionClearHandler{
		actionQueueUsecase: actionQueueUsecase,
	}
}

func (h *ActionClearHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeActionClear
}

func (h *ActionClearHandler) Validate(command *v1.WebSocketCommand) bool {
	return true
}

func (h *ActionClearHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	if err := h.actionQueueUsecase.Clear(ctx, req.CharacterID); err != nil {
		return err
	}
	queue, err := h.actionQueueUsecase.List(ctx, req.CharacterID)
	if err != nil {
		return err
	}
	req.Connection.Send(ctx, &usecase.WebSocketSendMessage{
		Type:    enum.WebSocketMessageTypeActionQueueUpdated,
		Payload: queue,
	})
	return nil
}
