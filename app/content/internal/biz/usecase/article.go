package usecase

import (
	commonclient "common/pkg/client"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	cerrors "common/proto/gen/common/errors"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"common/pkg/apperror"
	commonenums "common/proto/gen/common/enums"
	contentv1 "common/proto/gen/content/v1"
	economyv1 "common/proto/gen/economy/v1"
	economyv1enum "common/proto/gen/economy/v1/enum"

	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/config"
	"content/internal/enum"

	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/samber/lo"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticleUsecase struct {
	log  *slog.Logger
	conf *config.Bootstrap
	tx   base.Tx

	articleRepo          repo.ArticleRepo
	actionRecordRepo     repo.ArticleActionRecordRepo
	commentRepo          repo.CommentRepo
	articleQARepo        repo.ArticleQARepo
	articlePollRepo      repo.ArticlePollRepo
	articleLotteryRepo   repo.ArticleLotteryRepo
	outboxRepo           repo.OutboxEventRepo
	outboxUsecase        *OutboxUsecase
	moderationRecordRepo repo.ContentModerationRecordRepo
	viewCacheRepo        repo.ArticleViewCacheRepo
	viewRecordRepo       repo.ArticleViewRecordRepo
	delayedTaskClient    repo.DelayedTaskClient
	articleAccessUsecase *ArticleAccessUsecase
	thankUsecase         *ArticleThankUsecase
	rewardUsecase        *ArticleRewardUsecase
	dtm                  *commonclient.DTMClient
	branches             *commonclient.TCCBranchCaller
	business             *config.Business
}

func NewArticleUsecase(
	logger *slog.Logger,
	conf *config.Bootstrap,
	tx base.Tx,
	articleRepo repo.ArticleRepo,
	actionRecordRepo repo.ArticleActionRecordRepo,
	commentRepo repo.CommentRepo,
	articleQARepo repo.ArticleQARepo,
	articlePollRepo repo.ArticlePollRepo,
	articleLotteryRepo repo.ArticleLotteryRepo,
	outboxRepo repo.OutboxEventRepo,
	outboxUsecase *OutboxUsecase,
	moderationRecordRepo repo.ContentModerationRecordRepo,
	viewCacheRepo repo.ArticleViewCacheRepo,
	viewRecordRepo repo.ArticleViewRecordRepo,
	delayedTaskClient repo.DelayedTaskClient,
	articleAccessUsecase *ArticleAccessUsecase,
	thankUsecase *ArticleThankUsecase,
	rewardUsecase *ArticleRewardUsecase,
	dtm *commonclient.DTMClient,
	branches *commonclient.TCCBranchCaller,
	business *config.Business,
) *ArticleUsecase {
	return &ArticleUsecase{
		log:                  logger,
		conf:                 conf,
		tx:                   tx,
		articleRepo:          articleRepo,
		actionRecordRepo:     actionRecordRepo,
		commentRepo:          commentRepo,
		articleQARepo:        articleQARepo,
		articlePollRepo:      articlePollRepo,
		articleLotteryRepo:   articleLotteryRepo,
		outboxRepo:           outboxRepo,
		outboxUsecase:        outboxUsecase,
		moderationRecordRepo: moderationRecordRepo,
		viewCacheRepo:        viewCacheRepo,
		viewRecordRepo:       viewRecordRepo,
		delayedTaskClient:    delayedTaskClient,
		articleAccessUsecase: articleAccessUsecase,
		thankUsecase:         thankUsecase,
		rewardUsecase:        rewardUsecase,
		dtm:                  dtm,
		branches:             branches,
		business:             business,
	}
}

type ArticleAddReq struct {
	Access  *model.ContentAccess
	Article *model.Article
	Poll    *model.ArticlePollDraft
	Lottery *model.ArticleLotteryDraft
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
	if article.Type == enum.ArticleTypePoll {
		if err := req.Poll.Normalize(time.Now()); err != nil {
			return nil, err
		}
	}
	if article.Type == enum.ArticleTypeLottery {
		if err := req.Lottery.Normalize(time.Now()); err != nil {
			return nil, err
		}
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
		if err != nil {
			return err
		}
		if article.Type == enum.ArticleTypePoll {
			return d.articlePollRepo.SaveDraft(ctx, save.ID, req.Poll)
		}
		if article.Type == enum.ArticleTypeLottery {
			return d.articleLotteryRepo.SaveDraft(ctx, save.ID, req.Lottery)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return save, nil
}

type ArticleUpdateDraftReq struct {
	Access  *model.ContentAccess
	Article *model.Article
	Poll    *model.ArticlePollDraft
	Lottery *model.ArticleLotteryDraft
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
	if err = article.ValidateReward(); err != nil {
		return nil, err
	}
	if article.Type == enum.ArticleTypePoll {
		if err := req.Poll.Normalize(time.Now()); err != nil {
			return nil, err
		}
	}
	if article.Type == enum.ArticleTypeLottery {
		if err := req.Lottery.Normalize(time.Now()); err != nil {
			return nil, err
		}
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
		if err != nil {
			return err
		}
		if article.Type == enum.ArticleTypePoll {
			return d.articlePollRepo.SaveDraft(ctx, article.ID, req.Poll)
		}
		if article.Type == enum.ArticleTypeLottery {
			return d.articleLotteryRepo.SaveDraft(ctx, article.ID, req.Lottery)
		}
		if err = d.articlePollRepo.DeleteDraft(ctx, article.ID); err != nil {
			return err
		}
		return d.articleLotteryRepo.DeleteDraft(ctx, article.ID)
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
	interaction := articleInteraction{access: req.Access, articleID: req.ArticleID, userID: req.Access.ActorUserID, active: req.Active, action: enum.ArticleActionLike, stats: &repo.ArticleStatUpdate{LikeCount: 1}, event: d.articleLikedEvent}
	return d.changeInteraction(ctx, interaction)
}

type ArticleThankReq struct {
	Access          *model.ContentAccess
	ArticleID       int64
	UserID          int64
	Active          bool
	ClientRequestID string
}

func (d *ArticleUsecase) Thank(ctx context.Context, req *ArticleThankReq) (bool, error) {
	if !req.Active {
		return false, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return d.thankUsecase.Thank(ctx, req.Access, req.ArticleID, req.ClientRequestID)
}

type ArticleCollectReq struct {
	Access    *model.ContentAccess
	ArticleID int64
	UserID    int64
	Active    bool
}

func (d *ArticleUsecase) Collect(ctx context.Context, req *ArticleCollectReq) (bool, error) {
	interaction := articleInteraction{access: req.Access, articleID: req.ArticleID, userID: req.Access.ActorUserID, active: req.Active, action: enum.ArticleActionCollect, stats: &repo.ArticleStatUpdate{CollectCount: 1}, event: d.articleCollectedEvent}
	return d.changeInteraction(ctx, interaction)
}

type ArticleWatchReq struct {
	Access    *model.ContentAccess
	ArticleID int64
	UserID    int64
	Active    bool
}

func (d *ArticleUsecase) Watch(ctx context.Context, req *ArticleWatchReq) (bool, error) {
	interaction := articleInteraction{
		access:    req.Access,
		articleID: req.ArticleID,
		userID:    req.Access.ActorUserID,
		active:    req.Active,
		action:    enum.ArticleActionWatch,
		stats:     &repo.ArticleStatUpdate{WatchCount: 1},
		event: func(userID, articleID int64) *commonenums.Event {
			return &commonenums.Event{Type: commonenums.EventType_EVENT_TYPE_ARTICLE_WATCHED, Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_WATCHED, Payload: &commonenums.Event_ArticleWatched{ArticleWatched: &commonenums.ArticleWatchedPayload{SenderId: userID, ArticleId: articleID}}}
		},
	}
	return d.changeInteraction(ctx, interaction)
}

type articleInteraction struct {
	access    *model.ContentAccess
	articleID int64
	userID    int64
	active    bool
	action    enum.ArticleAction
	stats     *repo.ArticleStatUpdate
	event     func(int64, int64) *commonenums.Event
}

func (d *ArticleUsecase) changeInteraction(ctx context.Context, interaction articleInteraction) (bool, error) {
	var outboxEvent *repo.OutboxEvent
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: new(interaction.articleID)}})
		if err != nil {
			return err
		}
		if err = d.articleAccessUsecase.CanInteract(interaction.access, article); err != nil {
			return err
		}
		if interaction.active {
			created, err := d.actionRecordRepo.Save(ctx, &model.ArticleActionRecord{ArticleID: interaction.articleID, UserID: interaction.userID, Type: interaction.action, Status: enum.ContentActionStatusPublished})
			if err != nil || !created {
				return err
			}
			if interaction.stats != nil {
				if err = d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{ArticleID: interaction.articleID, Stats: *interaction.stats}); err != nil {
					return err
				}
			}
			if interaction.event == nil {
				return nil
			}
			outboxEvent, err = d.outboxRepo.Save(ctx, interaction.event(interaction.userID, interaction.articleID))
			return err
		}
		deleted, err := d.actionRecordRepo.Delete(ctx, &repo.ArticleActionRecordDeleteReq{ArticleID: interaction.articleID, UserID: interaction.userID, Action: interaction.action})
		if err != nil || deleted == 0 {
			return err
		}
		if interaction.stats == nil {
			return nil
		}
		return d.articleRepo.AddStats(ctx, &repo.ArticleAddStatsReq{ArticleID: interaction.articleID, Stats: d.negateArticleStats(*interaction.stats)})
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

func (*ArticleUsecase) negateArticleStats(stats repo.ArticleStatUpdate) repo.ArticleStatUpdate {
	stats.ViewCount = -stats.ViewCount
	stats.ThankCount = -stats.ThankCount
	stats.LikeCount = -stats.LikeCount
	stats.CollectCount = -stats.CollectCount
	stats.RewardCount = -stats.RewardCount
	stats.ReplyCount = -stats.ReplyCount
	return stats
}

func (*ArticleUsecase) articleLikedEvent(userID, articleID int64) *commonenums.Event {
	return &commonenums.Event{Type: commonenums.EventType_EVENT_TYPE_ARTICLE_LIKED, Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_LIKED, Payload: &commonenums.Event_ArticleLiked{ArticleLiked: &commonenums.ArticleLikedPayload{SenderId: userID, ArticleId: articleID}}}
}

func (*ArticleUsecase) articleThankedEvent(userID, articleID int64) *commonenums.Event {
	return &commonenums.Event{Type: commonenums.EventType_EVENT_TYPE_ARTICLE_THANKED, Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_THANKED, Payload: &commonenums.Event_ArticleThanked{ArticleThanked: &commonenums.ArticleThankedPayload{SenderId: userID, ArticleId: articleID}}}
}

func (*ArticleUsecase) articleCollectedEvent(userID, articleID int64) *commonenums.Event {
	return &commonenums.Event{Type: commonenums.EventType_EVENT_TYPE_ARTICLE_COLLECTED, Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_COLLECTED, Payload: &commonenums.Event_ArticleCollected{ArticleCollected: &commonenums.ArticleCollectedPayload{SenderId: userID, ArticleId: articleID}}}
}

type ArticleRewardReq struct {
	Access          *model.ContentAccess
	ArticleID       int64
	ClientRequestID string
}

func (d *ArticleUsecase) Reward(ctx context.Context, req *ArticleRewardReq) (bool, error) {
	return d.rewardUsecase.Reward(ctx, req.Access, req.ArticleID, req.ClientRequestID)
}

type ArticleViewReq struct {
	Access             *model.ContentAccess
	ArticleID          int64
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
	access, err := req.Access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return err
	}
	if err = article.CanView(access); err != nil {
		return err
	}
	if access.ActorUserID > 0 {
		if err = d.viewRecordRepo.Save(ctx, &model.ArticleViewRecord{
			ArticleID:          req.ArticleID,
			UserID:             access.ActorUserID,
			IP:                 req.IP,
			UserAgent:          req.UserAgent,
			BrowserFingerprint: req.BrowserFingerprint,
			ViewedAt:           new(time.Now()),
		}); err != nil {
			return err
		}
	}
	created, err := d.viewCacheRepo.Record(ctx, &repo.ArticleViewCacheRecordReq{
		ArticleID:          req.ArticleID,
		ViewerUserID:       new(access.ActorUserID),
		IP:                 req.IP,
		UserAgent:          req.UserAgent,
		BrowserFingerprint: req.BrowserFingerprint,
	})
	if err != nil || !created {
		return err
	}
	return nil
}

type ArticleViewHistoryRow struct {
	Article  *model.Article
	ViewedAt *time.Time
}

type ArticleViewHistoryPageResp struct {
	Rows []*ArticleViewHistoryRow
	Page *commonmodel.PageResp
}

func (d *ArticleUsecase) PageViewHistory(
	ctx context.Context,
	access *model.ContentAccess,
	page *commonmodel.PageReq,
	excludedAuthorIDs []int64,
) (*ArticleViewHistoryPageResp, error) {
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	if access.Scope != enum.ContentAccessScopeUser {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	history, err := d.viewRecordRepo.Page(ctx, &repo.ArticleViewRecordPageReq{
		UserID: access.ActorUserID,
		Page:   page,
	})
	if err != nil {
		return nil, err
	}
	articleIDs := make([]int64, 0, len(history.Rows))
	for _, row := range history.Rows {
		if row != nil && row.ArticleID > 0 {
			articleIDs = append(articleIDs, row.ArticleID)
		}
	}
	publishStatus := enum.ArticlePublishStatusPublished
	visibility := enum.ArticleVisibilityPublic
	restriction := enum.ContentRestrictionNone
	articles, err := d.articleRepo.Map(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{
		ArticleIDs:        articleIDs,
		PublishStatus:     &publishStatus,
		Visibility:        &visibility,
		Restriction:       &restriction,
		ExcludedAuthorIDs: excludedAuthorIDs,
	}})
	if err != nil {
		return nil, err
	}
	rows := make([]*ArticleViewHistoryRow, 0, len(history.Rows))
	for _, row := range history.Rows {
		if row == nil || articles[row.ArticleID] == nil {
			continue
		}
		rows = append(rows, &ArticleViewHistoryRow{Article: articles[row.ArticleID], ViewedAt: row.ViewedAt})
	}
	return &ArticleViewHistoryPageResp{Rows: rows, Page: history.Page}, nil
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
	Access          *model.ContentAccess
	ArticleID       int64
	ScheduledAt     *time.Time
	PublishedCity   *string
	ClientRequestID string
}

func (d *ArticleUsecase) Publish(ctx context.Context, req *ArticlePublishReq) error {
	articleID := req.ArticleID
	access, err := req.Access.Normalize(enum.ContentAccessScopeInternalTask)
	if err != nil {
		return err
	}
	now := time.Now()
	if req.ScheduledAt != nil && req.ScheduledAt.After(now) {
		if access.Scope != enum.ContentAccessScopeAuthor {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		operatorUserID := access.ActorUserID
		scheduledAt := req.ScheduledAt.UTC().Truncate(time.Second)
		var originalPublishedAt *time.Time
		var originalPublishedCity *string
		// content 与 scheduler 使用不同的数据存储，不能纳入同一数据库事务。
		// 先提交文章计划，随后登记 Scheduler；登记失败时只回滚仍属于本次计划的草稿。
		if err = d.tx(ctx, func(ctx context.Context) error {
			article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: new(articleID)}})
			if err != nil {
				return err
			}
			if err = article.CanPublish(access, new(scheduledAt), now); err != nil {
				return err
			}
			if err = d.validatePublishActivity(ctx, article, scheduledAt); err != nil {
				return err
			}
			originalPublishedAt = article.PublishedAt
			originalPublishedCity = article.PublishedCity
			return d.articleRepo.UpdatePublishStatus(ctx, &repo.ArticleUpdatePublishStatusReq{
				ArticleID: articleID, PublishStatus: enum.ArticlePublishStatusDraft, Visibility: enum.ArticleVisibilityPublic,
				PublishedAt: new(scheduledAt), PublishedCity: req.PublishedCity,
				ClearPublishedCity: req.PublishedCity == nil, UpdatedBy: new(operatorUserID),
			})
		}); err != nil {
			return err
		}
		if err = d.delayedTaskClient.RegisterPublishScheduledArticle(ctx, articleID, operatorUserID, scheduledAt); err != nil {
			// 重新读取并比对计划时间，避免并发重设后错误清除新计划。
			rollbackErr := d.tx(ctx, func(ctx context.Context) error {
				article, getErr := d.articleRepo.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: new(articleID)}})
				if getErr != nil || article.PublishedAt == nil ||
					!article.PublishedAt.UTC().Truncate(time.Second).Equal(scheduledAt) ||
					(article.PublishedCity == nil) != (req.PublishedCity == nil) ||
					(article.PublishedCity != nil && req.PublishedCity != nil && *article.PublishedCity != *req.PublishedCity) {
					return getErr
				}
				restore := &repo.ArticleUpdatePublishStatusReq{
					ArticleID: articleID, PublishStatus: enum.ArticlePublishStatusDraft, Visibility: enum.ArticleVisibilityPublic,
					PublishedAt: originalPublishedAt, ClearPublished: originalPublishedAt == nil,
					PublishedCity: originalPublishedCity, ClearPublishedCity: originalPublishedCity == nil,
					UpdatedBy: new(operatorUserID),
				}
				return d.articleRepo.UpdatePublishStatus(ctx, restore)
			})
			if rollbackErr != nil {
				d.log.WarnContext(ctx, "rollback scheduled article publish failed", slog.Int64("article_id", articleID), slog.Any("err", rollbackErr))
			}
			return err
		}
		return nil
	}
	var lotteryDrawAt *time.Time
	var pollEndsAt *time.Time
	rewardSetupArticle, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: new(articleID)}})
	if err != nil {
		if req.ScheduledAt != nil {
			if code, ok := apperror.BusinessCode(err); ok && code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND {
				return nil
			}
		}
		return err
	}
	if req.ScheduledAt != nil && access.Scope == enum.ContentAccessScopeInternalTask &&
		(rewardSetupArticle.PublishStatus != enum.ArticlePublishStatusDraft || rewardSetupArticle.PublishedAt == nil || !rewardSetupArticle.PublishedAt.UTC().Truncate(time.Second).Equal(req.ScheduledAt.UTC().Truncate(time.Second))) {
		return nil
	}
	if err = rewardSetupArticle.CanPublish(access, nil, now); err != nil {
		return err
	}
	if err = d.validatePublishActivity(ctx, rewardSetupArticle, now); err != nil {
		return err
	}
	clientRequestID := req.ClientRequestID
	if clientRequestID == "" && access.Scope == enum.ContentAccessScopeInternalTask && req.ScheduledAt != nil {
		clientRequestID = fmt.Sprintf("scheduled-article-publish:%d:%d", articleID, req.ScheduledAt.UTC().Unix())
	}
	if clientRequestID == "" || rewardSetupArticle.CreatedBy == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	var rewardSetupPoints *int64
	if rewardSetupArticle.RewardContent != nil && rewardSetupArticle.RewardPoints != nil {
		points := int64(20)
		if d.business != nil && d.business.GetArticle().GetRewardSetupPoints() > 0 {
			points = d.business.GetArticle().GetRewardSetupPoints()
		}
		rewardSetupPoints = new(points)
	}
	if rewardSetupArticle.Type == enum.ArticleTypeLottery {
		lottery, lotteryErr := d.articleLotteryRepo.Get(ctx, articleID, 0)
		if lotteryErr != nil || lottery == nil {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		lotteryDrawAt = new(lottery.DrawAt)
	}
	if rewardSetupArticle.Type == enum.ArticleTypePoll {
		poll, pollErr := d.articlePollRepo.Get(ctx, articleID, 0)
		if pollErr != nil || poll == nil {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		pollEndsAt = new(poll.EndsAt)
	}
	publish := &contentv1.TCCArticlePublish_Req{ArticleId: articleID, ClientRequestId: clientRequestID, OperatorUserId: *rewardSetupArticle.CreatedBy, City: req.PublishedCity, RewardSetupPoints: rewardSetupPoints, PublishedAt: timestamppb.New(now)}
	_, err = d.dtm.InTCC(ctx, func(tcc *dtmgrpc.TccGrpc) error {
		if err = d.branches.Call(ctx, tcc, constant.ContentServiceName.String(), publish, contentv1.TCCArticlePublishService_Try_FullMethodName, contentv1.TCCArticlePublishService_Confirm_FullMethodName, contentv1.TCCArticlePublishService_Cancel_FullMethodName, &contentv1.TCCArticlePublish_Resp{}); err != nil {
			return err
		}
		if rewardSetupPoints == nil {
			return nil
		}
		charge := &economyv1.TCCPoints_Req{UserId: *rewardSetupArticle.CreatedBy, Amount: *rewardSetupPoints, RecordType: economyv1enum.EconomyRecordType_ECONOMY_RECORD_TYPE_ARTICLE_REWARD_SETUP_EXPENSE, Remark: new(fmt.Sprintf("article reward setup:%d", articleID))}
		return d.branches.Call(ctx, tcc, constant.EconomyServiceName.String(), charge, economyv1.TCCPointsService_Try_FullMethodName, economyv1.TCCPointsService_Confirm_FullMethodName, economyv1.TCCPointsService_Cancel_FullMethodName, &economyv1.TCCPoints_Resp{})
	})
	if err != nil {
		return err
	}
	if lotteryDrawAt != nil {
		if err = d.delayedTaskClient.RegisterLotteryDraw(ctx, articleID, *lotteryDrawAt); err != nil {
			return err
		}
	}
	if pollEndsAt != nil {
		if err = d.delayedTaskClient.RegisterPollFinish(ctx, articleID, *pollEndsAt); err != nil {
			return err
		}
	}
	if req.ScheduledAt == nil {
		if cancelErr := d.delayedTaskClient.CancelPublishScheduledArticle(ctx, articleID); cancelErr != nil {
			d.log.WarnContext(ctx, "cancel scheduled article publish failed", slog.Int64("article_id", articleID), slog.Any("err", cancelErr))
		}
	}
	return nil
}

func (d *ArticleUsecase) validatePublishActivity(ctx context.Context, article *model.Article, publishedAt time.Time) error {
	if article.Type == enum.ArticleTypePoll {
		poll, err := d.articlePollRepo.Get(ctx, article.ID, 0)
		if err != nil {
			return err
		}
		if poll == nil || !poll.EndsAt.After(publishedAt) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
	}
	if article.Type == enum.ArticleTypeLottery {
		lottery, err := d.articleLotteryRepo.Get(ctx, article.ID, 0)
		if err != nil {
			return err
		}
		if lottery == nil || !lottery.DrawAt.After(publishedAt) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
	}
	return nil
}

type ArticleCancelPublishReq struct {
	Access    *model.ContentAccess
	ArticleID int64
}

func (d *ArticleUsecase) CancelPublish(ctx context.Context, req *ArticleCancelPublishReq) error {
	articleID := req.ArticleID
	access, err := req.Access.Normalize("")
	if err != nil {
		return err
	}
	operatorUserID := access.ActorUserID
	err = d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleID)},
		})
		if err != nil {
			return err
		}
		if err = article.CanCancelSchedule(operatorUserID); err != nil {
			return err
		}
		return d.articleRepo.UpdatePublishStatus(ctx, &repo.ArticleUpdatePublishStatusReq{
			ArticleID:          articleID,
			PublishStatus:      enum.ArticlePublishStatusDraft,
			Visibility:         article.Visibility,
			ClearPublished:     true,
			ClearPublishedCity: true,
			UpdatedBy:          new(operatorUserID),
		})
	})
	if err != nil {
		return err
	}
	return d.delayedTaskClient.CancelPublishScheduledArticle(ctx, articleID)
}

type ArticleMakePrivateReq struct {
	Access    *model.ContentAccess
	ArticleID int64
}

func (d *ArticleUsecase) MakePrivate(ctx context.Context, req *ArticleMakePrivateReq) error {
	articleID := req.ArticleID
	userID := req.Access.ActorUserID
	return d.updateVisibility(ctx, &articleUpdateVisibilityReq{
		ArticleID:  articleID,
		Visibility: enum.ArticleVisibilityPrivate,
		UserID:     userID,
	})
}

type ArticleMakePublicReq struct {
	Access    *model.ContentAccess
	ArticleID int64
}

func (d *ArticleUsecase) MakePublic(ctx context.Context, req *ArticleMakePublicReq) error {
	articleID := req.ArticleID
	userID := req.Access.ActorUserID
	return d.updateVisibility(ctx, &articleUpdateVisibilityReq{
		ArticleID:  articleID,
		Visibility: enum.ArticleVisibilityPublic,
		UserID:     userID,
	})
}

type articleUpdateVisibilityReq struct {
	ArticleID  int64
	Visibility enum.ArticleVisibility
	UserID     int64
}

func (d *ArticleUsecase) updateVisibility(ctx context.Context, req *articleUpdateVisibilityReq) error {
	articleID := req.ArticleID
	visibility := req.Visibility
	userID := req.UserID
	var outboxEvent *repo.OutboxEvent
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleID)},
		})
		if err != nil {
			return err
		}
		if !d.isAuthor(article, userID) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		if article.PublishStatus != enum.ArticlePublishStatusPublished || article.Restriction != enum.ContentRestrictionNone {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		if err := d.articleRepo.UpdateVisibility(ctx, &repo.ArticleUpdateVisibilityReq{
			ArticleID:  articleID,
			Visibility: visibility,
			UpdatedBy:  userID,
		}); err != nil {
			return err
		}
		outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_STATUS_UPDATED,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_STATUS_UPDATED,
			Payload: &commonenums.Event_ArticleStatusUpdated{
				ArticleStatusUpdated: &commonenums.ArticleStatusUpdatedPayload{
					SenderId:      userID,
					ArticleId:     articleID,
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
	articleID := req.ArticleID
	userID := req.Access.ActorUserID
	reason := req.Reason
	return d.updatePublishStatus(ctx, &articleUpdatePublishStatusReq{
		ArticleID:     articleID,
		PublishStatus: enum.ArticlePublishStatusArchived,
		UserID:        userID,
		Action:        "archived",
		Reason:        reason,
	})
}

// SetPinnedComment 设置或取消作者文章的一条置顶一级评论。
func (d *ArticleUsecase) SetPinnedComment(ctx context.Context, access *model.ContentAccess, articleID int64, commentID *int64) error {
	access, err := access.Normalize(enum.ContentAccessScopeAuthor)
	if err != nil {
		return err
	}
	return d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: &articleID}})
		if err != nil {
			return err
		}
		if !article.IsAuthor(access.ActorUserID) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		if commentID != nil {
			comment, err := d.commentRepo.Get(ctx, &repo.CommentGetReq{Filter: &model.CommentFilter{CommentID: commentID, ArticleID: &articleID, Restriction: new(enum.ContentRestrictionNone)}})
			if err != nil {
				return err
			}
			if comment.ParentID != nil {
				return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_INVALID)
			}
		}
		return d.articleRepo.UpdatePinnedComment(ctx, &repo.ArticleUpdatePinnedCommentReq{ArticleID: articleID, PinnedID: commentID, UpdatedBy: access.ActorUserID})
	})
}

func (d *ArticleUsecase) AcceptAnswer(ctx context.Context, access *model.ContentAccess, articleID, commentID int64) error {
	access, err := access.Normalize(enum.ContentAccessScopeAuthor)
	if err != nil {
		return err
	}
	var outboxEvent *repo.OutboxEvent
	err = d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: &articleID}})
		if err != nil {
			return err
		}
		if article.Type != enum.ArticleTypeQA || !article.IsAuthor(access.ActorUserID) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		if accepted, err := d.articleQARepo.GetAcceptedCommentID(ctx, articleID); err != nil {
			return err
		} else if accepted != nil {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		comment, err := d.commentRepo.Get(ctx, &repo.CommentGetReq{Filter: &model.CommentFilter{CommentID: &commentID, ArticleID: &articleID, Restriction: new(enum.ContentRestrictionNone)}})
		if err != nil {
			return err
		}
		if comment.ParentID != nil || comment.CreatedBy == nil || *comment.CreatedBy == access.ActorUserID {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_INVALID)
		}
		if err = d.articleQARepo.Accept(ctx, articleID, commentID, time.Now().UTC()); err != nil {
			return err
		}
		outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{Type: commonenums.EventType_EVENT_TYPE_ARTICLE_ACCEPTED_ANSWER, Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_ACCEPTED_ANSWER, Payload: &commonenums.Event_ArticleAcceptedAnswer{ArticleAcceptedAnswer: &commonenums.ArticleAcceptedAnswerPayload{SenderId: access.ActorUserID, ArticleId: articleID, CommentId: commentID}}})
		return err
	})
	if err != nil {
		return err
	}
	if outboxEvent != nil {
		if _, publishErr := d.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{ID: outboxEvent.ID}); publishErr != nil {
			d.log.WarnContext(ctx, "publish accepted answer outbox event failed", slog.Int64("outbox_id", outboxEvent.ID), slog.Any("err", publishErr))
		}
	}
	return nil
}

type ArticleUnarchiveReq struct {
	Access    *model.ContentAccess
	ArticleID int64
	Reason    *string
}

func (d *ArticleUsecase) Unarchive(ctx context.Context, req *ArticleUnarchiveReq) error {
	articleID := req.ArticleID
	userID := req.Access.ActorUserID
	reason := req.Reason
	return d.updatePublishStatus(ctx, &articleUpdatePublishStatusReq{
		ArticleID:     articleID,
		PublishStatus: enum.ArticlePublishStatusPublished,
		UserID:        userID,
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
	articleID := req.ArticleID
	publishStatus := req.PublishStatus
	userID := req.UserID
	action := req.Action
	reason := req.Reason
	var outboxEvent *repo.OutboxEvent
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleID)},
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
			ArticleID:     articleID,
			PublishStatus: publishStatus,
			Visibility:    article.Visibility,
			UpdatedBy:     new(userID),
		}); err != nil {
			return err
		}
		outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_STATUS_UPDATED,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_STATUS_UPDATED,
			Payload: &commonenums.Event_ArticleStatusUpdated{
				ArticleStatusUpdated: &commonenums.ArticleStatusUpdatedPayload{
					SenderId:      userID,
					ArticleId:     articleID,
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
	articleID := req.ArticleID
	userID := req.Access.ActorUserID
	reason := req.Reason
	return d.updateRestriction(ctx, &articleUpdateRestrictionReq{
		ArticleID:   articleID,
		Restriction: enum.ContentRestrictionHidden,
		UserID:      userID,
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
	articleID := req.ArticleID
	userID := req.Access.ActorUserID
	reason := req.Reason
	return d.updateRestriction(ctx, &articleUpdateRestrictionReq{
		ArticleID:   articleID,
		Restriction: enum.ContentRestrictionNone,
		UserID:      userID,
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
	articleID := req.ArticleID
	userID := req.Access.ActorUserID
	reason := req.Reason
	return d.updateRestriction(ctx, &articleUpdateRestrictionReq{
		ArticleID:   articleID,
		Restriction: enum.ContentRestrictionLocked,
		UserID:      userID,
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
	articleID := req.ArticleID
	userID := req.Access.ActorUserID
	reason := req.Reason
	return d.updateRestriction(ctx, &articleUpdateRestrictionReq{
		ArticleID:   articleID,
		Restriction: enum.ContentRestrictionNone,
		UserID:      userID,
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
	articleID := req.ArticleID
	restriction := req.Restriction
	userID := req.UserID
	action := req.Action
	reason := req.Reason
	var outboxEvent *repo.OutboxEvent
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleID)},
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
			ArticleID:   articleID,
			Restriction: restriction,
			UpdatedBy:   userID,
		}); err != nil {
			return err
		}
		if _, err := d.moderationRecordRepo.Save(ctx, &model.ContentModerationRecord{
			Target:     enum.ContentModerationTargetArticle,
			TargetID:   articleID,
			Action:     action,
			Reason:     reason,
			OperatorID: userID,
		}); err != nil {
			return err
		}
		outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{
			Type:    commonenums.EventType_EVENT_TYPE_ARTICLE_STATUS_UPDATED,
			Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_STATUS_UPDATED,
			Payload: &commonenums.Event_ArticleStatusUpdated{
				ArticleStatusUpdated: &commonenums.ArticleStatusUpdatedPayload{
					SenderId:      userID,
					ArticleId:     articleID,
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
	article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
		Filter: &model.ArticleFilter{ArticleID: new(articleID)},
	})
	if err != nil {
		return nil, err
	}
	return article, nil
}

type ArticleVotePollReq struct {
	Access    *model.ContentAccess
	ArticleID int64
	OptionIDs []int64
}

type ArticleFinishPollReq struct {
	ArticleID int64
	EndsAt    time.Time
}

func (d *ArticleUsecase) FinishPoll(ctx context.Context, req *ArticleFinishPollReq) error {
	var outboxEvent *repo.OutboxEvent
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: new(req.ArticleID)}})
		if err != nil {
			return err
		}
		if article.Type != enum.ArticleTypePoll || article.PublishStatus != enum.ArticlePublishStatusPublished || article.CreatedBy == nil {
			return nil
		}
		poll, err := d.articlePollRepo.Get(ctx, article.ID, 0)
		if err != nil {
			return err
		}
		if poll == nil || !poll.EndsAt.UTC().Equal(req.EndsAt.UTC()) || time.Now().UTC().Before(poll.EndsAt) {
			return nil
		}
		outboxEvent, err = d.outboxRepo.Save(ctx, &commonenums.Event{Type: commonenums.EventType_EVENT_TYPE_ARTICLE_POLL_ENDED, Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_POLL_ENDED, Payload: &commonenums.Event_ArticlePollEnded{ArticlePollEnded: &commonenums.ArticlePollEndedPayload{SenderId: *article.CreatedBy, ArticleId: article.ID}}})
		return err
	})
	if err != nil {
		return err
	}
	if outboxEvent != nil {
		if _, publishErr := d.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{ID: outboxEvent.ID}); publishErr != nil {
			d.log.WarnContext(ctx, "publish poll ended outbox event failed", slog.Int64("outbox_id", outboxEvent.ID), slog.Any("err", publishErr))
		}
	}
	return nil
}

type ArticleEnterLotteryReq struct {
	Access    *model.ContentAccess
	ArticleID int64
	Entered   bool
}

func (d *ArticleUsecase) EnterLottery(ctx context.Context, req *ArticleEnterLotteryReq) error {
	access, err := req.Access.Normalize("")
	if err != nil {
		return err
	}
	return d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: new(req.ArticleID)}})
		if err != nil {
			return err
		}
		if article.Type != enum.ArticleTypeLottery || article.IsAuthor(access.ActorUserID) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		if err = d.articleAccessUsecase.CanInteract(access, article); err != nil {
			return err
		}
		lottery, err := d.articleLotteryRepo.Get(ctx, article.ID, 0)
		if err != nil {
			return err
		}
		if lottery == nil {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		err = d.articleLotteryRepo.SetEntry(ctx, lottery.ID, access.ActorUserID, req.Entered, time.Now())
		if _, ok := errors.AsType[*repo.ArticleLotteryClosedError](err); ok {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		return err
	})
}

type ArticleDrawLotteryReq struct {
	ArticleID int64
	DrawAt    time.Time
}

func (d *ArticleUsecase) DrawLottery(ctx context.Context, req *ArticleDrawLotteryReq) error {
	outboxEvents := make([]*repo.OutboxEvent, 0)
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: new(req.ArticleID)}})
		if err != nil {
			return err
		}
		if article.Type != enum.ArticleTypeLottery || article.PublishStatus != enum.ArticlePublishStatusPublished {
			return nil
		}
		lottery, err := d.articleLotteryRepo.Get(ctx, article.ID, 0)
		if err != nil {
			return err
		}
		if lottery == nil || lottery.DrawnAt != nil || !lottery.DrawAt.UTC().Equal(req.DrawAt.UTC()) {
			return nil
		}
		winners, drawErr := d.articleLotteryRepo.Draw(ctx, lottery.ID, time.Now().UTC())
		if drawErr != nil {
			return drawErr
		}
		if article.CreatedBy == nil {
			return nil
		}
		for _, winner := range winners {
			if winner == nil || winner.UserID <= 0 {
				continue
			}
			event, saveErr := d.outboxRepo.Save(ctx, &commonenums.Event{Type: commonenums.EventType_EVENT_TYPE_ARTICLE_LOTTERY_WON, Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_LOTTERY_WON, Payload: &commonenums.Event_ArticleLotteryWon{ArticleLotteryWon: &commonenums.ArticleLotteryWonPayload{SenderId: *article.CreatedBy, ArticleId: article.ID, PrizeId: winner.PrizeID, WinnerId: winner.UserID}}})
			if saveErr != nil {
				return saveErr
			}
			outboxEvents = append(outboxEvents, event)
		}
		event, saveErr := d.outboxRepo.Save(ctx, &commonenums.Event{Type: commonenums.EventType_EVENT_TYPE_ARTICLE_LOTTERY_DRAWN, Subject: commonenums.EventSubject_EVENT_SUBJECT_ARTICLE_LOTTERY_DRAWN, Payload: &commonenums.Event_ArticleLotteryDrawn{ArticleLotteryDrawn: &commonenums.ArticleLotteryDrawnPayload{SenderId: *article.CreatedBy, ArticleId: article.ID}}})
		if saveErr != nil {
			return saveErr
		}
		outboxEvents = append(outboxEvents, event)
		return nil
	})
	if err != nil {
		return err
	}
	for _, event := range outboxEvents {
		if event != nil {
			if _, publishErr := d.outboxUsecase.Publish(ctx, &PublishOutboxEventReq{ID: event.ID}); publishErr != nil {
				d.log.WarnContext(ctx, "publish lottery winner outbox event failed", slog.Int64("outbox_id", event.ID), slog.Any("err", publishErr))
			}
		}
	}
	return nil
}

func (d *ArticleUsecase) VotePoll(ctx context.Context, req *ArticleVotePollReq) error {
	access, err := req.Access.Normalize("")
	if err != nil {
		return err
	}
	return d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{Filter: &model.ArticleFilter{ArticleID: new(req.ArticleID)}})
		if err != nil {
			return err
		}
		if article.Type != enum.ArticleTypePoll {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		if article.IsAuthor(access.ActorUserID) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		if err = d.articleAccessUsecase.CanInteract(access, article); err != nil {
			return err
		}
		poll, err := d.articlePollRepo.Get(ctx, article.ID, 0)
		if err != nil {
			return err
		}
		if err = poll.ValidateVote(req.OptionIDs, time.Now()); err != nil {
			return err
		}
		return d.articlePollRepo.Vote(ctx, poll.ID, access.ActorUserID, req.OptionIDs)
	})
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
	if article.Type == enum.ArticleTypeQA {
		article.AcceptedCommentID, err = d.articleQARepo.GetAcceptedCommentID(ctx, article.ID)
		if err != nil {
			return nil, err
		}
	}
	if article.Type == enum.ArticleTypePoll {
		article.Poll, err = d.articlePollRepo.Get(ctx, article.ID, req.Access.ActorUserID)
		if err != nil {
			return nil, err
		}
		if article.Poll != nil {
			article.Poll.ResultVisible = article.Poll.CanRevealResults(util.DerefOrDefault(article.CreatedBy, 0), req.Access.ActorUserID, time.Now())
			if !article.Poll.ResultVisible {
				for _, option := range article.Poll.Options {
					option.VoteCount = 0
				}
			}
		}
	}
	if article.Type == enum.ArticleTypeLottery {
		article.Lottery, err = d.articleLotteryRepo.Get(ctx, article.ID, req.Access.ActorUserID)
		if err != nil {
			return nil, err
		}
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
	scope, err := d.articleAccessUsecase.BuildScope(req.Access, req.Filter)
	if err != nil {
		return nil, err
	}
	return d.articleRepo.List(ctx, &repo.ArticleGetReq{Filter: req.Filter, Scope: scope})
}

type ArticlePageByAccessReq struct {
	Access *model.ContentAccess
	Filter *model.ArticleFilter
	Page   *commonmodel.PageReq
}

func (d *ArticleUsecase) PageByAccess(ctx context.Context, req *ArticlePageByAccessReq) (*ArticlePageResp, error) {
	if req == nil {
		req = &ArticlePageByAccessReq{}
	}
	scope, err := d.articleAccessUsecase.BuildScope(req.Access, req.Filter)
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
	PublishedAtEnd  *time.Time
	Scheduled       *bool
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
		PublishedAtEnd:  req.PublishedAtEnd,
		Scheduled:       req.Scheduled,
	}})
}

type ArticlePageReq struct {
	Page            *commonmodel.PageReq
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
	PublishedAtEnd  *time.Time
	Scheduled       *bool
	ArticleIDs      []int64
}

type ArticlePageResp struct {
	Rows []*model.Article
	Page *commonmodel.PageResp
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
		PublishedAtEnd:  req.PublishedAtEnd,
		Scheduled:       req.Scheduled,
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

func (d *ArticleUsecase) ListWatchers(ctx context.Context, articleID int64) ([]int64, error) {
	if articleID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	action := enum.ArticleActionWatch
	records, err := d.actionRecordRepo.List(ctx, &repo.ArticleActionRecordReq{ArticleIDs: []int64{articleID}, Type: &action})
	if err != nil {
		return nil, err
	}
	userIDs := make([]int64, 0, len(records))
	for _, record := range records {
		if record != nil && record.UserID > 0 {
			userIDs = append(userIDs, record.UserID)
		}
	}
	return lo.Uniq(userIDs), nil
}

func (d *ArticleUsecase) ListPollParticipants(ctx context.Context, articleID int64) ([]int64, error) {
	if articleID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return d.articlePollRepo.ListParticipantUserIDs(ctx, articleID)
}

func (d *ArticleUsecase) ListLotteryParticipants(ctx context.Context, articleID int64) ([]int64, error) {
	if articleID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return d.articleLotteryRepo.ListParticipantUserIDs(ctx, articleID)
}

func (d *ArticleUsecase) MapViewerActionStates(ctx context.Context, req *ArticleMapViewerActionStatesReq) (map[int64]*model.
	ArticleViewerActionState, error) {
	articleIDs := req.ArticleIDs
	userID := req.UserID
	articleIDs = lo.Uniq(articleIDs)
	if len(articleIDs) == 0 {
		return map[int64]*model.ArticleViewerActionState{}, nil
	}
	states := lo.SliceToMap(articleIDs, func(articleID int64) (int64, *model.ArticleViewerActionState) {
		return articleID, &model.ArticleViewerActionState{}
	})
	records, err := d.actionRecordRepo.List(ctx, &repo.ArticleActionRecordReq{
		ArticleIDs: articleIDs,
		UserID:     new(userID),
		Status:     new(enum.ContentActionStatusPublished),
		Types: []enum.ArticleAction{
			enum.ArticleActionLike,
			enum.ArticleActionThank,
			enum.ArticleActionCollect,
			enum.ArticleActionReward,
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
		case enum.ArticleActionReward:
			state.Rewarded = true
		case enum.ArticleActionWatch:
			state.Watched = true
		}
	}
	return states, nil
}

type ArticleDiscardDraftReq struct {
	Access    *model.ContentAccess
	ArticleID int64
}

func (d *ArticleUsecase) DiscardDraft(ctx context.Context, req *ArticleDiscardDraftReq) error {
	articleID := req.ArticleID
	userID := req.Access.ActorUserID
	err := d.tx(ctx, func(ctx context.Context) error {
		article, err := d.articleRepo.Get(ctx, &repo.ArticleGetReq{
			Filter: &model.ArticleFilter{ArticleID: new(articleID)},
		})
		if err != nil {
			return err
		}
		if !d.isAuthor(article, userID) {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		if article.PublishStatus != enum.ArticlePublishStatusDraft {
			return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_STATUS_CONFLICT)
		}
		return d.articleRepo.DiscardDraft(ctx, articleID)
	})
	if err != nil {
		return err
	}
	if err = d.delayedTaskClient.CancelPublishScheduledArticle(ctx, articleID); err != nil {
		d.log.WarnContext(ctx, "cancel scheduled article publish failed", slog.Int64("article_id", articleID), slog.Any("err", err))
	}
	return nil
}

func (d *ArticleUsecase) isAuthor(article *model.Article, userID int64) bool {
	return article != nil && article.CreatedBy != nil && *article.CreatedBy == userID
}
