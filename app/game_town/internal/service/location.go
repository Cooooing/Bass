package service

import (
	"context"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/game_town/v1"
	"game_town/internal/biz/usecase"
	"game_town/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type LocationService struct {
	v1.UnimplementedGameTownLocationServiceServer
	usecase *usecase.LocationUsecase
}

func NewLocationService(usecase *usecase.LocationUsecase) *LocationService {
	return &LocationService{usecase: usecase}
}

func (s *LocationService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterGameTownLocationServiceServer(server, s)
}

func (s *LocationService) RegisterHttp(*http.Server) {}

func (s *LocationService) Get(ctx context.Context, req *v1.GetGameTownLocation_Request) (*v1.GetGameTownLocation_Resp, error) {
	if req.GetWorldId() <= 0 || req.GetPlayerId() <= 0 || req.GetId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	row, err := s.usecase.Get(ctx, req.GetWorldId(), req.GetPlayerId(), req.GetId())
	if err != nil {
		return nil, err
	}

	return &v1.GetGameTownLocation_Resp{
		Row: &v1.GetGameTownLocation_Resp_Row{
			Id:                   row.ID,
			Code:                 row.Code,
			Name:                 row.Name,
			Description:          row.Description,
			Status:               enum.LocationStatusMap.MustToProto(row.Status),
			ControllingFactionId: row.ControllingFactionID,
			EnvironmentTags:      row.EnvironmentTags,
		},
	}, nil
}

func (s *LocationService) List(ctx context.Context, req *v1.ListGameTownLocations_Request) (*v1.ListGameTownLocations_Resp, error) {
	if req.GetWorldId() <= 0 || req.GetPlayerId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	rows, member, err := s.usecase.List(ctx, req.GetWorldId(), req.GetPlayerId())
	if err != nil {
		return nil, err
	}

	reply := &v1.ListGameTownLocations_Resp{
		Rows: make([]*v1.ListGameTownLocations_Resp_Row, 0, len(rows)),
	}
	for _, row := range rows {
		reply.Rows = append(reply.Rows, &v1.ListGameTownLocations_Resp_Row{
			Id:         row.ID,
			Code:       row.Code,
			Name:       row.Name,
			Status:     enum.LocationStatusMap.MustToProto(row.Status),
			Current:    row.ID == member.CurrentLocationID,
			Accessible: row.Accessible,
		})
	}
	return reply, nil
}
