package model

import (
	"time"

	"content/internal/enum"
)

type Domain struct {
	ID          int64
	Name        string
	Description *string
	Status      enum.DomainStatus
	URL         *string
	Icon        *string
	IsNav       bool
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	CreatedBy   *int64
	UpdatedBy   *int64

	Tags []*Tag
}
