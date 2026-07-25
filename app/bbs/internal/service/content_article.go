package service

import (
	"bbs/internal/biz/usecase"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	bbscontentv1enum "common/proto/gen/bbs/v1/content/enum"
	bbsuserv1enum "common/proto/gen/bbs/v1/user/enum"
	"common/proto/gen/common"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/proto"
)

func cloneMessage[T proto.Message](src proto.Message, dst T) T {
	if src == nil {
		return dst
	}
	data, err := proto.Marshal(src)
	if err != nil {
		return dst
	}
	_ = proto.Unmarshal(data, dst)
	return dst
}

type ContentArticleService struct {
	bbscontentv1.UnimplementedArticleServiceServer
	contentArticleUsecase *usecase.ContentArticleUsecase
}

func NewContentArticleService(
	contentArticleUsecase *usecase.ContentArticleUsecase,
) *ContentArticleService {
	return &ContentArticleService{
		contentArticleUsecase: contentArticleUsecase,
	}
}

func (s *ContentArticleService) RegisterGrpc(gs *grpc.Server) {
}

func (s *ContentArticleService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterArticleServiceHTTPServer(hs, s)
}

func (s *ContentArticleService) Create(ctx context.Context, req *bbscontentv1.CreateArticle_Req) (*bbscontentv1.CreateArticle_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	var article *usecase.ContentArticleSave
	if row := req.GetArticle(); row != nil {
		article = &usecase.ContentArticleSave{
			Title:         row.GetTitle(),
			Content:       row.GetContent(),
			RewardContent: row.RewardContent,
			RewardPoints:  row.RewardPoints,
			Type:          row.GetType(),
			BountyPoints:  row.BountyPoints,
			Statement:     row.Statement,
			Commentable:   row.Commentable,
			Anonymous:     row.Anonymous,
			TagIDs:        row.GetTagIds(),
		}
	}
	resp, err := s.contentArticleUsecase.CreateArticle(ctx, &usecase.CreateArticleReq{
		UserID:  userID,
		Article: article,
	})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.CreateArticle_Resp{
		Article: s.articleDetail(resp),
	}, nil
}

func (s *ContentArticleService) Update(ctx context.Context, req *bbscontentv1.UpdateArticle_Req) (*bbscontentv1.UpdateArticle_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	var article *usecase.ContentArticleSave
	if row := req.GetArticle(); row != nil {
		article = &usecase.ContentArticleSave{
			Title:         row.GetTitle(),
			Content:       row.GetContent(),
			RewardContent: row.RewardContent,
			RewardPoints:  row.RewardPoints,
			Type:          row.GetType(),
			BountyPoints:  row.BountyPoints,
			Statement:     row.Statement,
			Commentable:   row.Commentable,
			Anonymous:     row.Anonymous,
			TagIDs:        row.GetTagIds(),
		}
	}
	resp, err := s.contentArticleUsecase.UpdateArticle(ctx, &usecase.UpdateArticleReq{
		UserID:    userID,
		ArticleID: req.GetArticleId(),
		Article:   article,
	})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.UpdateArticle_Resp{
		Article: cloneMessage(s.articleDetail(resp), &bbscontentv1.UpdateArticle_Resp_ArticleDetail{}),
	}, nil
}

func (s *ContentArticleService) UpdateDraft(ctx context.Context, req *bbscontentv1.UpdateDraftArticle_Req) (*bbscontentv1.UpdateDraftArticle_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	var article *usecase.ContentArticleSave
	if row := req.GetArticle(); row != nil {
		article = &usecase.ContentArticleSave{
			Title:         row.GetTitle(),
			Content:       row.GetContent(),
			RewardContent: row.RewardContent,
			RewardPoints:  row.RewardPoints,
			Type:          row.GetType(),
			BountyPoints:  row.BountyPoints,
			Statement:     row.Statement,
			Commentable:   row.Commentable,
			Anonymous:     row.Anonymous,
			TagIDs:        row.GetTagIds(),
		}
	}
	resp, err := s.contentArticleUsecase.UpdateDraftArticle(ctx, &usecase.UpdateDraftArticleReq{
		UserID:    userID,
		ArticleID: req.GetArticleId(),
		Article:   article,
	})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.UpdateDraftArticle_Resp{
		Article: cloneMessage(s.articleDetail(resp), &bbscontentv1.UpdateDraftArticle_Resp_ArticleDetail{}),
	}, nil
}

func (s *ContentArticleService) Publish(ctx context.Context, req *bbscontentv1.PublishArticle_Req) (*bbscontentv1.PublishArticle_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	err = s.contentArticleUsecase.PublishArticle(ctx, &usecase.PublishArticleReq{
		UserID:     userID,
		ArticleID:  req.GetArticleId(),
		Visibility: req.GetVisibility(),
	})
	return &bbscontentv1.PublishArticle_Resp{}, err
}

func (s *ContentArticleService) DiscardDraft(ctx context.Context, req *bbscontentv1.DiscardDraftArticle_Req) (*bbscontentv1.DiscardDraftArticle_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	err = s.contentArticleUsecase.DiscardDraftArticle(ctx, &usecase.DiscardDraftArticleReq{
		UserID:    userID,
		ArticleID: req.GetArticleId(),
	})
	return &bbscontentv1.DiscardDraftArticle_Resp{}, err
}

func (s *ContentArticleService) List(ctx context.Context, req *bbscontentv1.ListArticles_Req) (*bbscontentv1.ListArticles_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.contentArticleUsecase.ListArticles(ctx, &usecase.ListArticlesReq{
		UserID: userID,
		Page:   req.GetPage(),
		Query:  req.GetQuery(),
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbscontentv1.ListArticles_Resp_ArticleListItem, 0, len(resp.Rows))
	for _, row := range resp.Rows {
		rows = append(rows, s.articleListItem(row))
	}
	var page *common.PageResp
	if resp.Page != nil {
		page = &common.PageResp{
			Page:  resp.Page.Page,
			Size:  resp.Page.Size,
			Total: resp.Page.Total,
		}
	}
	return &bbscontentv1.ListArticles_Resp{
		Page: page,
		Rows: rows,
	}, nil
}

func (s *ContentArticleService) Get(ctx context.Context, req *bbscontentv1.GetArticle_Req) (*bbscontentv1.GetArticle_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.contentArticleUsecase.GetArticle(ctx, &usecase.GetArticleReq{
		UserID:    userID,
		ArticleID: req.GetArticleId(),
	})
	if err != nil {
		return nil, err
	}
	article := resp
	if article.PublishStatus == int32(bbscontentv1enum.ArticlePublishStatus_ARTICLE_PUBLISH_STATUS_PUBLISHED) && article.Visibility == int32(bbscontentv1enum.ArticleVisibility_ARTICLE_VISIBILITY_PUBLIC) && article.Restriction == int32(bbscontentv1enum.ContentRestriction_CONTENT_RESTRICTION_NONE) {
		if err := s.contentArticleUsecase.ViewArticle(ctx, &usecase.ViewArticleReq{
			UserID:    userID,
			ArticleID: req.GetArticleId(),
		}); err != nil {
			return nil, err
		}
	}
	return &bbscontentv1.GetArticle_Resp{
		Article: cloneMessage(s.articleDetail(resp), &bbscontentv1.GetArticle_Resp_ArticleDetail{}),
	}, nil
}

func (s *ContentArticleService) Like(ctx context.Context, req *bbscontentv1.LikeArticle_Req) (*bbscontentv1.LikeArticle_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.contentArticleUsecase.LikeArticle(ctx, &usecase.LikeArticleReq{
		UserID:    userID,
		ArticleID: req.GetArticleId(),
		Active:    req.GetActive(),
	})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.LikeArticle_Resp{
		Liked: resp,
	}, nil
}

func (s *ContentArticleService) Thank(ctx context.Context, req *bbscontentv1.ThankArticle_Req) (*bbscontentv1.ThankArticle_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.contentArticleUsecase.ThankArticle(ctx, &usecase.ThankArticleReq{
		UserID:    userID,
		ArticleID: req.GetArticleId(),
		Active:    req.GetActive(),
	})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.ThankArticle_Resp{
		Thanked: resp,
	}, nil
}

func (s *ContentArticleService) Collect(ctx context.Context, req *bbscontentv1.CollectArticle_Req) (*bbscontentv1.CollectArticle_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.contentArticleUsecase.CollectArticle(ctx, &usecase.CollectArticleReq{
		UserID:    userID,
		ArticleID: req.GetArticleId(),
		Active:    req.GetActive(),
	})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.CollectArticle_Resp{
		Collected: resp,
	}, nil
}

func (s *ContentArticleService) Watch(ctx context.Context, req *bbscontentv1.WatchArticle_Req) (*bbscontentv1.WatchArticle_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := s.contentArticleUsecase.WatchArticle(ctx, &usecase.WatchArticleReq{
		UserID:    userID,
		ArticleID: req.GetArticleId(),
		Active:    req.GetActive(),
	})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.WatchArticle_Resp{
		Watched: resp,
	}, nil
}

func (s *ContentArticleService) Reward(ctx context.Context, req *bbscontentv1.RewardArticle_Req) (*bbscontentv1.RewardArticle_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	err = s.contentArticleUsecase.RewardArticle(ctx, &usecase.RewardArticleReq{
		UserID:    userID,
		ArticleID: req.GetArticleId(),
		Points:    req.GetPoints(),
	})
	return &bbscontentv1.RewardArticle_Resp{}, err
}

func (s *ContentArticleService) AcceptAnswer(ctx context.Context, req *bbscontentv1.AcceptAnswerArticle_Req) (*bbscontentv1.AcceptAnswerArticle_Resp, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	err = s.contentArticleUsecase.AcceptAnswerArticle(ctx, &usecase.AcceptAnswerArticleReq{
		UserID:    userID,
		ArticleID: req.GetArticleId(),
		CommentID: req.GetCommentId(),
	})
	return &bbscontentv1.AcceptAnswerArticle_Resp{}, err
}

func (s *ContentArticleService) articleListItem(row *usecase.ContentArticleListItem) *bbscontentv1.ListArticles_Resp_ArticleListItem {
	if row == nil {
		return nil
	}
	return &bbscontentv1.ListArticles_Resp_ArticleListItem{
		Id:                row.ID,
		Title:             row.Title,
		Content:           row.Content,
		ContentRender:     row.ContentRender,
		HasPostscript:     row.HasPostscript,
		HasReward:         row.HasReward,
		PublishStatus:     bbscontentv1enum.ArticlePublishStatus(row.PublishStatus),
		Visibility:        bbscontentv1enum.ArticleVisibility(row.Visibility),
		Restriction:       bbscontentv1enum.ContentRestriction(row.Restriction),
		Type:              bbscontentv1enum.ArticleType(row.Type),
		BountyPoints:      row.BountyPoints,
		AcceptedAnswerId:  row.AcceptedAnswerID,
		Statement:         row.Statement,
		Commentable:       row.Commentable,
		Anonymous:         row.Anonymous,
		ViewCount:         row.ViewCount,
		ThankCount:        row.ThankCount,
		LikeCount:         row.LikeCount,
		CollectCount:      row.CollectCount,
		WatchCount:        row.WatchCount,
		ReplyCount:        row.ReplyCount,
		CoverImageUrl:     row.CoverImageURL,
		ViewerActionState: cloneMessage(s.articleViewerActionState(row.ViewerActionState), &bbscontentv1.ListArticles_Resp_ArticleViewerActionState{}),
		LastReplyAt:       row.LastReplyAt,
		LastReplyUser:     cloneMessage(s.accountProfile(row.LastReplyUser), &bbscontentv1.ListArticles_Resp_AccountProfile{}),
		AuthorUser:        cloneMessage(s.accountProfile(row.AuthorUser), &bbscontentv1.ListArticles_Resp_AccountProfile{}),
		CreatedBy:         row.CreatedBy,
		UpdatedBy:         row.UpdatedBy,
		CreatedAt:         row.CreatedAt,
		UpdatedAt:         row.UpdatedAt,
		PublishedAt:       row.PublishedAt,
		EditedAt:          row.EditedAt,
	}
}

func (s *ContentArticleService) articleDetail(row *usecase.ContentArticleDetail) *bbscontentv1.CreateArticle_Resp_ArticleDetail {
	if row == nil {
		return nil
	}
	postscripts := make([]*bbscontentv1.CreateArticle_Resp_ArticlePostscript, 0, len(row.Postscripts))
	for _, item := range row.Postscripts {
		postscripts = append(postscripts, s.articlePostscript(item))
	}
	return &bbscontentv1.CreateArticle_Resp_ArticleDetail{
		Id:                  row.ID,
		Title:               row.Title,
		Content:             row.Content,
		ContentRender:       row.ContentRender,
		HasPostscript:       row.HasPostscript,
		HasReward:           row.HasReward,
		RewardContent:       row.RewardContent,
		RewardContentRender: row.RewardContentRender,
		RewardPoints:        row.RewardPoints,
		PublishStatus:       bbscontentv1enum.ArticlePublishStatus(row.PublishStatus),
		Visibility:          bbscontentv1enum.ArticleVisibility(row.Visibility),
		Restriction:         bbscontentv1enum.ContentRestriction(row.Restriction),
		Type:                bbscontentv1enum.ArticleType(row.Type),
		BountyPoints:        row.BountyPoints,
		AcceptedAnswerId:    row.AcceptedAnswerID,
		Statement:           row.Statement,
		Commentable:         row.Commentable,
		Anonymous:           row.Anonymous,
		ViewCount:           row.ViewCount,
		ThankCount:          row.ThankCount,
		LikeCount:           row.LikeCount,
		CollectCount:        row.CollectCount,
		WatchCount:          row.WatchCount,
		ReplyCount:          row.ReplyCount,
		CoverImageUrl:       row.CoverImageURL,
		ViewerActionState:   s.articleViewerActionState(row.ViewerActionState),
		LastReplyAt:         row.LastReplyAt,
		LastReplyUser:       s.accountProfile(row.LastReplyUser),
		Postscripts:         postscripts,
		AuthorUser:          s.accountProfile(row.AuthorUser),
		CreatedBy:           row.CreatedBy,
		UpdatedBy:           row.UpdatedBy,
		CreatedAt:           row.CreatedAt,
		UpdatedAt:           row.UpdatedAt,
		PublishedAt:         row.PublishedAt,
		EditedAt:            row.EditedAt,
	}
}

func (s *ContentArticleService) articleViewerActionState(row *usecase.ContentArticleViewerActionState) *bbscontentv1.CreateArticle_Resp_ArticleViewerActionState {
	if row == nil {
		return nil
	}
	return &bbscontentv1.CreateArticle_Resp_ArticleViewerActionState{
		Liked:     row.Liked,
		Thanked:   row.Thanked,
		Collected: row.Collected,
		Watched:   row.Watched,
	}
}

func (s *ContentArticleService) articlePostscript(row *usecase.ContentArticlePostscript) *bbscontentv1.CreateArticle_Resp_ArticlePostscript {
	if row == nil {
		return nil
	}
	return &bbscontentv1.CreateArticle_Resp_ArticlePostscript{
		Id:            row.ID,
		ArticleId:     row.ArticleID,
		Content:       row.Content,
		ContentRender: row.ContentRender,
		Restriction:   bbscontentv1enum.ContentRestriction(row.Restriction),
		CreatedBy:     row.CreatedBy,
		UpdatedBy:     row.UpdatedBy,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}
}

func (s *ContentArticleService) accountProfile(row *usecase.ContentAccountProfile) *bbscontentv1.CreateArticle_Resp_AccountProfile {
	if row == nil {
		return nil
	}
	return &bbscontentv1.CreateArticle_Resp_AccountProfile{
		Id:            row.ID,
		Name:          row.Name,
		Nickname:      row.Nickname,
		Url:           row.URL,
		AvatarUrl:     row.AvatarURL,
		Introduction:  row.Introduction,
		Mbti:          bbsuserv1enum.MBTI(row.MBTI),
		Status:        bbsuserv1enum.AccountStatus(row.Status),
		FollowCount:   row.FollowCount,
		FollowerCount: row.FollowerCount,
		CreatedAt:     row.CreatedAt,
		UpdatedAt:     row.UpdatedAt,
	}
}
