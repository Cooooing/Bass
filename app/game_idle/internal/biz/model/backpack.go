package model

import "time"

// CharacterItem 是角色背包里的物品余额快照。
type CharacterItem struct {
	ID            int64
	CharacterID   int64
	ItemID        string
	Quantity      int64
	TotalObtained int64
	TotalConsumed int64
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}

// BackpackItemChange 表示一次背包变更中的单个物品数量。
type BackpackItemChange struct {
	ItemID   string
	Quantity int64
}
