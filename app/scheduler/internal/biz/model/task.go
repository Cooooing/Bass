package model

import "time"

type Task struct {
	ID             int64
	Name           string
	Title          string
	Description    string
	Enabled        bool
	CronSpec       string
	Payload        string
	TimeoutSeconds int32
	AllowOverlap   bool
	AlertEnabled   bool
	Version        int64
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
