package command

import (
	"bff_game_idle/internal/biz/usecase"
	"bff_game_idle/internal/enum"
	"context"

	"google.golang.org/protobuf/proto"
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

func (h *ConfigGetHandler) Payload() proto.Message {
	return nil
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
