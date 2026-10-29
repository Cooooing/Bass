package repo

import (
	"context"
	"time"

	"common/pkg/client/rpc"
	contentv1 "common/proto/gen/content/v1"
	contentenum "common/proto/gen/content/v1/enum"
	imv1 "common/proto/gen/im/v1"

	"user/internal/biz/model"
	bizrepo "user/internal/biz/repo"

	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ bizrepo.LivenessSource = (*LivenessSource)(nil)

// LivenessSource 并行适配内容和聊天室的资源计数协议。
type LivenessSource struct {
	content *rpc.ContentClient
	im      *rpc.IMClient
}

// NewLivenessSource 创建跨领域资源计数读取器。
func NewLivenessSource(content *rpc.ContentClient, im *rpc.IMClient) bizrepo.LivenessSource {
	return &LivenessSource{content: content, im: im}
}

// BatchGet 并行读取一组账号、同一时间窗口的内容和聊天室原始计数。
func (s *LivenessSource) BatchGet(ctx context.Context, userIDs []int64, start, end time.Time) (map[int64]*model.LivenessDetail, error) {
	requestStart := timestamppb.New(start)
	requestEnd := timestamppb.New(end)
	group, groupCtx := errgroup.WithContext(ctx)

	var (
		contentRows *contentv1.BatchCountContentUserBehaviors_Resp
		messageRows *imv1.BatchCountGlobalChatMessages_Resp
	)

	group.Go(func() error {
		var err error
		contentRows, err = s.content.Statistics.BatchCountUserBehaviors(groupCtx, &contentv1.BatchCountContentUserBehaviors_Req{UserIds: userIDs, Start: requestStart, End: requestEnd, ActionStatuses: []contentenum.ContentActionStatus{contentenum.ContentActionStatus_CONTENT_ACTION_STATUS_PUBLISHED}})
		return err
	})
	group.Go(func() error {
		var err error
		messageRows, err = s.im.Statistics.BatchCountGlobalChatMessages(groupCtx, &imv1.BatchCountGlobalChatMessages_Req{UserIds: userIDs, Start: requestStart, End: requestEnd})
		return err
	})
	if err := group.Wait(); err != nil {
		return nil, err
	}
	result := make(map[int64]*model.LivenessDetail, len(userIDs))
	for _, userID := range userIDs {
		result[userID] = &model.LivenessDetail{}
	}
	for _, row := range contentRows.GetRows() {
		result[row.GetUserId()] = &model.LivenessDetail{ArticleCount: row.GetArticleCount(), CommentCount: row.GetCommentCount(), LikeCount: row.GetArticleLikeCount() + row.GetCommentLikeCount(), ThankCount: row.GetArticleThankCount() + row.GetCommentThankCount(), RewardCount: row.GetArticleRewardCount(), ViewCount: row.GetViewCount()}
	}
	for _, row := range messageRows.GetRows() {
		result[row.GetUserId()].GlobalChatMessageCount = row.GetCount()
	}
	return result, nil
}
