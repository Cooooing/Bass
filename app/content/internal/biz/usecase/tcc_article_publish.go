package usecase

import (
	"context"

	"common/pkg/apperror"
	commonclient "common/pkg/client"
	commonevent "common/proto/gen/common/enums/event/v1"
	cerrors "common/proto/gen/common/errors/v1"
	contentv1 "common/proto/gen/content/v1"

	"content/internal/biz/repo"
)

// TCCArticlePublishUsecase 执行文章发布 TCC 分支。
type TCCArticlePublishUsecase struct {
	articles repo.ArticleRepo
	outbox   repo.OutboxEventRepo
	branch   commonclient.TCCBranchExecutor
}

func NewTCCArticlePublishUsecase(articles repo.ArticleRepo, outbox repo.OutboxEventRepo, branch commonclient.TCCBranchExecutor) *TCCArticlePublishUsecase {
	return &TCCArticlePublishUsecase{articles: articles, outbox: outbox, branch: branch}
}

func (u *TCCArticlePublishUsecase) Try(ctx context.Context, req *contentv1.TCCArticlePublish_Req) error {
	return u.branch.Execute(ctx, func(ctx context.Context) error {
		if req.GetArticleId() <= 0 || req.GetOperatorUserId() <= 0 || req.GetPublishedAt() == nil || !req.GetPublishedAt().IsValid() {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		return u.articles.TryPublish(ctx, &repo.ArticleTryPublishReq{ArticleID: req.GetArticleId(), RewardSetupPoints: req.RewardSetupPoints})
	})
}

func (u *TCCArticlePublishUsecase) Confirm(ctx context.Context, req *contentv1.TCCArticlePublish_Req) error {
	return u.branch.Execute(ctx, func(ctx context.Context) error {
		if err := u.articles.ConfirmPublish(ctx, &repo.ArticleConfirmPublishReq{ArticleID: req.GetArticleId(), PublishedAt: req.GetPublishedAt().AsTime(), PublishedCity: req.City, UpdatedBy: new(req.OperatorUserId)}); err != nil {
			return err
		}
		_, err := u.outbox.Save(ctx, &commonevent.Event{Subject: commonevent.EventSubject_EVENT_SUBJECT_ARTICLE_PUBLISHED, Payload: &commonevent.Event_ArticlePublished{ArticlePublished: &commonevent.ArticlePublishedPayload{ArticleId: req.GetArticleId()}}})
		return err
	})
}

func (u *TCCArticlePublishUsecase) Cancel(ctx context.Context, req *contentv1.TCCArticlePublish_Req) error {
	return u.branch.Execute(ctx, func(ctx context.Context) error {
		return u.articles.CancelPublish(ctx, req.GetArticleId())
	})
}
