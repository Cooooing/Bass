package task

import (
	"common/pkg/client/rpc"
	contentv1 "common/proto/gen/content/v1"
	"context"
	"encoding/json"
	schedulerenum "scheduler/internal/enum"
	"strings"
	"time"
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
			TaskKey:       schedulerenum.TaskKeyContentFlushArticleViewsDefault,
			Title:         t.Title(),
			Description:   t.Description(),
			Enabled:       true,
			CronSpec:      "0/30 * * * * ?",
			Payload:       `{"limit":1000}`,
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

type contentFlushArticleViewsPayload struct {
	Limit int32 `json:"limit"`
}

func (t *ContentFlushArticleViews) Execute(ctx context.Context, payload string) error {
	limit := t.defaultLimit
	payload = strings.TrimSpace(payload)
	if payload != "" && payload != "{}" {
		data := &contentFlushArticleViewsPayload{}
		if err := json.Unmarshal([]byte(payload), data); err != nil {
			return err
		}
		if data.Limit > 0 {
			limit = data.Limit
		}
	}
	_, err := t.contentClient.Article.FlushViews(ctx, &contentv1.FlushArticleViews_Req{Limit: limit})
	return err
}
