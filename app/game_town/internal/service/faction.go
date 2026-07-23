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

type FactionService struct {
	v1.UnimplementedGameTownFactionServiceServer
	usecase *usecase.FactionUsecase
}

func NewFactionService(usecase *usecase.FactionUsecase) *FactionService {
	return &FactionService{usecase: usecase}
}

func (s *FactionService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterGameTownFactionServiceServer(server, s)
}

func (s *FactionService) RegisterHttp(*http.Server) {}

func (s *FactionService) Get(ctx context.Context, req *v1.GetGameTownFaction_Request) (*v1.GetGameTownFaction_Resp, error) {
	if req.GetWorldId() <= 0 || req.GetPlayerId() <= 0 || req.GetId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	row, err := s.usecase.Get(ctx, req.GetWorldId(), req.GetPlayerId(), req.GetId())
	if err != nil {
		return nil, err
	}

	return &v1.GetGameTownFaction_Resp{
		Row: &v1.GetGameTownFaction_Resp_Row{
			Id:          row.ID,
			Code:        row.Code,
			Name:        row.Name,
			Description: row.Description,
			PublicGoal:  row.PublicGoal,
			Status:      enum.FactionStatusMap.MustToProto(row.Status),
		},
	}, nil
}

func (s *FactionService) List(ctx context.Context, req *v1.ListGameTownFactions_Request) (*v1.ListGameTownFactions_Resp, error) {
	if req.GetWorldId() <= 0 || req.GetPlayerId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}

	rows, err := s.usecase.List(ctx, req.GetWorldId(), req.GetPlayerId())
	if err != nil {
		return nil, err
	}

	reply := &v1.ListGameTownFactions_Resp{
		Rows: make([]*v1.ListGameTownFactions_Resp_Row, 0, len(rows)),
	}
	for _, row := range rows {
		reply.Rows = append(reply.Rows, &v1.ListGameTownFactions_Resp_Row{
			Id:     row.ID,
			Code:   row.Code,
			Name:   row.Name,
			Status: enum.FactionStatusMap.MustToProto(row.Status),
		})
	}
	return reply, nil
}
