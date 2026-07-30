package model

import (
	"time"

	"content/internal/enum"
)

type Tag struct {
	ID           int64
	Code         string
	Name         string
	Description  *string
	DomainID     *int64
	Icon         *string
	Sort         int32
	ArticleCount int32
	Status       enum.TagStatus
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	CreatedBy    *int64
	UpdatedBy    *int64
}
