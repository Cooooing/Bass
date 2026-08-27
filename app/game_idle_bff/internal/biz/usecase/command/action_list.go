package command

import (
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
)

type ActionListHandler struct {
	actionQueueUsecase *usecase.ActionQueueUsecase
}

func NewActionListHandler(actionQueueUsecase *usecase.ActionQueueUsecase) *ActionListHandler {
	return &ActionListHandler{
		actionQueueUsecase: actionQueueUsecase,
	}
}

func (h *ActionListHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeActionList
}

func (h *ActionListHandler) Validate(command *v1.WebSocketCommand) bool {
	return true
}

func (h *ActionListHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	queue, err := h.actionQueueUsecase.List(ctx, req.CharacterID)
	if err != nil {
		return err
	}
	req.Connection.Send(ctx, &usecase.WebSocketSendMessage{
		Type:    enum.WebSocketMessageTypeActionListed,
		Payload: queue,
	})
	return nil
}
