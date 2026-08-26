package repo

import (
	commonclient "common/pkg/client"
	"context"
	"encoding/json"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/data/gen"
	recipeent "game_idle/internal/data/gen/recipe"
	recipeinputent "game_idle/internal/data/gen/recipeinput"
	recipeoutputent "game_idle/internal/data/gen/recipeoutput"
	"game_idle/internal/enum"
	"sync"

	"github.com/redis/go-redis/v9"
)

var _ bizrepo.RecipeRepo = (*RecipeRepo)(nil)

type RecipeRepo struct {
	mutex           sync.RWMutex
	db              *gen.Client
	redisClient     *commonclient.RedisClient
	recipes         map[string]*model.Recipe
	loaded          bool
	recipesRedisKey string
	loadedRedisKey  string
}

func NewRecipeRepo(db *gen.Client, redisClient *commonclient.RedisClient) (bizrepo.RecipeRepo, error) {
	repo := &RecipeRepo{
		db:              db,
		redisClient:     redisClient,
		recipes:         make(map[string]*model.Recipe),
		recipesRedisKey: "game_idle:config:recipes",
		loadedRedisKey:  "game_idle:config:recipes:loaded",
	}
	if _, err := repo.Refresh(context.Background()); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *RecipeRepo) Refresh(ctx context.Context) ([]*model.Recipe, error) {
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
		return nil, err
	}
	recipes := make([]*model.Recipe, 0, len(rows))
	recipeMap := make(map[string]*model.Recipe, len(rows))
	values := make(map[string]any, len(rows))
	for _, row := range rows {
		recipe := &model.Recipe{
			ID:              row.ID,
			Name:            row.Name,
			Description:     row.Description,
			Type:            enum.RecipeType(row.Type),
			GenerationTimes: row.GenerationTimes,
			Enabled:         row.Enabled,
			Inputs:          make([]*model.RecipeInput, 0, len(row.Edges.Inputs)),
			Outputs:         make([]*model.RecipeOutput, 0, len(row.Edges.Outputs)),
		}
		if recipe.ID == "" || recipe.GenerationTimes <= 0 || len(row.Edges.Outputs) == 0 {
			return nil, model.ErrRecipeInvalid
		}
		for _, input := range row.Edges.Inputs {
			if input.ItemID == "" || input.Quantity <= 0 {
				return nil, model.ErrRecipeInvalid
			}
			recipe.Inputs = append(recipe.Inputs, &model.RecipeInput{
				ID:       input.ID,
				RecipeID: input.RecipeID,
				ItemID:   input.ItemID,
				Quantity: input.Quantity,
				Sort:     input.Sort,
			})
		}
		for _, output := range row.Edges.Outputs {
			if output.ItemID == "" || output.MinQuantity <= 0 || output.MaxQuantity < output.MinQuantity || output.Weight <= 0 {
				return nil, model.ErrRecipeInvalid
			}
			recipe.TotalWeight += int64(output.Weight)
			recipe.Outputs = append(recipe.Outputs, &model.RecipeOutput{
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
			return nil, model.ErrRecipeInvalid
		}
		data, err := json.Marshal(recipe)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
		recipeMap[recipe.ID] = recipe
		values[recipe.ID] = data
	}
	_, err = r.redisClient.Client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Del(ctx, r.recipesRedisKey)
		if len(values) > 0 {
			pipe.HSet(ctx, r.recipesRedisKey, values)
		}
		pipe.Set(ctx, r.loadedRedisKey, "1", 0)
		return nil
	})
	if err != nil {
		return nil, err
	}
	r.mutex.Lock()
	r.recipes = recipeMap
	r.loaded = true
	r.mutex.Unlock()
	return recipes, err
}

func (r *RecipeRepo) Get(ctx context.Context, recipeID string) (*model.Recipe, error) {
	r.mutex.RLock()
	recipe, ok := r.recipes[recipeID]
	loaded := r.loaded
	r.mutex.RUnlock()
	if ok {
		return recipe, nil
	}
	if loaded {
		return nil, model.ErrRecipeInvalid
	}
	if _, err := r.Refresh(ctx); err != nil {
		return nil, err
	}
	r.mutex.RLock()
	recipe = r.recipes[recipeID]
	r.mutex.RUnlock()
	if recipe == nil {
		return nil, model.ErrRecipeInvalid
	}
	return recipe, nil
}

func (r *RecipeRepo) Map(ctx context.Context, recipeIDs []string) (map[string]*model.Recipe, error) {
	r.mutex.RLock()
	loaded := r.loaded
	r.mutex.RUnlock()
	if !loaded {
		if _, err := r.Refresh(ctx); err != nil {
			return nil, err
		}
	}
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	recipes := make(map[string]*model.Recipe)
	if len(recipeIDs) == 0 {
		for recipeID, recipe := range r.recipes {
			recipes[recipeID] = recipe
		}
		return recipes, nil
	}
	for _, recipeID := range recipeIDs {
		if recipe, ok := r.recipes[recipeID]; ok {
			recipes[recipeID] = recipe
		}
	}
	return recipes, nil
}
