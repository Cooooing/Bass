package command

import (
	"bff_game_idle/internal/biz/usecase"
	"bff_game_idle/internal/enum"
	v1 "common/proto/gen/bff_game_idle/v1"
	"context"

	"google.golang.org/protobuf/proto"
)

type ChatMessageSendHandler struct {
	chatUsecase *usecase.ChatUsecase
}

func NewChatMessageSendHandler(chatUsecase *usecase.ChatUsecase) *ChatMessageSendHandler {
	return &ChatMessageSendHandler{
		chatUsecase: chatUsecase,
	}
}

func (h *ChatMessageSendHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeChatMessageSend
}

func (h *ChatMessageSendHandler) Payload() proto.Message {
	return &v1.ChatMessageSendWebSocketPayload{}
}

func (h *ChatMessageSendHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload := req.Payload.(*v1.ChatMessageSendWebSocketPayload)
	_, err := h.chatUsecase.Send(ctx, &usecase.SendChatMessageReq{
		CharacterID:         req.CharacterID,
		ChannelType:         payload.GetChannelType(),
		ChannelID:           payload.GetChannelId(),
		ReceiverCharacterID: payload.GetReceiverCharacterId(),
		Content:             payload.GetContent(),
	})
	return err
}
