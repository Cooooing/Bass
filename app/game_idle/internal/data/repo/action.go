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
	actionent "game_idle/internal/data/gen/action"
	actionrecipeent "game_idle/internal/data/gen/actionrecipe"
	"game_idle/internal/enum"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var _ bizrepo.ActionRepo = (*ActionRepo)(nil)

type ActionRepo struct {
	mutex           sync.RWMutex
	db              *gen.Client
	redisClient     *commonclient.RedisClient
	actions         map[string]*model.Action
	loaded          bool
	actionsRedisKey string
	loadedRedisKey  string
	redisTTL        time.Duration
}

func NewActionRepo(db *gen.Client, redisClient *commonclient.RedisClient) (bizrepo.ActionRepo, error) {
	repo := &ActionRepo{
		db:              db,
		redisClient:     redisClient,
		actions:         make(map[string]*model.Action),
		actionsRedisKey: "game_idle:metadata:actions",
		loadedRedisKey:  "game_idle:metadata:actions:loaded",
		redisTTL:        2 * time.Hour,
	}
	if err := repo.RefreshLocal(context.Background()); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *ActionRepo) Refresh(ctx context.Context) ([]*model.Action, error) {
	actions, actionMap, err := r.loadFromDB(ctx)
	if err != nil {
		return nil, err
	}
	if err = r.saveRedis(ctx, actionMap); err != nil {
		return nil, err
	}
	r.replaceLocal(actionMap)
	return actions, nil
}

func (r *ActionRepo) RefreshLocal(ctx context.Context) error {
	actionMap, err := r.loadFromRedis(ctx)
	if err == nil {
		r.replaceLocal(actionMap)
		return nil
	}
	_, actionMap, err = r.loadFromDB(ctx)
	if err != nil {
		return err
	}
	r.replaceLocal(actionMap)
	_ = r.saveRedis(ctx, actionMap)
	return nil
}

func (r *ActionRepo) loadFromDB(ctx context.Context) ([]*model.Action, map[string]*model.Action, error) {
	rows, err := r.db.Action.Query().
		Where(actionent.DeletedAtIsNil()).
		WithRecipes(func(query *gen.ActionRecipeQuery) {
			query.Where(
				actionrecipeent.DeletedAtIsNil(),
				actionrecipeent.EnabledEQ(true),
			).Order(actionrecipeent.BySort(), actionrecipeent.ByID())
		}).
		Order(actionent.BySort(), actionent.ByID()).
		All(ctx)
	if err != nil {
		return nil, nil, err
	}
	actions := make([]*model.Action, 0, len(rows))
	actionMap := make(map[string]*model.Action, len(rows))
	for _, row := range rows {
		regionID := ""
		if row.RegionID != nil {
			regionID = *row.RegionID
		}
		abilityID := ""
		if row.AbilityID != nil {
			abilityID = string(*row.AbilityID)
		}
		action := &model.Action{
			ID:                   row.ID,
			Name:                 row.Name,
			Description:          row.Description,
			RegionID:             regionID,
			ActionKind:           enum.ActionKind(row.ActionKind),
			AbilityID:            abilityID,
			RequiredAbilityLevel: row.RequiredAbilityLevel,
			Recipes:              make([]*model.ActionRecipe, 0, len(row.Edges.Recipes)),
			Duration:             time.Duration(row.DurationSeconds) * time.Second,
			ExpReward:            row.ExpReward,
			Enabled:              row.Enabled,
			Sort:                 row.Sort,
		}
		if action.Enabled && (action.ID == "" || action.Duration <= 0 || len(row.Edges.Recipes) == 0) {
			return nil, nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
		}
		for _, relation := range row.Edges.Recipes {
			if relation.RecipeID == "" {
				return nil, nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
			}
			action.Recipes = append(action.Recipes, &model.ActionRecipe{
				ID:       relation.ID,
				ActionID: relation.ActionID,
				RecipeID: relation.RecipeID,
				Enabled:  relation.Enabled,
				Sort:     relation.Sort,
			})
		}
		actions = append(actions, action)
		actionMap[action.ID] = action
	}
	return actions, actionMap, nil
}

func (r *ActionRepo) loadFromRedis(ctx context.Context) (map[string]*model.Action, error) {
	loaded, err := r.redisClient.Client.Exists(ctx, r.loadedRedisKey).Result()
	if err != nil {
		return nil, err
	}
	if loaded == 0 {
		return nil, redis.Nil
	}
	values, err := r.redisClient.Client.HGetAll(ctx, r.actionsRedisKey).Result()
	if err != nil {
		return nil, err
	}
	actionMap := make(map[string]*model.Action, len(values))
	for actionID, text := range values {
		action := &model.Action{}
		if err = json.Unmarshal([]byte(text), action); err != nil {
			return nil, err
		}
		actionMap[actionID] = action
	}
	return actionMap, nil
}

func (r *ActionRepo) saveRedis(ctx context.Context, actionMap map[string]*model.Action) error {
	values := make(map[string]any, len(actionMap))
	for actionID, action := range actionMap {
		data, err := json.Marshal(action)
		if err != nil {
			return err
		}
		values[actionID] = data
	}
	existing, err := r.redisClient.Client.HKeys(ctx, r.actionsRedisKey).Result()
	if err != nil && err != redis.Nil {
		return err
	}
	stale := make([]string, 0)
	for _, actionID := range existing {
		if _, ok := actionMap[actionID]; !ok {
			stale = append(stale, actionID)
		}
	}
	_, err = r.redisClient.Client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		if len(values) > 0 {
			pipe.HSet(ctx, r.actionsRedisKey, values)
		}
		if len(stale) > 0 {
			pipe.HDel(ctx, r.actionsRedisKey, stale...)
		}
		pipe.Expire(ctx, r.actionsRedisKey, r.redisTTL)
		pipe.Set(ctx, r.loadedRedisKey, "1", r.redisTTL)
		return nil
	})
	return err
}

func (r *ActionRepo) replaceLocal(actionMap map[string]*model.Action) {
	r.mutex.Lock()
	r.actions = actionMap
	r.loaded = true
	r.mutex.Unlock()
}

func (r *ActionRepo) Get(ctx context.Context, actionID string) (*model.Action, error) {
	r.mutex.RLock()
	action, ok := r.actions[actionID]
	loaded := r.loaded
	r.mutex.RUnlock()
	if ok {
		return action, nil
	}
	if loaded {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
	}
	if err := r.RefreshLocal(ctx); err != nil {
		return nil, err
	}
	r.mutex.RLock()
	action = r.actions[actionID]
	r.mutex.RUnlock()
	if action == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
	}
	return action, nil
}

func (r *ActionRepo) Map(ctx context.Context, actionIDs []string) (map[string]*model.Action, error) {
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
	actions := make(map[string]*model.Action)
	if len(actionIDs) == 0 {
		for actionID, action := range r.actions {
			actions[actionID] = action
		}
		return actions, nil
	}
	for _, actionID := range actionIDs {
		if action, ok := r.actions[actionID]; ok {
			actions[actionID] = action
		}
	}
	return actions, nil
}
