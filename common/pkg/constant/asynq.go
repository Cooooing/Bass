package constant

type TaskName string

func (t TaskName) String() string {
	return string(t)
}
