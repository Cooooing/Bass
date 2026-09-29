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
	"strconv"
	"time"

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

type publishScheduledArticlePayload struct {
	ArticleID     int64     `json:"article_id"`
	AuthorUserID  int64     `json:"author_user_id"`
	ScheduledAt   time.Time `json:"scheduled_at"`
}

func (c *DelayedTaskClient) RegisterPublishScheduledArticle(ctx context.Context, articleID int64, authorUserID int64, publishAt time.Time) error {
	if publishAt.IsZero() {
		return errors.New("delayed task publish_at is required")
	}
	payload, err := json.Marshal(&publishScheduledArticlePayload{
		ArticleID: articleID,
		AuthorUserID: authorUserID,
		ScheduledAt: publishAt.UTC().Truncate(time.Second),
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
		IdempotencyKey: uuid.NewString(),
		BusinessKey:    publishScheduledArticleBusinessKey(articleID),
	})
	return err
}

func (c *DelayedTaskClient) CancelPublishScheduledArticle(ctx context.Context, articleID int64) error {
	_, err := c.schedulerClient.DelayedTask.CancelExecution(ctx, &schedulerv1.CancelSchedulerDelayedTaskExecution_Req{
		BusinessKey: publishScheduledArticleBusinessKey(articleID),
	})
	return err
}

func publishScheduledArticleBusinessKey(articleID int64) string {
	return "content.article.publish:" + strconv.FormatInt(articleID, 10)
}
