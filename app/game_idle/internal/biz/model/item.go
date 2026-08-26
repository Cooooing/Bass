package model

import "game_idle/internal/enum"

// Item 是游戏里所有资产的统一配置，包括货币、资源、装备和消耗品。
type Item struct {
	ID          string
	Name        string
	Type        enum.ItemType
	Description string
	Enabled     bool
	Sort        int32
}
