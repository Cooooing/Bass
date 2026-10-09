package model

import "time"

type Moonbreeze struct {
	ID        int64
	Content   string
	AuthorID  int64
	City      *string
	CreatedAt *time.Time
	Author    *AccountProfile
	ShowCity  bool
}

type PageMoonbreezesResp struct {
	Rows       []*Moonbreeze
	NextCursor *string
}
