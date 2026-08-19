package service

import (
	"bbs/internal/biz/usecase"
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bbscontentv1 "common/proto/gen/bbs/v1/content"
	bbscontentv1enum "common/proto/gen/bbs/v1/content/enum"
	cerrors "common/proto/gen/common/errors"
	"context"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ContentPostscriptService struct {
	bbscontentv1.UnimplementedPostscriptServiceServer
	contentPostscriptUsecase *usecase.ContentPostscriptUsecase
}

func NewContentPostscriptService(
	contentPostscriptUsecase *usecase.ContentPostscriptUsecase,
) *ContentPostscriptService {
	return &ContentPostscriptService{
		contentPostscriptUsecase: contentPostscriptUsecase,
	}
}

func (s *ContentPostscriptService) RegisterGrpc(gs *grpc.Server) {
}

func (s *ContentPostscriptService) RegisterHttp(hs *http.Server) {
	bbscontentv1.RegisterPostscriptServiceHTTPServer(hs, s)
}

func (s *ContentPostscriptService) Add(ctx context.Context, req *bbscontentv1.AddPostscript_Req) (*bbscontentv1.AddPostscript_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	resp, err := s.contentPostscriptUsecase.AddPostscript(ctx, &usecase.AddPostscriptReq{
		UserID:    user.ID,
		ArticleID: req.GetArticleId(),
		Content:   req.GetContent(),
	})
	if err != nil {
		return nil, err
	}
	postscript := &bbscontentv1.ArticlePostscript{
		Id:            resp.ID,
		ArticleId:     resp.ArticleID,
		Content:       resp.Content,
		ContentRender: resp.ContentRender,
		Restriction:   bbscontentv1enum.ContentRestriction(resp.Restriction),
		CreatedBy:     resp.CreatedBy,
		UpdatedBy:     resp.UpdatedBy,
	}
	if resp.CreatedAt != nil {
		postscript.CreatedAt = timestamppb.New(*resp.CreatedAt)
	}
	if resp.UpdatedAt != nil {
		postscript.UpdatedAt = timestamppb.New(*resp.UpdatedAt)
	}
	return &bbscontentv1.AddPostscript_Resp{
		Postscript: postscript,
	}, nil
}

func (s *ContentPostscriptService) List(ctx context.Context, req *bbscontentv1.ListPostscripts_Req) (*bbscontentv1.ListPostscripts_Resp, error) {
	userID := int64(0)
	if user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo); ok && user != nil {
		userID = user.ID
	}
	resp, err := s.contentPostscriptUsecase.ListPostscripts(ctx, &usecase.ListPostscriptsReq{
		ArticleID: req.GetArticleId(),
		UserID:    userID,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbscontentv1.ArticlePostscript, 0, len(resp))
	for _, item := range resp {
		row := &bbscontentv1.ArticlePostscript{
			Id:            item.ID,
			ArticleId:     item.ArticleID,
			Content:       item.Content,
			ContentRender: item.ContentRender,
			Restriction:   bbscontentv1enum.ContentRestriction(item.Restriction),
			CreatedBy:     item.CreatedBy,
			UpdatedBy:     item.UpdatedBy,
		}
		if item.CreatedAt != nil {
			row.CreatedAt = timestamppb.New(*item.CreatedAt)
		}
		if item.UpdatedAt != nil {
			row.UpdatedAt = timestamppb.New(*item.UpdatedAt)
		}
		rows = append(rows, row)
	}
	return &bbscontentv1.ListPostscripts_Resp{
		Rows: rows,
	}, nil
}
