package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"regexp"
	"strings"

	"bff_bbs/internal/biz/repo"
	"bff_bbs/internal/config"
)

type AssetUsecase struct {
	assetClient           repo.AssetClient
	assetHashPattern      *regexp.Regexp
	maxMarkdownImageBytes int64
	maxBadgeImageBytes    int64
	maxBadgeFontBytes     int64
}

func NewAssetUsecase(assetClient repo.AssetClient, conf *config.Bootstrap) *AssetUsecase {
	const defaultMaxBytes = 3 * 1024 * 1024
	maxBadgeImageBytes, maxBadgeFontBytes := int64(defaultMaxBytes), int64(defaultMaxBytes)
	if conf != nil && conf.GetBusiness() != nil && conf.GetBusiness().GetBadgeMakerUpload() != nil {
		settings := conf.GetBusiness().GetBadgeMakerUpload()
		if settings.GetMaxImageBytes() > 0 {
			maxBadgeImageBytes = settings.GetMaxImageBytes()
		}
		if settings.GetMaxFontBytes() > 0 {
			maxBadgeFontBytes = settings.GetMaxFontBytes()
		}
	}
	return &AssetUsecase{
		assetClient:           assetClient,
		assetHashPattern:      regexp.MustCompile(`^[a-f0-9]{64}$`),
		maxMarkdownImageBytes: 5 * 1024 * 1024,
		maxBadgeImageBytes:    maxBadgeImageBytes,
		maxBadgeFontBytes:     maxBadgeFontBytes,
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
	if req.Size > u.maxMarkdownImageBytes {
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

// PrepareBadgeImageUploadReq 描述 Maker 徽章图片的上传参数。
type PrepareBadgeImageUploadReq struct {
	UserID   int64
	Hash     string
	MimeType string
	Size     int64
}

// PrepareBadgeImageUploadResp 返回图片资源或直传凭证。
type PrepareBadgeImageUploadResp = PrepareMarkdownImageUploadResp

func (u *AssetUsecase) PrepareBadgeImageUpload(ctx context.Context, req *PrepareBadgeImageUploadReq) (*PrepareBadgeImageUploadResp, error) {
	if req == nil || req.UserID <= 0 || !u.validMarkdownImage(req.Hash, req.MimeType) || req.Size <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_BADGE_IMAGE_INVALID)
	}
	if req.Size > u.maxBadgeImageBytes {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_BADGE_IMAGE_TOO_LARGE)
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
	return &PrepareBadgeImageUploadResp{
		AssetID:    asset.AssetID,
		UploadURL:  asset.UploadURL,
		FormFields: asset.FormFields,
	}, nil
}

// CompleteBadgeImageUploadReq identifies a completed Maker image upload.
type CompleteBadgeImageUploadReq struct {
	Hash string
}

// CompleteBadgeImageUploadResp returns the public image URL used by RenderSpec.
type CompleteBadgeImageUploadResp = CompleteMarkdownImageUploadResp

func (u *AssetUsecase) CompleteBadgeImageUpload(ctx context.Context, req *CompleteBadgeImageUploadReq) (*CompleteBadgeImageUploadResp, error) {
	if req == nil || !u.assetHashPattern.MatchString(strings.ToLower(strings.TrimSpace(req.Hash))) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_BADGE_IMAGE_INVALID)
	}
	asset, err := u.assetClient.GetAvailableByHash(ctx, strings.ToLower(strings.TrimSpace(req.Hash)))
	if err != nil {
		return nil, err
	}
	if asset == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if !u.validMarkdownImage(req.Hash, asset.MimeType) || asset.Size <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_BADGE_IMAGE_INVALID)
	}
	if asset.Size > u.maxBadgeImageBytes {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_BADGE_IMAGE_TOO_LARGE)
	}
	return &CompleteBadgeImageUploadResp{
		AssetID:  asset.ID,
		ImageURL: asset.URL,
	}, nil
}

// PrepareBadgeFontUploadReq 描述 Maker 自定义字体的上传参数。
type PrepareBadgeFontUploadReq struct {
	UserID   int64
	Hash     string
	MimeType string
	Size     int64
}

// PrepareBadgeFontUploadResp 返回字体资源或直传凭证。
type PrepareBadgeFontUploadResp = PrepareMarkdownImageUploadResp

func (u *AssetUsecase) PrepareBadgeFontUpload(ctx context.Context, req *PrepareBadgeFontUploadReq) (*PrepareBadgeFontUploadResp, error) {
	if req == nil || req.UserID <= 0 || !u.validBadgeFont(req.Hash, req.MimeType) || req.Size <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_BADGE_FONT_INVALID)
	}
	if req.Size > u.maxBadgeFontBytes {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_BADGE_FONT_TOO_LARGE)
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
	return &PrepareBadgeFontUploadResp{
		AssetID:    asset.AssetID,
		UploadURL:  asset.UploadURL,
		FormFields: asset.FormFields,
	}, nil
}

// CompleteBadgeFontUploadReq identifies a completed Maker font upload.
type CompleteBadgeFontUploadReq struct {
	Hash string
}

// CompleteBadgeFontUploadResp returns the public font URL used by RenderSpec.
type CompleteBadgeFontUploadResp struct {
	AssetID int64
	FontURL string
}

func (u *AssetUsecase) CompleteBadgeFontUpload(ctx context.Context, req *CompleteBadgeFontUploadReq) (*CompleteBadgeFontUploadResp, error) {
	if req == nil || !u.assetHashPattern.MatchString(strings.ToLower(strings.TrimSpace(req.Hash))) {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_BADGE_FONT_INVALID)
	}
	asset, err := u.assetClient.GetAvailableByHash(ctx, strings.ToLower(strings.TrimSpace(req.Hash)))
	if err != nil {
		return nil, err
	}
	if asset == nil {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if !u.validBadgeFont(req.Hash, asset.MimeType) || asset.Size <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_BADGE_FONT_INVALID)
	}
	if asset.Size > u.maxBadgeFontBytes {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_PLATFORM_BADGE_FONT_TOO_LARGE)
	}
	return &CompleteBadgeFontUploadResp{
		AssetID: asset.ID,
		FontURL: asset.URL,
	}, nil
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
	if asset.Size > u.maxMarkdownImageBytes {
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

func (u *AssetUsecase) validBadgeFont(hash, mimeType string) bool {
	if !u.assetHashPattern.MatchString(strings.ToLower(strings.TrimSpace(hash))) {
		return false
	}
	switch strings.ToLower(strings.TrimSpace(mimeType)) {
	case "font/ttf", "font/otf", "font/woff", "font/woff2", "application/font-sfnt", "application/x-font-ttf", "application/x-font-opentype":
		return true
	}
	return false
}
