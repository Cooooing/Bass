package task

import (
	"common/pkg/client/rpc"
	contentv1 "common/proto/gen/content/v1"
	"context"
	"encoding/json"
	"fmt"
	schedulerenum "scheduler/internal/enum"
	"strings"
	"time"
)

type ContentPublishScheduledArticles struct {
	contentClient *rpc.ContentClient
	title         string
	description   string
}

func NewContentPublishScheduledArticles(
	contentClient *rpc.ContentClient,
) *ContentPublishScheduledArticles {
	return &ContentPublishScheduledArticles{
		contentClient: contentClient,
		title:         "内容定时发布",
		description:   "按文章延迟任务调用 content.ArticleService.Publish 发布定时文章",
	}
}

func (t *ContentPublishScheduledArticles) HandlerName() schedulerenum.TaskHandlerName {
	return schedulerenum.TaskHandlerNameContentPublishScheduledArticles
}

func (t *ContentPublishScheduledArticles) Title() string {
	return t.title
}

func (t *ContentPublishScheduledArticles) Description() string {
	return t.description
}

func (t *ContentPublishScheduledArticles) DefaultScheduledTasks() []*DefaultScheduledTask {
	return nil
}

func (t *ContentPublishScheduledArticles) DefaultDelayedTasks() []*DefaultDelayedTask {
	return []*DefaultDelayedTask{
		{
			TaskKey:       schedulerenum.TaskKeyContentPublishScheduledArticlesDefault,
			Title:         t.Title(),
			Description:   t.Description(),
			Enabled:       true,
			Timeout:       30 * time.Second,
			MaxAttempts:   3,
			MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteAll,
		},
	}
}

type contentPublishScheduledArticlesPayload struct {
	ArticleID int64 `json:"article_id"`
}

func (t *ContentPublishScheduledArticles) Execute(ctx context.Context, payload string) error {
	var data contentPublishScheduledArticlesPayload
	if err := json.Unmarshal([]byte(strings.TrimSpace(payload)), &data); err != nil {
		return err
	}
	if data.ArticleID <= 0 {
		return fmt.Errorf("invalid content publish scheduled article payload")
	}
	_, err := t.contentClient.Article.Publish(ctx, &contentv1.PublishArticle_Req{
		ArticleId: data.ArticleID,
	})
	return err
}
