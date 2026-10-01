package usecase

import (
	"bff_bbs/internal/biz/repo"
	"common/pkg/apperror"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
	"strings"
	"time"
)

type ContentArticleUsecase struct {
	contentArticleClient repo.ContentArticleClient
	assetClient          repo.AssetClient
	privacyClient        repo.PrivacySettingClient
	ipResolutionClient   repo.IPResolutionClient
}

func NewContentArticleUsecase(
	contentArticleClient repo.ContentArticleClient,
	assetClient repo.AssetClient,
	privacyClient repo.PrivacySettingClient,
	ipResolutionClient repo.IPResolutionClient,
) *ContentArticleUsecase {
	return &ContentArticleUsecase{
		contentArticleClient: contentArticleClient,
		assetClient:          assetClient,
		privacyClient:        privacyClient,
		ipResolutionClient:   ipResolutionClient,
	}
}

type ContentArticleSave struct {
	Title         string
	Content       string
	RewardContent *string
	RewardPoints  *int32
	Type          int32
	Statement     *string
	Commentable   *bool
}

type CreateDraftArticleReq struct {
	UserID  int64
	Article *ContentArticleSave
}

func (u *ContentArticleUsecase) CreateDraftArticle(ctx context.Context, req *CreateDraftArticleReq) (int64, error) {
	if req == nil || req.Article == nil {
		return 0, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_INVALID)
	}
	return u.contentArticleClient.CreateDraftArticle(ctx, &repo.CreateDraftArticleReq{UserID: req.UserID, Article: &repo.ArticleSave{Title: req.Article.Title, Content: req.Article.Content, RewardContent: req.Article.RewardContent, RewardPoints: req.Article.RewardPoints, Type: req.Article.Type, Statement: req.Article.Statement, Commentable: req.Article.Commentable}})
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
	return u.contentArticleClient.UpdateDraftArticle(ctx, &repo.UpdateDraftArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Article: &repo.ArticleSave{Title: req.Article.Title, Content: req.Article.Content, RewardContent: req.Article.RewardContent, RewardPoints: req.Article.RewardPoints, Type: req.Article.Type, Statement: req.Article.Statement, Commentable: req.Article.Commentable}})
}

type PublishArticleReq struct {
	UserID      int64
	ArticleID   int64
	ScheduledAt *time.Time
	IP          string
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

type ListArticlesReq struct {
	UserID int64
	Page   *common.PageReq
	Query  *ArticleQuery
}

type ArticleQuery struct {
	TagID           *int64
	DomainID        *int64
	Keyword         *string
	AuthorID        *int64
	Type            *int32
	Order           *int32
	PublishStatus   *int32
	PublishStatuses []int32
	Visibility      *int32
	Visibilities    []int32
	Scheduled       *bool
}

type ListArticlesResp struct {
	Page *repo.PageResp
	Rows []*repo.ArticleListItem
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
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()}
	}
	var query *repo.ArticleQuery
	if req.Query != nil {
		query = &repo.ArticleQuery{TagID: req.Query.TagID, DomainID: req.Query.DomainID, Keyword: req.Query.Keyword, AuthorID: req.Query.AuthorID, Type: req.Query.Type, Order: req.Query.Order, PublishStatus: req.Query.PublishStatus, PublishStatuses: req.Query.PublishStatuses, Visibility: req.Query.Visibility, Visibilities: req.Query.Visibilities, Scheduled: req.Query.Scheduled}
	}
	resp, err := u.contentArticleClient.ListArticles(ctx, &repo.ListArticlesReq{UserID: req.UserID, Page: page, Query: query})
	if err != nil {
		return nil, err
	}
	profiles := make([]*repo.AccountProfile, 0, len(resp.Rows)*2)
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

type GetArticleReq struct {
	UserID        int64
	ArticleID     int64
	PublishStatus *int32
}

func (u *ContentArticleUsecase) GetArticle(ctx context.Context, req *GetArticleReq) (*repo.ArticleDetail, error) {
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
	if req.UserID != *resp.CreatedBy {
		privacy, err := u.privacyClient.GetCurrentPrivacySetting(ctx, *resp.CreatedBy)
		if err != nil {
			return nil, err
		}
		if privacy != nil && privacy.PublicLocation != nil && !*privacy.PublicLocation {
			resp.City = nil
		}
	}
	if err = u.hydrateArticleProfiles(ctx, []*repo.AccountProfile{resp.AuthorUser, resp.LastReplyUser}); err != nil {
		return nil, err
	}
	return resp, nil
}

// hydrateArticleProfiles resolves display avatars only for article read models.
// Draft writes intentionally bypass it because their response exposes no profile.
func (u *ContentArticleUsecase) hydrateArticleProfiles(ctx context.Context, profiles []*repo.AccountProfile) error {
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
	return u.contentArticleClient.ViewArticle(ctx, &repo.ViewArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, IP: req.IP, UserAgent: req.UserAgent})
}

type LikeArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}
type ThankArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}
type CollectArticleReq struct {
	UserID    int64
	ArticleID int64
	Active    bool
}
type RewardArticleReq struct {
	UserID    int64
	ArticleID int64
	Points    int32
}

func (u *ContentArticleUsecase) LikeArticle(ctx context.Context, req *LikeArticleReq) (bool, error) {
	return u.contentArticleClient.LikeArticle(ctx, &repo.LikeArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Active: req.Active})
}

func (u *ContentArticleUsecase) ThankArticle(ctx context.Context, req *ThankArticleReq) (bool, error) {
	return u.contentArticleClient.ThankArticle(ctx, &repo.ThankArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Active: req.Active})
}

func (u *ContentArticleUsecase) CollectArticle(ctx context.Context, req *CollectArticleReq) (bool, error) {
	return u.contentArticleClient.CollectArticle(ctx, &repo.CollectArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Active: req.Active})
}

func (u *ContentArticleUsecase) RewardArticle(ctx context.Context, req *RewardArticleReq) error {
	return u.contentArticleClient.RewardArticle(ctx, &repo.RewardArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Points: req.Points})
}
