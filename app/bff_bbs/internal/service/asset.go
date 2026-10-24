package service

import (
	"common/pkg/apperror"
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	bffbbsv1 "common/proto/gen/bff_bbs/v1/platform"
	cerrors "common/proto/gen/common/errors"
	"context"

	"bff_bbs/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type AssetService struct {
	bffbbsv1.UnimplementedAssetServiceServer
	assetUsecase *usecase.AssetUsecase
}

func NewAssetService(assetUsecase *usecase.AssetUsecase) *AssetService {
	return &AssetService{assetUsecase: assetUsecase}
}

func (*AssetService) RegisterGrpc(*grpc.Server) {}

func (s *AssetService) RegisterHTTP(server *http.Server) {
	bffbbsv1.RegisterAssetServiceHTTPServer(server, s)
}

func (s *AssetService) PrepareMarkdownImageUpload(ctx context.Context, req *bffbbsv1.PrepareMarkdownImageUpload_Req) (*bffbbsv1.PrepareMarkdownImageUpload_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil || user.ID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_MARKDOWN_IMAGE_INVALID)
	}
	asset, err := s.assetUsecase.PrepareMarkdownImageUpload(ctx, &usecase.PrepareMarkdownImageUploadReq{
		UserID:   user.ID,
		Hash:     req.GetHash(),
		MimeType: req.GetMimeType(),
		Size:     req.GetSize(),
	})
	if err != nil {
		return nil, err
	}
	return &bffbbsv1.PrepareMarkdownImageUpload_Resp{AssetId: asset.AssetID, UploadUrl: asset.UploadURL, FormFields: asset.FormFields}, nil
}

func (s *AssetService) CompleteMarkdownImageUpload(ctx context.Context, req *bffbbsv1.CompleteMarkdownImageUpload_Req) (*bffbbsv1.CompleteMarkdownImageUpload_Resp, error) {
	user, ok := util.GetContextValue[*commonmodel.User](ctx, constant.CtxUserInfo)
	if !ok || user == nil || user.ID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_USER_TOKEN_REQUIRED)
	}
	if req == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_MARKDOWN_IMAGE_INVALID)
	}
	asset, err := s.assetUsecase.CompleteMarkdownImageUpload(ctx, &usecase.CompleteMarkdownImageUploadReq{Hash: req.GetHash()})
	if err != nil {
		return nil, err
	}
	return &bffbbsv1.CompleteMarkdownImageUpload_Resp{AssetId: asset.AssetID, ImageUrl: asset.ImageURL}, nil
}
