package model

// GameConfig 是前端低频配置快照。
type GameConfig struct {
	Regions    []*RegionConfig `json:"regions"`
	Actions    []*ActionConfig `json:"actions"`
	ServerTime int64           `json:"server_time"`
}

type RegionConfig struct {
	RegionID    string `json:"region_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Sort        int32  `json:"sort"`
}

type ActionConfig struct {
	ActionID             string `json:"action_id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	RegionID             string `json:"region_id"`
	ActionKind           string `json:"action_kind"`
	AbilityID            string `json:"ability_id"`
	RequiredAbilityLevel int32  `json:"required_ability_level"`
	DurationSeconds      int64  `json:"duration_seconds"`
	ExpReward            int64  `json:"exp_reward"`
	Enabled              bool   `json:"enabled"`
	Sort                 int32  `json:"sort"`
}

// ActionDetailConfig 是前端行动详情配置。
type ActionDetailConfig struct {
	Action  *ActionConfig         `json:"action"`
	Recipes []*ActionRecipeConfig `json:"recipes"`
}

type ActionRecipeConfig struct {
	RecipeID        string                `json:"recipe_id"`
	Name            string                `json:"name"`
	Description     string                `json:"description"`
	RecipeType      string                `json:"recipe_type"`
	GenerationTimes int32                 `json:"generation_times"`
	Inputs          []*RecipeInputConfig  `json:"inputs"`
	Outputs         []*RecipeOutputConfig `json:"outputs"`
}

type RecipeInputConfig struct {
	ItemID   string `json:"item_id"`
	ItemName string `json:"item_name"`
	ItemType string `json:"item_type"`
	Quantity int64  `json:"quantity"`
}

type RecipeOutputConfig struct {
	ItemID      string  `json:"item_id"`
	ItemName    string  `json:"item_name"`
	ItemType    string  `json:"item_type"`
	MinQuantity int64   `json:"min_quantity"`
	MaxQuantity int64   `json:"max_quantity"`
	Weight      int32   `json:"weight"`
	Probability float64 `json:"probability"`
}
