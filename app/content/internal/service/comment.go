package service

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	"common/pkg/util"
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
	user := util.MustGetUserInfo(ctx)
	_, err = s.commentDomain.Add(ctx, &model.Comment{
		ArticleID: req.ArticleId,
		CreatedBy: &user.ID,
		Content:   req.Content,
		ReplyID:   base.If(req.ReplyId != 0, &req.ReplyId, nil),
	})
	return &v1.AddCommentReply{}, err
}

func (s *CommentService) Page(ctx context.Context, req *v1.PageCommentRequest) (*v1.PageCommentReply, error) {
	return s.commentDomain.Page(ctx, req.Page, &repo.CommentGetReq{
		CommentId: req.CommentId,
		ArticleId: req.ArticleId,
		UserId:    req.UserId,
		Order:     req.Order,
	})
}

func (s *CommentService) Like(ctx context.Context, req *v1.LikeCommentRequest) (rsp *v1.LikeCommentReply, err error) {
	// user := s.tokenRepo.GetUserInfo(ctx)
	exist, err := s.commentRepo.Exist(ctx, s.db, req.Id)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, cv1.ErrorBadRequest("comment not exist")
	}

	err = s.commentRepo.UpdateStat(ctx, s.db, req.Id, cv1.CommentAction_CommentActionLike, base.If[int32](req.Active, 1, -1))
	return &v1.LikeCommentReply{}, err
}

func (s *CommentService) Thank(ctx context.Context, req *v1.ThankCommentRequest) (rsp *v1.ThankCommentReply, err error) {
	// TODO implement me
	panic("implement me")
}

func (s *CommentService) UpdateStatus(ctx context.Context, req *v1.UpdateStatusCommentRequest) (rsp *v1.UpdateStatusCommentReply, err error) {
	// TODO implement me
	panic("implement me")
}
