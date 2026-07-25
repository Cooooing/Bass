package usecase

import (
	"context"

	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
)

// LocationUsecase 只提供玩家已知地点的查询入口。
type LocationUsecase struct {
	locationRepo    repo.LocationRepo
	worldMemberRepo repo.WorldMemberRepo
}

func NewLocationUsecase(
	locationRepo repo.LocationRepo,
	worldMemberRepo repo.WorldMemberRepo,
) *LocationUsecase {
	return &LocationUsecase{
		locationRepo:    locationRepo,
		worldMemberRepo: worldMemberRepo,
	}
}

func (u *LocationUsecase) Get(
	ctx context.Context,
	worldID, playerID, locationID int64,
) (*model.Location, error) {
	if _, err := u.member(ctx, worldID, playerID); err != nil {
		return nil, err
	}
	return u.locationRepo.Get(ctx, &repo.LocationQuery{
		ID:      new(locationID),
		WorldID: new(worldID),
	})
}

func (u *LocationUsecase) List(
	ctx context.Context,
	worldID, playerID int64,
) ([]*model.Location, *model.WorldMember, error) {
	member, err := u.member(ctx, worldID, playerID)
	if err != nil {
		return nil, nil, err
	}
	rows, err := u.locationRepo.List(ctx, &repo.LocationQuery{
		WorldID: new(worldID),
	})
	return rows, member, err
}

func (u *LocationUsecase) member(
	ctx context.Context,
	worldID, playerID int64,
) (*model.WorldMember, error) {
	return u.worldMemberRepo.Get(ctx, &repo.WorldMemberQuery{
		WorldID:  new(worldID),
		PlayerID: new(playerID),
	})
}
