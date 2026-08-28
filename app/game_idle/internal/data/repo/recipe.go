package repo

import (
	"common/pkg/apperror"
	commonclient "common/pkg/client"
	cerrors "common/proto/gen/common/errors"
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
	"time"

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
	redisTTL        time.Duration
}

func NewRecipeRepo(db *gen.Client, redisClient *commonclient.RedisClient) (bizrepo.RecipeRepo, error) {
	repo := &RecipeRepo{
		db:              db,
		redisClient:     redisClient,
		recipes:         make(map[string]*model.Recipe),
		recipesRedisKey: "game_idle:metadata:recipes",
		loadedRedisKey:  "game_idle:metadata:recipes:loaded",
		redisTTL:        2 * time.Hour,
	}
	if err := repo.RefreshLocal(context.Background()); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *RecipeRepo) Refresh(ctx context.Context) ([]*model.Recipe, error) {
	recipes, recipeMap, err := r.loadFromDB(ctx)
	if err != nil {
		return nil, err
	}
	if err = r.saveRedis(ctx, recipeMap); err != nil {
		return nil, err
	}
	r.replaceLocal(recipeMap)
	return recipes, nil
}

func (r *RecipeRepo) RefreshLocal(ctx context.Context) error {
	recipeMap, err := r.loadFromRedis(ctx)
	if err == nil {
		r.replaceLocal(recipeMap)
		return nil
	}
	_, recipeMap, err = r.loadFromDB(ctx)
	if err != nil {
		return err
	}
	r.replaceLocal(recipeMap)
	_ = r.saveRedis(ctx, recipeMap)
	return nil
}

func (r *RecipeRepo) loadFromDB(ctx context.Context) ([]*model.Recipe, map[string]*model.Recipe, error) {
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
		return nil, nil, err
	}
	recipes := make([]*model.Recipe, 0, len(rows))
	recipeMap := make(map[string]*model.Recipe, len(rows))
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
			return nil, nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_RECIPE_INVALID)
		}
		for _, input := range row.Edges.Inputs {
			if input.ItemID == "" || input.Quantity <= 0 {
				return nil, nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_RECIPE_INVALID)
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
				return nil, nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_RECIPE_INVALID)
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
			return nil, nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_RECIPE_INVALID)
		}
		recipes = append(recipes, recipe)
		recipeMap[recipe.ID] = recipe
	}
	return recipes, recipeMap, nil
}

func (r *RecipeRepo) loadFromRedis(ctx context.Context) (map[string]*model.Recipe, error) {
	loaded, err := r.redisClient.Client.Exists(ctx, r.loadedRedisKey).Result()
	if err != nil {
		return nil, err
	}
	if loaded == 0 {
		return nil, redis.Nil
	}
	values, err := r.redisClient.Client.HGetAll(ctx, r.recipesRedisKey).Result()
	if err != nil {
		return nil, err
	}
	recipes := make(map[string]*model.Recipe, len(values))
	for recipeID, text := range values {
		recipe := &model.Recipe{}
		if err = json.Unmarshal([]byte(text), recipe); err != nil {
			return nil, err
		}
		recipes[recipeID] = recipe
	}
	return recipes, nil
}

func (r *RecipeRepo) saveRedis(ctx context.Context, recipeMap map[string]*model.Recipe) error {
	values := make(map[string]any, len(recipeMap))
	for recipeID, recipe := range recipeMap {
		data, err := json.Marshal(recipe)
		if err != nil {
			return err
		}
		values[recipeID] = data
	}
	existing, err := r.redisClient.Client.HKeys(ctx, r.recipesRedisKey).Result()
	if err != nil && err != redis.Nil {
		return err
	}
	stale := make([]string, 0)
	for _, recipeID := range existing {
		if _, ok := recipeMap[recipeID]; !ok {
			stale = append(stale, recipeID)
		}
	}
	_, err = r.redisClient.Client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		if len(values) > 0 {
			pipe.HSet(ctx, r.recipesRedisKey, values)
		}
		if len(stale) > 0 {
			pipe.HDel(ctx, r.recipesRedisKey, stale...)
		}
		pipe.Expire(ctx, r.recipesRedisKey, r.redisTTL)
		pipe.Set(ctx, r.loadedRedisKey, "1", r.redisTTL)
		return nil
	})
	return err
}

func (r *RecipeRepo) replaceLocal(recipeMap map[string]*model.Recipe) {
	r.mutex.Lock()
	r.recipes = recipeMap
	r.loaded = true
	r.mutex.Unlock()
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
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_RECIPE_INVALID)
	}
	if err := r.RefreshLocal(ctx); err != nil {
		return nil, err
	}
	r.mutex.RLock()
	recipe = r.recipes[recipeID]
	r.mutex.RUnlock()
	if recipe == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_RECIPE_INVALID)
	}
	return recipe, nil
}

func (r *RecipeRepo) Map(ctx context.Context, recipeIDs []string) (map[string]*model.Recipe, error) {
	r.mutex.RLock()
	loaded := r.loaded
	r.mutex.RUnlock()
	if !loaded {
		if err := r.RefreshLocal(ctx); err != nil {
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
