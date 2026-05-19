package service

import (
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/content/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	"common/pkg/util"

	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/biz/usecase"
	"content/internal/data/gen"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type CommentService struct {
	v1.UnimplementedContentCommentServiceServer

	commentDomain *usecase.CommentUsecase
	commentRepo   repo.CommentRepo
	articleRepo   repo.ArticleRepo
	db            *gen.Client
}

func (s *CommentService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentCommentServiceServer(gs, s)
}

func NewCommentService(
	commentDomain *usecase.CommentUsecase,
	commentRepo repo.CommentRepo,
	articleRepo repo.ArticleRepo,
	db *gen.Client,
) *CommentService {
	return &CommentService{
		commentDomain: commentDomain,
		commentRepo:   commentRepo,
		articleRepo:   articleRepo,
		db:            db,
	}
}

func (s *CommentService) AddComment(ctx context.Context, req *v1.AddComment_Request) (rsp *v1.AddComment_Reply, err error) {
	user, ok := util.GetContextValue[*commonModel.User](ctx, constant.CtxUserInfo)
	if !ok {
		return nil, cerrors.ErrorUnauthorized("user not login")
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
	exist, err := s.commentRepo.Exist(ctx, s.db, &repo.CommentGetReq{CommentId: new(req.Id)})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, cerrors.ErrorBadRequest("comment not exist")
	}

	err = s.commentRepo.UpdateStat(ctx, s.db, req.Id, v1.CommentAction_COMMENT_ACTION_LIKE, util.If[int32](req.Active, 1, -1))
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
