package repo

import (
	bbscontentv1 "common/api/gen/bbs/v1/content"
	"context"
)

type ContentRepo interface {
	CreateArticle(ctx context.Context, req *bbscontentv1.CreateArticle_Request) (*bbscontentv1.CreateArticle_Reply, error)
	UpdateDraftArticle(ctx context.Context, req *bbscontentv1.UpdateDraftArticle_Request) (*bbscontentv1.UpdateDraftArticle_Reply, error)
	PublishArticle(ctx context.Context, req *bbscontentv1.PublishArticle_Request) (*bbscontentv1.PublishArticle_Reply, error)
	DeleteArticle(ctx context.Context, req *bbscontentv1.DeleteArticle_Request) (*bbscontentv1.DeleteArticle_Reply, error)
	ListArticles(ctx context.Context, req *bbscontentv1.ListArticles_Request) (*bbscontentv1.ListArticles_Reply, error)
	GetArticle(ctx context.Context, req *bbscontentv1.GetArticle_Request) (*bbscontentv1.GetArticle_Reply, error)
	AddPostscript(ctx context.Context, req *bbscontentv1.AddPostscript_Request) (*bbscontentv1.AddPostscript_Reply, error)
	LikeArticle(ctx context.Context, req *bbscontentv1.LikeArticle_Request) (*bbscontentv1.LikeArticle_Reply, error)
	ThankArticle(ctx context.Context, req *bbscontentv1.ThankArticle_Request) (*bbscontentv1.ThankArticle_Reply, error)
	CollectArticle(ctx context.Context, req *bbscontentv1.CollectArticle_Request) (*bbscontentv1.CollectArticle_Reply, error)
	WatchArticle(ctx context.Context, req *bbscontentv1.WatchArticle_Request) (*bbscontentv1.WatchArticle_Reply, error)
	RewardArticle(ctx context.Context, req *bbscontentv1.RewardArticle_Request) (*bbscontentv1.RewardArticle_Reply, error)
	AcceptAnswerArticle(ctx context.Context, req *bbscontentv1.AcceptAnswerArticle_Request) (*bbscontentv1.AcceptAnswerArticle_Reply, error)
	CreateComment(ctx context.Context, req *bbscontentv1.CreateComment_Request) (*bbscontentv1.CreateComment_Reply, error)
	ListComments(ctx context.Context, req *bbscontentv1.ListComments_Request) (*bbscontentv1.ListComments_Reply, error)
	LikeComment(ctx context.Context, req *bbscontentv1.LikeComment_Request) (*bbscontentv1.LikeComment_Reply, error)
	ThankComment(ctx context.Context, req *bbscontentv1.ThankComment_Request) (*bbscontentv1.ThankComment_Reply, error)
	ListDomains(ctx context.Context, req *bbscontentv1.ListDomains_Request) (*bbscontentv1.ListDomains_Reply, error)
	ListTags(ctx context.Context, req *bbscontentv1.ListTags_Request) (*bbscontentv1.ListTags_Reply, error)
}
