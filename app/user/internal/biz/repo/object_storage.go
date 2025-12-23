package repo

import (
	cv1 "common/api/common/v1"
	"context"
	"user/internal/biz/model"
	"user/internal/data/ent/gen"
)

type ObjectStorageRepo interface {
	// Save 保存 oss 对象信息
	Save(ctx context.Context, tx *gen.Client, o *model.ObjectStorage) (*model.ObjectStorage, error)

	// UpdateAudit 更新对象审核信息
	UpdateAudit(ctx context.Context, tx *gen.Client, u *model.ObjectStorage) error

	Delete(ctx context.Context, tx *gen.Client, o *model.ObjectStorage) (int, error)

	Exist(ctx context.Context, tx *gen.Client, req *ObjectStorageGetReq) (bool, error)
	GetList(ctx context.Context, tx *gen.Client, req *ObjectStorageGetReq) ([]*model.ObjectStorage, error)
	GetPage(ctx context.Context, tx *gen.Client, page *cv1.PageRequest, req *ObjectStorageGetReq) ([]*model.ObjectStorage, *cv1.PageReply, error)
}

type ObjectStorageGetReq struct {
	Provider      *string
	Bucket        *string
	Key           *string
	MimeType      *string
	Size          *cv1.Int64Range
	Blocked       *bool
	BlockedByName *string
}

type ObjectStorageProvider interface {
	Name() string

	// UploadToken 获取上传token
	UploadToken(key string) string
	// Status 启用或禁用对象
	Status(ctx context.Context, key string, enable bool) error
	// VerifyCallback 验证回调来源
	VerifyCallback(ctx context.Context) error
}
