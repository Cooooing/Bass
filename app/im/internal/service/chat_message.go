package service

import (
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/im/v1"
	"common/pkg/apperror"
	"context"
	"im/internal/biz/usecase"
	"im/internal/enum"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ChatMessageService struct {
	v1.UnimplementedIMChatMessageServiceServer
	chatMessageUsecase *usecase.ChatMessageUsecase
}

func NewChatMessageService(chatMessageUsecase *usecase.ChatMessageUsecase) *ChatMessageService {
	return &ChatMessageService{
		chatMessageUsecase: chatMessageUsecase,
	}
}

func (s *ChatMessageService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterIMChatMessageServiceServer(gs, s)
}

func (s *ChatMessageService) RegisterHttp(hs *http.Server) {
}

// Send 发送消息。
func (s *ChatMessageService) Send(ctx context.Context, req *v1.SendChatMessage_Request) (*v1.SendChatMessage_Reply, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	receiverType, ok := enum.ReceiverTypeMap.ToEnum(req.GetReceiverType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	messageType, ok := enum.MessageTypeMap.ToEnum(req.GetType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.chatMessageUsecase.Send(ctx, req.GetUserId(), req.GetReceiverId(), receiverType, messageType, req.GetContent())
	if err != nil {
		return nil, err
	}
	return &v1.SendChatMessage_Reply{}, nil
}

// Revoke 撤回消息。
func (s *ChatMessageService) Revoke(ctx context.Context, req *v1.RevokeChatMessage_Request) (*v1.RevokeChatMessage_Reply, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.chatMessageUsecase.Revoke(ctx, req.GetId(), req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &v1.RevokeChatMessage_Reply{}, nil
}

// List 查询消息列表。
func (s *ChatMessageService) List(ctx context.Context, req *v1.ListChatMessages_Request) (*v1.ListChatMessages_Reply, error) {
	var ids []int64
	var sessionID *int64
	var senderID *int64
	if req.GetQuery() != nil {
		ids = req.GetQuery().GetIds()
		if req.GetQuery().GetSessionId() != 0 {
			sessionID = new(req.GetQuery().SessionId)
		}
		if req.GetQuery().GetSenderId() != 0 {
			senderID = new(req.GetQuery().SenderId)
		}
	}
	list, page, err := s.chatMessageUsecase.List(ctx, req.GetPage(), ids, sessionID, senderID)
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.ChatMessage, 0, len(list))
	for _, item := range list {
		rows = append(rows, toProtoChatMessage(item))
	}
	return &v1.ListChatMessages_Reply{
		Page: page,
		Rows: rows,
	}, nil
}
