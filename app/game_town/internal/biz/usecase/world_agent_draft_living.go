package usecase

import (
	"context"
	"fmt"
	"strings"

	"game_town/internal/biz/model"
	"game_town/internal/enum"
)

func (r *WorldAgentRunner) saveDraftFactions(
	ctx context.Context,
	worldID int64,
	draft *model.WorldDraft,
) (map[string]*model.Faction, error) {
	out := make(map[string]*model.Faction, len(draft.Factions))
	if draft == nil {
		return out, nil
	}

	for _, item := range draft.Factions {
		code := strings.TrimSpace(item.Code)
		name := normalizeModelText(item.Name, maxNpcNameRunes)
		if code == "" || len(code) > 64 || name == "" {
			return nil, fmt.Errorf("invalid faction")
		}
		if out[code] != nil {
			return nil, fmt.Errorf("duplicate faction code")
		}

		faction, err := r.factionRepo.Save(ctx, &model.Faction{
			WorldID:     worldID,
			Code:        code,
			Name:        name,
			Description: strings.TrimSpace(item.Description),
			PublicGoal:  strings.TrimSpace(item.PublicGoal),
			Status:      enum.FactionStatusActive,
			Attributes:  map[string]any{},
		})
		if err != nil {
			return nil, err
		}
		out[code] = faction
	}
	return out, nil
}

func defaultWorldRules(
	rules map[string]any,
) map[string]any {
	if rules == nil {
		rules = make(map[string]any)
	}
	if _, ok := rules["calendar"]; !ok {
		rules["calendar"] = map[string]any{"days_per_month": 30, "months_per_year": 12}
	}
	if _, ok := rules["relationship_dimensions"]; !ok {
		rules["relationship_dimensions"] = []any{"trust", "fear", "obligation"}
	}
	if _, ok := rules["actions"]; !ok {
		rules["actions"] = []any{"move", "talk", "investigate", "trade", "work", "rest"}
	}
	return rules
}
