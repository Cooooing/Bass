package service

import (
	cv1 "common/api/gen/common/v1"
	v1 "common/api/gen/content/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"

	"content/internal/biz/domain"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

type CommentService struct {
	v1.UnimplementedContentCommentServiceServer
	*BaseService

	commentDomain *domain.CommentDomain
	commentRepo   repo.CommentRepo
	articleRepo   repo.ArticleRepo
}

func (s *CommentService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentCommentServiceServer(gs, s)
}

func (s *CommentService) RegisterHttp(hs *http.Server) {
	v1.RegisterContentCommentServiceHTTPServer(hs, s)
}

func NewCommentService(baseService *BaseService, commentDomain *domain.CommentDomain, commentRepo repo.CommentRepo, articleRepo repo.ArticleRepo) *CommentService {
	return &CommentService{
		BaseService:   baseService,
		commentDomain: commentDomain,
		commentRepo:   commentRepo,
		articleRepo:   articleRepo,
	}
}

func (s *CommentService) AddComment(ctx context.Context, req *v1.AddComment_Request) (rsp *v1.AddComment_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cv1.ErrorUnauthorized("user not login")
	}
	comment, err := s.commentDomain.Add(ctx, &model.Comment{
		Comment: &gen.Comment{
			ArticleID: req.ArticleId,
			CreatedBy: &user.ID,
			Content:   req.Content,
			ReplyID:   util.If(req.ReplyId != 0, &req.ReplyId, nil),
		},
	})
	if err != nil {
		return nil, err
	}
	return &v1.AddComment_Reply{
		Comment: comment.ConvertToRpc(),
	}, err
}

func (s *CommentService) Page(ctx context.Context, req *v1.PageComment_Request) (*v1.PageComment_Reply, error) {
	req.Query = util.OrDefault(req.Query, &v1.CommentQueryParams{})
	page, comments, err := s.commentDomain.Page(ctx, req.Page, &repo.CommentGetReq{
		CommentId:   req.Query.CommentId,
		CommentIds:  nil,
		ParentId:    req.Query.ParentId,
		ReplyId:     req.Query.ReplyId,
		ArticleId:   req.Query.ArticleId,
		ArticleIds:  nil,
		CreatedBy:   req.Query.UserId,
		Status:      new(v1.CommentStatus_COMMENT_STATUS_NORMAL),
		Level:       req.Query.Level,
		Order:       (*int32)(req.Query.Order),
		WithArticle: req.Query.WithArticle,
	})
	return &v1.PageComment_Reply{
		Page: page,
		Rows: commonModel.ConvertToRpcList(comments),
	}, err
}

func (s *CommentService) Like(ctx context.Context, req *v1.LikeComment_Request) (rsp *v1.LikeComment_Reply, err error) {
	// user := s.tokenCache.GetUserInfo(ctx)
	exist, err := s.commentRepo.Exist(ctx, s.Db, &repo.CommentGetReq{CommentId: new(req.Id)})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, cv1.ErrorBadRequest("comment not exist")
	}

	err = s.commentRepo.UpdateStat(ctx, s.Db, req.Id, v1.CommentAction_COMMENT_ACTION_LIKE, util.If[int32](req.Active, 1, -1))
	return &v1.LikeComment_Reply{}, err
}

func (s *CommentService) Thank(ctx context.Context, req *v1.ThankComment_Request) (rsp *v1.ThankComment_Reply, err error) {
	// TODO implement me
	panic("implement me")
}

func (s *CommentService) UpdateStatus(ctx context.Context, req *v1.UpdateStatusComment_Request) (rsp *v1.UpdateStatusComment_Reply, err error) {
	// TODO implement me
	panic("implement me")
}
