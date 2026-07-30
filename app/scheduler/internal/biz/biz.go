package biz

import (
	"scheduler/internal/biz/usecase"
	"scheduler/internal/biz/usecase/task"

	"github.com/google/wire"
)

var BizProviderSet = wire.NewSet(
	task.NewNoop,
	task.NewUserUnbanAccounts,
	task.NewUserOutboxPublishBatch,
	task.NewContentOutboxPublishBatch,
	ProvideTasks,
	usecase.NewScheduledTaskUsecase,
	usecase.NewDelayedTaskUsecase,
)

func ProvideTasks(
	noop *task.Noop,
	userUnbanAccounts *task.UserUnbanAccounts,
	userOutboxPublishBatch *task.UserOutboxPublishBatch,
	contentOutboxPublishBatch *task.ContentOutboxPublishBatch,
) map[string]task.Task {
	return map[string]task.Task{
		noop.HandlerName().String():                      noop,
		userUnbanAccounts.HandlerName().String():         userUnbanAccounts,
		userOutboxPublishBatch.HandlerName().String():    userOutboxPublishBatch,
		contentOutboxPublishBatch.HandlerName().String(): contentOutboxPublishBatch,
	}
}
