package task

type TaskName string

const (
	TaskNodePing TaskName = "node_ping"
	TaskNodePow  TaskName = "node_pow"
)

func (t TaskName) String() string {
	return string(t)
}
