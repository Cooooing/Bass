package service

import (
	"context"

	imv1 "common/proto/gen/im/v1"

	"im/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

// TCCGlobalChatTopicService 适配全局聊天室话题 TCC 分支协议。
type TCCGlobalChatTopicService struct {
	imv1.UnimplementedTCCGlobalChatTopicServiceServer
	tccTopic *usecase.TCCGlobalChatTopicUsecase
}

func NewTCCGlobalChatTopicService(tccTopic *usecase.TCCGlobalChatTopicUsecase) *TCCGlobalChatTopicService {
	return &TCCGlobalChatTopicService{tccTopic: tccTopic}
}

func (s *TCCGlobalChatTopicService) RegisterGrpc(server *grpc.Server) {
	imv1.RegisterTCCGlobalChatTopicServiceServer(server, s)
}

func (*TCCGlobalChatTopicService) RegisterHTTP(*http.Server) {}

func (s *TCCGlobalChatTopicService) Try(ctx context.Context, req *imv1.TCCSetGlobalChatTopic_Req) (*imv1.TCCSetGlobalChatTopic_Resp, error) {
	if err := s.tccTopic.Try(ctx, req.GetUserId(), req.GetClientRequestId(), req.GetContent(), req.GetPointsCost()); err != nil {
		return nil, err
	}
	return &imv1.TCCSetGlobalChatTopic_Resp{}, nil
}

func (s *TCCGlobalChatTopicService) Confirm(ctx context.Context, req *imv1.TCCSetGlobalChatTopic_Req) (*imv1.TCCSetGlobalChatTopic_Resp, error) {
	if err := s.tccTopic.Confirm(ctx, req.GetClientRequestId()); err != nil {
		return nil, err
	}
	return &imv1.TCCSetGlobalChatTopic_Resp{}, nil
}

func (s *TCCGlobalChatTopicService) Cancel(ctx context.Context, req *imv1.TCCSetGlobalChatTopic_Req) (*imv1.TCCSetGlobalChatTopic_Resp, error) {
	if err := s.tccTopic.Cancel(ctx, req.GetClientRequestId()); err != nil {
		return nil, err
	}
	return &imv1.TCCSetGlobalChatTopic_Resp{}, nil
}
