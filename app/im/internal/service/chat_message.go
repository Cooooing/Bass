package service

import (
	v1 "common/api/im/v1"
	"context"
	"im/internal/biz/domain"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ChatMessageService struct {
	v1.UnimplementedIMChatMessageServiceServer
	*BaseService
	chatMessageDomain *domain.ChatMessageDomain
}

func NewChatMessageService(baseService *BaseService, chatMessageDomain *domain.ChatMessageDomain) *ChatMessageService {
	return &ChatMessageService{
		BaseService:       baseService,
		chatMessageDomain: chatMessageDomain,
	}
}

func (s *ChatMessageService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterIMChatMessageServiceServer(gs, s)
}

func (s *ChatMessageService) RegisterHttp(hs *http.Server) {
	v1.RegisterIMChatMessageServiceHTTPServer(hs, s)
}

func (s *ChatMessageService) Send(ctx context.Context, req *v1.ChatMessageSendRequest) (rsp *v1.ChatMessageSendReply, err error) {
	// TODO implement me
	panic("implement me")
}

func (s *ChatMessageService) Revoke(ctx context.Context, req *v1.ChatMessageRevokeRequest) (rsp *v1.ChatMessageRevokeReply, err error) {
	// TODO implement me
	panic("implement me")
}

func (s *ChatMessageService) Page(ctx context.Context, req *v1.ChatMessagePageRequest) (rsp *v1.ChatMessagePageReply, err error) {
	// TODO implement me
	panic("implement me")
}
