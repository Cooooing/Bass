package task

import (
	"common/pkg/client/rpc"
	contentv1 "common/proto/gen/content/v1"
	"context"
	"encoding/json"
	"fmt"
	"time"

	schedulerenum "scheduler/internal/enum"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ContentDrawArticleLotteries struct {
	contentClient *rpc.ContentClient
}

func NewContentDrawArticleLotteries(contentClient *rpc.ContentClient) *ContentDrawArticleLotteries {
	return &ContentDrawArticleLotteries{contentClient: contentClient}
}

func (*ContentDrawArticleLotteries) HandlerName() schedulerenum.TaskHandlerName {
	return schedulerenum.TaskHandlerNameContentDrawArticleLotteries
}

func (*ContentDrawArticleLotteries) Title() string { return "文章抽奖开奖" }

func (*ContentDrawArticleLotteries) Description() string {
	return "按延迟任务调用 Content 完成文章抽奖开奖。"
}

func (*ContentDrawArticleLotteries) DefaultScheduledTasks() []*DefaultScheduledTask { return nil }

func (t *ContentDrawArticleLotteries) DefaultDelayedTasks() []*DefaultDelayedTask {
	return []*DefaultDelayedTask{{TaskKey: schedulerenum.TaskKeyContentDrawArticleLotteriesDefault, Title: t.Title(), Description: t.Description(), Enabled: true, Timeout: 30 * time.Second, MaxAttempts: 3, MisfirePolicy: schedulerenum.TaskMisfirePolicyExecuteAll}}
}

func (t *ContentDrawArticleLotteries) Execute(ctx context.Context, payload string) error {
	var data struct {
		ArticleID int64     `json:"article_id"`
		DrawAt    time.Time `json:"draw_at"`
	}
	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		return err
	}
	if data.ArticleID <= 0 || data.DrawAt.IsZero() {
		return fmt.Errorf("invalid content article lottery draw payload")
	}
	_, err := t.contentClient.Article.DrawLottery(ctx, &contentv1.DrawArticleLottery_Req{ArticleId: data.ArticleID, DrawAt: timestamppb.New(data.DrawAt)})
	return err
}
