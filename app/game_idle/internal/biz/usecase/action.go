package usecase

import (
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/enum"
	"sort"
)

type ActionUsecase struct {
	itemRepo   repo.ItemRepo
	actionRepo repo.ActionRepo
	recipeRepo repo.RecipeRepo
}

func NewActionUsecase(
	itemRepo repo.ItemRepo,
	actionRepo repo.ActionRepo,
	recipeRepo repo.RecipeRepo,
) *ActionUsecase {
	return &ActionUsecase{
		itemRepo:   itemRepo,
		actionRepo: actionRepo,
		recipeRepo: recipeRepo,
	}
}

func (u *ActionUsecase) List(ctx context.Context) ([]*model.Action, error) {
	rows, err := u.actionRepo.Map(ctx, nil)
	if err != nil {
		return nil, err
	}
	out := make([]*model.Action, 0, len(rows))
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

func (u *ActionUsecase) GetDetail(ctx context.Context, actionID string) (*model.ActionDetail, error) {
	action, err := u.actionRepo.Get(ctx, actionID)
	if err != nil {
		return nil, err
	}
	recipeIDs := make([]string, 0, len(action.Recipes))
	for _, binding := range action.Recipes {
		recipeIDs = append(recipeIDs, binding.RecipeID)
	}
	recipes, err := u.recipeRepo.Map(ctx, recipeIDs)
	if err != nil {
		return nil, err
	}
	recipeRows := make([]*model.Recipe, 0, len(recipeIDs))
	itemIDs := make([]string, 0)
	for _, recipeID := range recipeIDs {
		recipe := recipes[recipeID]
		if recipe == nil {
			return nil, model.ErrRecipeInvalid
		}
		recipeRows = append(recipeRows, recipe)
		for _, input := range recipe.Inputs {
			itemIDs = append(itemIDs, input.ItemID)
		}
		for _, output := range recipe.Outputs {
			if output.ItemID == enum.ItemIDEmpty.String() {
				continue
			}
			itemIDs = append(itemIDs, output.ItemID)
		}
	}
	items, err := u.itemRepo.Map(ctx, itemIDs)
	if err != nil {
		return nil, err
	}
	return &model.ActionDetail{
		Action:  action,
		Recipes: recipeRows,
		Items:   items,
	}, nil
}
