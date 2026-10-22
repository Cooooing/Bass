package service

import (
	"context"
	"time"

	commonmodel "common/pkg/model"
	imv1 "common/proto/gen/im/v1"

	"im/internal/biz/model"
	"im/internal/biz/usecase"
	"im/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GlobalChatService struct {
	imv1.UnimplementedIMGlobalChatServiceServer
	globalChat *usecase.GlobalChatUsecase
}

func NewGlobalChatService(globalChat *usecase.GlobalChatUsecase) *GlobalChatService {
	return &GlobalChatService{globalChat: globalChat}
}

func (s *GlobalChatService) RegisterGrpc(server *grpc.Server) {
	imv1.RegisterIMGlobalChatServiceServer(server, s)
}

func (*GlobalChatService) RegisterHTTP(*http.Server) {}

func (s *GlobalChatService) Send(ctx context.Context, req *imv1.SendGlobalChatMessage_Req) (*imv1.SendGlobalChatMessage_Resp, error) {
	message, err := s.globalChat.Send(ctx, req.GetSenderId(), req.GetClientMessageId(), req.GetContent())
	if err != nil {
		return nil, err
	}
	return &imv1.SendGlobalChatMessage_Resp{Message: s.globalChatMessageProto(message)}, nil
}

func (s *GlobalChatService) Revoke(ctx context.Context, req *imv1.RevokeGlobalChatMessage_Req) (*imv1.RevokeGlobalChatMessage_Resp, error) {
	message, err := s.globalChat.Revoke(ctx, req.GetUserId(), req.GetMessageId(), req.GetAdministrator())
	if err != nil {
		return nil, err
	}
	return &imv1.RevokeGlobalChatMessage_Resp{Message: s.globalChatMessageProto(message)}, nil
}

func (s *GlobalChatService) List(ctx context.Context, req *imv1.ListGlobalChatMessages_Req) (*imv1.ListGlobalChatMessages_Resp, error) {
	rows, page, err := s.globalChat.List(ctx, req.GetBeforeId(), commonmodel.PageReqFromProto(req.GetPage()))
	if err != nil {
		return nil, err
	}
	response := make([]*imv1.GlobalChatMessage, 0, len(rows))
	for _, row := range rows {
		response = append(response, s.globalChatMessageProto(row))
	}
	return &imv1.ListGlobalChatMessages_Resp{Page: page.ToProto(), Rows: response}, nil
}

func (s *GlobalChatService) Mute(ctx context.Context, req *imv1.MuteGlobalChatUser_Req) (*imv1.MuteGlobalChatUser_Resp, error) {
	if err := s.globalChat.Mute(ctx, req.GetOperatorUserId(), req.GetUserId(), time.Duration(req.GetMutedSeconds())*time.Second, req.Reason); err != nil {
		return nil, err
	}
	return &imv1.MuteGlobalChatUser_Resp{}, nil
}

func (s *GlobalChatService) Unmute(ctx context.Context, req *imv1.UnmuteGlobalChatUser_Req) (*imv1.UnmuteGlobalChatUser_Resp, error) {
	if err := s.globalChat.Unmute(ctx, req.GetUserId()); err != nil {
		return nil, err
	}
	return &imv1.UnmuteGlobalChatUser_Resp{}, nil
}

func (*GlobalChatService) globalChatMessageProto(row *model.GlobalChatMessage) *imv1.GlobalChatMessage {
	if row == nil {
		return nil
	}
	messageType, ok := enum.MessageTypeMap.ToProto(row.Type)
	if !ok {
		return nil
	}
	response := &imv1.GlobalChatMessage{Id: row.ID, SenderId: row.SenderID, Content: row.Content, MessageType: messageType, Revoked: row.Revoked}
	if row.RevokedBy != nil {
		response.RevokedBy = *row.RevokedBy
	}
	if row.CreatedAt != nil {
		response.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.RevokedAt != nil {
		response.RevokedAt = timestamppb.New(*row.RevokedAt)
	}
	return response
}
