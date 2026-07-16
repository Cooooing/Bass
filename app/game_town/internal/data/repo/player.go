package repo

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/player"
)

type PlayerRepo struct{ *baseRepo }

func NewPlayerRepo(db *gen.Client) bizrepo.PlayerRepo {
	return &PlayerRepo{baseRepo: &baseRepo{db: db}}
}

func (r *PlayerRepo) GetPlayer(ctx context.Context, req *bizrepo.GetPlayerReq) (*bizrepo.GetPlayerResponse, error) {
	row, err := r.db.Player.Query().Where(player.ID(req.ID), player.DeletedAtIsNil()).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_PLAYER_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.GetPlayerResponse{Row: r.player(row)}, nil
}

func (r *PlayerRepo) GetPlayerByName(ctx context.Context, req *bizrepo.GetPlayerByNameReq) (*bizrepo.GetPlayerByNameResponse, error) {
	row, err := r.db.Player.Query().Where(player.Name(req.Name), player.DeletedAtIsNil()).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_PLAYER_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.GetPlayerByNameResponse{Row: r.player(row)}, nil
}

func (r *PlayerRepo) CreatePlayer(ctx context.Context, req *bizrepo.CreatePlayerReq) (*bizrepo.CreatePlayerResponse, error) {
	row, err := r.db.Player.Create().SetName(req.Name).SetDisplayName(req.DisplayName).SetStatus("active").Save(ctx)
	if gen.IsConstraintError(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_PLAYER_NAME_TAKEN)
	}
	if err != nil {
		return nil, err
	}
	return &bizrepo.CreatePlayerResponse{Row: r.player(row)}, nil
}
