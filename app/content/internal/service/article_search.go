package service

import (
	"context"

	v1 "common/proto/gen/content/v1"
	"content/internal/biz/model"
	"content/internal/biz/usecase"
	"content/internal/enum"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ArticleSearchService struct {
	v1.UnimplementedContentArticleSearchServiceServer
	articleSearchUsecase *usecase.ArticleSearchUsecase
}

func NewArticleSearchService(articleSearchUsecase *usecase.ArticleSearchUsecase) *ArticleSearchService {
	return &ArticleSearchService{articleSearchUsecase: articleSearchUsecase}
}

func (s *ArticleSearchService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentArticleSearchServiceServer(gs, s)
}

func (s *ArticleSearchService) RegisterHttp(hs *http.Server) {}

func (s *ArticleSearchService) SearchPublic(
	ctx context.Context,
	req *v1.SearchPublicArticles_Req,
) (*v1.SearchPublicArticles_Resp, error) {
	result, err := s.articleSearchUsecase.SearchPublic(ctx, &usecase.SearchPublicArticlesReq{
		Query:             req.GetQuery(),
		Cursor:            req.Cursor,
		Size:              int(req.GetSize()),
		ExcludedAuthorIDs: req.GetExcludedAuthorIds(),
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.Article, 0, len(result.Rows))
	for _, row := range result.Rows {
		rows = append(rows, s.articleReply(row))
	}
	return &v1.SearchPublicArticles_Resp{Rows: rows, NextCursor: result.NextCursor}, nil
}

func (*ArticleSearchService) articleReply(row *model.Article) *v1.Article {
	if row == nil {
		return nil
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
		PublishedCity: row.PublishedCity,
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
	return article
}
