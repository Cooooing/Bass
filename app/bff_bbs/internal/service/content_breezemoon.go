package service

import (
	"bff_bbs/internal/biz/repo"
	"bff_bbs/internal/biz/usecase"
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/server"
	"common/pkg/util"
	bbscontentv1 "common/proto/gen/bff_bbs/v1/content"
	bbsuserv1enum "common/proto/gen/bff_bbs/v1/user/enum"
	cerrors "common/proto/gen/common/errors"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ContentBreezemoonService struct {
	bbscontentv1.UnimplementedBreezemoonServiceServer
	breezemoonUsecase *usecase.ContentBreezemoonUsecase
}

func NewContentBreezemoonService(breezemoonUsecase *usecase.ContentBreezemoonUsecase) *ContentBreezemoonService {
	return &ContentBreezemoonService{
		breezemoonUsecase: breezemoonUsecase,
	}
}

func (s *ContentBreezemoonService) RegisterGrpc(gs *grpc.Server) {}

func (s *ContentBreezemoonService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterBreezemoonServiceHTTPServer(hs, s)
}

func (s *ContentBreezemoonService) Create(
	ctx context.Context,
	req *bbscontentv1.CreateBreezemoon_Req,
) (*bbscontentv1.CreateBreezemoon_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	row, err := s.breezemoonUsecase.Create(ctx, &usecase.CreateBreezemoonReq{
		UserID:  user.ID,
		Content: req.GetContent(),
		IP:      server.ClientIP(ctx),
	})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.CreateBreezemoon_Resp{Breezemoon: s.breezemoonView(row)}, nil
}

func (s *ContentBreezemoonService) PagePublic(
	ctx context.Context,
	req *bbscontentv1.PagePublicBreezemoons_Req,
) (*bbscontentv1.PageBreezemoons_Resp, error) {
	page, err := s.breezemoonUsecase.Page(ctx, &usecase.PageBreezemoonsReq{
		ViewerID: s.currentUserID(ctx),
		Cursor:   req.Cursor,
		Size:     req.GetSize(),
	})
	if err != nil {
		return nil, err
	}
	return s.pageReply(page), nil
}

func (s *ContentBreezemoonService) PageWatching(
	ctx context.Context,
	req *bbscontentv1.PageWatchingBreezemoons_Req,
) (*bbscontentv1.PageBreezemoons_Resp, error) {
	viewerID := s.currentUserID(ctx)
	page, err := s.breezemoonUsecase.PageWatching(ctx, viewerID, req.Cursor, req.GetSize())
	if err != nil {
		return nil, err
	}
	return s.pageReply(page), nil
}

func (s *ContentBreezemoonService) PageMember(
	ctx context.Context,
	req *bbscontentv1.PageMemberBreezemoons_Req,
) (*bbscontentv1.PageBreezemoons_Resp, error) {
	viewerID := s.currentUserID(ctx)
	page, err := s.breezemoonUsecase.PageMember(ctx, viewerID, req.GetName(), req.Cursor, req.GetSize())
	if err != nil {
		return nil, err
	}
	return s.pageReply(page), nil
}

func (s *ContentBreezemoonService) pageReply(page *repo.PageBreezemoonsResp) *bbscontentv1.PageBreezemoons_Resp {
	rows := make([]*bbscontentv1.Breezemoon, 0, len(page.Rows))
	for _, row := range page.Rows {
		rows = append(rows, s.breezemoonView(row))
	}
	return &bbscontentv1.PageBreezemoons_Resp{Rows: rows, NextCursor: page.NextCursor}
}

func (*ContentBreezemoonService) breezemoonView(row *repo.Breezemoon) *bbscontentv1.Breezemoon {
	if row == nil {
		return nil
	}
	view := &bbscontentv1.Breezemoon{
		Id:      row.ID,
		Content: row.Content,
	}
	if author := row.Author; author != nil {
		view.Author = &bbscontentv1.AccountProfile{
			Id:        author.ID,
			Name:      author.Name,
			Nickname:  author.Nickname,
			AvatarUrl: author.AvatarURL,
			Status:    bbsuserv1enum.AccountStatus(author.Status),
		}
	}
	if row.CreatedAt != nil {
		view.CreatedAt = timestamppb.New(*row.CreatedAt)
	}
	if row.ShowCity {
		view.City = row.City
	}
	return view
}

func (*ContentBreezemoonService) currentUserID(ctx context.Context) int64 {
	if user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo); ok && user != nil {
		return user.ID
	}
	return 0
}
