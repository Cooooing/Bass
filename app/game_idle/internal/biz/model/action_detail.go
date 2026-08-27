package model

// ActionDetail 是行动详情，包含行动绑定配方和展示物品信息。
type ActionDetail struct {
	Action  *Action
	Recipes []*Recipe
	Items   map[string]*Item
}
