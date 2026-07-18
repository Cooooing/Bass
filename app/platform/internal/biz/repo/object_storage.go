package repo

import (
	"common/proto/gen/common"
	"context"
	"platform/internal/biz/model"
	"platform/internal/enum"
)

type ObjectStorageRepo interface {
	Save(ctx context.Context, row *model.ObjectStorage) (*model.ObjectStorage, error)
	UpdateAudit(ctx context.Context, row *model.ObjectStorage) error
	Delete(ctx context.Context, row *model.ObjectStorage) (int, error)

	Exist(ctx context.Context, req *ObjectStorageGetReq) (bool, error)
	Get(ctx context.Context, req *ObjectStorageGetReq) (*model.ObjectStorage, error)
	List(ctx context.Context, req *ObjectStorageGetReq) ([]*model.ObjectStorage, error)
	Map(ctx context.Context, req *ObjectStorageGetReq) (map[int64]*model.ObjectStorage, error)
	Count(ctx context.Context, req *ObjectStorageGetReq) (int, error)
	Page(ctx context.Context, req *ObjectStoragePageReq) (*ObjectStoragePageResp, error)
}

type ObjectStorageGetReq struct {
	ID       *int64
	IDs      []int64
	Provider *enum.ObjectStorageProvider
	Bucket   *string
	Key      *string
	MimeType *string
	Size     *common.Int64Range
	Blocked  *bool
}

type ObjectStoragePageReq struct {
	Page *common.PageReq
	ObjectStorageGetReq
}

type ObjectStoragePageResp struct {
	Rows []*model.ObjectStorage
	Page *common.PageResp
}
