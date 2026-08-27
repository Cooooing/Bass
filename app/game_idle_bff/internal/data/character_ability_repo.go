package data

import (
	"common/pkg/client/rpc"
	gameidlev1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"
)

var _ repo.CharacterAbilityRepo = (*CharacterAbilityRepo)(nil)

type CharacterAbilityRepo struct {
	gameIdleClient *rpc.GameIdleClient
}

func NewCharacterAbilityRepo(
	gameIdleClient *rpc.GameIdleClient,
) repo.CharacterAbilityRepo {
	return &CharacterAbilityRepo{
		gameIdleClient: gameIdleClient,
	}
}

func (r *CharacterAbilityRepo) Map(
	ctx context.Context,
	characterID int64,
) (map[string]*model.CharacterAbility, error) {
	reply, err := r.gameIdleClient.Ability.Get(ctx, &gameidlev1.GetCharacterAbility_Request{
		CharacterId: characterID,
	})
	if err != nil {
		return nil, err
	}
	rows := make(map[string]*model.CharacterAbility, len(reply.GetRows()))
	for _, row := range reply.GetRows() {
		rows[row.GetAbilityId()] = &model.CharacterAbility{
			AbilityID:    row.GetAbilityId(),
			Level:        row.GetLevel(),
			Exp:          row.GetExp(),
			NextLevelExp: row.GetNextLevelExp(),
		}
	}
	return rows, nil
}
