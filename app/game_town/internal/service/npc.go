package service

import (
	"context"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/game_town/v1"
	"game_town/internal/biz/model"
	"game_town/internal/biz/usecase"
	"game_town/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NpcService struct {
	v1.UnimplementedGameTownNpcServiceServer
	usecase *usecase.NpcUsecase
}

func NewNpcService(
	usecase *usecase.NpcUsecase,
) *NpcService {
	return &NpcService{
		usecase: usecase,
	}
}

func (s *NpcService) RegisterGrpc(
	server *grpc.Server,
) {
	v1.RegisterGameTownNpcServiceServer(server, s)
}

func (s *NpcService) RegisterHttp(
	*http.Server,
) {
}

func (s *NpcService) Get(
	ctx context.Context,
	req *v1.GetGameTownNpc_Request,
) (*v1.GetGameTownNpc_Resp, error) {
	if req.GetWorldId() <= 0 || req.GetPlayerId() <= 0 || req.GetId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	row, err := s.usecase.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if row.WorldID != req.GetWorldId() {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}

	return &v1.GetGameTownNpc_Resp{
		Row: &v1.GetGameTownNpc_Resp_Row{
			Id:                  row.ID,
			WorldId:             row.WorldID,
			Code:                row.Code,
			Name:                row.Name,
			Role:                row.Role,
			Species:             row.Species,
			LifeStatus:          enum.NpcLifeStatusMap.MustToProto(row.LifeStatus),
			LastKnownLocationId: new(row.CurrentLocationID),
			KnownAt:             timestamppb.New(*row.UpdatedAt),
		},
	}, nil
}

func (s *NpcService) List(
	ctx context.Context,
	req *v1.ListGameTownNpcs_Request,
) (*v1.ListGameTownNpcs_Resp, error) {
	if req.GetWorldId() <= 0 || req.GetPlayerId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	rows, err := s.usecase.List(ctx, &usecase.ListNpcsReq{
		WorldID:    req.GetWorldId(),
		LocationID: req.LocationId,
	})
	if err != nil {
		return nil, err
	}

	return &v1.ListGameTownNpcs_Resp{
		Rows: lo.Map(rows, func(row *model.Npc, _ int) *v1.ListGameTownNpcs_Resp_Row {
			return &v1.ListGameTownNpcs_Resp_Row{
				Id:                  row.ID,
				Code:                row.Code,
				Name:                row.Name,
				Role:                row.Role,
				Species:             row.Species,
				LifeStatus:          enum.NpcLifeStatusMap.MustToProto(row.LifeStatus),
				LastKnownLocationId: new(row.CurrentLocationID),
				StateTags:           row.StateTags,
			}
		}),
	}, nil
}
