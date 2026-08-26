package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle/internal/biz/usecase"
	"sort"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type BackpackService struct {
	v1.UnimplementedGameIdleBackpackServiceServer
	backpackUsecase *usecase.BackpackUsecase
}

func NewBackpackService(backpackUsecase *usecase.BackpackUsecase) *BackpackService {
	return &BackpackService{
		backpackUsecase: backpackUsecase,
	}
}

func (s *BackpackService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterGameIdleBackpackServiceServer(server, s)
}

func (s *BackpackService) RegisterHttp(*http.Server) {
}

func (s *BackpackService) Get(ctx context.Context, req *v1.GetGameIdleBackpack_Request) (*v1.GetGameIdleBackpack_Resp, error) {
	if req.GetCharacterId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	rows, err := s.backpackUsecase.Map(ctx, &usecase.BackpackMapReq{
		CharacterID: req.GetCharacterId(),
		ItemIDs:     req.GetItemIds(),
	})
	if err != nil {
		return nil, err
	}
	itemIDs := make([]string, 0, len(rows))
	for itemID := range rows {
		itemIDs = append(itemIDs, itemID)
	}
	sort.Strings(itemIDs)
	items := make([]*v1.CharacterItem, 0, len(itemIDs))
	for _, itemID := range itemIDs {
		row := rows[itemID]
		items = append(items, &v1.CharacterItem{
			ItemId:   row.ItemID,
			Quantity: row.Quantity,
		})
	}
	return &v1.GetGameIdleBackpack_Resp{Rows: items}, nil
}
