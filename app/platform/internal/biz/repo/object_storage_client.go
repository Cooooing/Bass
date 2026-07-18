package repo

import (
	"context"
	"io"
	"platform/internal/enum"
)

type ObjectStorageClient interface {
	Name() string
	CreateBucket(ctx context.Context, bucket string) error
	DeleteBucket(ctx context.Context, bucket string) error
	Upload(ctx context.Context, req *ObjectStorageUploadReq) (*ObjectStorageUploadResp, error)
	StreamUpload(ctx context.Context, req *ObjectStorageStreamUploadReq) (*ObjectStorageStreamUploadResp, error)
	Download(ctx context.Context, key string) (*ObjectStorageDownloadResp, error)
	StreamDownload(ctx context.Context, key string) (*ObjectStorageStreamDownloadResp, error)
	UploadToken(ctx context.Context, req *ObjectStorageUploadTokenReq) (string, error)
	Status(ctx context.Context, req *ObjectStorageStatusReq) error
}

type ObjectStorageUploadReq struct {
	Key        string
	FileName   string
	MimeType   string
	Content    []byte
	UploaderID int64
}

type ObjectStorageUploadResp struct {
	Provider enum.ObjectStorageProvider
	Bucket   string
	Key      string
	MimeType string
	Size     int64
	Hash     string
}

type ObjectStorageStreamUploadReq struct {
	Key        string
	FileName   string
	MimeType   string
	Size       int64
	Body       io.Reader
	UploaderID int64
}

type ObjectStorageStreamUploadResp struct {
	Provider enum.ObjectStorageProvider
	Bucket   string
	Key      string
	MimeType string
	Size     int64
	Hash     string
}

type ObjectStorageDownloadResp struct {
	Key      string
	MimeType string
	Size     int64
	Content  []byte
}

type ObjectStorageStreamDownloadResp struct {
	Key      string
	MimeType string
	Size     int64
	Body     io.ReadCloser
}

type ObjectStorageUploadTokenReq struct {
	Key        string
	UploaderID int64
}

type ObjectStorageStatusReq struct {
	Key    string
	Enable bool
}
