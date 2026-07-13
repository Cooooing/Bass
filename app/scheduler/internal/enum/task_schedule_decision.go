package enum

type TaskScheduleDecision string

const (
	TaskScheduleDecisionSkip    TaskScheduleDecision = "skip"
	TaskScheduleDecisionRun     TaskScheduleDecision = "run"
	TaskScheduleDecisionOverlap TaskScheduleDecision = "overlap"
)
