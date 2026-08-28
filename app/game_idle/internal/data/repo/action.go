package repo

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/data/gen"
	actionent "game_idle/internal/data/gen/action"
	actionrecipeent "game_idle/internal/data/gen/actionrecipe"
	"game_idle/internal/enum"
	"sync"
	"time"
)

var _ bizrepo.ActionRepo = (*ActionRepo)(nil)

type ActionRepo struct {
	mutex   sync.RWMutex
	db      *gen.Client
	actions map[string]*model.Action
	loaded  bool
}

func NewActionRepo(db *gen.Client) (bizrepo.ActionRepo, error) {
	repo := &ActionRepo{
		db:      db,
		actions: make(map[string]*model.Action),
	}
	if _, err := repo.Refresh(context.Background()); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *ActionRepo) Refresh(ctx context.Context) ([]*model.Action, error) {
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
		return nil, err
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
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
		}
		for _, relation := range row.Edges.Recipes {
			if relation.RecipeID == "" {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_IDLE_ACTION_INVALID)
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
	r.mutex.Lock()
	r.actions = actionMap
	r.loaded = true
	r.mutex.Unlock()
	return actions, nil
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
	if _, err := r.Refresh(ctx); err != nil {
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
		if _, err := r.Refresh(ctx); err != nil {
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
