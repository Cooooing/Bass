package service

import (
	v1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type RegionService struct {
	v1.UnimplementedRegionServiceServer
	regionUsecase *usecase.RegionUsecase
}

func NewRegionService(regionUsecase *usecase.RegionUsecase) *RegionService {
	return &RegionService{
		regionUsecase: regionUsecase,
	}
}

func (s *RegionService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterRegionServiceServer(server, s)
}

func (s *RegionService) RegisterHttp(*http.Server) {
}

func (s *RegionService) List(ctx context.Context, req *v1.ListRegions_Request) (*v1.ListRegions_Resp, error) {
	rows, err := s.regionUsecase.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]*v1.Region, 0, len(rows))
	for _, row := range rows {
		out = append(out, &v1.Region{
			RegionId:    row.ID,
			Name:        row.Name,
			Description: row.Description,
			ActionKind:  row.ActionKind.String(),
			Enabled:     row.Enabled,
			Sort:        row.Sort,
		})
	}
	return &v1.ListRegions_Resp{Rows: out}, nil
}
