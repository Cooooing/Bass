package model

import "time"

// Moonbreeze 是清风明月短动态的领域事实。
type Moonbreeze struct {
	ID        int64
	Content   string
	AuthorID  int64
	City      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
