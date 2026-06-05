package minio

import (
	"common/pkg/constant"
	"context"
	"notify/internal/biz/model"
	"notify/internal/data/gen"
)

type Minio struct {
}

func NewMinio() *Minio {
	return &Minio{}
}

func (m *Minio) Name() string {
	return constant.Minio.String()
}

func (m *Minio) Save(ctx context.Context, tx *gen.Client, o *model.ObjectStorage) (*model.ObjectStorage, error) {
	panic("implement me")
}

func (m *Minio) UploadToken(ctx context.Context, key string, uploaderID int64, uploaderName string) (string, error) {
	panic("implement me")
}

func (m *Minio) Status(ctx context.Context, key string, enable bool) error {
	panic("implement me")
}
