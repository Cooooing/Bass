package model

import (
	"game_idle/internal/enum"
	"time"
)

// Character 是玩家在挂机游戏里的角色。
type Character struct {
	ID                  int64
	UserID              int64
	Slot                int32
	Name                string
	NameKey             string
	ActionQueueCapacity int32
	MaxOfflineDuration  time.Duration
	Status              enum.CharacterStatus
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
	LastOfflineAt       *time.Time
	DeletedAt           *time.Time
}

// CharacterSession 是角色在线会话快照。
type CharacterSession struct {
	CharacterID int64
	SessionID   string
	ExpiresIn   time.Duration
}

// CharacterCloseSessionEvent 表示需要关闭指定在线会话。
type CharacterCloseSessionEvent struct {
	SessionID       string
	Reason          enum.CharacterCloseSessionReason
	Message         string
	ShouldReconnect bool
}

// CharacterAbility 是角色能力等级和累计经验快照。
type CharacterAbility struct {
	ID           int64
	CharacterID  int64
	AbilityID    enum.Ability
	Level        int32
	Exp          int64
	NextLevelExp int64
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

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

// CharacterBackpackItemChange 表示一次背包变更中的单个物品数量。
type CharacterBackpackItemChange struct {
	ItemID   string
	Quantity int64
}

// CharacterActionQueueItem 是玩家提交到行动队列里的一个行动计划。
type CharacterActionQueueItem struct {
	ID        string
	ActionID  string
	Times     int64
	CreatedAt time.Time
}

// CharacterActionQueue 保存玩家行动计划的有序列表，第一项表示当前被时间轮调度的行动。
type CharacterActionQueue struct {
	CharacterID int64
	Items       []*CharacterActionQueueItem
}

// CharacterActionTask 是需要放入时间轮的运行时任务。
type CharacterActionTask struct {
	TaskID      string
	CharacterID int64
	ActionID    string
	DueAt       time.Time
}
