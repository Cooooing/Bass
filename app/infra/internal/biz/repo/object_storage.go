package repo

import (
	"common/api/gen/common"
	"context"
	"infra/internal/biz/model"
	"infra/internal/data/ent/gen"
)

type ObjectStorageRepo interface {
	// Save 保存 oss 对象信息
	Save(ctx context.Context, tx *gen.Client, o *model.ObjectStorage) (*model.ObjectStorage, error)

	// UpdateAudit 更新对象审核信息
	UpdateAudit(ctx context.Context, tx *gen.Client, u *model.ObjectStorage) error

	Delete(ctx context.Context, tx *gen.Client, o *model.ObjectStorage) (int, error)

	Exist(ctx context.Context, tx *gen.Client, req *ObjectStorageGetReq) (bool, error)
	GetList(ctx context.Context, tx *gen.Client, req *ObjectStorageGetReq) ([]*model.ObjectStorage, error)
	GetPage(ctx context.Context, tx *gen.Client, page *common.PageRequest, req *ObjectStorageGetReq) ([]*model.ObjectStorage, *common.PageReply, error)
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

	// UploadToken 获取上传token
	UploadToken(ctx context.Context, key string) (string, error)
	// Status 启用或禁用对象
	Status(ctx context.Context, key string, enable bool) error
}
