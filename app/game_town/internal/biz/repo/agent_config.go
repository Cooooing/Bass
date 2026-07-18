package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type AgentConfigRepo interface {
	CreateAgentConfig(ctx context.Context, row *model.AgentConfig) (*model.AgentConfig, error)
	GetAgentConfig(ctx context.Context, req *GetAgentConfigReq) (*model.AgentConfig, error)
	GetDefaultAgentConfig(ctx context.Context, playerID int64) (*model.AgentConfig, error)
	ListAgentConfigs(ctx context.Context, req *ListAgentConfigsReq) ([]*model.AgentConfig, error)
}

type GetAgentConfigReq struct {
	ID       int64
	PlayerID int64
}

type ListAgentConfigsReq struct {
	PlayerID int64
	Status   *string
}
