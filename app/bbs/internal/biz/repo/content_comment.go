package repo

import (
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"context"
)

type ContentCommentClient interface {
	CreateComment(ctx context.Context, req *bbscontentv1.CreateComment_Request) (*bbscontentv1.CreateComment_Reply, error)
	ListComments(ctx context.Context, req *bbscontentv1.ListComments_Request) (*bbscontentv1.ListComments_Reply, error)
	ListCommentThreads(ctx context.Context, req *bbscontentv1.ListCommentThreads_Request) (*bbscontentv1.ListCommentThreads_Reply, error)
	ListCommentReplies(ctx context.Context, req *bbscontentv1.ListCommentReplies_Request) (*bbscontentv1.ListCommentReplies_Reply, error)
	ListCommentTimeline(ctx context.Context, req *bbscontentv1.ListCommentTimeline_Request) (*bbscontentv1.ListCommentTimeline_Reply, error)
	LikeComment(ctx context.Context, req *bbscontentv1.LikeComment_Request) (*bbscontentv1.LikeComment_Reply, error)
	ThankComment(ctx context.Context, req *bbscontentv1.ThankComment_Request) (*bbscontentv1.ThankComment_Reply, error)
}
