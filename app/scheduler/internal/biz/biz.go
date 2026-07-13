package biz

import (
	"scheduler/internal/biz/usecase"
	"scheduler/internal/biz/usecase/task"

	"github.com/google/wire"
)

// BizProviderSet 是 biz 层依赖集合。
var BizProviderSet = wire.NewSet(
	task.NewNoop,
	ProvideTasks,
	usecase.NewTaskUsecase,
	usecase.NewSchedulerRunner,
)

func ProvideTasks(noop *task.Noop) map[string]task.Task {
	tasks := map[string]task.Task{}
	for _, item := range []task.Task{noop} {
		name := item.Name()
		if name == "" {
			panic("scheduler task name is empty")
		}
		if _, ok := tasks[name]; ok {
			panic("scheduler task name duplicated: " + name)
		}
		tasks[name] = item
	}
	return tasks
}
