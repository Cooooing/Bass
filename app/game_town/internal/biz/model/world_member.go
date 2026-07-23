package model

import (
	"time"

	"game_town/internal/enum"
)

type WorldMember struct {
	ID                  int64
	WorldID             int64
	PlayerID            int64
	CurrentLocationID   int64
	Role                enum.WorldMemberRole
	CharacterPreference string
	CharacterName       string
	CharacterBackground string
	CharacterGoal       string
	CharacterTraits     []string
	CharacterReady      bool
	JoinedAt            time.Time
	LastSeenAt          time.Time
	CreatedAt           *time.Time
	UpdatedAt           *time.Time
}

type CharacterDraft struct {
	Name       string   `json:"name"`
	Background string   `json:"background"`
	Goal       string   `json:"goal"`
	Traits     []string `json:"traits"`
}
