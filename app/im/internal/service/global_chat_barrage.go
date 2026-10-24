package service

import (
	"context"

	commonmodel "common/pkg/model"
	imv1 "common/proto/gen/im/v1"

	"im/internal/biz/model"
	"im/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GlobalChatBarrageService struct {
	imv1.UnimplementedIMGlobalChatBarrageServiceServer
	barrages *usecase.GlobalChatBarrageUsecase
}

func NewGlobalChatBarrageService(barrages *usecase.GlobalChatBarrageUsecase) *GlobalChatBarrageService {
	return &GlobalChatBarrageService{barrages: barrages}
}

func (s *GlobalChatBarrageService) RegisterGrpc(server *grpc.Server) {
	imv1.RegisterIMGlobalChatBarrageServiceServer(server, s)
}

func (*GlobalChatBarrageService) RegisterHTTP(*http.Server) {}

func (s *GlobalChatBarrageService) Send(ctx context.Context, req *imv1.SendGlobalChatBarrage_Req) (*imv1.SendGlobalChatBarrage_Resp, error) {
	barrage, err := s.barrages.Send(ctx, req.GetUserId(), req.GetClientRequestId(), req.GetContent())
	if err != nil {
		return nil, err
	}
	return &imv1.SendGlobalChatBarrage_Resp{Barrage: s.barrageProto(barrage)}, nil
}

func (s *GlobalChatBarrageService) List(ctx context.Context, req *imv1.ListGlobalChatBarrages_Req) (*imv1.ListGlobalChatBarrages_Resp, error) {
	rows, page, err := s.barrages.List(ctx, req.UserId, commonmodel.PageReqFromProto(req.GetPage()))
	if err != nil {
		return nil, err
	}
	response := make([]*imv1.GlobalChatBarrage, 0, len(rows))
	for _, row := range rows {
		response = append(response, s.barrageProto(row))
	}
	return &imv1.ListGlobalChatBarrages_Resp{Page: page.ToProto(), Rows: response}, nil
}

func (*GlobalChatBarrageService) barrageProto(row *model.GlobalChatBarrage) *imv1.GlobalChatBarrage {
	if row == nil {
		return nil
	}
	response := &imv1.GlobalChatBarrage{Id: row.ID, UserId: row.UserID, Content: row.Content, PointsCost: row.PointsCost}
	if row.CreatedAt != nil {
		response.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	return response
}
