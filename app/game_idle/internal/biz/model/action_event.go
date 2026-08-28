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
	AbilityChanges []*ActionCompletedAbilityChange
}

// ActionCompletedItemChange 表示行动结算带来的单个物品变化。
type ActionCompletedItemChange struct {
	ItemID        string
	QuantityDelta int64
	QuantityAfter int64
}

// ActionCompletedAbilityChange 表示行动结算带来的单个能力经验变化。
type ActionCompletedAbilityChange struct {
	AbilityID string
	ExpDelta  int64
	ExpAfter  int64
}

// AbilityLeveledUpEvent 表示行动结算触发的能力升级事件。
type AbilityLeveledUpEvent struct {
	CharacterID  int64
	AbilityID    string
	Level        int32
	Exp          int64
	NextLevelExp int64
}

// ActionQueueUpdatedEvent 表示行动队列已发生变化，需要同步前端快照。
type ActionQueueUpdatedEvent struct {
	CharacterID int64
	Items       []*ActionQueueItem
	Reason      string
	UpdatedAt   time.Time
}
