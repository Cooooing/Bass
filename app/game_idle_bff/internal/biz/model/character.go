package model

import (
	gameidleenum "common/proto/gen/game_idle/v1/enum"
	"time"
)

type Character struct {
	ID                  int64
	Name                string
	Status              gameidleenum.CharacterStatus
	Slot                int32
	ActionQueueCapacity int32
	MaxOfflineSeconds   int64
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
}
