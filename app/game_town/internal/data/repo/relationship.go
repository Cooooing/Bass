package repo

import (
	"context"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/relationship"
	"time"
)

type RelationshipRepo struct{ *baseRepo }

func NewRelationshipRepo(db *gen.Client) bizrepo.RelationshipRepo {
	return &RelationshipRepo{baseRepo: &baseRepo{db: db}}
}

func (r *RelationshipRepo) GetRelationship(ctx context.Context, req *bizrepo.GetRelationshipReq) (*model.Relationship, error) {
	row, err := r.db.Relationship.Query().Where(relationship.WorldID(req.WorldID), relationship.PlayerID(req.PlayerID), relationship.NpcID(req.NpcID)).Only(ctx)
	if gen.IsNotFound(err) {
		return &model.Relationship{WorldID: req.WorldID, PlayerID: req.PlayerID, NpcID: req.NpcID, Affinity: 50, Trust: 50, Tension: 0, CustomMetrics: map[string]any{}}, nil
	}
	if err != nil {
		return nil, err
	}
	return r.relationship(row), nil
}

func (r *RelationshipRepo) UpsertRelationship(ctx context.Context, row *model.Relationship) (*model.Relationship, error) {
	current, err := r.db.Relationship.Query().Where(relationship.WorldID(row.WorldID), relationship.PlayerID(row.PlayerID), relationship.NpcID(row.NpcID)).Only(ctx)
	now := time.Now()
	if gen.IsNotFound(err) {
		created, err := r.db.Relationship.Create().SetWorldID(row.WorldID).SetPlayerID(row.PlayerID).SetNpcID(row.NpcID).SetAffinity(row.Affinity).SetTrust(row.Trust).SetTension(row.Tension).SetCustomMetrics(row.CustomMetrics).SetLastInteractionAt(now).Save(ctx)
		if err != nil {
			return nil, err
		}
		return r.relationship(created), nil
	}
	if err != nil {
		return nil, err
	}
	updated, err := r.db.Relationship.UpdateOneID(current.ID).SetAffinity(row.Affinity).SetTrust(row.Trust).SetTension(row.Tension).SetCustomMetrics(row.CustomMetrics).SetLastInteractionAt(now).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.relationship(updated), nil
}
