package service

import (
	v1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type MetaRecipeService struct {
	v1.UnimplementedRecipeServiceServer
	recipeUsecase *usecase.MetaRecipeUsecase
}

func NewMetaRecipeService(recipeUsecase *usecase.MetaRecipeUsecase) *MetaRecipeService {
	return &MetaRecipeService{
		recipeUsecase: recipeUsecase,
	}
}

func (s *MetaRecipeService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterRecipeServiceServer(server, s)
}

func (s *MetaRecipeService) RegisterHttp(*http.Server) {
}

func (s *MetaRecipeService) Refresh(ctx context.Context, req *v1.RefreshRecipes_Request) (*v1.RefreshRecipes_Resp, error) {
	if err := s.recipeUsecase.Refresh(ctx); err != nil {
		return nil, err
	}
	return &v1.RefreshRecipes_Resp{}, nil
}
