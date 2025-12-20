package repo

import (
	"context"
	"user/internal/biz/model"
	"user/internal/data/ent/gen"
)

type ObjectStorageRepo interface {
	Name() string

	Save(ctx context.Context, tx *gen.Client, o *model.ObjectStorage) (*model.ObjectStorage, error)
	UploadToken(key string) string
	//UpdateAuditStatus(ctx context.Context, key string, status bool) error
}
