package model

import "time"

type TaskVersion struct {
	ID             int64
	TaskID         int64
	Version        int64
	Name           string
	Title          string
	Description    string
	Enabled        bool
	CronSpec       string
	Payload        string
	TimeoutSeconds int32
	AllowOverlap   bool
	AlertEnabled   bool
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}
