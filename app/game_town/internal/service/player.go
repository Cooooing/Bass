package service

import (
	"context"

	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/game_town/v1"
	"game_town/internal/biz/usecase"
	gameenum "game_town/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PlayerService struct {
	v1.UnimplementedGameTownPlayerServiceServer
	usecase *usecase.PlayerUsecase
}

func NewPlayerService(usecase *usecase.PlayerUsecase) *PlayerService {
	return &PlayerService{usecase: usecase}
}

func (s *PlayerService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterGameTownPlayerServiceServer(server, s)
}

func (s *PlayerService) RegisterHttp(*http.Server) {}

func (s *PlayerService) Register(ctx context.Context, req *v1.RegisterGameTownPlayer_Request) (*v1.RegisterGameTownPlayer_Resp, error) {
	row, err := s.usecase.Register(ctx, &usecase.RegisterPlayerReq{
		Name:        req.GetName(),
		DisplayName: req.GetDisplayName(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.RegisterGameTownPlayer_Resp{
		Row: &v1.RegisterGameTownPlayer_Resp_Row{
			Id:          row.ID,
			Name:        row.Name,
			DisplayName: row.DisplayName,
			Status:      gameenum.PlayerStatusMap.MustToProto(row.Status),
			CreatedAt:   timestamppb.New(*row.CreatedAt),
			UpdatedAt:   timestamppb.New(*row.UpdatedAt),
		},
	}, nil
}

func (s *PlayerService) Get(ctx context.Context, req *v1.GetGameTownPlayer_Request) (*v1.GetGameTownPlayer_Resp, error) {
	if req.GetId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row, err := s.usecase.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &v1.GetGameTownPlayer_Resp{
		Row: &v1.GetGameTownPlayer_Resp_Row{
			Id:          row.ID,
			Name:        row.Name,
			DisplayName: row.DisplayName,
			Status:      gameenum.PlayerStatusMap.MustToProto(row.Status),
			CreatedAt:   timestamppb.New(*row.CreatedAt),
			UpdatedAt:   timestamppb.New(*row.UpdatedAt),
		},
	}, nil
}
