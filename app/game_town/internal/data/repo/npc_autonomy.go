package repo

import (
	"context"
	"fmt"

	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen/npc"
)

func (r *NpcRepo) UpdateAutonomy(ctx context.Context, req *bizrepo.NpcAutonomyUpdateReq) (*model.Npc, error) {
	update := r.getClient(ctx).Npc.Update().Where(npc.ID(req.NpcID), npc.Version(req.Version), npc.DeletedAtIsNil())
	if req.Goal != "" {
		update.SetGoal(req.Goal)
	}
	if req.ContextSummary != "" {
		update.SetContextSummary(req.ContextSummary)
	}
	update.SetNillableNextDecisionAt(req.NextDecisionAt).SetNillableLastPlannedAt(req.LastPlannedAt).AddVersion(1)
	count, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("npc version conflict")
	}
	return r.Get(ctx, &bizrepo.NpcQuery{
		ID: new(req.NpcID),
	})
}

func (r *NpcRepo) UpdateState(ctx context.Context, req *bizrepo.NpcStateUpdateReq) (*model.Npc, error) {
	update := r.getClient(ctx).Npc.Update().Where(npc.ID(req.NpcID), npc.Version(req.Version), npc.DeletedAtIsNil())
	if req.CurrentLocationID != nil {
		update.SetCurrentLocationID(*req.CurrentLocationID)
	}
	if req.LifeStatus != nil {
		update.SetLifeStatus(npc.LifeStatus(*req.LifeStatus))
	}
	if req.StateTags != nil {
		update.SetStateTags(req.StateTags)
	}
	if req.Attributes != nil {
		update.SetAttributes(req.Attributes)
	}
	if req.DeathWorldTime != nil {
		update.SetDeathWorldTime(*req.DeathWorldTime)
	}
	update.AddVersion(1)
	count, err := update.Save(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, fmt.Errorf("npc version conflict")
	}
	return r.Get(ctx, &bizrepo.NpcQuery{
		ID: new(req.NpcID),
	})
}
