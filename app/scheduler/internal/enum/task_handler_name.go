package enum

import (
	commonenum "common/pkg/enum"
	schedulerv1 "common/proto/gen/scheduler/enum/v1"
)

// TaskHandlerName 定义代码内任务处理器名称。
type TaskHandlerName string

const (
	TaskHandlerNameNoop                            TaskHandlerName = "noop"                               // 空处理器，仅用于占位
	TaskHandlerNameUserOutboxPublishBatch          TaskHandlerName = "user.outbox_publish_batch"          // 批量发布用户发件箱事件
	TaskHandlerNameUserUnbanAccounts               TaskHandlerName = "user.unban_accounts"                // 解除到期账号封禁
	TaskHandlerNameUserSettleLiveness              TaskHandlerName = "user.settle_liveness"               // 结算用户活跃度奖励
	TaskHandlerNameContentOutboxPublishBatch       TaskHandlerName = "content.outbox_publish_batch"       // 批量发布内容发件箱事件
	TaskHandlerNameContentPublishScheduledArticles TaskHandlerName = "content.publish_scheduled_articles" // 发布定时文章
	TaskHandlerNameContentFlushArticleViews        TaskHandlerName = "content.flush_article_views"        // 刷新文章浏览量
	TaskHandlerNameContentDrawArticleLotteries     TaskHandlerName = "content.draw_article_lotteries"     // 开奖文章抽奖活动
	TaskHandlerNameContentFinishArticlePolls       TaskHandlerName = "content.finish_article_polls"       // 结束文章投票活动
	TaskHandlerNamePlatformRefreshHoliday          TaskHandlerName = "platform.refresh_holiday"           // 刷新节假日数据
	TaskHandlerNamePlatformRefreshIPData           TaskHandlerName = "platform.refresh_ip_data"           // 刷新 IP 地址数据
)

// TaskHandlerNameMap 将内部任务处理器名称映射到 proto 枚举。
var TaskHandlerNameMap = commonenum.NewMapping[TaskHandlerName, schedulerv1.SchedulerTaskHandlerName](
	map[TaskHandlerName]commonenum.Entry[TaskHandlerName, schedulerv1.SchedulerTaskHandlerName]{
		TaskHandlerNameNoop:                            {Proto: schedulerv1.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_NOOP},
		TaskHandlerNameUserOutboxPublishBatch:          {Proto: schedulerv1.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_USER_OUTBOX_PUBLISH_BATCH},
		TaskHandlerNameUserUnbanAccounts:               {Proto: schedulerv1.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_USER_UNBAN_ACCOUNTS},
		TaskHandlerNameUserSettleLiveness:              {Proto: schedulerv1.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_USER_SETTLE_LIVENESS},
		TaskHandlerNameContentOutboxPublishBatch:       {Proto: schedulerv1.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_CONTENT_OUTBOX_PUBLISH_BATCH},
		TaskHandlerNameContentPublishScheduledArticles: {Proto: schedulerv1.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_CONTENT_PUBLISH_SCHEDULED_ARTICLES},
		TaskHandlerNameContentFlushArticleViews:        {Proto: schedulerv1.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_CONTENT_FLUSH_ARTICLE_VIEWS},
		TaskHandlerNameContentDrawArticleLotteries:     {Proto: schedulerv1.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_CONTENT_DRAW_ARTICLE_LOTTERIES},
		TaskHandlerNameContentFinishArticlePolls:       {Proto: schedulerv1.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_CONTENT_FINISH_ARTICLE_POLLS},
		TaskHandlerNamePlatformRefreshHoliday:          {Proto: schedulerv1.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_PLATFORM_REFRESH_HOLIDAY},
		TaskHandlerNamePlatformRefreshIPData:           {Proto: schedulerv1.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_PLATFORM_REFRESH_IP_DATA},
	},
)

func (TaskHandlerName) Values() []string {
	return TaskHandlerNameMap.EnumValues()
}

func (e TaskHandlerName) String() string {
	return string(e)
}
