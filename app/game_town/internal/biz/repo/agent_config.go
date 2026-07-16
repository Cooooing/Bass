package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type AgentConfigRepo interface {
	CreateAgentConfig(ctx context.Context, req *CreateAgentConfigReq) (*CreateAgentConfigResponse, error)
	GetAgentConfig(ctx context.Context, req *GetAgentConfigReq) (*GetAgentConfigResponse, error)
	GetDefaultAgentConfig(ctx context.Context, req *GetDefaultAgentConfigReq) (*GetDefaultAgentConfigResponse, error)
	ListAgentConfigs(ctx context.Context, req *ListAgentConfigsReq) (*ListAgentConfigsResponse, error)
}

type CreateAgentConfigReq struct {
	Row *model.AgentConfig
}

type CreateAgentConfigResponse struct {
	Row *model.AgentConfig
}

type GetAgentConfigReq struct {
	ID       int64
	PlayerID int64
}

type GetAgentConfigResponse struct {
	Row *model.AgentConfig
}

type GetDefaultAgentConfigReq struct {
	PlayerID int64
}

type GetDefaultAgentConfigResponse struct {
	Row *model.AgentConfig
}

type ListAgentConfigsReq struct {
	PlayerID int64
	Status   *string
}

type ListAgentConfigsResponse struct {
	Rows []*model.AgentConfig
}
