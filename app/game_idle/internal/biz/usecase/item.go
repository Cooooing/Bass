package usecase

import (
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"sort"
)

type ItemUsecase struct {
	itemRepo repo.ItemRepo
}

func NewItemUsecase(itemRepo repo.ItemRepo) *ItemUsecase {
	return &ItemUsecase{
		itemRepo: itemRepo,
	}
}

func (u *ItemUsecase) List(ctx context.Context) ([]*model.Item, error) {
	rows, err := u.itemRepo.Map(ctx, nil)
	if err != nil {
		return nil, err
	}
	out := make([]*model.Item, 0, len(rows))
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
