package task

import "context"

type DefaultSchedule struct {
	Title          string
	Description    string
	Enabled        bool
	CronSpec       string
	Payload        string
	TimeoutSeconds int32
	AllowOverlap   bool
	AlertEnabled   bool
}

type Task interface {
	Name() string
	Title() string
	Description() string
	DefaultSchedules() []*DefaultSchedule
	Execute(ctx context.Context, payload string) error
}
