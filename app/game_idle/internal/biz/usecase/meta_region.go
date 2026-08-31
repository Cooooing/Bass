package usecase

import (
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"sort"
)

type MetaRegionUsecase struct {
	regionRepo repo.MetaRegionRepo
}

func NewMetaRegionUsecase(regionRepo repo.MetaRegionRepo) *MetaRegionUsecase {
	return &MetaRegionUsecase{
		regionRepo: regionRepo,
	}
}

func (u *MetaRegionUsecase) List(ctx context.Context) ([]*model.MetaRegion, error) {
	rows, err := u.regionRepo.Map(ctx, nil)
	if err != nil {
		return nil, err
	}
	out := make([]*model.MetaRegion, 0, len(rows))
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

func (u *MetaRegionUsecase) Refresh(ctx context.Context) error {
	return u.regionRepo.Refresh(ctx)
}
