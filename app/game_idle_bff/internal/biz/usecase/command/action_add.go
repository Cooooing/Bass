package command

import (
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
)

type ActionAddHandler struct {
	actionQueueUsecase *usecase.ActionQueueUsecase
}

func NewActionAddHandler(actionQueueUsecase *usecase.ActionQueueUsecase) *ActionAddHandler {
	return &ActionAddHandler{
		actionQueueUsecase: actionQueueUsecase,
	}
}

func (h *ActionAddHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeActionAdd
}

func (h *ActionAddHandler) Validate(command *v1.WebSocketCommand) bool {
	return command.GetPayload().GetActionAdd() != nil
}

func (h *ActionAddHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload := req.Command.GetPayload().GetActionAdd()
	if err := h.actionQueueUsecase.Add(ctx, &usecase.AddActionReq{
		CharacterID: req.CharacterID,
		ActionID:    payload.GetActionId(),
		Times:       payload.GetTimes(),
		Position:    payload.Position,
	}); err != nil {
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
