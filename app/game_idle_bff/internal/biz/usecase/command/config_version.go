package command

import (
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"

	"google.golang.org/protobuf/proto"
)

type ConfigVersionHandler struct {
	configUsecase *usecase.ConfigUsecase
}

func NewConfigVersionHandler(configUsecase *usecase.ConfigUsecase) *ConfigVersionHandler {
	return &ConfigVersionHandler{
		configUsecase: configUsecase,
	}
}

func (h *ConfigVersionHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeConfigVersion
}

func (h *ConfigVersionHandler) Payload() proto.Message {
	return nil
}

func (h *ConfigVersionHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	row, err := h.configUsecase.Version(ctx)
	if err != nil {
		return err
	}
	req.Connection.Send(ctx, &usecase.WebSocketSendMessage{
		Type:    enum.WebSocketMessageTypeConfigVersionCompleted,
		Payload: row,
	})
	return nil
}
