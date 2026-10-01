package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/content/v1"
	"content/internal/biz/model"
	"content/internal/biz/usecase"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BreezemoonService struct {
	v1.UnimplementedContentBreezemoonServiceServer
	breezemoonUsecase *usecase.BreezemoonUsecase
}

func NewBreezemoonService(breezemoonUsecase *usecase.BreezemoonUsecase) *BreezemoonService {
	return &BreezemoonService{breezemoonUsecase: breezemoonUsecase}
}

func (s *BreezemoonService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentBreezemoonServiceServer(gs, s)
}

func (s *BreezemoonService) RegisterHttp(hs *http.Server) {}

func (s *BreezemoonService) Create(ctx context.Context, req *v1.CreateBreezemoon_Req) (*v1.CreateBreezemoon_Resp, error) {
	if req.GetAuthorId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_BREEZEMOON_INVALID)
	}
	row, err := s.breezemoonUsecase.Create(ctx, &usecase.CreateBreezemoonReq{
		Content:  req.GetContent(),
		AuthorID: req.GetAuthorId(),
		City:     req.City,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateBreezemoon_Resp{Breezemoon: s.breezemoonReply(row)}, nil
}

func (s *BreezemoonService) Page(ctx context.Context, req *v1.PageBreezemoons_Req) (*v1.PageBreezemoons_Resp, error) {
	if req.GetSize() > 50 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_BREEZEMOON_INVALID)
	}
	page, err := s.breezemoonUsecase.Page(ctx, &usecase.PageBreezemoonsReq{
		AuthorIDs: req.GetAuthorIds(),
		Cursor:    req.Cursor,
		Size:      int(req.GetSize()),
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.Breezemoon, 0, len(page.Rows))
	for _, row := range page.Rows {
		rows = append(rows, s.breezemoonReply(row))
	}
	return &v1.PageBreezemoons_Resp{Rows: rows, NextCursor: page.NextCursor}, nil
}

func (*BreezemoonService) breezemoonReply(row *model.Breezemoon) *v1.Breezemoon {
	if row == nil {
		return nil
	}
	reply := &v1.Breezemoon{
		Id:       row.ID,
		Content:  row.Content,
		AuthorId: row.AuthorID,
		City:     row.City,
	}
	if row.CreatedAt != nil {
		reply.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.UpdatedAt != nil {
		reply.UpdatedAt = timestamppb.New(*row.UpdatedAt)
	}
	return reply
}
