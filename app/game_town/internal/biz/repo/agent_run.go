package repo

import (
	"context"
	"game_town/internal/biz/model"
)

type AgentRunRepo interface {
	CreateAgentRun(ctx context.Context, row *model.AgentRun) (*model.AgentRun, error)
}
