package usecase

import (
	"context"
	"game_idle/internal/biz/model"
	"time"
)

// StateEventType 表示状态机产生的领域事件类型。
type StateEventType string

const (
	StateEventTypeActionQueueChanged StateEventType = "action_queue.changed" // 行动队列变化
	StateEventTypeItemsChanged       StateEventType = "items.changed"        // 物品变化
	StateEventTypeAbilityExpGained   StateEventType = "ability.exp_gained"   // 能力经验变化
	StateEventTypeAbilityLeveledUp   StateEventType = "ability.leveled_up"   // 能力升级
)

// StateEventHandler 处理领域事件，并可继续产生新的领域事件。
type StateEventHandler interface {
	Handle(ctx context.Context, event *StateEvent) ([]*StateEvent, error)
}

// StateEvent 是状态机内部事件，不直接等同于 NATS 对外事件。
type StateEvent struct {
	Type               StateEventType
	ActionQueueChanged *ActionQueueChangedStateEvent
	ItemsChanged       *ItemsChangedStateEvent
	AbilityExpGained   *AbilityExpGainedStateEvent
	AbilityLeveledUp   *model.AbilityLeveledUpEvent
}

// ActionQueueChangedStateEvent 表示角色行动队列发生变化。
type ActionQueueChangedStateEvent struct {
	CharacterID int64
	Queue       *model.ActionQueue
	OldHeadID   string
	NewHeadID   string
	ChangedAt   time.Time
}

func (e *ActionQueueChangedStateEvent) HeadChanged() bool {
	return e != nil && e.OldHeadID != e.NewHeadID
}

// ItemsChangedStateEvent 表示角色背包物品发生变化。
type ItemsChangedStateEvent struct {
	CharacterID int64
	Changes     []*model.ActionCompletedItemChange
}

// AbilityExpGainedStateEvent 表示角色能力经验发生变化。
type AbilityExpGainedStateEvent struct {
	CharacterID int64
	Changes     []*model.ActionCompletedAbilityChange
}

// StateChangeSet 是一次命令产生的状态变化和领域事件集合。
type StateChangeSet struct {
	CharacterID          int64
	QueueCommandApplied  bool
	QueueChanged         *ActionQueueChangedStateEvent
	ActionTimesRemaining int64
	ItemChanges          []*model.ActionCompletedItemChange
	AbilityChanges       []*model.ActionCompletedAbilityChange
	Events               []*StateEvent
}

func (s *StateChangeSet) AbilityLeveledUpEvents() []*model.AbilityLeveledUpEvent {
	events := make([]*model.AbilityLeveledUpEvent, 0)
	for _, event := range s.Events {
		if event.AbilityLeveledUp != nil {
			events = append(events, event.AbilityLeveledUp)
		}
	}
	return events
}
