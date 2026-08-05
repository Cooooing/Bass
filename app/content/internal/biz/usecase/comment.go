package usecase

import (
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
	Page   *base.PageRequest
	Filter *model.CommentFilter
}

type CommentPageResp struct {
	Page *base.PageResp
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
	Access         *model.ContentAccess
	ArticleID      int64
	ParentIDs      []int64
	LimitPerParent int32
	Restriction    *enum.ContentRestriction
	Restrictions   []enum.ContentRestriction
	Order          *enum.CommentOrder
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
			ArticleID:    new(articleID),
			Restriction:  req.Restriction,
			Restrictions: req.Restrictions,
			Order:        req.Order,
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
			ParentID: preview.ParentId,
			Rows:     preview.Rows,
		}
	})
	return rows, nil
}

func (d *CommentUsecase) MapArticleLastComments(ctx context.Context, articleIDs []int64) (map[int64]*model.
	Comment, error) {
	articleIds := articleIDs
	if len(articleIds) == 0 {
		return map[int64]*model.Comment{}, nil
	}
	comments, err := d.commentRepo.MapArticleLastComments(ctx, &repo.CommentGetReq{
		Filter: &model.CommentFilter{
			ArticleIDs: lo.Uniq(articleIds),
			Restrictions: []enum.ContentRestriction{
				enum.ContentRestrictionNone,
				enum.ContentRestrictionLocked,
			},
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
	commentId := req.CommentID
	access, err := req.Access.Normalize("")
	if err != nil {
		return err
	}
	reason := req.Reason
	return d.updateRestriction(ctx, &commentUpdateRestrictionReq{
		CommentID:   commentId,
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
	commentId := req.CommentID
	access, err := req.Access.Normalize("")
	if err != nil {
		return err
	}
	reason := req.Reason
	return d.updateRestriction(ctx, &commentUpdateRestrictionReq{
		CommentID:   commentId,
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
	commentId := req.CommentID
	access, err := req.Access.Normalize("")
	if err != nil {
		return err
	}
	reason := req.Reason
	return d.updateRestriction(ctx, &commentUpdateRestrictionReq{
		CommentID:   commentId,
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
	commentId := req.CommentID
	access, err := req.Access.Normalize("")
	if err != nil {
		return err
	}
	reason := req.Reason
	return d.updateRestriction(ctx, &commentUpdateRestrictionReq{
		CommentID:   commentId,
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
	commentId := req.CommentID
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
			Filter: &model.CommentFilter{CommentID: new(commentId)},
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
			CommentID:   commentId,
			Restriction: restriction,
			UpdatedBy:   access.ActorUserID,
		}); err != nil {
			return err
		}
		if _, err := d.moderationRecordRepo.Save(ctx, &model.ContentModerationRecord{
			Target:     enum.ContentModerationTargetComment,
			TargetID:   commentId,
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
					CommentId:   commentId,
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
	commentId := req.CommentID
	access, err := req.Access.Normalize("")
	if err != nil {
		return false, err
	}
	userId := access.ActorUserID
	active := req.Active
	var outboxEvent *repo.OutboxEvent
	err = d.tx(ctx, func(ctx context.Context) error {
		commentResp, err := d.commentRepo.Get(ctx, &repo.CommentGetReq{
			Filter: &model.CommentFilter{CommentID: new(commentId)},
		})
		comment := commentResp
		if err != nil {
			return err
		}
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(comment.ArticleID)},
		})
		if err != nil {
			return err
		}
		if err := d.commentAccessUsecase.CanInteract(access, comment, article); err != nil {
			return err
		}
		if active {
			createdResp, err := d.commentActionRecordRepo.Save(ctx, &model.CommentActionRecord{
				CommentID: commentId,
				UserID:    userId,
				Type:      enum.CommentActionLike,
			})
			if err != nil {
				return err
			}
			if !createdResp {
				return nil
			}
			if err = d.commentRepo.AddStats(ctx, &repo.CommentAddStatsReq{
				CommentID: commentId,
				Stats: repo.CommentStatUpdate{
					LikeCount: 1,
				},
			}); err != nil {
				return err
			}
			outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_COMMENT_LIKED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_COMMENT_LIKED,
				Payload: &commonenums.Event_CommentLiked{
					CommentLiked: &commonenums.CommentLikedPayload{
						SenderId:  userId,
						CommentId: commentId,
					},
				},
			})
			return err
		}

		deletedResp, err := d.commentActionRecordRepo.Delete(ctx, &repo.CommentActionRecordDeleteReq{
			CommentID: commentId,
			UserID:    userId,
			Action:    enum.CommentActionLike,
		})
		if err != nil {
			return err
		}
		if deletedResp == 0 {
			return nil
		}
		err = d.commentRepo.AddStats(ctx, &repo.CommentAddStatsReq{
			CommentID: commentId,
			Stats: repo.CommentStatUpdate{
				LikeCount: -1,
			},
		})
		return err
	})
	if err != nil {
		return false, err
	}
	if outboxEvent != nil {
		if _, publishErr := d.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{ID: outboxEvent.ID}); publishErr != nil {
			d.log.WarnContext(ctx, "publish content outbox event failed", slog.Int64("outbox_id", outboxEvent.ID), slog.Any("err", publishErr))
		}
	}
	return active, nil
}

type CommentThankReq struct {
	CommentID int64
	Access    *model.ContentAccess
	Active    bool
}

func (d *CommentUsecase) Thank(ctx context.Context, req *CommentThankReq) (bool, error) {
	commentId := req.CommentID
	access, err := req.Access.Normalize("")
	if err != nil {
		return false, err
	}
	userId := access.ActorUserID
	active := req.Active
	var outboxEvent *repo.OutboxEvent
	err = d.tx(ctx, func(ctx context.Context) error {
		commentResp, err := d.commentRepo.Get(ctx, &repo.CommentGetReq{
			Filter: &model.CommentFilter{CommentID: new(commentId)},
		})
		comment := commentResp
		if err != nil {
			return err
		}
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(comment.ArticleID)},
		})
		if err != nil {
			return err
		}
		if err := d.commentAccessUsecase.CanInteract(access, comment, article); err != nil {
			return err
		}
		if active {
			createdResp, err := d.commentActionRecordRepo.Save(ctx, &model.CommentActionRecord{
				CommentID: commentId,
				UserID:    userId,
				Type:      enum.CommentActionThank,
			})
			if err != nil {
				return err
			}
			if !createdResp {
				return nil
			}
			if err = d.commentRepo.AddStats(ctx, &repo.CommentAddStatsReq{
				CommentID: commentId,
				Stats: repo.CommentStatUpdate{
					ThankCount: 1,
				},
			}); err != nil {
				return err
			}
			outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_COMMENT_THANKED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_COMMENT_THANKED,
				Payload: &commonenums.Event_CommentThanked{
					CommentThanked: &commonenums.CommentThankedPayload{
						SenderId:  userId,
						CommentId: commentId,
					},
				},
			})
			return err
		}

		deletedResp, err := d.commentActionRecordRepo.Delete(ctx, &repo.CommentActionRecordDeleteReq{
			CommentID: commentId,
			UserID:    userId,
			Action:    enum.CommentActionThank,
		})
		if err != nil {
			return err
		}
		if deletedResp == 0 {
			return nil
		}
		err = d.commentRepo.AddStats(ctx, &repo.CommentAddStatsReq{
			CommentID: commentId,
			Stats: repo.CommentStatUpdate{
				ThankCount: -1,
			},
		})
		return err
	})
	if err != nil {
		return false, err
	}
	if outboxEvent != nil {
		if _, publishErr := d.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{ID: outboxEvent.ID}); publishErr != nil {
			d.log.WarnContext(ctx, "publish content outbox event failed", slog.Int64("outbox_id", outboxEvent.ID), slog.Any("err", publishErr))
		}
	}
	return active, nil
}

type CommentMapViewerActionStatesReq struct {
	CommentIDs []int64
	UserID     int64
}

func (d *CommentUsecase) MapViewerActionStates(ctx context.Context, req *CommentMapViewerActionStatesReq) (map[int64]*model.
	CommentViewerActionState, error) {
	commentIds := req.CommentIDs
	userId := req.UserID
	commentIds = lo.Uniq(commentIds)
	if len(commentIds) == 0 {
		return map[int64]*model.CommentViewerActionState{}, nil
	}
	states := lo.SliceToMap(commentIds, func(commentID int64) (int64, *model.CommentViewerActionState) {
		return commentID, &model.CommentViewerActionState{}
	})
	records, err := d.commentActionRecordRepo.List(ctx, &repo.CommentActionRecordReq{
		CommentIds: commentIds,
		UserId:     new(userId),
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
