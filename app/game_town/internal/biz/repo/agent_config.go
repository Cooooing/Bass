package repo

import (
	"context"

	"game_town/internal/biz/base"
	"game_town/internal/biz/model"
	"game_town/internal/enum"
)

type AgentConfigRepo interface {
	Save(ctx context.Context, config *model.AgentConfig) (*model.AgentConfig, error)
	Get(ctx context.Context, req *AgentConfigQuery) (*model.AgentConfig, error)
	List(ctx context.Context, req *AgentConfigQuery) ([]*model.AgentConfig, error)
	Map(ctx context.Context, req *AgentConfigQuery) (map[int64]*model.AgentConfig, error)
	Count(ctx context.Context, req *AgentConfigQuery) (int, error)
	Page(ctx context.Context, req *AgentConfigPageReq) (*AgentConfigPageResp, error)
}

type AgentConfigQuery struct {
	ID       *int64
	IDs      []int64
	Provider *enum.AgentProvider
}

type AgentConfigPageReq struct {
	Page  base.PageRequest
	Query AgentConfigQuery
}

type AgentConfigPageResp struct {
	Rows []*model.AgentConfig
	Page base.PageResp
}
