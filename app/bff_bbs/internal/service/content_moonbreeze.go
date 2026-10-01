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

type ContentMoonbreezeService struct {
	bbscontentv1.UnimplementedMoonbreezeServiceServer
	moonbreezeUsecase *usecase.ContentMoonbreezeUsecase
}

func NewContentMoonbreezeService(moonbreezeUsecase *usecase.ContentMoonbreezeUsecase) *ContentMoonbreezeService {
	return &ContentMoonbreezeService{
		moonbreezeUsecase: moonbreezeUsecase,
	}
}

func (s *ContentMoonbreezeService) RegisterGrpc(gs *grpc.Server) {}

func (s *ContentMoonbreezeService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterMoonbreezeServiceHTTPServer(hs, s)
}

func (s *ContentMoonbreezeService) Create(
	ctx context.Context,
	req *bbscontentv1.CreateMoonbreeze_Req,
) (*bbscontentv1.CreateMoonbreeze_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	row, err := s.moonbreezeUsecase.Create(ctx, &usecase.CreateMoonbreezeReq{
		UserID:  user.ID,
		Content: req.GetContent(),
		IP:      server.ClientIP(ctx),
	})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.CreateMoonbreeze_Resp{Moonbreeze: s.moonbreezeView(row)}, nil
}

func (s *ContentMoonbreezeService) PagePublic(
	ctx context.Context,
	req *bbscontentv1.PagePublicMoonbreezes_Req,
) (*bbscontentv1.PageMoonbreezes_Resp, error) {
	page, err := s.moonbreezeUsecase.Page(ctx, &usecase.PageMoonbreezesReq{
		ViewerID: s.currentUserID(ctx),
		Cursor:   req.Cursor,
		Size:     req.GetSize(),
	})
	if err != nil {
		return nil, err
	}
	return s.pageReply(page), nil
}

func (s *ContentMoonbreezeService) PageWatching(
	ctx context.Context,
	req *bbscontentv1.PageWatchingMoonbreezes_Req,
) (*bbscontentv1.PageMoonbreezes_Resp, error) {
	viewerID := s.currentUserID(ctx)
	page, err := s.moonbreezeUsecase.PageWatching(ctx, viewerID, req.Cursor, req.GetSize())
	if err != nil {
		return nil, err
	}
	return s.pageReply(page), nil
}

func (s *ContentMoonbreezeService) PageMember(
	ctx context.Context,
	req *bbscontentv1.PageMemberMoonbreezes_Req,
) (*bbscontentv1.PageMoonbreezes_Resp, error) {
	viewerID := s.currentUserID(ctx)
	page, err := s.moonbreezeUsecase.PageMember(ctx, viewerID, req.GetName(), req.Cursor, req.GetSize())
	if err != nil {
		return nil, err
	}
	return s.pageReply(page), nil
}

func (s *ContentMoonbreezeService) pageReply(page *repo.PageMoonbreezesResp) *bbscontentv1.PageMoonbreezes_Resp {
	rows := make([]*bbscontentv1.Moonbreeze, 0, len(page.Rows))
	for _, row := range page.Rows {
		rows = append(rows, s.moonbreezeView(row))
	}
	return &bbscontentv1.PageMoonbreezes_Resp{Rows: rows, NextCursor: page.NextCursor}
}

func (*ContentMoonbreezeService) moonbreezeView(row *repo.Moonbreeze) *bbscontentv1.Moonbreeze {
	if row == nil {
		return nil
	}
	view := &bbscontentv1.Moonbreeze{
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

func (*ContentMoonbreezeService) currentUserID(ctx context.Context) int64 {
	if user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo); ok && user != nil {
		return user.ID
	}
	return 0
}
