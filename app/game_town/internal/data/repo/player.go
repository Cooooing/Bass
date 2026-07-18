package repo

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_town/internal/biz/model"
	bizrepo "game_town/internal/biz/repo"
	"game_town/internal/data/gen"
	"game_town/internal/data/gen/player"
)

type PlayerRepo struct{ *baseRepo }

func NewPlayerRepo(db *gen.Client) bizrepo.PlayerRepo {
	return &PlayerRepo{baseRepo: &baseRepo{db: db}}
}

func (r *PlayerRepo) GetPlayer(ctx context.Context, id int64) (*model.Player, error) {
	row, err := r.db.Player.Query().Where(player.ID(id), player.DeletedAtIsNil()).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_PLAYER_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.player(row), nil
}

func (r *PlayerRepo) GetPlayerByName(ctx context.Context, name string) (*model.Player, error) {
	row, err := r.db.Player.Query().Where(player.Name(name), player.DeletedAtIsNil()).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_PLAYER_NOT_FOUND)
	}
	if err != nil {
		return nil, err
	}
	return r.player(row), nil
}

func (r *PlayerRepo) CreatePlayer(ctx context.Context, req *bizrepo.CreatePlayerReq) (*model.Player, error) {
	row, err := r.db.Player.Create().SetName(req.Name).SetDisplayName(req.DisplayName).SetStatus("active").Save(ctx)
	if gen.IsConstraintError(err) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_PLAYER_NAME_TAKEN)
	}
	if err != nil {
		return nil, err
	}
	return r.player(row), nil
}
