package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/content/v1"
	"content/internal/biz/usecase"
	"content/internal/enum"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostscriptService struct {
	v1.UnimplementedContentPostscriptServiceServer

	postscriptUsecase *usecase.PostscriptUsecase
}

func NewPostscriptService(
	postscriptUsecase *usecase.PostscriptUsecase,
) *PostscriptService {
	return &PostscriptService{
		postscriptUsecase: postscriptUsecase,
	}
}

func (s *PostscriptService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentPostscriptServiceServer(gs, s)
}

func (s *PostscriptService) RegisterHttp(hs *http.Server) {
}

func (s *PostscriptService) Add(ctx context.Context, req *v1.AddPostscript_Req) (*v1.AddPostscript_Resp, error) {
	if req.GetArticleId() <= 0 || req.GetUserId() <= 0 || req.GetContent() == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row, err := s.postscriptUsecase.Add(ctx, &usecase.PostscriptAddReq{ArticleID: req.GetArticleId(), Content: req.GetContent(), UserID: req.GetUserId()})
	if err != nil {
		return nil, err
	}
	out := &v1.Postscript{Id: row.ID, ArticleId: row.ArticleID, Content: row.Content, Restriction: enum.ContentRestrictionMap.MustToProto(row.Restriction), CreatedBy: row.CreatedBy, UpdatedBy: row.UpdatedBy}
	if row.CreatedAt != nil {
		out.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		out.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return &v1.AddPostscript_Resp{Postscript: out}, nil
}

func (s *PostscriptService) List(ctx context.Context, req *v1.ListPostscripts_Req) (*v1.ListPostscripts_Resp, error) {
	if req.GetArticleId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	rows, err := s.postscriptUsecase.List(ctx, req.GetArticleId())
	if err != nil {
		return nil, err
	}
	out := make([]*v1.Postscript, 0, len(rows))
	for _, row := range rows {
		item := &v1.Postscript{Id: row.ID, ArticleId: row.ArticleID, Content: row.Content, Restriction: enum.ContentRestrictionMap.MustToProto(row.Restriction), CreatedBy: row.CreatedBy, UpdatedBy: row.UpdatedBy}
		if row.CreatedAt != nil {
			item.CreatedAt = timestamppb.New(*row.CreatedAt)
		}
		if row.UpdatedAt != nil {
			item.UpdatedAt = timestamppb.New(*row.UpdatedAt)
		}
		out = append(out, item)
	}
	return &v1.ListPostscripts_Resp{Rows: out}, nil
}
