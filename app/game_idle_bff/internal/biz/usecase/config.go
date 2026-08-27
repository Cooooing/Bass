package usecase

import (
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"
	"time"
)

type ConfigUsecase struct {
	regionRepo repo.RegionRepo
	actionRepo repo.ActionRepo
}

func NewConfigUsecase(
	regionRepo repo.RegionRepo,
	actionRepo repo.ActionRepo,
) *ConfigUsecase {
	return &ConfigUsecase{
		regionRepo: regionRepo,
		actionRepo: actionRepo,
	}
}

func (u *ConfigUsecase) Get(ctx context.Context) (*model.GameConfig, error) {
	regions, err := u.regionRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	actions, err := u.actionRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return &model.GameConfig{
		Regions:    regions,
		Actions:    actions,
		ServerTime: time.Now().Unix(),
	}, nil
}

func (u *ConfigUsecase) GetActionDetail(ctx context.Context, actionID string) (*model.ActionDetailConfig, error) {
	return u.actionRepo.GetDetail(ctx, actionID)
}
