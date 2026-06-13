package repo

import (
	"common/proto/gen/common"
	"context"
	"platform/internal/biz/model"
)

type ObjectStorageRepo interface {
	Save(ctx context.Context, o *model.ObjectStorage) (*model.ObjectStorage, error)

	UpdateAudit(ctx context.Context, u *model.ObjectStorage) error

	Delete(ctx context.Context, o *model.ObjectStorage) (int, error)

	Exist(ctx context.Context, req *ObjectStorageGetReq) (bool, error)
	Get(ctx context.Context, req *ObjectStorageGetReq) (*model.ObjectStorage, error)
	List(ctx context.Context, req *ObjectStorageGetReq) ([]*model.ObjectStorage, error)
	Map(ctx context.Context, req *ObjectStorageGetReq) (map[int64]*model.ObjectStorage, error)
	Count(ctx context.Context, req *ObjectStorageGetReq) (int, error)
	Page(ctx context.Context, page *common.PageRequest, req *ObjectStorageGetReq) ([]*model.ObjectStorage, *common.PageReply, error)
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

type ObjectStorageClient interface {
	Name() string
	UploadToken(ctx context.Context, key string, uploaderID int64, uploaderName string) (string, error)
	Status(ctx context.Context, key string, enable bool) error
}
