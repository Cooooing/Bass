package service

import (
	v1 "common/proto/gen/game_town/v1"
	"context"
	"game_town/internal/biz/usecase"
	gameenum "game_town/internal/enum"
	"time"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PlayerService struct {
	v1.UnimplementedGameTownPlayerServiceServer
	gameUsecase *usecase.GameUsecase
}

func NewPlayerService(gameUsecase *usecase.GameUsecase) *PlayerService {
	return &PlayerService{gameUsecase: gameUsecase}
}
func (s *PlayerService) RegisterGrpc(gs *grpc.Server) { v1.RegisterGameTownPlayerServiceServer(gs, s) }
func (s *PlayerService) RegisterHttp(hs *http.Server) {}
func (s *PlayerService) Register(ctx context.Context, req *v1.RegisterGameTownPlayer_Request) (*v1.RegisterGameTownPlayer_Response, error) {
	registerResponse, err := s.gameUsecase.RegisterPlayer(ctx, &usecase.RegisterPlayerReq{Name: req.GetName(), DisplayName: req.GetDisplayName()})
	if err != nil {
		return nil, err
	}
	timestamp := func(t *time.Time) *timestamppb.Timestamp {
		if t == nil {
			return nil
		}
		return timestamppb.New(*t)
	}
	reply := &v1.RegisterGameTownPlayer_Response{}
	if row := registerResponse.Row; row != nil {
		reply.Row = &v1.RegisterGameTownPlayer_Response_GameTownPlayer{
			CreatedAt:   timestamp(row.CreatedAt),
			UpdatedAt:   timestamp(row.UpdatedAt),
			Id:          row.ID,
			Name:        row.Name,
			DisplayName: row.DisplayName,
			Status:      gameenum.PlayerStatusMap.MustToProto(gameenum.PlayerStatus(row.Status)),
		}
	}
	return reply, nil
}
func (s *PlayerService) Get(ctx context.Context, req *v1.GetGameTownPlayer_Request) (*v1.GetGameTownPlayer_Response, error) {
	getResponse, err := s.gameUsecase.GetPlayer(ctx, &usecase.GetPlayerReq{ID: req.GetId()})
	if err != nil {
		return nil, err
	}
	timestamp := func(t *time.Time) *timestamppb.Timestamp {
		if t == nil {
			return nil
		}
		return timestamppb.New(*t)
	}
	reply := &v1.GetGameTownPlayer_Response{}
	if row := getResponse.Row; row != nil {
		reply.Row = &v1.GetGameTownPlayer_Response_GameTownPlayer{
			CreatedAt:   timestamp(row.CreatedAt),
			UpdatedAt:   timestamp(row.UpdatedAt),
			Id:          row.ID,
			Name:        row.Name,
			DisplayName: row.DisplayName,
			Status:      gameenum.PlayerStatusMap.MustToProto(gameenum.PlayerStatus(row.Status)),
		}
	}
	return reply, nil
}
