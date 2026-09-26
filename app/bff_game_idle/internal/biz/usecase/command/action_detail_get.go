package command

import (
	"bff_game_idle/internal/biz/usecase"
	"bff_game_idle/internal/enum"
	v1 "common/proto/gen/bff_game_idle/v1"
	"context"

	"google.golang.org/protobuf/proto"
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

func (h *ActionDetailGetHandler) Payload() proto.Message {
	return &v1.ActionDetailGetWebSocketPayload{}
}

func (h *ActionDetailGetHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload := req.Payload.(*v1.ActionDetailGetWebSocketPayload)
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
