package repo

import (
	"context"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"time"
)

type AgentRunRepo struct{ *baseRepo }

func NewAgentRunRepo(db *gen.Client) bizrepo.AgentRunRepo {
	return &AgentRunRepo{baseRepo: &baseRepo{db: db}}
}

func (r *AgentRunRepo) CreateAgentRun(ctx context.Context, row *model.AgentRun) (*model.AgentRun, error) {
	now := time.Now()
	created, err := r.db.AgentRun.Create().SetNillableWorldID(row.WorldID).SetNillableAgentConfigID(row.AgentConfigID).SetRunType(row.RunType).SetNillableCommandID(row.CommandID).SetNillableEventID(row.EventID).SetNillableNpcID(row.NpcID).SetModel(row.Model).SetInputJSON(row.InputJSON).SetOutputJSON(row.OutputJSON).SetStatus(row.Status).SetErrorSummary(row.ErrorSummary).SetLatencyMs(row.LatencyMS).SetCreatedAt(now).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.agentRun(created), nil
}
