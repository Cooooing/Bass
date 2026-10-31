package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"common/pkg/apperror"
	commonclient "common/pkg/client"
	"common/pkg/constant"
	cerrors "common/proto/gen/common/errors"
	contentv1 "common/proto/gen/content/v1"
	economyv1 "common/proto/gen/economy/v1"
	economyv1enum "common/proto/gen/economy/v1/enum"

	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/config"
	"content/internal/enum"

	"github.com/dtm-labs/client/dtmgrpc"
)

// ArticleThankUsecase 编排文章感谢的积分转账。
type ArticleThankUsecase struct {
	articles repo.ArticleRepo
	actions  repo.ArticleActionRecordRepo
	outbox   repo.OutboxEventRepo
	dtm      *commonclient.DTMClient
	branches *commonclient.TCCBranchCaller
	business *config.Business
}

func NewArticleThankUsecase(articles repo.ArticleRepo, actions repo.ArticleActionRecordRepo, outbox repo.OutboxEventRepo, dtm *commonclient.DTMClient, branches *commonclient.TCCBranchCaller, business *config.Business) *ArticleThankUsecase {
	return &ArticleThankUsecase{articles: articles, actions: actions, outbox: outbox, dtm: dtm, branches: branches, business: business}
}
func (u *ArticleThankUsecase) Thank(ctx context.Context, access *model.ContentAccess, articleID int64) (bool, error) {
	article, err := u.articles.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: &articleID}})
	if err != nil {
		return false, err
	}
	if err = (&ArticleAccessUsecase{}).CanInteract(access, article); err != nil {
		return false, err
	}
	if article.CreatedBy == nil {
		return false, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
	}
	actor := access.ActorUserID
	action := enum.ArticleActionThank
	status := enum.ContentActionStatusPublished
	exists, err := u.actions.Exist(ctx, &repo.ArticleActionRecordReq{ArticleID: &articleID, UserID: &actor, Type: &action, Status: &status})
	if err != nil || exists {
		return exists, err
	}
	points := int64(20)
	if u.business != nil && u.business.GetArticle().GetThankPoints() > 0 {
		points = u.business.GetArticle().GetThankPoints()
	}
	req := &contentv1.TCCArticleThank_Req{ArticleId: articleID, SenderUserId: actor, ReceiverUserId: *article.CreatedBy, Points: points}
	transfer := &economyv1.TCCPointsTransfer_Req{FromUserId: actor, ToUserId: *article.CreatedBy, Amount: points, OutRecordType: economyv1enum.EconomyRecordType_ECONOMY_RECORD_TYPE_ARTICLE_THANK_EXPENSE, InRecordType: economyv1enum.EconomyRecordType_ECONOMY_RECORD_TYPE_ARTICLE_THANK_REWARD, IdempotencyKey: new(fmt.Sprintf("article thank:%d", articleID))}
	_, err = u.dtm.InTCC(ctx, func(tcc *dtmgrpc.TccGrpc) error {
		if err = u.branches.Call(ctx, tcc, constant.ContentServiceName.String(), req, contentv1.TCCArticleThankService_Try_FullMethodName, contentv1.TCCArticleThankService_Confirm_FullMethodName, contentv1.TCCArticleThankService_Cancel_FullMethodName, &contentv1.TCCArticleThank_Resp{}); err != nil {
			return err
		}
		return u.branches.Call(ctx, tcc, constant.EconomyServiceName.String(), transfer, economyv1.TCCPointsTransferService_Try_FullMethodName, economyv1.TCCPointsTransferService_Confirm_FullMethodName, economyv1.TCCPointsTransferService_Cancel_FullMethodName, &economyv1.TCCPointsTransfer_Resp{})
	})
	return err == nil, err
}

// TCCArticleThankUsecase 执行文章感谢 TCC 分支。
type TCCArticleThankUsecase struct {
	articles  repo.ArticleRepo
	actions   repo.ArticleActionRecordRepo
	outbox    repo.OutboxEventRepo
	publisher *OutboxUsecase
	log       *slog.Logger
	branch    commonclient.TCCBranchExecutor
}

func NewTCCArticleThankUsecase(logger *slog.Logger, articles repo.ArticleRepo, actions repo.ArticleActionRecordRepo, outbox repo.OutboxEventRepo, publisher *OutboxUsecase, branch commonclient.TCCBranchExecutor) *TCCArticleThankUsecase {
	return &TCCArticleThankUsecase{articles: articles, actions: actions, outbox: outbox, publisher: publisher, log: logger, branch: branch}
}
func (u *TCCArticleThankUsecase) Try(ctx context.Context, r *contentv1.TCCArticleThank_Req) error {
	return u.branch.Execute(ctx, func(ctx context.Context) error {
		if r.GetArticleId() <= 0 || r.GetSenderUserId() <= 0 || r.GetReceiverUserId() <= 0 || r.GetSenderUserId() == r.GetReceiverUserId() || r.GetPoints() <= 0 {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		a, err := u.articles.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: new(r.ArticleId)}})
		if err != nil {
			return err
		}
		if a.CreatedBy == nil || *a.CreatedBy != r.GetReceiverUserId() || a.IsAuthor(r.GetSenderUserId()) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		if err = a.CanInteract(); err != nil {
			return err
		}
		action := enum.ArticleActionThank
		actionExists, err := u.actions.Exist(ctx, &repo.ArticleActionRecordReq{ArticleID: new(r.ArticleId), UserID: new(r.SenderUserId), Type: &action})
		if err != nil {
			return err
		}
		if actionExists {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
		}
		created, err := u.actions.Save(ctx, &model.ArticleActionRecord{ArticleID: r.GetArticleId(), UserID: r.GetSenderUserId(), Type: action, Points: new(r.Points), Status: enum.ContentActionStatusPending})
		if err != nil {
			return err
		}
		if !created {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
		}
		return nil
	})
}
func (u *TCCArticleThankUsecase) Confirm(ctx context.Context, r *contentv1.TCCArticleThank_Req) error {
	var outboxEvent *repo.OutboxEvent
	err := u.branch.Execute(ctx, func(ctx context.Context) error {
		action := enum.ArticleActionThank
		changed, err := u.actions.TransitionStatus(ctx, &repo.ArticleActionRecordReq{ArticleID: new(r.ArticleId), UserID: new(r.SenderUserId), Type: &action}, enum.ContentActionStatusPending, enum.ContentActionStatusPublished)
		if err != nil {
			return err
		}
		if !changed {
			return nil
		}
		if err := u.articles.AddStats(ctx, &repo.ArticleAddStatsReq{ArticleID: r.GetArticleId(), Stats: repo.ArticleStatUpdate{ThankCount: 1}}); err != nil {
			return err
		}
		outboxEvent, err = u.outbox.Save(ctx, (&ArticleUsecase{}).articleThankedEvent(r.GetSenderUserId(), r.GetArticleId()))
		return err
	})
	if err != nil || outboxEvent == nil || u.publisher == nil {
		return err
	}
	if _, publishErr := u.publisher.Publish(context.WithoutCancel(ctx), &PublishOutboxEventReq{ID: outboxEvent.ID}); publishErr != nil {
		u.log.WarnContext(ctx, "publish article thank outbox event failed", slog.Int64("outbox_id", outboxEvent.ID), slog.Any("err", publishErr))
	}
	return nil
}
func (u *TCCArticleThankUsecase) Cancel(ctx context.Context, r *contentv1.TCCArticleThank_Req) error {
	return u.branch.Execute(ctx, func(ctx context.Context) error {
		action := enum.ArticleActionThank
		_, err := u.actions.TransitionStatus(ctx, &repo.ArticleActionRecordReq{ArticleID: new(r.ArticleId), UserID: new(r.SenderUserId), Type: &action}, enum.ContentActionStatusPending, enum.ContentActionStatusCanceled)
		return err
	})
}
