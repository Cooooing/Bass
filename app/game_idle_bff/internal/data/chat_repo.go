package data

import (
	"common/pkg/apperror"
	"common/pkg/client/rpc"
	cerrors "common/proto/gen/common/errors"
	gameidlev1 "common/proto/gen/game_idle/v1"
	gameidleenum "common/proto/gen/game_idle/v1/enum"
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"
)

var _ repo.ChatRepo = (*ChatRepo)(nil)

type ChatRepo struct {
	gameIdleClient *rpc.GameIdleClient
}

func NewChatRepo(gameIdleClient *rpc.GameIdleClient) repo.ChatRepo {
	return &ChatRepo{
		gameIdleClient: gameIdleClient,
	}
}

func (r *ChatRepo) Send(ctx context.Context, req *repo.SendChatMessageReq) (*model.WebSocketChatMessage, error) {
	channelType := gameidleenum.ChatChannelType_CHAT_CHANNEL_TYPE_UNSPECIFIED
	switch req.ChannelType {
	case "world":
		channelType = gameidleenum.ChatChannelType_CHAT_CHANNEL_TYPE_WORLD
	default:
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	reply, err := r.gameIdleClient.Chat.Send(ctx, &gameidlev1.SendChatMessage_Request{
		CharacterId:         req.CharacterID,
		ChannelType:         channelType,
		ChannelId:           req.ChannelID,
		ReceiverCharacterId: req.ReceiverCharacterID,
		Content:             req.Content,
	})
	if err != nil {
		return nil, err
	}
	row := reply.GetRow()
	return &model.WebSocketChatMessage{
		MessageID:           row.GetId(),
		ChannelType:         req.ChannelType,
		ChannelID:           row.GetChannelId(),
		SenderCharacterID:   row.GetSenderCharacterId(),
		ReceiverCharacterID: row.GetReceiverCharacterId(),
		Content:             row.GetContent(),
	}, nil
}

func (r *ChatRepo) List(ctx context.Context, req *repo.ListChatMessagesReq) ([]*model.WebSocketChatMessage, error) {
	channelType := gameidleenum.ChatChannelType_CHAT_CHANNEL_TYPE_UNSPECIFIED
	switch req.ChannelType {
	case "world":
		channelType = gameidleenum.ChatChannelType_CHAT_CHANNEL_TYPE_WORLD
	default:
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	reply, err := r.gameIdleClient.Chat.List(ctx, &gameidlev1.ListChatMessages_Request{
		ChannelType: channelType,
		ChannelId:   req.ChannelID,
		BeforeId:    req.BeforeID,
		Size:        req.Size,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*model.WebSocketChatMessage, 0, len(reply.GetRows()))
	for _, row := range reply.GetRows() {
		rows = append(rows, &model.WebSocketChatMessage{
			MessageID:           row.GetId(),
			ChannelType:         req.ChannelType,
			ChannelID:           row.GetChannelId(),
			SenderCharacterID:   row.GetSenderCharacterId(),
			ReceiverCharacterID: row.GetReceiverCharacterId(),
			Content:             row.GetContent(),
		})
	}
	return rows, nil
}
