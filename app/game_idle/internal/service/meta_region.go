package service

import (
	v1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type MetaRegionService struct {
	v1.UnimplementedRegionServiceServer
	regionUsecase *usecase.MetaRegionUsecase
}

func NewMetaRegionService(regionUsecase *usecase.MetaRegionUsecase) *MetaRegionService {
	return &MetaRegionService{
		regionUsecase: regionUsecase,
	}
}

func (s *MetaRegionService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterRegionServiceServer(server, s)
}

func (s *MetaRegionService) RegisterHttp(*http.Server) {
}

func (s *MetaRegionService) List(ctx context.Context, req *v1.ListRegions_Request) (*v1.ListRegions_Resp, error) {
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

func (s *MetaRegionService) Refresh(ctx context.Context, req *v1.RefreshRegions_Request) (*v1.RefreshRegions_Resp, error) {
	if err := s.regionUsecase.Refresh(ctx); err != nil {
		return nil, err
	}
	return &v1.RefreshRegions_Resp{}, nil
}
