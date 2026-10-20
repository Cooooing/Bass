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

type ContentFinishArticlePolls struct {
	contentClient *rpc.ContentClient
}

func NewContentFinishArticlePolls(contentClient *rpc.ContentClient) *ContentFinishArticlePolls {
	return &ContentFinishArticlePolls{contentClient: contentClient}
}

func (*ContentFinishArticlePolls) HandlerName() schedulerenum.TaskHandlerName {
	return schedulerenum.TaskHandlerNameContentFinishArticlePolls
}

func (*ContentFinishArticlePolls) Title() string { return "文章投票截止" }

func (*ContentFinishArticlePolls) Description() string {
	return "按延迟任务完成文章投票截止。"
}

func (*ContentFinishArticlePolls) DefaultScheduledTasks() []*DefaultScheduledTask { return nil }

func (t *ContentFinishArticlePolls) DefaultDelayedTasks() []*DefaultDelayedTask {
	return []*DefaultDelayedTask{{Title: t.Title(), Description: t.Description(), Enabled: true, Timeout: 30 * time.Second, MaxAttempts: 3, MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteAll}}
}

func (t *ContentFinishArticlePolls) Execute(ctx context.Context, payload string) error {
	data := &schedulerv1.FinishArticlePollTaskPayload{}
	if err := commonscheduler.UnmarshalPayload(payload, data); err != nil {
		return err
	}
	if data.GetArticleId() <= 0 || data.GetEndsAt() == nil || !data.GetEndsAt().IsValid() {
		return fmt.Errorf("invalid content article poll finish payload")
	}
	_, err := t.contentClient.Article.FinishPoll(ctx, &contentv1.FinishArticlePoll_Req{ArticleId: data.GetArticleId(), EndsAt: timestamppb.New(data.GetEndsAt().AsTime())})
	return err
}
