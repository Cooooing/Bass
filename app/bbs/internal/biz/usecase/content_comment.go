package usecase

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/api/gen/bbs/v1/content"
	"context"
)

type ContentCommentUsecase struct {
	contentCommentRepo repo.ContentCommentRepo
}

func NewContentCommentUsecase(contentCommentRepo repo.ContentCommentRepo) *ContentCommentUsecase {
	return &ContentCommentUsecase{contentCommentRepo: contentCommentRepo}
}

func (u *ContentCommentUsecase) CreateComment(ctx context.Context, req *bbscontentv1.CreateComment_Request) (*bbscontentv1.CreateComment_Reply, error) {
	return u.contentCommentRepo.CreateComment(ctx, req)
}

func (u *ContentCommentUsecase) ListComments(ctx context.Context, req *bbscontentv1.ListComments_Request) (*bbscontentv1.ListComments_Reply, error) {
	return u.contentCommentRepo.ListComments(ctx, req)
}

func (u *ContentCommentUsecase) LikeComment(ctx context.Context, req *bbscontentv1.LikeComment_Request) (*bbscontentv1.LikeComment_Reply, error) {
	return u.contentCommentRepo.LikeComment(ctx, req)
}

func (u *ContentCommentUsecase) ThankComment(ctx context.Context, req *bbscontentv1.ThankComment_Request) (*bbscontentv1.ThankComment_Reply, error) {
	return u.contentCommentRepo.ThankComment(ctx, req)
}
