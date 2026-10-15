package repo

import (
	"common/pkg/client/rpc"
	commonscheduler "common/pkg/scheduler"
	schedulerv1 "common/proto/gen/scheduler/v1"
	schedulerv1enum "common/proto/gen/scheduler/v1/enum"
	"context"
	"errors"
	"strconv"
	"time"

	bizrepo "content/internal/biz/repo"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DelayedTaskClient struct {
	schedulerClient *rpc.SchedulerClient
}

func NewDelayedTaskClient(
	schedulerClient *rpc.SchedulerClient,
) bizrepo.DelayedTaskClient {
	return &DelayedTaskClient{
		schedulerClient: schedulerClient,
	}
}

func (c *DelayedTaskClient) RegisterPublishScheduledArticle(ctx context.Context, articleID int64, authorUserID int64, publishAt time.Time) error {
	if publishAt.IsZero() {
		return errors.New("delayed task publish_at is required")
	}
	payload, err := commonscheduler.MarshalPayload(&schedulerv1.PublishScheduledArticleTaskPayload{ArticleId: articleID, AuthorUserId: authorUserID, ScheduledAt: timestamppb.New(publishAt.UTC().Truncate(time.Second))})
	if err != nil {
		return err
	}
	_, err = c.schedulerClient.DelayedTask.Schedule(ctx, &schedulerv1.ScheduleSchedulerDelayedTask_Req{
		HandlerName:    schedulerv1enum.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_CONTENT_PUBLISH_SCHEDULED_ARTICLES,
		Payload:        payload,
		ScheduledAt:    timestamppb.New(publishAt),
		IdempotencyKey: uuid.NewString(),
		BusinessKey:    c.publishScheduledArticleBusinessKey(articleID),
	})
	return err
}

func (c *DelayedTaskClient) CancelPublishScheduledArticle(ctx context.Context, articleID int64) error {
	_, err := c.schedulerClient.DelayedTask.CancelExecution(ctx, &schedulerv1.CancelSchedulerDelayedTaskExecution_Req{
		BusinessKey: c.publishScheduledArticleBusinessKey(articleID),
	})
	return err
}

func (c *DelayedTaskClient) RegisterLotteryDraw(ctx context.Context, articleID int64, drawAt time.Time) error {
	payload, err := commonscheduler.MarshalPayload(&schedulerv1.DrawArticleLotteryTaskPayload{ArticleId: articleID, DrawAt: timestamppb.New(drawAt.UTC().Truncate(time.Second))})
	if err != nil {
		return err
	}
	_, err = c.schedulerClient.DelayedTask.Schedule(ctx, &schedulerv1.ScheduleSchedulerDelayedTask_Req{HandlerName: schedulerv1enum.SchedulerTaskHandlerName_SCHEDULER_TASK_HANDLER_NAME_CONTENT_DRAW_ARTICLE_LOTTERIES, Payload: payload, ScheduledAt: timestamppb.New(drawAt), IdempotencyKey: uuid.NewString(), BusinessKey: "content.article.lottery.draw:" + strconv.FormatInt(articleID, 10)})
	return err
}

func (*DelayedTaskClient) publishScheduledArticleBusinessKey(articleID int64) string {
	return "content.article.publish:" + strconv.FormatInt(articleID, 10)
}
