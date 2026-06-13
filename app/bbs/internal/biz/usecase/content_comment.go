package usecase

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	"context"
)

type ContentCommentUsecase struct {
	contentCommentClient repo.ContentCommentClient
}

func NewContentCommentUsecase(contentCommentClient repo.ContentCommentClient) *ContentCommentUsecase {
	return &ContentCommentUsecase{contentCommentClient: contentCommentClient}
}

func (u *ContentCommentUsecase) CreateComment(ctx context.Context, req *bbscontentv1.CreateComment_Request) (*bbscontentv1.CreateComment_Reply, error) {
	return u.contentCommentClient.CreateComment(ctx, req)
}

func (u *ContentCommentUsecase) ListComments(ctx context.Context, req *bbscontentv1.ListComments_Request) (*bbscontentv1.ListComments_Reply, error) {
	if req == nil {
		req = &bbscontentv1.ListComments_Request{}
	}
	if req.Query == nil {
		req.Query = &bbscontentv1.CommentQuery{}
	}
	req.Query.Restriction = nil
	req.Query.Restrictions = []bbscontentv1.ContentRestriction{
		bbscontentv1.ContentRestriction_CONTENT_RESTRICTION_NONE,
		bbscontentv1.ContentRestriction_CONTENT_RESTRICTION_LOCKED,
	}
	return u.contentCommentClient.ListComments(ctx, req)
}

func (u *ContentCommentUsecase) ListCommentThreads(ctx context.Context, req *bbscontentv1.ListCommentThreads_Request) (*bbscontentv1.ListCommentThreads_Reply, error) {
	return u.contentCommentClient.ListCommentThreads(ctx, req)
}

func (u *ContentCommentUsecase) ListCommentReplies(ctx context.Context, req *bbscontentv1.ListCommentReplies_Request) (*bbscontentv1.ListCommentReplies_Reply, error) {
	return u.contentCommentClient.ListCommentReplies(ctx, req)
}

func (u *ContentCommentUsecase) ListCommentTimeline(ctx context.Context, req *bbscontentv1.ListCommentTimeline_Request) (*bbscontentv1.ListCommentTimeline_Reply, error) {
	return u.contentCommentClient.ListCommentTimeline(ctx, req)
}

func (u *ContentCommentUsecase) LikeComment(ctx context.Context, req *bbscontentv1.LikeComment_Request) (*bbscontentv1.LikeComment_Reply, error) {
	return u.contentCommentClient.LikeComment(ctx, req)
}

func (u *ContentCommentUsecase) ThankComment(ctx context.Context, req *bbscontentv1.ThankComment_Request) (*bbscontentv1.ThankComment_Reply, error) {
	return u.contentCommentClient.ThankComment(ctx, req)
}
