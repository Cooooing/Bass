package model

import "time"

type Location struct {
	ID          int64
	WorldID     int64
	Code        string
	Name        string
	Description string
	Tags        map[string]any
	Sort        int32
	Enabled     bool
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
