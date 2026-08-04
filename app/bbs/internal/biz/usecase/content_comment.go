package usecase

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	bbscontentv1enum "common/proto/gen/bbs/v1/content/enum"
	"common/proto/gen/common"
	"context"
)

type ContentCommentUsecase struct {
	contentCommentClient repo.ContentCommentClient
	assetClient          repo.AssetClient
}

func NewContentCommentUsecase(
	contentCommentClient repo.ContentCommentClient,
	assetClient repo.AssetClient,
) *ContentCommentUsecase {
	return &ContentCommentUsecase{
		contentCommentClient: contentCommentClient,
		assetClient:          assetClient,
	}
}

type CreateCommentReq struct {
	UserID    int64
	ArticleID int64
	Content   string
	ReplyID   int64
}

func (u *ContentCommentUsecase) CreateComment(ctx context.Context, req *CreateCommentReq) (*repo.CommentDetail, error) {
	resp, err := u.contentCommentClient.CreateComment(ctx, &repo.CreateCommentReq{
		UserID:    req.UserID,
		ArticleID: req.ArticleID,
		Content:   req.Content,
		ReplyID:   req.ReplyID,
	})
	if err != nil {
		return nil, err
	}
	if resp != nil {
		profiles := []*repo.AccountProfile{resp.User, resp.ReplyUser}
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
	}
	return resp, nil
}

type ListCommentsReq struct {
	UserID int64
	Page   *common.PageReq
	Query  *bbscontentv1.ListComments_Req_CommentQuery
}

type ListCommentsResp struct {
	Page *repo.PageResp
	Rows []*repo.CommentListItem
}

func (u *ContentCommentUsecase) ListComments(ctx context.Context, req *ListCommentsReq) (*ListCommentsResp, error) {
	if req == nil {
		req = &ListCommentsReq{}
	}
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{
			Page: req.Page.GetPage(),
			Size: req.Page.GetSize(),
		}
	}
	query := &repo.CommentQuery{
		Restrictions: []int32{
			int32(bbscontentv1enum.ContentRestriction_CONTENT_RESTRICTION_NONE),
			int32(bbscontentv1enum.ContentRestriction_CONTENT_RESTRICTION_LOCKED),
		},
	}
	if req.Query != nil {
		query.CommentID = req.Query.CommentId
		query.ArticleID = req.Query.ArticleId
		query.ParentID = req.Query.ParentId
		query.ReplyID = req.Query.ReplyId
		query.Level = req.Query.Level
		query.UserID = req.Query.UserId
		if req.Query.Order != nil {
			query.Order = new(int32(*req.Query.Order))
		}
	}
	resp, err := u.contentCommentClient.ListComments(ctx, &repo.ListCommentsReq{
		UserID: req.UserID,
		Page:   page,
		Query:  query,
	})
	if err != nil {
		return nil, err
	}
	profiles := make([]*repo.AccountProfile, 0, len(resp.Rows)*2)
	for _, row := range resp.Rows {
		if row == nil {
			continue
		}
		profiles = append(profiles, row.User, row.ReplyUser)
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
	return &ListCommentsResp{
		Page: resp.Page,
		Rows: resp.Rows,
	}, nil
}

type ListCommentThreadsReq struct {
	UserID            int64
	Page              *common.PageReq
	ArticleID         int64
	Order             *bbscontentv1enum.CommentOrder
	ReplyPreviewLimit *int32
}

type ListCommentThreadsResp struct {
	Page *repo.PageResp
	Rows []*repo.CommentThread
}

func (u *ContentCommentUsecase) ListCommentThreads(ctx context.Context, req *ListCommentThreadsReq) (*ListCommentThreadsResp, error) {
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{
			Page: req.Page.GetPage(),
			Size: req.Page.GetSize(),
		}
	}
	var order *int32
	if req.Order != nil {
		order = new(int32(*req.Order))
	}
	resp, err := u.contentCommentClient.ListCommentThreads(ctx, &repo.ListCommentThreadsReq{
		UserID:            req.UserID,
		Page:              page,
		ArticleID:         req.ArticleID,
		Order:             order,
		ReplyPreviewLimit: req.ReplyPreviewLimit,
	})
	if err != nil {
		return nil, err
	}
	profiles := make([]*repo.AccountProfile, 0, len(resp.Rows)*4)
	for _, row := range resp.Rows {
		if row == nil {
			continue
		}
		if row.Root != nil {
			profiles = append(profiles, row.Root.User, row.Root.ReplyUser)
		}
		for _, reply := range row.PreviewReplies {
			if reply == nil {
				continue
			}
			profiles = append(profiles, reply.User, reply.ReplyUser)
		}
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
	return &ListCommentThreadsResp{
		Page: resp.Page,
		Rows: resp.Rows,
	}, nil
}

type ListCommentRepliesReq struct {
	UserID    int64
	Page      *common.PageReq
	ArticleID int64
	ParentID  int64
	Order     *bbscontentv1enum.CommentOrder
}

type ListCommentRepliesResp struct {
	Page *repo.PageResp
	Rows []*repo.CommentListItem
}

func (u *ContentCommentUsecase) ListCommentReplies(ctx context.Context, req *ListCommentRepliesReq) (*ListCommentRepliesResp, error) {
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{
			Page: req.Page.GetPage(),
			Size: req.Page.GetSize(),
		}
	}
	var order *int32
	if req.Order != nil {
		order = new(int32(*req.Order))
	}
	resp, err := u.contentCommentClient.ListCommentReplies(ctx, &repo.ListCommentRepliesReq{
		UserID:    req.UserID,
		Page:      page,
		ArticleID: req.ArticleID,
		ParentID:  req.ParentID,
		Order:     order,
	})
	if err != nil {
		return nil, err
	}
	profiles := make([]*repo.AccountProfile, 0, len(resp.Rows)*2)
	for _, row := range resp.Rows {
		if row == nil {
			continue
		}
		profiles = append(profiles, row.User, row.ReplyUser)
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
	return &ListCommentRepliesResp{
		Page: resp.Page,
		Rows: resp.Rows,
	}, nil
}

type ListCommentTimelineReq struct {
	UserID    int64
	Page      *common.PageReq
	ArticleID int64
	Order     *bbscontentv1enum.CommentOrder
}

type ListCommentTimelineResp struct {
	Page *repo.PageResp
	Rows []*repo.CommentListItem
}

func (u *ContentCommentUsecase) ListCommentTimeline(ctx context.Context, req *ListCommentTimelineReq) (*ListCommentTimelineResp, error) {
	var page *repo.PageReq
	if req.Page != nil {
		page = &repo.PageReq{
			Page: req.Page.GetPage(),
			Size: req.Page.GetSize(),
		}
	}
	var order *int32
	if req.Order != nil {
		order = new(int32(*req.Order))
	}
	resp, err := u.contentCommentClient.ListCommentTimeline(ctx, &repo.ListCommentTimelineReq{
		UserID:    req.UserID,
		Page:      page,
		ArticleID: req.ArticleID,
		Order:     order,
	})
	if err != nil {
		return nil, err
	}
	profiles := make([]*repo.AccountProfile, 0, len(resp.Rows)*2)
	for _, row := range resp.Rows {
		if row == nil {
			continue
		}
		profiles = append(profiles, row.User, row.ReplyUser)
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
	return &ListCommentTimelineResp{
		Page: resp.Page,
		Rows: resp.Rows,
	}, nil
}

type LikeCommentReq struct {
	UserID int64
	ID     int64
	Active bool
}

func (u *ContentCommentUsecase) LikeComment(ctx context.Context, req *LikeCommentReq) (bool, error) {
	resp, err := u.contentCommentClient.LikeComment(ctx, &repo.LikeCommentReq{
		UserID: req.UserID,
		ID:     req.ID,
		Active: req.Active,
	})
	if err != nil {
		return false, err
	}
	return resp, nil
}

type ThankCommentReq struct {
	UserID int64
	ID     int64
	Active bool
}

func (u *ContentCommentUsecase) ThankComment(ctx context.Context, req *ThankCommentReq) (bool, error) {
	resp, err := u.contentCommentClient.ThankComment(ctx, &repo.ThankCommentReq{
		UserID: req.UserID,
		ID:     req.ID,
		Active: req.Active,
	})
	if err != nil {
		return false, err
	}
	return resp, nil
}
