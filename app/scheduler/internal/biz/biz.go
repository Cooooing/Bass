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
	task.NewContentPublishScheduledArticles,
	task.NewContentFlushArticleViews,
	task.NewContentDrawArticleLotteries,
	task.NewContentFinishArticlePolls,
	task.NewPlatformRefreshHoliday,
	ProvideTasks,
	usecase.NewScheduledTaskUsecase,
	usecase.NewDelayedTaskUsecase,
)

func ProvideTasks(
	noop *task.Noop,
	userUnbanAccounts *task.UserUnbanAccounts,
	userOutboxPublishBatch *task.UserOutboxPublishBatch,
	contentOutboxPublishBatch *task.ContentOutboxPublishBatch,
	contentPublishScheduledArticles *task.ContentPublishScheduledArticles,
	contentFlushArticleViews *task.ContentFlushArticleViews,
	contentDrawArticleLotteries *task.ContentDrawArticleLotteries,
	contentFinishArticlePolls *task.ContentFinishArticlePolls,
	platformRefreshHoliday *task.PlatformRefreshHoliday,
) map[string]task.Task {
	return map[string]task.Task{
		noop.HandlerName().String():                            noop,
		userUnbanAccounts.HandlerName().String():               userUnbanAccounts,
		userOutboxPublishBatch.HandlerName().String():          userOutboxPublishBatch,
		contentOutboxPublishBatch.HandlerName().String():       contentOutboxPublishBatch,
		contentPublishScheduledArticles.HandlerName().String(): contentPublishScheduledArticles,
		contentFlushArticleViews.HandlerName().String():        contentFlushArticleViews,
		contentDrawArticleLotteries.HandlerName().String():     contentDrawArticleLotteries,
		contentFinishArticlePolls.HandlerName().String():       contentFinishArticlePolls,
		platformRefreshHoliday.HandlerName().String():          platformRefreshHoliday,
	}
}
