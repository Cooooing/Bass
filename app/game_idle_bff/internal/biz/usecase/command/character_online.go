package command

import (
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"

	"google.golang.org/protobuf/proto"
)

type CharacterOnlineHandler struct {
	characterUsecase *usecase.CharacterUsecase
}

func NewCharacterOnlineHandler(characterUsecase *usecase.CharacterUsecase) *CharacterOnlineHandler {
	return &CharacterOnlineHandler{
		characterUsecase: characterUsecase,
	}
}

func (h *CharacterOnlineHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeCharacterOnline
}

func (h *CharacterOnlineHandler) Payload() proto.Message {
	return nil
}

func (h *CharacterOnlineHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	session, err := h.characterUsecase.Online(ctx, &usecase.OnlineCharacterReq{
		CharacterID: req.CharacterID,
	})
	if err != nil {
		return err
	}
	if !req.BindOnline(ctx, req.Connection, session) {
		return h.characterUsecase.Offline(ctx, &usecase.OfflineCharacterReq{
			CharacterID: req.CharacterID,
			SessionID:   session.SessionID,
		})
	}
	req.Connection.Send(ctx, &usecase.WebSocketSendMessage{
		Type: enum.WebSocketMessageTypeCharacterOnlineDone,
		Payload: &model.WebSocketOnline{
			CharacterID:      req.CharacterID,
			ExpiresInSeconds: int64(session.RemainingDuration.Seconds()),
		},
	})
	return nil
}
