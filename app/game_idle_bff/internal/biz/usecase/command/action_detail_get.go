package command

import (
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
)

type ActionDetailGetHandler struct {
	configUsecase *usecase.ConfigUsecase
}

func NewActionDetailGetHandler(configUsecase *usecase.ConfigUsecase) *ActionDetailGetHandler {
	return &ActionDetailGetHandler{
		configUsecase: configUsecase,
	}
}

func (h *ActionDetailGetHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeActionDetailGet
}

func (h *ActionDetailGetHandler) Validate(command *v1.WebSocketCommand) bool {
	return command.GetPayload().GetActionDetailGet() != nil
}

func (h *ActionDetailGetHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload := req.Command.GetPayload().GetActionDetailGet()
	row, err := h.configUsecase.GetActionDetail(ctx, payload.GetActionId())
	if err != nil {
		return err
	}
	req.Connection.Send(ctx, &usecase.WebSocketSendMessage{
		Type:    enum.WebSocketMessageTypeActionDetailCompleted,
		Payload: row,
	})
	return nil
}
