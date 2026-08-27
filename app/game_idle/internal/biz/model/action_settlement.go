package model

// ActionSettlement 是一次行动完成后的原子结算结果。
type ActionSettlement struct {
	ItemChanges      []*ActionCompletedItemChange
	AbilityChanges   []*ActionCompletedAbilityChange
	AbilityLeveledUp *AbilityLeveledUpEvent
}
