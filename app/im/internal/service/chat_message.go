package service

import (
	v1 "common/api/gen/im/v1"
	"context"
	"im/internal/biz/domain"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type ChatMessageService struct {
	v1.UnimplementedIMChatMessageServiceServer
	chatMessageDomain *domain.ChatMessageDomain
}

func NewChatMessageService(chatMessageDomain *domain.ChatMessageDomain) *ChatMessageService {
	return &ChatMessageService{
		chatMessageDomain: chatMessageDomain,
	}
}

func (s *ChatMessageService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterIMChatMessageServiceServer(gs, s)
}

func (s *ChatMessageService) RegisterHttp(hs *http.Server) {
	v1.RegisterIMChatMessageServiceHTTPServer(hs, s)
}

func (s *ChatMessageService) Send(ctx context.Context, req *v1.SendChatMessage_Request) (rsp *v1.SendChatMessage_Reply, err error) {
	// TODO implement me
	panic("implement me")
}

func (s *ChatMessageService) Revoke(ctx context.Context, req *v1.RevokeChatMessage_Request) (rsp *v1.RevokeChatMessage_Reply, err error) {
	// TODO implement me
	panic("implement me")
}

func (s *ChatMessageService) Page(ctx context.Context, req *v1.PageChatMessage_Request) (rsp *v1.PageChatMessage_Reply, err error) {
	// TODO implement me
	panic("implement me")
}
