package usecase

import (
	commonmodel "common/pkg/model"
	cerrors "common/proto/gen/common/errors"
	"context"

	"common/pkg/apperror"
	commonenums "common/proto/gen/common/enums"

	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/enum"

	"log/slog"

	"github.com/samber/lo"
)

type CommentUsecase struct {
	log *slog.Logger
	tx  base.Tx

	commentRepo             repo.CommentRepo
	commentAccessUsecase    *CommentAccessUsecase
	commentActionRecordRepo repo.CommentActionRecordRepo
	articleRepo             repo.ArticleRepo
	outboxRepo              repo.OutboxEventRepo
	outboxUsecase           *OutboxUsecase
	moderationRecordRepo    repo.ContentModerationRecordRepo
	thankUsecase            *CommentThankUsecase
}

func NewCommentUsecase(
	logger *slog.Logger,
	tx base.Tx,
	commentRepo repo.CommentRepo,
	commentAccessUsecase *CommentAccessUsecase,
	commentActionRecordRepo repo.CommentActionRecordRepo,
	articleRepo repo.ArticleRepo,
	outboxRepo repo.OutboxEventRepo,
	outboxUsecase *OutboxUsecase,
	moderationRecordRepo repo.ContentModerationRecordRepo,
	thankUsecase *CommentThankUsecase,
) *CommentUsecase {
	return &CommentUsecase{
		log:                     logger,
		tx:                      tx,
		commentRepo:             commentRepo,
		commentAccessUsecase:    commentAccessUsecase,
		commentActionRecordRepo: commentActionRecordRepo,
		articleRepo:             articleRepo,
		outboxRepo:              outboxRepo,
		outboxUsecase:           outboxUsecase,
		moderationRecordRepo:    moderationRecordRepo,
		thankUsecase:            thankUsecase,
	}
}

type CommentAddReq struct {
	Access  *model.ContentAccess
	Comment *model.Comment
}

func (d *CommentUsecase) Add(ctx context.Context, req *CommentAddReq) (*model.Comment, error) {
	access, err := req.Access.Normalize("")
	if err != nil {
		return nil, err
	}
	comment := req.Comment
	var (
		c           *model.Comment
		outboxEvent *repo.OutboxEvent
	)
	err = d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(comment.ArticleID)},
		})
		if err != nil {
			return err
		}
		if err := d.commentAccessUsecase.CanCreate(access, article); err != nil {
			return err
		}

		replyComment := &model.Comment{}
		var parentID *int64
		if comment.ReplyID != nil {
			replyComment, err = d.commentRepo.Get(ctx, &repo.CommentGetReq{
				Filter: &model.CommentFilter{
					CommentID:   comment.ReplyID,
					ArticleID:   new(comment.ArticleID),
					Restriction: new(enum.ContentRestrictionNone),
				},
			})
			if err != nil {
				return err
			}
			if replyComment.ParentID == nil {
				parentID = new(replyComment.ID)
			} else {
				parentID = replyComment.ParentID
			}
			err = d.commentRepo.AddStats(ctx, &repo.CommentAddStatsReq{
				CommentID: *parentID,
				Stats: repo.CommentStatUpdate{
					ReplyCount: 1,
				},
			})
			if err != nil {
				return err
			}
		}

		err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{
			ArticleID: article.ID,
			Stats: repo.ArticleStatUpdate{
				ReplyCount: 1,
			},
		})
		if err != nil {
			return err
		}

		save := &model.Comment{
			ArticleID:   comment.ArticleID,
			Content:     comment.Content,
			Level:       replyComment.Level + 1,
			ParentID:    parentID,
			ReplyID:     comment.ReplyID,
			Restriction: enum.ContentRestrictionNone,
			ReplyUserID: replyComment.CreatedBy,
			CreatedBy:   new(access.ActorUserID),
			UpdatedBy:   new(access.ActorUserID),
		}
		save.FormatContent()
		commentResp, commentErr := d.commentRepo.Save(ctx, save)
		if commentErr != nil {
			return commentErr
		}
		c = commentResp
		c.ReplyUserID = replyComment.CreatedBy
		outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_COMMENT_PUBLISHED,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_COMMENT_PUBLISHED,
			Payload: &commonenums.Event_CommentPublished{
				CommentPublished: &commonenums.CommentPublishedPayload{
					CommentId: c.ID,
				},
			},
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	if outboxEvent != nil {
		if _, publishErr := d.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{ID: outboxEvent.ID}); publishErr != nil {
			d.log.WarnContext(ctx, "publish content outbox event failed", slog.Int64("outbox_id", outboxEvent.ID), slog.Any("err", publishErr))
		}
	}
	return c, nil
}

type CommentPageReq struct {
	Access *model.ContentAccess
	Page   *commonmodel.PageReq
	Filter *model.CommentFilter
}

type CommentPageResp struct {
	Page *commonmodel.PageResp
	Rows []*model.Comment
}

func (d *CommentUsecase) Page(ctx context.Context, req *CommentPageReq) (*CommentPageResp, error) {
	if req == nil {
		req = &CommentPageReq{}
	}
	scope, err := d.commentAccessUsecase.BuildScope(req.Access)
	if err != nil {
		return nil, err
	}
	pageResp, err := d.commentRepo.Page(ctx, &repo.CommentGetReq{
		Page:   req.Page,
		Filter: req.Filter,
		Scope:  scope,
	})
	if err != nil {
		return nil, err
	}
	return &CommentPageResp{
		Page: pageResp.Page,
		Rows: pageResp.Rows,
	}, nil
}

type CommentListReplyPreviewsReq struct {
	Access          *model.ContentAccess
	ArticleID       int64
	ParentIDs       []int64
	LimitPerParent  int32
	Restriction     *enum.ContentRestriction
	Restrictions    []enum.ContentRestriction
	Order           *enum.CommentOrder
	ExcludedUserIDs []int64
}

type CommentChildPreview struct {
	ParentID int64
	Rows     []*model.Comment
}

func (d *CommentUsecase) ListReplyPreviews(ctx context.Context, req *CommentListReplyPreviewsReq) ([]*CommentChildPreview, error) {
	articleID := req.ArticleID
	scope, err := d.commentAccessUsecase.BuildScope(req.Access)
	if err != nil {
		return nil, err
	}
	previews, err := d.commentRepo.ListReplyPreviews(ctx, &repo.CommentReplyPreviewReq{
		Filter: &model.CommentFilter{
			ArticleID:       new(articleID),
			Restriction:     req.Restriction,
			Restrictions:    req.Restrictions,
			Order:           req.Order,
			ExcludedUserIDs: req.ExcludedUserIDs,
		},
		Scope:          scope,
		ParentIDs:      req.ParentIDs,
		LimitPerParent: req.LimitPerParent,
	})
	if err != nil {
		return nil, err
	}
	rows := lo.Map(previews, func(preview *repo.CommentReplyPreview, _ int) *CommentChildPreview {
		return &CommentChildPreview{
			ParentID: preview.ParentID,
			Rows:     preview.Rows,
		}
	})
	return rows, nil
}

func (d *CommentUsecase) MapArticleLastComments(ctx context.Context, articleIDs []int64, excludedUserIDs []int64) (map[int64]*model.
	Comment, error) {
	if len(articleIDs) == 0 {
		return map[int64]*model.Comment{}, nil
	}
	comments, err := d.commentRepo.MapArticleLastComments(ctx, &repo.CommentGetReq{
		Filter: &model.CommentFilter{
			ArticleIDs: lo.Uniq(articleIDs),
			Restrictions: []enum.ContentRestriction{
				enum.ContentRestrictionNone,
				enum.ContentRestrictionLocked,
			},
			ExcludedUserIDs: excludedUserIDs,
		},
	})
	if err != nil {
		return nil, err
	}
	return comments, nil
}

type CommentHideReq struct {
	CommentID int64
	Access    *model.ContentAccess
	Reason    *string
}

func (d *CommentUsecase) Hide(ctx context.Context, req *CommentHideReq) error {
	commentID := req.CommentID
	access, err := req.Access.Normalize("")
	if err != nil {
		return err
	}
	reason := req.Reason
	return d.updateRestriction(ctx, &commentUpdateRestrictionReq{
		CommentID:   commentID,
		Restriction: enum.ContentRestrictionHidden,
		Access:      access,
		Action:      enum.ContentModerationActionHide,
		Reason:      reason,
	})
}

type CommentUnhideReq struct {
	CommentID int64
	Access    *model.ContentAccess
	Reason    *string
}

func (d *CommentUsecase) Unhide(ctx context.Context, req *CommentUnhideReq) error {
	commentID := req.CommentID
	access, err := req.Access.Normalize("")
	if err != nil {
		return err
	}
	reason := req.Reason
	return d.updateRestriction(ctx, &commentUpdateRestrictionReq{
		CommentID:   commentID,
		Restriction: enum.ContentRestrictionNone,
		Access:      access,
		Action:      enum.ContentModerationActionUnhide,
		Reason:      reason,
	})
}

type CommentLockReq struct {
	CommentID int64
	Access    *model.ContentAccess
	Reason    *string
}

func (d *CommentUsecase) Lock(ctx context.Context, req *CommentLockReq) error {
	commentID := req.CommentID
	access, err := req.Access.Normalize("")
	if err != nil {
		return err
	}
	reason := req.Reason
	return d.updateRestriction(ctx, &commentUpdateRestrictionReq{
		CommentID:   commentID,
		Restriction: enum.ContentRestrictionLocked,
		Access:      access,
		Action:      enum.ContentModerationActionLock,
		Reason:      reason,
	})
}

type CommentUnlockReq struct {
	CommentID int64
	Access    *model.ContentAccess
	Reason    *string
}

func (d *CommentUsecase) Unlock(ctx context.Context, req *CommentUnlockReq) error {
	commentID := req.CommentID
	access, err := req.Access.Normalize("")
	if err != nil {
		return err
	}
	reason := req.Reason
	return d.updateRestriction(ctx, &commentUpdateRestrictionReq{
		CommentID:   commentID,
		Restriction: enum.ContentRestrictionNone,
		Access:      access,
		Action:      enum.ContentModerationActionUnlock,
		Reason:      reason,
	})
}

type commentUpdateRestrictionReq struct {
	CommentID   int64
	Restriction enum.ContentRestriction
	Access      *model.ContentAccess
	Action      enum.ContentModerationAction
	Reason      *string
}

func (d *CommentUsecase) updateRestriction(ctx context.Context, req *commentUpdateRestrictionReq) error {
	commentID := req.CommentID
	restriction := req.Restriction
	access := req.Access
	action := req.Action
	reason := req.Reason
	if err := d.commentAccessUsecase.CanManage(access); err != nil {
		return err
	}
	var outboxEvent *repo.OutboxEvent
	err := d.tx(ctx, func(ctx context.Context) error {
		commentResp, err := d.commentRepo.Get(ctx, &repo.CommentGetReq{
			Filter: &model.CommentFilter{CommentID: new(commentID)},
		})
		comment := commentResp
		if err != nil {
			return err
		}
		switch action {
		case enum.ContentModerationActionHide, enum.ContentModerationActionLock:
			if comment.Restriction != enum.ContentRestrictionNone {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_STATUS)
			}
		case enum.ContentModerationActionUnhide:
			if comment.Restriction != enum.ContentRestrictionHidden {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_STATUS)
			}
		case enum.ContentModerationActionUnlock:
			if comment.Restriction != enum.ContentRestrictionLocked {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_COMMENT_STATUS)
			}
		}
		if err := d.commentRepo.UpdateRestriction(ctx, &repo.CommentUpdateRestrictionReq{
			CommentID:   commentID,
			Restriction: restriction,
			UpdatedBy:   access.ActorUserID,
		}); err != nil {
			return err
		}
		if _, err := d.moderationRecordRepo.Save(ctx, &model.ContentModerationRecord{
			Target:     enum.ContentModerationTargetComment,
			TargetID:   commentID,
			Action:     action,
			Reason:     reason,
			OperatorID: access.ActorUserID,
		}); err != nil {
			return err
		}
		outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_COMMENT_STATUS_UPDATED,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_COMMENT_STATUS_UPDATED,
			Payload: &commonenums.Event_CommentStatusUpdated{
				CommentStatusUpdated: &commonenums.CommentStatusUpdatedPayload{
					SenderId:    access.ActorUserID,
					CommentId:   commentID,
					Action:      action.String(),
					Restriction: restriction.String(),
					Reason:      reason,
				},
			},
		})
		return err
	})
	if err != nil {
		return err
	}
	if outboxEvent != nil {
		if _, publishErr := d.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{ID: outboxEvent.ID}); publishErr != nil {
			d.log.WarnContext(ctx, "publish content outbox event failed", slog.Int64("outbox_id", outboxEvent.ID), slog.Any("err", publishErr))
		}
	}
	return nil
}

type CommentLikeReq struct {
	CommentID int64
	Access    *model.ContentAccess
	Active    bool
}

func (d *CommentUsecase) Like(ctx context.Context, req *CommentLikeReq) (bool, error) {
	interaction := commentInteraction{commentID: req.CommentID, access: req.Access, active: req.Active, action: enum.CommentActionLike, stats: repo.CommentStatUpdate{LikeCount: 1}, event: d.commentLikedEvent}
	return d.changeInteraction(ctx, interaction)
}

type CommentThankReq struct {
	CommentID       int64
	Access          *model.ContentAccess
	Active          bool
	ClientRequestID string
}

func (d *CommentUsecase) Thank(ctx context.Context, req *CommentThankReq) (bool, error) {
	if !req.Active {
		return false, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return d.thankUsecase.Thank(ctx, req.Access, req.CommentID, req.ClientRequestID)
}

type commentInteraction struct {
	commentID int64
	access    *model.ContentAccess
	active    bool
	action    enum.CommentAction
	stats     repo.CommentStatUpdate
	event     func(int64, int64) *commonenums.Event
}

func (d *CommentUsecase) changeInteraction(ctx context.Context, interaction commentInteraction) (bool, error) {
	access, err := interaction.access.Normalize("")
	if err != nil {
		return false, err
	}
	var outboxEvent *repo.OutboxEvent
	err = d.tx(ctx, func(ctx context.Context) error {
		comment, err := d.commentRepo.Get(ctx, &repo.CommentGetReq{Filter: &model.CommentFilter{CommentID: new(interaction.commentID)}})
		if err != nil {
			return err
		}
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: new(comment.ArticleID)}})
		if err != nil {
			return err
		}
		if err = d.commentAccessUsecase.CanInteract(access, comment, article); err != nil {
			return err
		}
		if interaction.active {
			created, err := d.commentActionRecordRepo.Save(ctx, &model.CommentActionRecord{CommentID: interaction.commentID, UserID: access.ActorUserID, Type: interaction.action, Status: enum.ContentActionStatusPublished})
			if err != nil || !created {
				return err
			}
			if err = d.commentRepo.AddStats(ctx, &repo.CommentAddStatsReq{CommentID: interaction.commentID, Stats: interaction.stats}); err != nil {
				return err
			}
			outboxEvent, err = d.outboxRepo.Save(ctx, interaction.event(access.ActorUserID, interaction.commentID))
			return err
		}
		deleted, err := d.commentActionRecordRepo.Delete(ctx, &repo.CommentActionRecordDeleteReq{CommentID: interaction.commentID, UserID: access.ActorUserID, Action: interaction.action})
		if err != nil || deleted == 0 {
			return err
		}
		return d.commentRepo.AddStats(ctx, &repo.CommentAddStatsReq{CommentID: interaction.commentID, Stats: d.negateCommentStats(interaction.stats)})
	})
	if err != nil {
		return false, err
	}
	if outboxEvent != nil {
		if _, publishErr := d.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{ID: outboxEvent.ID}); publishErr != nil {
			d.log.WarnContext(ctx, "publish content outbox event failed", slog.Int64("outbox_id", outboxEvent.ID), slog.Any("err", publishErr))
		}
	}
	return interaction.active, nil
}

func (*CommentUsecase) negateCommentStats(stats repo.CommentStatUpdate) repo.CommentStatUpdate {
	stats.ThankCount = -stats.ThankCount
	stats.LikeCount = -stats.LikeCount
	stats.ReplyCount = -stats.ReplyCount
	return stats
}

func (*CommentUsecase) commentLikedEvent(userID, commentID int64) *commonenums.Event {
	return &commonenums.Event{Type: commonenums.EventType_EVENT_TYPE_COMMENT_LIKED, Subject: commonenums.EventSubject_EVENT_SUBJECT_COMMENT_LIKED, Payload: &commonenums.Event_CommentLiked{CommentLiked: &commonenums.CommentLikedPayload{SenderId: userID, CommentId: commentID}}}
}

func (*CommentUsecase) commentThankedEvent(userID, commentID int64) *commonenums.Event {
	return &commonenums.Event{Type: commonenums.EventType_EVENT_TYPE_COMMENT_THANKED, Subject: commonenums.EventSubject_EVENT_SUBJECT_COMMENT_THANKED, Payload: &commonenums.Event_CommentThanked{CommentThanked: &commonenums.CommentThankedPayload{SenderId: userID, CommentId: commentID}}}
}

type CommentMapViewerActionStatesReq struct {
	CommentIDs []int64
	UserID     int64
}

func (d *CommentUsecase) MapViewerActionStates(ctx context.Context, req *CommentMapViewerActionStatesReq) (map[int64]*model.
	CommentViewerActionState, error) {
	commentIDs := req.CommentIDs
	userID := req.UserID
	commentIDs = lo.Uniq(commentIDs)
	if len(commentIDs) == 0 {
		return map[int64]*model.CommentViewerActionState{}, nil
	}
	states := lo.SliceToMap(commentIDs, func(commentID int64) (int64, *model.CommentViewerActionState) {
		return commentID, &model.CommentViewerActionState{}
	})
	records, err := d.commentActionRecordRepo.List(ctx, &repo.CommentActionRecordReq{
		CommentIDs: commentIDs,
		UserID:     new(userID),
		Status:     new(enum.ContentActionStatusPublished),
		Types: []enum.CommentAction{
			enum.CommentActionLike,
			enum.CommentActionThank,
		},
	})
	if err != nil {
		return nil, err
	}
	for _, record := range records {
		state := states[record.CommentID]
		if state == nil {
			state = &model.CommentViewerActionState{}
			states[record.CommentID] = state
		}
		switch record.Type {
		case enum.CommentActionLike:
			state.Liked = true
		case enum.CommentActionThank:
			state.Thanked = true
		}
	}
	return states, nil
}
