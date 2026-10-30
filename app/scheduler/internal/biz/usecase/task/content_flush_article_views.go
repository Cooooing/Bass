package task

import (
	"common/pkg/client/rpc"
	commonscheduler "common/pkg/scheduler"
	contentv1 "common/proto/gen/content/v1"
	schedulerv1 "common/proto/gen/scheduler/v1"
	"context"
	"fmt"
	"time"

	schedulerenum "scheduler/internal/enum"
)

type ContentFlushArticleViews struct {
	contentClient *rpc.ContentClient
	title         string
	description   string
	defaultLimit  int32
}

func NewContentFlushArticleViews(
	contentClient *rpc.ContentClient,
) *ContentFlushArticleViews {
	return &ContentFlushArticleViews{
		contentClient: contentClient,
		title:         "刷新文章浏览量",
		description:   "调用 content.ArticleService.FlushViews 将 Redis 浏览增量刷新到数据库。",
		defaultLimit:  1000,
	}
}

func (t *ContentFlushArticleViews) HandlerName() schedulerenum.TaskHandlerName {
	return schedulerenum.TaskHandlerNameContentFlushArticleViews
}

func (t *ContentFlushArticleViews) Title() string {
	return t.title
}

func (t *ContentFlushArticleViews) Description() string {
	return t.description
}

func (t *ContentFlushArticleViews) DefaultScheduledTasks() []*DefaultScheduledTask {
	return []*DefaultScheduledTask{
		{
			Title:         t.Title(),
			Description:   t.Description(),
			Enabled:       true,
			CronSpec:      "0/30 * * * * ?",
			Payload:       commonscheduler.MustMarshalPayload(&schedulerv1.FlushArticleViewsTaskPayload{Limit: t.defaultLimit}),
			Timeout:       30 * time.Second,
			MaxAttempts:   1,
			MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteLatest,
			AllowOverlap:  false,
		},
	}
}

func (t *ContentFlushArticleViews) DefaultDelayedTasks() []*DefaultDelayedTask {
	return nil
}

func (t *ContentFlushArticleViews) Execute(ctx context.Context, execution Execution) (*Result, error) {
	limit := t.defaultLimit
	data := &schedulerv1.FlushArticleViewsTaskPayload{}
	if err := commonscheduler.UnmarshalPayload(execution.Payload, data); err != nil {
		return nil, err
	}
	if data.GetLimit() > 0 {
		limit = data.GetLimit()
	}
	response, err := t.contentClient.Article.FlushViews(ctx, &contentv1.FlushArticleViews_Req{Limit: limit})
	if err != nil {
		return nil, err
	}
	return &Result{Summary: fmt.Sprintf("flushed=%d", response.GetFlushed())}, nil
}
