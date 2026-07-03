package usecase

import (
	cerrors "common/proto/gen/common/errors"
	"context"

	"common/pkg/apperror"
	"common/proto/gen/common"
	commonenums "common/proto/gen/common/enums"
	base "content/internal/biz/base"
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

func (d *CommentUsecase) Add(ctx context.Context, comment *model.Comment) (c *model.Comment, err error) {
	err = d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: new(comment.ArticleID),
		})
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
			replyComment, err = d.commentRepo.Get(ctx, &repo.CommentGetReq{
				CommentId:   comment.ReplyID,
				ArticleId:   new(comment.ArticleID),
				Restriction: &commentStatus,
			})
			if err != nil {
				return err
			}
			if replyComment.ParentID == nil {
				parentID = new(replyComment.ID)
			} else {
				parentID = replyComment.ParentID
			}
			err = d.commentRepo.AddStats(ctx, *parentID, repo.CommentStatUpdate{ReplyCount: 1}, nil)
			if err != nil {
				return err
			}
		}

		err = d.articleRepo.AddStats(ctx, article.ID, repo.ArticleStatUpdate{ReplyCount: 1}, nil)
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
		c, err = d.commentRepo.Save(ctx, save)
		if err != nil {
			return err
		}
		c.ReplyUserID = replyComment.CreatedBy
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
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
	})
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (d *CommentUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.CommentGetReq) (*common.PageReply, []*model.Comment, error) {
	comments, pageReply, err := d.commentRepo.Page(ctx, page, req)
	return pageReply, comments, err
}

func (d *CommentUsecase) ListReplyPreviews(ctx context.Context, articleID int64, parentIDs []int64, limitPerParent int32, restriction *enum.ContentRestriction, restrictions []enum.ContentRestriction, order *enum.CommentOrder) ([]*repo.CommentReplyPreview, error) {
	if _, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleID)}); err != nil {
		return nil, err
	}
	return d.commentRepo.ListReplyPreviews(ctx, &repo.CommentReplyPreviewReq{
		ArticleId:      articleID,
		ParentIds:      parentIDs,
		LimitPerParent: limitPerParent,
		Restriction:    restriction,
		Restrictions:   restrictions,
		Order:          order,
	})
}

func (d *CommentUsecase) MapArticleLastComments(ctx context.Context, articleIds []int64) (map[int64]*model.Comment, error) {
	if len(articleIds) == 0 {
		return map[int64]*model.Comment{}, nil
	}
	return d.commentRepo.MapArticleLastComments(ctx, &repo.CommentGetReq{
		ArticleIds: lo.Uniq(articleIds),
		Restrictions: []enum.ContentRestriction{
			enum.ContentRestrictionNone,
			enum.ContentRestrictionLocked,
		},
	})
}

func (d *CommentUsecase) Hide(ctx context.Context, commentId int64, userId int64, reason *string) error {
	return d.updateRestriction(ctx, commentId, enum.ContentRestrictionHidden, userId, enum.ContentModerationActionHide, reason)
}

func (d *CommentUsecase) Unhide(ctx context.Context, commentId int64, userId int64, reason *string) error {
	return d.updateRestriction(ctx, commentId, enum.ContentRestrictionNone, userId, enum.ContentModerationActionUnhide, reason)
}

func (d *CommentUsecase) Lock(ctx context.Context, commentId int64, userId int64, reason *string) error {
	return d.updateRestriction(ctx, commentId, enum.ContentRestrictionLocked, userId, enum.ContentModerationActionLock, reason)
}

func (d *CommentUsecase) Unlock(ctx context.Context, commentId int64, userId int64, reason *string) error {
	return d.updateRestriction(ctx, commentId, enum.ContentRestrictionNone, userId, enum.ContentModerationActionUnlock, reason)
}

func (d *CommentUsecase) updateRestriction(ctx context.Context, commentId int64, restriction enum.ContentRestriction, userId int64, action enum.ContentModerationAction, reason *string) error {
	return d.tx(ctx, func(ctx context.Context) error {
		comment, err := d.commentRepo.Get(ctx, &repo.CommentGetReq{CommentId: new(commentId)})
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
		if err := d.commentRepo.UpdateRestriction(ctx, commentId, restriction, userId); err != nil {
			return err
		}
		if _, err := d.moderationRecordRepo.Save(ctx, &model.ContentModerationRecord{
			Target:     enum.ContentModerationTargetComment,
			TargetID:   commentId,
			Action:     action,
			Reason:     reason,
			OperatorID: userId,
		}); err != nil {
			return err
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
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
	})
}

func (d *CommentUsecase) Like(ctx context.Context, commentId int64, userId int64, active bool) (bool, error) {
	err := d.tx(ctx, func(ctx context.Context) error {
		comment, err := d.commentRepo.Get(ctx, &repo.CommentGetReq{
			CommentId: &commentId,
		})
		if err != nil {
			return err
		}
		if comment.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_NOT_FOUND)
		}
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: new(comment.ArticleID),
		})
		if err != nil {
			return err
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished ||
			article.Visibility != enum.ArticleVisibilityPublic ||
			article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_NOT_FOUND)
		}
		if active {
			created, err := d.commentActionRecordRepo.Save(ctx, &model.CommentActionRecord{
				CommentID: commentId,
				UserID:    userId,
				Type:      enum.CommentActionLike,
			})
			if err != nil {
				return err
			}
			if !created {
				return nil
			}
			if err = d.commentRepo.AddStats(ctx, commentId, repo.CommentStatUpdate{LikeCount: 1}, nil); err != nil {
				return err
			}
			return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
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

		deleted, err := d.commentActionRecordRepo.Delete(ctx, commentId, userId, enum.CommentActionLike)
		if err != nil {
			return err
		}
		if deleted == 0 {
			return nil
		}
		return d.commentRepo.AddStats(ctx, commentId, repo.CommentStatUpdate{LikeCount: -1}, nil)
	})
	return active, err
}

func (d *CommentUsecase) Thank(ctx context.Context, commentId int64, userId int64, active bool) (bool, error) {
	err := d.tx(ctx, func(ctx context.Context) error {
		comment, err := d.commentRepo.Get(ctx, &repo.CommentGetReq{
			CommentId: &commentId,
		})
		if err != nil {
			return err
		}
		if comment.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_NOT_FOUND)
		}
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: new(comment.ArticleID),
		})
		if err != nil {
			return err
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished ||
			article.Visibility != enum.ArticleVisibilityPublic ||
			article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_NOT_FOUND)
		}
		if active {
			created, err := d.commentActionRecordRepo.Save(ctx, &model.CommentActionRecord{
				CommentID: commentId,
				UserID:    userId,
				Type:      enum.CommentActionThank,
			})
			if err != nil {
				return err
			}
			if !created {
				return nil
			}
			if err = d.commentRepo.AddStats(ctx, commentId, repo.CommentStatUpdate{ThankCount: 1}, nil); err != nil {
				return err
			}
			return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
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

		deleted, err := d.commentActionRecordRepo.Delete(ctx, commentId, userId, enum.CommentActionThank)
		if err != nil {
			return err
		}
		if deleted == 0 {
			return nil
		}
		return d.commentRepo.AddStats(ctx, commentId, repo.CommentStatUpdate{ThankCount: -1}, nil)
	})
	return active, err
}

func (d *CommentUsecase) MapViewerActionStates(ctx context.Context, commentIds []int64, userId int64) (map[int64]*model.CommentViewerActionState, error) {
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
