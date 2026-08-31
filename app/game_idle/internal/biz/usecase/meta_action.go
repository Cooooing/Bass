package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/repo"
	"game_idle/internal/enum"
	"sort"
)

type MetaActionUsecase struct {
	itemRepo   repo.MetaItemRepo
	actionRepo repo.MetaActionRepo
	recipeRepo repo.MetaRecipeRepo
}

func NewMetaActionUsecase(
	itemRepo repo.MetaItemRepo,
	actionRepo repo.MetaActionRepo,
	recipeRepo repo.MetaRecipeRepo,
) *MetaActionUsecase {
	return &MetaActionUsecase{
		itemRepo:   itemRepo,
		actionRepo: actionRepo,
		recipeRepo: recipeRepo,
	}
}

func (u *MetaActionUsecase) List(ctx context.Context) ([]*model.MetaAction, error) {
	rows, err := u.actionRepo.Map(ctx, nil)
	if err != nil {
		return nil, err
	}
	out := make([]*model.MetaAction, 0, len(rows))
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

func (u *MetaActionUsecase) Refresh(ctx context.Context) error {
	return u.actionRepo.Refresh(ctx)
}

func (u *MetaActionUsecase) GetDetail(ctx context.Context, actionID string) (*model.MetaActionDetail, error) {
	action, err := u.actionRepo.Get(ctx, actionID)
	if err != nil {
		return nil, err
	}
	if action == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
	}
	recipeIDs := make([]string, 0, len(action.Recipes))
	for _, binding := range action.Recipes {
		recipeIDs = append(recipeIDs, binding.RecipeID)
	}
	recipes, err := u.recipeRepo.Map(ctx, recipeIDs)
	if err != nil {
		return nil, err
	}
	recipeRows := make([]*model.MetaRecipe, 0, len(recipeIDs))
	itemIDs := make([]string, 0)
	for _, recipeID := range recipeIDs {
		recipe := recipes[recipeID]
		if recipe == nil {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_RECIPE_INVALID)
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
	return &model.MetaActionDetail{
		Action:  action,
		Recipes: recipeRows,
		Items:   items,
	}, nil
}
