package model

import "time"

// Breezemoon 是明月清风短动态的领域事实。
type Breezemoon struct {
	ID        int64
	Content   string
	AuthorID  int64
	City      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
