package repo

import (
	"context"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"time"
)

type AgentRunRepo struct{ *baseRepo }

func NewAgentRunRepo(db *gen.Client) bizrepo.AgentRunRepo {
	return &AgentRunRepo{baseRepo: &baseRepo{db: db}}
}

func (r *AgentRunRepo) CreateAgentRun(ctx context.Context, req *bizrepo.CreateAgentRunReq) (*bizrepo.CreateAgentRunResponse, error) {
	now := time.Now()
	row := req.Row
	created, err := r.db.AgentRun.Create().SetNillableWorldID(row.WorldID).SetNillableAgentConfigID(row.AgentConfigID).SetRunType(row.RunType).SetNillableCommandID(row.CommandID).SetNillableEventID(row.EventID).SetNillableNpcID(row.NpcID).SetModel(row.Model).SetInputJSON(row.InputJSON).SetOutputJSON(row.OutputJSON).SetStatus(row.Status).SetErrorSummary(row.ErrorSummary).SetLatencyMs(row.LatencyMS).SetCreatedAt(now).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &bizrepo.CreateAgentRunResponse{Row: r.agentRun(created)}, nil
}
