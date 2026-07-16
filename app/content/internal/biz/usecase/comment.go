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

	"github.com/samber/lo"
	"log/slog"
)

type CommentUsecase struct {
	log *slog.Logger
	tx  base.Tx

	commentRepo             repo.CommentRepo
	commentActionRecordRepo repo.CommentActionRecordRepo
	articleRepo             repo.ArticleRepo
	outboxRepo              repo.OutboxEventRepo
	moderationRecordRepo    repo.ContentModerationRecordRepo
}

func NewCommentUsecase(
	logger *slog.Logger,
	tx base.Tx,
	commentRepo repo.CommentRepo,
	commentActionRecordRepo repo.CommentActionRecordRepo,
	articleRepo repo.ArticleRepo,
	outboxRepo repo.OutboxEventRepo,
	moderationRecordRepo repo.ContentModerationRecordRepo,
) *CommentUsecase {
	return &CommentUsecase{
		log:                     logger,
		tx:                      tx,
		commentRepo:             commentRepo,
		commentActionRecordRepo: commentActionRecordRepo,
		articleRepo:             articleRepo,
		outboxRepo:              outboxRepo,
		moderationRecordRepo:    moderationRecordRepo,
	}
}

type CommentAddReq struct {
	Comment *model.Comment
}

type CommentAddResponse struct {
	Comment *model.Comment
}

func (d *CommentUsecase) Add(ctx context.Context, req *CommentAddReq) (*CommentAddResponse, error) {
	comment := req.Comment
	var c *model.Comment
	err := d.tx(ctx, func(ctx context.Context) error {
		articleResponse, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: new(comment.ArticleID),
		})
		article := articleResponse.Article
		if err != nil {
			return err
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished ||
			article.Visibility != enum.ArticleVisibilityPublic ||
			article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		if !article.Commentable {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_COMMENTABLE)
		}

		replyComment := &model.Comment{}
		var parentID *int64
		if comment.ReplyID != nil {
			commentStatus := enum.ContentRestrictionNone
			replyCommentResponse, replyCommentErr := d.commentRepo.Get(ctx, &repo.CommentGetReq{
				CommentId:   comment.ReplyID,
				ArticleId:   new(comment.ArticleID),
				Restriction: &commentStatus,
			})
			if replyCommentErr != nil {
				return replyCommentErr
			}
			replyComment = replyCommentResponse.Comment
			if replyComment.ParentID == nil {
				parentID = new(replyComment.ID)
			} else {
				parentID = replyComment.ParentID
			}
			_, err = d.commentRepo.AddStats(ctx, &repo.CommentAddStatsReq{CommentID: *parentID, Stats: repo.CommentStatUpdate{ReplyCount: 1}})
			if err != nil {
				return err
			}
		}

		_, err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{ArticleID: article.ID, Stats: repo.ArticleStatUpdate{ReplyCount: 1}})
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
			CreatedBy:   comment.CreatedBy,
			UpdatedBy:   comment.UpdatedBy,
		}
		save.FormatContent()
		commentResponse, commentErr := d.commentRepo.Save(ctx, &repo.CommentSaveReq{Comment: save})
		if commentErr != nil {
			return commentErr
		}
		c = commentResponse.Comment
		if err != nil {
			return err
		}
		c.ReplyUserID = replyComment.CreatedBy
		_, err = d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_COMMENT_PUBLISHED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_COMMENT_PUBLISHED,
				Payload: &commonenums.Event_CommentPublished{
					CommentPublished: &commonenums.CommentPublishedPayload{
						CommentId: c.ID,
					},
				},
			},
		})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &CommentAddResponse{Comment: c}, nil
}

type CommentPageReq struct {
	Page         *base.PageRequest
	CommentID    *int64
	ParentID     *int64
	ReplyID      *int64
	ArticleID    *int64
	CreatedBy    *int64
	Restriction  *enum.ContentRestriction
	Restrictions []enum.ContentRestriction
	Level        *int32
	Order        *enum.CommentOrder
}

type CommentPageResponse struct {
	Page *base.PageResponse
	Rows []*model.Comment
}

func (d *CommentUsecase) Page(ctx context.Context, req *CommentPageReq) (*CommentPageResponse, error) {
	if req == nil {
		req = &CommentPageReq{}
	}
	pageResponse, err := d.commentRepo.Page(ctx, &repo.CommentGetReq{
		Page:         req.Page,
		CommentId:    req.CommentID,
		ParentId:     req.ParentID,
		ReplyId:      req.ReplyID,
		ArticleId:    req.ArticleID,
		CreatedBy:    req.CreatedBy,
		Restriction:  req.Restriction,
		Restrictions: req.Restrictions,
		Level:        req.Level,
		Order:        req.Order,
	})
	if err != nil {
		return nil, err
	}
	return &CommentPageResponse{Page: pageResponse.Page, Rows: pageResponse.Rows}, nil
}

type CommentListReplyPreviewsReq struct {
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

type CommentListReplyPreviewsResponse struct {
	Rows []*CommentChildPreview
}

func (d *CommentUsecase) ListReplyPreviews(ctx context.Context, req *CommentListReplyPreviewsReq) (*CommentListReplyPreviewsResponse, error) {
	articleID := req.ArticleID
	parentIDs := req.ParentIDs
	limitPerParent := req.LimitPerParent
	restriction := req.Restriction
	restrictions := req.Restrictions
	order := req.Order
	if _, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleID)}); err != nil {
		return nil, err
	}
	previews, err := d.commentRepo.ListReplyPreviews(ctx, &repo.CommentReplyPreviewReq{
		ArticleId:      articleID,
		ParentIds:      parentIDs,
		LimitPerParent: limitPerParent,
		Restriction:    restriction,
		Restrictions:   restrictions,
		Order:          order,
	})
	if err != nil {
		return nil, err
	}
	rows := lo.Map(previews.Rows, func(preview *repo.CommentReplyPreview, _ int) *CommentChildPreview {
		return &CommentChildPreview{ParentID: preview.ParentId, Rows: preview.Rows}
	})
	return &CommentListReplyPreviewsResponse{Rows: rows}, nil
}

type CommentMapArticleLastCommentsReq struct {
	ArticleIDs []int64
}

type CommentMapArticleLastCommentsResponse struct {
	Comments map[int64]*model.Comment
}

func (d *CommentUsecase) MapArticleLastComments(ctx context.Context, req *CommentMapArticleLastCommentsReq) (*CommentMapArticleLastCommentsResponse, error) {
	articleIds := req.ArticleIDs
	if len(articleIds) == 0 {
		return &CommentMapArticleLastCommentsResponse{Comments: map[int64]*model.Comment{}}, nil
	}
	comments, err := d.commentRepo.MapArticleLastComments(ctx, &repo.CommentGetReq{
		ArticleIds: lo.Uniq(articleIds),
		Restrictions: []enum.ContentRestriction{
			enum.ContentRestrictionNone,
			enum.ContentRestrictionLocked,
		},
	})
	if err != nil {
		return nil, err
	}
	return &CommentMapArticleLastCommentsResponse{Comments: comments.Rows}, nil
}

type CommentHideReq struct {
	CommentID int64
	UserID    int64
	Reason    *string
}

func (d *CommentUsecase) Hide(ctx context.Context, req *CommentHideReq) error {
	commentId := req.CommentID
	userId := req.UserID
	reason := req.Reason
	return d.updateRestriction(ctx, &commentUpdateRestrictionReq{CommentID: commentId, Restriction: enum.ContentRestrictionHidden, UserID: userId, Action: enum.ContentModerationActionHide, Reason: reason})
}

type CommentUnhideReq struct {
	CommentID int64
	UserID    int64
	Reason    *string
}

func (d *CommentUsecase) Unhide(ctx context.Context, req *CommentUnhideReq) error {
	commentId := req.CommentID
	userId := req.UserID
	reason := req.Reason
	return d.updateRestriction(ctx, &commentUpdateRestrictionReq{CommentID: commentId, Restriction: enum.ContentRestrictionNone, UserID: userId, Action: enum.ContentModerationActionUnhide, Reason: reason})
}

type CommentLockReq struct {
	CommentID int64
	UserID    int64
	Reason    *string
}

func (d *CommentUsecase) Lock(ctx context.Context, req *CommentLockReq) error {
	commentId := req.CommentID
	userId := req.UserID
	reason := req.Reason
	return d.updateRestriction(ctx, &commentUpdateRestrictionReq{CommentID: commentId, Restriction: enum.ContentRestrictionLocked, UserID: userId, Action: enum.ContentModerationActionLock, Reason: reason})
}

type CommentUnlockReq struct {
	CommentID int64
	UserID    int64
	Reason    *string
}

func (d *CommentUsecase) Unlock(ctx context.Context, req *CommentUnlockReq) error {
	commentId := req.CommentID
	userId := req.UserID
	reason := req.Reason
	return d.updateRestriction(ctx, &commentUpdateRestrictionReq{CommentID: commentId, Restriction: enum.ContentRestrictionNone, UserID: userId, Action: enum.ContentModerationActionUnlock, Reason: reason})
}

type commentUpdateRestrictionReq struct {
	CommentID   int64
	Restriction enum.ContentRestriction
	UserID      int64
	Action      enum.ContentModerationAction
	Reason      *string
}

func (d *CommentUsecase) updateRestriction(ctx context.Context, req *commentUpdateRestrictionReq) error {
	commentId := req.CommentID
	restriction := req.Restriction
	userId := req.UserID
	action := req.Action
	reason := req.Reason
	return d.tx(ctx, func(ctx context.Context) error {
		commentResponse, err := d.commentRepo.Get(ctx, &repo.CommentGetReq{CommentId: new(commentId)})
		comment := commentResponse.Comment
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
		if _, err := d.commentRepo.UpdateRestriction(ctx, &repo.CommentUpdateRestrictionReq{CommentID: commentId, Restriction: restriction, UpdatedBy: userId}); err != nil {
			return err
		}
		if _, err := d.moderationRecordRepo.Save(ctx, &repo.ContentModerationRecordSaveReq{Record: &model.ContentModerationRecord{
			Target:     enum.ContentModerationTargetComment,
			TargetID:   commentId,
			Action:     action,
			Reason:     reason,
			OperatorID: userId,
		}}); err != nil {
			return err
		}
		_, err = d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_COMMENT_STATUS_UPDATED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_COMMENT_STATUS_UPDATED,
				Payload: &commonenums.Event_CommentStatusUpdated{
					CommentStatusUpdated: &commonenums.CommentStatusUpdatedPayload{
						SenderId:    userId,
						CommentId:   commentId,
						Action:      string(action),
						Restriction: string(restriction),
						Reason:      reason,
					},
				},
			},
		})
		return err
	})
}

type CommentLikeReq struct {
	CommentID int64
	UserID    int64
	Active    bool
}

type CommentLikeResponse struct {
	Liked bool
}

func (d *CommentUsecase) Like(ctx context.Context, req *CommentLikeReq) (*CommentLikeResponse, error) {
	commentId := req.CommentID
	userId := req.UserID
	active := req.Active
	err := d.tx(ctx, func(ctx context.Context) error {
		commentResponse, err := d.commentRepo.Get(ctx, &repo.CommentGetReq{
			CommentId: &commentId,
		})
		comment := commentResponse.Comment
		if err != nil {
			return err
		}
		if comment.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_NOT_FOUND)
		}
		articleResponse, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: new(comment.ArticleID),
		})
		article := articleResponse.Article
		if err != nil {
			return err
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished ||
			article.Visibility != enum.ArticleVisibilityPublic ||
			article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_NOT_FOUND)
		}
		if active {
			createdResponse, err := d.commentActionRecordRepo.Save(ctx, &repo.CommentActionRecordSaveReq{Record: &model.CommentActionRecord{
				CommentID: commentId,
				UserID:    userId,
				Type:      enum.CommentActionLike,
			}})
			if err != nil {
				return err
			}
			if !createdResponse.Created {
				return nil
			}
			if _, err = d.commentRepo.AddStats(ctx, &repo.CommentAddStatsReq{CommentID: commentId, Stats: repo.CommentStatUpdate{LikeCount: 1}}); err != nil {
				return err
			}
			_, err = d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
				Event: &commonenums.Event{
					Type:    commonenums.EventType_EVENT_TYPE_COMMENT_LIKED,
					Subject: commonenums.EventSubject_EVENT_SUBJECT_COMMENT_LIKED,
					Payload: &commonenums.Event_CommentLiked{
						CommentLiked: &commonenums.CommentLikedPayload{
							SenderId:  userId,
							CommentId: commentId,
						},
					},
				},
			})
		}

		deletedResponse, err := d.commentActionRecordRepo.Delete(ctx, &repo.CommentActionRecordDeleteReq{CommentID: commentId, UserID: userId, Action: enum.CommentActionLike})
		if err != nil {
			return err
		}
		if deletedResponse.Deleted == 0 {
			return nil
		}
		_, err = d.commentRepo.AddStats(ctx, &repo.CommentAddStatsReq{CommentID: commentId, Stats: repo.CommentStatUpdate{LikeCount: -1}})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &CommentLikeResponse{Liked: active}, nil
}

type CommentThankReq struct {
	CommentID int64
	UserID    int64
	Active    bool
}

type CommentThankResponse struct {
	Thanked bool
}

func (d *CommentUsecase) Thank(ctx context.Context, req *CommentThankReq) (*CommentThankResponse, error) {
	commentId := req.CommentID
	userId := req.UserID
	active := req.Active
	err := d.tx(ctx, func(ctx context.Context) error {
		commentResponse, err := d.commentRepo.Get(ctx, &repo.CommentGetReq{
			CommentId: &commentId,
		})
		comment := commentResponse.Comment
		if err != nil {
			return err
		}
		if comment.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_NOT_FOUND)
		}
		articleResponse, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: new(comment.ArticleID),
		})
		article := articleResponse.Article
		if err != nil {
			return err
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished ||
			article.Visibility != enum.ArticleVisibilityPublic ||
			article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_NOT_FOUND)
		}
		if active {
			createdResponse, err := d.commentActionRecordRepo.Save(ctx, &repo.CommentActionRecordSaveReq{Record: &model.CommentActionRecord{
				CommentID: commentId,
				UserID:    userId,
				Type:      enum.CommentActionThank,
			}})
			if err != nil {
				return err
			}
			if !createdResponse.Created {
				return nil
			}
			if _, err = d.commentRepo.AddStats(ctx, &repo.CommentAddStatsReq{CommentID: commentId, Stats: repo.CommentStatUpdate{ThankCount: 1}}); err != nil {
				return err
			}
			_, err = d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
				Event: &commonenums.Event{
					Type:    commonenums.EventType_EVENT_TYPE_COMMENT_THANKED,
					Subject: commonenums.EventSubject_EVENT_SUBJECT_COMMENT_THANKED,
					Payload: &commonenums.Event_CommentThanked{
						CommentThanked: &commonenums.CommentThankedPayload{
							SenderId:  userId,
							CommentId: commentId,
						},
					},
				},
			})
		}

		deletedResponse, err := d.commentActionRecordRepo.Delete(ctx, &repo.CommentActionRecordDeleteReq{CommentID: commentId, UserID: userId, Action: enum.CommentActionThank})
		if err != nil {
			return err
		}
		if deletedResponse.Deleted == 0 {
			return nil
		}
		_, err = d.commentRepo.AddStats(ctx, &repo.CommentAddStatsReq{CommentID: commentId, Stats: repo.CommentStatUpdate{ThankCount: -1}})
		return err
	})
	if err != nil {
		return nil, err
	}
	return &CommentThankResponse{Thanked: active}, nil
}

type CommentMapViewerActionStatesReq struct {
	CommentIDs []int64
	UserID     int64
}

type CommentMapViewerActionStatesResponse struct {
	States map[int64]*model.CommentViewerActionState
}

func (d *CommentUsecase) MapViewerActionStates(ctx context.Context, req *CommentMapViewerActionStatesReq) (*CommentMapViewerActionStatesResponse, error) {
	commentIds := req.CommentIDs
	userId := req.UserID
	commentIds = lo.Uniq(commentIds)
	if len(commentIds) == 0 {
		return &CommentMapViewerActionStatesResponse{States: map[int64]*model.CommentViewerActionState{}}, nil
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
	for _, record := range records.Rows {
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
	return &CommentMapViewerActionStatesResponse{States: states}, nil
}
