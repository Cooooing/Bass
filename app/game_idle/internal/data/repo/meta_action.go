package repo

import (
	"context"
	"fmt"
	"game_idle/internal/biz/model"
	bizrepo "game_idle/internal/biz/repo"
	"game_idle/internal/data/gen"
	actionent "game_idle/internal/data/gen/action"
	actionrecipeent "game_idle/internal/data/gen/actionrecipe"
	"game_idle/internal/enum"
	"sync"
	"time"
)

var _ bizrepo.MetaActionRepo = (*MetaActionRepo)(nil)

type MetaActionRepo struct {
	mutex   sync.RWMutex
	db      *gen.Client
	actions map[string]*model.MetaAction
}

func NewMetaActionRepo(db *gen.Client) (bizrepo.MetaActionRepo, error) {
	repo := &MetaActionRepo{
		db: db,
	}
	if err := repo.Refresh(context.Background()); err != nil {
		return nil, err
	}
	return repo, nil
}

func (r *MetaActionRepo) Refresh(ctx context.Context) error {
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
		return err
	}
	actions := make(map[string]*model.MetaAction, len(rows))
	for _, row := range rows {
		regionID := ""
		if row.RegionID != nil {
			regionID = *row.RegionID
		}
		abilityID := ""
		if row.AbilityID != nil {
			abilityID = string(*row.AbilityID)
		}
		action := &model.MetaAction{
			ID:                   row.ID,
			Name:                 row.Name,
			Description:          row.Description,
			RegionID:             regionID,
			ActionKind:           enum.ActionKind(row.ActionKind),
			AbilityID:            abilityID,
			RequiredAbilityLevel: row.RequiredAbilityLevel,
			Recipes:              make([]*model.MetaActionRecipe, 0, len(row.Edges.Recipes)),
			Duration:             time.Duration(row.DurationSeconds) * time.Second,
			ExpReward:            row.ExpReward,
			Enabled:              row.Enabled,
			Sort:                 row.Sort,
		}
		if action.Enabled && (action.ID == "" || action.Duration <= 0 || len(row.Edges.Recipes) == 0) {
			return fmt.Errorf("game idle action config invalid: %s", action.ID)
		}
		for _, relation := range row.Edges.Recipes {
			if relation.RecipeID == "" {
				return fmt.Errorf("game idle action recipe config invalid: %s", action.ID)
			}
			action.Recipes = append(action.Recipes, &model.MetaActionRecipe{
				ID:       relation.ID,
				ActionID: relation.ActionID,
				RecipeID: relation.RecipeID,
				Enabled:  relation.Enabled,
				Sort:     relation.Sort,
			})
		}
		actions[action.ID] = action
	}
	r.mutex.Lock()
	r.actions = actions
	r.mutex.Unlock()
	return nil
}

func (r *MetaActionRepo) Get(ctx context.Context, actionID string) (*model.MetaAction, error) {
	r.mutex.RLock()
	action := r.actions[actionID]
	r.mutex.RUnlock()
	return action, nil
}

func (r *MetaActionRepo) Map(ctx context.Context, actionIDs []string) (map[string]*model.MetaAction, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	if len(actionIDs) == 0 {
		actions := make(map[string]*model.MetaAction, len(r.actions))
		for actionID, action := range r.actions {
			actions[actionID] = action
		}
		return actions, nil
	}
	actions := make(map[string]*model.MetaAction, len(actionIDs))
	for _, actionID := range actionIDs {
		if action := r.actions[actionID]; action != nil {
			actions[actionID] = action
		}
	}
	return actions, nil
}
