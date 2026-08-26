package model

import "game_idle/internal/enum"

// Recipe 是可配置行为的核心规则，既能表达无消耗采集，也能表达有消耗制造。
type Recipe struct {
	ID              string
	Name            string
	Description     string
	Type            enum.RecipeType
	GenerationTimes int32
	TotalWeight     int64
	Enabled         bool
	Inputs          []*RecipeInput
	Outputs         []*RecipeOutput
}

// RecipeInput 是配方运行前必须消耗的物品。
type RecipeInput struct {
	ID       string
	RecipeID string
	ItemID   string
	Quantity int64
	Sort     int32
}

// RecipeOutput 是配方运行后的产出候选项；同一配方每次生成只会按权重命中一个候选。
type RecipeOutput struct {
	ID          string
	RecipeID    string
	ItemID      string
	MinQuantity int64
	MaxQuantity int64
	Weight      int32
	WeightLimit int64
	Sort        int32
}
