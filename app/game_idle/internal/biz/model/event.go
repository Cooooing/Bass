package model

// GameIdleEvent 是挂机游戏统一事件载体。
type GameIdleEvent struct {
	ChatMessage      *ChatMessage
	CloseSession     *CharacterCloseSessionEvent
	ActionCompleted  *ActionCompletedEvent
	AbilityLeveledUp *AbilityLeveledUpEvent
}
