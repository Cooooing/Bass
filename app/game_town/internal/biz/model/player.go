package model

import "time"

type Player struct {
	ID          int64
	Name        string
	DisplayName string
	Status      string
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}
