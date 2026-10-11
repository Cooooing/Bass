package model

import (
	"time"

	"bff_bbs/internal/enum"
)

type Relation struct {
	ID        int64
	Type      enum.RelationType
	ActorID   int64
	TargetID  int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
