package repo

import (
	"common/proto/gen/common"
	"context"
	"platform/internal/biz/model"
)

type ObjectStorageRepo interface {
	Save(ctx context.Context, req *ObjectStorageSaveReq) (*ObjectStorageSaveResponse, error)
	UpdateAudit(ctx context.Context, req *ObjectStorageUpdateAuditReq) (*ObjectStorageUpdateAuditResponse, error)
	Delete(ctx context.Context, req *ObjectStorageDeleteReq) (*ObjectStorageDeleteResponse, error)

	Exist(ctx context.Context, req *ObjectStorageGetReq) (*ObjectStorageExistResponse, error)
	Get(ctx context.Context, req *ObjectStorageGetReq) (*ObjectStorageGetResponse, error)
	List(ctx context.Context, req *ObjectStorageGetReq) (*ObjectStorageListResponse, error)
	Map(ctx context.Context, req *ObjectStorageGetReq) (*ObjectStorageMapResponse, error)
	Count(ctx context.Context, req *ObjectStorageGetReq) (*ObjectStorageCountResponse, error)
	Page(ctx context.Context, req *ObjectStoragePageReq) (*ObjectStoragePageResponse, error)
}

type ObjectStorageGetReq struct {
	ID            *int64
	IDs           []int64
	Provider      *string
	Bucket        *string
	Key           *string
	MimeType      *string
	Size          *common.Int64Range
	Blocked       *bool
	BlockedByName *string
}

type ObjectStorageSaveReq struct {
	Row *model.ObjectStorage
}

type ObjectStorageSaveResponse struct {
	Row *model.ObjectStorage
}

type ObjectStorageUpdateAuditReq struct {
	Row *model.ObjectStorage
}

type ObjectStorageUpdateAuditResponse struct{}

type ObjectStorageDeleteReq struct {
	Row *model.ObjectStorage
}

type ObjectStorageDeleteResponse struct {
	Count int
}

type ObjectStorageExistResponse struct {
	Exists bool
}

type ObjectStorageGetResponse struct {
	Row *model.ObjectStorage
}

type ObjectStorageListResponse struct {
	Rows []*model.ObjectStorage
}

type ObjectStorageMapResponse struct {
	Rows map[int64]*model.ObjectStorage
}

type ObjectStorageCountResponse struct {
	Count int
}

type ObjectStoragePageReq struct {
	Page *common.PageRequest
	ObjectStorageGetReq
}

type ObjectStoragePageResponse struct {
	Rows []*model.ObjectStorage
	Page *common.PageResponse
}

type ObjectStorageClient interface {
	Name() string
	UploadToken(ctx context.Context, req *ObjectStorageUploadTokenReq) (*ObjectStorageUploadTokenResponse, error)
	Status(ctx context.Context, req *ObjectStorageStatusReq) (*ObjectStorageStatusResponse, error)
}

type ObjectStorageUploadTokenReq struct {
	Key          string
	UploaderID   int64
	UploaderName string
}

type ObjectStorageUploadTokenResponse struct {
	Token string
}

type ObjectStorageStatusReq struct {
	Key    string
	Enable bool
}

type ObjectStorageStatusResponse struct{}
