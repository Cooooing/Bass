package service

import (
	v1 "common/proto/gen/game_town/v1"
	"context"
	"game_town/internal/biz/usecase"
	gameenum "game_town/internal/enum"
	"time"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type NpcService struct {
	v1.UnimplementedGameTownNpcServiceServer
	gameUsecase *usecase.GameUsecase
}

func NewNpcService(gameUsecase *usecase.GameUsecase) *NpcService {
	return &NpcService{gameUsecase: gameUsecase}
}
func (s *NpcService) RegisterGrpc(gs *grpc.Server) { v1.RegisterGameTownNpcServiceServer(gs, s) }
func (s *NpcService) RegisterHttp(hs *http.Server) {}
func (s *NpcService) Get(ctx context.Context, req *v1.GetGameTownNpc_Request) (*v1.GetGameTownNpc_Response, error) {
	getResponse, err := s.gameUsecase.GetNpc(ctx, &usecase.GetNpcReq{ID: req.GetId()})
	if err != nil {
		return nil, err
	}
	timestamp := func(t *time.Time) *timestamppb.Timestamp {
		if t == nil {
			return nil
		}
		return timestamppb.New(*t)
	}
	structValue := func(values map[string]any) *structpb.Struct {
		st, err := structpb.NewStruct(values)
		if err != nil {
			return &structpb.Struct{}
		}
		return st
	}
	reply := &v1.GetGameTownNpc_Response{}
	if row := getResponse.Row; row != nil {
		reply.Row = &v1.GetGameTownNpc_Response_GameTownNpc{
			CreatedAt:         timestamp(row.CreatedAt),
			UpdatedAt:         timestamp(row.UpdatedAt),
			Id:                row.ID,
			WorldId:           row.WorldID,
			Code:              row.Code,
			Name:              row.Name,
			Role:              row.Role,
			Personality:       row.Personality,
			Goal:              row.Goal,
			Background:        row.Background,
			CurrentLocationId: row.CurrentLocationID,
			State:             gameenum.NpcStateMap.MustToProto(gameenum.NpcState(row.State)),
			SystemPrompt:      row.SystemPrompt,
			GeneratedProfile:  structValue(row.GeneratedProfile),
			Enabled:           row.Enabled,
		}
	}
	return reply, nil
}
func (s *NpcService) List(ctx context.Context, req *v1.ListGameTownNpcs_Request) (*v1.ListGameTownNpcs_Response, error) {
	listResponse, err := s.gameUsecase.ListNpcs(ctx, &usecase.ListNpcsReq{WorldID: req.GetWorldId(), LocationID: req.LocationId})
	if err != nil {
		return nil, err
	}
	timestamp := func(t *time.Time) *timestamppb.Timestamp {
		if t == nil {
			return nil
		}
		return timestamppb.New(*t)
	}
	structValue := func(values map[string]any) *structpb.Struct {
		st, err := structpb.NewStruct(values)
		if err != nil {
			return &structpb.Struct{}
		}
		return st
	}
	reply := &v1.ListGameTownNpcs_Response{Rows: make([]*v1.ListGameTownNpcs_Response_GameTownNpc, 0, len(listResponse.Rows))}
	for _, row := range listResponse.Rows {
		if row == nil {
			reply.Rows = append(reply.Rows, nil)
			continue
		}
		reply.Rows = append(reply.Rows, &v1.ListGameTownNpcs_Response_GameTownNpc{
			CreatedAt:         timestamp(row.CreatedAt),
			UpdatedAt:         timestamp(row.UpdatedAt),
			Id:                row.ID,
			WorldId:           row.WorldID,
			Code:              row.Code,
			Name:              row.Name,
			Role:              row.Role,
			Personality:       row.Personality,
			Goal:              row.Goal,
			Background:        row.Background,
			CurrentLocationId: row.CurrentLocationID,
			State:             gameenum.NpcStateMap.MustToProto(gameenum.NpcState(row.State)),
			SystemPrompt:      row.SystemPrompt,
			GeneratedProfile:  structValue(row.GeneratedProfile),
			Enabled:           row.Enabled,
		})
	}
	return reply, nil
}
