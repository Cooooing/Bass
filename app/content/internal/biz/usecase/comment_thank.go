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

type CommentThankUsecase struct {
	comments repo.CommentRepo
	articles repo.ArticleRepo
	actions  repo.CommentActionRecordRepo
	dtm      *commonclient.DTMClient
	branches *commonclient.TCCBranchCaller
	business *config.Business
}

func NewCommentThankUsecase(comments repo.CommentRepo, articles repo.ArticleRepo, actions repo.CommentActionRecordRepo, dtm *commonclient.DTMClient, branches *commonclient.TCCBranchCaller, business *config.Business) *CommentThankUsecase {
	return &CommentThankUsecase{comments, articles, actions, dtm, branches, business}
}
func (u *CommentThankUsecase) Thank(ctx context.Context, access *model.ContentAccess, commentID int64) (bool, error) {
	c, err := u.comments.Get(ctx, &repo.CommentGetReq{Filter: &model.CommentFilter{CommentID: &commentID}})
	if err != nil {
		return false, err
	}
	a, err := u.articles.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: &c.ArticleID}})
	if err != nil {
		return false, err
	}
	if err = (&CommentAccessUsecase{}).CanInteract(access, c, a); err != nil {
		return false, err
	}
	if c.CreatedBy == nil || *c.CreatedBy == access.ActorUserID {
		return false, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	actor := access.ActorUserID
	action := enum.CommentActionThank
	status := enum.ContentActionStatusPublished
	exists, err := u.actions.Exist(ctx, &repo.CommentActionRecordReq{CommentID: &commentID, UserID: &actor, Type: &action, Status: &status})
	if err != nil || exists {
		return exists, err
	}
	points := int64(20)
	if u.business != nil && u.business.GetComment().GetThankPoints() > 0 {
		points = u.business.GetComment().GetThankPoints()
	}
	req := &contentv1.TCCCommentThank_Req{CommentId: commentID, SenderUserId: actor, ReceiverUserId: *c.CreatedBy, Points: points}
	transfer := &economyv1.TCCPointsTransfer_Req{FromUserId: actor, ToUserId: *c.CreatedBy, Amount: points, OutRecordType: economyv1enum.EconomyRecordType_ECONOMY_RECORD_TYPE_COMMENT_THANK_EXPENSE, InRecordType: economyv1enum.EconomyRecordType_ECONOMY_RECORD_TYPE_COMMENT_THANK_REWARD, IdempotencyKey: new(fmt.Sprintf("comment thank:%d", commentID))}
	_, err = u.dtm.InTCC(ctx, func(t *dtmgrpc.TccGrpc) error {
		if err = u.branches.Call(ctx, t, constant.ContentServiceName.String(), req, contentv1.TCCCommentThankService_Try_FullMethodName, contentv1.TCCCommentThankService_Confirm_FullMethodName, contentv1.TCCCommentThankService_Cancel_FullMethodName, &contentv1.TCCCommentThank_Resp{}); err != nil {
			return err
		}
		return u.branches.Call(ctx, t, constant.EconomyServiceName.String(), transfer, economyv1.TCCPointsTransferService_Try_FullMethodName, economyv1.TCCPointsTransferService_Confirm_FullMethodName, economyv1.TCCPointsTransferService_Cancel_FullMethodName, &economyv1.TCCPointsTransfer_Resp{})
	})
	return err == nil, err
}

type TCCCommentThankUsecase struct {
	comments  repo.CommentRepo
	actions   repo.CommentActionRecordRepo
	outbox    repo.OutboxEventRepo
	publisher *OutboxUsecase
	log       *slog.Logger
	branch    commonclient.TCCBranchExecutor
}

func NewTCCCommentThankUsecase(logger *slog.Logger, comments repo.CommentRepo, actions repo.CommentActionRecordRepo, outbox repo.OutboxEventRepo, publisher *OutboxUsecase, branch commonclient.TCCBranchExecutor) *TCCCommentThankUsecase {
	return &TCCCommentThankUsecase{comments: comments, actions: actions, outbox: outbox, publisher: publisher, log: logger, branch: branch}
}
func (u *TCCCommentThankUsecase) Try(ctx context.Context, r *contentv1.TCCCommentThank_Req) error {
	return u.branch.Execute(ctx, func(ctx context.Context) error {
		if r.GetCommentId() <= 0 || r.GetSenderUserId() <= 0 || r.GetReceiverUserId() <= 0 || r.GetSenderUserId() == r.GetReceiverUserId() || r.GetPoints() <= 0 {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
		}
		action := enum.CommentActionThank
		exists, err := u.actions.Exist(ctx, &repo.CommentActionRecordReq{CommentID: new(r.CommentId), UserID: new(r.SenderUserId), Type: &action})
		if err != nil || exists {
			return err
		}
		created, err := u.actions.Save(ctx, &model.CommentActionRecord{CommentID: r.GetCommentId(), UserID: r.GetSenderUserId(), Type: action, Points: new(r.Points), Status: enum.ContentActionStatusPending})
		if err != nil {
			return err
		}
		if !created {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_CONFLICT)
		}
		return nil
	})
}
func (u *TCCCommentThankUsecase) Confirm(ctx context.Context, r *contentv1.TCCCommentThank_Req) error {
	var outboxEvent *repo.OutboxEvent
	err := u.branch.Execute(ctx, func(ctx context.Context) error {
		action := enum.CommentActionThank
		changed, err := u.actions.TransitionStatus(ctx, &repo.CommentActionRecordReq{CommentID: new(r.CommentId), UserID: new(r.SenderUserId), Type: &action}, enum.ContentActionStatusPending, enum.ContentActionStatusPublished)
		if err != nil || !changed {
			return err
		}
		if err = u.comments.AddStats(ctx, &repo.CommentAddStatsReq{CommentID: r.GetCommentId(), Stats: repo.CommentStatUpdate{ThankCount: 1}}); err != nil {
			return err
		}
		outboxEvent, err = u.outbox.Save(ctx, (&CommentUsecase{}).commentThankedEvent(r.GetSenderUserId(), r.GetCommentId()))
		return err
	})
	if err != nil || outboxEvent == nil || u.publisher == nil {
		return err
	}
	if _, publishErr := u.publisher.Publish(context.WithoutCancel(ctx), &PublishOutboxEventReq{ID: outboxEvent.ID}); publishErr != nil {
		u.log.WarnContext(ctx, "publish comment thank outbox event failed", slog.Int64("outbox_id", outboxEvent.ID), slog.Any("err", publishErr))
	}
	return nil
}
func (u *TCCCommentThankUsecase) Cancel(ctx context.Context, r *contentv1.TCCCommentThank_Req) error {
	return u.branch.Execute(ctx, func(ctx context.Context) error {
		action := enum.CommentActionThank
		_, err := u.actions.TransitionStatus(ctx, &repo.CommentActionRecordReq{CommentID: new(r.CommentId), UserID: new(r.SenderUserId), Type: &action}, enum.ContentActionStatusPending, enum.ContentActionStatusCanceled)
		return err
	})
}
