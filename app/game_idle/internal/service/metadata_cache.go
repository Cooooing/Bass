package service

import (
	v1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type MetadataCacheService struct {
	v1.UnimplementedMetadataCacheServiceServer
	metadataCacheUsecase *usecase.MetadataCacheUsecase
}

func NewMetadataCacheService(metadataCacheUsecase *usecase.MetadataCacheUsecase) *MetadataCacheService {
	return &MetadataCacheService{
		metadataCacheUsecase: metadataCacheUsecase,
	}
}

func (s *MetadataCacheService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterMetadataCacheServiceServer(server, s)
}

func (s *MetadataCacheService) RegisterHttp(*http.Server) {
}

func (s *MetadataCacheService) Refresh(ctx context.Context, req *v1.RefreshMetadataCache_Request) (*v1.RefreshMetadataCache_Resp, error) {
	if err := s.metadataCacheUsecase.RefreshRedis(ctx); err != nil {
		return nil, err
	}
	return &v1.RefreshMetadataCache_Resp{}, nil
}
