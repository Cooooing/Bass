package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type AgentRunRepo interface {
	CreateAgentRun(ctx context.Context, req *CreateAgentRunReq) (*CreateAgentRunResponse, error)
}

type CreateAgentRunReq struct {
	Row *model.AgentRun
}

type CreateAgentRunResponse struct {
	Row *model.AgentRun
}
