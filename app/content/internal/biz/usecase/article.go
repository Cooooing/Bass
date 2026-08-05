package usecase

import (
	cerrors "common/proto/gen/common/errors"
	"context"
	"log/slog"
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
	log  *slog.Logger
	conf *config.Bootstrap
	tx   base.Tx

	articleRepo          repo.ArticleRepo
	actionRecordRepo     repo.ArticleActionRecordRepo
	commentRepo          repo.CommentRepo
	outboxRepo           repo.OutboxEventRepo
	outboxUsecase        *OutboxUsecase
	moderationRecordRepo repo.ContentModerationRecordRepo
	viewCacheRepo        repo.ArticleViewCacheRepo
	viewRecordRepo       repo.ArticleViewRecordRepo
	delayedTaskClient    repo.DelayedTaskClient
	articleAccessUsecase *ArticleAccessUsecase
}

func NewArticleUsecase(
	logger *slog.Logger,
	conf *config.Bootstrap,
	tx base.Tx,
	articleRepo repo.ArticleRepo,
	actionRecordRepo repo.ArticleActionRecordRepo,
	commentRepo repo.CommentRepo,
	outboxRepo repo.OutboxEventRepo,
	outboxUsecase *OutboxUsecase,
	moderationRecordRepo repo.ContentModerationRecordRepo,
	viewCacheRepo repo.ArticleViewCacheRepo,
	viewRecordRepo repo.ArticleViewRecordRepo,
	delayedTaskClient repo.DelayedTaskClient,
	articleAccessUsecase *ArticleAccessUsecase,
) *ArticleUsecase {
	return &ArticleUsecase{
		log:                  logger,
		conf:                 conf,
		tx:                   tx,
		articleRepo:          articleRepo,
		actionRecordRepo:     actionRecordRepo,
		commentRepo:          commentRepo,
		outboxRepo:           outboxRepo,
		outboxUsecase:        outboxUsecase,
		moderationRecordRepo: moderationRecordRepo,
		viewCacheRepo:        viewCacheRepo,
		viewRecordRepo:       viewRecordRepo,
		delayedTaskClient:    delayedTaskClient,
		articleAccessUsecase: articleAccessUsecase,
	}
}

type ArticleAddReq struct {
	Access  *model.ContentAccess
	Article *model.Article
}

func (d *ArticleUsecase) Add(ctx context.Context, req *ArticleAddReq) (*model.Article, error) {
	article := req.Article
	access, err := req.Access.Normalize("")
	if err != nil {
		return nil, err
	}
	if err = d.articleAccessUsecase.CanCreateDraft(access); err != nil {
		return nil, err
	}
	if err := article.CanCreateDraft(); err != nil {
		return nil, err
	}
	article.PublishStatus = enum.ArticlePublishStatusDraft
	if article.Visibility == "" {
		article.Visibility = enum.ArticleVisibilityPublic
	}
	article.Restriction = enum.ContentRestrictionNone
	article.EditedAt = new(time.Now())
	article.CreatedBy = new(access.ActorUserID)
	article.UpdatedBy = new(access.ActorUserID)
	article.FormatContent()

	var save *model.Article
	err = d.tx(ctx, func(ctx context.Context) error {
		save, err = d.articleRepo.Save(ctx, article)
		return err
	})
	if err != nil {
		return nil, err
	}
	return save, nil
}

type ArticleUpdateDraftReq struct {
	Access  *model.ContentAccess
	Article *model.Article
}

func (d *ArticleUsecase) Update(ctx context.Context, req *ArticleUpdateDraftReq) (*model.Article, error) {
	article := req.Article
	access, err := req.Access.Normalize("")
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var save *model.Article
	err = d.tx(ctx, func(ctx context.Context) error {
		current, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: new(article.ID)}})
		if err != nil {
			return err
		}
		if err = d.articleAccessUsecase.CanUpdateDraft(access, current); err != nil {
			return err
		}
		article.PublishStatus = current.PublishStatus
		article.Visibility = current.Visibility
		article.Restriction = current.Restriction
		article.PublishedAt = current.PublishedAt
		article.EditedAt = new(time.Now())
		article.UpdatedBy = new(access.ActorUserID)
		article.FormatContent()
		save, err = d.articleRepo.Update(ctx, article)
		return err
	})
	if err != nil {
		return nil, err
	}
	return save, nil
}

type ArticleLikeReq struct {
	Access    *model.ContentAccess
	ArticleID int64
	UserID    int64
	Active    bool
}

func (d *ArticleUsecase) Like(ctx context.Context, req *ArticleLikeReq) (bool, error) {
	articleId := req.ArticleID
	userId := req.Access.ActorUserID
	active := req.Active
	var outboxEvent *repo.OutboxEvent
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleId)},
		})
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
			if err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{
				ArticleID: articleId,
				Stats: repo.ArticleStatUpdate{
					LikeCount: 1,
				},
			}); err != nil {
				return err
			}
			outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_LIKED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_LIKED,
				Payload: &commonenums.Event_ArticleLiked{
					ArticleLiked: &commonenums.ArticleLikedPayload{
						SenderId:  userId,
						ArticleId: articleId,
					},
				},
			})
			return err
		}

		deletedResp, err := d.actionRecordRepo.Delete(ctx, &repo.ArticleActionRecordDeleteReq{
			ArticleID: articleId,
			UserID:    userId,
			Action:    enum.ArticleActionLike,
		})
		if err != nil {
			return err
		}
		if deletedResp == 0 {
			return nil
		}
		err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{
			ArticleID: articleId,
			Stats: repo.ArticleStatUpdate{
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

type ArticleThankReq struct {
	Access    *model.ContentAccess
	ArticleID int64
	UserID    int64
	Active    bool
}

func (d *ArticleUsecase) Thank(ctx context.Context, req *ArticleThankReq) (bool, error) {
	articleId := req.ArticleID
	userId := req.Access.ActorUserID
	active := req.Active
	var outboxEvent *repo.OutboxEvent
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleId)},
		})
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
			if err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{
				ArticleID: articleId,
				Stats: repo.ArticleStatUpdate{
					ThankCount: 1,
				},
			}); err != nil {
				return err
			}
			outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_THANKED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_THANKED,
				Payload: &commonenums.Event_ArticleThanked{
					ArticleThanked: &commonenums.ArticleThankedPayload{
						SenderId:  userId,
						ArticleId: articleId,
					},
				},
			})
			return err
		}

		deletedResp, err := d.actionRecordRepo.Delete(ctx, &repo.ArticleActionRecordDeleteReq{
			ArticleID: articleId,
			UserID:    userId,
			Action:    enum.ArticleActionThank,
		})
		if err != nil {
			return err
		}
		if deletedResp == 0 {
			return nil
		}
		err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{
			ArticleID: articleId,
			Stats: repo.ArticleStatUpdate{
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

type ArticleCollectReq struct {
	Access    *model.ContentAccess
	ArticleID int64
	UserID    int64
	Active    bool
}

func (d *ArticleUsecase) Collect(ctx context.Context, req *ArticleCollectReq) (bool, error) {
	articleId := req.ArticleID
	userId := req.Access.ActorUserID
	active := req.Active
	var outboxEvent *repo.OutboxEvent
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleId)},
		})
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
			if err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{
				ArticleID: articleId,
				Stats: repo.ArticleStatUpdate{
					CollectCount: 1,
				},
			}); err != nil {
				return err
			}
			outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{
				Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_COLLECTED,
				Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_COLLECTED,
				Payload: &commonenums.Event_ArticleCollected{
					ArticleCollected: &commonenums.ArticleCollectedPayload{
						SenderId:  userId,
						ArticleId: articleId,
					},
				},
			})
			return err
		}

		deletedResp, err := d.actionRecordRepo.Delete(ctx, &repo.ArticleActionRecordDeleteReq{
			ArticleID: articleId,
			UserID:    userId,
			Action:    enum.ArticleActionCollect,
		})
		if err != nil {
			return err
		}
		if deletedResp == 0 {
			return nil
		}
		err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{
			ArticleID: articleId,
			Stats: repo.ArticleStatUpdate{
				CollectCount: -1,
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

type ArticleRewardReq struct {
	Access    *model.ContentAccess
	ArticleID int64
	UserID    int64
	Points    int32
}

func (d *ArticleUsecase) Reward(ctx context.Context, req *ArticleRewardReq) error {
	articleId := req.ArticleID
	userId := req.Access.ActorUserID
	return d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleId)},
		})
		if err != nil {
			return err
		}
		if err = d.articleAccessUsecase.CanInteract(req.Access, article); err != nil {
			return err
		}
		created, err := d.actionRecordRepo.Save(ctx, &model.ArticleActionRecord{
			ArticleID: articleId,
			UserID:    userId,
			Type:      enum.ArticleActionReward,
		})
		if err != nil || !created {
			return err
		}
		return d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{
			ArticleID: articleId,
			Stats: repo.ArticleStatUpdate{
				RewardCount: 1,
			},
		})
	})
}

type ArticleViewReq struct {
	Access             *model.ContentAccess
	ArticleID          int64
	ViewerUserID       *int64
	IP                 *string
	UserAgent          *string
	BrowserFingerprint *string
}

func (d *ArticleUsecase) View(ctx context.Context, req *ArticleViewReq) error {
	article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
		Filter: &model.ArticleFilter{ArticleID: new(req.ArticleID)},
	})
	if err != nil {
		return err
	}
	viewerID := int64(0)
	if req.ViewerUserID != nil {
		viewerID = *req.ViewerUserID
	}
	if err = article.CanView(viewerID); err != nil {
		return err
	}
	created, err := d.viewCacheRepo.Record(ctx, &repo.ArticleViewCacheRecordReq{
		ArticleID:          req.ArticleID,
		ViewerUserID:       req.ViewerUserID,
		IP:                 req.IP,
		UserAgent:          req.UserAgent,
		BrowserFingerprint: req.BrowserFingerprint,
	})
	if err != nil || !created {
		return err
	}
	if viewerID <= 0 {
		return nil
	}
	return d.viewRecordRepo.Save(ctx, &model.ArticleViewRecord{
		ArticleID:          req.ArticleID,
		UserID:             viewerID,
		IP:                 req.IP,
		UserAgent:          req.UserAgent,
		BrowserFingerprint: req.BrowserFingerprint,
		ViewedAt:           new(time.Now()),
	})
}

type ArticleFlushViewsReq struct {
	Limit int32
}

func (d *ArticleUsecase) FlushViews(ctx context.Context, req *ArticleFlushViewsReq) (int32, error) {
	counts, err := d.viewCacheRepo.PopCounts(ctx, req.Limit)
	if err != nil {
		return 0, err
	}
	flushed := int32(0)
	for articleID, count := range counts {
		if count <= 0 {
			continue
		}
		if err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{
			ArticleID: articleID,
			Stats: repo.ArticleStatUpdate{
				ViewCount: count,
			},
		}); err != nil {
			return flushed, err
		}
		flushed++
	}
	return flushed, nil
}

type ArticlePublishReq struct {
	Access     *model.ContentAccess
	ArticleID  int64
	Visibility enum.ArticleVisibility
}

func (d *ArticleUsecase) Publish(ctx context.Context, req *ArticlePublishReq) error {
	articleId := req.ArticleID
	access, err := req.Access.Normalize(enum.ContentAccessScopeInternalTask)
	if err != nil {
		return err
	}
	var operatorUserId *int64
	if access.ActorUserID > 0 {
		operatorUserId = new(access.ActorUserID)
	}
	visibility := req.Visibility
	var outboxEvent *repo.OutboxEvent
	err = d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleId)},
		})
		if err != nil {
			return err
		}
		if err = article.CanPublish(access, nil, time.Now()); err != nil {
			return err
		}
		switch visibility {
		case enum.ArticleVisibilityPublic, enum.ArticleVisibilityPrivate:
		default:
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		if err = d.articleRepo.UpdatePublishStatus(ctx, &repo.ArticleUpdatePublishStatusReq{
			ArticleID:     articleId,
			PublishStatus: enum.ArticlePublishStatusPublished,
			Visibility:    visibility,
			PublishedAt:   new(time.Now()),
			UpdatedBy:     operatorUserId,
		}); err != nil {
			return err
		}
		outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_PUBLISHED,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_PUBLISHED,
			Payload: &commonenums.Event_ArticlePublished{
				ArticlePublished: &commonenums.ArticlePublishedPayload{
					ArticleId: articleId,
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

type ArticleSchedulePublishReq struct {
	Access    *model.ContentAccess
	ArticleID int64
	PublishAt time.Time
}

func (d *ArticleUsecase) SchedulePublish(ctx context.Context, req *ArticleSchedulePublishReq) error {
	articleId := req.ArticleID
	access, err := req.Access.Normalize("")
	if err != nil {
		return err
	}
	operatorUserId := access.ActorUserID
	publishAt := req.PublishAt
	err = d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleId)},
		})
		if err != nil {
			return err
		}
		if err = article.CanPublish(access, new(publishAt), time.Now()); err != nil {
			return err
		}
		return d.articleRepo.UpdatePublishStatus(ctx, &repo.ArticleUpdatePublishStatusReq{
			ArticleID:     articleId,
			PublishStatus: enum.ArticlePublishStatusScheduled,
			Visibility:    article.Visibility,
			PublishedAt:   new(publishAt),
			UpdatedBy:     new(operatorUserId),
		})
	})
	if err != nil {
		return err
	}
	return d.delayedTaskClient.RegisterPublishScheduledArticle(ctx, articleId, publishAt)
}

type ArticleCancelPublishReq struct {
	Access    *model.ContentAccess
	ArticleID int64
}

func (d *ArticleUsecase) CancelPublish(ctx context.Context, req *ArticleCancelPublishReq) error {
	articleId := req.ArticleID
	access, err := req.Access.Normalize("")
	if err != nil {
		return err
	}
	operatorUserId := access.ActorUserID
	err = d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleId)},
		})
		if err != nil {
			return err
		}
		if err = article.CanCancelSchedule(operatorUserId); err != nil {
			return err
		}
		return d.articleRepo.UpdatePublishStatus(ctx, &repo.ArticleUpdatePublishStatusReq{
			ArticleID:      articleId,
			PublishStatus:  enum.ArticlePublishStatusDraft,
			Visibility:     article.Visibility,
			ClearPublished: true,
			UpdatedBy:      new(operatorUserId),
		})
	})
	if err != nil {
		return err
	}
	return d.delayedTaskClient.CancelPublishScheduledArticle(ctx, articleId)
}

type ArticleMakePrivateReq struct {
	Access    *model.ContentAccess
	ArticleID int64
}

func (d *ArticleUsecase) MakePrivate(ctx context.Context, req *ArticleMakePrivateReq) error {
	articleId := req.ArticleID
	userId := req.Access.ActorUserID
	return d.updateVisibility(ctx, &articleUpdateVisibilityReq{
		ArticleID:  articleId,
		Visibility: enum.ArticleVisibilityPrivate,
		UserID:     userId,
	})
}

type ArticleMakePublicReq struct {
	Access    *model.ContentAccess
	ArticleID int64
}

func (d *ArticleUsecase) MakePublic(ctx context.Context, req *ArticleMakePublicReq) error {
	articleId := req.ArticleID
	userId := req.Access.ActorUserID
	return d.updateVisibility(ctx, &articleUpdateVisibilityReq{
		ArticleID:  articleId,
		Visibility: enum.ArticleVisibilityPublic,
		UserID:     userId,
	})
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
	var outboxEvent *repo.OutboxEvent
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleId)},
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
		if err := d.articleRepo.UpdateVisibility(ctx, &repo.ArticleUpdateVisibilityReq{
			ArticleID:  articleId,
			Visibility: visibility,
			UpdatedBy:  userId,
		}); err != nil {
			return err
		}
		outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_STATUS_UPDATED,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_STATUS_UPDATED,
			Payload: &commonenums.Event_ArticleStatusUpdated{
				ArticleStatusUpdated: &commonenums.ArticleStatusUpdatedPayload{
					SenderId:      userId,
					ArticleId:     articleId,
					Action:        "visibility_updated",
					PublishStatus: article.PublishStatus.String(),
					Visibility:    visibility.String(),
					Restriction:   article.Restriction.String(),
				},
			},
		},
		)
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

type ArticleArchiveReq struct {
	Access    *model.ContentAccess
	ArticleID int64
	Reason    *string
}

func (d *ArticleUsecase) Archive(ctx context.Context, req *ArticleArchiveReq) error {
	articleId := req.ArticleID
	userId := req.Access.ActorUserID
	reason := req.Reason
	return d.updatePublishStatus(ctx, &articleUpdatePublishStatusReq{
		ArticleID:     articleId,
		PublishStatus: enum.ArticlePublishStatusArchived,
		UserID:        userId,
		Action:        "archived",
		Reason:        reason,
	})
}

type ArticleUnarchiveReq struct {
	Access    *model.ContentAccess
	ArticleID int64
	Reason    *string
}

func (d *ArticleUsecase) Unarchive(ctx context.Context, req *ArticleUnarchiveReq) error {
	articleId := req.ArticleID
	userId := req.Access.ActorUserID
	reason := req.Reason
	return d.updatePublishStatus(ctx, &articleUpdatePublishStatusReq{
		ArticleID:     articleId,
		PublishStatus: enum.ArticlePublishStatusPublished,
		UserID:        userId,
		Action:        "unarchived",
		Reason:        reason,
	})
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
	var outboxEvent *repo.OutboxEvent
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleId)},
		})
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
		if err := d.articleRepo.UpdatePublishStatus(ctx, &repo.ArticleUpdatePublishStatusReq{
			ArticleID:     articleId,
			PublishStatus: publishStatus,
			Visibility:    article.Visibility,
			UpdatedBy:     new(userId),
		}); err != nil {
			return err
		}
		outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_STATUS_UPDATED,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_STATUS_UPDATED,
			Payload: &commonenums.Event_ArticleStatusUpdated{
				ArticleStatusUpdated: &commonenums.ArticleStatusUpdatedPayload{
					SenderId:      userId,
					ArticleId:     articleId,
					Action:        action,
					PublishStatus: publishStatus.String(),
					Visibility:    article.Visibility.String(),
					Restriction:   article.Restriction.String(),
					Reason:        reason,
				},
			},
		},
		)
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

type ArticleHideReq struct {
	Access    *model.ContentAccess
	ArticleID int64
	Reason    *string
}

func (d *ArticleUsecase) Hide(ctx context.Context, req *ArticleHideReq) error {
	articleId := req.ArticleID
	userId := req.Access.ActorUserID
	reason := req.Reason
	return d.updateRestriction(ctx, &articleUpdateRestrictionReq{
		ArticleID:   articleId,
		Restriction: enum.ContentRestrictionHidden,
		UserID:      userId,
		Action:      enum.ContentModerationActionHide,
		Reason:      reason,
	})
}

type ArticleUnhideReq struct {
	Access    *model.ContentAccess
	ArticleID int64
	Reason    *string
}

func (d *ArticleUsecase) Unhide(ctx context.Context, req *ArticleUnhideReq) error {
	articleId := req.ArticleID
	userId := req.Access.ActorUserID
	reason := req.Reason
	return d.updateRestriction(ctx, &articleUpdateRestrictionReq{
		ArticleID:   articleId,
		Restriction: enum.ContentRestrictionNone,
		UserID:      userId,
		Action:      enum.ContentModerationActionUnhide,
		Reason:      reason,
	})
}

type ArticleLockReq struct {
	Access    *model.ContentAccess
	ArticleID int64
	Reason    *string
}

func (d *ArticleUsecase) Lock(ctx context.Context, req *ArticleLockReq) error {
	articleId := req.ArticleID
	userId := req.Access.ActorUserID
	reason := req.Reason
	return d.updateRestriction(ctx, &articleUpdateRestrictionReq{
		ArticleID:   articleId,
		Restriction: enum.ContentRestrictionLocked,
		UserID:      userId,
		Action:      enum.ContentModerationActionLock,
		Reason:      reason,
	})
}

type ArticleUnlockReq struct {
	Access    *model.ContentAccess
	ArticleID int64
	Reason    *string
}

func (d *ArticleUsecase) Unlock(ctx context.Context, req *ArticleUnlockReq) error {
	articleId := req.ArticleID
	userId := req.Access.ActorUserID
	reason := req.Reason
	return d.updateRestriction(ctx, &articleUpdateRestrictionReq{
		ArticleID:   articleId,
		Restriction: enum.ContentRestrictionNone,
		UserID:      userId,
		Action:      enum.ContentModerationActionUnlock,
		Reason:      reason,
	})
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
	var outboxEvent *repo.OutboxEvent
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleId)},
		})
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
		if err := d.articleRepo.UpdateRestriction(ctx, &repo.ArticleUpdateRestrictionReq{
			ArticleID:   articleId,
			Restriction: restriction,
			UpdatedBy:   userId,
		}); err != nil {
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
		outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_STATUS_UPDATED,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_STATUS_UPDATED,
			Payload: &commonenums.Event_ArticleStatusUpdated{
				ArticleStatusUpdated: &commonenums.ArticleStatusUpdatedPayload{
					SenderId:      userId,
					ArticleId:     articleId,
					Action:        action.String(),
					PublishStatus: article.PublishStatus.String(),
					Visibility:    article.Visibility.String(),
					Restriction:   restriction.String(),
					Reason:        reason,
				},
			},
		},
		)
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

func (d *ArticleUsecase) Get(ctx context.Context, articleID int64) (*model.Article, error) {
	articleId := articleID
	article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
		Filter: &model.ArticleFilter{ArticleID: new(articleId)},
	})
	if err != nil {
		return nil, err
	}
	return article, nil
}

type ArticleGetByAccessReq struct {
	ArticleID int64
	Access    *model.ContentAccess
}

func (d *ArticleUsecase) GetByAccess(ctx context.Context, req *ArticleGetByAccessReq) (*model.Article, error) {
	article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: new(req.ArticleID)}})
	if err != nil {
		return nil, err
	}
	if err = d.articleAccessUsecase.CanGet(req.Access, article); err != nil {
		return nil, err
	}
	return article, nil
}

type ArticleListByAccessReq struct {
	Access *model.ContentAccess
	Filter *model.ArticleFilter
}

func (d *ArticleUsecase) ListByAccess(ctx context.Context, req *ArticleListByAccessReq) ([]*model.Article, error) {
	if req == nil {
		req = &ArticleListByAccessReq{}
	}
	scope, err := d.articleAccessUsecase.BuildScope(req.Access)
	if err != nil {
		return nil, err
	}
	return d.articleRepo.List(ctx, &repo.ArticleGetReq{Filter: req.Filter, Scope: scope})
}

type ArticlePageByAccessReq struct {
	Access *model.ContentAccess
	Filter *model.ArticleFilter
	Page   *base.PageRequest
}

func (d *ArticleUsecase) PageByAccess(ctx context.Context, req *ArticlePageByAccessReq) (*ArticlePageResp, error) {
	if req == nil {
		req = &ArticlePageByAccessReq{}
	}
	scope, err := d.articleAccessUsecase.BuildScope(req.Access)
	if err != nil {
		return nil, err
	}
	pageResp, err := d.articleRepo.Page(ctx, &repo.ArticleGetReq{Page: req.Page, Filter: req.Filter, Scope: scope})
	if err != nil {
		return nil, err
	}
	return &ArticlePageResp{Rows: pageResp.Rows, Page: pageResp.Page}, nil
}

type ArticleListReq struct {
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
	PublishedAtEnd  *time.Time
	ArticleIDs      []int64
}

func (d *ArticleUsecase) List(ctx context.Context, req *ArticleListReq) ([]*model.Article, error) {
	if req == nil {
		req = &ArticleListReq{}
	}
	return d.articleRepo.List(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{
		TagID:           req.TagID,
		DomainID:        req.DomainID,
		ArticleIDs:      req.ArticleIDs,
		PublishStatus:   req.PublishStatus,
		PublishStatuses: req.PublishStatuses,
		Visibility:      req.Visibility,
		Visibilities:    req.Visibilities,
		Restriction:     req.Restriction,
		Restrictions:    req.Restrictions,
		AuthorID:        req.AuthorID,
		Order:           req.Order,
		Type:            req.Type,
		Keyword:         req.Keyword,
		PublishedAtEnd:  req.PublishedAtEnd,
	}})
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
	PublishedAtEnd  *time.Time
	ArticleIDs      []int64
}

type ArticlePageResp struct {
	Rows []*model.Article
	Page *base.PageResp
}

func (d *ArticleUsecase) Page(ctx context.Context, req *ArticlePageReq) (*ArticlePageResp, error) {
	if req == nil {
		req = &ArticlePageReq{}
	}
	pageResp, err := d.articleRepo.Page(ctx, &repo.ArticleGetReq{Page: req.Page, Filter: &model.ArticleFilter{
		TagID:           req.TagID,
		DomainID:        req.DomainID,
		ArticleIDs:      req.ArticleIDs,
		PublishStatus:   req.PublishStatus,
		PublishStatuses: req.PublishStatuses,
		Visibility:      req.Visibility,
		Visibilities:    req.Visibilities,
		Restriction:     req.Restriction,
		Restrictions:    req.Restrictions,
		AuthorID:        req.AuthorID,
		Order:           req.Order,
		Type:            req.Type,
		Keyword:         req.Keyword,
		PublishedAtEnd:  req.PublishedAtEnd,
	}})
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
			enum.ArticleActionReward,
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
		case enum.ArticleActionReward:
			state.Rewarded = true
		}
	}
	return states, nil
}

type ArticleDiscardDraftReq struct {
	Access    *model.ContentAccess
	ArticleID int64
}

func (d *ArticleUsecase) DiscardDraft(ctx context.Context, req *ArticleDiscardDraftReq) error {
	articleId := req.ArticleID
	userId := req.Access.ActorUserID
	return d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleId)},
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
