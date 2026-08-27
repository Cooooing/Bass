package command

import (
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
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

func (h *ChatMessageSendHandler) Validate(command *v1.WebSocketCommand) bool {
	return command.GetPayload().GetChatMessageSend() != nil
}

func (h *ChatMessageSendHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload := req.Command.GetPayload().GetChatMessageSend()
	_, err := h.chatUsecase.Send(ctx, &usecase.SendChatMessageReq{
		CharacterID:         req.CharacterID,
		ChannelType:         payload.GetChannelType(),
		ChannelID:           payload.GetChannelId(),
		ReceiverCharacterID: payload.GetReceiverCharacterId(),
		Content:             payload.GetContent(),
	})
	return err
}
