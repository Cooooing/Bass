package model

import (
	gameidleenum "common/proto/gen/game_idle/v1/enum"
	"time"
)

type Character struct {
	ID                  int64                        `json:"id"`
	Name                string                       `json:"name"`
	Status              gameidleenum.CharacterStatus `json:"status"`
	Slot                int32                        `json:"slot"`
	ActionQueueCapacity int32                        `json:"action_queue_capacity"`
	MaxOfflineSeconds   int64                        `json:"max_offline_seconds"`
	CreatedAt           *time.Time                   `json:"created_at,omitempty"`
	UpdatedAt           *time.Time                   `json:"updated_at,omitempty"`
}
