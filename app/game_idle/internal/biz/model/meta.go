package model

import (
	"game_idle/internal/enum"
	"time"
)

// MetaItem 是游戏里所有资产的统一配置，包括货币、资源、装备和消耗品。
type MetaItem struct {
	ID          string
	Name        string
	Type        enum.ItemType
	Description string
	Enabled     bool
	Sort        int32
}

// MetaRegion 是前端展示区域配置，不参与行动调度。
type MetaRegion struct {
	ID          string
	Name        string
	Description string
	ActionKind  enum.ActionKind
	Enabled     bool
	Sort        int32
}

// MetaAction 是队列中可执行的行动配置；区域只作为展示分类，不参与核心调度语义。
type MetaAction struct {
	ID                   string
	Name                 string
	Description          string
	RegionID             string
	ActionKind           enum.ActionKind
	AbilityID            string
	RequiredAbilityLevel int32
	Recipes              []*MetaActionRecipe
	Duration             time.Duration
	ExpReward            int64
	Enabled              bool
	Sort                 int32
}

// MetaActionRecipe 是行动与配方的绑定关系，允许一个行动顺序执行多个配方。
type MetaActionRecipe struct {
	ID       string
	ActionID string
	RecipeID string
	Enabled  bool
	Sort     int32
}

// MetaActionDetail 是行动详情，包含行动绑定配方和展示物品信息。
type MetaActionDetail struct {
	Action  *MetaAction
	Recipes []*MetaRecipe
	Items   map[string]*MetaItem
}

// MetaRecipe 是可配置行为的核心规则，既能表达无消耗采集，也能表达有消耗制造。
type MetaRecipe struct {
	ID              string
	Name            string
	Description     string
	Type            enum.RecipeType
	GenerationTimes int32
	TotalWeight     int64
	Enabled         bool
	Inputs          []*MetaRecipeInput
	Outputs         []*MetaRecipeOutput
}

// MetaRecipeInput 是配方运行前必须消耗的物品。
type MetaRecipeInput struct {
	ID       string
	RecipeID string
	ItemID   string
	Quantity int64
	Sort     int32
}

// MetaRecipeOutput 是配方运行后的产出候选项；同一配方每次生成只会按权重命中一个候选。
type MetaRecipeOutput struct {
	ID          string
	RecipeID    string
	ItemID      string
	MinQuantity int64
	MaxQuantity int64
	Weight      int32
	WeightLimit int64
	Sort        int32
}
