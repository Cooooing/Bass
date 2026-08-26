package model

import (
	"game_idle/internal/enum"
	"time"
)

// Action 是队列中可执行的行动配置；区域只作为展示分类，不参与核心调度语义。
type Action struct {
	ID                   string
	Name                 string
	Description          string
	RegionID             string
	ActionKind           enum.ActionKind
	AbilityID            string
	RequiredAbilityLevel int32
	Recipes              []*ActionRecipe
	Duration             time.Duration
	ExpReward            int64
	Enabled              bool
}

// ActionRecipe 是行动与配方的绑定关系，允许一个行动顺序执行多个配方。
type ActionRecipe struct {
	ID       string
	ActionID string
	RecipeID string
	Enabled  bool
	Sort     int32
}
