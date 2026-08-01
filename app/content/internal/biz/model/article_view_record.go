package model

import "time"

type ArticleViewRecord struct {
	ID                 int64
	ArticleID          int64
	UserID             int64
	IP                 *string
	UserAgent          *string
	BrowserFingerprint *string
	ViewedAt           *time.Time
	CreatedAt          *time.Time
	UpdatedAt          *time.Time
}
