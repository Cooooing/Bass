package task

type TaskName string

const (
	TaskNodePing TaskName = "node:ping"
	TaskNodePow  TaskName = "node:pow"
)

func (t TaskName) String() string {
	return string(t)
}
