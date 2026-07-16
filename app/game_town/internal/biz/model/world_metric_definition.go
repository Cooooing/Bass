package model

import "time"

type WorldMetricDefinition struct {
	ID           int64
	WorldID      int64
	Key          string
	Name         string
	Description  string
	MinValue     int32
	MaxValue     int32
	InitialValue int32
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
