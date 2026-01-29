package constant

type TaskName string

const (
	TaskSignalNodePing    TaskName = "signal:node:ping"
	TaskSignalNodePow     TaskName = "signal:node:pow"
	TaskSignalNodeSession TaskName = "signal:node:session"

	TaskConnectorRegister TaskName = "connector:register"
)

func (t TaskName) String() string {
	return string(t)
}
