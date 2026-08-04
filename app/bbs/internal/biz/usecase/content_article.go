package usecase

import (
	"bbs/internal/biz/repo"
	"common/pkg/apperror"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"
	"time"
)

type ContentArticleUsecase struct {
	contentArticleClient repo.ContentArticleClient
	assetClient          repo.AssetClient
}

func NewContentArticleUsecase(
	contentArticleClient repo.ContentArticleClient,
	assetClient repo.AssetClient,
) *ContentArticleUsecase {
	return &ContentArticleUsecase{
		contentArticleClient: contentArticleClient,
		assetClient:          assetClient,
	}
}

type ContentArticleViewerActionState = repo.ArticleViewerActionState
type ContentArticlePostscript = repo.ArticlePostscript
type ContentArticleListItem = repo.ArticleListItem
type ContentArticleDetail = repo.ArticleDetail
type ContentAccountProfile = repo.AccountProfile
type ContentPageResp = repo.PageResp

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

func (u *ContentArticleUsecase) CreateDraftArticle(ctx context.Context, req *CreateDraftArticleReq) (*ContentArticleDetail, error) {
	if req == nil || req.Article == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := u.contentArticleClient.CreateDraftArticle(ctx, &repo.CreateDraftArticleReq{UserID: req.UserID, Article: &repo.ArticleSave{Title: req.Article.Title, Content: req.Article.Content, RewardContent: req.Article.RewardContent, RewardPoints: req.Article.RewardPoints, Type: req.Article.Type, Statement: req.Article.Statement, Commentable: req.Article.Commentable}})
	if err != nil {
		return nil, err
	}
	if resp != nil {
		profiles := []*repo.AccountProfile{resp.AuthorUser, resp.LastReplyUser}
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
				return nil, err
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
	}
	return resp, nil
}

type UpdateDraftArticleReq struct {
	UserID    int64
	ArticleID int64
	Article   *ContentArticleSave
}

func (u *ContentArticleUsecase) UpdateDraftArticle(ctx context.Context, req *UpdateDraftArticleReq) (*ContentArticleDetail, error) {
	if req == nil || req.Article == nil || req.ArticleID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := u.contentArticleClient.UpdateDraftArticle(ctx, &repo.UpdateDraftArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Article: &repo.ArticleSave{Title: req.Article.Title, Content: req.Article.Content, RewardContent: req.Article.RewardContent, RewardPoints: req.Article.RewardPoints, Type: req.Article.Type, Statement: req.Article.Statement, Commentable: req.Article.Commentable}})
	if err != nil {
		return nil, err
	}
	if resp != nil {
		profiles := []*repo.AccountProfile{resp.AuthorUser, resp.LastReplyUser}
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
				return nil, err
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
	}
	return resp, nil
}

type PublishArticleReq struct {
	UserID     int64
	ArticleID  int64
	Visibility int32
}

func (u *ContentArticleUsecase) PublishArticle(ctx context.Context, req *PublishArticleReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.contentArticleClient.PublishArticle(ctx, &repo.PublishArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, Visibility: req.Visibility})
}

type SchedulePublishArticleReq struct {
	UserID    int64
	ArticleID int64
	PublishAt time.Time
}

func (u *ContentArticleUsecase) SchedulePublishArticle(ctx context.Context, req *SchedulePublishArticleReq) error {
	if req == nil || req.PublishAt.IsZero() {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.contentArticleClient.SchedulePublishArticle(ctx, &repo.SchedulePublishArticleReq{UserID: req.UserID, ArticleID: req.ArticleID, PublishAt: req.PublishAt})
}

type CancelPublishArticleReq struct {
	UserID    int64
	ArticleID int64
}

func (u *ContentArticleUsecase) CancelPublishArticle(ctx context.Context, req *CancelPublishArticleReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.contentArticleClient.CancelPublishArticle(ctx, &repo.CancelPublishArticleReq{UserID: req.UserID, ArticleID: req.ArticleID})
}

type DiscardDraftArticleReq struct {
	UserID    int64
	ArticleID int64
}

func (u *ContentArticleUsecase) DiscardDraftArticle(ctx context.Context, req *DiscardDraftArticleReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
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
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
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
}

type ListArticlesResp struct {
	Page *ContentPageResp
	Rows []*ContentArticleListItem
}

func (u *ContentArticleUsecase) ListArticles(ctx context.Context, req *ListArticlesReq) (*ListArticlesResp, error) {
	if req == nil {
		req = &ListArticlesReq{}
	}
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{Page: req.Page.GetPage(), Size: req.Page.GetSize()}
	}
	var query *repo.ArticleQuery
	if req.Query != nil {
		query = &repo.ArticleQuery{TagID: req.Query.TagID, DomainID: req.Query.DomainID, Keyword: req.Query.Keyword, AuthorID: req.Query.AuthorID, Type: req.Query.Type, Order: req.Query.Order, PublishStatus: req.Query.PublishStatus, PublishStatuses: req.Query.PublishStatuses, Visibility: req.Query.Visibility, Visibilities: req.Query.Visibilities}
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
		assets, err = u.assetClient.Map(ctx, &repo.AssetGetReq{IDs: assetIDs})
		if err != nil {
			return nil, err
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
	return &ListArticlesResp{Page: resp.Page, Rows: resp.Rows}, nil
}

type GetArticleReq struct {
	UserID    int64
	ArticleID int64
}

func (u *ContentArticleUsecase) GetArticle(ctx context.Context, req *GetArticleReq) (*ContentArticleDetail, error) {
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	resp, err := u.contentArticleClient.GetArticle(ctx, &repo.GetArticleReq{UserID: req.UserID, ArticleID: req.ArticleID})
	if err != nil {
		return nil, err
	}
	if resp != nil {
		profiles := []*repo.AccountProfile{resp.AuthorUser, resp.LastReplyUser}
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
				return nil, err
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
	}
	return resp, nil
}

type ViewArticleReq struct {
	UserID    int64
	ArticleID int64
	IP        *string
	UserAgent *string
}

func (u *ContentArticleUsecase) ViewArticle(ctx context.Context, req *ViewArticleReq) error {
	if req == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
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
