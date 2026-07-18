package usecase

import (
	cerrors "common/proto/gen/common/errors"
	"context"
	"time"

	"common/pkg/apperror"
	commonenums "common/proto/gen/common/enums"
	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/config"
	"content/internal/enum"

	"github.com/samber/lo"
)

type ArticleUsecase struct {
	conf *config.Bootstrap
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
	conf *config.Bootstrap,
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

type ArticleAddReq struct {
	Article *model.Article
	TagIDs  []int64
}

func (d *ArticleUsecase) Add(ctx context.Context, req *ArticleAddReq) (*model.Article, error) {
	article := req.Article
	var (
		save *model.Article
		err  error
	)
	err = d.tx(ctx, func(ctx context.Context) error {
		bindTagIDs := lo.Uniq(req.TagIDs)
		if len(bindTagIDs) > 0 {
			tagStatus := enum.TagStatusEnabled
			countResp, err := d.tagRepo.Count(ctx, &repo.TagGetReq{TagIds: bindTagIDs, Status: &tagStatus})
			if err != nil {
				return err
			}
			if countResp != len(bindTagIDs) {
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
		saveResp, saveErr := d.articleRepo.Save(ctx, article)
		if saveErr != nil {
			return saveErr
		}
		save = saveResp
		if err != nil {
			return err
		}
		if len(bindTagIDs) > 0 {
			if err = d.articleRepo.ReplaceTags(ctx, &repo.ArticleReplaceTagsReq{ArticleID: save.ID, TagIDs: bindTagIDs}); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return save, nil
}

type ArticleAddPostscriptReq struct {
	ArticleID int64
	Content   string
	UserID    int64
}

func (d *ArticleUsecase) AddPostscript(ctx context.Context, req *ArticleAddPostscriptReq) (*model.ArticlePostscript, error) {
	articleId := req.ArticleID
	content := req.Content
	userId := req.UserID
	var save *model.ArticlePostscript
	err := d.tx(ctx, func(ctx context.Context) error {
		articleResp, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: &articleId,
		})
		article := articleResp
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
		postscriptResp, postscriptErr := d.postscriptRepo.Save(ctx, postscript)
		if postscriptErr != nil {
			return postscriptErr
		}
		save = postscriptResp
		if err != nil {
			return err
		}
		if err = d.articleRepo.UpdateHasPostscript(ctx, &repo.ArticleUpdateHasPostscriptReq{ArticleID: articleId, HasPostscript: true, UpdatedBy: userId}); err != nil {
			return err
		}
		err = d.outboxRepo.Save(ctx, &commonenums.Event{
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
		)
		return err
	})
	if err != nil {
		return nil, err
	}
	return save, nil
}

func (d *ArticleUsecase) ListPostscripts(ctx context.Context, articleID int64) ([]*model.ArticlePostscript, error) {
	articleId := articleID
	if _, err := d.Get(ctx, articleId); err != nil {
		return nil, err
	}
	restriction := enum.ContentRestrictionNone
	listResp, err := d.postscriptRepo.List(ctx, &repo.ArticlePostscriptGetReq{
		ArticleID:   new(articleId),
		Restriction: &restriction,
	})
	if err != nil {
		return nil, err
	}
	return listResp, nil
}

func (d *ArticleUsecase) Update(ctx context.Context, article *model.Article) (*model.Article, error) {
	var (
		save *model.Article
		err  error
	)
	err = d.tx(ctx, func(ctx context.Context) error {
		currentResp, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: new(article.ID),
		})
		current := currentResp
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

		now := time.Now()
		article.PublishStatus = current.PublishStatus
		article.Visibility = current.Visibility
		article.Restriction = current.Restriction
		article.PublishedAt = current.PublishedAt
		article.EditedAt = new(now)
		article.FormatContent()
		updateResp, updateErr := d.articleRepo.Update(ctx, article)
		if updateErr != nil {
			return updateErr
		}
		save = updateResp
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return save, nil
}

type ArticleRewardReq struct {
	ArticleID int64
	UserID    int64
	Points    int32
}

func (d *ArticleUsecase) Reward(ctx context.Context, req *ArticleRewardReq) error {
	return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_REWARD_NOT_IMPLEMENTED)
}

type ArticleLikeReq struct {
	ArticleID int64
	UserID    int64
	Active    bool
}

func (d *ArticleUsecase) Like(ctx context.Context, req *ArticleLikeReq) (bool, error) {
	articleId := req.ArticleID
	userId := req.UserID
	active := req.Active
	err := d.tx(ctx, func(ctx context.Context) error {
		articleResp, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		article := articleResp
		if err != nil {
			return err
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished || article.Visibility != enum.ArticleVisibilityPublic || article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}

		if active {
			createdResp, err := d.actionRecordRepo.Save(ctx, &model.ArticleActionRecord{
				ArticleID: articleId,
				UserID:    userId,
				Type:      enum.ArticleActionLike,
			})
			if err != nil {
				return err
			}
			if !createdResp {
				return nil
			}
			if err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{ArticleID: articleId, Stats: repo.ArticleStatUpdate{LikeCount: 1}}); err != nil {
				return err
			}
			err = d.outboxRepo.Save(ctx, &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_LIKED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_LIKED,
				Payload: &commonenums.Event_ArticleLiked{
					ArticleLiked: &commonenums.ArticleLikedPayload{
						SenderId:  userId,
						ArticleId: articleId,
					},
				},
			})
		}

		deletedResp, err := d.actionRecordRepo.Delete(ctx, &repo.ArticleActionRecordDeleteReq{ArticleID: articleId, UserID: userId, Action: enum.ArticleActionLike})
		if err != nil {
			return err
		}
		if deletedResp == 0 {
			return nil
		}
		err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{ArticleID: articleId, Stats: repo.ArticleStatUpdate{LikeCount: -1}})
		return err
	})
	if err != nil {
		return false, err
	}
	return active, nil
}

type ArticleThankReq struct {
	ArticleID int64
	UserID    int64
	Active    bool
}

func (d *ArticleUsecase) Thank(ctx context.Context, req *ArticleThankReq) (bool, error) {
	articleId := req.ArticleID
	userId := req.UserID
	active := req.Active
	err := d.tx(ctx, func(ctx context.Context) error {
		articleResp, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		article := articleResp
		if err != nil {
			return err
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished || article.Visibility != enum.ArticleVisibilityPublic || article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}

		if active {
			createdResp, err := d.actionRecordRepo.Save(ctx, &model.ArticleActionRecord{
				ArticleID: articleId,
				UserID:    userId,
				Type:      enum.ArticleActionThank,
			})
			if err != nil {
				return err
			}
			if !createdResp {
				return nil
			}
			if err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{ArticleID: articleId, Stats: repo.ArticleStatUpdate{ThankCount: 1}}); err != nil {
				return err
			}
			err = d.outboxRepo.Save(ctx, &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_THANKED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_THANKED,
				Payload: &commonenums.Event_ArticleThanked{
					ArticleThanked: &commonenums.ArticleThankedPayload{
						SenderId:  userId,
						ArticleId: articleId,
					},
				},
			})
		}

		deletedResp, err := d.actionRecordRepo.Delete(ctx, &repo.ArticleActionRecordDeleteReq{ArticleID: articleId, UserID: userId, Action: enum.ArticleActionThank})
		if err != nil {
			return err
		}
		if deletedResp == 0 {
			return nil
		}
		err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{ArticleID: articleId, Stats: repo.ArticleStatUpdate{ThankCount: -1}})
		return err
	})
	if err != nil {
		return false, err
	}
	return active, nil
}

type ArticleCollectReq struct {
	ArticleID int64
	UserID    int64
	Active    bool
}

func (d *ArticleUsecase) Collect(ctx context.Context, req *ArticleCollectReq) (bool, error) {
	articleId := req.ArticleID
	userId := req.UserID
	active := req.Active
	err := d.tx(ctx, func(ctx context.Context) error {
		articleResp, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		article := articleResp
		if err != nil {
			return err
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished || article.Visibility != enum.ArticleVisibilityPublic || article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}

		if active {
			createdResp, err := d.actionRecordRepo.Save(ctx, &model.ArticleActionRecord{
				ArticleID: articleId,
				UserID:    userId,
				Type:      enum.ArticleActionCollect,
			})
			if err != nil {
				return err
			}
			if !createdResp {
				return nil
			}
			if err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{ArticleID: articleId, Stats: repo.ArticleStatUpdate{CollectCount: 1}}); err != nil {
				return err
			}
			err = d.outboxRepo.Save(ctx, &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_COLLECTED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_COLLECTED,
				Payload: &commonenums.Event_ArticleCollected{
					ArticleCollected: &commonenums.ArticleCollectedPayload{
						SenderId:  userId,
						ArticleId: articleId,
					},
				},
			})
		}

		deletedResp, err := d.actionRecordRepo.Delete(ctx, &repo.ArticleActionRecordDeleteReq{ArticleID: articleId, UserID: userId, Action: enum.ArticleActionCollect})
		if err != nil {
			return err
		}
		if deletedResp == 0 {
			return nil
		}
		err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{ArticleID: articleId, Stats: repo.ArticleStatUpdate{CollectCount: -1}})
		return err
	})
	if err != nil {
		return false, err
	}
	return active, nil
}

type ArticleWatchReq struct {
	ArticleID int64
	UserID    int64
	Active    bool
}

func (d *ArticleUsecase) Watch(ctx context.Context, req *ArticleWatchReq) (bool, error) {
	articleId := req.ArticleID
	userId := req.UserID
	active := req.Active
	err := d.tx(ctx, func(ctx context.Context) error {
		articleResp, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		article := articleResp
		if err != nil {
			return err
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished || article.Visibility != enum.ArticleVisibilityPublic || article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}

		if active {
			createdResp, err := d.actionRecordRepo.Save(ctx, &model.ArticleActionRecord{
				ArticleID: articleId,
				UserID:    userId,
				Type:      enum.ArticleActionWatch,
			})
			if err != nil {
				return err
			}
			if !createdResp {
				return nil
			}
			if err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{ArticleID: articleId, Stats: repo.ArticleStatUpdate{WatchCount: 1}}); err != nil {
				return err
			}
			err = d.outboxRepo.Save(ctx, &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_WATCHED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_WATCHED,
				Payload: &commonenums.Event_ArticleWatched{
					ArticleWatched: &commonenums.ArticleWatchedPayload{
						SenderId:  userId,
						ArticleId: articleId,
					},
				},
			})
		}

		deletedResp, err := d.actionRecordRepo.Delete(ctx, &repo.ArticleActionRecordDeleteReq{ArticleID: articleId, UserID: userId, Action: enum.ArticleActionWatch})
		if err != nil {
			return err
		}
		if deletedResp == 0 {
			return nil
		}
		err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{ArticleID: articleId, Stats: repo.ArticleStatUpdate{WatchCount: -1}})
		return err
	})
	if err != nil {
		return false, err
	}
	return active, nil
}

type ArticleViewReq struct {
	ArticleID    int64
	ViewerUserID *int64
}

func (d *ArticleUsecase) View(ctx context.Context, req *ArticleViewReq) error {
	articleId := req.ArticleID
	return d.tx(ctx, func(ctx context.Context) error {
		articleResp, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		article := articleResp
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
		err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{ArticleID: articleId, Stats: repo.ArticleStatUpdate{ViewCount: 1}})
		return err
	})
}

type ArticlePublishReq struct {
	ArticleID  int64
	UserID     int64
	Visibility enum.ArticleVisibility
}

func (d *ArticleUsecase) Publish(ctx context.Context, req *ArticlePublishReq) error {
	articleId := req.ArticleID
	userId := req.UserID
	visibility := req.Visibility
	return d.tx(ctx, func(ctx context.Context) error {
		articleResp, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		article := articleResp
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
		if err = d.articleRepo.UpdatePublishStatus(ctx, &repo.ArticleUpdatePublishStatusReq{ArticleID: articleId, PublishStatus: enum.ArticlePublishStatusPublished, Visibility: visibility, PublishedAt: new(now), UpdatedBy: userId}); err != nil {
			return err
		}
		err = d.outboxRepo.Save(ctx, &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_PUBLISHED,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_PUBLISHED,
			Payload: &commonenums.Event_ArticlePublished{
				ArticlePublished: &commonenums.ArticlePublishedPayload{
					ArticleId: articleId,
				},
			},
		},
		)
		return err
	})
}

type ArticleAcceptAnswerReq struct {
	ArticleID int64
	CommentID int64
	UserID    int64
}

func (d *ArticleUsecase) AcceptAnswer(ctx context.Context, req *ArticleAcceptAnswerReq) error {
	articleId := req.ArticleID
	commentId := req.CommentID
	userId := req.UserID
	return d.tx(ctx, func(ctx context.Context) error {
		articleResp, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: new(articleId),
		})
		article := articleResp
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
		if _, err = d.articleRepo.UpdateAcceptedAnswerID(ctx, &repo.ArticleUpdateAcceptedAnswerIDReq{ArticleID: articleId, CommentID: commentId, UpdatedBy: userId}); err != nil {
			return err
		}
		err = d.outboxRepo.Save(ctx, &commonenums.Event{
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
		)
		return err
	})
}

type ArticleMakePrivateReq struct {
	ArticleID int64
	UserID    int64
}

func (d *ArticleUsecase) MakePrivate(ctx context.Context, req *ArticleMakePrivateReq) error {
	articleId := req.ArticleID
	userId := req.UserID
	return d.updateVisibility(ctx, &articleUpdateVisibilityReq{ArticleID: articleId, Visibility: enum.ArticleVisibilityPrivate, UserID: userId})
}

type ArticleMakePublicReq struct {
	ArticleID int64
	UserID    int64
}

func (d *ArticleUsecase) MakePublic(ctx context.Context, req *ArticleMakePublicReq) error {
	articleId := req.ArticleID
	userId := req.UserID
	return d.updateVisibility(ctx, &articleUpdateVisibilityReq{ArticleID: articleId, Visibility: enum.ArticleVisibilityPublic, UserID: userId})
}

type articleUpdateVisibilityReq struct {
	ArticleID  int64
	Visibility enum.ArticleVisibility
	UserID     int64
}

func (d *ArticleUsecase) updateVisibility(ctx context.Context, req *articleUpdateVisibilityReq) error {
	articleId := req.ArticleID
	visibility := req.Visibility
	userId := req.UserID
	return d.tx(ctx, func(ctx context.Context) error {
		articleResp, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		article := articleResp
		if err != nil {
			return err
		}
		if !d.isAuthor(article, userId) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished || article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		if err := d.articleRepo.UpdateVisibility(ctx, &repo.ArticleUpdateVisibilityReq{ArticleID: articleId, Visibility: visibility, UpdatedBy: userId}); err != nil {
			return err
		}
		err = d.outboxRepo.Save(ctx, &commonenums.Event{
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
		)
		return err
	})
}

type ArticleArchiveReq struct {
	ArticleID int64
	UserID    int64
	Reason    *string
}

func (d *ArticleUsecase) Archive(ctx context.Context, req *ArticleArchiveReq) error {
	articleId := req.ArticleID
	userId := req.UserID
	reason := req.Reason
	return d.updatePublishStatus(ctx, &articleUpdatePublishStatusReq{ArticleID: articleId, PublishStatus: enum.ArticlePublishStatusArchived, UserID: userId, Action: "archived", Reason: reason})
}

type ArticleUnarchiveReq struct {
	ArticleID int64
	UserID    int64
	Reason    *string
}

func (d *ArticleUsecase) Unarchive(ctx context.Context, req *ArticleUnarchiveReq) error {
	articleId := req.ArticleID
	userId := req.UserID
	reason := req.Reason
	return d.updatePublishStatus(ctx, &articleUpdatePublishStatusReq{ArticleID: articleId, PublishStatus: enum.ArticlePublishStatusPublished, UserID: userId, Action: "unarchived", Reason: reason})
}

type articleUpdatePublishStatusReq struct {
	ArticleID     int64
	PublishStatus enum.ArticlePublishStatus
	UserID        int64
	Action        string
	Reason        *string
}

func (d *ArticleUsecase) updatePublishStatus(ctx context.Context, req *articleUpdatePublishStatusReq) error {
	articleId := req.ArticleID
	publishStatus := req.PublishStatus
	userId := req.UserID
	action := req.Action
	reason := req.Reason
	return d.tx(ctx, func(ctx context.Context) error {
		articleResp, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		article := articleResp
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
		if err := d.articleRepo.UpdatePublishStatus(ctx, &repo.ArticleUpdatePublishStatusReq{ArticleID: articleId, PublishStatus: publishStatus, Visibility: article.Visibility, UpdatedBy: userId}); err != nil {
			return err
		}
		err = d.outboxRepo.Save(ctx, &commonenums.Event{
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
		)
		return err
	})
}

type ArticleHideReq struct {
	ArticleID int64
	UserID    int64
	Reason    *string
}

func (d *ArticleUsecase) Hide(ctx context.Context, req *ArticleHideReq) error {
	articleId := req.ArticleID
	userId := req.UserID
	reason := req.Reason
	return d.updateRestriction(ctx, &articleUpdateRestrictionReq{ArticleID: articleId, Restriction: enum.ContentRestrictionHidden, UserID: userId, Action: enum.ContentModerationActionHide, Reason: reason})
}

type ArticleUnhideReq struct {
	ArticleID int64
	UserID    int64
	Reason    *string
}

func (d *ArticleUsecase) Unhide(ctx context.Context, req *ArticleUnhideReq) error {
	articleId := req.ArticleID
	userId := req.UserID
	reason := req.Reason
	return d.updateRestriction(ctx, &articleUpdateRestrictionReq{ArticleID: articleId, Restriction: enum.ContentRestrictionNone, UserID: userId, Action: enum.ContentModerationActionUnhide, Reason: reason})
}

type ArticleLockReq struct {
	ArticleID int64
	UserID    int64
	Reason    *string
}

func (d *ArticleUsecase) Lock(ctx context.Context, req *ArticleLockReq) error {
	articleId := req.ArticleID
	userId := req.UserID
	reason := req.Reason
	return d.updateRestriction(ctx, &articleUpdateRestrictionReq{ArticleID: articleId, Restriction: enum.ContentRestrictionLocked, UserID: userId, Action: enum.ContentModerationActionLock, Reason: reason})
}

type ArticleUnlockReq struct {
	ArticleID int64
	UserID    int64
	Reason    *string
}

func (d *ArticleUsecase) Unlock(ctx context.Context, req *ArticleUnlockReq) error {
	articleId := req.ArticleID
	userId := req.UserID
	reason := req.Reason
	return d.updateRestriction(ctx, &articleUpdateRestrictionReq{ArticleID: articleId, Restriction: enum.ContentRestrictionNone, UserID: userId, Action: enum.ContentModerationActionUnlock, Reason: reason})
}

type articleUpdateRestrictionReq struct {
	ArticleID   int64
	Restriction enum.ContentRestriction
	UserID      int64
	Action      enum.ContentModerationAction
	Reason      *string
}

func (d *ArticleUsecase) updateRestriction(ctx context.Context, req *articleUpdateRestrictionReq) error {
	articleId := req.ArticleID
	restriction := req.Restriction
	userId := req.UserID
	action := req.Action
	reason := req.Reason
	return d.tx(ctx, func(ctx context.Context) error {
		articleResp, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{ArticleId: new(articleId)})
		article := articleResp
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
		if err := d.articleRepo.UpdateRestriction(ctx, &repo.ArticleUpdateRestrictionReq{ArticleID: articleId, Restriction: restriction, UpdatedBy: userId}); err != nil {
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
		err = d.outboxRepo.Save(ctx, &commonenums.Event{
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
		)
		return err
	})
}

func (d *ArticleUsecase) Get(ctx context.Context, articleID int64) (*model.Article, error) {
	articleId := articleID
	articleResp, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
		ArticleId: new(articleId),
	})
	article := articleResp
	if err != nil {
		return nil, err
	}
	return article, nil
}

type ArticlePageReq struct {
	Page            *base.PageRequest
	TagID           *int64
	DomainID        *int64
	PublishStatus   *enum.ArticlePublishStatus
	PublishStatuses []enum.ArticlePublishStatus
	Visibility      *enum.ArticleVisibility
	Visibilities    []enum.ArticleVisibility
	Restriction     *enum.ContentRestriction
	Restrictions    []enum.ContentRestriction
	AuthorID        *int64
	Order           *enum.ArticleOrder
	Type            *enum.ArticleType
	Keyword         *string
}

type ArticlePageResp struct {
	Rows []*model.Article
	Page *base.PageResp
}

func (d *ArticleUsecase) Page(ctx context.Context, req *ArticlePageReq) (*ArticlePageResp, error) {
	if req == nil {
		req = &ArticlePageReq{}
	}
	pageResp, err := d.articleRepo.Page(ctx, &repo.ArticleGetReq{
		Page:            req.Page,
		TagId:           req.TagID,
		DomainId:        req.DomainID,
		PublishStatus:   req.PublishStatus,
		PublishStatuses: req.PublishStatuses,
		Visibility:      req.Visibility,
		Visibilities:    req.Visibilities,
		Restriction:     req.Restriction,
		Restrictions:    req.Restrictions,
		AuthorId:        req.AuthorID,
		Order:           req.Order,
		Type:            req.Type,
		Keyword:         req.Keyword,
	})
	if err != nil {
		return nil, err
	}
	return &ArticlePageResp{Rows: pageResp.Rows, Page: pageResp.Page}, nil
}

type ArticleMapViewerActionStatesReq struct {
	ArticleIDs []int64
	UserID     int64
}

func (d *ArticleUsecase) MapViewerActionStates(ctx context.Context, req *ArticleMapViewerActionStatesReq) (map[int64]*model.
	ArticleViewerActionState, error) {
	articleIds := req.ArticleIDs
	userId := req.UserID
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

type ArticleDiscardDraftReq struct {
	ArticleID int64
	UserID    int64
}

func (d *ArticleUsecase) DiscardDraft(ctx context.Context, req *ArticleDiscardDraftReq) error {
	articleId := req.ArticleID
	userId := req.UserID
	return d.tx(ctx, func(ctx context.Context) error {
		articleResp, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			ArticleId: &articleId,
		})
		article := articleResp
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
