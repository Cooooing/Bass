package biz

import (
	"scheduler/internal/biz/usecase"
	"scheduler/internal/biz/usecase/task"

	"github.com/google/wire"
)

var BizProviderSet = wire.NewSet(
	task.NewNoop,
	task.NewUserUnbanExpired,
	ProvideTasks,
	usecase.NewTaskUsecase,
	usecase.NewDelayedTaskUsecase,
	usecase.NewSchedulerRunner,
	usecase.NewDelayedTaskRunner,
)

func ProvideTasks(noop *task.Noop, userUnbanExpired *task.UserUnbanExpired) map[string]task.Task {
	tasks := map[string]task.Task{}
	for _, item := range []task.Task{noop, userUnbanExpired} {
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
