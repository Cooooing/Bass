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

type AgentConfigService struct {
	v1.UnimplementedGameTownAgentConfigServiceServer
	gameUsecase *usecase.GameUsecase
}

func NewAgentConfigService(gameUsecase *usecase.GameUsecase) *AgentConfigService {
	return &AgentConfigService{gameUsecase: gameUsecase}
}
func (s *AgentConfigService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterGameTownAgentConfigServiceServer(gs, s)
}
func (s *AgentConfigService) RegisterHttp(hs *http.Server) {}
func (s *AgentConfigService) Create(ctx context.Context, req *v1.CreateGameTownAgentConfig_Request) (*v1.CreateGameTownAgentConfig_Response, error) {
	createResponse, err := s.gameUsecase.CreateAgentConfig(ctx, &usecase.CreateAgentConfigReq{PlayerID: req.GetPlayerId(), Name: req.GetName(), Provider: req.GetProvider(), ModelName: req.GetModel(), BaseURL: req.GetBaseUrl(), APIKey: req.GetApiKey(), TimeoutSeconds: req.GetTimeoutSeconds(), IsDefault: req.GetIsDefault()})
	if err != nil {
		return nil, err
	}
	timestamp := func(t *time.Time) *timestamppb.Timestamp {
		if t == nil {
			return nil
		}
		return timestamppb.New(*t)
	}
	reply := &v1.CreateGameTownAgentConfig_Response{}
	if row := createResponse.Row; row != nil {
		reply.Row = &v1.CreateGameTownAgentConfig_Response_GameTownAgentConfig{
			CreatedAt:      timestamp(row.CreatedAt),
			UpdatedAt:      timestamp(row.UpdatedAt),
			Id:             row.ID,
			PlayerId:       row.PlayerID,
			Name:           row.Name,
			Provider:       row.Provider,
			Model:          row.Model,
			BaseUrl:        row.BaseURL,
			HasApiKey:      row.APIKey != "",
			TimeoutSeconds: row.TimeoutSeconds,
			IsDefault:      row.IsDefault,
			Status:         gameenum.AgentConfigStatusMap.MustToProto(gameenum.AgentConfigStatus(row.Status)),
		}
	}
	return reply, nil
}
func (s *AgentConfigService) Get(ctx context.Context, req *v1.GetGameTownAgentConfig_Request) (*v1.GetGameTownAgentConfig_Response, error) {
	getResponse, err := s.gameUsecase.GetAgentConfig(ctx, &usecase.GetAgentConfigReq{ID: req.GetId(), PlayerID: req.GetPlayerId()})
	if err != nil {
		return nil, err
	}
	timestamp := func(t *time.Time) *timestamppb.Timestamp {
		if t == nil {
			return nil
		}
		return timestamppb.New(*t)
	}
	reply := &v1.GetGameTownAgentConfig_Response{}
	if row := getResponse.Row; row != nil {
		reply.Row = &v1.GetGameTownAgentConfig_Response_GameTownAgentConfig{
			CreatedAt:      timestamp(row.CreatedAt),
			UpdatedAt:      timestamp(row.UpdatedAt),
			Id:             row.ID,
			PlayerId:       row.PlayerID,
			Name:           row.Name,
			Provider:       row.Provider,
			Model:          row.Model,
			BaseUrl:        row.BaseURL,
			HasApiKey:      row.APIKey != "",
			TimeoutSeconds: row.TimeoutSeconds,
			IsDefault:      row.IsDefault,
			Status:         gameenum.AgentConfigStatusMap.MustToProto(gameenum.AgentConfigStatus(row.Status)),
		}
	}
	return reply, nil
}
func (s *AgentConfigService) List(ctx context.Context, req *v1.ListGameTownAgentConfigs_Request) (*v1.ListGameTownAgentConfigs_Response, error) {
	var status *string
	if req.Status != nil && *req.Status != v1.GameTownAgentConfigStatus_GAME_TOWN_AGENT_CONFIG_STATUS_UNSPECIFIED {
		value, ok := gameenum.AgentConfigStatusMap.ToEnum(*req.Status)
		if ok {
			statusValue := string(value)
			status = &statusValue
		}
	}
	listResponse, err := s.gameUsecase.ListAgentConfigs(ctx, &usecase.ListAgentConfigsReq{PlayerID: req.GetPlayerId(), Status: status})
	if err != nil {
		return nil, err
	}
	timestamp := func(t *time.Time) *timestamppb.Timestamp {
		if t == nil {
			return nil
		}
		return timestamppb.New(*t)
	}
	reply := &v1.ListGameTownAgentConfigs_Response{Rows: make([]*v1.ListGameTownAgentConfigs_Response_GameTownAgentConfig, 0, len(listResponse.Rows))}
	for _, row := range listResponse.Rows {
		if row == nil {
			reply.Rows = append(reply.Rows, nil)
			continue
		}
		reply.Rows = append(reply.Rows, &v1.ListGameTownAgentConfigs_Response_GameTownAgentConfig{
			CreatedAt:      timestamp(row.CreatedAt),
			UpdatedAt:      timestamp(row.UpdatedAt),
			Id:             row.ID,
			PlayerId:       row.PlayerID,
			Name:           row.Name,
			Provider:       row.Provider,
			Model:          row.Model,
			BaseUrl:        row.BaseURL,
			HasApiKey:      row.APIKey != "",
			TimeoutSeconds: row.TimeoutSeconds,
			IsDefault:      row.IsDefault,
			Status:         gameenum.AgentConfigStatusMap.MustToProto(gameenum.AgentConfigStatus(row.Status)),
		})
	}
	return reply, nil
}
