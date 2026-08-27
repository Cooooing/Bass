package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle/internal/biz/usecase"
	"game_idle/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ChatService struct {
	v1.UnimplementedChatServiceServer
	chatUsecase *usecase.ChatUsecase
}

func NewChatService(chatUsecase *usecase.ChatUsecase) *ChatService {
	return &ChatService{
		chatUsecase: chatUsecase,
	}
}

func (s *ChatService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterChatServiceServer(server, s)
}

func (s *ChatService) RegisterHttp(*http.Server) {
}

func (s *ChatService) Send(ctx context.Context, req *v1.SendChatMessage_Request) (*v1.SendChatMessage_Resp, error) {
	if req.GetCharacterId() <= 0 || req.GetContent() == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	channelType, ok := enum.ChatChannelTypeMap.ToEnum(req.GetChannelType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var receiverCharacterID *int64
	if req.GetReceiverCharacterId() > 0 {
		receiverCharacterID = new(req.GetReceiverCharacterId())
	}
	row, err := s.chatUsecase.Send(ctx, &usecase.SendMessageReq{
		CharacterID:         req.GetCharacterId(),
		ChannelType:         channelType,
		ChannelID:           req.GetChannelId(),
		ReceiverCharacterID: receiverCharacterID,
		Content:             req.GetContent(),
	})
	if err != nil {
		return nil, err
	}
	message := &v1.ChatMessage{
		Id:                  row.ID,
		ChannelType:         enum.ChatChannelTypeMap.MustToProto(row.ChannelType),
		ChannelId:           row.ChannelID,
		SenderCharacterId:   row.SenderCharacterID,
		ReceiverCharacterId: 0,
		Content:             row.Content,
		Status:              enum.ChatMessageStatusMap.MustToProto(row.Status),
	}
	if row.ReceiverCharacterID != nil {
		message.ReceiverCharacterId = *row.ReceiverCharacterID
	}
	if row.CreatedAt != nil {
		message.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	return &v1.SendChatMessage_Resp{Row: message}, nil
}

func (s *ChatService) List(ctx context.Context, req *v1.ListChatMessages_Request) (*v1.ListChatMessages_Resp, error) {
	channelType, ok := enum.ChatChannelTypeMap.ToEnum(req.GetChannelType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	rows, err := s.chatUsecase.List(ctx, &usecase.ListMessagesReq{
		ChannelType: channelType,
		ChannelID:   req.GetChannelId(),
		BeforeID:    req.GetBeforeId(),
		Size:        int(req.GetSize()),
	})
	if err != nil {
		return nil, err
	}
	messages := make([]*v1.ChatMessage, 0, len(rows))
	for _, row := range rows {
		message := &v1.ChatMessage{
			Id:                  row.ID,
			ChannelType:         enum.ChatChannelTypeMap.MustToProto(row.ChannelType),
			ChannelId:           row.ChannelID,
			SenderCharacterId:   row.SenderCharacterID,
			ReceiverCharacterId: 0,
			Content:             row.Content,
			Status:              enum.ChatMessageStatusMap.MustToProto(row.Status),
		}
		if row.ReceiverCharacterID != nil {
			message.ReceiverCharacterId = *row.ReceiverCharacterID
		}
		if row.CreatedAt != nil {
			message.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		messages = append(messages, message)
	}
	return &v1.ListChatMessages_Resp{Rows: messages}, nil
}
