package model

import "time"

// ActionCompletedEvent 表示一次行动结算成功后产生的前端同步事件。
type ActionCompletedEvent struct {
	CharacterID    int64
	ActionID       string
	TimesFinished  int64
	TimesRemaining int64
	StartedAt      time.Time
	CompletedAt    time.Time
	ItemChanges    []*ActionCompletedItemChange
}

// ActionCompletedItemChange 表示行动结算带来的单个物品变化。
type ActionCompletedItemChange struct {
	ItemID        string
	QuantityDelta int64
	QuantityAfter int64
}
