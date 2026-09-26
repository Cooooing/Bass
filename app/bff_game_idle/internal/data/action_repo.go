package data

import (
	"bff_game_idle/internal/biz/model"
	"bff_game_idle/internal/biz/repo"
	"common/pkg/client/rpc"
	gameidlev1 "common/proto/gen/game_idle/v1"
	"context"
)

var _ repo.ActionRepo = (*ActionRepo)(nil)

type ActionRepo struct {
	gameIdleClient *rpc.GameIdleClient
}

func NewActionRepo(gameIdleClient *rpc.GameIdleClient) repo.ActionRepo {
	return &ActionRepo{
		gameIdleClient: gameIdleClient,
	}
}

func (r *ActionRepo) List(ctx context.Context) ([]*model.ActionConfig, error) {
	reply, err := r.gameIdleClient.Action.List(ctx, &gameidlev1.ListActions_Request{})
	if err != nil {
		return nil, err
	}
	rows := make([]*model.ActionConfig, 0, len(reply.GetRows()))
	for _, row := range reply.GetRows() {
		rows = append(rows, &model.ActionConfig{
			ActionID:             row.GetActionId(),
			Name:                 row.GetName(),
			Description:          row.GetDescription(),
			RegionID:             row.GetRegionId(),
			ActionKind:           row.GetActionKind(),
			AbilityID:            row.GetAbilityId(),
			RequiredAbilityLevel: row.GetRequiredAbilityLevel(),
			DurationSeconds:      row.GetDurationSeconds(),
			ExpReward:            row.GetExpReward(),
			Enabled:              row.GetEnabled(),
			Sort:                 row.GetSort(),
		})
	}
	return rows, nil
}

func (r *ActionRepo) GetDetail(ctx context.Context, actionID string) (*model.ActionDetailConfig, error) {
	reply, err := r.gameIdleClient.Action.GetDetail(ctx, &gameidlev1.GetActionDetail_Request{
		ActionId: actionID,
	})
	if err != nil {
		return nil, err
	}
	row := reply.GetRow()
	out := &model.ActionDetailConfig{
		Action: &model.ActionConfig{
			ActionID:             row.GetAction().GetActionId(),
			Name:                 row.GetAction().GetName(),
			Description:          row.GetAction().GetDescription(),
			RegionID:             row.GetAction().GetRegionId(),
			ActionKind:           row.GetAction().GetActionKind(),
			AbilityID:            row.GetAction().GetAbilityId(),
			RequiredAbilityLevel: row.GetAction().GetRequiredAbilityLevel(),
			DurationSeconds:      row.GetAction().GetDurationSeconds(),
			ExpReward:            row.GetAction().GetExpReward(),
			Enabled:              row.GetAction().GetEnabled(),
			Sort:                 row.GetAction().GetSort(),
		},
		Recipes: make([]*model.ActionRecipeConfig, 0, len(row.GetRecipes())),
	}
	for _, recipe := range row.GetRecipes() {
		recipeRow := &model.ActionRecipeConfig{
			RecipeID:        recipe.GetRecipeId(),
			Name:            recipe.GetName(),
			Description:     recipe.GetDescription(),
			RecipeType:      recipe.GetRecipeType(),
			GenerationTimes: recipe.GetGenerationTimes(),
			Inputs:          make([]*model.RecipeInputConfig, 0, len(recipe.GetInputs())),
			Outputs:         make([]*model.RecipeOutputConfig, 0, len(recipe.GetOutputs())),
		}
		for _, input := range recipe.GetInputs() {
			recipeRow.Inputs = append(recipeRow.Inputs, &model.RecipeInputConfig{
				ItemID:   input.GetItemId(),
				ItemName: input.GetItemName(),
				ItemType: input.GetItemType(),
				Quantity: input.GetQuantity(),
			})
		}
		for _, output := range recipe.GetOutputs() {
			recipeRow.Outputs = append(recipeRow.Outputs, &model.RecipeOutputConfig{
				ItemID:      output.GetItemId(),
				ItemName:    output.GetItemName(),
				ItemType:    output.GetItemType(),
				MinQuantity: output.GetMinQuantity(),
				MaxQuantity: output.GetMaxQuantity(),
				Weight:      output.GetWeight(),
				Probability: output.GetProbability(),
			})
		}
		out.Recipes = append(out.Recipes, recipeRow)
	}
	return out, nil
}
