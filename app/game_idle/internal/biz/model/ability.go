package model

import (
	"game_idle/internal/enum"
	"time"
)

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
