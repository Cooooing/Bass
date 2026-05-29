package service

import (
	cerrors "common/api/gen/common/errors"
	v1 "common/api/gen/content/v1"
	"common/pkg/util"

	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/biz/usecase"
	"content/internal/enum"
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CommentService struct {
	v1.UnimplementedContentCommentServiceServer

	commentUsecase *usecase.CommentUsecase
}

func (s *CommentService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentCommentServiceServer(gs, s)
}

func NewCommentService(
	commentUsecase *usecase.CommentUsecase,
) *CommentService {
	return &CommentService{
		commentUsecase: commentUsecase,
	}
}

func (s *CommentService) Create(ctx context.Context, req *v1.CreateComment_Request) (rsp *v1.CreateComment_Reply, err error) {
	comment, err := s.commentUsecase.Add(ctx, req.UserId, &model.Comment{
		ArticleID: req.ArticleId,
		Content:   req.Content,
		ReplyID:   util.If(req.ReplyId != 0, &req.ReplyId, nil),
		CreatedBy: new(req.UserId),
		UpdatedBy: new(req.UserId),
	})
	if err != nil {
		return nil, err
	}
	comment.ParseContent()
	reply := &v1.Comment{
		CreatedAt:     timestamppb.New(*comment.CreatedAt),
		UpdatedAt:     timestamppb.New(*comment.UpdatedAt),
		CreatedBy:     comment.CreatedBy,
		UpdatedBy:     comment.UpdatedBy,
		Id:            comment.ID,
		ArticleId:     comment.ArticleID,
		Content:       comment.Content,
		ContentRender: comment.ContentRender,
		Level:         comment.Level,
		ParentId:      comment.ParentID,
		ReplyId:       comment.ReplyID,
		Status:        enum.CommentStatusMap.MustToProto(enum.CommentStatus(comment.Status)),
		ThankCount:    comment.ThankCount,
		LikeCount:     comment.LikeCount,
		ReplyCount:    comment.ReplyCount,
		User:          comment.User,
		ReplyUser:     comment.ReplyUser,
	}
	return &v1.CreateComment_Reply{
		Comment: reply,
	}, err
}

func (s *CommentService) List(ctx context.Context, req *v1.ListComments_Request) (*v1.ListComments_Reply, error) {
	req.Query = util.OrDefault(req.Query, &v1.CommentQueryParams{})
	if req.Query.Status != nil {
		if _, ok := enum.CommentStatusMap.ToEnum(*req.Query.Status); !ok {
			return nil, cerrors.ErrorBadRequest("invalid comment status")
		}
	}
	page, comments, err := s.commentUsecase.Page(ctx, req.Page, &repo.CommentGetReq{
		CommentId:  req.Query.CommentId,
		CommentIds: nil,
		ParentId:   req.Query.ParentId,
		ReplyId:    req.Query.ReplyId,
		ArticleId:  req.Query.ArticleId,
		ArticleIds: nil,
		CreatedBy:  req.Query.UserId,
		Status:     req.Query.Status,
		Level:      req.Query.Level,
		Order:      req.Query.Order,
	})
	rows := make([]*v1.Comment, 0, len(comments))
	for _, comment := range comments {
		comment.ParseContent()
		row := &v1.Comment{
			CreatedAt:     timestamppb.New(*comment.CreatedAt),
			UpdatedAt:     timestamppb.New(*comment.UpdatedAt),
			CreatedBy:     comment.CreatedBy,
			UpdatedBy:     comment.UpdatedBy,
			Id:            comment.ID,
			ArticleId:     comment.ArticleID,
			Content:       comment.Content,
			ContentRender: comment.ContentRender,
			Level:         comment.Level,
			ParentId:      comment.ParentID,
			ReplyId:       comment.ReplyID,
			Status:        enum.CommentStatusMap.MustToProto(enum.CommentStatus(comment.Status)),
			ThankCount:    comment.ThankCount,
			LikeCount:     comment.LikeCount,
			ReplyCount:    comment.ReplyCount,
			User:          comment.User,
			ReplyUser:     comment.ReplyUser,
		}
		rows = append(rows, row)
	}
	return &v1.ListComments_Reply{
		Page: page,
		Rows: rows,
	}, err
}

func (s *CommentService) Like(ctx context.Context, req *v1.LikeComment_Request) (rsp *v1.LikeComment_Reply, err error) {
	err = s.commentUsecase.UpdateStat(ctx, req.Id, req.UserId, v1.CommentAction_COMMENT_ACTION_LIKE, req.Active)
	return &v1.LikeComment_Reply{}, err
}

func (s *CommentService) Thank(ctx context.Context, req *v1.ThankComment_Request) (rsp *v1.ThankComment_Reply, err error) {
	err = s.commentUsecase.UpdateStat(ctx, req.Id, req.UserId, v1.CommentAction_COMMENT_ACTION_THANK, req.Active)
	return &v1.ThankComment_Reply{}, err
}

func (s *CommentService) UpdateStatus(ctx context.Context, req *v1.UpdateStatusComment_Request) (rsp *v1.UpdateStatusComment_Reply, err error) {
	err = s.commentUsecase.UpdateStatus(ctx, req.Id, req.UserId, req.Status)
	return &v1.UpdateStatusComment_Reply{}, err
}
