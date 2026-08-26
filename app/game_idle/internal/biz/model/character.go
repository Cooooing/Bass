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
	DeletedAt           *time.Time
}
