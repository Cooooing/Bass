package command

import (
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
)

type ConfigGetHandler struct {
	configUsecase *usecase.ConfigUsecase
}

func NewConfigGetHandler(configUsecase *usecase.ConfigUsecase) *ConfigGetHandler {
	return &ConfigGetHandler{
		configUsecase: configUsecase,
	}
}

func (h *ConfigGetHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeConfigGet
}

func (h *ConfigGetHandler) Validate(command *v1.WebSocketCommand) bool {
	return true
}

func (h *ConfigGetHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	row, err := h.configUsecase.Get(ctx)
	if err != nil {
		return err
	}
	req.Connection.Send(ctx, &usecase.WebSocketSendMessage{
		Type:    enum.WebSocketMessageTypeConfigCompleted,
		Payload: row,
	})
	return nil
}
