package usecase

import (
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"sort"
)

type MetaItemUsecase struct {
	itemRepo repo.MetaItemRepo
}

func NewMetaItemUsecase(itemRepo repo.MetaItemRepo) *MetaItemUsecase {
	return &MetaItemUsecase{
		itemRepo: itemRepo,
	}
}

func (u *MetaItemUsecase) List(ctx context.Context) ([]*model.MetaItem, error) {
	rows, err := u.itemRepo.Map(ctx, nil)
	if err != nil {
		return nil, err
	}
	out := make([]*model.MetaItem, 0, len(rows))
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

func (u *MetaItemUsecase) Refresh(ctx context.Context) error {
	return u.itemRepo.Refresh(ctx)
}
