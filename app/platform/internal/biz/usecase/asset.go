package usecase

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"platform/internal/biz/model"
	"platform/internal/biz/repo"
	"platform/internal/config"
	"platform/internal/enum"
	"regexp"
	"strings"
	"time"
)

var sha256Pattern = regexp.MustCompile(`^[a-f0-9]{64}$`)

type AssetUsecase struct {
	assetRepo            repo.AssetRepo
	storageClient        repo.StorageClient
	assetUploadRateLimit repo.AssetUploadRateLimitCache
	rateLimitEnabled     bool
	rateLimitWindow      time.Duration
	rateLimitMaxCount    int64
}

func NewAssetUsecase(
	assetRepo repo.AssetRepo,
	storageClient repo.StorageClient,
	assetUploadRateLimit repo.AssetUploadRateLimitCache,
	conf *config.Bootstrap,
) *AssetUsecase {
	rateLimitEnabled := true
	rateLimitWindow := 5 * time.Minute
	rateLimitMaxCount := int64(5)
	if conf != nil && conf.GetPlatform() != nil {
		settings := conf.GetPlatform().GetAssetUploadRateLimit()
		if settings != nil {
			rateLimitEnabled = settings.GetEnable()
			if settings.GetWindow() != nil && settings.GetWindow().AsDuration() > 0 {
				rateLimitWindow = settings.GetWindow().AsDuration()
			}
			if settings.GetMaxCount() > 0 {
				rateLimitMaxCount = settings.GetMaxCount()
			}
		}
	}
	return &AssetUsecase{
		assetRepo:            assetRepo,
		storageClient:        storageClient,
		assetUploadRateLimit: assetUploadRateLimit,
		rateLimitEnabled:     rateLimitEnabled,
		rateLimitWindow:      rateLimitWindow,
		rateLimitMaxCount:    rateLimitMaxCount,
	}
}

type PrepareDirectAssetUploadReq struct {
	Hash       string
	MimeType   string
	Size       int64
	UploadByID int64
}

type PrepareDirectAssetUploadResp struct {
	Asset *model.Asset
	URL   string
	Form  map[string]string
}

// PrepareDirectUpload never creates an asset record. The record is created by
// the MinIO completion event after a successful object write. Existing content
// is returned immediately so identical bytes are uploaded only once.
func (u *AssetUsecase) PrepareDirectUpload(ctx context.Context, req *PrepareDirectAssetUploadReq) (*PrepareDirectAssetUploadResp, error) {
	if req == nil || strings.TrimSpace(req.MimeType) == "" || req.Size <= 0 || req.UploadByID <= 0 {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	hash, ok := normalizeSHA256(req.Hash)
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	asset, err := u.assetRepo.GetByHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	if asset != nil {
		if asset.Status != enum.AssetStatusAvailable {
			return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
		}
		return &PrepareDirectAssetUploadResp{Asset: asset}, nil
	}
	if err := u.consumeUploadAllowance(ctx, req.UploadByID); err != nil {
		return nil, err
	}
	prepared, err := u.storageClient.PrepareDirectUpload(ctx, &repo.DirectAssetUploadReq{
		Hash:       hash,
		MimeType:   strings.ToLower(strings.TrimSpace(req.MimeType)),
		Size:       req.Size,
		UploadByID: req.UploadByID,
	})
	if err != nil {
		return nil, err
	}
	return &PrepareDirectAssetUploadResp{URL: prepared.URL, Form: prepared.FormFields}, nil
}

// RecordMinioUpload creates the immutable asset fact from a MinIO-created
// event delivered through NATS. It performs StatObject only: application
// services never download the uploaded bytes during this flow.
func (u *AssetUsecase) RecordMinioUpload(ctx context.Context, objectKey string) (*model.Asset, error) {
	hash, ok := hashFromObjectKey(objectKey)
	if !ok {
		return nil, fmt.Errorf("%w: malformed object key", repo.ErrInvalidAssetUploadEvent)
	}
	object, err := u.storageClient.Stat(ctx, objectKey)
	if err != nil {
		return nil, err
	}
	expectedChecksum, err := sha256Base64(hash)
	if err != nil {
		return nil, err
	}
	if object.ChecksumSHA256 != expectedChecksum {
		return nil, fmt.Errorf("%w: object checksum does not match object key", repo.ErrInvalidAssetUploadEvent)
	}
	if object.UploadByID <= 0 {
		return nil, fmt.Errorf("%w: object uploader metadata is missing", repo.ErrInvalidAssetUploadEvent)
	}
	return u.assetRepo.CreateOrGet(ctx, &model.Asset{
		Provider:     object.Provider,
		Bucket:       object.Bucket,
		ObjectKey:    object.ObjectKey,
		Hash:         new(hash),
		UploadByID:   new(object.UploadByID),
		ProviderETag: stringPointer(object.ProviderETag),
		MimeType:     strings.ToLower(strings.TrimSpace(object.MimeType)),
		Size:         object.Size,
		Status:       enum.AssetStatusAvailable,
	})
}

// RecordAssetUploadEvent validates the event against the asset directory's
// bucket and records the completed object. Consumer ACK policy belongs to the
// infrastructure adapter, not to this business operation.
func (u *AssetUsecase) RecordAssetUploadEvent(
	ctx context.Context,
	event *repo.AssetUploadEvent,
) error {
	if event == nil || event.Bucket != u.StorageBucket() || !isAssetObjectKey(event.ObjectKey) {
		return repo.ErrInvalidAssetUploadEvent
	}
	_, err := u.RecordMinioUpload(ctx, event.ObjectKey)
	return err
}

func (u *AssetUsecase) GetAvailableByHash(ctx context.Context, hash string) (*model.Asset, error) {
	canonicalHash, ok := normalizeSHA256(hash)
	if !ok {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	asset, err := u.assetRepo.GetByHash(ctx, canonicalHash)
	if err != nil {
		return nil, err
	}
	if asset == nil {
		// The browser can finish its POST slightly before MinIO delivers the
		// callback. Callers may retry this precise not-found condition briefly.
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_NOT_FOUND)
	}
	if asset.Status != enum.AssetStatusAvailable {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	return asset, nil
}

func (u *AssetUsecase) ResolvePublic(ctx context.Context, ids []int64) (map[int64]string, error) {
	assets, err := u.assetRepo.Map(ctx, ids)
	if err != nil {
		return nil, err
	}
	urls := make(map[int64]string, len(assets))
	for id, asset := range assets {
		if asset.Status == enum.AssetStatusAvailable {
			urls[id] = u.storageClient.PublicURL(asset.ObjectKey)
		}
	}
	return urls, nil
}

func (u *AssetUsecase) ValidateAvailable(ctx context.Context, assetID int64) error {
	asset, err := u.assetRepo.Get(ctx, assetID)
	if err != nil {
		return err
	}
	if asset == nil || asset.Status != enum.AssetStatusAvailable {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_FORBIDDEN)
	}
	return nil
}

// Block changes only the global availability fact. It deliberately does not
// delete or rewrite domain references, because those relationships belong to
// the services that own the corresponding business objects.
func (u *AssetUsecase) Block(ctx context.Context, assetID int64, reason string, operatorUserID int64) error {
	if assetID <= 0 || operatorUserID <= 0 {
		return apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	return u.assetRepo.Block(ctx, assetID, strings.TrimSpace(reason), operatorUserID)
}

// StorageBucket returns the single bucket governed by this asset directory.
// The NATS consumer uses it to reject events emitted by unrelated buckets
// before it asks the provider for object metadata.
func (u *AssetUsecase) StorageBucket() string {
	return u.storageClient.Bucket()
}

func (u *AssetUsecase) PublicURL(asset *model.Asset) string {
	if asset == nil || asset.Status != enum.AssetStatusAvailable {
		return ""
	}
	return u.storageClient.PublicURL(asset.ObjectKey)
}

func (u *AssetUsecase) consumeUploadAllowance(ctx context.Context, uploadByID int64) error {
	if !u.rateLimitEnabled {
		return nil
	}
	if u.assetUploadRateLimit == nil {
		return fmt.Errorf("asset upload rate limiter is required")
	}
	state, err := u.assetUploadRateLimit.Allow(ctx, uploadByID, u.rateLimitWindow, u.rateLimitMaxCount)
	if err != nil {
		return err
	}
	if state == nil || !state.Allowed {
		retryAfterSeconds := int64(0)
		if state != nil && state.RetryAfter > 0 {
			retryAfterSeconds = int64((state.RetryAfter + time.Second - 1) / time.Second)
		}
		return apperror.New(
			cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_TOO_MANY_ReqS,
			apperror.WithData(&cerrors.RetryAfterErrorData{RetryAfterSeconds: retryAfterSeconds}),
		)
	}
	return nil
}

func normalizeSHA256(value string) (string, bool) {
	hash := strings.ToLower(strings.TrimSpace(value))
	return hash, sha256Pattern.MatchString(hash)
}

func hashFromObjectKey(objectKey string) (string, bool) {
	const prefix = "assets/sha256/"
	hash, ok := strings.CutPrefix(objectKey, prefix)
	return hash, ok && sha256Pattern.MatchString(hash)
}

func isAssetObjectKey(objectKey string) bool {
	_, ok := hashFromObjectKey(objectKey)
	return ok
}

func sha256Base64(hash string) (string, error) {
	decoded, err := hex.DecodeString(hash)
	if err != nil || len(decoded) != 32 {
		return "", fmt.Errorf("invalid SHA-256")
	}
	return base64.StdEncoding.EncodeToString(decoded), nil
}

func stringPointer(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
