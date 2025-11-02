package service

import (
	v1 "common/api/content/v1"
	"common/pkg/util/base"
	"content/internal/biz"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type CommentService struct {
	v1.UnimplementedContentCommentServiceServer
	*BaseService

	commentDomain *biz.CommentDomain
	commentRepo   repo.CommentRepo
	articleRepo   repo.ArticleRepo
}

func (s *CommentService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentCommentServiceServer(gs, s)
}

func (s *CommentService) RegisterHttp(hs *http.Server) {
	v1.RegisterContentCommentServiceHTTPServer(hs, s)
}

func NewCommentService(baseService *BaseService, commentDomain *biz.CommentDomain, commentRepo repo.CommentRepo, articleRepo repo.ArticleRepo) *CommentService {
	return &CommentService{
		BaseService:   baseService,
		commentDomain: commentDomain,
		commentRepo:   commentRepo,
		articleRepo:   articleRepo,
	}
}

func (s *CommentService) Add(ctx context.Context, req *v1.AddCommentRequest) (rsp *v1.AddCommentReply, err error) {
	user := s.tokenRepo.GetUserInfo(ctx)
	_, err = s.commentDomain.Add(ctx, &model.Comment{
		ArticleID: req.ArticleId,
		UserID:    user.ID,
		Content:   req.Content,
		ReplyID:   base.If(req.ReplyId != 0, &req.ReplyId, nil),
	})
	return &v1.AddCommentReply{}, err
}

func (s *CommentService) Get(ctx context.Context, req *v1.GetCommentRequest) (rsp *v1.GetCommentReply, err error) {
	return s.commentDomain.Get(ctx, req)
}

func (s *CommentService) Like(ctx context.Context, req *v1.LikeCommentRequest) (rsp *v1.LikeCommentReply, err error) {
	// TODO implement me
	panic("implement me")
}

func (s *CommentService) Thank(ctx context.Context, req *v1.ThankCommentRequest) (rsp *v1.ThankCommentReply, err error) {
	// TODO implement me
	panic("implement me")
}

func (s *CommentService) UpdateStatus(ctx context.Context, req *v1.UpdateStatusCommentRequest) (rsp *v1.UpdateStatusCommentReply, err error) {
	// TODO implement me
	panic("implement me")
}
