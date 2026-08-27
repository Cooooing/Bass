package data

import (
	"common/pkg/client/rpc"
	gameidlev1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle_bff/internal/biz/model"
	"game_idle_bff/internal/biz/repo"
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
		rows = append(rows, r.buildActionConfig(row))
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
		Action:  r.buildActionConfig(row.GetAction()),
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

func (r *ActionRepo) buildActionConfig(row *gameidlev1.Action) *model.ActionConfig {
	return &model.ActionConfig{
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
	}
}
