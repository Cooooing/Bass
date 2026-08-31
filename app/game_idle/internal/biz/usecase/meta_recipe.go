package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_idle/internal/biz/repo"
	"game_idle/internal/enum"
	"math/rand"
)

type MetaRecipeUsecase struct {
	recipeRepo repo.MetaRecipeRepo
}

func NewMetaRecipeUsecase(recipeRepo repo.MetaRecipeRepo) *MetaRecipeUsecase {
	return &MetaRecipeUsecase{recipeRepo: recipeRepo}
}

func (u *MetaRecipeUsecase) Refresh(ctx context.Context) error {
	return u.recipeRepo.Refresh(ctx)
}

type RollRecipeReq struct {
	RecipeID string
	// TODO 后续在这里承载角色属性、装备、Buff 等产出修正参数。
}

func (u *MetaRecipeUsecase) RollNormal(ctx context.Context, req *RollRecipeReq) (map[string]int64, error) {
	recipe, err := u.recipeRepo.Get(ctx, req.RecipeID)
	if err != nil {
		return nil, err
	}
	if recipe == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_RECIPE_INVALID)
	}

	outputs := recipe.Outputs
	itemQuantities := make(map[string]int64, len(outputs))
	for generationIndex := int32(0); generationIndex < recipe.GenerationTimes; generationIndex++ {
		cursor := rand.Int63n(recipe.TotalWeight) + 1
		left := 0
		right := len(outputs) - 1
		for left < right {
			middle := int(uint(left+right) >> 1)
			if cursor <= outputs[middle].WeightLimit {
				right = middle
			} else {
				left = middle + 1
			}
		}
		selected := outputs[left]
		if selected.ItemID == enum.ItemIDEmpty.String() {
			continue
		}
		itemQuantities[selected.ItemID] += selected.MinQuantity + rand.Int63n(selected.MaxQuantity-selected.MinQuantity+1)
	}
	return itemQuantities, nil
}
