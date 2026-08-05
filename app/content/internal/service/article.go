package service

import (
	"common/pkg/apperror"
	"common/pkg/util"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/content/v1"
	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/biz/usecase"
	"content/internal/enum"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticleService struct {
	v1.UnimplementedContentArticleServiceServer

	articleUsecase *usecase.ArticleUsecase
}

func NewArticleService(
	articleUsecase *usecase.ArticleUsecase,
) *ArticleService {
	return &ArticleService{
		articleUsecase: articleUsecase,
	}
}

func (s *ArticleService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentArticleServiceServer(gs, s)
}

func (s *ArticleService) RegisterHttp(hs *http.Server) {
}

func (s *ArticleService) CreateDraft(ctx context.Context, req *v1.CreateDraftArticle_Req) (*v1.CreateDraftArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	articleReq := req.GetArticle()
	if articleReq == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	articleType, ok := enum.ArticleTypeMap.ToEnum(articleReq.GetType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
	}
	row, err := s.articleUsecase.Add(ctx, &usecase.ArticleAddReq{
		Access: access,
		Article: &model.Article{
			Title:         articleReq.GetTitle(),
			Content:       articleReq.GetContent(),
			RewardContent: articleReq.RewardContent,
			RewardPoints:  articleReq.RewardPoints,
			Type:          articleType,
			Statement:     articleReq.Statement,
			Commentable:   util.DerefOrDefault(articleReq.Commentable, true),
		},
	})
	if err != nil {
		return nil, err
	}
	article := &v1.Article{
		Id:            row.ID,
		Title:         row.Title,
		Content:       row.Content,
		RewardContent: row.RewardContent,
		RewardPoints:  row.RewardPoints,
		HasPostscript: row.HasPostscript,
		HasReward:     row.RewardPoints != nil,
		Type:          enum.ArticleTypeMap.MustToProto(row.Type),
		Statement:     row.Statement,
		Commentable:   row.Commentable,
		PublishStatus: enum.ArticlePublishStatusMap.MustToProto(row.PublishStatus),
		Visibility:    enum.ArticleVisibilityMap.MustToProto(row.Visibility),
		Restriction:   enum.ContentRestrictionMap.MustToProto(row.Restriction),
		ViewCount:     row.ViewCount,
		ThankCount:    row.ThankCount,
		LikeCount:     row.LikeCount,
		CollectCount:  row.CollectCount,
		RewardCount:   row.RewardCount,
		ReplyCount:    row.ReplyCount,
		CreatedBy:     row.CreatedBy,
		UpdatedBy:     row.UpdatedBy,
	}
	if row.CreatedAt != nil {
		article.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		article.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	if row.PublishedAt != nil {
		article.PublishedAt = timestamppb.New(*row.PublishedAt)
	}
	if row.EditedAt != nil {
		article.EditedAt = timestamppb.New(*row.EditedAt)
	}
	return &v1.CreateDraftArticle_Resp{Article: article}, nil
}

func (s *ArticleService) UpdateDraft(ctx context.Context, req *v1.UpdateDraftArticle_Req) (*v1.UpdateDraftArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	articleReq := req.GetArticle()
	if articleReq == nil || req.GetArticleId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	articleType, ok := enum.ArticleTypeMap.ToEnum(articleReq.GetType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
	}
	row, err := s.articleUsecase.Update(ctx, &usecase.ArticleUpdateDraftReq{
		Access: access,
		Article: &model.Article{
			ID:            req.GetArticleId(),
			Title:         articleReq.GetTitle(),
			Content:       articleReq.GetContent(),
			RewardContent: articleReq.RewardContent,
			RewardPoints:  articleReq.RewardPoints,
			Type:          articleType,
			Statement:     articleReq.Statement,
			Commentable:   util.DerefOrDefault(articleReq.Commentable, true),
			UpdatedBy:     new(access.ActorUserID),
		},
	})
	if err != nil {
		return nil, err
	}
	article := &v1.Article{
		Id:            row.ID,
		Title:         row.Title,
		Content:       row.Content,
		RewardContent: row.RewardContent,
		RewardPoints:  row.RewardPoints,
		HasPostscript: row.HasPostscript,
		HasReward:     row.RewardPoints != nil,
		Type:          enum.ArticleTypeMap.MustToProto(row.Type),
		Statement:     row.Statement,
		Commentable:   row.Commentable,
		PublishStatus: enum.ArticlePublishStatusMap.MustToProto(row.PublishStatus),
		Visibility:    enum.ArticleVisibilityMap.MustToProto(row.Visibility),
		Restriction:   enum.ContentRestrictionMap.MustToProto(row.Restriction),
		ViewCount:     row.ViewCount,
		ThankCount:    row.ThankCount,
		LikeCount:     row.LikeCount,
		CollectCount:  row.CollectCount,
		RewardCount:   row.RewardCount,
		ReplyCount:    row.ReplyCount,
		CreatedBy:     row.CreatedBy,
		UpdatedBy:     row.UpdatedBy,
	}
	if row.CreatedAt != nil {
		article.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		article.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	if row.PublishedAt != nil {
		article.PublishedAt = timestamppb.New(*row.PublishedAt)
	}
	if row.EditedAt != nil {
		article.EditedAt = timestamppb.New(*row.EditedAt)
	}
	return &v1.UpdateDraftArticle_Resp{Article: article}, nil
}

func (s *ArticleService) Publish(ctx context.Context, req *v1.PublishArticle_Req) (*v1.PublishArticle_Resp, error) {
	access := &model.ContentAccess{Scope: enum.ContentAccessScopeInternalTask}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize(enum.ContentAccessScopeInternalTask)
	if err != nil {
		return nil, err
	}
	visibility := enum.ArticleVisibilityPublic
	if req.Visibility != 0 {
		item, ok := enum.ArticleVisibilityMap.ToEnum(req.Visibility)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		visibility = item
	}
	if err := s.articleUsecase.Publish(ctx, &usecase.ArticlePublishReq{Access: access, ArticleID: req.GetArticleId(), Visibility: visibility}); err != nil {
		return nil, err
	}
	return &v1.PublishArticle_Resp{}, nil
}

func (s *ArticleService) SchedulePublish(ctx context.Context, req *v1.SchedulePublishArticle_Req) (*v1.SchedulePublishArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	if req.GetPublishAt() == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.articleUsecase.SchedulePublish(ctx, &usecase.ArticleSchedulePublishReq{Access: access, ArticleID: req.GetArticleId(), PublishAt: req.GetPublishAt().AsTime()}); err != nil {
		return nil, err
	}
	return &v1.SchedulePublishArticle_Resp{}, nil
}

func (s *ArticleService) CancelPublish(ctx context.Context, req *v1.CancelPublishArticle_Req) (*v1.CancelPublishArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.articleUsecase.CancelPublish(ctx, &usecase.ArticleCancelPublishReq{Access: access, ArticleID: req.GetArticleId()}); err != nil {
		return nil, err
	}
	return &v1.CancelPublishArticle_Resp{}, nil
}

func (s *ArticleService) DiscardDraft(ctx context.Context, req *v1.DiscardDraftArticle_Req) (*v1.DiscardDraftArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.DiscardDraftArticle_Resp{}, s.articleUsecase.DiscardDraft(ctx, &usecase.ArticleDiscardDraftReq{Access: access, ArticleID: req.GetArticleId()})
}

func (s *ArticleService) MakePrivate(ctx context.Context, req *v1.MakePrivateArticle_Req) (*v1.MakePrivateArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.MakePrivateArticle_Resp{}, s.articleUsecase.MakePrivate(ctx, &usecase.ArticleMakePrivateReq{Access: access, ArticleID: req.GetArticleId()})
}

func (s *ArticleService) MakePublic(ctx context.Context, req *v1.MakePublicArticle_Req) (*v1.MakePublicArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.MakePublicArticle_Resp{}, s.articleUsecase.MakePublic(ctx, &usecase.ArticleMakePublicReq{Access: access, ArticleID: req.GetArticleId()})
}

func (s *ArticleService) Archive(ctx context.Context, req *v1.ArchiveArticle_Req) (*v1.ArchiveArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.ArchiveArticle_Resp{}, s.articleUsecase.Archive(ctx, &usecase.ArticleArchiveReq{Access: access, ArticleID: req.GetArticleId(), Reason: req.Reason})
}

func (s *ArticleService) Unarchive(ctx context.Context, req *v1.UnarchiveArticle_Req) (*v1.UnarchiveArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.UnarchiveArticle_Resp{}, s.articleUsecase.Unarchive(ctx, &usecase.ArticleUnarchiveReq{Access: access, ArticleID: req.GetArticleId(), Reason: req.Reason})
}

func (s *ArticleService) Hide(ctx context.Context, req *v1.HideArticle_Req) (*v1.HideArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.HideArticle_Resp{}, s.articleUsecase.Hide(ctx, &usecase.ArticleHideReq{Access: access, ArticleID: req.GetArticleId(), Reason: req.Reason})
}

func (s *ArticleService) Unhide(ctx context.Context, req *v1.UnhideArticle_Req) (*v1.UnhideArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.UnhideArticle_Resp{}, s.articleUsecase.Unhide(ctx, &usecase.ArticleUnhideReq{Access: access, ArticleID: req.GetArticleId(), Reason: req.Reason})
}

func (s *ArticleService) Lock(ctx context.Context, req *v1.LockArticle_Req) (*v1.LockArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.LockArticle_Resp{}, s.articleUsecase.Lock(ctx, &usecase.ArticleLockReq{Access: access, ArticleID: req.GetArticleId(), Reason: req.Reason})
}

func (s *ArticleService) Unlock(ctx context.Context, req *v1.UnlockArticle_Req) (*v1.UnlockArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.UnlockArticle_Resp{}, s.articleUsecase.Unlock(ctx, &usecase.ArticleUnlockReq{Access: access, ArticleID: req.GetArticleId(), Reason: req.Reason})
}

func (s *ArticleService) Like(ctx context.Context, req *v1.LikeArticle_Req) (*v1.LikeArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	liked, err := s.articleUsecase.Like(ctx, &usecase.ArticleLikeReq{Access: access, ArticleID: req.GetArticleId(), UserID: access.ActorUserID, Active: req.GetLiked()})
	return &v1.LikeArticle_Resp{Liked: liked}, err
}

func (s *ArticleService) Thank(ctx context.Context, req *v1.ThankArticle_Req) (*v1.ThankArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	thanked, err := s.articleUsecase.Thank(ctx, &usecase.ArticleThankReq{Access: access, ArticleID: req.GetArticleId(), UserID: access.ActorUserID, Active: req.GetThanked()})
	return &v1.ThankArticle_Resp{Thanked: thanked}, err
}

func (s *ArticleService) Collect(ctx context.Context, req *v1.CollectArticle_Req) (*v1.CollectArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	collected, err := s.articleUsecase.Collect(ctx, &usecase.ArticleCollectReq{Access: access, ArticleID: req.GetArticleId(), UserID: access.ActorUserID, Active: req.GetCollected()})
	return &v1.CollectArticle_Resp{Collected: collected}, err
}

func (s *ArticleService) Reward(ctx context.Context, req *v1.RewardArticle_Req) (*v1.RewardArticle_Resp, error) {
	access := &model.ContentAccess{Scope: ""}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize("")
	if err != nil {
		return nil, err
	}
	return &v1.RewardArticle_Resp{}, s.articleUsecase.Reward(ctx, &usecase.ArticleRewardReq{Access: access, ArticleID: req.GetArticleId(), UserID: access.ActorUserID, Points: req.GetPoints()})
}

func (s *ArticleService) View(ctx context.Context, req *v1.ViewArticle_Req) (*v1.ViewArticle_Resp, error) {
	access := &model.ContentAccess{Scope: enum.ContentAccessScopeGuest}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return nil, err
	}
	return &v1.ViewArticle_Resp{}, s.articleUsecase.View(ctx, &usecase.ArticleViewReq{Access: access, ArticleID: req.GetArticleId(), IP: req.Ip, UserAgent: req.UserAgent, BrowserFingerprint: req.BrowserFingerprint})
}

func (s *ArticleService) FlushViews(ctx context.Context, req *v1.FlushArticleViews_Req) (*v1.FlushArticleViews_Resp, error) {
	flushed, err := s.articleUsecase.FlushViews(ctx, &usecase.ArticleFlushViewsReq{Limit: req.GetLimit()})
	if err != nil {
		return nil, err
	}
	return &v1.FlushArticleViews_Resp{Flushed: flushed}, nil
}

func (s *ArticleService) Get(ctx context.Context, req *v1.GetArticle_Req) (*v1.GetArticle_Resp, error) {
	access := &model.ContentAccess{Scope: enum.ContentAccessScopeGuest}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return nil, err
	}
	row, err := s.articleUsecase.GetByAccess(ctx, &usecase.ArticleGetByAccessReq{ArticleID: req.GetArticleId(), Access: access})
	if err != nil {
		return nil, err
	}
	article := &v1.Article{
		Id:            row.ID,
		Title:         row.Title,
		Content:       row.Content,
		RewardContent: row.RewardContent,
		RewardPoints:  row.RewardPoints,
		HasPostscript: row.HasPostscript,
		HasReward:     row.RewardPoints != nil,
		Type:          enum.ArticleTypeMap.MustToProto(row.Type),
		Statement:     row.Statement,
		Commentable:   row.Commentable,
		PublishStatus: enum.ArticlePublishStatusMap.MustToProto(row.PublishStatus),
		Visibility:    enum.ArticleVisibilityMap.MustToProto(row.Visibility),
		Restriction:   enum.ContentRestrictionMap.MustToProto(row.Restriction),
		ViewCount:     row.ViewCount,
		ThankCount:    row.ThankCount,
		LikeCount:     row.LikeCount,
		CollectCount:  row.CollectCount,
		RewardCount:   row.RewardCount,
		ReplyCount:    row.ReplyCount,
		CreatedBy:     row.CreatedBy,
		UpdatedBy:     row.UpdatedBy,
	}
	if row.CreatedAt != nil {
		article.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		article.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	if row.PublishedAt != nil {
		article.PublishedAt = timestamppb.New(*row.PublishedAt)
	}
	if row.EditedAt != nil {
		article.EditedAt = timestamppb.New(*row.EditedAt)
	}
	return &v1.GetArticle_Resp{Article: article}, nil
}

func (s *ArticleService) List(ctx context.Context, req *v1.ListArticles_Req) (*v1.ListArticles_Resp, error) {
	access := &model.ContentAccess{Scope: enum.ContentAccessScopeGuest}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return nil, err
	}
	query := &model.ArticleFilter{}
	if req.GetQuery() != nil {
		query.TagID = req.GetQuery().TagId
		query.DomainID = req.GetQuery().DomainId
		query.ArticleIDs = req.GetQuery().GetArticleIds()
		query.AuthorID = req.GetQuery().AuthorId
		query.Keyword = req.GetQuery().Keyword
		if req.GetQuery().Type != nil {
			v, ok := enum.ArticleTypeMap.ToEnum(req.GetQuery().GetType())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
			}
			query.Type = new(v)
		}
		if req.GetQuery().Order != nil {
			v, ok := enum.ArticleOrderMap.ToEnum(req.GetQuery().GetOrder())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			query.Order = new(v)
		}
		if req.GetQuery().PublishStatus != nil {
			v, ok := enum.ArticlePublishStatusMap.ToEnum(req.GetQuery().GetPublishStatus())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
			}
			query.PublishStatus = new(v)
		}
		for _, item := range req.GetQuery().GetPublishStatuses() {
			v, ok := enum.ArticlePublishStatusMap.ToEnum(item)
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
			}
			query.PublishStatuses = append(query.PublishStatuses, v)
		}
		if req.GetQuery().Visibility != nil {
			v, ok := enum.ArticleVisibilityMap.ToEnum(req.GetQuery().GetVisibility())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
			}
			query.Visibility = new(v)
		}
		for _, item := range req.GetQuery().GetVisibilities() {
			v, ok := enum.ArticleVisibilityMap.ToEnum(item)
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
			}
			query.Visibilities = append(query.Visibilities, v)
		}
		if req.GetQuery().Restriction != nil {
			v, ok := enum.ContentRestrictionMap.ToEnum(req.GetQuery().GetRestriction())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			query.Restriction = new(v)
		}
		for _, item := range req.GetQuery().GetRestrictions() {
			v, ok := enum.ContentRestrictionMap.ToEnum(item)
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			query.Restrictions = append(query.Restrictions, v)
		}
		if req.GetQuery().GetPublishedAtEnd() != nil {
			end := req.GetQuery().GetPublishedAtEnd().AsTime()
			query.PublishedAtEnd = new(end)
		}
	}
	list, err := s.articleUsecase.ListByAccess(ctx, &usecase.ArticleListByAccessReq{Access: access, Filter: query})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.Article, 0, len(list))
	for _, row := range list {
		article := &v1.Article{
			Id:            row.ID,
			Title:         row.Title,
			Content:       row.Content,
			RewardContent: row.RewardContent,
			RewardPoints:  row.RewardPoints,
			HasPostscript: row.HasPostscript,
			HasReward:     row.RewardPoints != nil,
			Type:          enum.ArticleTypeMap.MustToProto(row.Type),
			Statement:     row.Statement,
			Commentable:   row.Commentable,
			PublishStatus: enum.ArticlePublishStatusMap.MustToProto(row.PublishStatus),
			Visibility:    enum.ArticleVisibilityMap.MustToProto(row.Visibility),
			Restriction:   enum.ContentRestrictionMap.MustToProto(row.Restriction),
			ViewCount:     row.ViewCount,
			ThankCount:    row.ThankCount,
			LikeCount:     row.LikeCount,
			CollectCount:  row.CollectCount,
			RewardCount:   row.RewardCount,
			ReplyCount:    row.ReplyCount,
			CreatedBy:     row.CreatedBy,
			UpdatedBy:     row.UpdatedBy,
		}
		if row.CreatedAt != nil {
			article.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			article.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		if row.PublishedAt != nil {
			article.PublishedAt = timestamppb.New(*row.PublishedAt)
		}
		if row.EditedAt != nil {
			article.EditedAt = timestamppb.New(*row.EditedAt)
		}
		rows = append(rows, article)
	}
	return &v1.ListArticles_Resp{Rows: rows}, nil
}

func (s *ArticleService) Page(ctx context.Context, req *v1.PageArticles_Req) (*v1.PageArticles_Resp, error) {
	access := &model.ContentAccess{Scope: enum.ContentAccessScopeGuest}
	if req.GetAccess() != nil {
		if req.GetAccess().GetScope() != 0 {
			scope, ok := enum.ContentAccessScopeMap.ToEnum(req.GetAccess().GetScope())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			access.Scope = scope
		}
		if req.GetAccess().ActorUserId != nil {
			access.ActorUserID = req.GetAccess().GetActorUserId()
		}
	}
	access, err := access.Normalize(enum.ContentAccessScopeGuest)
	if err != nil {
		return nil, err
	}
	query := &model.ArticleFilter{}
	if req.GetQuery() != nil {
		query.TagID = req.GetQuery().TagId
		query.DomainID = req.GetQuery().DomainId
		query.ArticleIDs = req.GetQuery().GetArticleIds()
		query.AuthorID = req.GetQuery().AuthorId
		query.Keyword = req.GetQuery().Keyword
		if req.GetQuery().Type != nil {
			v, ok := enum.ArticleTypeMap.ToEnum(req.GetQuery().GetType())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
			}
			query.Type = new(v)
		}
		if req.GetQuery().Order != nil {
			v, ok := enum.ArticleOrderMap.ToEnum(req.GetQuery().GetOrder())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			query.Order = new(v)
		}
		if req.GetQuery().PublishStatus != nil {
			v, ok := enum.ArticlePublishStatusMap.ToEnum(req.GetQuery().GetPublishStatus())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
			}
			query.PublishStatus = new(v)
		}
		for _, item := range req.GetQuery().GetPublishStatuses() {
			v, ok := enum.ArticlePublishStatusMap.ToEnum(item)
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
			}
			query.PublishStatuses = append(query.PublishStatuses, v)
		}
		if req.GetQuery().Visibility != nil {
			v, ok := enum.ArticleVisibilityMap.ToEnum(req.GetQuery().GetVisibility())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
			}
			query.Visibility = new(v)
		}
		for _, item := range req.GetQuery().GetVisibilities() {
			v, ok := enum.ArticleVisibilityMap.ToEnum(item)
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
			}
			query.Visibilities = append(query.Visibilities, v)
		}
		if req.GetQuery().Restriction != nil {
			v, ok := enum.ContentRestrictionMap.ToEnum(req.GetQuery().GetRestriction())
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			query.Restriction = new(v)
		}
		for _, item := range req.GetQuery().GetRestrictions() {
			v, ok := enum.ContentRestrictionMap.ToEnum(item)
			if !ok {
				return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
			}
			query.Restrictions = append(query.Restrictions, v)
		}
		if req.GetQuery().GetPublishedAtEnd() != nil {
			end := req.GetQuery().GetPublishedAtEnd().AsTime()
			query.PublishedAtEnd = new(end)
		}
	}
	pageReq := req.GetPage()
	pageResp, err := s.articleUsecase.PageByAccess(ctx, &usecase.ArticlePageByAccessReq{Access: access, Filter: query, Page: &base.PageRequest{Page: int64(pageReq.GetPage()), Size: int64(pageReq.GetSize())}})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.Article, 0, len(pageResp.Rows))
	for _, row := range pageResp.Rows {
		article := &v1.Article{
			Id:            row.ID,
			Title:         row.Title,
			Content:       row.Content,
			RewardContent: row.RewardContent,
			RewardPoints:  row.RewardPoints,
			HasPostscript: row.HasPostscript,
			HasReward:     row.RewardPoints != nil,
			Type:          enum.ArticleTypeMap.MustToProto(row.Type),
			Statement:     row.Statement,
			Commentable:   row.Commentable,
			PublishStatus: enum.ArticlePublishStatusMap.MustToProto(row.PublishStatus),
			Visibility:    enum.ArticleVisibilityMap.MustToProto(row.Visibility),
			Restriction:   enum.ContentRestrictionMap.MustToProto(row.Restriction),
			ViewCount:     row.ViewCount,
			ThankCount:    row.ThankCount,
			LikeCount:     row.LikeCount,
			CollectCount:  row.CollectCount,
			RewardCount:   row.RewardCount,
			ReplyCount:    row.ReplyCount,
			CreatedBy:     row.CreatedBy,
			UpdatedBy:     row.UpdatedBy,
		}
		if row.CreatedAt != nil {
			article.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			article.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		if row.PublishedAt != nil {
			article.PublishedAt = timestamppb.New(*row.PublishedAt)
		}
		if row.EditedAt != nil {
			article.EditedAt = timestamppb.New(*row.EditedAt)
		}
		rows = append(rows, article)
	}
	return &v1.PageArticles_Resp{Rows: rows, Page: &common.PageResp{Page: uint32(pageResp.Page.Page), Size: uint32(pageResp.Page.Size), Total: uint32(pageResp.Page.Total)}}, nil
}

func (s *ArticleService) MapViewerActionStates(ctx context.Context, req *v1.MapArticleViewerActionStates_Req) (*v1.MapArticleViewerActionStates_Resp, error) {
	states, err := s.articleUsecase.MapViewerActionStates(ctx, &usecase.ArticleMapViewerActionStatesReq{ArticleIDs: req.GetArticleIds(), UserID: req.GetUserId()})
	if err != nil {
		return nil, err
	}
	reply := make(map[int64]*v1.MapArticleViewerActionStates_Resp_ArticleViewerActionState, len(states))
	for articleID, state := range states {
		reply[articleID] = &v1.MapArticleViewerActionStates_Resp_ArticleViewerActionState{Liked: state.Liked, Thanked: state.Thanked, Collected: state.Collected, Rewarded: state.Rewarded}
	}
	return &v1.MapArticleViewerActionStates_Resp{States: reply}, nil
}
