package usecase

import (
	"context"

	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
)

// FactionUsecase 只返回公开阵营资料，不暴露秘密关系或内部数值。
type FactionUsecase struct {
	factionRepo     repo.FactionRepo
	worldMemberRepo repo.WorldMemberRepo
}

func NewFactionUsecase(
	factionRepo repo.FactionRepo,
	worldMemberRepo repo.WorldMemberRepo,
) *FactionUsecase {
	return &FactionUsecase{
		factionRepo:     factionRepo,
		worldMemberRepo: worldMemberRepo,
	}
}

func (u *FactionUsecase) Get(ctx context.Context, worldID, playerID, factionID int64) (*model.Faction, error) {
	if err := u.requireMember(ctx, worldID, playerID); err != nil {
		return nil, err
	}
	return u.factionRepo.Get(ctx, &repo.FactionQuery{
		ID:      new(factionID),
		WorldID: new(worldID),
	})
}

func (u *FactionUsecase) List(ctx context.Context, worldID, playerID int64) ([]*model.Faction, error) {
	if err := u.requireMember(ctx, worldID, playerID); err != nil {
		return nil, err
	}
	return u.factionRepo.List(ctx, &repo.FactionQuery{
		WorldID: new(worldID),
	})
}

func (u *FactionUsecase) requireMember(ctx context.Context, worldID, playerID int64) error {
	_, err := u.worldMemberRepo.Get(ctx, &repo.WorldMemberQuery{
		WorldID:  new(worldID),
		PlayerID: new(playerID),
	})
	return err
}
