package model

import "time"

// ActionQueueItem 是玩家提交到行动队列里的一个行动计划。
type ActionQueueItem struct {
	ID        string
	ActionID  string
	Times     int64
	CreatedAt time.Time
}

// ActionQueue 保存玩家行动计划的有序列表，第一项表示当前被时间轮调度的行动。
type ActionQueue struct {
	CharacterID int64
	Items       []*ActionQueueItem
}

// ActionTask 是需要放入时间轮的运行时任务。
type ActionTask struct {
	TaskID      string
	CharacterID int64
	ActionID    string
	DueAt       time.Time
}
