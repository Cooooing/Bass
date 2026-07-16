package model

import "time"

type AgentRun struct {
	ID            int64
	WorldID       *int64
	AgentConfigID *int64
	RunType       string
	CommandID     *int64
	EventID       *int64
	NpcID         *int64
	Model         string
	InputJSON     map[string]any
	OutputJSON    map[string]any
	Status        string
	ErrorSummary  string
	LatencyMS     int64
	CreatedAt     time.Time
}
