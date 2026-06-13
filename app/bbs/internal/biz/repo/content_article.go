package repo

import (
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"context"
)

type ContentArticleClient interface {
	CreateArticle(ctx context.Context, req *bbscontentv1.CreateArticle_Request) (*bbscontentv1.CreateArticle_Reply, error)
	UpdateArticle(ctx context.Context, req *bbscontentv1.UpdateArticle_Request) (*bbscontentv1.UpdateArticle_Reply, error)
	UpdateDraftArticle(ctx context.Context, req *bbscontentv1.UpdateDraftArticle_Request) (*bbscontentv1.UpdateDraftArticle_Reply, error)
	PublishArticle(ctx context.Context, req *bbscontentv1.PublishArticle_Request) (*bbscontentv1.PublishArticle_Reply, error)
	DiscardDraftArticle(ctx context.Context, req *bbscontentv1.DiscardDraftArticle_Request) (*bbscontentv1.DiscardDraftArticle_Reply, error)
	ListArticles(ctx context.Context, req *bbscontentv1.ListArticles_Request) (*bbscontentv1.ListArticles_Reply, error)
	GetArticle(ctx context.Context, req *bbscontentv1.GetArticle_Request) (*bbscontentv1.GetArticle_Reply, error)
	ViewArticle(ctx context.Context, articleID int64) error
	LikeArticle(ctx context.Context, req *bbscontentv1.LikeArticle_Request) (*bbscontentv1.LikeArticle_Reply, error)
	ThankArticle(ctx context.Context, req *bbscontentv1.ThankArticle_Request) (*bbscontentv1.ThankArticle_Reply, error)
	CollectArticle(ctx context.Context, req *bbscontentv1.CollectArticle_Request) (*bbscontentv1.CollectArticle_Reply, error)
	WatchArticle(ctx context.Context, req *bbscontentv1.WatchArticle_Request) (*bbscontentv1.WatchArticle_Reply, error)
	RewardArticle(ctx context.Context, req *bbscontentv1.RewardArticle_Request) (*bbscontentv1.RewardArticle_Reply, error)
	AcceptAnswerArticle(ctx context.Context, req *bbscontentv1.AcceptAnswerArticle_Request) (*bbscontentv1.AcceptAnswerArticle_Reply, error)
}
