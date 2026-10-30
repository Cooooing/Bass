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

	"google.golang.org/protobuf/types/known/timestamppb"
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
			Title:         t.Title(),
			Description:   t.Description(),
			Enabled:       true,
			Timeout:       30 * time.Second,
			MaxAttempts:   3,
			MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteAll,
		},
	}
}

func (t *ContentPublishScheduledArticles) Execute(ctx context.Context, execution Execution) (*Result, error) {
	data := &schedulerv1.PublishScheduledArticleTaskPayload{}
	if err := commonscheduler.UnmarshalPayload(execution.Payload, data); err != nil {
		return nil, err
	}
	if data.GetArticleId() <= 0 || data.GetScheduledAt() == nil || !data.GetScheduledAt().IsValid() {
		return nil, fmt.Errorf("invalid content publish scheduled article payload")
	}
	response, err := t.contentClient.Article.Publish(ctx, &contentv1.PublishArticle_Req{
		ArticleId:   data.GetArticleId(),
		ScheduledAt: timestamppb.New(data.GetScheduledAt().AsTime()),
	})
	if err != nil || response == nil {
		if err == nil {
			err = fmt.Errorf("scheduled article publish returned an empty response")
		}
		return nil, err
	}
	return &Result{Summary: fmt.Sprintf("article_id=%d", data.GetArticleId())}, nil
}
