package usecase

import (
	cerrors "common/proto/gen/common/errors"
	"context"
	"time"

	"common/pkg/apperror"
	"common/proto/gen/common"
	commonenums "common/proto/gen/common/enums"
	base "content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/conf"
	"content/internal/enum"

	"github.com/samber/lo"
)

type ArticleUsecase struct {
	conf *conf.Bootstrap
	tx   base.Tx

	articleRepo          repo.ArticleRepo
	postscriptRepo       repo.ArticlePostscriptRepo
	actionRecordRepo     repo.ArticleActionRecordRepo
	commentRepo          repo.CommentRepo
	tagRepo              repo.TagRepo
	outboxRepo           repo.OutboxEventRepo
	moderationRecordRepo repo.ContentModerationRecordRepo
}

func NewArticleUsecase(
	conf *conf.Bootstrap,
	tx base.Tx,
	articleRepo repo.ArticleRepo,
	postscriptRepo repo.ArticlePostscriptRepo,
	actionRecordRepo repo.ArticleActionRecordRepo,
	commentRepo repo.CommentRepo,
	tagRepo repo.TagRepo,
	outboxRepo repo.OutboxEventRepo,
	moderationRecordRepo repo.ContentModerationRecordRepo,
) *ArticleUsecase {
	return &ArticleUsecase{
		conf:                 conf,
		tx:                   tx,
		articleRepo:          articleRepo,
		postscriptRepo:       postscriptRepo,
		actionRecordRepo:     actionRecordRepo,
		commentRepo:          commentRepo,
		tagRepo:              tagRepo,
		outboxRepo:           outboxRepo,
		moderationRecordRepo: moderationRecordRepo,
	}
}

func (d *ArticleUsecase) Add(ctx context.Context, article *model.Article, tagIds []int64) (*model.Article, error) {
	var (
		save *model.Article
		err  error
	)
	err = d.tx(ctx, func(ctx context.Context) error {
		bindTagIDs := lo.Uniq(tagIds)
		if len(bindTagIDs) > 0 {
			tagStatus := enum.TagStatusEnabled
			count, err := d.tagRepo.Count(ctx, &repo.TagGetReq{TagIds: bindTagIDs, Status: &tagStatus})
			if err != nil {
				return err
			}
			if count != len(bindTagIDs) {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_NOT_FOUND)
			}
		}
		now := time.Now()
		article.PublishStatus = enum.ArticlePublishStatusDraft
		if article.Visibility == "" {
			article.Visibility = enum.ArticleVisibilityPublic
		}
		article.Restriction = enum.ContentRestrictionNone
		article.EditedAt = new(now)
		article.FormatContent()
		save, err = d.articleRepo.Save(ctx, article)
		if err != nil {
			return err
		}
		if len(bindTagIDs) > 0 {
			if err = d.articleRepo.ReplaceTags(ctx, save.ID, bindTagIDs); err != nil {
				return err
			}
		}
		return nil
	})
	return save, err
}

func (d *ArticleUsecase) AddPostscript(ctx context.Context, articleId int64, content string, userId int64) (*model.ArticlePostscript, error) {
	var save *model.ArticlePostscript
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: &articleId,
		})
		if err != nil {
			return err
		}
		if !d.isAuthor(article, userId) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished || article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		postscript := &model.ArticlePostscript{
			ArticleID:   articleId,
			Content:     content,
			Restriction: enum.ContentRestrictionNone,
			CreatedBy:   new(userId),
			UpdatedBy:   new(userId),
		}
		postscript.FormatContent()
		save, err = d.postscriptRepo.Save(ctx, postscript)
		if err != nil {
			return err
		}
		if err = d.articleRepo.UpdateHasPostscript(ctx, articleId, true, userId); err != nil {
			return err
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_POSTSCRIPT_ADDED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_POSTSCRIPT_ADDED,
				Payload: &commonenums.Event_ArticlePostscriptAdded{
					ArticlePostscriptAdded: &commonenums.ArticlePostscriptAddedPayload{
						SenderId:     userId,
						ArticleId:    articleId,
						PostscriptId: save.ID,
					},
				},
			},
		})
	})
	return save, err
}

func (d *ArticleUsecase) ListPostscripts(ctx context.Context, articleId int64) ([]*model.ArticlePostscript, error) {
	if _, err := d.Get(ctx, articleId); err != nil {
		return nil, err
	}
	restriction := enum.ContentRestrictionNone
	return d.postscriptRepo.List(ctx, &repo.ArticlePostscriptGetReq{
		ArticleID:   new(articleId),
		Restriction: &restriction,
	})
}

func (d *ArticleUsecase) Update(ctx context.Context, article *model.Article, tagIds []int64) (*model.Article, error) {
	var (
		save *model.Article
		err  error
	)
	err = d.tx(ctx, func(ctx context.Context) error {
		current, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: new(article.ID),
		})
		if err != nil {
			return err
		}
		if article.UpdatedBy == nil || !d.isAuthor(current, *article.UpdatedBy) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		if current.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		switch current.PublishStatus {
		case enum.ArticlePublishStatusDraft:
		case enum.ArticlePublishStatusPublished:
			editWindow := 10 * time.Minute
			maxViewCount := int32(100)
			if d.conf != nil && d.conf.GetBusiness() != nil && d.conf.GetBusiness().GetArticle() != nil {
				articleConf := d.conf.GetBusiness().GetArticle()
				if articleConf.GetPublishedEditWindow() != nil && articleConf.GetPublishedEditWindow().AsDuration() > 0 {
					editWindow = articleConf.GetPublishedEditWindow().AsDuration()
				}
				if articleConf.GetPublishedEditMaxViewCount() > 0 {
					maxViewCount = articleConf.GetPublishedEditMaxViewCount()
				}
			}
			if current.ViewCount >= maxViewCount || current.PublishedAt == nil || time.Since(*current.PublishedAt) > editWindow {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
			}
		default:
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}

		bindTagIDs := lo.Uniq(tagIds)
		if len(bindTagIDs) > 0 {
			tagStatus := enum.TagStatusEnabled
			count, err := d.tagRepo.Count(ctx, &repo.TagGetReq{TagIds: bindTagIDs, Status: &tagStatus})
			if err != nil {
				return err
			}
			if count != len(bindTagIDs) {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_TAG_NOT_FOUND)
			}
		}

		now := time.Now()
		article.PublishStatus = current.PublishStatus
		article.Visibility = current.Visibility
		article.Restriction = current.Restriction
		article.PublishedAt = current.PublishedAt
		article.EditedAt = new(now)
		article.FormatContent()
		save, err = d.articleRepo.Update(ctx, article)
		if err != nil {
			return err
		}
		if err = d.articleRepo.ReplaceTags(ctx, save.ID, bindTagIDs); err != nil {
			return err
		}
		return nil
	})
	return save, err
}

func (d *ArticleUsecase) Reward(ctx context.Context, articleId int64, userId int64, points int32) error {
	return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_REWARD_NOT_IMPLEMENTED)
}

func (d *ArticleUsecase) Like(ctx context.Context, articleId int64, userId int64, active bool) (bool, error) {
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		if err != nil {
			return err
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished || article.Visibility != enum.ArticleVisibilityPublic || article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}

		if active {
			created, err := d.actionRecordRepo.Save(ctx, &model.ArticleActionRecord{
				ArticleID: articleId,
				UserID:    userId,
				Type:      enum.ArticleActionLike,
			})
			if err != nil {
				return err
			}
			if !created {
				return nil
			}
			if err = d.articleRepo.AddStats(ctx, articleId, repo.ArticleStatUpdate{LikeCount: 1}, nil); err != nil {
				return err
			}
			return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
				Event: &commonenums.Event{
					Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_LIKED,
					Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_LIKED,
					Payload: &commonenums.Event_ArticleLiked{
						ArticleLiked: &commonenums.ArticleLikedPayload{
							SenderId:  userId,
							ArticleId: articleId,
						},
					},
				},
			})
		}

		deleted, err := d.actionRecordRepo.Delete(ctx, articleId, userId, enum.ArticleActionLike)
		if err != nil {
			return err
		}
		if deleted == 0 {
			return nil
		}
		return d.articleRepo.AddStats(ctx, articleId, repo.ArticleStatUpdate{LikeCount: -1}, nil)
	})
	return active, err
}

func (d *ArticleUsecase) Thank(ctx context.Context, articleId int64, userId int64, active bool) (bool, error) {
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		if err != nil {
			return err
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished || article.Visibility != enum.ArticleVisibilityPublic || article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}

		if active {
			created, err := d.actionRecordRepo.Save(ctx, &model.ArticleActionRecord{
				ArticleID: articleId,
				UserID:    userId,
				Type:      enum.ArticleActionThank,
			})
			if err != nil {
				return err
			}
			if !created {
				return nil
			}
			if err = d.articleRepo.AddStats(ctx, articleId, repo.ArticleStatUpdate{ThankCount: 1}, nil); err != nil {
				return err
			}
			return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
				Event: &commonenums.Event{
					Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_THANKED,
					Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_THANKED,
					Payload: &commonenums.Event_ArticleThanked{
						ArticleThanked: &commonenums.ArticleThankedPayload{
							SenderId:  userId,
							ArticleId: articleId,
						},
					},
				},
			})
		}

		deleted, err := d.actionRecordRepo.Delete(ctx, articleId, userId, enum.ArticleActionThank)
		if err != nil {
			return err
		}
		if deleted == 0 {
			return nil
		}
		return d.articleRepo.AddStats(ctx, articleId, repo.ArticleStatUpdate{ThankCount: -1}, nil)
	})
	return active, err
}

func (d *ArticleUsecase) Collect(ctx context.Context, articleId int64, userId int64, active bool) (bool, error) {
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		if err != nil {
			return err
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished || article.Visibility != enum.ArticleVisibilityPublic || article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}

		if active {
			created, err := d.actionRecordRepo.Save(ctx, &model.ArticleActionRecord{
				ArticleID: articleId,
				UserID:    userId,
				Type:      enum.ArticleActionCollect,
			})
			if err != nil {
				return err
			}
			if !created {
				return nil
			}
			if err = d.articleRepo.AddStats(ctx, articleId, repo.ArticleStatUpdate{CollectCount: 1}, nil); err != nil {
				return err
			}
			return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
				Event: &commonenums.Event{
					Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_COLLECTED,
					Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_COLLECTED,
					Payload: &commonenums.Event_ArticleCollected{
						ArticleCollected: &commonenums.ArticleCollectedPayload{
							SenderId:  userId,
							ArticleId: articleId,
						},
					},
				},
			})
		}

		deleted, err := d.actionRecordRepo.Delete(ctx, articleId, userId, enum.ArticleActionCollect)
		if err != nil {
			return err
		}
		if deleted == 0 {
			return nil
		}
		return d.articleRepo.AddStats(ctx, articleId, repo.ArticleStatUpdate{CollectCount: -1}, nil)
	})
	return active, err
}

func (d *ArticleUsecase) Watch(ctx context.Context, articleId int64, userId int64, active bool) (bool, error) {
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		if err != nil {
			return err
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished || article.Visibility != enum.ArticleVisibilityPublic || article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}

		if active {
			created, err := d.actionRecordRepo.Save(ctx, &model.ArticleActionRecord{
				ArticleID: articleId,
				UserID:    userId,
				Type:      enum.ArticleActionWatch,
			})
			if err != nil {
				return err
			}
			if !created {
				return nil
			}
			if err = d.articleRepo.AddStats(ctx, articleId, repo.ArticleStatUpdate{WatchCount: 1}, nil); err != nil {
				return err
			}
			return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
				Event: &commonenums.Event{
					Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_WATCHED,
					Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_WATCHED,
					Payload: &commonenums.Event_ArticleWatched{
						ArticleWatched: &commonenums.ArticleWatchedPayload{
							SenderId:  userId,
							ArticleId: articleId,
						},
					},
				},
			})
		}

		deleted, err := d.actionRecordRepo.Delete(ctx, articleId, userId, enum.ArticleActionWatch)
		if err != nil {
			return err
		}
		if deleted == 0 {
			return nil
		}
		return d.articleRepo.AddStats(ctx, articleId, repo.ArticleStatUpdate{WatchCount: -1}, nil)
	})
	return active, err
}

func (d *ArticleUsecase) View(ctx context.Context, articleId int64, viewerUserId *int64) error {
	return d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		if err != nil {
			return err
		}
		switch article.PublishStatus {
		case enum.ArticlePublishStatusPublished, enum.ArticlePublishStatusArchived:
		default:
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		if article.Restriction == enum.ContentRestrictionHidden {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		return d.articleRepo.AddStats(ctx, articleId, repo.ArticleStatUpdate{ViewCount: 1}, nil)
	})
}

func (d *ArticleUsecase) Publish(ctx context.Context, articleId int64, userId int64, visibility enum.ArticleVisibility) error {
	return d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		if err != nil {
			return err
		}
		if !d.isAuthor(article, userId) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		if article.PublishStatus != enum.ArticlePublishStatusDraft || article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		switch visibility {
		case enum.ArticleVisibilityPublic, enum.ArticleVisibilityPrivate:
		default:
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		now := time.Now()
		if err = d.articleRepo.UpdatePublishStatus(ctx, articleId, enum.ArticlePublishStatusPublished, visibility, new(now), userId); err != nil {
			return err
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_PUBLISHED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_PUBLISHED,
				Payload: &commonenums.Event_ArticlePublished{
					ArticlePublished: &commonenums.ArticlePublishedPayload{
						ArticleId: articleId,
					},
				},
			},
		})
	})
}

func (d *ArticleUsecase) AcceptAnswer(ctx context.Context, articleId int64, commentId int64, userId int64) error {
	return d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: new(articleId),
		})
		if err != nil {
			return err
		}
		if !d.isAuthor(article, userId) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished ||
			article.Visibility != enum.ArticleVisibilityPublic ||
			article.Restriction != enum.ContentRestrictionNone ||
			article.Type != enum.ArticleTypeQA ||
			article.AcceptedAnswerID != nil {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		commentStatus := enum.ContentRestrictionNone
		exist, err := d.commentRepo.Exist(ctx, &repo.CommentGetReq{
			CommentId:   &commentId,
			ArticleId:   &articleId,
			Restriction: &commentStatus,
		})
		if err != nil {
			return err
		}
		if !exist {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_NOT_FOUND)
		}
		if _, err = d.articleRepo.UpdateAcceptedAnswerID(ctx, articleId, commentId, userId); err != nil {
			return err
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_ACCEPTED_ANSWER,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_ACCEPTED_ANSWER,
				Payload: &commonenums.Event_ArticleAcceptedAnswer{
					ArticleAcceptedAnswer: &commonenums.ArticleAcceptedAnswerPayload{
						SenderId:  userId,
						ArticleId: articleId,
						CommentId: commentId,
					},
				},
			},
		})
	})
}

func (d *ArticleUsecase) MakePrivate(ctx context.Context, articleId int64, userId int64) error {
	return d.updateVisibility(ctx, articleId, enum.ArticleVisibilityPrivate, userId)
}

func (d *ArticleUsecase) MakePublic(ctx context.Context, articleId int64, userId int64) error {
	return d.updateVisibility(ctx, articleId, enum.ArticleVisibilityPublic, userId)
}

func (d *ArticleUsecase) updateVisibility(ctx context.Context, articleId int64, visibility enum.ArticleVisibility, userId int64) error {
	return d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		if err != nil {
			return err
		}
		if !d.isAuthor(article, userId) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished || article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		if err := d.articleRepo.UpdateVisibility(ctx, articleId, visibility, userId); err != nil {
			return err
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_STATUS_UPDATED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_STATUS_UPDATED,
				Payload: &commonenums.Event_ArticleStatusUpdated{
					ArticleStatusUpdated: &commonenums.ArticleStatusUpdatedPayload{
						SenderId:      userId,
						ArticleId:     articleId,
						Action:        "visibility_updated",
						PublishStatus: string(article.PublishStatus),
						Visibility:    string(visibility),
						Restriction:   string(article.Restriction),
					},
				},
			},
		})
	})
}

func (d *ArticleUsecase) Archive(ctx context.Context, articleId int64, userId int64, reason *string) error {
	return d.updatePublishStatus(ctx, articleId, enum.ArticlePublishStatusArchived, userId, "archived", reason)
}

func (d *ArticleUsecase) Unarchive(ctx context.Context, articleId int64, userId int64, reason *string) error {
	return d.updatePublishStatus(ctx, articleId, enum.ArticlePublishStatusPublished, userId, "unarchived", reason)
}

func (d *ArticleUsecase) updatePublishStatus(ctx context.Context, articleId int64, publishStatus enum.ArticlePublishStatus, userId int64, action string, reason *string) error {
	return d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		if err != nil {
			return err
		}
		if article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		switch action {
		case "archived":
			if article.PublishStatus != enum.ArticlePublishStatusPublished {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
			}
		case "unarchived":
			if article.PublishStatus != enum.ArticlePublishStatusArchived {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
			}
		}
		if err := d.articleRepo.UpdatePublishStatus(ctx, articleId, publishStatus, article.Visibility, nil, userId); err != nil {
			return err
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_STATUS_UPDATED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_STATUS_UPDATED,
				Payload: &commonenums.Event_ArticleStatusUpdated{
					ArticleStatusUpdated: &commonenums.ArticleStatusUpdatedPayload{
						SenderId:      userId,
						ArticleId:     articleId,
						Action:        action,
						PublishStatus: string(publishStatus),
						Visibility:    string(article.Visibility),
						Restriction:   string(article.Restriction),
						Reason:        reason,
					},
				},
			},
		})
	})
}

func (d *ArticleUsecase) Hide(ctx context.Context, articleId int64, userId int64, reason *string) error {
	return d.updateRestriction(ctx, articleId, enum.ContentRestrictionHidden, userId, enum.ContentModerationActionHide, reason)
}

func (d *ArticleUsecase) Unhide(ctx context.Context, articleId int64, userId int64, reason *string) error {
	return d.updateRestriction(ctx, articleId, enum.ContentRestrictionNone, userId, enum.ContentModerationActionUnhide, reason)
}

func (d *ArticleUsecase) Lock(ctx context.Context, articleId int64, userId int64, reason *string) error {
	return d.updateRestriction(ctx, articleId, enum.ContentRestrictionLocked, userId, enum.ContentModerationActionLock, reason)
}

func (d *ArticleUsecase) Unlock(ctx context.Context, articleId int64, userId int64, reason *string) error {
	return d.updateRestriction(ctx, articleId, enum.ContentRestrictionNone, userId, enum.ContentModerationActionUnlock, reason)
}

func (d *ArticleUsecase) updateRestriction(ctx context.Context, articleId int64, restriction enum.ContentRestriction, userId int64, action enum.ContentModerationAction, reason *string) error {
	return d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		if err != nil {
			return err
		}
		if article.PublishStatus == enum.ArticlePublishStatusDraft {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		switch action {
		case enum.ContentModerationActionHide:
			if article.Restriction != enum.ContentRestrictionNone {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
			}
		case enum.ContentModerationActionLock:
			if article.PublishStatus != enum.ArticlePublishStatusPublished ||
				article.Visibility != enum.ArticleVisibilityPublic ||
				article.Restriction != enum.ContentRestrictionNone {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
			}
		case enum.ContentModerationActionUnhide:
			if article.Restriction != enum.ContentRestrictionHidden {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
			}
		case enum.ContentModerationActionUnlock:
			if article.Restriction != enum.ContentRestrictionLocked {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
			}
		}
		if err := d.articleRepo.UpdateRestriction(ctx, articleId, restriction, userId); err != nil {
			return err
		}
		if _, err := d.moderationRecordRepo.Save(ctx, &model.ContentModerationRecord{
			Target:     enum.ContentModerationTargetArticle,
			TargetID:   articleId,
			Action:     action,
			Reason:     reason,
			OperatorID: userId,
		}); err != nil {
			return err
		}
		return d.outboxRepo.Save(ctx, &repo.OutboxEventSave{
			Event: &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_STATUS_UPDATED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_STATUS_UPDATED,
				Payload: &commonenums.Event_ArticleStatusUpdated{
					ArticleStatusUpdated: &commonenums.ArticleStatusUpdatedPayload{
						SenderId:      userId,
						ArticleId:     articleId,
						Action:        string(action),
						PublishStatus: string(article.PublishStatus),
						Visibility:    string(article.Visibility),
						Restriction:   string(restriction),
						Reason:        reason,
					},
				},
			},
		})
	})
}

func (d *ArticleUsecase) Get(ctx context.Context, articleId int64) (*model.Article, error) {
	return d.articleRepo.Get(ctx, &repo.ArticleGetReq{
		ArticleId: new(articleId),
	})
}

func (d *ArticleUsecase) Page(ctx context.Context, page *common.PageRequest, req *repo.ArticleGetReq) ([]*model.Article, *common.PageReply, error) {
	return d.articleRepo.Page(ctx, page, req)
}

func (d *ArticleUsecase) MapViewerActionStates(ctx context.Context, articleIds []int64, userId int64) (map[int64]*model.ArticleViewerActionState, error) {
	articleIds = lo.Uniq(articleIds)
	if len(articleIds) == 0 {
		return map[int64]*model.ArticleViewerActionState{}, nil
	}
	states := lo.SliceToMap(articleIds, func(articleID int64) (int64, *model.ArticleViewerActionState) {
		return articleID, &model.ArticleViewerActionState{}
	})
	records, err := d.actionRecordRepo.List(ctx, &repo.ArticleActionRecordReq{
		ArticleIds: articleIds,
		UserId:     new(userId),
		Types: []enum.ArticleAction{
			enum.ArticleActionLike,
			enum.ArticleActionThank,
			enum.ArticleActionCollect,
			enum.ArticleActionWatch,
		},
	})
	if err != nil {
		return nil, err
	}
	for _, record := range records {
		state := states[record.ArticleID]
		if state == nil {
			continue
		}
		switch record.Type {
		case enum.ArticleActionLike:
			state.Liked = true
		case enum.ArticleActionThank:
			state.Thanked = true
		case enum.ArticleActionCollect:
			state.Collected = true
		case enum.ArticleActionWatch:
			state.Watched = true
		}
	}
	return states, nil
}

func (d *ArticleUsecase) DiscardDraft(ctx context.Context, articleId int64, userId int64) error {
	return d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: &articleId,
		})
		if err != nil {
			return err
		}
		if !d.isAuthor(article, userId) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		if article.PublishStatus != enum.ArticlePublishStatusDraft {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		return d.articleRepo.DiscardDraft(ctx, articleId)
	})
}

func (d *ArticleUsecase) isAuthor(article *model.Article, userId int64) bool {
	return article != nil && article.CreatedBy != nil && *article.CreatedBy == userId
}
