package model

import (
	"time"

	"content/internal/enum"
)

type Domain struct {
	ID          int64
	Code        string
	Name        string
	Description *string
	Status      enum.DomainStatus
	URL         *string
	Icon        *string
	IsNav       bool
	Sort        int32
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	CreatedBy   *int64
	UpdatedBy   *int64
}
