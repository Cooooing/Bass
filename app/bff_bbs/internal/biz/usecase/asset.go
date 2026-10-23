package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"regexp"
	"strings"

	"bff_bbs/internal/biz/repo"
)

type AssetUsecase struct {
	assetClient      repo.AssetClient
	assetHashPattern *regexp.Regexp
	maxImageSize     int64
}

func NewAssetUsecase(assetClient repo.AssetClient) *AssetUsecase {
	return &AssetUsecase{
		assetClient:      assetClient,
		assetHashPattern: regexp.MustCompile(`^[a-f0-9]{64}$`),
		maxImageSize:     5 * 1024 * 1024,
	}
}

type PrepareMarkdownImageUploadReq struct {
	UserID   int64
	Hash     string
	MimeType string
	Size     int64
}

type PrepareMarkdownImageUploadResp struct {
	AssetID    *int64
	UploadURL  string
	FormFields map[string]string
}

func (u *AssetUsecase) PrepareMarkdownImageUpload(ctx context.Context, req *PrepareMarkdownImageUploadReq) (*PrepareMarkdownImageUploadResp, error) {
	if req == nil || req.UserID <= 0 || !u.validMarkdownImage(req.Hash, req.MimeType) || req.Size <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_MARKDOWN_IMAGE_INVALID)
	}
	if req.Size > u.maxImageSize {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_MARKDOWN_IMAGE_TOO_LARGE)
	}
	asset, err := u.assetClient.PrepareDirectUpload(ctx, &repo.PrepareDirectAssetUploadReq{
		Hash:       strings.ToLower(strings.TrimSpace(req.Hash)),
		MimeType:   strings.ToLower(strings.TrimSpace(req.MimeType)),
		Size:       req.Size,
		UploadByID: req.UserID,
	})
	if err != nil {
		return nil, err
	}
	return &PrepareMarkdownImageUploadResp{AssetID: asset.AssetID, UploadURL: asset.UploadURL, FormFields: asset.FormFields}, nil
}

type CompleteMarkdownImageUploadReq struct {
	Hash string
}

type CompleteMarkdownImageUploadResp struct {
	AssetID  int64
	ImageURL string
}

func (u *AssetUsecase) CompleteMarkdownImageUpload(ctx context.Context, req *CompleteMarkdownImageUploadReq) (*CompleteMarkdownImageUploadResp, error) {
	if req == nil || !u.assetHashPattern.MatchString(strings.ToLower(strings.TrimSpace(req.Hash))) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_MARKDOWN_IMAGE_INVALID)
	}
	asset, err := u.assetClient.GetAvailableByHash(ctx, strings.ToLower(strings.TrimSpace(req.Hash)))
	if err != nil {
		return nil, err
	}
	if asset == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if !u.validMarkdownImage(req.Hash, asset.MimeType) || asset.Size <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_MARKDOWN_IMAGE_INVALID)
	}
	if asset.Size > u.maxImageSize {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_MARKDOWN_IMAGE_TOO_LARGE)
	}
	return &CompleteMarkdownImageUploadResp{AssetID: asset.ID, ImageURL: asset.URL}, nil
}

func (u *AssetUsecase) validMarkdownImage(hash, mimeType string) bool {
	if !u.assetHashPattern.MatchString(strings.ToLower(strings.TrimSpace(hash))) {
		return false
	}
	switch strings.ToLower(strings.TrimSpace(mimeType)) {
	case "image/gif", "image/jpeg", "image/png", "image/webp":
		return true
	}
	return false
}
