package command

import (
	v1 "common/proto/gen/game_idle_bff/v1"
	"context"
	"game_idle_bff/internal/biz/usecase"
	"game_idle_bff/internal/enum"
)

type ChatMessageListHandler struct {
	chatUsecase *usecase.ChatUsecase
}

func NewChatMessageListHandler(chatUsecase *usecase.ChatUsecase) *ChatMessageListHandler {
	return &ChatMessageListHandler{
		chatUsecase: chatUsecase,
	}
}

func (h *ChatMessageListHandler) Type() enum.WebSocketMessageType {
	return enum.WebSocketMessageTypeChatMessageList
}

func (h *ChatMessageListHandler) Validate(command *v1.WebSocketCommand) bool {
	return command.GetPayload().GetChatMessageList() != nil
}

func (h *ChatMessageListHandler) Handle(ctx context.Context, req *usecase.WebSocketCommandReq) error {
	payload := req.Command.GetPayload().GetChatMessageList()
	rows, err := h.chatUsecase.List(ctx, &usecase.ListChatMessagesReq{
		ChannelType: payload.GetChannelType(),
		ChannelID:   payload.GetChannelId(),
		BeforeID:    payload.GetBeforeId(),
		Size:        payload.GetSize(),
	})
	if err != nil {
		return err
	}
	req.Connection.Send(ctx, &usecase.WebSocketSendMessage{
		Type:    enum.WebSocketMessageTypeChatMessageListed,
		Payload: rows,
	})
	return nil
}
