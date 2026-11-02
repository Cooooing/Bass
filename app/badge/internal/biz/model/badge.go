package model

import "time"

// Badge is the service-owned reusable medal.
type Badge struct {
	ID             int64
	Code           *string
	Name           string
	Description    string
	RenderParams   []byte
	DesignerUserID int64
	Status         string
	PublishedAt    *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
