package task

type TaskName string

const (
	TaskNodePing    TaskName = "node:ping"
	TaskNodePow     TaskName = "node:pow"
	TaskNodeSession TaskName = "node:session"
)

func (t TaskName) String() string {
	return string(t)
}
