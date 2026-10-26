package service

import (
	"context"

	imv1 "common/proto/gen/im/v1"

	"im/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type TCCGlobalChatBarrageService struct {
	imv1.UnimplementedTCCGlobalChatBarrageServiceServer
	barrages *usecase.TCCGlobalChatBarrageUsecase
}

func NewTCCGlobalChatBarrageService(barrages *usecase.TCCGlobalChatBarrageUsecase) *TCCGlobalChatBarrageService {
	return &TCCGlobalChatBarrageService{barrages: barrages}
}

func (s *TCCGlobalChatBarrageService) RegisterGrpc(server *grpc.Server) {
	imv1.RegisterTCCGlobalChatBarrageServiceServer(server, s)
}

func (*TCCGlobalChatBarrageService) RegisterHTTP(*http.Server) {}

func (s *TCCGlobalChatBarrageService) Try(ctx context.Context, req *imv1.TCCSendGlobalChatBarrage_Req) (*imv1.TCCSendGlobalChatBarrage_Resp, error) {
	if err := s.barrages.Try(ctx, req.GetBarrageId(), req.GetUserId(), req.GetContent(), req.GetPointsCost()); err != nil {
		return nil, err
	}
	return &imv1.TCCSendGlobalChatBarrage_Resp{}, nil
}

func (s *TCCGlobalChatBarrageService) Confirm(ctx context.Context, req *imv1.TCCSendGlobalChatBarrage_Req) (*imv1.TCCSendGlobalChatBarrage_Resp, error) {
	if err := s.barrages.Confirm(ctx, req.GetBarrageId()); err != nil {
		return nil, err
	}
	return &imv1.TCCSendGlobalChatBarrage_Resp{}, nil
}

func (s *TCCGlobalChatBarrageService) Cancel(ctx context.Context, req *imv1.TCCSendGlobalChatBarrage_Req) (*imv1.TCCSendGlobalChatBarrage_Resp, error) {
	if err := s.barrages.Cancel(ctx, req.GetBarrageId()); err != nil {
		return nil, err
	}
	return &imv1.TCCSendGlobalChatBarrage_Resp{}, nil
}
