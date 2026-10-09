package command

import (
	"common/pkg/apperror"
	v1 "common/proto/gen/bff_game_idle/v1"
	cerrors "common/proto/gen/common/errors"
	"context"

	"bff_game_idle/internal/biz/usecase"
	"bff_game_idle/internal/enum"

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
	payload, ok := req.Payload.(*v1.ChatMessageSendWebSocketPayload)
	if !ok {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	_, err := h.chatUsecase.Send(ctx, &usecase.SendChatMessageReq{
		CharacterID:         req.CharacterID,
		ChannelType:         payload.GetChannelType(),
		ChannelID:           payload.GetChannelId(),
		ReceiverCharacterID: payload.GetReceiverCharacterId(),
		Content:             payload.GetContent(),
	})
	return err
}
