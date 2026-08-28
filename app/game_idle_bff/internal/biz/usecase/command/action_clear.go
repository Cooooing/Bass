package command

import (
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"

	"google.golang.org/protobuf/proto"
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

func (h *ActionClearHandler) Payload() proto.Message {
	return nil
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
