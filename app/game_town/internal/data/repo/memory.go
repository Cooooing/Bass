package repo

import (
	"context"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/memory"
)

type MemoryRepo struct{ *baseRepo }

func NewMemoryRepo(db *gen.Client) bizrepo.MemoryRepo {
	return &MemoryRepo{baseRepo: &baseRepo{db: db}}
}

func (r *MemoryRepo) ListMemories(ctx context.Context, queryReq *bizrepo.MemoryQuery) ([]*model.Memory, error) {
	query := r.db.Memory.Query().Where(memory.WorldID(queryReq.WorldID), memory.PlayerID(queryReq.PlayerID), memory.DeletedAtIsNil())
	if queryReq.NpcID != nil {
		query = query.Where(memory.NpcID(*queryReq.NpcID))
	}
	if queryReq.Type != nil {
		query = query.Where(memory.Type(*queryReq.Type))
	}
	rows, err := query.Order(gen.Desc(memory.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Memory, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.memory(row))
	}
	return result, nil
}

func (r *MemoryRepo) CreateMemory(ctx context.Context, row *model.Memory) (*model.Memory, error) {
	created, err := r.db.Memory.Create().SetWorldID(row.WorldID).SetPlayerID(row.PlayerID).SetNpcID(row.NpcID).SetType(row.Type).SetContent(row.Content).SetImportance(row.Importance).SetNillableSourceEventID(row.SourceEventID).SetNillableLastRecalledAt(row.LastRecalledAt).SetNillableExpiresAt(row.ExpiresAt).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.memory(created), nil
}
