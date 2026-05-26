package repo

import (
	"common/api/gen/common"
	"context"
	"notify/internal/biz/model"
)

type ObjectStorageRepo interface {
	Save(ctx context.Context, o *model.ObjectStorage) (*model.ObjectStorage, error)

	UpdateAudit(ctx context.Context, u *model.ObjectStorage) error

	Delete(ctx context.Context, o *model.ObjectStorage) (int, error)

	Exist(ctx context.Context, req *ObjectStorageGetReq) (bool, error)
	GetList(ctx context.Context, req *ObjectStorageGetReq) ([]*model.ObjectStorage, error)
	GetPage(ctx context.Context, page *common.PageRequest, req *ObjectStorageGetReq) ([]*model.ObjectStorage, *common.PageReply, error)
}

type ObjectStorageGetReq struct {
	Provider      *string
	Bucket        *string
	Key           *string
	MimeType      *string
	Size          *common.Int64Range
	Blocked       *bool
	BlockedByName *string
}

type ObjectStorageProvider interface {
	Name() string
	UploadToken(ctx context.Context, key string) (string, error)
	Status(ctx context.Context, key string, enable bool) error
}
