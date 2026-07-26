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
	var postscript *bbscontentv1.AddPostscript_Resp_ArticlePostscript
	if resp != nil {
		postscript = &bbscontentv1.AddPostscript_Resp_ArticlePostscript{
			Id:            resp.ID,
			ArticleId:     resp.ArticleID,
			Content:       resp.Content,
			ContentRender: resp.ContentRender,
			Restriction:   bbscontentv1enum.ContentRestriction(resp.Restriction),
			CreatedBy:     resp.CreatedBy,
			UpdatedBy:     resp.UpdatedBy,
			CreatedAt:     timestamppb.New(*resp.CreatedAt),
			UpdatedAt:     timestamppb.New(*resp.UpdatedAt),
		}
	}
	return &bbscontentv1.AddPostscript_Resp{
		Postscript: postscript,
	}, nil
}
