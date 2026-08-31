package service

import (
	v1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle/internal/biz/usecase"
	"game_idle/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type MetaItemService struct {
	v1.UnimplementedItemServiceServer
	itemUsecase *usecase.MetaItemUsecase
}

func NewMetaItemService(itemUsecase *usecase.MetaItemUsecase) *MetaItemService {
	return &MetaItemService{
		itemUsecase: itemUsecase,
	}
}

func (s *MetaItemService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterItemServiceServer(server, s)
}

func (s *MetaItemService) RegisterHttp(*http.Server) {
}

func (s *MetaItemService) List(ctx context.Context, req *v1.ListItems_Request) (*v1.ListItems_Resp, error) {
	rows, err := s.itemUsecase.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]*v1.Item, 0, len(rows))
	for _, row := range rows {
		if row.ID == enum.ItemIDEmpty.String() {
			continue
		}
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

func (s *MetaItemService) Refresh(ctx context.Context, req *v1.RefreshItems_Request) (*v1.RefreshItems_Resp, error) {
	if err := s.itemUsecase.Refresh(ctx); err != nil {
		return nil, err
	}
	return &v1.RefreshItems_Resp{}, nil
}
