package model

import (
	"time"

	"game_town/internal/enum"
)

type Player struct {
	ID          int64
	Name        string
	DisplayName string
	Status      enum.PlayerStatus
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
