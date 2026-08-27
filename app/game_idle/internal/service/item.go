package service

import (
	v1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type ItemService struct {
	v1.UnimplementedItemServiceServer
	itemUsecase *usecase.ItemUsecase
}

func NewItemService(itemUsecase *usecase.ItemUsecase) *ItemService {
	return &ItemService{
		itemUsecase: itemUsecase,
	}
}

func (s *ItemService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterItemServiceServer(server, s)
}

func (s *ItemService) RegisterHttp(*http.Server) {
}

func (s *ItemService) List(ctx context.Context, req *v1.ListItems_Request) (*v1.ListItems_Resp, error) {
	rows, err := s.itemUsecase.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]*v1.Item, 0, len(rows))
	for _, row := range rows {
		out = append(out, &v1.Item{
			ItemId:      row.ID,
			Name:        row.Name,
			ItemType:    row.Type.String(),
			Description: row.Description,
			Enabled:     row.Enabled,
			Sort:        row.Sort,
		})
	}
	return &v1.ListItems_Resp{Rows: out}, nil
}
