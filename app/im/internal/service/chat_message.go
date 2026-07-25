package service

import (
	"common/proto/gen/common"
	"context"
	"im/internal/biz/base"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/im/v1"
	"im/internal/biz/usecase"
	"im/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type ChatMessageService struct {
	v1.UnimplementedIMChatMessageServiceServer
	chatMessageUsecase *usecase.ChatMessageUsecase
}

func NewChatMessageService(
	chatMessageUsecase *usecase.ChatMessageUsecase,
) *ChatMessageService {
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
func (s *ChatMessageService) Send(ctx context.Context, req *v1.SendChatMessage_Req) (*v1.SendChatMessage_Resp, error) {
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
	err := s.chatMessageUsecase.Send(ctx, &usecase.SendReq{
		SenderID:     req.GetUserId(),
		ReceiverID:   req.GetReceiverId(),
		ReceiverType: receiverType,
		MessageType:  messageType,
		Content:      req.GetContent(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.SendChatMessage_Resp{}, nil
}

// Revoke 撤回消息。
func (s *ChatMessageService) Revoke(ctx context.Context, req *v1.RevokeChatMessage_Req) (*v1.RevokeChatMessage_Resp, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.chatMessageUsecase.Revoke(ctx, &usecase.RevokeReq{
		MessageID: req.GetId(),
		SenderID:  req.GetUserId(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.RevokeChatMessage_Resp{}, nil
}

// List 查询消息列表。
func (s *ChatMessageService) List(ctx context.Context, req *v1.ListChatMessages_Req) (*v1.ListChatMessages_Resp, error) {
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
	resp, err := s.chatMessageUsecase.List(ctx, &usecase.ChatMessageListReq{
		Page: &base.PageRequest{
			Page: int64(req.GetPage().GetPage()),
			Size: int64(req.GetPage().GetSize()),
		},
		IDs:       ids,
		SessionID: sessionID,
		SenderID:  senderID,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.ListChatMessages_Resp_ChatMessage, 0, len(resp.List))
	for _, item := range resp.List {
		msgType := enum.MessageTypeMap.MustToProto(item.Type)
		status := enum.MessageStatusMap.MustToProto(item.Status)
		var receiverUserID int64
		if item.ReceiverID != nil {
			receiverUserID = *item.ReceiverID
		}
		var receiverGroupID int64
		if item.GroupID != nil {
			receiverGroupID = *item.GroupID
		}
		rows = append(rows, &v1.ListChatMessages_Resp_ChatMessage{
			Id:              item.ID,
			ReceiverUserId:  receiverUserID,
			ReceiverGroupId: receiverGroupID,
			Type:            msgType,
			Content:         item.Content,
			Status:          status,
		})
	}
	return &v1.ListChatMessages_Resp{
		Page: &common.PageResp{
			Page:  uint32(resp.Page.Page),
			Size:  uint32(resp.Page.Size),
			Total: uint32(resp.Page.Total),
		},
		Rows: rows,
	}, nil
}
