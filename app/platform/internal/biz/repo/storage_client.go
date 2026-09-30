package repo

import (
	"context"
	"platform/internal/enum"
)

// StorageClient hides provider-specific object-storage APIs. Both the asset
// directory and platform-owned data sets use the same physical storage while
// retaining separate business models above this boundary.
type StorageClient interface {
	Name() string
	Bucket() string
	PrepareDirectUpload(ctx context.Context, req *DirectAssetUploadReq) (*DirectAssetUploadResp, error)
	Stat(ctx context.Context, objectKey string) (*AssetObject, error)
	PublicURL(objectKey string) string
	Get(ctx context.Context, objectKey string) (*StoredObject, error)
	Put(ctx context.Context, req *PutStoredObjectReq) (*StoredObject, error)
}

type DirectAssetUploadReq struct {
	Hash       string
	MimeType   string
	Size       int64
	UploadByID int64
}

type DirectAssetUploadResp struct {
	URL        string
	FormFields map[string]string
}

type AssetObject struct {
	Provider       enum.AssetProvider
	Bucket         string
	ObjectKey      string
	MimeType       string
	Size           int64
	ProviderETag   string
	ChecksumSHA256 string
	UploadByID     int64
}

type StoredObject struct {
	ObjectKey string
	MimeType  string
	Size      int64
	Content   []byte
}

type PutStoredObjectReq struct {
	ObjectKey string
	MimeType  string
	Content   []byte
}
