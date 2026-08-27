package usecase

import (
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"sort"
)

type RegionUsecase struct {
	regionRepo repo.RegionRepo
}

func NewRegionUsecase(regionRepo repo.RegionRepo) *RegionUsecase {
	return &RegionUsecase{
		regionRepo: regionRepo,
	}
}

func (u *RegionUsecase) List(ctx context.Context) ([]*model.Region, error) {
	rows, err := u.regionRepo.Map(ctx, nil)
	if err != nil {
		return nil, err
	}
	out := make([]*model.Region, 0, len(rows))
	for _, row := range rows {
		out = append(out, row)
	}
	sort.SliceStable(out, func(left, right int) bool {
		if out[left].Sort == out[right].Sort {
			return out[left].ID < out[right].ID
		}
		return out[left].Sort < out[right].Sort
	})
	return out, nil
}
