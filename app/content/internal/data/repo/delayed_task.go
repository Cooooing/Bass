package repo

import (
	"common/pkg/client/rpc"
	commonscheduler "common/pkg/scheduler"
	schedulerv1 "common/proto/gen/scheduler/v1"
	schedulerv1enum "common/proto/gen/scheduler/v1/enum"
	bizrepo "content/internal/biz/repo"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

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

type publishScheduledArticlePayload struct {
	ArticleID int64 `json:"article_id"`
}

func (c *DelayedTaskClient) RegisterPublishScheduledArticle(ctx context.Context, articleID int64, publishAt time.Time) error {
	if publishAt.IsZero() {
		return errors.New("delayed task publish_at is required")
	}
	payload, err := json.Marshal(&publishScheduledArticlePayload{
		ArticleID: articleID,
	})
	if err != nil {
		return err
	}
	_, err = c.schedulerClient.DelayedTask.Schedule(ctx, &schedulerv1.ScheduleSchedulerDelayedTask_Req{
		TaskKey: commonscheduler.TaskKeyMap.MustToEnum(
			schedulerv1enum.SchedulerTaskKey_SCHEDULER_TASK_KEY_CONTENT_PUBLISH_SCHEDULED_ARTICLES_DEFAULT,
		).String(),
		Payload:        string(payload),
		ScheduledAt:    timestamppb.New(publishAt),
		IdempotencyKey: fmt.Sprintf("content:article:publish:%d", articleID),
	})
	return err
}

func (c *DelayedTaskClient) CancelPublishScheduledArticle(ctx context.Context, articleID int64) error {
	_, err := c.schedulerClient.DelayedTask.CancelExecution(ctx, &schedulerv1.CancelSchedulerDelayedTaskExecution_Req{
		IdempotencyKey: fmt.Sprintf("content:article:publish:%d", articleID),
	})
	return err
}
