package service

import (
	"bbs/internal/biz/usecase"
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/server"
	"common/pkg/util"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	bbscontentv1enum "common/proto/gen/bbs/v1/content/enum"
	bbsuserv1enum "common/proto/gen/bbs/v1/user/enum"
	"common/proto/gen/common"
	cerrors "common/proto/gen/common/errors"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ContentArticleService struct {
	bbscontentv1.UnimplementedArticleServiceServer
	contentArticleUsecase *usecase.ContentArticleUsecase
}

func NewContentArticleService(
	contentArticleUsecase *usecase.ContentArticleUsecase,
) *ContentArticleService {
	return &ContentArticleService{contentArticleUsecase: contentArticleUsecase}
}

func (s *ContentArticleService) RegisterGrpc(gs *grpc.Server) {}

func (s *ContentArticleService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterArticleServiceHTTPServer(hs, s)
}

func (s *ContentArticleService) CreateDraft(ctx context.Context, req *bbscontentv1.CreateDraftArticle_Req) (*bbscontentv1.CreateDraftArticle_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	article := req.GetArticle()
	if article == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row, err := s.contentArticleUsecase.CreateDraftArticle(ctx, &usecase.CreateDraftArticleReq{UserID: user.ID, Article: &usecase.ContentArticleSave{Title: article.GetTitle(), Content: article.GetContent(), RewardContent: article.RewardContent, RewardPoints: article.RewardPoints, Type: int32(article.GetType()), Statement: article.Statement, Commentable: article.Commentable}})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.CreateDraftArticle_Resp{Article: s.articleDetail(row)}, nil
}

func (s *ContentArticleService) UpdateDraft(ctx context.Context, req *bbscontentv1.UpdateDraftArticle_Req) (*bbscontentv1.UpdateDraftArticle_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	article := req.GetArticle()
	if article == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row, err := s.contentArticleUsecase.UpdateDraftArticle(ctx, &usecase.UpdateDraftArticleReq{UserID: user.ID, ArticleID: req.GetArticleId(), Article: &usecase.ContentArticleSave{Title: article.GetTitle(), Content: article.GetContent(), RewardContent: article.RewardContent, RewardPoints: article.RewardPoints, Type: int32(article.GetType()), Statement: article.Statement, Commentable: article.Commentable}})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.UpdateDraftArticle_Resp{Article: s.articleDetail(row)}, nil
}

func (s *ContentArticleService) Publish(ctx context.Context, req *bbscontentv1.PublishArticle_Req) (*bbscontentv1.PublishArticle_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	err := s.contentArticleUsecase.PublishArticle(ctx, &usecase.PublishArticleReq{UserID: user.ID, ArticleID: req.GetArticleId(), Visibility: int32(req.GetVisibility())})
	return &bbscontentv1.PublishArticle_Resp{}, err
}

func (s *ContentArticleService) SchedulePublish(ctx context.Context, req *bbscontentv1.SchedulePublishArticle_Req) (*bbscontentv1.SchedulePublishArticle_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	if req.GetPublishAt() == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	err := s.contentArticleUsecase.SchedulePublishArticle(ctx, &usecase.SchedulePublishArticleReq{UserID: user.ID, ArticleID: req.GetArticleId(), PublishAt: req.GetPublishAt().AsTime()})
	return &bbscontentv1.SchedulePublishArticle_Resp{}, err
}

func (s *ContentArticleService) CancelPublish(ctx context.Context, req *bbscontentv1.CancelPublishArticle_Req) (*bbscontentv1.CancelPublishArticle_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	err := s.contentArticleUsecase.CancelPublishArticle(ctx, &usecase.CancelPublishArticleReq{UserID: user.ID, ArticleID: req.GetArticleId()})
	return &bbscontentv1.CancelPublishArticle_Resp{}, err
}

func (s *ContentArticleService) DiscardDraft(ctx context.Context, req *bbscontentv1.DiscardDraftArticle_Req) (*bbscontentv1.DiscardDraftArticle_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	err := s.contentArticleUsecase.DiscardDraftArticle(ctx, &usecase.DiscardDraftArticleReq{UserID: user.ID, ArticleID: req.GetArticleId()})
	return &bbscontentv1.DiscardDraftArticle_Resp{}, err
}

func (s *ContentArticleService) Archive(ctx context.Context, req *bbscontentv1.ArchiveArticle_Req) (*bbscontentv1.ArchiveArticle_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	err := s.contentArticleUsecase.ArchiveArticle(ctx, &usecase.ArchiveArticleReq{UserID: user.ID, ArticleID: req.GetArticleId(), Reason: req.Reason})
	return &bbscontentv1.ArchiveArticle_Resp{}, err
}

func (s *ContentArticleService) List(ctx context.Context, req *bbscontentv1.ListArticles_Req) (*bbscontentv1.ListArticles_Resp, error) {
	var userID int64
	if user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo); ok && user != nil {
		userID = user.ID
	}
	var query *usecase.ArticleQuery
	if req.GetQuery() != nil {
		q := req.GetQuery()
		query = &usecase.ArticleQuery{TagID: q.TagId, DomainID: q.DomainId, Keyword: q.Keyword, AuthorID: q.AuthorId}
		if q.Type != nil {
			query.Type = new(int32(*q.Type))
		}
		if q.Order != nil {
			query.Order = new(int32(*q.Order))
		}
		if q.PublishStatus != nil {
			query.PublishStatus = new(int32(*q.PublishStatus))
		}
		if q.Visibility != nil {
			query.Visibility = new(int32(*q.Visibility))
		}
	}
	resp, err := s.contentArticleUsecase.ListArticles(ctx, &usecase.ListArticlesReq{UserID: userID, Page: req.GetPage(), Query: query})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbscontentv1.ArticleListItem, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		rows = append(rows, s.articleListItem(row))
	}
	var page *common.PageResp
	if resp.Page != nil {
		page = &common.PageResp{Page: resp.Page.Page, Size: resp.Page.Size, Total: resp.Page.Total}
	}
	return &bbscontentv1.ListArticles_Resp{Page: page, Rows: rows}, nil
}

func (s *ContentArticleService) Get(ctx context.Context, req *bbscontentv1.GetArticle_Req) (*bbscontentv1.GetArticle_Resp, error) {
	var userID int64
	if user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo); ok && user != nil {
		userID = user.ID
	}
	row, err := s.contentArticleUsecase.GetArticle(ctx, &usecase.GetArticleReq{UserID: userID, ArticleID: req.GetArticleId()})
	if err != nil {
		return nil, err
	}
	if row.PublishStatus == int32(bbscontentv1enum.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_PUBLISHED) && row.Visibility == int32(bbscontentv1enum.ArticleVisibility_ARTICLE_VISIBILITY_PUBLIC) && row.Restriction == int32(bbscontentv1enum.ContentRestriction_CONTENT_RESTRICTION_NONE) {
		ip := server.ClientIP(ctx)
		userAgent := server.GetHeader(ctx, constant.HeaderUserAgent)
		_ = s.contentArticleUsecase.ViewArticle(ctx, &usecase.ViewArticleReq{UserID: userID, ArticleID: req.GetArticleId(), IP: &ip, UserAgent: &userAgent})
	}
	return &bbscontentv1.GetArticle_Resp{Article: s.articleDetail(row)}, nil
}

func (s *ContentArticleService) Like(ctx context.Context, req *bbscontentv1.LikeArticle_Req) (*bbscontentv1.LikeArticle_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	v, err := s.contentArticleUsecase.LikeArticle(ctx, &usecase.LikeArticleReq{UserID: user.ID, ArticleID: req.GetArticleId(), Active: req.GetActive()})
	return &bbscontentv1.LikeArticle_Resp{Liked: v}, err
}
func (s *ContentArticleService) Thank(ctx context.Context, req *bbscontentv1.ThankArticle_Req) (*bbscontentv1.ThankArticle_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	v, err := s.contentArticleUsecase.ThankArticle(ctx, &usecase.ThankArticleReq{UserID: user.ID, ArticleID: req.GetArticleId(), Active: req.GetActive()})
	return &bbscontentv1.ThankArticle_Resp{Thanked: v}, err
}
func (s *ContentArticleService) Collect(ctx context.Context, req *bbscontentv1.CollectArticle_Req) (*bbscontentv1.CollectArticle_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	v, err := s.contentArticleUsecase.CollectArticle(ctx, &usecase.CollectArticleReq{UserID: user.ID, ArticleID: req.GetArticleId(), Active: req.GetActive()})
	return &bbscontentv1.CollectArticle_Resp{Collected: v}, err
}
func (s *ContentArticleService) Reward(ctx context.Context, req *bbscontentv1.RewardArticle_Req) (*bbscontentv1.RewardArticle_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	err := s.contentArticleUsecase.RewardArticle(ctx, &usecase.RewardArticleReq{UserID: user.ID, ArticleID: req.GetArticleId(), Points: req.GetPoints()})
	return &bbscontentv1.RewardArticle_Resp{}, err
}

func (s *ContentArticleService) articleListItem(row *usecase.ContentArticleListItem) *bbscontentv1.ArticleListItem {
	if row == nil {
		return nil
	}
	out := &bbscontentv1.ArticleListItem{Id: row.ID, Title: row.Title, Content: row.Content, ContentRender: row.ContentRender, HasPostscript: row.HasPostscript, HasReward: row.HasReward, PublishStatus: bbscontentv1enum.ArticlePublishStatus(row.PublishStatus), Visibility: bbscontentv1enum.ArticleVisibility(row.Visibility), Restriction: bbscontentv1enum.ContentRestriction(row.Restriction), Type: bbscontentv1enum.ArticleType(row.Type), Statement: row.Statement, Commentable: row.Commentable, ViewCount: row.ViewCount, ThankCount: row.ThankCount, LikeCount: row.LikeCount, CollectCount: row.CollectCount, RewardCount: row.RewardCount, ReplyCount: row.ReplyCount, CoverImageUrl: row.CoverImageURL, ViewerActionState: s.articleViewerActionState(row.ViewerActionState), LastReplyUser: s.accountProfile(row.LastReplyUser), AuthorUser: s.accountProfile(row.AuthorUser), CreatedBy: row.CreatedBy, UpdatedBy: row.UpdatedBy}
	if row.LastReplyAt != nil {
		out.LastReplyAt = timestamppb.New(*row.LastReplyAt)
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

func (s *ContentArticleService) articleDetail(row *usecase.ContentArticleDetail) *bbscontentv1.ArticleDetail {
	if row == nil {
		return nil
	}
	postscripts := make([]*bbscontentv1.ArticlePostscript, 0, len(row.Postscripts))
	for _, item := range row.Postscripts {
		postscripts = append(postscripts, s.articlePostscript(item))
	}
	out := &bbscontentv1.ArticleDetail{Id: row.ID, Title: row.Title, Content: row.Content, ContentRender: row.ContentRender, HasPostscript: row.HasPostscript, HasReward: row.HasReward, RewardContent: row.RewardContent, RewardContentRender: row.RewardContentRender, RewardPoints: row.RewardPoints, PublishStatus: bbscontentv1enum.ArticlePublishStatus(row.PublishStatus), Visibility: bbscontentv1enum.ArticleVisibility(row.Visibility), Restriction: bbscontentv1enum.ContentRestriction(row.Restriction), Type: bbscontentv1enum.ArticleType(row.Type), Statement: row.Statement, Commentable: row.Commentable, ViewCount: row.ViewCount, ThankCount: row.ThankCount, LikeCount: row.LikeCount, CollectCount: row.CollectCount, RewardCount: row.RewardCount, ReplyCount: row.ReplyCount, CoverImageUrl: row.CoverImageURL, ViewerActionState: s.articleViewerActionState(row.ViewerActionState), LastReplyUser: s.accountProfile(row.LastReplyUser), Postscripts: postscripts, AuthorUser: s.accountProfile(row.AuthorUser), CreatedBy: row.CreatedBy, UpdatedBy: row.UpdatedBy}
	if row.LastReplyAt != nil {
		out.LastReplyAt = timestamppb.New(*row.LastReplyAt)
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

func (s *ContentArticleService) articleViewerActionState(row *usecase.ContentArticleViewerActionState) *bbscontentv1.ArticleViewerActionState {
	if row == nil {
		return nil
	}
	return &bbscontentv1.ArticleViewerActionState{Liked: row.Liked, Thanked: row.Thanked, Collected: row.Collected, Rewarded: row.Rewarded}
}
func (s *ContentArticleService) articlePostscript(row *usecase.ContentArticlePostscript) *bbscontentv1.ArticlePostscript {
	if row == nil {
		return nil
	}
	out := &bbscontentv1.ArticlePostscript{Id: row.ID, ArticleId: row.ArticleID, Content: row.Content, ContentRender: row.ContentRender, Restriction: bbscontentv1enum.ContentRestriction(row.Restriction), CreatedBy: row.CreatedBy, UpdatedBy: row.UpdatedBy}
	if row.CreatedAt != nil {
		out.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		out.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return out
}
func (s *ContentArticleService) accountProfile(row *usecase.ContentAccountProfile) *bbscontentv1.AccountProfile {
	if row == nil {
		return nil
	}
	out := &bbscontentv1.AccountProfile{Id: row.ID, Name: row.Name, Nickname: row.Nickname, Url: row.URL, AvatarUrl: row.AvatarURL, Introduction: row.Introduction, Mbti: bbsuserv1enum.MBTI(row.MBTI), Status: bbsuserv1enum.AccountStatus(row.Status), FollowCount: row.FollowCount, FollowerCount: row.FollowerCount}
	if row.CreatedAt != nil {
		out.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		out.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return out
}
