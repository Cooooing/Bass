package usecase

import (
	"common/pkg/apperror"
	commonmodel "common/pkg/model"
	cerrors "common/proto/gen/common/errors"
	"context"
	"fmt"
	"strings"
	"time"

	"bff_bbs/internal/biz/model"

	"bff_bbs/internal/biz/repo"
)

type ContentArticleUsecase struct {
	contentArticleClient         repo.ContentArticleClient
	assetClient                  repo.AssetClient
	privacyClient                repo.PrivacySettingClient
	relationClient               repo.RelationClient
	locationClient               repo.LocationClient
	ipResolutionClient           repo.IPResolutionClient
	articleSearchRateLimitCache  repo.ArticleSearchRateLimitCache
	articleSearchShortQueryLimit int64
	articleSearchLimit           int64
}

func NewContentArticleUsecase(
	contentArticleClient repo.ContentArticleClient,
	assetClient repo.AssetClient,
	privacyClient repo.PrivacySettingClient,
	relationClient repo.RelationClient,
	locationClient repo.LocationClient,
	ipResolutionClient repo.IPResolutionClient,
	articleSearchRateLimitCache repo.ArticleSearchRateLimitCache,
) *ContentArticleUsecase {
	return &ContentArticleUsecase{
		contentArticleClient:         contentArticleClient,
		assetClient:                  assetClient,
		privacyClient:                privacyClient,
		relationClient:               relationClient,
		locationClient:               locationClient,
		ipResolutionClient:           ipResolutionClient,
		articleSearchRateLimitCache:  articleSearchRateLimitCache,
		articleSearchShortQueryLimit: 10,
		articleSearchLimit:           30,
	}
}

func (u *ContentArticleUsecase) ListNearbyArticles(ctx context.Context, userID int64, page *commonmodel.PageReq) (*ListArticlesResp, error) {
	if userID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	location, err := u.locationClient.GetCurrentLocation(ctx, userID)
	if err != nil {
		return nil, err
	}
	if location == nil || location.City == nil || strings.TrimSpace(*location.City) == "" {
		return &ListArticlesResp{Page: &commonmodel.PageResp{Page: commonmodel.NormalizePage(page).Page, Size: commonmodel.NormalizePage(page).Size}}, nil
	}
	city := strings.TrimSpace(*location.City)
	return u.ListArticles(ctx, &ListArticlesReq{UserID: userID, Page: page, Query: &ArticleQuery{PublishedCity: &city}})
}

func (u *ContentArticleUsecase) ListFollowingArticles(ctx context.Context, userID int64, page *commonmodel.PageReq) (*ListArticlesResp, error) {
	if userID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	var authorIDs []int64
	for currentPage := commonmodel.DefaultPage; ; currentPage++ {
		result, err := u.relationClient.ListFollowing(ctx, &repo.ListFollowingRelationsReq{ActorID: userID, Page: &commonmodel.PageReq{Page: currentPage, Size: 100}})
		if err != nil {
			return nil, err
		}
		for _, row := range result.Rows {
			if row != nil && row.TargetID > 0 {
				authorIDs = append(authorIDs, row.TargetID)
			}
		}
		if result.Page == nil || result.Page.Size == 0 || len(result.Rows) == 0 || currentPage >= (result.Page.Total+result.Page.Size-1)/result.Page.Size {
			break
		}
	}
	if len(authorIDs) == 0 {
		normalizedPage := commonmodel.NormalizePage(page)
		return &ListArticlesResp{Page: &commonmodel.PageResp{Page: normalizedPage.Page, Size: normalizedPage.Size}}, nil
	}
	return u.ListArticles(ctx, &ListArticlesReq{UserID: userID, Page: page, Query: &ArticleQuery{AuthorIDs: authorIDs}})
}

type ContentArticleSave struct {
	Title         string
	Content       string
	RewardContent *string
	RewardPoints  *int32
	Type          int32
	Statement     *string
	Commentable   *bool
	Poll          *model.ArticlePollDraft
	Lottery       *model.ArticleLotteryDraft
}

type CreateDraftArticleReq struct {
	UserID  int64
	Article *ContentArticleSave
}

func (u *ContentArticleUsecase) CreateDraftArticle(ctx context.Context, req *CreateDraftArticleReq) (int64, error) {
	if req == nil || req.Article == nil {
		return 0, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_INVALID)
	}
	return u.contentArticleClient.CreateDraftArticle(ctx, &repo.CreateDraftArticleReq{UserID: req.UserID, Article: &repo.ArticleSave{Title: req.Article.Title, Content: req.Article.Content, RewardContent: req.Article.RewardContent, RewardPoints: req.Article.RewardPoints, Type: req.Article.Type, Statement: req.Article.Statement, Commentable: req.Article.Commentable, Poll: req.Article.Poll, Lottery: req.Article.Lottery}})
}

type UpdateDraftArticleReq struct {
	UserID    int64
	ArticleID int64
	Article   *ContentArticleSave
}

func (u *ContentArticleUsecase) UpdateDraftArticle(ctx context.Context, req *UpdateDraftArticleReq) (int64, error) {
	if req == nil || req.Article == nil || req.ArticleID <= 0 {
		return 0, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_INVALID)
	}
	return u.contentArticleClient.UpdateDraftArticle(ctx, &repo.UpdateDraftArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Article: &repo.ArticleSave{Title: req.Article.Title, Content: req.Article.Content, RewardContent: req.Article.RewardContent, RewardPoints: req.Article.RewardPoints, Type: req.Article.Type, Statement: req.Article.Statement, Commentable: req.Article.Commentable, Poll: req.Article.Poll, Lottery: req.Article.Lottery}})
}

type VoteArticlePollReq struct {
	UserID    int64
	ArticleID int64
	OptionIDs []int64
}

func (u *ContentArticleUsecase) VoteArticlePoll(ctx context.Context, req *VoteArticlePollReq) error {
	if req == nil || req.ArticleID <= 0 || len(req.OptionIDs) == 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_INVALID)
	}
	return u.contentArticleClient.VoteArticlePoll(ctx, &repo.VoteArticlePollReq{UserID: req.UserID, ArticleID: req.ArticleID, OptionIDs: req.OptionIDs})
}

type EnterArticleLotteryReq struct {
	UserID    int64
	ArticleID int64
	Entered   bool
}

func (u *ContentArticleUsecase) EnterArticleLottery(ctx context.Context, req *EnterArticleLotteryReq) error {
	if req == nil || req.ArticleID <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_INVALID)
	}
	return u.contentArticleClient.EnterArticleLottery(ctx, &repo.EnterArticleLotteryReq{UserID: req.UserID, ArticleID: req.ArticleID, Entered: req.Entered})
}

type PublishArticleReq struct {
	UserID          int64
	ArticleID       int64
	ScheduledAt     *time.Time
	IP              string
}

func (u *ContentArticleUsecase) PublishArticle(ctx context.Context, req *PublishArticleReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_INVALID)
	}
	var city *string
	if ip := strings.TrimSpace(req.IP); ip != "" && u.ipResolutionClient != nil {
		resolved, err := u.ipResolutionClient.Resolve(ctx, ip)
		if err == nil && resolved != nil {
			value := []rune(strings.TrimSpace(resolved.City))
			if len(value) > 128 {
				value = value[:128]
			}
			if len(value) > 0 {
				resolvedCity := string(value)
				city = &resolvedCity
			}
		}
	}
	return u.contentArticleClient.PublishArticle(ctx, &repo.PublishArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, ScheduledAt: req.ScheduledAt, City: city})
}

type CancelPublishArticleReq struct {
	UserID    int64
	ArticleID int64
}

func (u *ContentArticleUsecase) CancelPublishArticle(ctx context.Context, req *CancelPublishArticleReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_INVALID)
	}
	return u.contentArticleClient.CancelPublishArticle(ctx, &repo.CancelPublishArticleReq{UserID: req.UserID, ArticleID: req.ArticleID})
}

type DiscardDraftArticleReq struct {
	UserID    int64
	ArticleID int64
}

func (u *ContentArticleUsecase) DiscardDraftArticle(ctx context.Context, req *DiscardDraftArticleReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_INVALID)
	}
	return u.contentArticleClient.DiscardDraftArticle(ctx, &repo.DiscardDraftArticleReq{UserID: req.UserID, ArticleID: req.ArticleID})
}

type ArchiveArticleReq struct {
	UserID    int64
	ArticleID int64
	Reason    *string
}

func (u *ContentArticleUsecase) ArchiveArticle(ctx context.Context, req *ArchiveArticleReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_INVALID)
	}
	return u.contentArticleClient.ArchiveArticle(ctx, &repo.ArchiveArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Reason: req.Reason})
}

type SetPinnedArticleCommentReq struct {
	UserID    int64
	ArticleID int64
	CommentID *int64
}

func (u *ContentArticleUsecase) SetPinnedArticleComment(ctx context.Context, req *SetPinnedArticleCommentReq) error {
	if req == nil || req.ArticleID <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_INVALID)
	}
	return u.contentArticleClient.SetPinnedArticleComment(ctx, &repo.SetPinnedArticleCommentReq{UserID: req.UserID, ArticleID: req.ArticleID, CommentID: req.CommentID})
}

func (u *ContentArticleUsecase) AcceptArticleAnswer(ctx context.Context, userID, articleID, commentID int64) error {
	if articleID <= 0 || commentID <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_INVALID)
	}
	return u.contentArticleClient.AcceptArticleAnswer(ctx, &repo.AcceptArticleAnswerReq{UserID: userID, ArticleID: articleID, CommentID: commentID})
}

type ListArticlesReq struct {
	UserID int64
	Page   *commonmodel.PageReq
	Query  *ArticleQuery
}

type ArticleQuery struct {
	TagID           *int64
	DomainID        *int64
	AuthorID        *int64
	AuthorIDs       []int64
	PublishedCity   *string
	Type            *int32
	Order           *int32
	PublishStatus   *int32
	PublishStatuses []int32
	Visibility      *int32
	Visibilities    []int32
	Scheduled       *bool
}

type ListArticlesResp struct {
	Page *commonmodel.PageResp
	Rows []*model.ArticleListItem
}

type SearchArticlesReq struct {
	UserID   int64
	ClientIP string
	Query    string
	Cursor   *string
	Size     uint32
}

type SearchArticlesResp struct {
	Rows       []*model.ArticleListItem
	NextCursor *string
}

type ArticleViewHistoryPageResp struct {
	Rows []*repo.ArticleViewHistoryRow
	Page *commonmodel.PageResp
}

func (u *ContentArticleUsecase) PageViewHistory(
	ctx context.Context,
	userID int64,
	page *commonmodel.PageReq,
) (*ArticleViewHistoryPageResp, error) {
	if userID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	var requestPage *commonmodel.PageReq
	if page != nil {
		requestPage = &commonmodel.PageReq{Page: page.Page, Size: page.Size}
	}
	excludedAuthorIDs := []int64{}
	if userID > 0 && u.relationClient != nil {
		var err error
		excludedAuthorIDs, err = u.relationClient.ListBlockedCounterparts(ctx, userID)
		if err != nil {
			return nil, err
		}
	}
	result, err := u.contentArticleClient.PageViewHistory(ctx, userID, requestPage, excludedAuthorIDs)
	if err != nil {
		return nil, err
	}
	profiles := make([]*model.AccountProfile, 0, len(result.Rows)*2)
	for _, row := range result.Rows {
		if row == nil || row.Article == nil {
			continue
		}
		profiles = append(profiles, row.Article.AuthorUser, row.Article.LastReplyUser)
	}
	if err = u.hydrateArticleProfiles(ctx, profiles); err != nil {
		return nil, err
	}
	return &ArticleViewHistoryPageResp{Rows: result.Rows, Page: result.Page}, nil
}

func (u *ContentArticleUsecase) ListArticles(ctx context.Context, req *ListArticlesReq) (*ListArticlesResp, error) {
	if req == nil {
		req = &ListArticlesReq{}
	}
	if req.Query != nil && req.Query.AuthorID != nil && req.UserID != *req.Query.AuthorID {
		privacy, err := u.privacyClient.GetCurrentPrivacySetting(ctx, *req.Query.AuthorID)
		if err != nil {
			return nil, err
		}
		if privacy.PublicArticleList != nil && !*privacy.PublicArticleList {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_ARTICLE_LIST_PRIVATE)
		}
	}
	var page *commonmodel.PageReq
	if req.Page != nil {
		page = &commonmodel.PageReq{Page: req.Page.Page, Size: req.Page.Size}
	}
	var query *repo.ArticleQuery
	if req.Query != nil {
		query = &repo.ArticleQuery{TagID: req.Query.TagID, DomainID: req.Query.DomainID, AuthorID: req.Query.AuthorID, AuthorIDs: req.Query.AuthorIDs, PublishedCity: req.Query.PublishedCity, Type: req.Query.Type, Order: req.Query.Order, PublishStatus: req.Query.PublishStatus, PublishStatuses: req.Query.PublishStatuses, Visibility: req.Query.Visibility, Visibilities: req.Query.Visibilities, Scheduled: req.Query.Scheduled}
	}
	if query != nil {
		excludedAuthorIDs := []int64{}
		if req.UserID > 0 && u.relationClient != nil {
			var err error
			excludedAuthorIDs, err = u.relationClient.ListBlockedCounterparts(ctx, req.UserID)
			if err != nil {
				return nil, err
			}
		}
		query.ExcludedAuthorIDs = excludedAuthorIDs
	}
	resp, err := u.contentArticleClient.ListArticles(ctx, &repo.ListArticlesReq{UserID: req.UserID, Page: page, Query: query})
	if err != nil {
		return nil, err
	}
	profiles := make([]*model.AccountProfile, 0, len(resp.Rows)*2)
	for _, row := range resp.Rows {
		if row == nil {
			continue
		}
		profiles = append(profiles, row.AuthorUser, row.LastReplyUser)
	}
	if err = u.hydrateArticleProfiles(ctx, profiles); err != nil {
		return nil, err
	}
	return &ListArticlesResp{Page: resp.Page, Rows: resp.Rows}, nil
}

func (u *ContentArticleUsecase) SearchArticles(ctx context.Context, req *SearchArticlesReq) (*SearchArticlesResp, error) {
	if req == nil {
		req = &SearchArticlesReq{}
	}
	if err := u.allowArticleSearch(ctx, req); err != nil {
		return nil, err
	}
	excludedAuthorIDs := []int64{}
	if req.UserID > 0 && u.relationClient != nil {
		var err error
		excludedAuthorIDs, err = u.relationClient.ListBlockedCounterparts(ctx, req.UserID)
		if err != nil {
			return nil, err
		}
	}
	result, err := u.contentArticleClient.SearchArticles(ctx, &repo.SearchArticlesReq{
		UserID:            req.UserID,
		ClientIP:          req.ClientIP,
		Query:             req.Query,
		Cursor:            req.Cursor,
		Size:              req.Size,
		ExcludedAuthorIDs: excludedAuthorIDs,
	})
	if err != nil {
		return nil, err
	}
	profiles := make([]*model.AccountProfile, 0, len(result.Rows)*2)
	for _, row := range result.Rows {
		if row != nil {
			profiles = append(profiles, row.AuthorUser, row.LastReplyUser)
		}
	}
	if err = u.hydrateArticleProfiles(ctx, profiles); err != nil {
		return nil, err
	}
	return &SearchArticlesResp{Rows: result.Rows, NextCursor: result.NextCursor}, nil
}

func (u *ContentArticleUsecase) ensureUnblocked(ctx context.Context, userID int64, targetID int64) error {
	if userID <= 0 || targetID <= 0 || userID == targetID || u.relationClient == nil {
		return nil
	}
	statuses, err := u.relationClient.MapStatus(ctx, userID, []int64{targetID})
	if err != nil {
		return err
	}
	if status := statuses[targetID]; status != nil && (status.Blocking || status.BlockedBy) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND)
	}
	return nil
}

func (u *ContentArticleUsecase) allowArticleSearch(ctx context.Context, req *SearchArticlesReq) error {
	if u.articleSearchRateLimitCache == nil {
		return nil
	}
	principal := "anonymous"
	if req.UserID > 0 {
		principal = fmt.Sprintf("user:%d", req.UserID)
	} else if ip := strings.TrimSpace(req.ClientIP); ip != "" {
		principal = "ip:" + ip
	}
	limit := u.articleSearchLimit
	if len([]rune(strings.TrimSpace(req.Query))) == 2 {
		limit = u.articleSearchShortQueryLimit
	}
	allowed, err := u.articleSearchRateLimitCache.Allow(ctx, &repo.ArticleSearchRateLimitReq{
		Key:    principal,
		Window: time.Minute,
		Limit:  limit,
	})
	if err != nil {
		return err
	}
	if !allowed {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_TOO_MANY_ReqS)
	}
	return nil
}

type GetArticleReq struct {
	UserID        int64
	ArticleID     int64
	PublishStatus *int32
}

func (u *ContentArticleUsecase) GetArticle(ctx context.Context, req *GetArticleReq) (*model.ArticleDetail, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_INVALID)
	}
	resp, err := u.contentArticleClient.GetArticle(ctx, &repo.GetArticleReq{UserID: req.UserID, ArticleID: req.ArticleID})
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND)
	}
	// A route declares the lifecycle state it serves. Do not let an author's
	// draft become readable from a public article URL merely because the caller
	// also owns that draft.
	if req.PublishStatus != nil && resp.PublishStatus != *req.PublishStatus {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND)
	}
	if resp.CreatedBy == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND)
	}
	if err = u.ensureUnblocked(ctx, req.UserID, *resp.CreatedBy); err != nil {
		return nil, err
	}
	if req.UserID != *resp.CreatedBy {
		privacy, err := u.privacyClient.GetCurrentPrivacySetting(ctx, *resp.CreatedBy)
		if err != nil {
			return nil, err
		}
		if privacy != nil && privacy.PublicLocation != nil && !*privacy.PublicLocation {
			resp.City = nil
		}
	}
	if err = u.hydrateArticleProfiles(ctx, []*model.AccountProfile{resp.AuthorUser, resp.LastReplyUser}); err != nil {
		return nil, err
	}
	return resp, nil
}

// hydrateArticleProfiles resolves display avatars only for article read models.
// Draft writes intentionally bypass it because their response exposes no profile.
func (u *ContentArticleUsecase) hydrateArticleProfiles(ctx context.Context, profiles []*model.AccountProfile) error {
	assetIDs := make([]int64, 0, len(profiles))
	seen := map[int64]struct{}{}
	for _, profile := range profiles {
		if profile == nil || profile.AvatarAssetID == nil || *profile.AvatarAssetID <= 0 {
			continue
		}
		if _, ok := seen[*profile.AvatarAssetID]; ok {
			continue
		}
		seen[*profile.AvatarAssetID] = struct{}{}
		assetIDs = append(assetIDs, *profile.AvatarAssetID)
	}
	assets := map[int64]*repo.Asset{}
	if len(assetIDs) > 0 && u.assetClient != nil {
		var err error
		assets, err = u.assetClient.Map(ctx, &repo.AssetGetReq{IDs: assetIDs})
		if err != nil {
			return err
		}
	}
	for _, profile := range profiles {
		if profile == nil {
			continue
		}
		avatarURL := "/v1/user/account/avatar?name=" + profile.Name
		if profile.AvatarAssetID != nil {
			if asset := assets[*profile.AvatarAssetID]; asset != nil && asset.URL != "" {
				avatarURL = asset.URL
			}
		}
		profile.AvatarURL = &avatarURL
	}
	return nil
}

type ViewArticleReq struct {
	UserID    int64
	ArticleID int64
	IP        *string
	UserAgent *string
}

func (u *ContentArticleUsecase) ViewArticle(ctx context.Context, req *ViewArticleReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_INVALID)
	}
	if err := u.ensureArticleUnblocked(ctx, req.UserID, req.ArticleID); err != nil {
		return err
	}
	return u.contentArticleClient.ViewArticle(ctx, &repo.ViewArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, IP: req.IP, UserAgent: req.UserAgent})
}

type LikeArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}
type ThankArticleReq struct {
	UserID          int64
	ArticleID       int64
}
type CollectArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}

type WatchArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}
type RewardArticleReq struct {
	UserID          int64
	ArticleID       int64
}

func (u *ContentArticleUsecase) LikeArticle(ctx context.Context, req *LikeArticleReq) (bool, error) {
	if err := u.ensureArticleUnblocked(ctx, req.UserID, req.ArticleID); err != nil {
		return false, err
	}
	return u.contentArticleClient.LikeArticle(ctx, &repo.LikeArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Active: req.Active})
}

func (u *ContentArticleUsecase) ThankArticle(ctx context.Context, req *ThankArticleReq) (bool, error) {
	if err := u.ensureArticleUnblocked(ctx, req.UserID, req.ArticleID); err != nil {
		return false, err
	}
	return u.contentArticleClient.ThankArticle(ctx, &repo.ThankArticleReq{UserID: req.UserID, ArticleID: req.ArticleID})
}

func (u *ContentArticleUsecase) CollectArticle(ctx context.Context, req *CollectArticleReq) (bool, error) {
	if err := u.ensureArticleUnblocked(ctx, req.UserID, req.ArticleID); err != nil {
		return false, err
	}
	return u.contentArticleClient.CollectArticle(ctx, &repo.CollectArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Active: req.Active})
}

func (u *ContentArticleUsecase) WatchArticle(ctx context.Context, req *WatchArticleReq) (bool, error) {
	if err := u.ensureArticleUnblocked(ctx, req.UserID, req.ArticleID); err != nil {
		return false, err
	}
	return u.contentArticleClient.WatchArticle(ctx, &repo.WatchArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Active: req.Active})
}

func (u *ContentArticleUsecase) RewardArticle(ctx context.Context, req *RewardArticleReq) (bool, error) {
	if err := u.ensureArticleUnblocked(ctx, req.UserID, req.ArticleID); err != nil {
		return false, err
	}
	return u.contentArticleClient.RewardArticle(ctx, &repo.RewardArticleReq{UserID: req.UserID, ArticleID: req.ArticleID})
}

func (u *ContentArticleUsecase) ensureArticleUnblocked(ctx context.Context, userID int64, articleID int64) error {
	if userID <= 0 || articleID <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND)
	}
	article, err := u.contentArticleClient.GetArticle(ctx, &repo.GetArticleReq{UserID: userID, ArticleID: articleID})
	if err != nil {
		return err
	}
	if article == nil || article.CreatedBy == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND)
	}
	return u.ensureUnblocked(ctx, userID, *article.CreatedBy)
}
