package repo

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/npc"
)

type NpcRepo struct{ *baseRepo }

func NewNpcRepo(db *gen.Client) bizrepo.NpcRepo {
	return &NpcRepo{baseRepo: &baseRepo{db: db}}
}

func (r *NpcRepo) ListNpcs(ctx context.Context, req *bizrepo.ListNpcsReq) (*bizrepo.ListNpcsResponse, error) {
	query := r.db.Npc.Query().Where(npc.WorldID(req.WorldID), npc.DeletedAtIsNil(), npc.Enabled(true))
	if req.LocationID != nil {
		query = query.Where(npc.CurrentLocationID(*req.LocationID))
	}
	rows, err := query.Order(npc.ByID()).All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Npc, 0, len(rows))
	for _, row := range rows {
		result = append(result, r.npc(row))
	}
	return &bizrepo.ListNpcsResponse{Rows: result}, nil
}

func (r *NpcRepo) GetNpc(ctx context.Context, req *bizrepo.GetNpcReq) (*bizrepo.GetNpcResponse, error) {
	row, err := r.db.Npc.Query().Where(npc.ID(req.ID), npc.DeletedAtIsNil()).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_NPC_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.GetNpcResponse{Row: r.npc(row)}, nil
}

func (r *NpcRepo) GetNpcByCode(ctx context.Context, req *bizrepo.GetNpcByCodeReq) (*bizrepo.GetNpcByCodeResponse, error) {
	row, err := r.db.Npc.Query().Where(npc.WorldID(req.WorldID), npc.Code(req.Code), npc.DeletedAtIsNil()).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_NPC_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.GetNpcByCodeResponse{Row: r.npc(row)}, nil
}
