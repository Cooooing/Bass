package model

import (
	"time"

	"content/internal/enum"
)

type ContentModerationRecord struct {
	ID         int64
	Target     enum.ContentModerationTarget
	TargetID   int64
	Action     enum.ContentModerationAction
	ReasonCode *string
	Reason     *string
	OperatorID int64
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
}
