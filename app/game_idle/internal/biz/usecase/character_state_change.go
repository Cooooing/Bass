package usecase

import (
	"game_idle/internal/biz/model"
	"time"
)

// CharacterActionQueueChangedStateEvent 表示角色行动队列发生变化。
type CharacterActionQueueChangedStateEvent struct {
	CharacterID int64
	// Queue 是变更后的完整队列快照，BFF 可直接转发给前端覆盖本地状态。
	Queue *model.CharacterActionQueue
	// OldHeadID 和 NewHeadID 用于判断是否需要停止旧时间轮任务并启动新队首。
	OldHeadID string
	NewHeadID string
	ChangedAt time.Time
}

func (e *CharacterActionQueueChangedStateEvent) HeadChanged() bool {
	return e != nil && e.OldHeadID != e.NewHeadID
}

// CharacterStateChangeSet 是一次命令产生的状态变化和领域事件集合。
type CharacterStateChangeSet struct {
	CharacterID int64
	// QueueCommandApplied 为 false 表示时间轮任务已经过期，调用方不应继续发奖励或推进队列。
	QueueCommandApplied  bool
	QueueChanged         *CharacterActionQueueChangedStateEvent
	ActionTimesRemaining int64
	// ItemChanges 和 AbilityChanges 用于构造行动完成事件，通知前端做增量刷新。
	ItemChanges    []*model.ActionCompletedItemChange
	AbilityChanges []*model.ActionCompletedAbilityChange
	// AbilityLeveledUp 是由经验结算派生出来的领域事件，可独立推送给客户端。
	AbilityLeveledUp []*model.AbilityLeveledUpEvent
}
