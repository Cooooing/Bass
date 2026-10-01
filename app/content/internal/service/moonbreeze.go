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

type MoonbreezeService struct {
	v1.UnimplementedContentMoonbreezeServiceServer
	moonbreezeUsecase *usecase.MoonbreezeUsecase
}

func NewMoonbreezeService(moonbreezeUsecase *usecase.MoonbreezeUsecase) *MoonbreezeService {
	return &MoonbreezeService{moonbreezeUsecase: moonbreezeUsecase}
}

func (s *MoonbreezeService) RegisterGrpc(gs *grpc.Server) {
	v1.RegisterContentMoonbreezeServiceServer(gs, s)
}

func (s *MoonbreezeService) RegisterHttp(hs *http.Server) {}

func (s *MoonbreezeService) Create(ctx context.Context, req *v1.CreateMoonbreeze_Req) (*v1.CreateMoonbreeze_Resp, error) {
	if req.GetAuthorId() <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_MOONBREEZE_INVALID)
	}
	row, err := s.moonbreezeUsecase.Create(ctx, &usecase.CreateMoonbreezeReq{
		Content:  req.GetContent(),
		AuthorID: req.GetAuthorId(),
		City:     req.City,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateMoonbreeze_Resp{Moonbreeze: s.moonbreezeReply(row)}, nil
}

func (s *MoonbreezeService) Page(ctx context.Context, req *v1.PageMoonbreezes_Req) (*v1.PageMoonbreezes_Resp, error) {
	if req.GetSize() > 50 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_CONTENT_MOONBREEZE_INVALID)
	}
	page, err := s.moonbreezeUsecase.Page(ctx, &usecase.PageMoonbreezesReq{
		AuthorIDs: req.GetAuthorIds(),
		Cursor:    req.Cursor,
		Size:      int(req.GetSize()),
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*v1.Moonbreeze, 0, len(page.Rows))
	for _, row := range page.Rows {
		rows = append(rows, s.moonbreezeReply(row))
	}
	return &v1.PageMoonbreezes_Resp{Rows: rows, NextCursor: page.NextCursor}, nil
}

func (*MoonbreezeService) moonbreezeReply(row *model.Moonbreeze) *v1.Moonbreeze {
	if row == nil {
		return nil
	}
	reply := &v1.Moonbreeze{
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
