package service

import (
	"context"

	commonmodel "common/pkg/model"
	imv1 "common/proto/gen/im/v1"

	"im/internal/biz/model"
	"im/internal/biz/usecase"
	"im/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PrivateChatService struct {
	imv1.UnimplementedIMPrivateChatServiceServer
	privateChat *usecase.PrivateChatUsecase
}

func NewPrivateChatService(privateChat *usecase.PrivateChatUsecase) *PrivateChatService {
	return &PrivateChatService{privateChat: privateChat}
}
func (s *PrivateChatService) RegisterGrpc(server *grpc.Server) {
	imv1.RegisterIMPrivateChatServiceServer(server, s)
}
func (*PrivateChatService) RegisterHTTP(*http.Server) {}
func (s *PrivateChatService) Send(ctx context.Context, req *imv1.SendPrivateChatMessage_Req) (*imv1.SendPrivateChatMessage_Resp, error) {
	row, err := s.privateChat.Send(ctx, req.GetSenderId(), req.GetReceiverId(), req.GetClientMessageId(), req.GetContent())
	if err != nil {
		return nil, err
	}
	return &imv1.SendPrivateChatMessage_Resp{Message: s.privateChatMessageProto(row)}, nil
}
func (s *PrivateChatService) Revoke(ctx context.Context, req *imv1.RevokePrivateChatMessage_Req) (*imv1.RevokePrivateChatMessage_Resp, error) {
	row, err := s.privateChat.Revoke(ctx, req.GetUserId(), req.GetMessageId())
	if err != nil {
		return nil, err
	}
	return &imv1.RevokePrivateChatMessage_Resp{Message: s.privateChatMessageProto(row)}, nil
}
func (s *PrivateChatService) ListSessions(ctx context.Context, req *imv1.ListPrivateChatSessions_Req) (*imv1.ListPrivateChatSessions_Resp, error) {
	rows, page, err := s.privateChat.ListSessions(ctx, req.GetUserId(), commonmodel.PageReqFromProto(req.GetPage()))
	if err != nil {
		return nil, err
	}
	out := make([]*imv1.PrivateChatSession, 0, len(rows))
	for _, row := range rows {
		out = append(out, s.privateChatSessionProto(row))
	}
	return &imv1.ListPrivateChatSessions_Resp{Page: page.ToProto(), Rows: out}, nil
}
func (s *PrivateChatService) ListMessages(ctx context.Context, req *imv1.ListPrivateChatMessages_Req) (*imv1.ListPrivateChatMessages_Resp, error) {
	rows, page, err := s.privateChat.ListMessages(ctx, req.GetUserId(), req.GetPeerUserId(), req.GetBeforeId(), commonmodel.PageReqFromProto(req.GetPage()))
	if err != nil {
		return nil, err
	}
	out := make([]*imv1.PrivateChatMessage, 0, len(rows))
	for _, row := range rows {
		out = append(out, s.privateChatMessageProto(row))
	}
	return &imv1.ListPrivateChatMessages_Resp{Page: page.ToProto(), Rows: out}, nil
}

func (s *PrivateChatService) CountUnread(ctx context.Context, req *imv1.CountUnreadPrivateChatMessages_Req) (*imv1.CountUnreadPrivateChatMessages_Resp, error) {
	count, err := s.privateChat.CountUnread(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	return &imv1.CountUnreadPrivateChatMessages_Resp{Count: count}, nil
}
func (s *PrivateChatService) MarkRead(ctx context.Context, req *imv1.MarkPrivateChatSessionRead_Req) (*imv1.MarkPrivateChatSessionRead_Resp, error) {
	return &imv1.MarkPrivateChatSessionRead_Resp{}, s.privateChat.MarkRead(ctx, req.GetUserId(), req.GetPeerUserId(), req.GetMessageId())
}
func (s *PrivateChatService) SetMuted(ctx context.Context, req *imv1.SetPrivateChatSessionMuted_Req) (*imv1.SetPrivateChatSessionMuted_Resp, error) {
	return &imv1.SetPrivateChatSessionMuted_Resp{}, s.privateChat.SetMuted(ctx, req.GetUserId(), req.GetPeerUserId(), req.GetMuted())
}
func (s *PrivateChatService) SetPinned(ctx context.Context, req *imv1.SetPrivateChatSessionPinned_Req) (*imv1.SetPrivateChatSessionPinned_Resp, error) {
	return &imv1.SetPrivateChatSessionPinned_Resp{}, s.privateChat.SetPinned(ctx, req.GetUserId(), req.GetPeerUserId(), req.GetPinned())
}

func (*PrivateChatService) privateChatMessageProto(row *model.PrivateChatMessage) *imv1.PrivateChatMessage {
	if row == nil {
		return nil
	}
	messageType, ok := enum.MessageTypeMap.ToProto(row.Type)
	if !ok {
		return nil
	}
	result := &imv1.PrivateChatMessage{Id: row.ID, SenderId: row.SenderID, ReceiverId: row.ReceiverID, Content: row.Content, MessageType: messageType, Revoked: row.Revoked}
	if row.CreatedAt != nil {
		result.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.RevokedAt != nil {
		result.RevokedAt = timestamppb.New(*row.RevokedAt)
	}
	return result
}
func (s *PrivateChatService) privateChatSessionProto(row *model.PrivateChatSession) *imv1.PrivateChatSession {
	if row == nil {
		return nil
	}
	result := &imv1.PrivateChatSession{Id: row.ID, PeerUserId: row.PeerUserID, UnreadCount: row.UnreadCount, Muted: row.Muted, Pinned: row.Pinned, LastMessage: s.privateChatMessageProto(row.LastMessage)}
	if row.LastReadMessageID != nil {
		result.LastReadMessageId = *row.LastReadMessageID
	}
	if row.LastMessageAt != nil {
		result.LastMessageAt = timestamppb.New(*row.LastMessageAt)
	}
	return result
}
