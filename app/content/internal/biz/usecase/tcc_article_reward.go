package usecase

import (
	"context"
	"log/slog"

	"common/pkg/apperror"
	commonclient "common/pkg/client"
	commonenums "common/proto/gen/common/enums"
	cerrors "common/proto/gen/common/errors"
	contentv1 "common/proto/gen/content/v1"

	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/enum"
)

// TCCArticleRewardUsecase 执行文章打赏 TCC 分支。
type TCCArticleRewardUsecase struct {
	articles  repo.ArticleRepo
	actions   repo.ArticleActionRecordRepo
	outbox    repo.OutboxEventRepo
	publisher *OutboxUsecase
	log       *slog.Logger
	branch    commonclient.TCCBranchExecutor
}

func NewTCCArticleRewardUsecase(logger *slog.Logger, articles repo.ArticleRepo, actions repo.ArticleActionRecordRepo, outbox repo.OutboxEventRepo, publisher *OutboxUsecase, branch commonclient.TCCBranchExecutor) *TCCArticleRewardUsecase {
	return &TCCArticleRewardUsecase{articles: articles, actions: actions, outbox: outbox, publisher: publisher, log: logger, branch: branch}
}

func (u *TCCArticleRewardUsecase) Try(ctx context.Context, req *contentv1.TCCArticleReward_Req) error {
	return u.branch.Execute(ctx, func(ctx context.Context) error {
		if req.GetArticleId() <= 0 || req.GetSenderUserId() <= 0 || req.GetReceiverUserId() <= 0 || req.GetSenderUserId() == req.GetReceiverUserId() || req.GetPoints() <= 0 {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		article, err := u.articles.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: new(req.ArticleId)}})
		if err != nil {
			return err
		}
		if article.CreatedBy == nil || *article.CreatedBy != req.GetReceiverUserId() || article.IsAuthor(req.GetSenderUserId()) || article.RewardPoints == nil || int64(*article.RewardPoints) != req.GetPoints() {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		action := enum.ArticleActionReward
		created, err := u.actions.Save(ctx, &model.ArticleActionRecord{ArticleID: req.GetArticleId(), UserID: req.GetSenderUserId(), Type: action, Points: new(req.Points), Status: enum.ContentActionStatusPending})
		if err != nil {
			return err
		}
		if !created {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
		}
		return nil
	})
}

func (u *TCCArticleRewardUsecase) Confirm(ctx context.Context, req *contentv1.TCCArticleReward_Req) error {
	var event *repo.OutboxEvent
	err := u.branch.Execute(ctx, func(ctx context.Context) error {
		action := enum.ArticleActionReward
		changed, err := u.actions.TransitionStatus(ctx, &repo.ArticleActionRecordReq{ArticleID: new(req.ArticleId), UserID: new(req.SenderUserId), Type: &action}, enum.ContentActionStatusPending, enum.ContentActionStatusPublished)
		if err != nil || !changed {
			return err
		}
		if err = u.articles.AddStats(ctx, &repo.ArticleAddStatsReq{ArticleID: req.GetArticleId(), Stats: repo.ArticleStatUpdate{RewardCount: 1}}); err != nil {
			return err
		}
		event, err = u.outbox.Save(ctx, &commonenums.Event{Type: commonenums.EventType_EVENT_TYPE_ARTICLE_REWARDED, Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_REWARDED, Payload: &commonenums.Event_ArticleRewarded{ArticleRewarded: &commonenums.ArticleRewardedPayload{SenderId: req.GetSenderUserId(), ArticleId: req.GetArticleId()}}})
		return err
	})
	if err != nil || event == nil || u.publisher == nil {
		return err
	}
	if _, publishErr := u.publisher.Publish(context.WithoutCancel(ctx), &PublishOutboxEventReq{ID: event.ID}); publishErr != nil {
		u.log.WarnContext(ctx, "publish article reward outbox event failed", slog.Int64("outbox_id", event.ID), slog.Any("err", publishErr))
	}
	return nil
}

func (u *TCCArticleRewardUsecase) Cancel(ctx context.Context, req *contentv1.TCCArticleReward_Req) error {
	return u.branch.Execute(ctx, func(ctx context.Context) error {
		action := enum.ArticleActionReward
		_, err := u.actions.TransitionStatus(ctx, &repo.ArticleActionRecordReq{ArticleID: new(req.ArticleId), UserID: new(req.SenderUserId), Type: &action}, enum.ContentActionStatusPending, enum.ContentActionStatusCanceled)
		return err
	})
}
