package repo

import (
	"context"
	"fmt"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/data/gen"
	recipeent "game_idle/internal/data/gen/recipe"
	recipeinputent "game_idle/internal/data/gen/recipeinput"
	recipeoutputent "game_idle/internal/data/gen/recipeoutput"
	"game_idle/internal/enum"
	"sync"
)

var _ bizrepo.MetaRecipeRepo = (*MetaRecipeRepo)(nil)

type MetaRecipeRepo struct {
	mutex   sync.RWMutex
	db      *gen.Client
	recipes map[string]*model.MetaRecipe
}

func NewMetaRecipeRepo(db *gen.Client) (bizrepo.MetaRecipeRepo, error) {
	repo := &MetaRecipeRepo{
		db: db,
	}
	if err := repo.Refresh(context.Background()); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *MetaRecipeRepo) Refresh(ctx context.Context) error {
	rows, err := r.db.Recipe.Query().
		Where(recipeent.DeletedAtIsNil()).
		WithInputs(func(query *gen.RecipeInputQuery) {
			query.Where(recipeinputent.DeletedAtIsNil()).
				Order(recipeinputent.BySort(), recipeinputent.ByID())
		}).
		WithOutputs(func(query *gen.RecipeOutputQuery) {
			query.Where(recipeoutputent.DeletedAtIsNil()).
				Order(recipeoutputent.BySort(), recipeoutputent.ByID())
		}).
		Order(recipeent.BySort(), recipeent.ByID()).
		All(ctx)
	if err != nil {
		return err
	}
	recipes := make(map[string]*model.MetaRecipe, len(rows))
	for _, row := range rows {
		recipe := &model.MetaRecipe{
			ID:              row.ID,
			Name:            row.Name,
			Description:     row.Description,
			Type:            enum.RecipeType(row.Type),
			GenerationTimes: row.GenerationTimes,
			Enabled:         row.Enabled,
			Inputs:          make([]*model.MetaRecipeInput, 0, len(row.Edges.Inputs)),
			Outputs:         make([]*model.MetaRecipeOutput, 0, len(row.Edges.Outputs)),
		}
		if recipe.ID == "" || recipe.GenerationTimes <= 0 || len(row.Edges.Outputs) == 0 {
			return fmt.Errorf("game idle recipe config invalid: %s", recipe.ID)
		}
		for _, input := range row.Edges.Inputs {
			if input.ItemID == "" || input.Quantity <= 0 {
				return fmt.Errorf("game idle recipe input config invalid: %s", recipe.ID)
			}
			recipe.Inputs = append(recipe.Inputs, &model.MetaRecipeInput{
				ID:       input.ID,
				RecipeID: input.RecipeID,
				ItemID:   input.ItemID,
				Quantity: input.Quantity,
				Sort:     input.Sort,
			})
		}
		for _, output := range row.Edges.Outputs {
			if output.ItemID == "" || output.MinQuantity <= 0 || output.MaxQuantity < output.MinQuantity || output.Weight <= 0 {
				return fmt.Errorf("game idle recipe output config invalid: %s", recipe.ID)
			}
			recipe.TotalWeight += int64(output.Weight)
			recipe.Outputs = append(recipe.Outputs, &model.MetaRecipeOutput{
				ID:          output.ID,
				RecipeID:    output.RecipeID,
				ItemID:      output.ItemID,
				MinQuantity: output.MinQuantity,
				MaxQuantity: output.MaxQuantity,
				Weight:      output.Weight,
				WeightLimit: recipe.TotalWeight,
				Sort:        output.Sort,
			})
		}
		if recipe.TotalWeight <= 0 {
			return fmt.Errorf("game idle recipe output weight invalid: %s", recipe.ID)
		}
		recipes[recipe.ID] = recipe
	}
	r.mutex.Lock()
	r.recipes = recipes
	r.mutex.Unlock()
	return nil
}

func (r *MetaRecipeRepo) Get(ctx context.Context, recipeID string) (*model.MetaRecipe, error) {
	r.mutex.RLock()
	recipe := r.recipes[recipeID]
	r.mutex.RUnlock()
	return recipe, nil
}

func (r *MetaRecipeRepo) Map(ctx context.Context, recipeIDs []string) (map[string]*model.MetaRecipe, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if len(recipeIDs) == 0 {
		recipes := make(map[string]*model.MetaRecipe, len(r.recipes))
		for recipeID, recipe := range r.recipes {
			recipes[recipeID] = recipe
		}
		return recipes, nil
	}
	recipes := make(map[string]*model.MetaRecipe, len(recipeIDs))
	for _, recipeID := range recipeIDs {
		if recipe := r.recipes[recipeID]; recipe != nil {
			recipes[recipeID] = recipe
		}
	}
	return recipes, nil
}
