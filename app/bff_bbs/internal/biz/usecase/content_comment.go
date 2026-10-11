package usecase

import (
	"common/pkg/apperror"
	commonmodel "common/pkg/model"
	bbscontentv1 "common/proto/gen/bff_bbs/v1/content"
	bbscontentv1enum "common/proto/gen/bff_bbs/v1/content/enum"
	contentmodel "common/proto/gen/bff_bbs/v1/content/model"
	bbsuserv1 "common/proto/gen/bff_bbs/v1/user"
	bbsuserv1enum "common/proto/gen/bff_bbs/v1/user/enum"
	cerrors "common/proto/gen/common/errors"
	"context"

	"bff_bbs/internal/biz/model"
	"bff_bbs/internal/biz/repo"
	"bff_bbs/internal/enum"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type ContentCommentUsecase struct {
	contentCommentClient repo.ContentCommentClient
	contentArticleClient repo.ContentArticleClient
	assetClient          repo.AssetClient
	relationClient       repo.RelationClient
	privacyClient        repo.PrivacySettingClient
}

func NewContentCommentUsecase(
	contentCommentClient repo.ContentCommentClient,
	contentArticleClient repo.ContentArticleClient,
	assetClient repo.AssetClient,
	relationClient repo.RelationClient,
	privacyClient repo.PrivacySettingClient,
) *ContentCommentUsecase {
	return &ContentCommentUsecase{
		contentCommentClient: contentCommentClient,
		contentArticleClient: contentArticleClient,
		assetClient:          assetClient,
		relationClient:       relationClient,
		privacyClient:        privacyClient,
	}
}

type CreateCommentReq struct {
	UserID    int64
	ArticleID int64
	Content   string
	ReplyID   int64
}

func (u *ContentCommentUsecase) CreateComment(ctx context.Context, req *CreateCommentReq) (*repo.CommentDetail, error) {
	if err := u.ensureArticleUnblocked(ctx, req.UserID, req.ArticleID); err != nil {
		return nil, err
	}
	if err := u.ensureCommentTargetUnblocked(ctx, req.UserID, req.ReplyID); err != nil {
		return nil, err
	}
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
		profiles := []*model.AccountProfile{resp.User, resp.ReplyUser}
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

func (u *ContentCommentUsecase) ensureArticleUnblocked(ctx context.Context, userID int64, articleID int64) error {
	if userID <= 0 || articleID <= 0 || u.relationClient == nil || u.contentArticleClient == nil {
		return nil
	}
	article, err := u.contentArticleClient.GetArticle(ctx, &repo.GetArticleReq{UserID: userID, ArticleID: articleID})
	if err != nil {
		return err
	}
	if article == nil || article.CreatedBy == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND)
	}
	if *article.CreatedBy == userID {
		return nil
	}
	statuses, err := u.relationClient.MapStatus(ctx, userID, []int64{*article.CreatedBy})
	if err != nil {
		return err
	}
	if status := statuses[*article.CreatedBy]; status != nil && (status.Blocking || status.BlockedBy) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_ARTICLE_NOT_FOUND)
	}
	return nil
}

type ListCommentsReq struct {
	UserID int64
	Page   *commonmodel.PageReq
	Query  *bbscontentv1.ListComments_Req_CommentQuery
}

type ListCommentsResp struct {
	Page *commonmodel.PageResp
	Rows []*model.CommentListItem
}

func (u *ContentCommentUsecase) ListComments(ctx context.Context, req *ListCommentsReq) (*ListCommentsResp, error) {
	if req == nil {
		req = &ListCommentsReq{}
	}
	var page *commonmodel.PageReq
	if req.Page != nil {
		page = &commonmodel.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
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
	if query.UserID != nil && req.UserID != *query.UserID {
		privacy, err := u.privacyClient.GetCurrentPrivacySetting(ctx, *query.UserID)
		if err != nil {
			return nil, err
		}
		if privacy.PublicCommentList != nil && !*privacy.PublicCommentList {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_PROFILE_COMMENT_LIST_PRIVATE)
		}
	}
	excludedUserIDs := []int64{}
	if req.UserID > 0 && u.relationClient != nil {
		var err error
		excludedUserIDs, err = u.relationClient.ListBlockedCounterparts(ctx, req.UserID)
		if err != nil {
			return nil, err
		}
	}
	query.ExcludedUserIDs = excludedUserIDs
	resp, err := u.contentCommentClient.ListComments(ctx, &repo.ListCommentsReq{
		UserID: req.UserID,
		Page:   page,
		Query:  query,
	})
	if err != nil {
		return nil, err
	}
	profiles := make([]*model.AccountProfile, 0, len(resp.Rows)*2)
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
	Page              *commonmodel.PageReq
	ArticleID         int64
	Order             *bbscontentv1enum.CommentOrder
	ReplyPreviewLimit *int32
}

type ListCommentThreadsResp struct {
	Page *commonmodel.PageResp
	Rows []*repo.CommentThread
}

func (u *ContentCommentUsecase) ListCommentThreads(ctx context.Context, req *ListCommentThreadsReq) (*ListCommentThreadsResp, error) {
	excludedUserIDs := []int64{}
	if req.UserID > 0 && u.relationClient != nil {
		var err error
		excludedUserIDs, err = u.relationClient.ListBlockedCounterparts(ctx, req.UserID)
		if err != nil {
			return nil, err
		}
	}
	var page *commonmodel.PageReq
	if req.Page != nil {
		page = &commonmodel.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
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
		ExcludedUserIDs:   excludedUserIDs,
	})
	if err != nil {
		return nil, err
	}
	profiles := make([]*model.AccountProfile, 0, len(resp.Rows)*4)
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
	Page      *commonmodel.PageReq
	ArticleID int64
	ParentID  int64
	Order     *bbscontentv1enum.CommentOrder
}

type ListCommentRepliesResp struct {
	Page *commonmodel.PageResp
	Rows []*model.CommentListItem
}

func (u *ContentCommentUsecase) ListCommentReplies(ctx context.Context, req *ListCommentRepliesReq) (*ListCommentRepliesResp, error) {
	excludedUserIDs := []int64{}
	if req.UserID > 0 && u.relationClient != nil {
		var err error
		excludedUserIDs, err = u.relationClient.ListBlockedCounterparts(ctx, req.UserID)
		if err != nil {
			return nil, err
		}
	}
	var page *commonmodel.PageReq
	if req.Page != nil {
		page = &commonmodel.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	var order *int32
	if req.Order != nil {
		order = new(int32(*req.Order))
	}
	resp, err := u.contentCommentClient.ListCommentReplies(ctx, &repo.ListCommentRepliesReq{
		UserID:          req.UserID,
		Page:            page,
		ArticleID:       req.ArticleID,
		ParentID:        req.ParentID,
		Order:           order,
		ExcludedUserIDs: excludedUserIDs,
	})
	if err != nil {
		return nil, err
	}
	profiles := make([]*model.AccountProfile, 0, len(resp.Rows)*2)
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
	Page      *commonmodel.PageReq
	ArticleID int64
	Order     *bbscontentv1enum.CommentOrder
}

type ListCommentTimelineResp struct {
	Page *commonmodel.PageResp
	Rows []*model.CommentListItem
}

func (u *ContentCommentUsecase) ListCommentTimeline(ctx context.Context, req *ListCommentTimelineReq) (*ListCommentTimelineResp, error) {
	excludedUserIDs := []int64{}
	if req.UserID > 0 && u.relationClient != nil {
		var err error
		excludedUserIDs, err = u.relationClient.ListBlockedCounterparts(ctx, req.UserID)
		if err != nil {
			return nil, err
		}
	}
	var page *commonmodel.PageReq
	if req.Page != nil {
		page = &commonmodel.PageReq{
			Page: req.Page.Page,
			Size: req.Page.Size,
		}
	}
	var order *int32
	if req.Order != nil {
		order = new(int32(*req.Order))
	}
	resp, err := u.contentCommentClient.ListCommentTimeline(ctx, &repo.ListCommentTimelineReq{
		UserID:          req.UserID,
		Page:            page,
		ArticleID:       req.ArticleID,
		Order:           order,
		ExcludedUserIDs: excludedUserIDs,
	})
	if err != nil {
		return nil, err
	}
	profiles := make([]*model.AccountProfile, 0, len(resp.Rows)*2)
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

func (u *ContentCommentUsecase) ListCommentRepliesProto(resp *ListCommentRepliesResp) *bbscontentv1.ListCommentReplies_Resp {
	if resp == nil {
		return &bbscontentv1.ListCommentReplies_Resp{}
	}
	rows := make([]*contentmodel.CommentListItem, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		rows = append(rows, u.commentListItemProto(row))
	}
	return &bbscontentv1.ListCommentReplies_Resp{Page: resp.Page.ToProto(), Rows: rows}
}

func (u *ContentCommentUsecase) ListCommentTimelineProto(resp *ListCommentTimelineResp) *bbscontentv1.ListCommentTimeline_Resp {
	if resp == nil {
		return &bbscontentv1.ListCommentTimeline_Resp{}
	}
	rows := make([]*contentmodel.CommentListItem, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		rows = append(rows, u.commentListItemProto(row))
	}
	return &bbscontentv1.ListCommentTimeline_Resp{Page: resp.Page.ToProto(), Rows: rows}
}

func (u *ContentCommentUsecase) commentListItemProto(row *model.CommentListItem) *contentmodel.CommentListItem {
	if row == nil {
		return nil
	}
	item := &contentmodel.CommentListItem{Id: row.ID, ArticleId: row.ArticleID, Content: row.Content, ContentRender: row.ContentRender, Level: row.Level, ParentId: row.ParentID, ReplyId: row.ReplyID, ReplyCount: row.ReplyCount, LikeCount: row.LikeCount, ThankCount: row.ThankCount, Restriction: contentmodel.ContentRestriction(enum.ContentRestrictionMap.MustToProto(row.Restriction)), CreatedBy: row.CreatedBy, UpdatedBy: row.UpdatedBy, CreatedAt: timestamppb.New(*row.CreatedAt), UpdatedAt: timestamppb.New(*row.UpdatedAt)}
	if row.ViewerActionState != nil {
		item.ViewerActionState = &contentmodel.CommentViewerActionState{Liked: row.ViewerActionState.Liked, Thanked: row.ViewerActionState.Thanked}
	}
	if account := row.User; account != nil {
		item.User = &bbsuserv1.AccountProfile{Id: account.ID, Name: account.Name, Nickname: account.Nickname, Url: account.URL, AvatarUrl: account.AvatarURL, Introduction: account.Introduction, Status: bbsuserv1enum.AccountStatus(account.Status), Mbti: account.MBTI, FollowCount: account.FollowCount, FollowerCount: account.FollowerCount, CreatedAt: timestamppb.New(*account.CreatedAt), UpdatedAt: timestamppb.New(*account.UpdatedAt)}
	}
	if account := row.ReplyUser; account != nil {
		item.ReplyUser = &bbsuserv1.AccountProfile{Id: account.ID, Name: account.Name, Nickname: account.Nickname, Url: account.URL, AvatarUrl: account.AvatarURL, Introduction: account.Introduction, Status: bbsuserv1enum.AccountStatus(account.Status), Mbti: account.MBTI, FollowCount: account.FollowCount, FollowerCount: account.FollowerCount, CreatedAt: timestamppb.New(*account.CreatedAt), UpdatedAt: timestamppb.New(*account.UpdatedAt)}
	}
	return item
}

type LikeCommentReq struct {
	UserID int64
	ID     int64
	Active bool
}

func (u *ContentCommentUsecase) LikeComment(ctx context.Context, req *LikeCommentReq) (bool, error) {
	if err := u.ensureCommentTargetUnblocked(ctx, req.UserID, req.ID); err != nil {
		return false, err
	}
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
	if err := u.ensureCommentTargetUnblocked(ctx, req.UserID, req.ID); err != nil {
		return false, err
	}
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

func (u *ContentCommentUsecase) ensureCommentTargetUnblocked(ctx context.Context, userID int64, commentID int64) error {
	if userID <= 0 || commentID <= 0 || u.relationClient == nil {
		return nil
	}
	commentIDValue := commentID
	result, err := u.contentCommentClient.ListComments(ctx, &repo.ListCommentsReq{
		UserID: userID,
		Page:   &commonmodel.PageReq{Page: 1, Size: 1},
		Query:  &repo.CommentQuery{CommentID: &commentIDValue},
	})
	if err != nil {
		return err
	}
	if result == nil || len(result.Rows) == 0 || result.Rows[0] == nil || result.Rows[0].User == nil {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_NOT_FOUND)
	}
	if result.Rows[0].User.ID == userID {
		return nil
	}
	targetID := result.Rows[0].User.ID
	statuses, err := u.relationClient.MapStatus(ctx, userID, []int64{targetID})
	if err != nil {
		return err
	}
	if status := statuses[targetID]; status != nil && (status.Blocking || status.BlockedBy) {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_COMMENT_NOT_FOUND)
	}
	return nil
}
