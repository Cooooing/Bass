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
	article := req.GetArticle()
	if article == nil || req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	articleType, ok := enum.ArticleTypeMap.ToEnum(article.GetType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
	}
	row, err := s.articleUsecase.Add(ctx, &usecase.ArticleAddReq{
		Article: &model.Article{
			Title:         article.GetTitle(),
			Content:       article.GetContent(),
			RewardContent: article.RewardContent,
			RewardPoints:  article.RewardPoints,
			Type:          articleType,
			Statement:     article.Statement,
			Commentable:   util.DerefOrDefault(article.Commentable, true),
			CreatedBy:     new(req.UserId),
			UpdatedBy:     new(req.UserId),
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateDraftArticle_Resp{Article: s.article(row)}, nil
}

func (s *ArticleService) UpdateDraft(ctx context.Context, req *v1.UpdateDraftArticle_Req) (*v1.UpdateDraftArticle_Resp, error) {
	article := req.GetArticle()
	if article == nil || req.GetUserId() <= 0 || req.GetArticleId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	articleType, ok := enum.ArticleTypeMap.ToEnum(article.GetType())
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_TYPE)
	}
	row, err := s.articleUsecase.Update(ctx, &model.Article{
		ID:            req.GetArticleId(),
		Title:         article.GetTitle(),
		Content:       article.GetContent(),
		RewardContent: article.RewardContent,
		RewardPoints:  article.RewardPoints,
		Type:          articleType,
		Statement:     article.Statement,
		Commentable:   util.DerefOrDefault(article.Commentable, true),
		UpdatedBy:     new(req.UserId),
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateDraftArticle_Resp{Article: s.article(row)}, nil
}

func (s *ArticleService) Publish(ctx context.Context, req *v1.PublishArticle_Req) (*v1.PublishArticle_Resp, error) {
	visibility := enum.ArticleVisibilityPublic
	if req.Visibility != 0 {
		item, ok := enum.ArticleVisibilityMap.ToEnum(req.Visibility)
		if !ok {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_INVALID_ARTICLE_STATUS)
		}
		visibility = item
	}
	if err := s.articleUsecase.Publish(ctx, &usecase.ArticlePublishReq{ArticleID: req.GetArticleId(), OperatorUserID: req.OperatorUserId, Visibility: visibility}); err != nil {
		return nil, err
	}
	return &v1.PublishArticle_Resp{}, nil
}

func (s *ArticleService) SchedulePublish(ctx context.Context, req *v1.SchedulePublishArticle_Req) (*v1.SchedulePublishArticle_Resp, error) {
	if req.GetOperatorUserId() <= 0 || req.GetPublishAt() == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.articleUsecase.SchedulePublish(ctx, &usecase.ArticleSchedulePublishReq{ArticleID: req.GetArticleId(), OperatorUserID: req.GetOperatorUserId(), PublishAt: req.GetPublishAt().AsTime()}); err != nil {
		return nil, err
	}
	return &v1.SchedulePublishArticle_Resp{}, nil
}

func (s *ArticleService) CancelPublish(ctx context.Context, req *v1.CancelPublishArticle_Req) (*v1.CancelPublishArticle_Resp, error) {
	if req.GetOperatorUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	if err := s.articleUsecase.CancelPublish(ctx, &usecase.ArticleCancelPublishReq{ArticleID: req.GetArticleId(), OperatorUserID: req.GetOperatorUserId()}); err != nil {
		return nil, err
	}
	return &v1.CancelPublishArticle_Resp{}, nil
}

func (s *ArticleService) DiscardDraft(ctx context.Context, req *v1.DiscardDraftArticle_Req) (*v1.DiscardDraftArticle_Resp, error) {
	if req.GetUserId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return &v1.DiscardDraftArticle_Resp{}, s.articleUsecase.DiscardDraft(ctx, &usecase.ArticleDiscardDraftReq{ArticleID: req.GetArticleId(), UserID: req.GetUserId()})
}

func (s *ArticleService) MakePrivate(ctx context.Context, req *v1.MakePrivateArticle_Req) (*v1.MakePrivateArticle_Resp, error) {
	return &v1.MakePrivateArticle_Resp{}, s.articleUsecase.MakePrivate(ctx, &usecase.ArticleMakePrivateReq{ArticleID: req.GetArticleId(), UserID: req.GetUserId()})
}

func (s *ArticleService) MakePublic(ctx context.Context, req *v1.MakePublicArticle_Req) (*v1.MakePublicArticle_Resp, error) {
	return &v1.MakePublicArticle_Resp{}, s.articleUsecase.MakePublic(ctx, &usecase.ArticleMakePublicReq{ArticleID: req.GetArticleId(), UserID: req.GetUserId()})
}

func (s *ArticleService) Archive(ctx context.Context, req *v1.ArchiveArticle_Req) (*v1.ArchiveArticle_Resp, error) {
	return &v1.ArchiveArticle_Resp{}, s.articleUsecase.Archive(ctx, &usecase.ArticleArchiveReq{ArticleID: req.GetArticleId(), UserID: req.GetUserId(), Reason: req.Reason})
}

func (s *ArticleService) Unarchive(ctx context.Context, req *v1.UnarchiveArticle_Req) (*v1.UnarchiveArticle_Resp, error) {
	return &v1.UnarchiveArticle_Resp{}, s.articleUsecase.Unarchive(ctx, &usecase.ArticleUnarchiveReq{ArticleID: req.GetArticleId(), UserID: req.GetUserId(), Reason: req.Reason})
}

func (s *ArticleService) Hide(ctx context.Context, req *v1.HideArticle_Req) (*v1.HideArticle_Resp, error) {
	return &v1.HideArticle_Resp{}, s.articleUsecase.Hide(ctx, &usecase.ArticleHideReq{ArticleID: req.GetArticleId(), UserID: req.GetUserId(), Reason: req.Reason})
}

func (s *ArticleService) Unhide(ctx context.Context, req *v1.UnhideArticle_Req) (*v1.UnhideArticle_Resp, error) {
	return &v1.UnhideArticle_Resp{}, s.articleUsecase.Unhide(ctx, &usecase.ArticleUnhideReq{ArticleID: req.GetArticleId(), UserID: req.GetUserId(), Reason: req.Reason})
}

func (s *ArticleService) Lock(ctx context.Context, req *v1.LockArticle_Req) (*v1.LockArticle_Resp, error) {
	return &v1.LockArticle_Resp{}, s.articleUsecase.Lock(ctx, &usecase.ArticleLockReq{ArticleID: req.GetArticleId(), UserID: req.GetUserId(), Reason: req.Reason})
}

func (s *ArticleService) Unlock(ctx context.Context, req *v1.UnlockArticle_Req) (*v1.UnlockArticle_Resp, error) {
	return &v1.UnlockArticle_Resp{}, s.articleUsecase.Unlock(ctx, &usecase.ArticleUnlockReq{ArticleID: req.GetArticleId(), UserID: req.GetUserId(), Reason: req.Reason})
}

func (s *ArticleService) Like(ctx context.Context, req *v1.LikeArticle_Req) (*v1.LikeArticle_Resp, error) {
	liked, err := s.articleUsecase.Like(ctx, &usecase.ArticleLikeReq{ArticleID: req.GetArticleId(), UserID: req.GetUserId(), Active: req.GetLiked()})
	return &v1.LikeArticle_Resp{Liked: liked}, err
}

func (s *ArticleService) Thank(ctx context.Context, req *v1.ThankArticle_Req) (*v1.ThankArticle_Resp, error) {
	thanked, err := s.articleUsecase.Thank(ctx, &usecase.ArticleThankReq{ArticleID: req.GetArticleId(), UserID: req.GetUserId(), Active: req.GetThanked()})
	return &v1.ThankArticle_Resp{Thanked: thanked}, err
}

func (s *ArticleService) Collect(ctx context.Context, req *v1.CollectArticle_Req) (*v1.CollectArticle_Resp, error) {
	collected, err := s.articleUsecase.Collect(ctx, &usecase.ArticleCollectReq{ArticleID: req.GetArticleId(), UserID: req.GetUserId(), Active: req.GetCollected()})
	return &v1.CollectArticle_Resp{Collected: collected}, err
}

func (s *ArticleService) Reward(ctx context.Context, req *v1.RewardArticle_Req) (*v1.RewardArticle_Resp, error) {
	return &v1.RewardArticle_Resp{}, s.articleUsecase.Reward(ctx, &usecase.ArticleRewardReq{ArticleID: req.GetArticleId(), UserID: req.GetUserId(), Points: req.GetPoints()})
}

func (s *ArticleService) View(ctx context.Context, req *v1.ViewArticle_Req) (*v1.ViewArticle_Resp, error) {
	return &v1.ViewArticle_Resp{}, s.articleUsecase.View(ctx, &usecase.ArticleViewReq{ArticleID: req.GetArticleId(), ViewerUserID: req.ViewerUserId, IP: req.Ip, UserAgent: req.UserAgent, BrowserFingerprint: req.BrowserFingerprint})
}

func (s *ArticleService) FlushViews(ctx context.Context, req *v1.FlushArticleViews_Req) (*v1.FlushArticleViews_Resp, error) {
	flushed, err := s.articleUsecase.FlushViews(ctx, &usecase.ArticleFlushViewsReq{Limit: req.GetLimit()})
	if err != nil {
		return nil, err
	}
	return &v1.FlushArticleViews_Resp{Flushed: flushed}, nil
}

func (s *ArticleService) Get(ctx context.Context, req *v1.GetArticle_Req) (*v1.GetArticle_Resp, error) {
	row, err := s.articleUsecase.Get(ctx, req.GetArticleId())
	if err != nil {
		return nil, err
	}
	return &v1.GetArticle_Resp{Article: s.article(row)}, nil
}

func (s *ArticleService) List(ctx context.Context, req *v1.ListArticles_Req) (*v1.ListArticles_Resp, error) {
	pageResp, err := s.articleUsecase.Page(ctx, &usecase.ArticlePageReq{Page: &base.PageRequest{Page: 1, Size: 1000}})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.Article, 0, len(pageResp.Rows))
	for _, row := range pageResp.Rows {
		rows = append(rows, s.article(row))
	}
	return &v1.ListArticles_Resp{Rows: rows}, nil
}

func (s *ArticleService) Page(ctx context.Context, req *v1.PageArticles_Req) (*v1.PageArticles_Resp, error) {
	query := req.GetQuery()
	pageReq := req.GetPage()
	usecaseReq := &usecase.ArticlePageReq{Page: &base.PageRequest{Page: int64(pageReq.GetPage()), Size: int64(pageReq.GetSize())}}
	if query != nil {
		usecaseReq.TagID = query.TagId
		usecaseReq.DomainID = query.DomainId
		usecaseReq.AuthorID = query.AuthorId
		usecaseReq.Keyword = query.Keyword
		if query.GetPublishedAtEnd() != nil {
			end := query.GetPublishedAtEnd().AsTime()
			usecaseReq.PublishedAtEnd = new(end)
		}
	}
	pageResp, err := s.articleUsecase.Page(ctx, usecaseReq)
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.Article, 0, len(pageResp.Rows))
	for _, row := range pageResp.Rows {
		rows = append(rows, s.article(row))
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

func (s *ArticleService) article(row *model.Article) *v1.Article {
	if row == nil {
		return nil
	}
	out := &v1.Article{
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
		out.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		out.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	if row.PublishedAt != nil {
		out.PublishedAt = timestamppb.New(*row.PublishedAt)
	}
	if row.EditedAt != nil {
		out.EditedAt = timestamppb.New(*row.EditedAt)
	}
	return out
}
